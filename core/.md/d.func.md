func IsKind() (pairIs func(Pair) bool)
func IsNested() (pairIs func(Pair) bool)
func StringOfOnes(a interface{}) string
func StringOfPair(a Pair) string
func StringOfTwos(a, b interface{}) string
func KindOfPair(a Pair) kind
    func LengthOfPair(a Pair) (length Cardinality)
type Head func() Pair
    func At(i int) Index
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
