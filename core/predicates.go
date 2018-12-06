// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// IsKind returns a predicate to discriminate
// whether a given Pair is a kind.
func IsKind() (pairIs func(Pair) bool) {

	return func(a Pair) (isKind bool) {
		_, isKind = a.(kind)
		return
	}
}

// IsNested returns a predicate to discriminate
// whether a given Pair is a nested pair of pairs.
func IsNested() (pairIs func(Pair) bool) {

	return func(a Pair) (isNest bool) {
		_, isNest = a.(nest)
		return
	}
}

// IsEqual returns a predicate to discriminate
// whether a given Pair is equal to the original Pair a.
func IsEqual(a Pair) (pairIs func(Pair) bool) {
	aten, apep := a.Both()
	return func(item Pair) (isEqual bool) {
		// TODO: Do we need reflect.DeepEqual here?
		if item == a { return true}
		iten, ipep := item.Both()
		if aten == iten && apep == ipep { return true}
		return false
	}
}
