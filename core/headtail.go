// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// Head is a thunk
// which evaluates to
// a Pair.
type Head func() Pair

// Tail is a thunk
// which evaluates to
// a Head and to another Tail.
type Tail func() (Head, Tail)

// NilTail is a helper, it returns the Tail which evaluates to nil, nil.
//  Note: this Tail is unique. So its the tail, not a tail.
func NilTail() Tail { return nilTail }

// ===========================================================================

// Iter permits to iterate through given pairs - Usage pattern:
//  for head, tail := Iter(a, b, c); head != nil; head, tail = tail() {
//      // Beware: head is a function, a thunk, an accessor!
//      _ = head() // do sth with Head's content ...
//      // or pass head to another function
//  }
func Iter(pairs ...Pair) (tail Tail) {
	if len(pairs) < 1 { return NilTail() }
	return tailRecurse(pairs...)
}

func tailRecurse(pairs ...Pair) (tail Tail) {
	return func() (head Head, tail Tail) {
		if len(pairs) < 1 { return NilTail()() }
		head = func() Pair { return pairs[0] }
		tail = tailRecurse(pairs[1:]...)
		return head, tail
	}
}

// ===========================================================================

// Join helps Tail to become a Monad.
func (a Tail) Join(tails ...Iterable) Tail {
	return func() (head Head, tail Tail) {
		head, tail = a()
		if head == nil {
			if len(tails) > 0 {
				head, tail = tails[0].Tail()()
				tail = tail.Join(tails[1:]...)
				return
			}
			return NilTail()()
		}
		tail = tail.Join(tails...)
		return
	}
}

// ===========================================================================

// Range returns a channel to range over.
func (a Tail) Range() <-chan Pair {
	pairs := make(chan Pair)

	go func(pairs chan<- Pair, tail Tail) {
		defer close(pairs)
		for head, tail := tail(); head != nil; head, tail = tail() {
			pairs <- head()
		}
	}(pairs, a)
	return pairs
}

// ===========================================================================

// Sew walks a Tail and applies f to every head().
// Think of a thread and a needle...
//
// The original Tail is returned in order to ease chaining with other methods.
func (a Tail) Sew(needle func(Pair)) Tail {
	for head, tail := a(); head != nil; head, tail = tail() {
		needle(head())
	}
	return a
}

// ===========================================================================

// Skip returns
// a Tail
// which allows to traverse only those heads of the Tail of the original iterable
// which evaluate to a pair
// that does not satisfy the given boolean predicate function.
func Skip(iter Iterable, pairIs func(Pair) bool) Tail {
	return func() (head Head, tail Tail) {
		for head, tail = iter.Tail()(); head != nil && pairIs(head()); head, tail = tail() { /* noop */ }
		if head == nil { return NilTail()() }
		return
	}
}

// Only returns
// a Tail
// which allows to traverse only those heads of the Tail of the original iterable
// which evaluate to a pair
// that satisfies the given boolean predicate function.
func Only(iter Iterable, pairIs func(Pair) bool) Tail {
	return Skip(iter, func(a Pair) bool { return !pairIs(a) })
}

// ===========================================================================

// Skip returns
// a Tail
// which allows to traverse those heads only
// which evaluate to a pair
// that does not satisfy the given boolean predicate function.
func (a Tail) Skip(pairIs func(Pair) bool) Tail {
	return Skip(a, pairIs)
}

// Only returns
// a Tail
// which allows to traverse those heads only
// which evaluate to a pair
// which satisfies the given boolean predicate function.
func (a Tail) Only(pairIs func(Pair) bool) Tail {
	return Only(a, pairIs)
}

// ===========================================================================

// X returns the (right-associative) cartesian product
// of a with the given iterables.
// A nil a is ignored when composing the product.
func (a Tail) X(iters ...Iterable) Tail {
	if a != nil {
		iters = append([]Iterable{a}, iters...)
	}
	return X(iters...)
}

// ===========================================================================
