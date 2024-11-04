// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package msh

import (
	"math"
	"testing"

	"github.com/lei006/gomath/la"
)

var (
	benchmarkMesh *Mesh
	benchmarkMint *MeshIntegrator
	benchmarkIx   float64
)

func init() {
	r, R := 1.0, 3.0
	nr, na := 10, 35
	benchmarkMesh = GenRing2d(TypeQua17, nr, na, r, R, 2.0*math.Pi)
	benchmarkMint = NewMeshIntegrator(benchmarkMesh, 1)
}

func BenchmarkInteg(b *testing.B) {
	fcnIx := func(x la.Vector) (f float64) {
		f = x[1] * x[1]
		return
	}
	var Ix float64
	for i := 0; i < b.N; i++ {
		Ix = benchmarkMint.IntegrateSv(0, fcnIx)
	}
	benchmarkIx = Ix
}
