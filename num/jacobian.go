// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package num

import (
	"math"
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/fun"
	"github.com/lei006/gomath/la"
	"github.com/lei006/gomath/utl"
)

// Jacobian computes Jacobian (sparse) matrix
//
//	    Calculates (with N=n-1):
//	        df0dx0, df0dx1, df0dx2, ... df0dxN
//	        df1dx0, df1dx1, df1dx2, ... df1dxN
//	             . . . . . . . . . . . . .
//	        dfNdx0, dfNdx1, dfNdx2, ... dfNdxN
//	INPUT:
//	    ffcn : f(x) function
//	    x    : station where dfdx has to be calculated
//	    fx   : f @ x
//	    w    : workspace with size == n == len(x)
//	RETURNS:
//	    J : dfdx @ x [must be pre-allocated]
func Jacobian(J *la.Triplet, ffcn fun.Vv, x, fx, w []float64) {
	ndim := len(x)
	start, endp1 := 0, ndim
	if J.Max() == 0 {
		J.Init(ndim, ndim, ndim*ndim)
	}
	J.Start()
	var df float64
	for col := 0; col < ndim; col++ {
		xsafe := x[col]
		delta := math.Sqrt(MACHEPS * utl.Max(1e-5, math.Abs(xsafe)))
		x[col] = xsafe + delta
		ffcn(w, x) // w := f(x+δx[col])
		for row := start; row < endp1; row++ {
			df = w[row] - fx[row]
			J.Put(row, col, df/delta)
		}
		x[col] = xsafe
	}
}

// CompareJac compares Jacobian matrix (e.g. for testing)
func CompareJac(tst *testing.T, ffcn fun.Vv, Jfcn fun.Tv, x []float64, tol float64) {

	// numerical
	n := len(x)
	fx := make([]float64, n)
	w := make([]float64, n) // workspace
	ffcn(fx, x)
	var Jnum la.Triplet
	Jnum.Init(n, n, n*n)
	Jacobian(&Jnum, ffcn, x, fx, w)

	// analytical
	var Jana la.Triplet
	Jana.Init(n, n, n*n)
	Jfcn(&Jana, x)

	// compare
	chk.Deep2(tst, "J", tol, Jana.ToDense().GetDeep2(), Jnum.ToDense().GetDeep2())
}

// CompareJacDense compares Jacobian matrix (e.g. for testing) in dense format
func CompareJacDense(tst *testing.T, ffcn fun.Vv, Jfcn fun.Mv, x []float64, tol float64) {

	// numerical
	n := len(x)
	fx := make([]float64, n)
	w := make([]float64, n) // workspace
	ffcn(fx, x)
	var Jnum la.Triplet
	Jnum.Init(n, n, n*n)
	Jacobian(&Jnum, ffcn, x, fx, w)

	// analytical
	Jana := la.NewMatrix(n, n)
	Jfcn(Jana, x)

	// compare
	chk.Deep2(tst, "J", tol, Jana.GetDeep2(), Jnum.ToDense().GetDeep2())
}
