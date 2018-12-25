// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"math/big"
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
type ordinalNumber struct { *big.Int }

// Ordinal returns an Index representing the number a.
func Ordinal(a int) Index {return Index{big.NewInt(int64(a))}}

// ===========================================================================

// Cardinality represents a cardinal number such as the #-of items in a Pile.
type Cardinality = cardinalNumber
type cardinalNumber struct { *big.Int }

// Cardinal returns a Cardinality representing the number a.
func Cardinal(a int) Cardinality {return Cardinality{big.NewInt(int64(a))}}

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
func(a *Index)Unit() Index {return Ordinal(1)}

// Unit returns the unit value
// of type Cardinality which
// represents the cardinality of the empty set.
// It resembles Zero/Null - the neutral element of addition.
func(a *Cardinality)Unit() Cardinality {return Cardinal(0)}

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
	return a.Int.Cmp(b.Int)
}

// Cmp compares a and b and returns:
//
//  -1 if a <  b
//   0 if a == b
//  +1 if a >  b
func(a Cardinality)Cmp(b Cardinality) (r int) {
	return a.Int.Cmp(b.Int)
}

// ===========================================================================

// AsOffset returns the slice-offset corresponding to ordinal number a: a-1.
// If a-1 exceeds maxInt -1 is returned.
func (a Index) AsOffset() int {
	i := Ordinal(-1)
	i = i.Add(a, i)
	return i.AsInt()
}

// At returns the Index corresponding to slice-offset i: i+1
func At(i int) Index {
	idx := Ordinal(i) 
	return idx.Add(idx, Ordinal(1))
}

// ===========================================================================
// const maxInt = int(^uint(0) >> 1)

// AsInt returns a as an int, or -1 iff a exceeds maxInt
func (a Index) AsInt() int {
	if !a.IsInt64() {
		return -1
	}
	max := Ordinal(int(^uint(0) >> 1)) // maxInt
	if a.Cmp(max) > 0 {
		return -1
	}
	return int(a.Int64())
}

// AsInt returns a as an int, or -1 iff a exceeds maxInt
func (a Cardinality) AsInt() int {
	if !a.IsInt64() {
		return -1
	}
	max := Cardinal(int(^uint(0) >> 1)) // maxInt
	if a.Cmp(max) > 0 {
		return -1
	}
	return int(a.Int64())
}

// ===========================================================================

// Add sets a to the sum x+y and returns a.
func (a Index) Add(x, y Index) Index {

	a.Int = a.Int.Add(x.Int, y.Int) 
	return a
}

// Add sets a to the sum x+y and returns a.
func (a Cardinality) Add(x, y Cardinality) Cardinality {

	a.Int = a.Int.Add(x.Int, y.Int) 
	return a
}

// Mul sets a to the product x*y and returns a.
func (a Cardinality) Mul(x, y Cardinality) Cardinality {

	a.Int = a.Int.Mul(x.Int, y.Int) 
	return a
}

// ===========================================================================
// <eom>
