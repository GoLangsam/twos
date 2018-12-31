// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	. "github.com/GoLangsam/twos/core"
)

// ===========================================================================

type tailPair struct {
	Cols Tail
	Rows Tail
}

func (a tailPair) Both() (aten, apep interface{}) { return a.Cols, a.Rows }
func (a tailPair) Size() Cardinality {
	size := Cardinal(0)
	return size.Mul(a.Cols.Size(), a.Rows.Size())
}
func (a tailPair) Tail() Tail { return X(a.Cols, a.Rows) }

// ---------------------------------------------------------------------------

type colsRows struct{
	ID
	tailPair
}

func ColsRows(name ID, cols, rows Iterable) *colsRows {
	return &colsRows{name, tailPair{cols.Tail(), rows.Tail()}}
}

func (a colsRows) Both() (aten, apep interface{}) { return a.ID, a.tailPair }
func (a colsRows) Size() Cardinality { return a.tailPair.Size() }

// ===========================================================================

type IndexedPile interface{
//	Both() (aten, apep interface{})	// Pair
	Size() Cardinality              // Pile
	Tail() Tail                     // Iterable
	Of(Index) Head                  // Indexed
}

// ---------------------------------------------------------------------------

type pilePair struct {
	Cols IndexedPile
	Rows IndexedPile
}

func (a pilePair) Both() (aten, apep interface{}) { return a.Cols, a.Rows }
func (a pilePair) Size() Cardinality {
	size := Cardinal(0)
	return size.Mul(a.Cols.Size(), a.Rows.Size())
}

func (a pilePair) Of(x, y Index) Head {
	return func() Pair {
		return Join(a.Cols.Of(x)(), a.Rows.Of(y)())
	}
}
func (a pilePair) Tail() Tail { return X(a.Cols, a.Rows) }

// ---------------------------------------------------------------------------

type rectangle struct{
	ID
	pilePair
}

func Rectangle(name ID, cols, rows IndexedPile) *rectangle {
	return &rectangle{name, pilePair{cols, rows}}
}

func (a rectangle) Both() (aten, apep interface{}) { return a.ID, a.pilePair }
func (a rectangle) Size() Cardinality { return a.pilePair.Size() }
func (a rectangle) Tail() Tail        { return a.pilePair.Tail() }

// ===========================================================================
