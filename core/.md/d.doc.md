package core // import "github.com/GoLangsam/anda/twos/core"

func IsKind() (pairIs func(Pair) bool)
func IsNested() (pairIs func(Pair) bool)
func StringOfOnes(a interface{}) string
func StringOfPair(a Pair) string
func StringOfTwos(a, b interface{}) string
func KindOfPair(a Pair) kind
type Cardinality = cardinalNumber
    func LengthOfPair(a Pair) (length Cardinality)
type Head func() Pair
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
    func Mult(factors ...Iterable) (tail Tail)
    func NilTail() Tail
    func Only(iter Iterable, pairIs func(Pair) bool) Tail
    func Prod(a, b Iterable) (tail Tail)
    func Skip(iter Iterable, pairIs func(Pair) bool) Tail
type Type = reflect.Type
    func TypeOf(a interface{}) Type
