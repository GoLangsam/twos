// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

// ===========================================================================
// a little gem found @ "github.com/bradfitz/iter"

// N returns a slice of n 0-sized elements, suitable for ranging over:
//    for i := range N(10) {
//        fmt.Println(i)
//    }
// ... will print 0 to 9, inclusive.
//
// It does not cause any allocations.
func N(n int) []struct{} {
	return make([]struct{}, n)
}

// ===========================================================================
