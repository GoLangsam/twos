// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file uses geanny to pull the type specific generic code

// core datatypes: nest,kind,Cardinality,Index,Name
// core functypes: Head,Tail
// core interface: Iterable,Kind,Type,Pair
// void interface: interface{}

//go:generate genny	-in ..\internal\bool.go		-out gen.bool.go	-pkg bool gen "anyType=interface{},Cardinality,Index,ID,Type,Iterable,Pair,Pile"

package bool
