// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// X returns a Tail to the cartesian product of the factors
// presented as right-associated nested pairs of pairs of items.
//
// Heads of X(I1, I2, I3, I4) will look like { i1 | { i2 | { i3 | i4 }}}.
//
// No factors evaluates to the NilTail.
// A single factor evaluates to its Tail.
func X(factors ...Iterable) (tail Tail) {
	l := len(factors)
	switch {
	case l < 1: return NilTail()
	case l < 2: return factors[0].Tail()
	case l < 3: return mult(factors[0], factors[1])
	default:    return mult(factors[0], X(factors[1:]...))
	}
}

// mult implements the Cartesian product.
func mult(a, b Iterable) (tail Tail) {
	if a == nil || b == nil { return NilTail() }

	aHead, _ := a.Tail()()
	aTail, bTail, reset := a.Tail(), NilTail(), b.Tail()

	tail = func() (Head, Tail) { return prod(aHead, aTail, bTail, reset) }
	return
}

// prod implements the recursive part of the Cartesian product.
func prod(aHead Head, aTail, bTail, reset Tail) (head Head, tail Tail) {
	bHead, bTail := bTail()
	if bHead == nil {
		aHead, aTail = aTail() // advance a
		bHead, bTail = reset() // reset b
	}

	if aHead == nil || bHead == nil { return NilTail()() }

	head = func() Pair         { return Join(aHead, bHead) }
	tail = func() (Head, Tail) { return prod(aHead, aTail, bTail, reset) }
	return
}
