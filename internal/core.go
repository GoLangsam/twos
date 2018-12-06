// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

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

// Cardinality represents a cardinal number such as the #-of items in a Pile.
type Cardinality = core.Cardinality

// ===========================================================================

// Head is a thunk
// which evaluates to
// a Pair.
type Head = core.Head

// Tail is a thunk
// which evaluates to
// a Head and to another Tail.
type Tail = core.Tail

// ===========================================================================
// Interfaces

// Pair has two sides: Aten & Apep. It may be atomic, or composed, nested.
type Pair = core.Pair

// ===========================================================================

var (
	NilTail = core.NilTail
	StringOfTwos = core.StringOfTwos
	StringOfOnes = core.StringOfOnes
	StringOfPair = core.StringOfPair
)