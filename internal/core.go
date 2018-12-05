// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/GoLangsam/twos/core"
)

// ===========================================================================

// ID as a unique identifier - unique among its Kind.
// Implements Named.
type ID = core.ID

// ===========================================================================

// Index represents an ordinal number and may be used
// to enumerate a collention or access some item.
// Intentionally, the first item has Index = 1.
// This is more intuitive for users.
// (Well, programmers prefer offsets over ordinals).
type Index = core.Index

// Indices represents a stream of ordinal numbers
// implemented as a (read-only) channel of Index
// which can be ranged over.
// type Indices = core.Indices

/*
// At returns the Index corresponding to slice-offset i: i+1
func At(i int) Index {
	return core.At(i)
}
*/

// ===========================================================================

// Cardinality represents a cardinal number such as the #-of items in a Pile.
type Cardinality = core.Cardinality

// Cardinalities represents a stream of cardinal numbers
// implemented as a (read-only) channel of Cardinality
// which can be ranged over.
// type Cardinalities = core.Cardinalities

// ===========================================================================

// Type is the reflect.Type
type Type = core.Type

/*
// TypeOf returns the Type of a
func TypeOf(a interface{}) Type {
	return core.TypeOf(a)
}
*/

// ===========================================================================

// Pair has two sides: Aten & Apep. It may be atomic, or composed, nested.
type Pair = core.Pair

// Pairs represents a stream of pairs
// implemented as a (read-only) channel of Pair
// which can be ranged over.
// type Pairs core.Pairs

// ===========================================================================

// Head is a thunk
// which evaluates to
// a Pair.
type Head = core.Head

// Tail is a thunk
// which evaluates to
// a Head and to another Tail.
type Tail = core.Tail

// NilTail is a helper, it returns the Tail which evaluates to nil, nil.
//  Note: this Tail is unique. So its the tail, not a tail.
func NilTail() Tail { return core.NilTail() }

// ===========================================================================
// Interfaces

// Pile holds Length items.
type Pile interface {
	Length() Cardinality
}

// Named represents a named type.
type Named interface {
	Name() string
}

// Indexed allows to retrieve a Head Of any item by Index.
type Indexed interface {
	Of(Index) Head
}

// Iterable represents a tailor-made collection, a thread for a tailor.
type Iterable interface {
	Tail() Tail
}

// Container can tell for any item whether it contains this item or not.
type Container interface {
	Contains(item interface{}) bool
}

// ===========================================================================

// Pairs represents a stream of pairs
// implemented as a (read-only) channel of Pair
// which can be ranged over.
type Pairs <-chan Pair

// ===========================================================================

// StringOfTwos returns the string of two items.
func StringOfTwos(a, b interface{}) string { return core.StringOfTwos(a, b) }

// StringOfOnes returns the string of one item.
func StringOfOnes(a interface{}) string { return core.StringOfOnes(a) }

// StringOfPair returns the string of a Pair
func StringOfPair(a Pair) string { return core.StringOfPair(a) }
