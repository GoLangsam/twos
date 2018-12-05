// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

// ===========================================================================

// C allows to range over the first N cardinal numbers:
//    for c := range C(10) {
//        fmt.Println(c)
//    }
// ... will print 0 to 9, inclusive.
func C(N int) Cardinalities {
	Cs := make(chan Cardinality)
	go func(Cs chan<- Cardinality, N int) {
		defer close(Cs)
		for c := 0; c < N; c++ {
			Cs <- Cardinality(c)
		}
	}(Cs, N)
	return Cs
}

// ===========================================================================
