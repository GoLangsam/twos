// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// HeadS represents a series of Head.
type HeadS []Head

// TailS represents a series of Tail.
type TailS []Tail

// Both implements Pair
// by returning both parts of a,
// a slice of the first item and a slice of the remaining items.
func (a HeadS) Both() (aten, apep interface{}) {
	if len(a) > 0 {
	return a[:1], a[1:] }
	return a[:0], a[0:] }

// Both implements Pair
// by returning both parts of a,
// a slice of the first item and a slice of the remaining items.
func (a TailS) Both() (aten, apep interface{}) {
	if len(a) > 0 {
	return a[:1], a[1:] }
	return a[:0], a[0:] }

/*
// String implements fmt.Stringer.
func (a HeadS) String() string {
	if len(a) > 0 {
	return StringOfTwos(a[:1], a[1:]) }
	return StringOfTwos("<nilHead>", "[<nilHead>...]") }

// String implements fmt.Stringer.
func (a TailS) String() string {
	if len(a) > 0 {
	return StringOfTwos(a[:1], a[1:]) }
	return StringOfTwos("<nilTail>", "[<nilTail>...]") }
*/
// ===========================================================================

// Len reports the length.
func (a HeadS) Len() int { return len(a) }

// Len reports the length.
func (a TailS) Len() int { return len(a) }

// Length implements Pile by returning the length.
func (a HeadS) Length() Cardinality { return Cardinality(len(a)) }

// Length implements Pile by returning the length.
func (a TailS) Length() Cardinality { return Cardinality(len(a)) }

// ===========================================================================
