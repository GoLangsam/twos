// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/id/i.go"

package data

// ===========================================================================

// I allows to range over the first N ordinal numbers:
//    for i := range I(10) {
//        fmt.Println(i)
//    }
// ... will print 1 to 10, inclusive.
func I(N int) Indices {
	Is := make(chan Index)
	go func(Is chan<- Index, N int) {
		defer close(Is)
		for i := 1; i < N+1; i++ {
			Is <- Ordinal(i)
		}
	}(Is, N)
	return Is
}

// ===========================================================================
