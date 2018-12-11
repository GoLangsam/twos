package core // import "github.com/GoLangsam/twos/core"

const itemFmt = "%+v"
const nilName = name("<noName>")
const nodeFmt = "{%+v<%+v>%+v}"
const tailBeg = "["
const tailEnd = "]"
const twosFmt = "{ %+v | %+v }"
var nilHead Head = func() Pair { ... } ...
var TypeOf = reflect.TypeOf
func IsEqual(a Pair) (pairIs func(Pair) bool)
func IsKind() (pairIs func(Pair) bool)
func IsNested() (pairIs func(Pair) bool)
func StringOfOnes(a interface{}) string
func StringOfPair(a Pair) string
func StringOfTwos(a, b interface{}) string
func prod(aHead Head, aTail, bTail, reset Tail) (head Head, tail Tail)
type Cardinality = cardinalNumber
    func LengthOfPair(a Pair) (length Cardinality)
type Head func() Pair
type HeadS []Head
type ID = name
type Index = ordinalNumber
    func At(i int) Index
type Iterable interface{ ... }
type Pair interface{ ... }
    func Apep(a Pair) Pair
    func Aten(a Pair) Pair
    func Join(a, b Pair) Pair
    func Swap(a Pair) Pair
type Tail func() (Head, Tail)
    func Iter(pairs ...Pair) (tail Tail)
    func NilTail() Tail
    func Only(iter Iterable, pairIs func(Pair) bool) Tail
    func Skip(iter Iterable, pairIs func(Pair) bool) Tail
    func X(factors ...Iterable) (tail Tail)
    func mult(a, b Iterable) (tail Tail)
    func tailRecurse(pairs ...Pair) (tail Tail)
type TailS []Tail
type Type = reflect.Type
type cardinalNumber uint64
type kind struct{ ... }
    func KindOfPair(a Pair) kind
type name string
type nest struct{ ... }
type nilPair struct{}
type ordinalNumber uint64
type stringer interface{ ... }
