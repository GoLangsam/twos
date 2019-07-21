package core

import (
	"github.com/mauricelam/genny/generic"
)

// ===========================================================================
// Beg of generation for anyType

// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func boolanyType() { // to fool genny
	type anyType generic.Type
}

// ===========================================================================

// anyTypeIs represents a boolean predicate of anyType
// implemented as a boolean function for anything of type anyType.
type anyTypeIs func(anyType) bool

func isanyTypeTrue()  anyTypeIs { return func(a anyType) bool { return true  } }
func isanyTypeFalse() anyTypeIs { return func(a anyType) bool { return false } }

// And evaluates to true, iff all predicates do so.
// Thus: And()() == true, as And() == isanyTypeTrue().
func (a anyTypeIs) And(predicates ...func(anyType) bool) anyTypeIs {
	if a != nil {
		predicates = append([]func(anyType) bool{a}, predicates...)
	}
	if len(predicates) < 1 {
		return isanyTypeTrue()
	}
	return func(a anyType) bool {
		for _, predicate := range predicates {
			if !predicate(a) {
				return false
			}
		}
		return true
	}
}

// Or evaluates to true, iff at least one of the predicates does so.
// Thus: Or()() == false, as Or() == isanyTypeFalse().
func (a anyTypeIs) Or(predicates ...func(anyType) bool) anyTypeIs {
	if a != nil {
		predicates = append([]func(anyType) bool{a}, predicates...)
	}
	if len(predicates) < 1 {
		return isanyTypeFalse()
	}
	return func(a anyType) bool {
		for _, predicate := range predicates {
			if predicate(a) {
				return true
			}
		}
		return false
	}
}

// Not evaluates to true, iff a evaluates to false or is nil.
// Thus: Not(nil)() == true, as Not(nil) == isanyTypeTrue().
func (a anyTypeIs) Not() anyTypeIs {
	if a == nil {
		return isanyTypeTrue()
	}
	return func(arg anyType) bool {
		return !a(arg)
	}
}

// Is returns the predicate a unless it is nil
// in which case the isanyTypeFalse() predicate
// is returned which evaluates to false always.
func (a anyTypeIs) Is() anyTypeIs {
	if a == nil {
		return isanyTypeFalse()
	}
	return a
}

// Eval returns false if a is nil and returns
// the result of what a evaluates to
// for the given argument arg.
func (a anyTypeIs) Eval(arg anyType) bool {
	if a == nil {
		return false
	}
	return a(arg)
}

// End of generation for anyType
// ===========================================================================
