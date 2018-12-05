// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// nest represents a nested pair, a pair of pairs.
type nest struct {
	Aten Pair
	Apep Pair
}

// Aten returns the first constituent of Pair a.
//
// In case this constituent does not implement Pair,
// the kind of Pair a is returned instead.
//
// Intentionally it is not named Fst (for First)
// as Fst may easily be confused with Fist.
func Aten(a Pair) Pair {
	aten, _ := a.Both()
	if pair, ok := aten.(Pair); ok {
		return pair
	}
	return KindOfPair(a)
}

// Apep returns the second constituent of Pair a.
//
// In case this constituent does not implement Pair,
// the Pair a itself is returned instead.
//
// Intentionally it is not named Snd (for Second)
// as Snd may easily be confused with the opposite
// of Rcv - Send and Receive.
func Apep(a Pair) Pair {
	apep, _ := a.Both()
	if pair, ok := apep.(Pair); ok {
		return pair
	}
	return a
}

// Swap reverses a Pair.
func Swap(a Pair) Pair {
	switch a := a.(type) {
	case nest:
		return Join(a.Apep, a.Aten)
	default:
		return Join(Apep(a), Aten(a))
	}
}

// Join returns a pair of pairs - a nested pair.
func Join(a, b Pair) Pair {
	return nest{a, b}
}

// Contains discriminates iff item == a
// or is identical to one of the two pairs of a.
func (a nest) Contains(item interface{}) (contains bool) {
	if n, contains := item.(nest); contains {
		if contains = a == n; !contains {
			if p, contains := item.(Pair); contains {
				if contains = a.Aten == p; !contains {
					return a.Apep == p
				}
			}
		}
	}
	return
}
