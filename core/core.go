// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"reflect"
)

// ===========================================================================

// ID as a unique identifier - unique among its Kind.
// Implements Named.
type ID = name

// name as a unique identifier - unique among its Kind.
type name string

// ===========================================================================

// Index represents an ordinal number and may be used
// to enumerate a collention or access some item.
// Intentionally, the first item has Index = 1.
// This is more intuitive for users.
// (Well, programmers prefer offsets over ordinals).
type Index = ordinalNumber
type ordinalNumber uint64

// ===========================================================================

// Cardinality represents a cardinal number such as the #-of items in a Pile.
type Cardinality = cardinalNumber
type cardinalNumber uint64

// ===========================================================================
// Interfaces

// Type is the reflect.Type
type Type = reflect.Type

// TypeOf returns the Type of a
var TypeOf = reflect.TypeOf

// Pair has two sides: Aten & Apep. It may be atomic, or composed, nested.
type Pair interface {
	Both() (aten, apep interface{}) // both sides - whatever each type is
}

// Iterable represents a tailor-made collection, a thread for a tailor.
type Iterable interface {
	Tail() Tail
}

// ===========================================================================

// Unit returns the unit value
// of type ID which
// resembles an empty string - the neutral element of concatination.
func(a *name)Unit() name {return ID("")}

// Unit returns the unit value
// of type Index which
// resembles One (1) - the neutral element of multiplication.
func(a *Index)Unit() Index {return Index(1)}

// Unit returns the unit value
// of type Cardinality which
// represents the cardinality of the empty set.
// It resembles Zero/Null - the neutral element of addition.
func(a *Cardinality)Unit() Cardinality {return Cardinality(0)}

// Cmp compares a and b and returns:
//
//  -1 if a <  b
//   0 if a == b
//  +1 if a >  b
func(a name)Cmp(b name) (r int) {
	switch {
	case a < b:
		return -1
	case a > b:
		return +1
	default:
		return 0
	}
}

// Cmp compares a and b and returns:
//
//  -1 if a <  b
//   0 if a == b
//  +1 if a >  b
func(a Index)Cmp(b Index) (r int) {
	switch {
	case a < b:
		return -1
	case a > b:
		return +1
	default:
		return 0
	}
}

// Cmp compares a and b and returns:
//
//  -1 if a <  b
//   0 if a == b
//  +1 if a >  b
func(a Cardinality)Cmp(b Cardinality) (r int) {
	switch {
	case a < b:
		return -1
	case a > b:
		return +1
	default:
		return 0
	}
}

// ===========================================================================

// AsOffset returns the slice-offset corresponding to ordinal number i: i-1.
func (a Index) AsOffset() int {
	return int(a - 1)
}

// At returns the Index corresponding to slice-offset i: i+1
func At(i int) Index {
	return Index(i + 1)
}

