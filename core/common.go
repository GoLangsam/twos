// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// ===========================================================================

// kind implements a named type
type kind struct {
	ID
	Type
}

const nilName = name("<noName>")

type nilPair struct{}

// I wish this could be constant
var (
	nilHead Head = func() Pair { return nilPair{} }
	nilTail Tail = func() (Head, Tail) { return nil, func() (Head, Tail) {return nil, nil} }

	theKindOfNest = kind{name(TypeOf(nest{}).Name()), TypeOf(nest{})}
	theKindOfKind = kind{name(TypeOf(kind{}).Name()), TypeOf(kind{})}
	theKindOfName = kind{name(TypeOf(name("")).Name()), TypeOf(name(""))}
	theKindOfIndex = kind{name(TypeOf(Index(1)).Name()), TypeOf(Index(1))}
	theKindOfCardinality = kind{name(TypeOf(Cardinality(0)).Name()), TypeOf(Cardinality(0))}
)

// KindOfPair returns the kind of a Pair,
// a special Pair which contains the ID and the Type.
func KindOfPair(a Pair)	kind {
	if a == nil {return kind{nilName, TypeOf(nilPair{})} }
	type Named interface { Name() string }
	switch a := a.(type) {
	case kind:
		if a.ID == ""    {a.ID   = nilName}
		if a.Type == nil {a.Type = TypeOf(kind{})}
		return a
	case nest:
		return theKindOfNest
	case name:
		return theKindOfName
	case Index:
		return theKindOfIndex
	case Cardinality:
		return theKindOfCardinality
	case Head:
		if a == nil {return kind{nilName, TypeOf(nilHead)}   }
		return kind{name(TypeOf(a).Name()), TypeOf(a)}
	case Tail:
		if a == nil {return kind{nilName, TypeOf(nilTail)}   }
		return kind{name(TypeOf(a).Name()), TypeOf(a)}
	case Named:
		return kind{name(a.Name()), TypeOf(a)}
	default:
		return kind{name(TypeOf(a).Name()), TypeOf(a)}
	}
}

/* Name implements Named by returning the Name of the TypeOf(a).	*/func (a nilPair)     Name() string { return TypeOf(a).Name() }
/* Name implements Named by returning the Name of the TypeOf(a).	*/func (a kind)        Name() string { return string(a.ID) }
/* Name implements Named by returning the Name of the TypeOf(a).	*/func (a name)        Name() string { return string(a) }
/* Name implements Named by returning the Name of the TypeOf(a).	*/func (a Index)       Name() string { return TypeOf(a).Name() }
/* Name implements Named by returning the Name of the TypeOf(a).	*/func (a Cardinality) Name() string { return TypeOf(a).Name() }
/* Name implements Named by returning the Name of the TypeOf(a).	*/func (a nest)        Name() string { return TypeOf(a).Name() }
/* Name implements Named by returning the Name of the TypeOf(a).	*/func (a Head)        Name() string { return TypeOf(a).Name() }
/* Name implements Named by returning the Name of the TypeOf(a).	*/func (a Tail)        Name() string { return TypeOf(a).Name() }

// ===========================================================================

/* Both implements Pair by returning the Kind and the Type.		*/func (a nilPair)     Both() (aten, apep interface{}) { return KindOfPair(a),     nil }
/* Both implements Pair by returning the Name and the Type.		*/func (a kind)        Both() (aten, apep interface{}) { return KindOfPair(a).ID, KindOfPair(a).Type }
/* Both implements Pair by returning the Kind and the value.		*/func (a name)        Both() (aten, apep interface{}) { return KindOfPair(a),       a }
/* Both implements Pair by returning the Kind and the value.		*/func (a Index)       Both() (aten, apep interface{}) { return KindOfPair(a),       a }
/* Both implements Pair by returning the Kind and the value.		*/func (a Cardinality) Both() (aten, apep interface{}) { return KindOfPair(a),       a }
/* Both implements Pair by returning both parts of a.                   */func (a nest)        Both() (aten, apep interface{}) { return a.Aten, a.Apep }
/* Both implements Pair by returning Both() of the evaluated Head.	*/func (a Head)        Both() (aten, apep interface{}) {
	if a == nil {return KindOfPair(a), nilHead}
	if a() == nil {return nilPair{}.Both()};	return a().Both() }
/* Both implements Pair by returning the Head and Tail a evaluates to.	*/func (a Tail)        Both() (aten, apep interface{}) {
	if a == nil {return KindOfPair(a), nil};	return a() }

// ===========================================================================

// Tail implements Iterable
// by returning
// the unique NilTail().
func (a nilPair) Tail() Tail { return NilTail() }

// Tail implements Iterable
// by returning
// a head which evaluates to a ( head() == Pair(a) ) and
// as tail the unique NilTail().
func (a kind) Tail() Tail                          { return func() (Head, Tail) { return func() Pair { return a }, NilTail() } }

// Tail implements Iterable
// by returning
// a head which evaluates to a ( head() == Pair(a) ) and
// as tail the unique NilTail().
func (a name) Tail() Tail                          { return func() (Head, Tail) { return func() Pair { return a }, NilTail() } }

// Tail implements Iterable
// by returning
// a head which evaluates to a ( head() == Pair(a) ) and
// as tail the unique NilTail().
func (a Index) Tail() Tail                         { return func() (Head, Tail) { return func() Pair { return a }, NilTail() } }

// Tail implements Iterable
// by returning
// a head which evaluates to a ( head() == Pair(a) ) and
// as tail the unique NilTail().
func (a Cardinality) Tail() Tail                   { return func() (Head, Tail) { return func() Pair { return a }, NilTail() } }

// Tail implements Iterable
// by returning
// a head for the first Pair and
// a tail for the second Pair.
func (a nest) Tail() Tail                          { return func() (Head, Tail) { return func() Pair { return a.Aten }, func() (Head, Tail) { return func() Pair { return a.Apep }, NilTail() } } }

// Tail implements Iterable
// by returning
// a head which evaluates to a ( head() == Pair(a) ) and
// as tail the unique NilTail().
func (a Head) Tail() Tail                          { return func() (Head, Tail) { return func() Pair { return a }, NilTail() } }

// Tail implements Iterable
// by returning
// itself - idempotent, so to say.
func (a Tail) Tail() Tail                          { return a }

// ===========================================================================

// LengthOfPair reports the length (as a Cardinality).
func LengthOfPair(a Pair ) (length Cardinality) {
	// Pile holds Length items.
	type Pile interface {
		Length() Cardinality
	}
	if p, ok := a.(Pile); ok { return p.Length() }
	return
}

/* Length implements Pile by returning 0 */func (a nilPair)     Length() Cardinality { return 0 }
/* Length implements Pile by returning 1 */func (a kind)        Length() Cardinality { return 1 }
/* Length implements Pile by returning 1 */func (a name)        Length() Cardinality { return 1 }

// Length implements Pile by returning 1.
func (a Index)       Length() Cardinality { return 1 }

// Length implements Pile by returning 1
func (a Cardinality) Length() Cardinality { return 1 }

// Length implements Pile
// by returning
// the sum of the length of the two pairs.
func (a nest) Length() Cardinality { return LengthOfPair(a.Aten) + LengthOfPair(a.Apep) }

// Length implements Pile
// by returning
// the length of the pair a evaluates to
// or zero for nil
func (a Head) Length() Cardinality {
	if a == nil { return 0 }
	if a() == nil { return 0 }
	return LengthOfPair(a())
}

// Length implements Pile
// by incrementing upon traversing a.
//  Note: Complexity is O(n) due to complete traversal.
func (a Tail) Length() Cardinality {
	var i Cardinality
	for head, tail := a(); head != nil; head, tail = tail() {
		i++
	}
	return i
}

// LengthRecursive is a recursive implementation to determine the Length
// by returning zero for nil and 1 plus the recursive length of the tail a evaluates to.
//  Note: Complexity is O(n) due to recursion.
func (a Tail) LengthRecursive() Cardinality {
	head, tail := a()
	if head == nil {
		return 0
	}
	return 1 + tail.LengthRecursive()
}

// ===========================================================================

// Contains TODO

// ===========================================================================

// Of returns a nilHead.
func (a nilPair)     Of(index Index) Head { return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a kind)        Of(index Index) Head { if index == 1 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a name)        Of(index Index) Head { if index == 1 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a Index)       Of(index Index) Head { if index == 1 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a Cardinality) Of(index Index) Head { if index == 1 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a nest)        Of(index Index) Head { if index == 1 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a Head)        Of(index Index) Head { if index == 1 {return func() Pair {return a}}; return nilHead }

// Of returns a Head for the item of position index, of nilHead iff the index is out of bounds.
func (a Tail)        Of(index Index) Head {
	var current Index
	for head, tail := a(); head != nil && current < index; head, tail = tail() {
		current++
		if index == current {return func() Pair {return head}}
	}
	return nilHead
}
