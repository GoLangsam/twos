var nilHead Head = func() Pair { ... } ...
func IsEqual(a Pair) (pairIs func(Pair) bool)
func IsKind() (pairIs func(Pair) bool)
func IsNested() (pairIs func(Pair) bool)
func StringOfOnes(a interface{}) string
func StringOfPair(a Pair) string
func StringOfTwos(a, b interface{}) string
func prod(aHead Head, aTail, bTail, reset Tail) (head Head, tail Tail)
    func Cardinal(a int) Cardinality
    func SizeOfPair(a Pair) (size Cardinality)
type Head func() Pair
    func At(i int) Index
    func Ordinal(a int) Index
    func Apep(a Pair) Pair
    func Aten(a Pair) Pair
    func Join(a, b Pair) Pair
    func Swap(a Pair) Pair
type Predicate func(Pair) bool
    func Is(a func(Pair) bool) Predicate
type Tail func() (Head, Tail)
    func Iter(pairs ...Pair) (tail Tail)
    func NilTail() Tail
    func Only(iter Iterable, pairIs func(Pair) bool) Tail
    func Skip(iter Iterable, pairIs func(Pair) bool) Tail
    func X(factors ...Iterable) (tail Tail)
    func mult(a, b Iterable) (tail Tail)
    func tailRecurse(pairs ...Pair) (tail Tail)
    func KindOfPair(a Pair) kind
