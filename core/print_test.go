// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
)

// ===========================================================================

func ExampleID_Unit() {
	var i ID
	i = i.Unit()
	fmt.Println(i)
	// Output:
	// <noName>
}

func ExampleIndex_Unit() {
	var i Index
	i = i.Unit()
	fmt.Println(i)
	// Output:
	// 1
}

func ExampleCardinality_Unit() {
	var c Cardinality
	c = c.Unit()
	fmt.Println(c)
	// Output:
	// 0
}

// ===========================================================================

func ExampleKindOfPair() {

	fmt.Println(KindOfPair(ID("TestName")))
	fmt.Println(KindOfPair(Index{}))
	fmt.Println(KindOfPair(Cardinality{}))
	// Output:
	// { name | core.name }
	// { ordinalNumber | core.ordinalNumber }
	// { cardinalNumber | core.cardinalNumber }
}

// ===========================================================================

func ExampleStringOfPair() {
	fmt.Println(StringOfPair(ID("TestName")))
	// Output:
	// TestName
}

// ===========================================================================

func ExampleStringOfPair_nil() {

	var namE name
	var indx Index
	var card Cardinality
	var kinD kind
	var head Head
	var tail Tail
	var nesT nest
	var pair Pair

	fmt.Println("name:", namE, "==", StringOfPair(namE))
	fmt.Println("indx:", indx, "==", StringOfPair(indx))
	fmt.Println("card:", card, "==", StringOfPair(card))
	fmt.Println("kind:", kinD, "==", StringOfPair(kinD))
	fmt.Println("head:", head, "==", StringOfPair(head))
	fmt.Println("tail:", tail, "==", StringOfPair(tail))
	fmt.Println("nest:", nesT, "==", StringOfPair(nesT))
	fmt.Println()
	fmt.Println("pair:", pair, "!=", StringOfPair(pair))

	// Output:
	// name: <noName> == <noName>
	// indx: <nil> == <nil>
	// card: <nil> == <nil>
	// kind: { <noName> | <nil> } == { <noName> | <nil> }
	// head: (<nilHead>) == (<nilHead>)
	// tail: [<nilTail>] == [<nilTail>]
	// nest: { <nilPair> | <nilPair> } == { <nilPair> | <nilPair> }
	//
	// pair: <nil> != <nilPair>
}

func ExampleKindOfPair_nil() {

	var name ID
	var index Index
	var cardinality Cardinality

	var kinD kind
	var head Head
	var tail Tail
	//	var pair Pair

	{
		k := KindOfPair(name)
		fmt.Println("name.Kind():     ", k.ID, k.Type)
	}
	{
		k := KindOfPair(index)
		fmt.Println("index.Kind():    ", k.ID, k.Type)
	}
	{
		k := KindOfPair(cardinality)
		fmt.Println("card.Kind():     ", k.ID, k.Type)
	}
	fmt.Println()
	{
		k := KindOfPair(kinD)
		fmt.Println("kinD.Kind():     ", k.ID, k.Type)
	}
	{
		k := KindOfPair(head)
		fmt.Println("head.Kind():     ", k.ID, k.Type)
	}
	{
		k := KindOfPair(tail)
		fmt.Println("tail.Kind():     ", k.ID, k.Type)
	}

	// Output:
	// name.Kind():      name core.name
	// index.Kind():     ordinalNumber core.ordinalNumber
	// card.Kind():      cardinalNumber core.cardinalNumber
	//
	// kinD.Kind():      <noName> core.kind
	// head.Kind():      <noName> core.Head
	// tail.Kind():      <noName> core.Tail
}

func ExampleNilTail() {

	tail := NilTail() // Note: Any tail implements Pair

	fmt.Println(tail)
	fmt.Println(tail.String())
	fmt.Println(StringOfPair(tail))

	head, tail := tail()

	fmt.Println(head)
	fmt.Println(head.String())
	fmt.Println(StringOfPair(head))

	fmt.Println(tail)
	fmt.Println(tail.String())
	fmt.Println(StringOfPair(tail))

	// Output:
	// []
	// []
	// []
	// (<nilHead>)
	// (<nilHead>)
	// (<nilHead>)
	// []
	// []
	// []
}
