// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
	"github.com/lei006/gomath/la"
)

func TestHL01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("HL01. Van der Pol with 1 cycle/stationary point")

	// constants
	T := 6.6632868593231301896996820305
	A := 2.00861986087484313650940188
	y := la.NewVectorSlice([]float64{A, 0})
	xf := T
	dx := 0.1

	// function
	fcn := func(f la.Vector, dx, x float64, y la.Vector) {
		f[0] = y[1]
		f[1] = (1.0-y[0]*y[0])*y[1] - y[0]
	}

	// solve
	yf := y.GetCopy()
	atol, rtol := 1e-5, 1e-5
	numJac, fixedStep, saveStep, saveCont := false, false, true, false
	stat, _ := Solve("dopri5", fcn, nil, yf, xf, dx, atol, rtol, numJac, fixedStep, saveStep, saveCont)

	// results
	io.Pf("yf = %v\n", yf)
	stat.Print(false)

	// check
	chk.AnaNum(tst, "dopri5: y0", 1e-4, yf[0], y[0], chk.Verbose)
	chk.AnaNum(tst, "dopri5: y1", 1e-4, yf[1], y[1], chk.Verbose)

	// using simple version
	yf2 := y.GetCopy()
	Dopri5simple(fcn, yf2, xf, atol)
	chk.Array(tst, "dopri5: yf2", 1e-17, yf, yf2)

	// dopri8
	yf3 := y.GetCopy()
	Dopri8simple(fcn, yf3, xf, atol)
	chk.AnaNum(tst, "dopri8: y0", 1e-7, yf3[0], y[0], chk.Verbose)
	chk.AnaNum(tst, "dopri8: y1", 1e-5, yf3[1], y[1], chk.Verbose)

	// radau5
	yf4 := y.GetCopy()
	Radau5simple(fcn, nil, yf4, xf, atol)
	chk.AnaNum(tst, "radau5: y0", 1e-6, yf4[0], y[0], chk.Verbose)
	chk.AnaNum(tst, "radau5: y1", 1e-5, yf4[1], y[1], chk.Verbose)
}
