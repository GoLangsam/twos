// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

/* Quote:
   Type is the representation of a Go type.

   Not all methods apply to all kinds of types. Restrictions, if any, are noted
   in the documentation for each method. Use the Kind method to find out the
   kind of type before calling kind-specific methods. Calling a method
   inappropriate to the kind of type causes a run-time panic.

   Type values are comparable, such as with the == operator, so they can be
   used as map keys. Two Type values are equal if they represent identical
   types.
*/

// reflect.Type. FieldByNameFunc(match func(string) bool) (StructField, bool)
//
// FieldByNameFunc returns the struct field with a name
// that satisfies the match function and a boolean indicating if
// the field was found.
//
// FieldByNameFunc considers the fields in the struct itself
// and then the fields in any embedded structs, in breadth first order,
// stopping at the shallowest nesting depth containing one or more
// fields satisfying the match function. If multiple fields at that depth
// satisfy the match function, they cancel each other
// and FieldByNameFunc returns no match.
// This behavior mirrors Go's handling of name lookup in
// structs containing embedded fields.
