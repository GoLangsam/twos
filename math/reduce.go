// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	. "github.com/GoLangsam/twos/core"
)

// ===========================================================================

// Reduce returns the result of applying f along the iterable.
func Reduce(iter Iterable, f func(Pair, interface{}) interface{}, init interface{}) interface{} {
	for head, tail := iter.Tail()(); head != nil; head, tail = tail() {
		init = f(head(), init)
	}
	return init
}

// ReduceInt returns the result of applying f along the iterable.
func ReduceInt(iter Iterable, f func(Pair, int) int, init int) int {
	for head, tail := iter.Tail()(); head != nil; head, tail = tail() {
		init = f(head(), init)
	}
	return init
}

// ReduceBool returns the result of applying f along the iterable.
func ReduceBool(iter Iterable, f func(Pair, bool) bool, init bool) bool {
	for head, tail := iter.Tail()(); head != nil; head, tail = tail() {
		init = f(head(), init)
	}
	return init
}

// ReducePair returns the result of applying f along the iterable.
func ReducePair(iter Iterable, f func(Pair, Pair) Pair, init Pair) Pair {
	for head, tail := iter.Tail()(); head != nil; head, tail = tail() {
		init = f(head(), init)
	}
	return init
}

// ReduceString returns the result of applying f along the iterable.
func ReduceString(iter Iterable, f func(Pair, string) string, init string) string {
	for head, tail := iter.Tail()(); head != nil; head, tail = tail() {
		init = f(head(), init)
	}
	return init
}

// ReduceCardinality returns the result of applying f along the iterable.
func ReduceCardinality(iter Iterable, f func(Pair, Cardinality) Cardinality, init Cardinality) Cardinality {
	for head, tail := iter.Tail()(); head != nil; head, tail = tail() {
		init = f(head(), init)
	}
	return init
}

// ===========================================================================
