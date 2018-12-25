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
type cardinalNumber struct{ ... }
type kind struct{ ... }
type name string
type nest struct{ ... }
type nilPair struct{}
type ordinalNumber struct{ ... }
type stringer interface{ ... }
