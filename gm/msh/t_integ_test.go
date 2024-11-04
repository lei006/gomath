// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package msh

import (
	"math"
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
	"github.com/lei006/gomath/la"
)

func TestInteg01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Integ01. integration over rotated square")

	// vertices (diamond shape)
	X := la.NewMatrixDeep2([][]float64{
		{0.0, +0.0},
		{1.0, -1.0},
		{2.0, +0.0},
		{1.0, +1.0},
	})

	// allocate cell integrator with default integration points
	o := NewIntegrator(TypeQua4, nil, "")
	chk.Int(tst, "Nverts", o.Nverts, 4)
	chk.Int(tst, "Ndim", o.Ndim, 2)
	chk.Int(tst, "Npts", o.Npts, 4)

	// integrand function
	fcn := func(x la.Vector) (f float64) {
		f = x[0]*x[0] + x[1]*x[1]
		return
	}

	// perform integration
	res := o.IntegrateSv(X, fcn)
	io.Pforan("1: res = %v\n", res)
	chk.Float64(tst, "∫(x²+y²)dxdy (default)", 1e-15, res, 8.0/3.0)

	// reset integration points
	o.ResetP(nil, "legendre_9")

	// perform integration again
	res = o.IntegrateSv(X, fcn)
	io.Pforan("2: res = %v\n", res)
	chk.Float64(tst, "∫(x²+y²)dxdy (legendre 9)", 1e-15, res, 8.0/3.0)

	// reset integration points
	o.ResetP(nil, "wilson5corner_5")

	// perform integration again
	res = o.IntegrateSv(X, fcn)
	io.Pforan("3: res = %v\n", res)
	chk.Float64(tst, "∫(x²+y²)dxdy (wilson5corner)", 1e-15, res, 8.0/3.0)
}

func TestInteg02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Integ02. integration over trapezium")

	// vertices (trapezium)
	a, b, h := 3.0, 0.5, 1.0
	X := la.NewMatrixDeep2([][]float64{
		{-a / 2.0, -h / 2.0},
		{+a / 2.0, -h / 2.0},
		{+b / 2.0, +h / 2.0},
		{-b / 2.0, +h / 2.0},
	})

	// allocate cell integrator with default integration points
	o := NewIntegrator(TypeQua4, nil, "legendre_4")

	// integrand function for second moment of inertia about x-axis: Ix
	fcnIx := func(x la.Vector) (f float64) {
		f = x[1] * x[1]
		return
	}

	// integrand function for second moment of inertia about y-axis: Iy
	fcnIy := func(x la.Vector) (f float64) {
		f = x[0] * x[0]
		return
	}

	// integrand function for second moment of inertia about the origin: I0
	fcnI0 := func(x la.Vector) (f float64) {
		f = (x[0]*x[0] + x[1]*x[1])
		return
	}

	// analytical solutions
	anaIx := (a + b) * math.Pow(h, 3) / 24.0
	anaIy := h * (math.Pow(a, 4) - math.Pow(b, 4)) / (48.0 * (a - b))
	anaI0 := anaIx + anaIy

	// compute Ix
	Ix := o.IntegrateSv(X, fcnIx)
	io.Pforan("Ix = %v\n", Ix)
	chk.Float64(tst, "Ix", 1e-15, Ix, anaIx)

	// compute Iy
	Iy := o.IntegrateSv(X, fcnIy)
	io.Pforan("Iy = %v\n", Iy)
	chk.Float64(tst, "Iy", 1e-15, Iy, anaIy)

	// compute I0
	I0 := o.IntegrateSv(X, fcnI0)
	io.Pforan("I0 = %v\n", I0)
	chk.Float64(tst, "I0", 1e-15, I0, anaI0)
}

func TestInteg03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Integ03. 2nd mom inertia: quarter of circle")

	// integrand function for second moment of inertia about x-axis: Ix
	fcnIx := func(x la.Vector) (f float64) {
		f = x[1] * x[1]
		return
	}

	// constants
	anaIx := math.Pi / 16.0 // analytical solution
	r, R := 0.0, 1.0
	nr, na := 5, 5

	// run for many quads
	//tols := []float64{0.0014, 1e-6, 1e-6, 1e-7, 1e-7, 1e-10} // 11 x 11
	tols := []float64{0.007, 1e-5, 1e-5, 1e-5, 1e-5, 1e-8} // 5 x 5
	ctypes := []int{TypeQua4, TypeQua8, TypeQua9, TypeQua12, TypeQua16, TypeQua17}
	for i, ctype := range ctypes {
		mesh := GenRing2d(ctype, nr, na, r, R, math.Pi/2.0)

		// allocate cell integrator with default integration points
		o := NewMeshIntegrator(mesh, 1)

		// compute Ix
		Ix := o.IntegrateSv(0, fcnIx)
		typekey := TypeIndexToKey[ctype]
		io.Pf("%s : Ix = %v  error = %v\n", typekey, Ix, math.Abs(Ix-anaIx))
		chk.Float64(tst, "Ix", tols[i], Ix, anaIx)
	}
}

func TestInteg04(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Integ04. 2nd mom inertia: ring")

	// integrand function for second moment of inertia about x-axis: Ix
	fcnIx := func(x la.Vector) (f float64) {
		f = x[1] * x[1]
		return
	}

	// constants
	r, R := 1.0, 3.0
	nr, na := 4, 13
	anaIx := math.Pi * (math.Pow(R, 4) - math.Pow(r, 4)) / 4.0

	// run for many quads
	//tols := []float64{2.0, 0.003, 0.003, 0.004, 0.004, 1e-6} // 5 x 21
	tols := []float64{5, 0.02, 0.02, 0.003, 0.003, 1e-5} // 4 x 13
	ctypes := []int{TypeQua4, TypeQua8, TypeQua9, TypeQua12, TypeQua16, TypeQua17}
	for i, ctype := range ctypes {
		mesh := GenRing2d(ctype, nr, na, r, R, 2.0*math.Pi)

		// allocate cell integrator with default integration points
		o := NewMeshIntegrator(mesh, 1)

		// compute Ix
		Ix := o.IntegrateSv(0, fcnIx)
		typekey := TypeIndexToKey[ctype]
		io.Pf("%s : Ix = %v  error = %v\n", typekey, Ix, math.Abs(Ix-anaIx))
		chk.Float64(tst, "Ix", tols[i], Ix, anaIx)
	}
}
