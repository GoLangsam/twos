// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"strings"
)

const nodeFmt = "{%+v<%+v>%+v}"
const tailBeg = "["
const tailEnd = "]"

// ===========================================================================

type stringer interface {
	String() string
}

// String implements fmt.Stringer
func (a nilPair) String() string { return StringOfOnes("<nilPair>") }

// String implements fmt.Stringer
func (a name) String() string {
	if a == "" {
		a = nilName
	}
	return StringOfOnes(string(a))
}

/* String implements fmt.Stringer */ func (a Index) String() string { return StringOfOnes(int(a)) }

/* String implements fmt.Stringer */ func (a Cardinality) String() string { return StringOfOnes(int(a)) }

/* String implements fmt.Stringer */ func (a nest) String() string { return StringOfTwos(StringOfPair(a.Aten), StringOfPair(a.Apep)) }

/* String implements fmt.Stringer */ func (a kind) String() string { return StringOfTwos(a.ID, a.Type) }

/* String implements fmt.Stringer */ func (a Head) String() string {
	if a == nil {
		return StringOfOnes("(<nilHead>)")
	}
	return StringOfPair(a())
}

/* String implements fmt.Stringer */ func (a Tail) String() string {
	var b strings.Builder
	fmt.Fprint(&b, tailBeg)
	if a == nil {
		fmt.Fprintf(&b, StringOfOnes("<nilTail>"))
	} else {
		for head, tail := a(); head != nil; head, tail = tail() {
			fmt.Fprintf(&b, StringOfPair(head()))
		}
	}
	fmt.Fprint(&b, tailEnd)
	return b.String()
}

// ===========================================================================

// Printer returns a new Tail which, upon traversal,
// will print the Pair of each Head using stringer.
// Example:
//  tail := tail.Printer(DeKind)
func (a Tail) Printer(stringer func(a Pair) string) Tail {
	return a.FmapPair(func(a Pair) Pair {
		fmt.Print(stringer(a))
		return a
	})
}

// ===========================================================================

// StringOfPair returns the string of a Pair.
func StringOfPair(a Pair) string {
	if a == nil {
		a = nilPair{}
	} // Note: now StringOfPair(a) differs from fmt.Print(a)

	switch t := a.(type) {
	case fmt.Stringer: // nilPair, name, ID, Index, Cardinality, nest, kind, Head, Tail ...
		return t.String()
	case Pair:
		aten, apep := t.Both()
		return StringOfTwos(aten, apep)
	default:
		return StringOfOnes(t)
	}
}

// ===========================================================================

const itemFmt = "%+v"
const twosFmt = "{ %+v | %+v }"

// StringOfTwos returns the string of two items.
func StringOfTwos(a, b interface{}) string { return fmt.Sprintf(twosFmt, a, b) }

// StringOfOnes returns the string of one item.
func StringOfOnes(a interface{}) string { return fmt.Sprintf(itemFmt, a) }
