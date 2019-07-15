# TODO pile

## Cleanup

- Testdata for Iterable; genny,go

- see also ## Observations

- Head Tail String() gave recoursive trouble

- SizeOfPair is Zero iff it's not a Pile. This seems wrong!
- SizeOfPile (new) should do all the inspection!

- FAQ: Indexs, not Indices; Cardinalitys, not Cardinalities

- TODO: In order to extend a pile, we must make a new pile. What a pitty!

## Extend
- migrate to "github.com/ncw/gotemplate"

- func Of(indexes []Index, indices ...Indexed} Head
- change Contains into Predicate
- test Contains for all basic types

- C, I, Names as Tail

- Make Tree correct: Node(interface{}
- Node.Tail must walk the tree!

---
## Future maybe

- include PipeOf - from pipe.m for all? types? One package? One per type?
- allow for small and growing lookers

---
## Observations:
- Types: Cardinality, Index, Name
	- all: Kind/String
	- all: KindOf*()
	- all: some kind of range: C, I, N, Names
	- func At(i int) Index

- Kind: kind
	- Both/Kind/String
	- New:Kind/Type/Name					TODO: better names

- Thunk: Head, Tail
	- all: Both/Kind/Length/String/Tail			TODO: LengthRecursive benchmark
	- all: Fmap
	- Tail: FmapPair/Head, Join, Range

- Inter: Type, Kind, Pair, Pile, Interable


- PileOf:
	-ones: Both/Kind/Length/Tail				no Range + Fmap + more
	-twos: Both/----/Length/Tail				no Range + Fmap + more

	-pair: Both/Kind/Length/Tail/Range + Fmap + more

	-look: ----/----/Length/----/----- + Fmap + more	TODO:Range? Test, if external pack may range Duplicates
	-pile: Both/Kind/Length/Tail/Range + look*** + Fmap ...	TODO:Join: same+same, JoinPile: same+any

- Permuter.Length(): see math/big: func (z *Int) Binomial(n, k int64) *Int

---
## Package level functions:

### TODO: Which of these may go up into twos?

    func BothSameType(a Pair) bool
    func BothTypes(a Pair) (Type, Type)
    func TypePair(a Pair) Pair

    func Flat(a Pair) (values []interface{})

    func HaveSametree(p1, p2 Pair) (same bool)

    func IsAtom(a Pair) (IsAtom bool)
    func IsAtomApep(a Pair) (isAtomApep bool)
    func IsAtomAten(a Pair) (IsAtomAten bool)
    func IsNested(a Pair) (IsNested bool)
    func IsPairOfPairs(a Pair) (IsPairOfPairs bool)

    func join(a, b Pair) Pair	see Mult & Prod		=> rename PairOf(a, b)

    func Fmap(f func(Pair) Pair, tail Tail) Tail

    func Iter(a ...Pair) (tail Tail)

    func Mult(a, b Iterable) Tail
    func Prod(a, b Iterable) Tail

---
### keep?: would look odd, if no re-alias for Cardinality, Index, Name. But: need re-alias anyway!
    func C(N int) <-chan Cardinality
    func I(N int) <-chan Index
    func At(i int) Index
    func Idx(i Index) int
    func N(n int) []struct{}

### keep: internal need:
    func NilTail() Tail


### keep: generated:
    func FmapCardinalitys(f func(Cardinality) Cardinality, Cardinalitys ...Cardinality) []Cardinality
    func FmapIndexs(f func(Index) Index, Indexs ...Index) []Index
    func FmapInterfaces(f func(interface{}) interface{}, interfaces ...interface{}) []interface{}
    func FmapInts(f func(int) int, ints ...int) []int
    func FmapIterable(f func(Iterable) Iterable, Iterables ...Iterable) []Iterable
    func FmapKinds(f func(Kind) Kind, Kinds ...Kind) []Kind
    func FmapNames(f func(Name) Name, Names ...Name) []Name
    func FmapPairs(f func(Pair) Pair, Pairs ...Pair) []Pair
    func FmapPiles(f func(Pile) Pile, Piles ...Pile) []Pile
    func FmapStrings(f func(string) string, strings ...string) []string
    func FmapTimeWeekdays(f func(time.Weekday) time.Weekday, timeWeekdays ...time.Weekday) []time.Weekday
    func FmapTypes(f func(Type) Type, Types ...Type) []Type

    func Join----S(ss [][]----) []----

---
## Discarded ideas

- rename FmapIndexs		=> FmapIndices
- rename FmapCardinalitys	=> FmapCardinalities
=> No! it's generated! 
=> TODO:FAQ

- move twos/pile	=  twos/internal/pile ??? Better not: Nobody could use it!
- Ranger for TypeS ? makes little sense only! (Interface!)

- *kind! - implement special kind for each pile? More clear, may be.
  PileOf & onesOf have SAME *kind - is this ok?
  Better: Copy
  
  And register each into core a la image.RegisterFormat, see image/image.go

  v := ValueOf(???).Convert(t Type)
