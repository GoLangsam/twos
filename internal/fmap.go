package core

import (
	"github.com/mauricelam/genny/generic"
)

// ===========================================================================
// Beg of generation for anyType

// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func fmapanyType() { // to fool genny
	type anyType generic.Type
}

// ===========================================================================

// Naming conventions:
//  anyTypes   - a variadic sequence of anyType
//  anyTypeS   - a slice of anyType's
//  anyTypeSS  - a slice of slices of anyType
//  anyTypeRoC - a read-only channel of anyType

// FmapanyTypes establishes []anyType as a Functor.
func FmapanyTypes(f func(anyType) anyType, anyTypes ...anyType) (anyTypeS []anyType) {
	anyTypeS = make([]anyType, len(anyTypes))
	for _, i := range anyTypes {
		anyTypeS = append(anyTypeS, f(i))
	}
	return
}

// JoinanyTypeS helps []anyType to become a Monad.
// It's just a list comprehension - aka Concat.
func JoinanyTypeS(anyTypeSS [][]anyType) (anyTypeS []anyType) {
	for i := range anyTypeSS {
		anyTypeS = append(anyTypeS, anyTypeSS[i]...)
	}
	return
}

// FmapanyTypeRoC establishes <-chan anyType as a Functor.
func FmapanyTypeRoC(f func(anyType) anyType, RoCs ...<-chan anyType) <-chan anyType {
	out := make(chan anyType)
	go func(WoC chan<- anyType, f func(anyType) anyType, RoCs ...<-chan anyType) {
		defer close(WoC)
		for _, RoC := range RoCs {
			for i := range RoC {
				WoC <- f(i)
			}
		}
	}(out, f, RoCs...)
	return out
}

// End of generation for anyType
// ===========================================================================