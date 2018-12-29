// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// IsLess returns a Predicate
// which is useful to disriminate
// if some Pair is less than a or not.
func(a Index)IsLess() Predicate {
	return func (p Pair) (less bool) {
		if x, less := p.(Index); less {
			return x.Int.Cmp(a.Int) < 0
		}
		return
	}
}

// IsEq returns a Predicate
// which is useful to disriminate
// if some Pair is equal to a or not.
func(a Index)IsEq() Predicate {
	return func (p Pair) (equal bool) {
		if x, equal := p.(Index); equal {
			return x.Int.Cmp(a.Int) == 0
		}
		return
	}
}

// IsNul returns a Predicate
// which is useful to disriminate
// if some Pair represents the zero value.
func(*Index)IsNul() Predicate {
	return func (p Pair) (isNul bool) {
		if x, isNul := p.(Index); isNul {
			return x.IsEq()(Ordinal(0))
		}
		return
	}
}

// IsOne returns a Predicate
// which is useful to disriminate
// if some Pair represents the multiplicative unit.
func(*Index)IsOne() Predicate {
	return func (p Pair) (isOne bool) {
		if x, isOne := p.(Index); isOne {
			return x.IsEq()(Ordinal(1))
		}
		return
	}
}

// IsEqInt returns a Predicate
// which is useful to disriminate
// if some int number is equal to a or not.
//
// IsEqInt is a conveninence method - syntactical sugar.
func(a Index)IsEqInt() func(int) bool {
	return func (p int) bool {
		return a.IsEq()(Ordinal(p))
	}
}

// ===========================================================================
