// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/mauricelam/genny/generic"
	"sort"
)

func pileanyType() { // to fool genny
	type anyType generic.Type
}

// ===========================================================================
// generated for anyType:

type onesOfanyType struct {
	ID
	Apep anyType
}

// ===========================================================================

type lookUpanyType struct{ look map[anyType]Index }

// Idx returns the index of item iff applicable.
func (a lookUpanyType) Idx(item anyType) (idx Index, found bool) { idx, found = a.look[item]; return }
func (a *lookUpanyType) put(item anyType, idx Index)             { a.look[item] = idx }
func (a lookUpanyType) Len() int                                 { return len(a.look) }
func (a lookUpanyType) Length() Cardinality                      { return Cardinality(len(a.look)) }

// Random returns a channel to range over.
func (a lookUpanyType) Random() <-chan anyType {
	items := make(chan anyType)

	go func(items chan<- anyType, a lookUpanyType) {
		defer close(items)
		for item := range a.look {
			items <- item
		}
	}(items, a)
	return items
}

// ===========================================================================

// pairOfanyType is implemented by onesOfanyType and PileOfanyType.
type pairOfanyType interface {
	Pair
	Tail() Tail                // Iterable
	Name() string              // Named
	String() string            // Stringer
	Of(Index) Head             // Indexed
	Contains(interface{}) bool // Container
}

// pilerOfanyType is implemented by *PileOfanyType
type pilerOfanyType interface {
	Pair
	Tail() Tail                // Iterable
	Name() string              // Named
	String() string            // Stringer
	Of(Index) Head             // Indexed
	Contains(interface{}) bool // Container

	At(Index) anyType
	Duplicates() map[anyType]Index
	Fmap(f func(anyType) anyType) *PileOfanyType
	Range() <-chan anyType
	S() []anyType
	Sort(less func(i, j int) bool) *PileOfanyType

	lookerOfanyType // Pile: Length()

	add(item anyType) *PileOfanyType
	append(items ...anyType) *PileOfanyType
}

// lookerOfanyType is implemented by lookUpanyType
type lookerOfanyType interface {
	Idx(item anyType) (idx Index, found bool) // returns the index of item iff applicable
	Len() int                                 // current Length
	Length() Cardinality                      // current Length
	Random() <-chan anyType                   // a la Range, just: in random order
	put(item anyType, idx Index)              // may panic("Overflow")
}

// dummy func avoids to see these upon use of go doc -u
func assertPileOfanyTypeInterfaces() {
	var (
		_, _, _ pairOfanyType = onesOfanyType{}, &onesOfanyType{}, new(onesOfanyType)

		_, _ lookerOfanyType = &lookUpanyType{}, new(lookUpanyType)
		_, _ pilerOfanyType  = &PileOfanyType{}, new(PileOfanyType)
	)
}

// ===========================================================================

type anyTypeS []anyType

// PileOfanyType holds different items of some kind, accessible in order of arrival
// and by index as well as via reverse look-up: Idx(item) returns the index (if
// any). It may be seen as a finite ordered indexed immutable set.
//
// Intentionally there is no removal, neither are add nor append exported.
type PileOfanyType struct {
	ID
	anyTypeS

	lookUpanyType
	duplicates lookUpanyType
}

// NewPileOfanyType returns a named Pile of items.
// There must be at least one item.
func NewPileOfanyType(name ID, items ...anyType) *PileOfanyType {
	soManyItems := len(items)

	pile := &PileOfanyType{
		name,
		make([]anyType, 0, soManyItems),
		lookUpanyType{make(map[anyType]Index, soManyItems)},
		lookUpanyType{make(map[anyType]Index)},
	}
	pile = pile.append(items...)
	return pile
}

// ===========================================================================

// Both implements Pair
// by returning both parts of a,
// it's ID and it's anyType value.
func (a onesOfanyType) Both() (aten, apep interface{}) { return a.ID, a.Apep }

// Both implements Pair
// by returning both parts of a,
// it's ID and it's items (as slice of anyType).
func (a PileOfanyType) Both() (aten, apep interface{}) { return a.ID, a.anyTypeS }

// Both implements Pair
// by returning both parts of a,
// a slice of the first item and a slice of the remaining items.
func (a anyTypeS) Both() (aten, apep interface{}) { return a[:0], a[1:] }

// String implements fmt.Stringer.
func (a onesOfanyType) String() string { return StringOfTwos(a.ID, a.Apep) }

// String implements fmt.Stringer.
func (a PileOfanyType) String() string { return StringOfTwos(a.ID, a.anyTypeS) }

// String implements fmt.Stringer.
func (a anyTypeS) String() string { return StringOfTwos(a[:0], a[1:]) }

// ===========================================================================

// Length implements Pile by returning 1.
func (a onesOfanyType) Length() Cardinality { return 1 }

// Length implements Pile by returning the length of the Pile.
func (a PileOfanyType) Length() Cardinality { return a.lookUpanyType.Length() }

func (a onesOfanyType) Of(index Index) Head {
	if index == 1 {
		return func() Pair { return a }
	}
	nilHead, _ := NilTail()()
	return nilHead
}

// ===========================================================================

// Tail implements Iterable
// by returning
// a head which evaluates to a ( head() == Pair(a) ) and
// as tail the unique NilTail().
func (a onesOfanyType) Tail() Tail {
	return func() (Head, Tail) { return func() Pair { return a }, NilTail() }
}

// Tail implements Iterable
// by sucessively returning
// the items
func (a PileOfanyType) Tail() Tail { var idx int; return a.tail(idx) }

func (a PileOfanyType) head(idx int) Head {
	if idx < len(a.anyTypeS) {
		return func() Pair {
			return onesOfanyType{a.ID, a.anyTypeS[idx]}
		}
	}

	return func() Pair {
		head, _ := NilTail()()
		return head()
	}
}

func (a PileOfanyType) tail(idx int) Tail {
	if idx < len(a.anyTypeS) {
		return func() (head Head, tail Tail) {
			head = a.head(idx)
			tail = a.tail(idx + 1)
			return head, tail
		}
	}
	return NilTail()
}

// ===========================================================================

// Contains implements Container
// by telling whether the given item is of suitable type
// and if so, whether a contains this item.
func (a onesOfanyType) Contains(item interface{}) (contains bool) {
	if i, contains := item.(onesOfanyType); contains && i == a {
		return contains
	}
	if i, contains := item.(anyType); contains && i == a.Apep {
		return contains
	}
	return false
}

// Contains implements Container
// by telling whether the given item is of suitable type
// and if so, whether a contains this item.
func (a PileOfanyType) Contains(item interface{}) (contains bool) {
	var i anyType
	if i, contains = item.(anyType); contains {
		_, contains = a.Idx(i)
	}
	return
}

// ===========================================================================
// PileOfanyType specific

// At returns the item at Index idx - and panics iff idx < 1 or > Len().
func (a PileOfanyType) At(idx Index) anyType { return a.anyTypeS[idx.AsOffset()] }

// Of returns a Head to the item at Index idx - and panics iff idx < 1 or > Len().
// Iff the does not implement Pair a nilHead is returned.
func (a PileOfanyType) Of(idx Index) Head { return a.head(idx.AsOffset()) }

// S returns a slice of item in order of arrival.
//  Note: This is also returned by Both() as second parameter.
//  Just: S allows convenient and type-safe access.
func (a PileOfanyType) S() []anyType { return a.anyTypeS }

// Duplicates returns a map of the duplicate items.
// The Index tells how often the item was received as a duplicate.
//   Hint: The len(of this map) tells how manyType different item values were seen more than once.
//   Note: Clients should check and consider anyType non-zero result as an error.
func (a PileOfanyType) Duplicates() map[anyType]Index { return a.duplicates.look }

// append
func (a *PileOfanyType) append(items ...anyType) *PileOfanyType {
	for _, item := range items {
		a = a.add(item)
	}
	return a
}

func (a *PileOfanyType) add(item anyType) *PileOfanyType {
	if idx, duplicate := a.Idx(item); duplicate {
		a.put(item, idx)
		idx, _ := a.duplicates.Idx(item)
		idx++
		a.duplicates.put(item, idx)
	} else {
		a.anyTypeS = append(a.anyTypeS, item)
		a.put(item, Index(len(a.anyTypeS)))
	}
	return a
}

// Range returns a channel to range over.
func (a PileOfanyType) Range() <-chan anyType {
	items := make(chan anyType)

	go func(items chan<- anyType, a PileOfanyType) {
		defer close(items)
		for _, item := range a.anyTypeS {
			items <- item
		}
	}(items, a)
	return items
}

// ===========================================================================

// Fmap returns a new pile with f applied to each item.
func (a PileOfanyType) Fmap(f func(anyType) anyType) *PileOfanyType {
	cap := len(a.anyTypeS)
	pile := &PileOfanyType{
		a.ID,
		make([]anyType, 0, cap),
		lookUpanyType{make(map[anyType]Index, cap)},
		lookUpanyType{make(map[anyType]Index)},
	}
	for _, item := range a.anyTypeS {
		pile = pile.add(f(item))
	}
	return pile
}

// ===========================================================================

// Sort returns a new slice sorted according to less.
func (a anyTypeS) Sort(less func(i, j int) bool) anyTypeS {
	target := make([]anyType, 0, len(a))
	copy(target, a)
	sort.Slice(a, less) // No need for stable sort as there are no duplicates.
	return a
}

// Sort returns another pile sorted according to less.
func (a PileOfanyType) Sort(less func(i, j int) bool) *PileOfanyType {
	cap := len(a.anyTypeS)
	pile := &PileOfanyType{
		a.ID,
		a.anyTypeS.Sort(less),
		lookUpanyType{make(map[anyType]Index, cap)},
		lookUpanyType{make(map[anyType]Index)},
	}
	for i, item := range pile.anyTypeS {
		pile.lookUpanyType.look[item] = Index(i + 1)
	}
	return pile
}

// ===========================================================================
