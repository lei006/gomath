// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pde

import (
	"math"
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/fun"
	"github.com/lei006/gomath/gm"
	"github.com/lei006/gomath/io"
	"github.com/lei006/gomath/la"
	"github.com/lei006/gomath/utl"
)

func TestSpc01a(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Spc01a. Auu matrix after homogeneous bcs")

	// lagrange interpolators (5x5 grid)
	l := fun.NewLagIntSet(2, []int{4, 4}, []string{"cgl", "cgl"})

	// grid
	g := new(gm.Grid)
	g.RectSet2d(l[0].X, l[1].X)

	// solver
	p := utl.Params{{N: "kx", V: 1}, {N: "ky", V: 1}}
	s := NewSpcLaplacian(p, l, g, nil)

	// homogeneous boundary conditions
	s.SetHbc()

	// assemble
	s.Assemble(false)
	Duu := s.Eqs.Auu.ToDense()
	io.Pf("%v\n", Duu.Print("%7.2f"))

	// check
	zzz := 0.0
	chk.Deep2(tst, "Auu", 1e-14, Duu.GetDeep2(), [][]float64{
		{-28, +6., -2. /* */, +6., zzz, zzz /* */, -2., zzz, zzz},
		{+4., -20, +4. /* */, zzz, +6., zzz /* */, zzz, -2., zzz},
		{-2., +6., -28 /* */, zzz, zzz, +6. /* */, zzz, zzz, -2.},

		{+4., zzz, zzz /* */, -20, +6., -2. /* */, +4., zzz, zzz},
		{zzz, +4., zzz /* */, +4., -12, +4. /* */, zzz, +4., zzz},
		{zzz, zzz, +4. /* */, -2., +6., -20 /* */, zzz, zzz, +4.},

		{-2., zzz, zzz /* */, +6., zzz, zzz /* */, -28, +6., -2.},
		{zzz, -2., zzz /* */, zzz, +6., zzz /* */, +4., -20, +4.},
		{zzz, zzz, -2. /* */, zzz, zzz, +6. /* */, -2., +6., -28},
	})
}

func TestSpc01b(tst *testing.T) {
	//verbose()
	chk.PrintTitle("Spc01b. panic on params")
	defer chk.RecoverTstPanicIsOK(tst)
	NewSpcLaplacian(nil, nil, nil, nil)
}

func TestSpc02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Spc02. simple Dirichlet problem (unif-grid / Laplace)")

	// solve problem
	//    ∂²u     ∂²u
	//    ———  +  ——— = 0    with   u(x,0)=1   u(3,y)=2   u(x,3)=2   u(0,y)=1
	//    ∂x²     ∂y²               (bottom)   (right)    (top)      (left)

	// Lagrange interpolators
	lis := fun.NewLagIntSet(2, []int{3, 3}, []string{"uni", "uni"})

	// grid
	g := new(gm.Grid)
	g.RectSet2dU([]float64{0, 0}, []float64{2, 2}, lis[0].X, lis[1].X)

	// solver
	p := utl.Params{{N: "kx", V: 1}, {N: "ky", V: 1}}
	s := NewSpcLaplacian(p, lis, g, nil)

	// essential boundary conditions
	s.AddEbc(10, 1.0, nil) // left
	s.AddEbc(11, 2.0, nil) // right
	s.AddEbc(20, 1.0, nil) // bottom
	s.AddEbc(21, 2.0, nil) // top

	// set equations and assemble A matrix
	reactions := true
	s.Assemble(reactions)

	// check equations (must be after Assemble)
	chk.Int(tst, "number of equations == number of nodes", s.Eqs.N, 16)
	chk.Ints(tst, "UtoF", s.Eqs.UtoF, []int{5, 6, 9, 10})
	chk.Ints(tst, "KtoF", s.Eqs.KtoF, []int{0, 1, 2, 3, 4, 7, 8, 11, 12, 13, 14, 15})

	// check Duu
	Duu := s.Eqs.Auu.ToDense()
	io.Pf("Auu =\n%v\n", Duu.Print("%7.2f"))
	chk.Deep2(tst, "Auu", 1e-17, Duu.GetDeep2(), [][]float64{
		{-9.00, +2.25, +2.25, +0.00}, // 0 ⇒ node (1,1) 5
		{+2.25, -9.00, +0.00, +2.25}, // 1 ⇒ node (2,1) 6
		{+2.25, +0.00, -9.00, +2.25}, // 2 ⇒ node (1,2) 9
		{+0.00, +2.25, +2.25, -9.00}, // 3 ⇒ node (2,2) 10
	})

	// solve
	u, f := s.SolveSteady(reactions)
	chk.Array(tst, "xk", 1e-17, s.Eqs.Xk, []float64{1, 1, 1, 1, 1, 2, 1, 2, 2, 2, 2, 2})
	chk.Array(tst, "u", 1e-15, u, []float64{1, 1, 1, 1, 1, 1.25, 1.5, 2, 1, 1.5, 1.75, 2, 2, 2, 2, 2})

	// check f
	sFull := NewSpcLaplacian(p, lis, g, nil)
	sFull.Assemble(false)
	K := sFull.Eqs.Auu.ToMatrix(nil)
	Fref := la.NewVector(g.Size())
	la.SpMatVecMul(Fref, 1.0, K, u)
	chk.Array(tst, "f", 1e-14, f, Fref)

	// get results over grid
	uu := g.MapMeshgrid2d(u)
	chk.Deep2(tst, "uu", 1e-15, uu, [][]float64{
		{1.00, 1.00, 1.00, 1.00},
		{1.00, 1.25, 1.50, 2.00},
		{1.00, 1.50, 1.75, 2.00},
		{2.00, 2.00, 2.00, 2.00},
	})
}

func TestSpc03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Spc03. Trefethen's p16")

	// solve problem
	//    ∂²u     ∂²u
	//    ———  +  ——— = 10 sin(8x⋅(y-1))    with   homogeneous BCs
	//    ∂x²     ∂y²

	// Lagrange interpolators
	lis := fun.NewLagIntSet(2, []int{4, 4}, []string{"cgl", "cgl"})

	// grid
	g := new(gm.Grid)
	g.RectSet2dU([]float64{-1, -1}, []float64{+1, +1}, lis[0].X, lis[1].X)

	// solver
	p := utl.Params{{N: "kx", V: 1}, {N: "ky", V: 1}}
	source := func(x la.Vector, t float64) float64 {
		return 10.0 * math.Sin(8.0*x[0]*(x[1]-1.0))
	}
	s := NewSpcLaplacian(p, lis, g, source)

	// homogeneous boundary conditions
	s.SetHbc()

	// set equations and assemble A matrix
	reactions := false
	s.Assemble(reactions)

	// solve
	u, _ := s.SolveSteady(reactions)

	// check
	uu := g.MapMeshgrid2d(u)
	chk.Deep2(tst, "uu", 1e-14, uu, [][]float64{
		{0, +0.000000000000000, +0, +0.000000000000000, 0},
		{0, +0.181363633964132, +0, -0.181363633964131, 0},
		{0, +0.292713394079481, +0, -0.292713394079479, 0},
		{0, -0.329593843114906, +0, +0.329593843114906, 0},
		{0, +0.000000000000000, +0, +0.000000000000000, 0},
	})
}

func TestSpc04(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Spc04. Kopriva 7.1.4 p259")

	// auxiliary
	polar := func(x []float64) (r, θ float64) {
		r = math.Sqrt(x[0]*x[0] + x[1]*x[1])
		θ = math.Atan2(x[1], x[0])
		return
	}

	// Lagrange interpolators
	lis := fun.NewLagIntSet(2, []int{8, 8}, []string{"cgl", "cgl"})

	// grid
	trf := gm.FactoryTfinite.Surf2dQuarterRing(1.0, 3.0)
	g := new(gm.Grid)
	g.SetTransfinite2d(trf, lis[0].X, lis[1].X)

	// solver
	p := utl.Params{{N: "kx", V: 1}, {N: "ky", V: 1}}
	source := func(x la.Vector, t float64) float64 {
		r, θ := polar(x)
		return -16.0 * math.Log(r) * math.Sin(4.0*θ) / (r * r)
	}
	s := NewSpcLaplacian(p, lis, g, source)

	// essential boundary conditions
	s.AddEbc(10, 0.0, nil) // left
	s.AddEbc(11, 0.0, func(x la.Vector, t float64) float64 {
		r, θ := polar(x)
		return math.Log(r) * math.Sin(4.0*θ)
	}) // right
	s.AddEbc(20, 0.0, nil) // bottom
	s.AddEbc(21, 0.0, nil) // top

	// solve
	reactions := false
	s.Assemble(reactions)
	u, _ := s.SolveSteady(reactions)

	// check
	ana := func(x []float64) float64 {
		r, θ := polar(x)
		return math.Log(r) * math.Sin(4.0*θ)
	}
	for n := 0; n < g.Npts(1); n++ {
		for m := 0; m < g.Npts(0); m++ {
			I := g.IndexMNPtoI(m, n, 0)
			x := g.X(m, n, 0)
			chk.AnaNum(tst, io.Sf("u(%5.2f,%5.2f)", x[0], x[1]), 1e-3, u[I], ana(x), chk.Verbose)
		}
	}
}

func TestSpc05(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Spc05. Kopriva 7.1.5 p261")

	// Lagrange interpolators
	lis := fun.NewLagIntSet(2, []int{8, 8}, []string{"cgl", "cgl"})

	// grid
	r0, rinf := 0.5, 10.0
	trf := gm.FactoryTfinite.Surf2dHalfRing(r0, rinf)
	g := new(gm.Grid)
	g.SetTransfinite2d(trf, lis[0].X, lis[1].X)

	// solver
	p := utl.Params{{N: "kx", V: 1}, {N: "ky", V: 1}}
	s := NewSpcLaplacian(p, lis, g, nil)

	// essential boundary conditions
	Vinf := 0.5
	s.AddNbc(10, 0.0, nil) // inner circle
	s.AddEbc(11, 0.0, func(x la.Vector, t float64) float64 {
		return Vinf * x[0]
	}) // outer circle
	s.AddNbc(20, 0.0, nil) // right horizontal line
	s.AddNbc(21, 0.0, nil) // left horizontal line

	// solve
	reactions := false
	s.Assemble(reactions)
	u, _ := s.SolveSteady(reactions)

	// check
	ana := func(x []float64) float64 {
		r := math.Sqrt(x[0]*x[0] + x[1]*x[1])
		θ := math.Atan2(x[1], x[0])
		return Vinf * (r + r0*r0/r) * math.Cos(θ)
	}
	for n := 0; n < g.Npts(1); n++ {
		for m := 0; m < g.Npts(0); m++ {
			I := g.IndexMNPtoI(m, n, 0)
			x := g.X(m, n, 0)
			//chk.PrintAnaNum(io.Sf("u(%5.2f,%5.2f)", x[0], x[1]), 1e-3, u[I], ana(x), chk.Verbose)
			chk.AnaNum(tst, io.Sf("u(%5.2f,%5.2f)", x[0], x[1]), 0.053, u[I], ana(x), chk.Verbose)
		}
	}
}
