// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package core provides generic implementations such as
//   - bool.go for boolean operations (And, Or, Not) on attribute functions
//   - fmap.go for Fmap... and Join... for slices and read-only channels
//   - pile.go for PileOf...
//
// as well as
//   - pile_test.go for basic tests of PileOf... functionalities
//
// This package is not imported anwhere.
// See genny.go elsewhere for it's usage.
//
// Note: anyType_test.go provides some test data - any client package
// must provide similar data, appropriate for the target type,
// if pile_test.go is generated into such package.
package core
