// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// IsLess returns a Predicate
// which is useful to disriminate
// if a is less than some Pair or not.
func (a name) IsLess() Predicate {
	return func(p Pair) (less bool) {
		if x, less := p.(name); less {
			return x > a
		}
		return
	}
}

// IsEq returns a Predicate
// which is useful to disriminate
// if some Pair is equal to a or not.
func (a name) IsEq() Predicate {
	return func(p Pair) (equal bool) {
		if x, equal := p.(name); equal {
			return x == a
		}
		return
	}
}

// IsNul returns a Predicate
// which is useful to disriminate
// if some Pair represents the zero value.
func (*name) IsNul() Predicate {
	return func(p Pair) (isNul bool) {
		if x, isNul := p.(name); isNul {
			return x == ""
		}
		return
	}
}

// ===========================================================================
