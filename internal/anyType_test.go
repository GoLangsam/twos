// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"time"
)

func testanyType() (string, []anyType) {
	name := "any"
	data := []anyType{
		int(4),
		Cardinality(7),
		time.Monday,
		time.Monday,
		"Thursday",
		ID("Friday"),
	}
	return name, data
}
