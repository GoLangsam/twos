// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data

import (
	"fmt"
	"testing"
)

func ExampleNames() {
	for i := range Names("ID-", 4) {
		fmt.Println(i)
	}
	// Output:
	// ID-1
	// ID-2
	// ID-3
	// ID-4
}

func ExampleC() {
	for i := range C(4) {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
}

func ExampleI() {
	for i := range I(4) {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

// stolen from "github.com/bradfitz/iter"

// ===========================================================================
func ExampleN() {
	for i := range N(4) {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
}

func TestAllocs(t *testing.T) {
	var x []struct{}
	allocs := testing.AllocsPerRun(500, func() {
		x = N(1e9)
		_ = x
	})
	if allocs > 0.1 {
		t.Errorf("allocs = %v", allocs)
	}
}
