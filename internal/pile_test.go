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
			fmt.Println("Idx failed to find", item, "at index", i, )
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
	for _ = range pile.Range() {
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
	for _ = range pile.Random() {
		i++
	}

	fmt.Println(pile.Length(), "==", i)
	// Output:
	// 5
	// 5 == 5
}
