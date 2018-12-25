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

	theKindOfNest        = kind{name(TypeOf(nest{}).Name()),      TypeOf(nest{})}
	theKindOfKind        = kind{name(TypeOf(kind{}).Name()),      TypeOf(kind{})}
	theKindOfName        = kind{name(TypeOf(name("")).Name()),    TypeOf(name(""))}
	theKindOfIndex       = kind{name(TypeOf(Ordinal(1)).Name()),  TypeOf(Ordinal(1))}
	theKindOfCardinality = kind{name(TypeOf(Cardinal(0)).Name()), TypeOf(Cardinal(0))}
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

// SizeOfPair reports the length (as a Cardinality).
func SizeOfPair(a Pair ) (size Cardinality) {
	// Pile holds Size items.
	type Pile interface {
		Size() Cardinality
	}
	if p, ok := a.(Pile); ok { return p.Size() }
	return
}

/* Size implements Pile by returning 0 */func (a nilPair)     Size() Cardinality { return Cardinal(0) }
/* Size implements Pile by returning 1 */func (a kind)        Size() Cardinality { return Cardinal(1) }
/* Size implements Pile by returning 1 */func (a name)        Size() Cardinality { return Cardinal(1) }

// Size implements Pile by returning 1.
func (a Index)       Size() Cardinality { return Cardinal(1) }

// Size implements Pile by returning 1
func (a Cardinality) Size() Cardinality { return Cardinal(1) }

// Size implements Pile
// by returning
// the sum of the length of the two pairs.
func (a nest) Size() Cardinality {
	size := Cardinal(0)
	size = size.Add(SizeOfPair(a.Aten), SizeOfPair(a.Apep))
	return size
}

// Size implements Pile
// by returning
// the length of the pair a evaluates to
// or zero for nil
func (a Head) Size() Cardinality {
	if a == nil { return Cardinal(0) }
	if a() == nil { return Cardinal(1) }
	return SizeOfPair(a())
}

// Size implements Pile
// by incrementing upon traversing a.
//  Note: Complexity is O(n) due to complete traversal.
func (a Tail) Size() Cardinality {
	size, step := Cardinal(0), Cardinal(1)
	for head, tail := a(); head != nil; head, tail = tail() {
		size = size.Add(size, step)
	}
	return size
}

// SizeRecursive is a recursive implementation to determine the Size
// by returning zero for nil and 1 plus the recursive length of the tail a evaluates to.
//  Note: Complexity is O(n) due to recursion.
func (a Tail) SizeRecursive() Cardinality {
	head, tail := a()
	size, step := Cardinal(0), Cardinal(1)
	if head == nil {
		return size
	}
	size = size.Add(step, tail.SizeRecursive())
	return size
}

// ===========================================================================

// Contains implements Container
// by telling whether the given item is of suitable type
// and if so, whether a contains this item.
func (a nilPair) Contains(item interface{}) (contains bool) {
	if i, contains := item.(nilPair); contains && i == a { return contains }
	return false
}

// Contains implements Container
// by telling whether the given item is of suitable type
// and if so, whether a contains this item.
func (a name) Contains(item interface{}) (contains bool) {
	if i, contains := item.(name); contains && i == a { return contains }
	return false
}

// Contains implements Container
// by telling whether the given item is of suitable type
// and if so, whether a contains this item.
func (a Index) Contains(item interface{}) (contains bool) {
	if i, contains := item.(Index); contains && i == a { return contains }
	return false
}

// Contains implements Container
// by telling whether the given item is of suitable type
// and if so, whether a contains this item.
func (a Cardinality) Contains(item interface{}) (contains bool) {
	if i, contains := item.(Cardinality); contains && i == a { return contains }
	return false
}

// Contains implements Container
// by telling whether the given item is of suitable type
// and if so, whether a contains this item.
func (a Head) Contains(item interface{}) (contains bool) {
	if a == nil { return item == nil }
	if i, contains := item.(Pair); contains && i == a() { return contains }
	return false
}

// Contains implements Container
// by telling whether the given item implements Pair
// and if so, whether some Head evaluates to a Pair which IsEqual to this item.
func (a Tail) Contains(item interface{}) (contains bool) {
	if a == nil { return item == nil }

	if pair, is := item.(Pair); is {
		eq := IsEqual(pair)
		for head, tail := a(); head != nil; head, tail = tail() {
			if eq(head()) { return true}
		}
	}
	return false
}

// ===========================================================================

// Of returns a nilHead.
func (a nilPair)     Of(index Index) Head { return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a kind)        Of(index Index) Head { if index.Cmp(Ordinal(1)) == 0 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a name)        Of(index Index) Head { if index.Cmp(Ordinal(1)) == 0 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a Index)       Of(index Index) Head { if index.Cmp(Ordinal(1)) == 0 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a Cardinality) Of(index Index) Head { if index.Cmp(Ordinal(1)) == 0 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a nest)        Of(index Index) Head { if index.Cmp(Ordinal(1)) == 0 {return func() Pair {return a}}; return nilHead }
// Of returns a of the Pair iff index is one and nilHead otherwise.
func (a Head)        Of(index Index) Head { if index.Cmp(Ordinal(1)) == 0 {return func() Pair {return a}}; return nilHead }

// Of returns a Head for the item of position index, of nilHead iff the index is out of bounds.
func (a Tail)        Of(index Index) Head {
	current, step := Ordinal(0), Ordinal(1)
	for head, tail := a(); head != nil && current.Cmp(index) < 0; head, tail = tail() {
		current = current.Add(current, step)
		if current.Cmp(index) == 0 {return func() Pair {return head}}
	}
	return nilHead
}

// ===========================================================================
