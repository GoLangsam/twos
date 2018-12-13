package core // import "github.com/GoLangsam/twos/core"

type Tail func() (Head, Tail)
    Tail is a thunk which evaluates to a Head and to another Tail.


func Iter(pairs ...Pair) (tail Tail)
func NilTail() Tail
func Only(iter Iterable, pairIs func(Pair) bool) Tail
func Skip(iter Iterable, pairIs func(Pair) bool) Tail
func X(factors ...Iterable) (tail Tail)
func mult(a, b Iterable) (tail Tail)
func tailRecurse(pairs ...Pair) (tail Tail)
func (a Tail) Both() (aten, apep interface{})
func (a Tail) Contains(item interface{}) (contains bool)
func (a Tail) Fmap(f func(Tail) Tail) Tail
func (a Tail) FmapHead(f func(Head) Head) Tail
func (a Tail) FmapPair(f func(Pair) Pair) Tail
func (a Tail) Join(tails ...Iterable) Tail
func (a Tail) Name() string
func (a Tail) Of(index Index) Head
func (a Tail) Only(pairIs func(Pair) bool) Tail
func (a Tail) Printer(stringer func(a Pair) string) Tail
func (a Tail) Range() <-chan Pair
func (a Tail) Reduce(f func(Pair, interface{}) interface{}, init interface{}) interface{}
func (a Tail) ReduceBool(f func(Pair, bool) bool, init bool) bool
func (a Tail) ReduceCardinality(f func(Pair, Cardinality) Cardinality, init Cardinality) Cardinality
func (a Tail) ReduceInt(f func(Pair, int) int, init int) int
func (a Tail) ReducePair(f func(Pair, Pair) Pair, init Pair) Pair
func (a Tail) ReduceString(f func(Pair, string) string, init string) string
func (a Tail) Sew(needle func(Pair)) Tail
func (a Tail) Size() Cardinality
func (a Tail) SizeRecursive() Cardinality
func (a Tail) Skip(pairIs func(Pair) bool) Tail
func (a Tail) String() string
func (a Tail) Tail() Tail
func (a Tail) X(iters ...Iterable) Tail
