// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package la

import (
	"sort"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
	"github.com/lei006/gomath/utl"
)

// Equations organises the identification numbers (IDs) of equations in a linear system:
// [A] ⋅ {x} = {b}. Some of the components of the {x} vector may be known in advance—i.e. some
// {x} components are "prescribed"—therefore the system of equations must be reduced first
// before its solution. The equations corresponding to the known {x} values are called
// "known equations" whereas the equations corresponding to unknown {x} values are called
// "unknown equations".
//
// The right-hand-side vector {b} will also be partitioned into two sets. However, the
// "unknown equations" part will have known values of {b}. If needed, the other {b}
// values can be computed from the "known equations".
//
// The structure Equations will then hold the IDs of two sets of equations:
// "u" -- unknown {x} values with given {b} values
// "k" -- known {x} values (the corresponding {b} may be post-computed)
//
//	In summary:
//
//	 * We define the u-reduced system of equations with the unknown {x} (and known {b})
//	                 ┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄                       ┄┄┄┄┄┄┄
//	   i.e. the system that needs to be solved is:  [Auu]⋅{xu} = {bu} - [Auk]⋅{xk}
//
//	 * We define the k-reduced system of equations with the known {x} values.
//	                 ┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄                       ┄┄┄┄┄
//
//	 * We call the system [A] ⋅ {x} = {b} the "full" system, with all equations.
//
//	The "full" linear system is:
//
//	                   [A]⋅{x} = {b}   or   [ Auu Auk ]⋅{ xu ? } = { bu ✓ }
//	                                        [ Aku Akk ] { xk ✓ }   { bk ? }
//
//	NOTE: the {bu} part is known and the {bk} can be (post-)computed if needed.
//
//	The partitioned system is symbolized as follows:
//
//	                              Auu Auk  u─────> unknown
//	                              Aku Akk  k─────> known
//	                                u   k
//	                                │   └────────> known
//	                                └────────────> unknown
//
//	The Equations structure uses arrays to map the indices of equations from the "full"
//	system to the reduced systems and vice-versa. We call these arrays "maps" but they
//	are not Go maps; they are simple slices of integers and should give better
//	performance than maps.
//
//	Four "maps" (slices of integers) are built:
//	 * UtoF -- reduced u-system to full-system
//	 * FtoU -- full-system to reduced u-system
//	 * KtoF -- reduced k-system to full-system
//	 * FtoK -- full-system to reduced k-system
//
//
//	EXAMPLE:
//
//	 N  = 9          ⇒  total number of equations
//	 kx = [0  3  6]  ⇒  known x-components (i.e. "prescribed" equations)
//
//	   0  1  2  3  4  5  6  7  8                  0   1  2   3   4  5   6   7  8
//	 0 +--------+--------+------ 0  ◀ known ▶  0 [kk] ku ku [kk] ku ku [kk] ku ku 0
//	 1 |        |        |       1             1  uk  ○  ○   uk  ○  ○   uk  ○  ○  1
//	 2 |        |        |       2             2  uk  ○  ○   uk  ○  ○   uk  ○  ○  2
//	 3 +--------+--------+------ 3  ◀ known ▶  3 [kk] ku ku [kk] ku ku [kk] ku ku 3
//	 4 |        |        |       4             4  uk  ○  ○   uk  ○  ○   uk  ○  ○  4
//	 5 |        |        |       5             5  uk  ○  ○   uk  ○  ○   uk  ○  ○  5
//	 6 +--------+--------+------ 6  ◀ known ▶  6 [kk] ku ku [kk] ku ku [kk] ku ku 6
//	 7 |        |        |       7             7  uk  ○  ○   uk  ○  ○   uk  ○  ○  7
//	 8 |        |        |       8             8  uk  ○  ○   uk  ○  ○   uk  ○  ○  8
//	   0  1  2  3  4  5  6  7  8                  0   1  2   3   4  5   6   7  8
//	   ▲        ▲        ▲
//	 known    known    known                             ○ means uu equations
//
//	 Nu = 6 => number of "unknown equations"
//	 Nk = 3 => number of "known equations"
//	 Nu + Nk = N
//
//	 reduced:0  1  2  3  4  5              Example:
//	 UtoF = [1  2  4  5  7  8]           ⇒ UtoF[3] returns equation # 5 of full system
//
//	         0  1  2  3  4  5  6  7  8
//	 FtoU = [   0  1     2  3     4  5]  ⇒ FtoU[5] returns equation # 3 of reduced system
//	        -1       -1       -1         ⇒ -1 indicates 'value not set'
//
//	 reduced:0  1  2
//	 KtoF = [0  3  6]                    ⇒ KtoF[1] returns equation # 3 of full system
//
//	         0  1  2  3  4  5  6  7  8
//	 FtoK = [0        1        2      ]  ⇒ FtoK[3] returns equation # 1 of reduced system
//	           -1 -1    -1 -1    -1 -1   ⇒ -1 indicates 'value not set'
type Equations struct {

	// essential
	N    int   // total number of equations
	Nu   int   // number of unknowns (size of u-system)
	Nk   int   // number of known values (size of k-system)
	UtoF []int // reduced u-system to full-system
	FtoU []int // full-system to reduced u-system
	KtoF []int // reduced k-system to full system
	FtoK []int // full-system to reduced k-system

	// convenience
	Auu, Auk, Aku, Akk *Triplet // the partitioned system in sparse format
	Duu, Duk, Dku, Dkk *Matrix  // the partitioned system in dense format
	Bu, Bk, Xu, Xk     Vector   // partitioned rhs and unknowns vector
}

// NewEquations creates a new Equations structure
//
//	n  -- total number of equations; i.e. len({x}) in [A]⋅{x}={b}
//	kx -- known x-components ⇒ "known equations" [may be unsorted]
func NewEquations(n int, kx []int) (o *Equations) {
	sortedkx := utl.IntUnique(kx)
	o = new(Equations)
	o.N = n
	o.Nk = len(sortedkx)
	o.Nu = o.N - o.Nk
	if o.N < 1 {
		chk.Panic("the number of equations must be greater than 0. N=%d is invalid\n", o.N)
	}
	if o.Nu <= 0 {
		chk.Panic("at least one unknown equation is required. Nu=%d is invalid. Nk=%d. N=%d\n", o.Nu, o.Nk, o.N)
	}
	for _, eq := range kx {
		if eq < 0 || eq >= o.N {
			chk.Panic("known equation number is out of bounds. eq=%d must be in [0,%d]\n", eq, o.N-1)
		}
	}
	o.UtoF = make([]int, o.Nu)
	o.FtoU = make([]int, o.N)
	o.KtoF = make([]int, o.Nk)
	o.FtoK = make([]int, o.N)
	var iu, ik int
	for eq := 0; eq < n; eq++ {
		o.FtoU[eq] = -1 // -1 indicates "not set"
		o.FtoK[eq] = -1
		idx := sort.SearchInts(sortedkx, eq)
		if idx < len(sortedkx) && sortedkx[idx] == eq { // found known x-component
			o.KtoF[ik] = eq
			o.FtoK[eq] = ik
			ik++
		} else { // not found ⇒ unknown x-component
			o.UtoF[iu] = eq
			o.FtoU[eq] = iu
			iu++
		}
	}
	return
}

// Alloc allocates the A matrices in triplet format (sparse format)
//
//	INPUT:
//	  nnz -- total number of nonzeros in each part [nnz(Auu), nnz(Auk), nnz(Aku), nnz(Akk)]
//	         nnz may be nil, in this case, the following values are assumed:
//	                  nnz(Auu) = Nu ⋅ Nu
//	                  nnz(Auk) = Nu ⋅ Nk
//	                  nnz(Aku) = Nk ⋅ Nu
//	                  nnz(Akk) = Nk ⋅ Nk
//	         Thus, memory is wasted as the size of a fully dense system is considered.
//	 kparts -- also allocates Aku and Akk
//	 vectors -- also allocates the partitioned vectors Bu, Bk, Xu, Xk
//
// OUTPUT:
//
//	The partitioned system (Auu, Auk, Aku, Akk) is stored as member of this object
func (o *Equations) Alloc(nnz []int, kparts, vectors bool) {
	if nnz == nil {
		nnz = []int{o.Nu * o.Nu, o.Nu * o.Nk, o.Nk * o.Nu, o.Nk * o.Nk}
	}
	o.Auu = NewTriplet(o.Nu, o.Nu, nnz[0])
	o.Auk = NewTriplet(o.Nu, o.Nk, nnz[1])
	if kparts {
		o.Aku = NewTriplet(o.Nk, o.Nu, nnz[2])
		o.Akk = NewTriplet(o.Nk, o.Nk, nnz[3])
	}
	if vectors {
		o.Bu = NewVector(o.Nu)
		o.Bk = NewVector(o.Nk)
		o.Xu = NewVector(o.Nu)
		o.Xk = NewVector(o.Nk)
	}
}

// Start (re)starts index for inserting items using the Put command
func (o *Equations) Start() {
	o.Auu.Start()
	o.Auk.Start()
	if o.Aku != nil {
		o.Aku.Start()
		o.Akk.Start()
	}
}

// Put puts component into the right place in partitioned Triplet (Auu, Auk, Aku, Akk)
//
//	NOTE: (1) I and J are the equation numbers in the FULL system
//	      (2) Aku and Akk are ignored if the "kparts" have not been allocated
func (o *Equations) Put(I, J int, value float64) {
	i := o.FtoU[I]
	j := o.FtoU[J]
	if i >= 0 { // u-row
		if j >= 0 { // u-column
			o.Auu.Put(i, j, value)
			return
		}
		j = o.FtoK[J] // k-column
		o.Auk.Put(i, j, value)
		return
	}
	if o.Aku == nil {
		return
	}
	i = o.FtoK[I] // k-row
	if j >= 0 {   // u-column
		o.Aku.Put(i, j, value)
		return
	}
	j = o.FtoK[J] // k-column
	o.Akk.Put(i, j, value)
}

// GetAmat returns the full A matrix (sparse/triplet format) made by Auu, Auk, Aku and Akk
// (e.g. for debugging)
func (o *Equations) GetAmat() (A *Triplet) {
	nnz := o.Auu.max + o.Auk.max + o.Aku.max + o.Akk.max
	A = NewTriplet(o.N, o.N, nnz)
	for k := 0; k < o.Auu.pos; k++ {
		A.Put(o.UtoF[o.Auu.i[k]], o.UtoF[o.Auu.j[k]], o.Auu.x[k])
	}
	for k := 0; k < o.Auk.pos; k++ {
		A.Put(o.UtoF[o.Auk.i[k]], o.KtoF[o.Auk.j[k]], o.Auk.x[k])
	}
	for k := 0; k < o.Aku.pos; k++ {
		A.Put(o.KtoF[o.Aku.i[k]], o.UtoF[o.Aku.j[k]], o.Aku.x[k])
	}
	for k := 0; k < o.Akk.pos; k++ {
		A.Put(o.KtoF[o.Akk.i[k]], o.KtoF[o.Akk.j[k]], o.Akk.x[k])
	}
	return
}

// JoinVector joins unknown with known parts of vector
//
//	INPUT:
//	 bu, bk -- partitioned vectors; e.g. o.Bu, and o.Bk or o.Xu, o.Xk
//	OUTPUT:
//	 b -- pre-allocated b vector such that b = join(bu,bk)
func (o *Equations) JoinVector(b, bu, bk Vector) {
	for i, I := range o.UtoF {
		b[I] = bu[i]
	}
	for i, I := range o.KtoF {
		b[I] = bk[i]
	}
}

// SplitVector splits full vector b into known and unknown parts
//
//	INPUT:
//	 b -- full b vector. b = join(bu,bk)
//	OUTPUT:
//	 bu, bk -- partitioned vectors; e.g. o.Bu, and o.Bk or o.Xu, o.Xk
func (o *Equations) SplitVector(bu, bk, b Vector) {
	for i, I := range o.UtoF {
		bu[i] = b[I]
	}
	for i, I := range o.KtoF {
		bk[i] = b[I]
	}
}

// AllocDense allocates the A matrices in dense format
//
//	INPUT:
//	 kparts -- also allocates Aku and Akk
//	OUTPUT:
//	 The partitioned system (Duu, Duk, Dku, Dkk) is stored as member of this object
func (o *Equations) AllocDense(kparts bool) {
	o.Duu = NewMatrix(o.Nu, o.Nu)
	o.Duk = NewMatrix(o.Nu, o.Nk)
	if kparts {
		o.Dku = NewMatrix(o.Nk, o.Nu)
		o.Dkk = NewMatrix(o.Nk, o.Nk)
	}
}

// SetDense allocates and sets partitioned system in dense format
//
//	kparts -- also computes Aku and Akk
func (o *Equations) SetDense(A *Matrix, kparts bool) {
	o.AllocDense(kparts)
	for i, I := range o.UtoF {
		for j, J := range o.UtoF {
			o.Duu.Set(i, j, A.Get(I, J)) // uu
		}
		for j, J := range o.KtoF {
			o.Duk.Set(i, j, A.Get(I, J)) // uk
		}
	}
	if kparts {
		for i, I := range o.KtoF {
			for j, J := range o.UtoF {
				o.Dku.Set(i, j, A.Get(I, J)) // ku
			}
			for j, J := range o.KtoF {
				o.Dkk.Set(i, j, A.Get(I, J)) // kk
			}
		}
	}
}

// Solve solves linear system (represented by sparse matrices)
//
//	Solve:               {xu} = [Auu]⁻¹ ⋅ ( {bu} - [Auk]⋅{xk} )
//	and, if Aku != nil:  {bk} = [Aku]⋅{xu} + [Akk]⋅{xk}
//
//	Input:
//	 solver -- a pre-configured SparseSolver
//	 t      -- [optional] a scalar (e.g. time) to be used with calcXk and calcBu
//	 calcXk -- [optional] a function to calculate the known values of X
//	 calcBu -- [optional] a function to calculate the right-hand-side Bu
//
//	NOTE: (1) the following must be computed already:
//	             [Auu], [Auk] (optionally, [Aku] and [Akk] as well)
//	      (2) if calcXk is nil, the current values in Xk will be used
//	      (3) if calcBu is nil, the current values in Bu will be used
//	      (4) on exit, {bu} is the modified value {bu} - [Auk]⋅{xk} if [Auk] == !nil
//
//	Instead of providing the functions calcXk and calcBu, the vectors {xk} and {bu} can be
//	pre-computed. For example, in a FDM grid, the following loops can be used:
//
//	      // compute known X values (e.g. Dirichlet boundary conditions)
//	      for i, I := range KtoF { // i:known, I:full
//	          xk[i] = dirichlet(node(I), time)
//	      }
//
//	      // compute RHS
//	      for i, I := range UtoF { // i:unknown, I:full
//	          bu[i] = source(node(I), time)
//	      }
func (o *Equations) Solve(solver SparseSolver, t float64, calcXk, calcBu func(I int, t float64) float64) {

	// compute known X values (e.g. Dirichlet boundary conditions)
	if calcXk != nil {
		for i, I := range o.KtoF { // i:known, I:full
			o.Xk[i] = calcXk(I, t)
		}
	}

	// compute RHS
	if calcBu != nil {
		for i, I := range o.UtoF { // i:unknown, I:full
			o.Bu[i] = calcBu(I, t)
		}
	}

	// fix RHS vector: bu -= Auk⋅xk
	if o.Nk > 0 {
		auk := o.Auk.ToMatrix(nil)
		SpMatVecMulAdd(o.Bu, -1.0, auk, o.Xk)
	}

	// solve system
	solver.Solve(o.Xu, o.Bu)

	// calc {bk}
	if o.Nk > 0 && o.Aku != nil {
		aku := o.Aku.ToMatrix(nil)
		akk := o.Akk.ToMatrix(nil)
		SpMatVecMul(o.Bk, 1.0, aku, o.Xu)    // {bk} = [Aku]⋅{xu}
		SpMatVecMulAdd(o.Bk, 1.0, akk, o.Xk) // {bk} += [Akk]⋅{xk}
	}
}

// SolveOnce solves linear system just once; thus allocating and discarding a linear solver
// (umfpack) internally. See method Solve() for more details
func (o *Equations) SolveOnce(calcXk, calcBu func(I int, t float64) float64) {
	s := NewSparseSolver("umfpack")
	defer s.Free()
	s.Init(o.Auu, nil)
	s.Fact()
	o.Solve(s, 0, calcXk, calcBu)
}

// Info prints information about Equations
func (o *Equations) Info(full bool) {
	io.Pf("number of unknown x-components: Nu = %d\n", o.Nu)
	io.Pf("number of known x-components:   Nk = %d\n", o.Nk)
	io.Pf("total number of equations:      N  = %d\n", o.N)
	if full {
		io.Pf("reduced u-system to full-system map:\nUtoF =\n")
		io.Pf("%v\n", o.UtoF)
		io.Pf("full-system to reduced u-system map:\nFtoU =\n")
		io.Pf("%v\n", o.FtoU)
		io.Pf("reduced k-system to full system map:\nKtoF =\n")
		io.Pf("%v\n", o.KtoF)
		io.Pf("full-system to reduced k-system map:\nFtoK =\n")
		io.Pf("%v\n", o.FtoK)
	}
}
