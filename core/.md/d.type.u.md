type Cardinality = cardinalNumber
type Head func() Pair
type HeadS []Head
type ID = name
type Index = ordinalNumber
type Iterable interface{ ... }
type Pair interface{ ... }
type Tail func() (Head, Tail)
type TailS []Tail
type Type = reflect.Type
type cardinalNumber uint64
type kind struct{ ... }
type name string
type nest struct{ ... }
type nilPair struct{}
type ordinalNumber uint64
type stringer interface{ ... }
