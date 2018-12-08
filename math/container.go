// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	. "github.com/GoLangsam/twos/core"
)

// ===========================================================================

// Container can tell for any item whether it contains this item or not.
type Container interface {
	Contains(item interface{}) bool
}

// ===========================================================================

// IsInSome returns a function to discriminate
// whether some of the containers contains a given Pair.
// If there are no containers the returned function will evaluate to false always.
func IsInSome(containers ...Container) (pairIs func(Pair) bool) {

	return func(a Pair) bool {
		for _, container := range containers {
			if container.Contains(a) {return true}
		}
		return false
	}
}

// IsInNone returns a function to discriminate
// whether none of the containers contains a given Pair.
// If there are no containers the returned function will evaluate to true always.
func IsInNone(containers ...Container) (pairIs func(Pair) bool) {

	return func(a Pair) bool {
		for _, container := range containers {
			if container.Contains(a) {return false}
		}
		return true
	}
}

// ===========================================================================
