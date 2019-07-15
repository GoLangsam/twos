// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/id/c.go"

package data

// ===========================================================================

// C allows to range over the first cardinal numbers upto (and including) N:
//    for c := range C(10) {
//        fmt.Println(c)
//    }
// ... will print 0 to 10, inclusive.
func C(N int) Cardinalities {
	Cs := make(chan Cardinality)
	go func(Cs chan<- Cardinality, N int) {
		defer close(Cs)
		for c := 0; c <= N; c++ {
			Cs <- Cardinal(c)
		}
	}(Cs, N)
	return Cs
}

// ===========================================================================
