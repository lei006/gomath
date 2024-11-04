// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fun (functions) implements special functions such as elliptical, orthogonal polynomials,
// Bessel, discrete Fourier transform, polynomial interpolators, and more.
package fun

import (
	"math"

	"github.com/lei006/gomath/la"
)

// π = 3.141592653589...
const π = math.Pi

// Ss defines a scalar function f(s) of a scalar argument s (scalar scalar)
//
//	Input:
//	  s -- input scalar
//	Returns:
//	  scalar
type Ss func(s float64) float64

// Sss defines a scalar function f(r,s) of two scalar arguments (scalar scalar scalar)
//
//	Input:
//	  s1, s2 -- input scalar
//	Returns:
//	  scalar
type Sss func(s1, s2 float64) float64

// Sv defines a scalar function f(v) of a vector argument v (scalar vector)
//
//	Input:
//	  v -- input vector
//	Returns:
//	  scalar
type Sv func(v la.Vector) float64

// Vs defines a vector function f(s) of a scalar argument s (vector scalar)
//
//	Input:
//	  s -- input scalar
//	Output:
//	  f -- output vector
type Vs func(f la.Vector, s float64)

// Vss defines a vector function f(a,b) of two scalar arguments (vector scalar scalar)
//
//	Input:
//	  a -- first input scalar
//	  b -- second input scalar
//	Output:
//	  f -- output vector
type Vss func(f la.Vector, a, b float64)

// Vvss defines two vector functions u(a,b) and v(a,b) of two scalar arguments
// (vector vector scalar scalar)
//
//	Input:
//	  a -- first input scalar
//	  b -- second input scalar
//	Output:
//	  u -- first output vector
//	  v -- second output vector
type Vvss func(u, v la.Vector, a, b float64)

// Vvvss defines three vector functions u(a,b), v(a,b) and w(a,b) of two scalar arguments
// (vector vector vector scalar scalar)
//
//	Input:
//	  a -- first input scalar
//	  b -- second input scalar
//	Output:
//	  u -- first output vector
//	  v -- second output vector
//	  w -- second output vector
type Vvvss func(u, v, w la.Vector, a, b float64)

// Mss defines a matrix function f(a,b) of two scalar arguments (matrix scalar scalar)
//
//	Input:
//	  a -- first input scalar
//	  b -- second input scalar
//	Output:
//	  m -- output matrix
type Mss func(m *la.Matrix, a, b float64)

// Vv defines a vector function f(v) of a vector argument v (vector vector)
//
//	Input:
//	  v -- input vector
//	Output:
//	  f -- output vector
type Vv func(f, v la.Vector)

// Mv defines a matrix function f(v) of a vector argument v (matrix vector))
//
//	Input:
//	  v -- input vector
//	Output:
//	  f -- output matrix
type Mv func(f *la.Matrix, v la.Vector)

// Mm defines a matrix function f(m) of a matrix argument m (matrix matrix))
//
//	Input:
//	  m -- input matrix
//	Output:
//	  M -- output matrix
type Mm func(f, m *la.Matrix)

// Tv defines a triplet (matrix) function f(v) of a vector argument v (triplet vector)
//
//	Input:
//	  v -- input vector
//	Output:
//	  f -- output triplet
type Tv func(f *la.Triplet, v la.Vector)

// Tt defines a triplet (matrix) function f(t) of a triplet (matrix) argument t (triplet triplet)
//
//	Input:
//	  t -- input triplet
//	Output:
//	  f -- output triplet
type Tt func(f, t *la.Triplet)

// Svs defines a scalar function f(v,s) of a vector and a scalar
//
//	Input:
//	  s -- the scalar
//	  v -- the vector
//	Returns:
//	  scalar
type Svs func(v la.Vector, s float64) float64
