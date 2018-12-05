// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"github.com/GoLangsam/twos/core"
)

// ===========================================================================

// ID as a unique identifier - unique among its Kind.
// Implements Named.
type ID = core.ID

// ===========================================================================

// Index represents an ordinal number and may be used
// to enumerate a collention or access some item.
// Intentionally, the first item has Index = 1.
// This is more intuitive for users.
// (Well, programmers prefer offsets over ordinals).
type Index = core.Index

// Indices represents a stream of ordinal numbers
// implemented as a (read-only) channel of Index
// which can be ranged over.
type Indices <-chan Index

// ===========================================================================

// Cardinality represents a cardinal number such as the #-of items in a Pile.
type Cardinality = core.Cardinality

// Cardinalities represents a stream of cardinal numbers
// implemented as a (read-only) channel of Cardinality
// which can be ranged over.
type Cardinalities <-chan Cardinality

// ===========================================================================
