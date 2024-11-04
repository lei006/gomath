// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"testing"

	"github.com/lei006/gomath/chk"
)

func TestBwEuler01a(tst *testing.T) {

	//verbose()
	chk.PrintTitle("BwEuler01a. Backward-Euler (analytical Jacobian)")

	// problem
	p := ProbHwEq11()

	// configuration
	conf := NewConfig("bweuler", "")
	conf.SetFixedH(p.Dx, p.Xf)
	conf.SetStepOut(true, nil)

	// solver
	sol := NewSolver(p.Ndim, conf, p.Fcn, p.Jac, nil)
	defer sol.Free()

	// solve ODE
	sol.Solve(p.Y, 0.0, p.Xf)

	// check Stat
	chk.Int(tst, "number of F evaluations ", sol.Stat.Nfeval, 80)
	chk.Int(tst, "number of J evaluations ", sol.Stat.Njeval, 40)
	chk.Int(tst, "total number of steps   ", sol.Stat.Nsteps, 40)
	chk.Int(tst, "number of accepted steps", sol.Stat.Naccepted, 0)
	chk.Int(tst, "number of rejected steps", sol.Stat.Nrejected, 0)
	chk.Int(tst, "number of decompositions", sol.Stat.Ndecomp, 40)
	chk.Int(tst, "number of lin solutions ", sol.Stat.Nlinsol, 40)
	chk.Int(tst, "max number of iterations", sol.Stat.Nitmax, 2)

	// check results
	chk.Float64(tst, "yFin", 1e-4, p.Y[0], p.CalcYana(0, p.Xf))
}

func TestBwEuler01b(tst *testing.T) {

	//verbose()
	chk.PrintTitle("BwEuler01b. Backward-Euler (numerical Jacobian)")

	// problem
	p := ProbHwEq11()

	// configuration
	conf := NewConfig("bweuler", "")
	conf.SetFixedH(p.Dx, p.Xf)
	conf.SetStepOut(true, nil)

	// solver
	sol := NewSolver(p.Ndim, conf, p.Fcn, nil, nil)
	defer sol.Free()

	// solve ODE
	sol.Solve(p.Y, 0.0, p.Xf)

	// check Stat
	chk.Int(tst, "number of F evaluations ", sol.Stat.Nfeval, 120)
	chk.Int(tst, "number of J evaluations ", sol.Stat.Njeval, 40)
	chk.Int(tst, "total number of steps   ", sol.Stat.Nsteps, 40)
	chk.Int(tst, "number of accepted steps", sol.Stat.Naccepted, 0)
	chk.Int(tst, "number of rejected steps", sol.Stat.Nrejected, 0)
	chk.Int(tst, "number of decompositions", sol.Stat.Ndecomp, 40)
	chk.Int(tst, "number of lin solutions ", sol.Stat.Nlinsol, 40)
	chk.Int(tst, "max number of iterations", sol.Stat.Nitmax, 2)

	// check results
	chk.Float64(tst, "yFin", 1e-4, p.Y[0], p.CalcYana(0, p.Xf))
}
