// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// IsLess returns a Predicate
// which is useful to disriminate
// if some Pair is less than a or not.
func(a Cardinality)IsLess() Predicate {
	return func (p Pair) (less bool) {
		if x, less := p.(Cardinality); less {
			return x.Int.Cmp(a.Int) < 0
		}
		return
	}
}

// IsEq returns a Predicate
// which is useful to disriminate
// if some Pair is equal to a or not.
func(a Cardinality)IsEq() Predicate {
	return func (p Pair) (equal bool) {
		if x, equal := p.(Cardinality); equal {
			return x.Int.Cmp(a.Int) == 0
		}
		return
	}
}

// IsNul returns a Predicate
// which is useful to disriminate
// if some Pair represents the zero value.
func(*Cardinality)IsNul() Predicate {
	return func (p Pair) (isNul bool) {
		if x, isNul := p.(Cardinality); isNul {
			return x.IsEq()(Cardinal(0))
		}
		return
	}
}

// IsOne returns a Predicate
// which is useful to disriminate
// if some Pair represents the multiplicative unit.
func(*Cardinality)IsOne() Predicate {
	return func (p Pair) (isOne bool) {
		if x, isOne := p.(Cardinality); isOne {
			return x.IsEq()(Cardinal(1))
		}
		return
	}
}

// IsEqInt returns a Predicate
// which is useful to disriminate
// if some int number is equal to a or not.
//
// IsEqInt is a conveninence method - syntactical sugar.
func(a Cardinality)IsEqInt() func(int) bool {
	return func (x int) bool {
		return a.IsEq()(Cardinal(x))
	}
}

// ===========================================================================
