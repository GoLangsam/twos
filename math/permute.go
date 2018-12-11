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
	many  int   // how many tails
	curr  int   // where to progress
}

// Permuter provides an (opaque and lazy) iterator.
// Its Next() method succesively produces
// a HeadS thru all permutations of the items
// reachable thru the given TailS.
func Permuter(tails TailS) *permuter {
	many := len(tails)
	if many < 1 {
		return nil
	}

	a := &permuter{
		tails,
		TailS{},
		HeadS{},
		many,
		0,
	}

	for curr := 0; curr < len(a.reset); curr++ {
		head, tail := a.reset[curr].Tail()()
		a.HeadS = append(a.HeadS, head)
		a.TailS = append(a.TailS, tail)
	}

	return a.upto(a.many)
}

// Next iterates: returns another HeadS and reports iff there is such.
func (a *permuter) Next() (HeadS, bool) {
	if len(a.HeadS) < 1 {
		return nil, false
	}

	headS := HeadS{}
	for _, head := range a.HeadS {
		headS = append(headS, head)
	} // copy(headS, a.HeadS)
	// Note: Intentionally use of copy is avoided: it would need to specify the underlying type.

	a.next(a.curr)

	return headS, true
}

// Upto resets tails up to (but not including) given offset.
func (a *permuter) upto(upto int) *permuter {
	for curr := 0; curr < upto && curr < len(a.HeadS); curr++ {
		a.HeadS[curr], a.TailS[curr] = a.reset[curr].Tail()() // reset all lower
	}
	a.curr = 0
	return a
}

// Next evaluates the Tail here and carries over if need.
// On overflow, the output buffer is shrunk to zero size.
func (a *permuter) next(here int) *permuter {
	if here < len(a.HeadS) {

		a.HeadS[here], a.TailS[here] = a.TailS[here].Tail()()
		if a.HeadS[here] == nil {
			here++
			if here >= a.many { // too many: Overflow!
				a.HeadS = HeadS{} // clear output buffer
				return a
			}
			a.curr = 0                           // reset position
			for curr := 0; curr < here; curr++ { // reset all lower
				a.HeadS[curr], a.TailS[curr] = a.reset[curr].Tail()()
			}
			return a.next(here) // advance first - recurse
		}
	}
	return a
}
