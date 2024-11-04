// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows
// +build !windows

package dsfmt

import (
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
)

func Test_dsfmt01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("dsfmt01. floats")

	Init(1234)

	nsamples := 10
	for i := 0; i < nsamples; i++ {
		gen := Rand(10, 20)
		io.Pforan("gen = %v\n", gen)
	}
}
