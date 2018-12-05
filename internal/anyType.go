// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/mauricelam/genny/generic"
)

type anyType generic.Type

type someType struct{ a interface{} }

func (a someType) Both() (aten, apep interface{}) { return a, a }
