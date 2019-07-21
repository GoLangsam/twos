// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// Fmap establishes Head as Endo-Functor F<Head>.
//   If a == nil this is returned directly - f is not evaluated for nil.
//   If f == nil the identity function is applied.
func (a Head) Fmap(f func(Head) Head) Head {
	if f == nil { f = func(a Head) Head { return a } }
	return func() Pair {
		if a == nil { return nil }
		return f(a)
	}
}

// FmapPair establishes Head as a Functor.
//   If a == nil this is returned directly - f is not evaluated for nil.
//   If f == nil the identity function is applied.
func (a Head) FmapPair(f func(Pair) Pair) Head {
	if f == nil { f = func(a Pair) Pair { return a } }
	return func() Pair {
		if a == nil { return nil }
		return f(a())
	}
}

// ===========================================================================

// Fmap establishes Tail as Endo-Functor F<Tail>.
//   If f == nil the identity functions is applied.
func (a Tail) Fmap(f func(Tail) Tail) Tail {
	if a == nil { return NilTail() }
	if f == nil { f = func(a Tail) Tail { return a } }
	return func() (head Head, tail Tail) {
		head, tail = f(a)()
		if head == nil { return NilTail()() }
		tail = tail.Fmap(f)
		return
	}
}

// FmapHead returns a Tail with f applied to each of its heads.
//   If f == nil the identity functions is applied.
func (a Tail) FmapHead(f func(Head) Head) Tail {
	if a == nil { return NilTail() }
	if f == nil { f = func(a Head) Head { return a } }
	return func() (head Head, tail Tail) {
		aHead, tail := a()
		if aHead == nil { return NilTail()() }
		head = func() Pair { return f(aHead)() }
		tail = tail.FmapHead(f)
		return
	}
}

// FmapPair returns a Tail with f applied to each pair its heads evaluate to.
//   If f == nil the identity functions is applied.
//   If some head evaluates to a Pair == nil this is returned directly - f is not evaluated for nil.
func (a Tail) FmapPair(f func(Pair) Pair) Tail {
	if a == nil { return NilTail() }
	if f == nil { f = func(a Pair) Pair { return a } }
	return func() (head Head, tail Tail) {
		aHead, tail := a()
		if aHead == nil { return NilTail()() }
		head = func() (pair Pair) {
			pair = aHead()
			if pair != nil { pair = f(pair) }
			return
		}
		tail = tail.FmapPair(f)
		return
	}
}

// ===========================================================================
