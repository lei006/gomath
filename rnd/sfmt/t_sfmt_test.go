// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows
// +build !windows

package sfmt

import (
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
)

func Test_sfmt01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sfmt01. integers")

	Init(4321)

	PrintIDString()
	//TODO: add SFMT original test here
	//io.Pf("64 bit generated randoms\n")
	//io.Pf("init_gen_rand__________\n")

	nsamples := 10
	for i := 0; i < nsamples; i++ {
		gen := Rand(0, 10)
		io.Pforan("gen = %v\n", gen)
	}
}
