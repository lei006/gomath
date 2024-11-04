// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/fun"
	"github.com/lei006/gomath/io"
	"github.com/lei006/gomath/la"
)

func TestErk01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Erk01. coefficients")

	methods := []string{
		"moeuler", "rk2",
		"rk3", "heun3",
		"rk4", "rk4-3/8",
		"merson4", "zonneveld4", "fehlberg4",
		"dopri5", "verner6",
		"fehlberg7", "dopri8",
	}

	tols1 := []float64{
		1e-17, 1e-17,
		1e-17, 1e-17,
		1e-17, 1e-15,
		1e-17, 1e-17, 1e-15,
		1e-15, 1e-15,
		1e-14, 1e-14,
	}

	tols2 := []float64{
		1e-17, 1e-17,
		1e-15, 1e-17,
		1e-15, 1e-16,
		1e-15, 1e-15, 1e-15,
		1e-15, 1e-16,
		1e-16, 1e-15,
	}

	for im, method := range methods {
		io.Pf("\n------------------------------------ %s ---------------------------------------\n", method)
		o := newERK(method).(*ExplicitRK)
		chk.Int(tst, "len(A)=nstg", len(o.A), o.Nstg)
		chk.Int(tst, "len(B)=nstg", len(o.B), o.Nstg)
		chk.Int(tst, "len(C)=nstg", len(o.C), o.Nstg)
		if o.Embedded {
			chk.Int(tst, "len(E)=nstg", len(o.E), o.Nstg)
		}
		for i := 0; i < o.Nstg; i++ {
			chk.Int(tst, "ncol(A)=nstg", len(o.A[i]), o.Nstg)
		}

		io.Pf("\nc coefficients: ci = Σ_j aij\n")
		var sumrow float64
		for i := 0; i < o.Nstg; i++ {
			sumrow = 0.0
			for j := 0; j < i; j++ {
				sumrow += o.A[i][j]
			}
			chk.AnaNum(tst, io.Sf("Σa%2dj", i), tols1[im], sumrow, o.C[i], chk.Verbose)
		}
		if o.Embedded && len(o.Be) == len(o.E) {
			io.Pf("\nerror estimator\n")
			for i := 0; i < o.Nstg; i++ {
				chk.AnaNum(tst, io.Sf("E%d=B%d-Be%d", i, i, i), 1e-15, o.E[i], o.B[i]-o.Be[i], chk.Verbose)
			}
		}

		io.Pf("\nEquations (1.11) of [1, page 135-136] and (5.20) of [1, page 181-182]\n")
		SbcPow := 0.0
		for p := 1; p <= o.P; p++ {
			SbcPow = 0.0
			for i := 0; i < o.Nstg; i++ {
				SbcPow += o.B[i] * fun.PowP(o.C[i], uint32(p-1))
			}
			chk.AnaNum(tst, io.Sf("Σbi⋅ciⁿ⁻¹     =1/%d ", p), tols2[im], SbcPow, 1.0/float64(p), chk.Verbose)
		}
		if o.P < 3 {
			continue
		}
		Sbac := 0.0
		Sbcac := 0.0
		Sbac2 := 0.0
		Sbaac := 0.0
		for i := 0; i < o.Nstg; i++ {
			for j := 0; j < o.Nstg; j++ {
				Sbac += o.B[i] * o.A[i][j] * o.C[j]
				Sbcac += o.B[i] * o.C[i] * o.A[i][j] * o.C[j]
				Sbac2 += o.B[i] * o.A[i][j] * o.C[j] * o.C[j]
				for k := 0; k < o.Nstg; k++ {
					Sbaac += o.B[i] * o.A[i][j] * o.A[j][k] * o.C[k]
				}
			}
		}
		chk.AnaNum(tst, "Σbi⋅aij⋅cj    =1/6 ", tols2[im], Sbac, 1.0/6.0, chk.Verbose)
		if o.P < 4 {
			continue
		}
		chk.AnaNum(tst, "Σbi⋅ci⋅aij⋅cj =1/8 ", tols2[im], Sbcac, 1.0/8.0, chk.Verbose)
		chk.AnaNum(tst, "Σbi⋅aij⋅cj²   =1/12", tols2[im], Sbac2, 1.0/12.0, chk.Verbose)
		chk.AnaNum(tst, "Σbi⋅aij⋅ajk⋅ck=1/24", tols2[im], Sbaac, 1.0/24.0, chk.Verbose)
	}
}

func TestErk02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Erk02. simple problem")

	// problem
	p := ProbSimpleNdim2()
	p.Dx = 0.17

	// try methods
	tols := []float64{0.11, 0.11, 0.0077, 0.009, 5.8e-4, 5.8e-4, 1.3e-4, 5.8e-4}
	for im, method := range []string{"moeuler", "rk2", "rk3", "heun3", "rk4", "rk4-3/8", "merson4", "zonneveld4"} {

		// solve problem
		y, _, out := p.Solve(method, true, false)

		// results
		X := out.GetStepX()
		chk.Float64(tst, "xf == X[last]", 1e-15, p.Xf, X[out.StepIdx-1])
		chk.AnaNum(tst, "y0(xf)", tols[im], y[0], p.CalcYana(0, p.Xf), chk.Verbose)
		chk.AnaNum(tst, "y1(xf)", tols[im], y[1], p.CalcYana(1, p.Xf), chk.Verbose)
	}
}

func TestErk03a(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Erk03a. rk-fours. problem a")

	// problem
	p := ProbSimpleNdim4a()

	// reference solution
	yExact := la.NewVector(p.Ndim)
	p.Yana(yExact, p.Xf)

	// run test
	methods := []string{"rk4", "rk4-3/8", "merson4", "zonneveld4"}
	orders := []float64{4, 4, 4, 4}
	tols := []float64{0.011, 0.023, 0.00471, 0.011}
	p.ConvergenceTest(tst, 1e-3, 1e-2, 3, yExact, methods, orders, tols)
}

func TestErk03b(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Erk03b. rk-fours. problem b")

	// problem
	p := ProbSimpleNdim4b()

	// reference solution
	yExact := la.NewVector(p.Ndim)
	p.Yana(yExact, p.Xf)

	// run test
	methods := []string{"rk4", "rk4-3/8", "merson4", "zonneveld4"}
	orders := []float64{4, 4, 4, 4}
	tols := []float64{0.086, 0.164, 0.07, 0.09}
	p.ConvergenceTest(tst, 1e-3, 1e-2, 3, yExact, methods, orders, tols)
}

func TestErk04(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Erk04. problem a. methods with p<5")

	// problem
	p := ProbSimpleNdim4a()

	// reference solution
	yExact := la.NewVector(p.Ndim)
	p.Yana(yExact, p.Xf)

	// run test
	methods := []string{"fweuler", "bweuler", "moeuler", "rk2", "rk3", "heun3",
		"rk4", "rk4-3/8", "merson4", "zonneveld4", "fehlberg4"}
	orders := []float64{1, 1, 2, 2, 3, 3,
		4, 4, 4, 4, 4}
	tols := []float64{0.043, 0.0176, 0.11, 0.049, 0.016, 0.0023,
		0.086, 0.164, 0.07, 0.09, 0.005}
	p.ConvergenceTest(tst, 1e-3, 1e-2, 3, yExact, methods, orders, tols)
}

func TestErk05(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Erk05. Convergence of high-order methods")

	// problem
	p := ProbVanDerPol(0, true)

	// reference solution
	yExact := p.Y.GetCopy()

	// test
	methods := []string{"rk4", "fehlberg4", "dopri5", "verner6", "fehlberg7", "dopri8"}
	orders := []float64{3.9874, 4.6923, 6.3012, 6.5916, 8.0283, 6.8463}
	tols := []float64{1e-4, 1e-4, 1e-4, 1e-4, 1e-4, 1e-4}
	p.ConvergenceTest(tst, 8e-2, 3e-1, 3, yExact, methods, orders, tols)
}
