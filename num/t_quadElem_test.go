// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package num

import (
	"math"
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
)

func Test_QuadElem01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("QuadElem01. Trapz and Simpson Elementary")

	y := func(x float64) (res float64) {
		res = math.Sqrt(1.0 + math.Pow(math.Sin(x), 3.0))
		return
	}
	Acor := 1.08268158558

	// trapezoidal rule
	var T QuadElementary
	T = new(ElementaryTrapz)
	T.Init(y, 0, 1, 1e-11)
	A := T.Integrate()
	io.Pforan("A  = %v\n", A)
	chk.Float64(tst, "A", 1e-11, A, Acor)

	// Simpson's rule
	var S QuadElementary
	S = new(ElementarySimpson)
	S.Init(y, 0, 1, 1e-11)
	A = S.Integrate()
	io.Pforan("A  = %v\n", A)
	chk.Float64(tst, "A", 1e-11, A, Acor)
}
