// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// IsKind returns a function to discriminate
// whether a given Pair is a kind.
func IsKind() (pairIs func(Pair) bool) {

	return func(a Pair) (isKind bool) {
		_, isKind = a.(kind)
		return
	}
}

// IsNested returns a function to discriminate
// whether a given Pair is a nested pair of pairs.
func IsNested() (pairIs func(Pair) bool) {

	return func(a Pair) (isNest bool) {
		_, isNest = a.(nest)
		return
	}
}
