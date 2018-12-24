// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	. "github.com/GoLangsam/twos/core"
)

// ===========================================================================

type permuter struct {
	reset TailS // for reset
	TailS       // where we are
	HeadS       // what we have
	curr  int   // where to progress
}

// Permuter provides an (opaque and lazy) iterator.
//
// Its Next() method succesively produces
// a HeadS thru all permutations of the items
// reachable thru the given TailS.
//
// Its Tail() method implements Iterator and allows composition,
// but may need a type assertion upon evaluation of some head.
//
// Along the way, Next() may be mixed with evaluation of some tail().
func Permuter(tails TailS) *permuter {
	a := &permuter{
		tails,
		TailS{},
		HeadS{},
		0,
	}

	for curr := 0; curr < len(a.reset); curr++ {
		head, tail := a.reset[curr].Tail()()
		a.HeadS = append(a.HeadS, head)
		a.TailS = append(a.TailS, tail)
	}

	return a.upto(len(a.reset))
}

// Tail implements Iterable
// by returning all permutations sucessively.
func (a *permuter) Tail() Tail {
	return func() (head Head, tail Tail) {
		if len(a.HeadS) < 1 {
			return NilTail()()
		}

		return func() Pair { // head
			if headS, ok := a.Next(); ok {
				return headS
			}
			return nil
		}, a.Tail() // tail - tail recursiv
	}
}

// Next allows to iterate thru all permutations
// until nil, false.
func (a *permuter) Next() (HeadS, bool) {
	if len(a.HeadS) < 1 {
		return nil, false
	}
	headS := HeadS{}
	for _, head := range a.HeadS {
		headS = append(headS, head)
	} // copy(headS, a.HeadS)
	// Note: Intentionally use of copy is avoided: it would need to specify the underlying type.

	a = a.next(a.curr)
	return headS, true
}

// Upto resets tails up to (but not including) given offset.
func (a *permuter) upto(upto int) *permuter {
	if size := len(a.HeadS); size < upto {
		upto = size // upto = min(size, upto)
	}
	a.curr = 0                           // reset position
	for curr := 0; curr < upto; curr++ { // reset all lower
		a.HeadS[curr], a.TailS[curr] = a.reset[curr].Tail()()
	}
	return a
}

// Next evaluates the Tail here and carries over if need.
// On overflow, the output buffer is shrunk to zero size.
func (a *permuter) next(here int) *permuter {
	if here < len(a.HeadS) {
		a.HeadS[here], a.TailS[here] = a.TailS[here].Tail()()
		if a.HeadS[here] == nil {
			here++
			return a.upto(here).next(here) // advance first - tail recurse
		}
	} else {
		a.HeadS = HeadS{}
	} // clear output buffer
	return a
}

// ===========================================================================
// common

// Len reports the product of the length of the Tails.
func (a *permuter) Len() int {
	var size = 1
	for _, t := range a.reset {
		size = size * int(t.Size())
	}
	return size
}

// Size implements Pile
// by returning
// the product of the length of the Tails.
func (a *permuter) Size() Cardinality {
	var size Cardinality = 1
	for _, t := range a.reset {
		size = size * t.Size()
	}
	return size
}
