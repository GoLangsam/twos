// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

func(a kind		)Lift(as ...kind) 	[]kind		{return append([]kind{a},		as...)}
func(a name		)Lift(as ...name) 	[]name		{return append([]name{a},		as...)}
func(a Index		)Lift(as ...Index)	[]Index		{return append([]Index{a},		as...)}
func(a Cardinality	)Lift(as ...Cardinality)[]Cardinality	{return append([]Cardinality{a},	as...)}
func(a nest		)Lift(as ...nest) 	[]nest		{return append([]nest{a},		as...)}
func(a Head		)Lift(as ...Head) 	[]Head		{return append([]Head{a},		as...)}
func(a Tail		)Lift(as ...Tail) 	[]Tail		{return append([]Tail{a},		as...)}

// ===========================================================================
