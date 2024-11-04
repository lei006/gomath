// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gm

import (
	"math"
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
)

func Test_octree01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("octree01. PointN")

	p := &PointN{X: []float64{1, 2, 3}}

	io.Pforan("p = %+v\n", p)
	chk.Array(tst, "p.X", 1e-15, p.X, []float64{1, 2, 3})
	q := &PointN{X: []float64{2, 2, 1}}
	io.Pforan("q = %+v\n", q)
	chk.Array(tst, "q.X", 1e-15, q.X, []float64{2, 2, 1})
	if p.ExactlyTheSameX(q) {
		tst.Errorf("ExactlyTheSame should return false because points are indeed different")
		return
	}
	if p.AlmostTheSameX(q, 1e-15) {
		tst.Errorf("AlmostTheSame should return false because points are different within given tolerance (1e-15)")
		return
	}
	if p.AlmostTheSameX(q, 1.0) {
		tst.Errorf("AlmostTheSame should return false because points are different within given tolerance (1.0)")
		return
	}
	if !p.AlmostTheSameX(q, 2.0) {
		tst.Errorf("AlmostTheSame should return true because points are different within given tolerance (2.0)")
		return
	}

	a := p.GetCloneX()
	chk.Array(tst, "a == p", 1e-15, a.X, []float64{1, 2, 3})

	dap := DistPointPointN(a, p)
	chk.Float64(tst, "dist(a,p)", 1e-15, dap, 0)

	dpq := DistPointPointN(p, q)
	chk.Float64(tst, "dist(p,q)", 1e-15, dpq, math.Sqrt(5.0))

	c1 := NewPointNdim(2)
	c1.X[0] = 1
	c1.X[1] = 2
	chk.Array(tst, "c1", 1e-15, c1.X, []float64{1, 2})

	c2 := NewPointN(1, 2)
	chk.Array(tst, "c2", 1e-15, c2.X, []float64{1, 2})

	c3 := NewPointN(1, 2, 3)
	chk.Array(tst, "c3", 1e-15, c3.X, []float64{1, 2, 3})
}

func Test_octree02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("octree02. BoxN and DistPointPointN")

	b := &BoxN{&PointN{X: []float64{-1, -2, -3}}, &PointN{X: []float64{3, 2, 1}}, 0}
	mid := b.GetMid()
	delta := b.GetSize()
	chk.Array(tst, "mid", 1e-15, mid, []float64{1, 0, -1})
	chk.Array(tst, "delta", 1e-15, delta, []float64{4, 4, 4})

	p := &PointN{X: []float64{-2, 0, 0}}
	dist := DistPointBoxN(p, b)
	io.Pforan("dist = %v\n", dist)
	chk.Float64(tst, "dist(p,b)", 1e-15, dist, 1.0)

	if b.IsInside(p) {
		tst.Errorf("is inside box failed")
		return
	}

	q := &PointN{X: []float64{-2, 3, 0}}
	dist = DistPointBoxN(q, b)
	io.Pforan("dist = %v\n", dist)
	chk.Float64(tst, "dist(q,b)", 1e-15, dist, math.Sqrt2)

	if b.IsInside(q) {
		tst.Errorf("is inside box failed")
		return
	}

	r := &PointN{X: []float64{-2, 3, 2}}
	dist = DistPointBoxN(r, b)
	io.Pforan("dist = %v\n", dist)
	chk.Float64(tst, "dist(r,b)", 1e-15, dist, math.Sqrt(3.0))

	if b.IsInside(r) {
		tst.Errorf("is inside box failed")
		return
	}

	s := &PointN{X: []float64{0, 0, 0}}
	dist = DistPointBoxN(s, b)
	io.Pforan("dist = %v\n", dist)
	chk.Float64(tst, "dist(s,b)", 1e-15, dist, 0)

	if !b.IsInside(s) {
		tst.Errorf("is inside box failed")
		return
	}

	s.X[0] = 1.0
	dist = DistPointBoxN(s, b)
	io.Pforan("dist = %v\n", dist)
	chk.Float64(tst, "dist(s,b)", 1e-15, dist, 0)

	if !b.IsInside(s) {
		tst.Errorf("is inside box failed")
		return
	}

	b1 := NewBoxN(0, 1, 2, 3)       // xmin,xmax, ymin,ymax
	b2 := NewBoxN(0, 1, 2, 3, 4, 5) // xmin,xmax, ymin,ymax
	chk.Array(tst, "b1.Lo", 1e-15, b1.Lo.X, []float64{0, 2})
	chk.Array(tst, "b1.Hi", 1e-15, b1.Hi.X, []float64{1, 3})
	chk.Array(tst, "b2.Lo", 1e-15, b2.Lo.X, []float64{0, 2, 4})
	chk.Array(tst, "b2.Hi", 1e-15, b2.Hi.X, []float64{1, 3, 5})
}

func Test_octree03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("octree03. qobox")

	o := NewOctree(-1, 1, -2, 1) // 4 numbers => 2D
	chk.Array(tst, "blo", 1e-15, o.blo, []float64{-1, -2})
	chk.Array(tst, "bscale", 1e-15, o.bscale, []float64{2, 3})

	b1 := o.qobox(1)
	b2 := o.qobox(2)
	b3 := o.qobox(3)
	b4 := o.qobox(4)
	b5 := o.qobox(5)
	b6 := o.qobox(6)
	b7 := o.qobox(7)
	b15 := o.qobox(15)
	b21 := o.qobox(21)
	b22 := o.qobox(22)
	b41 := o.qobox(41)
	b45 := o.qobox(45)
	b52 := o.qobox(52)
	chk.Array(tst, "1: lo", 1e-15, b1.Lo.X, []float64{-1, -2})
	chk.Array(tst, "1: hi", 1e-15, b1.Hi.X, []float64{1, 1})
	chk.Array(tst, "2: lo", 1e-15, b2.Lo.X, []float64{-1, -2})
	chk.Array(tst, "2: hi", 1e-15, b2.Hi.X, []float64{0, -0.5})
	chk.Array(tst, "3: lo", 1e-15, b3.Lo.X, []float64{0, -2})
	chk.Array(tst, "3: hi", 1e-15, b3.Hi.X, []float64{1, -0.5})
	chk.Array(tst, "4: lo", 1e-15, b4.Lo.X, []float64{-1, -0.5})
	chk.Array(tst, "4: hi", 1e-15, b4.Hi.X, []float64{0, 1.0})
	chk.Array(tst, "5: lo", 1e-15, b5.Lo.X, []float64{0, -0.5})
	chk.Array(tst, "5: hi", 1e-15, b5.Hi.X, []float64{1, 1})
	chk.Array(tst, "6: lo", 1e-15, b6.Lo.X, []float64{-1, -2})
	chk.Array(tst, "6: hi", 1e-15, b6.Hi.X, []float64{-0.5, -1.25})
	chk.Array(tst, "7: lo", 1e-15, b7.Lo.X, []float64{-0.5, -2})
	chk.Array(tst, "7: hi", 1e-15, b7.Hi.X, []float64{0, -1.25})
	chk.Array(tst, "15: lo", 1e-15, b15.Lo.X, []float64{-0.5, -0.5})
	chk.Array(tst, "15: hi", 1e-15, b15.Hi.X, []float64{0, 0.25})
	chk.Array(tst, "21: lo", 1e-15, b21.Lo.X, []float64{0.5, 0.25})
	chk.Array(tst, "21: hi", 1e-15, b21.Hi.X, []float64{1, 1})
	chk.Array(tst, "22: lo", 1e-15, b22.Lo.X, []float64{-1, -2})
	chk.Array(tst, "22: hi", 1e-15, b22.Hi.X, []float64{-0.75, -1.625})
	chk.Array(tst, "41: lo", 1e-15, b41.Lo.X, []float64{0.25, -1.625})
	chk.Array(tst, "41: hi", 1e-15, b41.Hi.X, []float64{0.5, -1.25})
	chk.Array(tst, "45: lo", 1e-15, b45.Lo.X, []float64{0.75, -1.625})
	chk.Array(tst, "45: hi", 1e-15, b45.Hi.X, []float64{1, -1.25})
	chk.Array(tst, "52: lo", 1e-15, b52.Lo.X, []float64{0.5, -0.875})
	chk.Array(tst, "52: hi", 1e-15, b52.Hi.X, []float64{0.75, -0.5})
}
