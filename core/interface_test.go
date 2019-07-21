// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

func ExamplePair_interface() {

	var _ Pair = ID("Test")
	var _ Pair = Ordinal(4711)
	var _ Pair = Cardinal(4711)

	var _, _ Pair = nilPair{}, &nilPair{}
	var _, _ Pair = nest{}, &nest{}
	var _, _ Pair = kind{}, &kind{}

	var _ Pair = NilTail()
	var _, _ Pair = NilTail()()

}

func ExampleIterable_interface() {

	var _ Iterable = ID("Test")
	var _ Iterable = Ordinal(4711)
	var _ Iterable = Cardinal(4711)

	var _, _ Iterable = nilPair{}, &nilPair{}
	var _, _ Iterable = nest{}, &nest{}
	var _, _ Iterable = kind{}, &kind{}

	var _ Iterable = NilTail()
	var _, _ Iterable = NilTail()()

}

func ExampleNamed_interface() {

	// Named represents a named type.
	type Named interface{ Name() string }

	var _ Named = ID("Test")
	var _ Named = Ordinal(4711)
	var _ Named = Cardinal(4711)

	var _, _ Named = nilPair{}, &nilPair{}
	var _, _ Named = nest{}, &nest{}
	var _, _ Named = kind{}, &kind{}

	var _ Named = NilTail()
	var _, _ Named = NilTail()()

}

func ExampleStringer_interface() {

	// Stringer represents a type which can express itself as a string.
	type Stringer interface{ String() string }

	var _ Stringer = ID("Test")
	var _ Stringer = Ordinal(4711)
	var _ Stringer = Cardinal(4711)

	var _, _ Stringer = nilPair{}, &nilPair{}
	var _, _ Stringer = nest{}, &nest{}
	var _, _ Stringer = kind{}, &kind{}

	var _ Stringer = NilTail()
	var _, _ Stringer = NilTail()()

}

func ExamplePile_interface() {

	// Pile holds Size items.
	type Pile interface { Size() Cardinality }

	var _ Pile = ID("Test")
	var _ Pile = Ordinal(4711)
	var _ Pile = Cardinal(4711)

	var _, _ Pile = nilPair{}, &nilPair{}
	var _, _ Pile = nest{}, &nest{}
	var _, _ Pile = kind{}, &kind{}

	var _ Pile = NilTail()
	var _, _ Pile = NilTail()()

}

func ExampleIndexed_interface() {

	// Indexed allows to retrieve a Head Of any item by Index.
	type Indexed interface { Of(Index) Head }

	var _ Indexed = ID("Test")
	var _ Indexed = Ordinal(4711)
	var _ Indexed = Cardinal(4711)

	var _, _ Indexed = nilPair{}, &nilPair{}
	var _, _ Indexed = nest{}, &nest{}
	var _, _ Indexed = kind{}, &kind{}

	var _ Indexed = NilTail()
	var _, _ Indexed = NilTail()()
}

func ExampleContainer_interface() {

	// Container can tell for any item whether it contains this item or not.
	type Container interface { Contains(item interface{}) bool }

	var _ Container = ID("Test")
	var _ Container = Ordinal(4711)
	var _ Container = Cardinal(4711)

	var _, _ Container = nilPair{}, &nilPair{}
	var _, _ Container = nest{}, &nest{}
	var _, _ Container = kind{}, &kind{}

	var _ Container = NilTail()
	var _, _ Container = NilTail()()
}
