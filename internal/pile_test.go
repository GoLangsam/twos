// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"testing"

	"github.com/mauricelam/genny/generic"
)

// let's fool genny
func dummytestanyType(t *testing.T) {
	t.Helper()
	type anyType generic.Type
}

func ExamplePileOfanyType_interface() {

	// pairOfanyType is implemented by onesOfanyType and PileOfanyType.
	type pairOfanyType interface {
		Pair
		Tail() Tail                // Iterable
		Name() string              // Named
		String() string            // Stringer
		Of(Index) Head             // Indexed
		Contains(interface{}) bool // Container
	}

	// lookerOfanyType is implemented by lookUpanyType
	type lookerOfanyType interface {
		Idx(item anyType) (idx Index, found bool) // returns the index of item iff applicable
		Len() int                                 // current Length
		Length() Cardinality                      // current Length
		Random() <-chan anyType                   // a la Range, just: in random order
		put(item anyType, idx Index)              // may panic("Overflow")
	}

	// pilerOfanyType is implemented by *PileOfanyType
	type pilerOfanyType interface {
		pairOfanyType

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

	var (
		_, _, _ pairOfanyType = onesOfanyType{}, &onesOfanyType{}, new(onesOfanyType)

		_, _ lookerOfanyType = &lookUpanyType{}, new(lookUpanyType)
		_, _ pilerOfanyType  = &PileOfanyType{}, new(PileOfanyType)
	)
}

func ExampleNewPileOfanyType() {
	name, data := testanyType()
	pile := NewPileOfanyType(name, data...)
	fmt.Println(pile.Length())
	// Output:
	// 5
}

func ExamplePileOfanyType_Length() {
	name, data := testanyType()
	pile := NewPileOfanyType(name, data...)
	fmt.Println(pile.Length())

	var i Cardinality
	for head, tail := pile.Tail()(); head != nil; head, tail = tail() {
		i++
		aten, apep := head().Both()
		_, _ = aten, apep
		// fmt.Println(aten) // TODO: not easy to print this generic
		// fmt.Println(IsAtomAten(head()), IsAtomApep(head()))
	}

	fmt.Println(pile.Length(), "==", i)

	// Output:
	// 5
	// 5 == 5
}

func ExamplePileOfanyType_Both() {
	name, data := testanyType()
	pile := NewPileOfanyType(name, data...)
	fmt.Println(pile.Length())

	aten, apep := pile.Both()

	_, _ = aten, apep // TODO: not easy to print this generic
	// Output:
	// 5
}

func ExamplePileOfanyType_At() {
	name, data := testanyType()
	pile := NewPileOfanyType(name, data...)
	fmt.Println(pile.Length())

	var i Index
	for i = 1; i < 4; i++ {
		item := data[int(i)-1]
		if pile.At(i) != item {
			fmt.Println("At failed", i, pile.At(i), item)
		}
	}
	// Output:
	// 5
}

func ExamplePileOfanyType_Idx() {
	name, data := testanyType()
	pile := NewPileOfanyType(name, data...)
	fmt.Println(pile.Length())

	var i Index
	for i = 1; i < 4; i++ {
		item := data[int(i)-1]

		idx, ok := pile.Idx(item)
		if !ok {
			fmt.Println("Idx failed to find", item, "at index", i)
		} else if idx != i {
			fmt.Println("Idx returned", idx, "instead of", i)
		}
	}
	// Output:
	// 5
}

func ExamplePileOfanyType_Range() {
	name, data := testanyType()
	pile := NewPileOfanyType(name, data...)
	fmt.Println(pile.Length())

	i := 0
	for range pile.Range() {
		i++
	}

	fmt.Println(pile.Length(), "==", i)
	// Output:
	// 5
	// 5 == 5
}

func ExamplePileOfanyType_Random() {
	name, data := testanyType()
	pile := NewPileOfanyType(name, data...)
	fmt.Println(pile.Length())

	i := 0
	for range pile.Random() {
		i++
	}

	fmt.Println(pile.Length(), "==", i)
	// Output:
	// 5
	// 5 == 5
}
