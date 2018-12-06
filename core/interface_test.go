// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

func ExamplePair_interface() {

	var _ Pair = ID("Test")
	var _ Pair = Index(4711)
	var _ Pair = Cardinality(4711)

	var _, _ Pair = nilPair{}, &nilPair{}
	var _, _ Pair = nest{}, &nest{}
	var _, _ Pair = kind{}, &kind{}

	var _ Pair = NilTail()
	var _, _ Pair = NilTail()()

}

func ExampleIterable_interface() {

	var _ Iterable = ID("Test")
	var _ Iterable = Index(4711)
	var _ Iterable = Cardinality(4711)

	var _, _ Iterable = nilPair{}, &nilPair{}
	var _, _ Iterable = nest{}, &nest{}
	var _, _ Iterable = kind{}, &kind{}

	var _ Iterable = NilTail()
	var _, _ Iterable = NilTail()()

}

// Named represents a named type.
type Named interface{ Name() string }

func ExampleNamed_interface() {

	var _ Named = ID("Test")
	var _ Named = Index(4711)
	var _ Named = Cardinality(4711)

	var _, _ Named = nilPair{}, &nilPair{}
	var _, _ Named = nest{}, &nest{}
	var _, _ Named = kind{}, &kind{}

	var _ Named = NilTail()
	var _, _ Named = NilTail()()

}

// Pile holds Length items.
type Pile interface {
	Length() Cardinality
}

func ExamplePile_interface() {

	var _ Pile = ID("Test")
	var _ Pile = Index(4711)
	var _ Pile = Cardinality(4711)

	var _, _ Pile = nilPair{}, &nilPair{}
	var _, _ Pile = nest{}, &nest{}
	var _, _ Pile = kind{}, &kind{}

	var _ Pile = NilTail()
	var _, _ Pile = NilTail()()

}

// Indexed allows to retrieve a Head Of any item by Index.
type Indexed interface {
	Of(Index) Head
}

func ExampleIndexed_interface() {

	var _ Indexed = ID("Test")
	var _ Indexed = Index(4711)
	var _ Indexed = Cardinality(4711)

	var _, _ Indexed = nilPair{}, &nilPair{}
	var _, _ Indexed = nest{}, &nest{}
	var _, _ Indexed = kind{}, &kind{}

	var _ Indexed = NilTail()
	var _, _ Indexed = NilTail()()
}

// Container can tell for any item whether it contains this item or not.
type Container interface {
	Contains(item interface{}) bool
}

func ExampleContainer_interface() {

	var _ Container = ID("Test")
	var _ Container = Index(4711)
	var _ Container = Cardinality(4711)

	var _ Container = NilTail()
	var _, _ Container = NilTail()()

	var _, _ Container = nilPair{}, &nilPair{}
	var _, _ Container = nest{}, &nest{}
	var _, _ Container = kind{}, &kind{}
}
