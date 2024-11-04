// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package msh

import (
	"math"
	"strings"
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
	"github.com/lei006/gomath/num"
)

func TestQuadpts01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("quadpts01. quadrature points")

	// compute 1D Gauss-Legendre points for comparison
	degreeMax := 5
	xref := make([][]float64, degreeMax+1)
	wref := make([][]float64, degreeMax+1)
	for n := 1; n <= degreeMax; n++ {
		xref[n], wref[n] = num.GaussLegendreXW(-1, 1, n)
	}

	// other reference values
	wilson5corner := [][]float64{
		{-1, -1, 0, 1.0 / 3.0},
		{+1, -1, 0, 1.0 / 3.0},
		{+0, +0, 0, 8.0 / 3.0},
		{-1, +1, 0, 1.0 / 3.0},
		{+1, +1, 0, 1.0 / 3.0},
	}
	wilson5stable := [][]float64{
		{-0.5776391, -0.5776391, 0, 0.999},
		{+0.5776391, -0.5776391, 0, 0.999},
		{+0.0000000, +0.0000000, 0, 0.004},
		{-0.5776391, +0.5776391, 0, 0.999},
		{+0.5776391, +0.5776391, 0, 0.999},
	}
	a := math.Sqrt(7.0 / 9.0)
	b := math.Sqrt(7.0 / 15.0)
	wa := 9.0 / 49.0
	wb := 40.0 / 49.0
	wilson8default := [][]float64{
		{-a, -a, 0, wa},
		{+0, -b, 0, wb},
		{+a, -a, 0, wa},
		{-b, +0, 0, wb},
		{+b, +0, 0, wb},
		{-a, +a, 0, wa},
		{+0, +b, 0, wb},
		{+a, +a, 0, wa},
	}

	for kind, allPts := range IntPoints {

		io.PfYel("\n--------------------------------- %d ---------------------------------\n", kind)

		switch kind {
		case KindLin:
			for key, pts := range allPts {
				res := strings.Split(key, "_")
				rule, n := res[0], io.Atoi(res[1])
				x := make([]float64, n)
				w := make([]float64, n)
				sumW := 0.0
				for i := 0; i < n; i++ {
					x[i] = pts[i][0]
					w[i] = pts[i][3]
					sumW += pts[i][3]
				}
				io.Pfblue2("\nrule = %v\n", rule)
				chk.Float64(tst, "sumW", 1e-15, sumW, 2)
				io.Pf("x = %v\n", x)
				io.Pfgreen("    %v\n", xref[n])
				io.Pf("w = %v\n", w)
				io.Pfgreen("    %v\n", wref[n])
				chk.Array(tst, io.Sf("lin:%d x", n), 1e-15, x, xref[n])
				chk.Array(tst, io.Sf("lin:%d w", n), 1e-15, w, wref[n])
			}

		case KindQua:
			for key, pts := range allPts {
				res := strings.Split(key, "_")
				rule, n := res[0], io.Atoi(res[1])
				sumW := 0.0
				for _, p := range pts {
					sumW += p[3]
				}
				io.Pfblue2("\nrule = %v\n", rule)
				chk.Float64(tst, "sumW", 1e-14, sumW, 4)
				switch rule {
				case "legendre":
					n1d := int(math.Sqrt(float64(n)))
					x1d := xref[n1d]
					w1d := wref[n1d]
					for j := 0; j < n1d; j++ {
						for i := 0; i < n1d; i++ {
							m := i + n1d*j
							x := pts[m][:2]
							v := pts[m][3]
							y := []float64{x1d[i], x1d[j]}
							w := w1d[i] * w1d[j]
							io.Pf("  %d%d x = %23v  w = %23v\n", i, j, x, v)
							io.Pfgreen("         %23v      %23v\n", y, w)
							chk.Array(tst, "x", 1e-15, x, y)
							chk.Float64(tst, "w", 1e-15, v, w)
						}
					}
				case "wilson5corner":
					for i, p := range pts {
						chk.Array(tst, io.Sf("wilson5corner %d", i), 1e-15, p, wilson5corner[i])
					}
				case "wilson5stable":
					for i, p := range pts {
						chk.Array(tst, io.Sf("wilson5stable %d", i), 1e-15, p, wilson5stable[i])
					}
				case "wilson8default":
					for i, p := range pts {
						chk.Array(tst, io.Sf("wilson8default %d", i), 1e-15, p, wilson8default[i])
					}
				default:
					tst.Errorf("cannot check rule %q\n", rule)
					return
				}
			}

		case KindHex:
			for key, pts := range allPts {
				res := strings.Split(key, "_")
				rule, n := res[0], io.Atoi(res[1])
				sumW := 0.0
				for _, p := range pts {
					sumW += p[3]
				}
				io.Pfblue2("\nrule = %v\n", rule)
				chk.Float64(tst, "sumW", 1e-14, sumW, 8)
				switch rule {
				case "legendre":
					n1d := int(math.Floor(math.Pow(float64(n), 1.0/3.0) + 0.5))
					x1d := xref[n1d]
					w1d := wref[n1d]
					for k := 0; k < n1d; k++ {
						for j := 0; j < n1d; j++ {
							for i := 0; i < n1d; i++ {
								m := i + n1d*j + (n1d*n1d)*k
								x := pts[m][:3]
								v := pts[m][3]
								y := []float64{x1d[i], x1d[j], x1d[k]}
								w := w1d[i] * w1d[j] * w1d[k]
								io.Pf("%d%d x=%18v w=%18v\n", i, j, x, v)
								io.Pfgreen("     %18v   %18v\n", y, w)
								chk.Array(tst, "x", 1e-15, x, y)
								chk.Float64(tst, "w", 1e-14, v, w)
							}
						}
					}
				}
			}
		}
	}
}
