// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utl

import (
	"testing"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
)

func Test_sort01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sort01")

	i := []int{33, 0, 7, 8}
	x := []float64{1000.33, 0, -77.7, 88.8}

	io.Pforan("by 'i'\n")
	I, X, _, _ := SortQuadruples(i, x, nil, nil, "i")
	chk.Ints(tst, "i", I, []int{0, 7, 8, 33})
	chk.Array(tst, "x", 1e-16, X, []float64{0, -77.7, 88.8, 1000.33})

	io.Pforan("by 'x'\n")
	I, X, _, _ = SortQuadruples(i, x, nil, nil, "x")
	chk.Ints(tst, "i", I, []int{7, 0, 8, 33})
	chk.Array(tst, "x", 1e-16, X, []float64{-77.7, 0.0, 88.8, 1000.33})

	x = []float64{1000.33, 0, -77.7, 88.8}
	Sort4(&x[0], &x[1], &x[2], &x[3])
	io.Pforan("x = %v\n", x)
	chk.Array(tst, "x", 1e-16, x, []float64{-77.7, 0.0, 88.8, 1000.33})

	x = []float64{1, 10.33, 0, -8.7}
	Sort4(&x[0], &x[1], &x[2], &x[3])
	io.Pforan("x = %v\n", x)
	chk.Array(tst, "x", 1e-16, x, []float64{-8.7, 0, 1, 10.33})

	x = []float64{100.33, 10, -77.7, 8.8}
	Sort4(&x[0], &x[1], &x[2], &x[3])
	io.Pforan("x = %v\n", x)
	chk.Array(tst, "x", 1e-16, x, []float64{-77.7, 8.8, 10, 100.33})

	x = []float64{-10.33, 0, 7.7, -8.8}
	Sort4(&x[0], &x[1], &x[2], &x[3])
	io.Pforan("x = %v\n", x)
	chk.Array(tst, "x", 1e-16, x, []float64{-10.33, -8.8, 0, 7.7})

	x = []float64{-1000.33, 8, -177.7, 0.8}
	Sort4(&x[0], &x[1], &x[2], &x[3])
	io.Pforan("x = %v\n", x)
	chk.Array(tst, "x", 1e-16, x, []float64{-1000.33, -177.7, 0.8, 8})
}

func Test_sort02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sort02")

	i := []int{33, 0, 7, 8}
	x := []float64{1000.33, 0, -77.7, 88.8}
	y := []float64{1e-5, 1e-7, 1e-2, 1e-9}
	z := []float64{-8000, -7000, 0, -1}

	io.Pforan("by 'i'\n")
	I, X, Y, Z := SortQuadruples(i, x, y, z, "i")
	chk.Ints(tst, "i", I, []int{0, 7, 8, 33})
	chk.Array(tst, "x", 1e-16, X, []float64{0, -77.7, 88.8, 1000.33})
	chk.Array(tst, "y", 1e-16, Y, []float64{1e-7, 1e-2, 1e-9, 1e-5})
	chk.Array(tst, "z", 1e-16, Z, []float64{-7000, 0, -1, -8000})

	io.Pforan("by 'x'\n")
	I, X, Y, Z = SortQuadruples(i, x, y, z, "x")
	chk.Ints(tst, "i", I, []int{7, 0, 8, 33})
	chk.Array(tst, "x", 1e-16, X, []float64{-77.7, 0.0, 88.8, 1000.33})
	chk.Array(tst, "y", 1e-16, Y, []float64{1e-2, 1e-7, 1e-9, 1e-5})
	chk.Array(tst, "z", 1e-16, Z, []float64{0, -7000, -1, -8000})

	io.Pforan("by 'y'\n")
	I, X, Y, Z = SortQuadruples(i, x, y, z, "y")
	chk.Ints(tst, "i", I, []int{8, 0, 33, 7})
	chk.Array(tst, "x", 1e-16, X, []float64{88.8, 0, 1000.33, -77.7})
	chk.Array(tst, "y", 1e-16, Y, []float64{1e-9, 1e-7, 1e-5, 1e-2})
	chk.Array(tst, "z", 1e-16, Z, []float64{-1, -7000, -8000, 0})

	io.Pforan("by 'z'\n")
	I, X, Y, Z = SortQuadruples(i, x, y, z, "z")
	chk.Ints(tst, "i", I, []int{33, 0, 8, 7})
	chk.Array(tst, "x", 1e-16, X, []float64{1000.33, 0, 88.8, -77.7})
	chk.Array(tst, "y", 1e-16, Y, []float64{1e-5, 1e-7, 1e-9, 1e-2})
	chk.Array(tst, "z", 1e-16, Z, []float64{-8000, -7000, -1, 0})
}

func Test_sort03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sort03")

	a, b, c := 8.0, -5.5, 4.0
	Sort3(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{-5.5, 4, 8})
	Sort3Desc(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{8, 4, -5.5})

	a, b, c = -18.0, -5.5, 4.0
	Sort3(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{-18, -5.5, 4})
	Sort3Desc(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{4, -5.5, -18})

	a, b, c = 1.0, 2.0, 3.0
	Sort3(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{1, 2, 3})
	Sort3Desc(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{3, 2, 1})

	a, b, c = 1.0, 3.0, 2.0
	Sort3(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{1, 2, 3})
	Sort3Desc(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{3, 2, 1})

	a, b, c = 3.0, 2.0, 1.0
	Sort3(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{1, 2, 3})
	Sort3Desc(&a, &b, &c)
	io.Pforan("a b c = %v %v %v\n", a, b, c)
	chk.Array(tst, "a b c", 1e-16, []float64{a, b, c}, []float64{3, 2, 1})
}

func Test_sort04(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sort04")

	A := []int{-3, -7, 8, 11, 3, 0, -11, 8}
	B := IntGetSorted(A)
	chk.Ints(tst, "sorted A", B, []int{-11, -7, -3, 0, 3, 8, 8, 11})

	a := []float64{-3, -7, 8, 11, 3, 0, -11, 8}
	b := GetSorted(a)
	chk.Array(tst, "sorted a", 1e-16, b, []float64{-11, -7, -3, 0, 3, 8, 8, 11})
}

func Test_sort05(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sort05")

	a := map[string]int{"a": 1, "z": 2, "c": 3, "y": 4, "d": 5, "b": 6, "x": 7}
	b := map[string]float64{"a": 1, "z": 2, "c": 3, "y": 4, "d": 5, "b": 6, "x": 7}
	c := map[string]bool{"a": false, "z": true, "c": false, "y": true, "d": true, "b": false, "x": true}
	ka := StrIntMapSort(a)
	kb := StrFltMapSort(b)
	kc := StrBoolMapSort(c)
	io.Pforan("sorted_keys(a) = %v\n", ka)
	io.Pforan("sorted_keys(b) = %v\n", kb)
	io.Pforan("sorted_keys(c) = %v\n", kc)
	chk.Strings(tst, "ka", ka, []string{"a", "b", "c", "d", "x", "y", "z"})
	chk.Strings(tst, "kb", kb, []string{"a", "b", "c", "d", "x", "y", "z"})
	chk.Strings(tst, "kc", kc, []string{"a", "b", "c", "d", "x", "y", "z"})

	ka, va := StrIntMapSortSplit(a)
	io.Pfpink("sorted_keys(a) = %v\n", ka)
	io.Pfpink("sorted_vals(a) = %v\n", va)
	chk.Strings(tst, "ka", ka, []string{"a", "b", "c", "d", "x", "y", "z"})
	chk.Ints(tst, "va", va, []int{1, 6, 3, 5, 7, 4, 2})

	kb, vb := StrFltMapSortSplit(b)
	io.Pfcyan("sorted_keys(b) = %v\n", kb)
	io.Pfcyan("sorted_vals(b) = %v\n", vb)
	chk.Strings(tst, "kb", kb, []string{"a", "b", "c", "d", "x", "y", "z"})
	chk.Array(tst, "vb", 1e-16, vb, []float64{1, 6, 3, 5, 7, 4, 2})

	kc, vc := StrBoolMapSortSplit(c)
	io.Pfcyan("sorted_keys(c) = %v\n", kc)
	io.Pfcyan("sorted_vals(c) = %v\n", vc)
	chk.Strings(tst, "kc", kc, []string{"a", "b", "c", "d", "x", "y", "z"})
	chk.Bools(tst, "vc", vc, []bool{false, false, false, true, true, true, true})
}

func Test_sort06(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sort06. int => ??? maps")

	a := map[int]bool{100: true, 101: false, 102: true, 10: false, 9: true, 8: false, 0: true}
	k := IntBoolMapSort(a)
	io.Pforan("sorted_keys(a) = %v\n", k)
	chk.Ints(tst, "k", k, []int{0, 8, 9, 10, 100, 101, 102})
}

func Test_sort07(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sort07. sort 3 ints")

	x := []int{0, 10, 1, 3}
	IntSort3(&x[0], &x[1], &x[2])
	chk.Ints(tst, "sort3(x)", x, []int{0, 1, 10, 3})

	x = []int{0, 10, 1, 3}
	IntSort4(&x[0], &x[1], &x[2], &x[3])
	chk.Ints(tst, "sort4(x)", x, []int{0, 1, 3, 10})
}

func Test_sort08(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sort08. sort pair of slices")

	a := []float64{1, 5, 2, 8, 2, 9, 4, 0, -2}
	Qsort(a)
	chk.Array(tst, "a", 1e-15, a, []float64{-2, 0, 1, 2, 2, 4, 5, 8, 9})

	b := []float64{9, 8, 7, 0, 3, -7, 6, 1, -1}
	Qsort(b)
	chk.Array(tst, "b", 1e-15, b, []float64{-7, -1, 0, 1, 3, 6, 7, 8, 9})

	a = []float64{1, 5, 2, 8, 2, 9, 4, 0, -2}
	b = []float64{9, 8, 7, 0, 3, -7, 6, 1, -1}
	Qsort2(a, b)
	chk.Array(tst, "a", 1e-15, a, []float64{-2, 0, 1, 2, 2, 4, 5, 8, 9})
	chk.Array(tst, "b", 1e-15, b, []float64{-1, 1, 9, 7, 3, 6, 8, 0, -7})

	a = []float64{1, 5, 2, 8, 2, 9, 4, 0, -2}
	b = []float64{9, 8, 7, 0, 3, -7, 6, 1, -1}
	Qsort2(b, a)
	chk.Array(tst, "a", 1e-15, a, []float64{9, -2, 8, 0, 2, 4, 2, 5, 1})
	chk.Array(tst, "b", 1e-15, b, []float64{-7, -1, 0, 1, 3, 6, 7, 8, 9})

	var s Sorter

	a = []float64{1, 5, 2, 8, 2, 9, 4, 0, -2}
	s.Init(3, func(i, j int) bool { return a[i] < a[j] })
	chk.Ints(tst, "indx", s.Index, []int{0, 2, 1}) // 3 => smaller set

	a = []float64{1, 5, 2, 8, 2, 9, 4, 0, -2}
	as := s.GetSorted(a)
	chk.Array(tst, "as", 1e-15, as, []float64{1, 2, 5}) // smaller set

	b = []float64{9, 8, 7, 0, 3, -7, 6, 1, -1}
	s.Init(5, func(i, j int) bool { return b[i] < b[j] }) // 5 => smaller set
	chk.Ints(tst, "indx", s.Index, []int{3, 4, 2, 1, 0})

	b = []float64{9, 8, 7, 0, 3, -7, 6, 1, -1}
	bs := s.GetSorted(b)
	chk.Array(tst, "bs", 1e-15, bs, []float64{0, 3, 7, 8, 9}) // smaller set

	a = []float64{1, 5, 2, 8, 2, 9, 4, 0, -2}
	s.Init(len(a), func(i, j int) bool { return a[i] < a[j] })
	chk.Ints(tst, "indx", s.Index, []int{8, 7, 0, 4, 2, 6, 1, 3, 5})

	a = []float64{1, 5, 2, 8, 2, 9, 4, 0, -2}
	b = []float64{9, 8, 7, 0, 3, -7, 6, 1, -1}
	as = s.GetSorted(a)
	bs = s.GetSorted(b)
	chk.Array(tst, "as", 1e-15, as, []float64{-2, 0, 1, 2, 2, 4, 5, 8, 9})
	chk.Array(tst, "bs", 1e-15, bs, []float64{-1, 1, 9, 3, 7, 6, 8, 0, -7})

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	cs := s.GetSortedI(c)
	chk.Ints(tst, "cs", cs, []int{9, 8, 1, 5, 3, 7, 2, 4, 6})
}
