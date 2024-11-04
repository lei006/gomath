// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rnd implements random numbers generators (wrapping the standard functions or the Mersenne
// Twister library). It also implements probability distribution functions.
package rnd

import (
	"math/rand"
	"time"

	"github.com/lei006/gomath/utl"
)

// Init initializes random numbers generators
//
//	Input:
//	 seed -- seed value; use seed <= 0 to use current time
func Init(seed int) {
	if seed <= 0 {
		seed = int(time.Now().Unix())
	}
	rand.Seed(int64(seed))
}

// Int generates pseudo random integer between low and high.
//
//	Input:
//	 low  -- lower limit
//	 high -- upper limit
//	Output:
//	 random integer
func Int(low, high int) int {
	return rand.Int()%(high-low+1) + low
}

// Ints generates pseudo random integers between low and high.
//
//	Input:
//	 low    -- lower limit
//	 high   -- upper limit
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Ints(values []int, low, high int) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = Int(low, high)
	}
}

// Float64 generates a pseudo random real number between low and high; i.e. in [low, right)
//
//	Input:
//	 low  -- lower limit (closed)
//	 high -- upper limit (open)
//	Output:
//	 random float64
func Float64(low, high float64) float64 {
	return low + (high-low)*rand.Float64()
}

// Float64s generates pseudo random real numbers between low and high; i.e. in [low, right)
//
//	Input:
//	 low  -- lower limit (closed)
//	 high -- upper limit (open)
//	Output:
//	 values -- slice to be filled with len(values) numbers
func Float64s(values []float64, low, high float64) {
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*rand.Float64()
	}
}

// FlipCoin generates a Bernoulli variable; throw a coin with probability p
func FlipCoin(p float64) bool {
	if p == 1.0 {
		return true
	}
	if p == 0.0 {
		return false
	}
	if rand.Float64() <= p {
		return true
	}
	return false
}

// IntGetUnique randomly selects n items in a list avoiding duplicates
//
//	Note: using the 'reservoir sampling' method; see Wikipedia:
//	      https://en.wikipedia.org/wiki/Reservoir_sampling
func IntGetUnique(values []int, n int) (selected []int) {
	if n < 1 {
		return
	}
	if n >= len(values) {
		return IntGetShuffled(values)
	}
	selected = make([]int, n)
	for i := 0; i < n; i++ {
		selected[i] = values[i]
	}
	var j int
	for i := n; i < len(values); i++ {
		j = rand.Intn(i + 1)
		if j < n {
			selected[j] = values[i]
		}
	}
	return
}

// IntGetUniqueN randomly selects n items from start to endp1-1 avoiding duplicates
//
//	Note: using the 'reservoir sampling' method; see Wikipedia:
//	      https://en.wikipedia.org/wiki/Reservoir_sampling
func IntGetUniqueN(start, endp1, n int) (selected []int) {
	if n < 1 {
		return
	}
	size := endp1 - start
	if n >= size {
		selected = utl.IntRange2(start, endp1)
		IntShuffle(selected)
		return
	}
	selected = make([]int, n)
	for i := 0; i < n; i++ {
		selected[i] = start + i
	}
	var j int
	for i := n; i < size; i++ {
		j = rand.Intn(i + 1)
		if j < n {
			selected[j] = start + i
		}
	}
	return
}

// IntShuffle shuffles a slice of integers
func IntShuffle(values []int) {
	var j, tmp int
	for i := len(values) - 1; i > 0; i-- {
		j = rand.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// IntGetShuffled returns a shufled slice of integers
func IntGetShuffled(values []int) (shuffled []int) {
	shuffled = make([]int, len(values))
	copy(shuffled, values)
	IntShuffle(shuffled)
	return
}

// Shuffle shuffles a slice of float point numbers
func Shuffle(values []float64) {
	var tmp float64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = rand.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// IntGetGroups randomly selects indices from pool separating them in groups
//
//	Input:
//	  pool -- all ints.
//	Output:
//	  groups -- [ngroups][size_of_group] pre-allocated slices
func IntGetGroups(groups [][]int, pool []int) {
	ngroups := len(groups)
	sizeg := len(groups[0])
	indices := IntGetShuffled(pool)
	var k int
	for i := 0; i < ngroups; i++ {
		for j := 0; j < sizeg; j++ {
			groups[i][j] = indices[k]
			k++
		}
	}
}
