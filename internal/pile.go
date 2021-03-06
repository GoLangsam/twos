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
func (a lookUpanyType) Size() Cardinality                        { return Cardinal(len(a.look)) }

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
	duplicates map[anyType]Cardinality
}

// NewPileOfanyType returns a named Pile of items.
// There must be at least one item.
func NewPileOfanyType(name string, items ...anyType) *PileOfanyType {
	soManyItems := len(items)

	pile := &PileOfanyType{
		ID(name),
		make([]anyType, 0, soManyItems),
		lookUpanyType{make(map[anyType]Index, soManyItems)},
		make(map[anyType]Cardinality),
	}
	pile = pile.append(items...)
	return pile
}

// NewanyTypePile returns a named Pile of items.
// given an accessible finite channel of them.
func NewanyTypePile(name string, items <-chan anyType) *PileOfanyType {

	pile := &PileOfanyType{
		ID(name),
		[]anyType{},
		lookUpanyType{make(map[anyType]Index)},
		make(map[anyType]Cardinality),
	}
	for item := range items {
		pile = pile.append(item)
	}
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
func (a anyTypeS) Both() (aten, apep interface{}) {
	if len(a) > 0 {
	return a[:1], a[1:] }
	return nil, nil }

// String implements fmt.Stringer.
func (a onesOfanyType) String() string { return StringOfTwos(a.ID, a.Apep) }

// String implements fmt.Stringer.
func (a PileOfanyType) String() string { return StringOfTwos(a.ID, a.anyTypeS) }

// ===========================================================================

// Len reports the length.
func (a anyTypeS) Len() int { return len(a) }

// Size implements Pile by returning the length.
func (a anyTypeS) Size() Cardinality { return Cardinal(len(a)) }

// Size implements Pile by returning 1.
func (a onesOfanyType) Size() Cardinality { return Cardinal(1) }

// Len reports the length.
func (a PileOfanyType) Len() int { return a.anyTypeS.Len() }
// Note: Inherited both from anyTypeS and from lookUpanyType

// Size implements Pile by returning the length of the Pile.
func (a PileOfanyType) Size() Cardinality { return a.lookUpanyType.Size() }

// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a onesOfanyType) Of(index Index) Head {
	if index.IsOne()(index) {
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
		nilHead, _ := NilTail()()
		return nilHead()
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
// The Cardinality tells how often the item was received as a duplicate.
//   Hint: The len(of this map) tells how manyType different item values were seen more than once.
//   Note: Clients should check and consider any non-zero result as an error.
func (a PileOfanyType) Duplicates() map[anyType]Cardinality { return a.duplicates }

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
		cnt := Cardinal(1)
		if idx, found := a.duplicates[item]; found {
			cnt = cnt.Add(cnt, idx)
		}
		a.duplicates[item] = cnt
	} else {
		a.anyTypeS = append(a.anyTypeS, item)
		a.put(item, Ordinal(len(a.anyTypeS)))
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
		make(map[anyType]Cardinality),
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
		make(map[anyType]Cardinality),
	}
	for i, item := range pile.anyTypeS {
		pile.lookUpanyType.look[item] = Ordinal(i + 1)
	}
	return pile
}

// ===========================================================================
