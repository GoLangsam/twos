package pile // import "github.com/GoLangsam/twos/pile"

var TypeOf = core.TypeOf ...
type Cardinality = core.Cardinality
type CardinalityS []Cardinality
type Container interface{ ... }
type Head = core.Head
type ID = core.ID
type IDS []ID
type Index = core.Index
type IndexS []Index
type Indexed interface{ ... }
type Iterable interface{ ... }
type IterableS []Iterable
type Named interface{ ... }
type Pair = core.Pair
type PairS []Pair
type Pile interface{ ... }
type PileOfCardinality struct{ ... }
    func NewPileOfCardinality(name ID, items ...Cardinality) *PileOfCardinality
type PileOfID struct{ ... }
    func NewPileOfID(name ID, items ...ID) *PileOfID
type PileOfIndex struct{ ... }
    func NewPileOfIndex(name ID, items ...Index) *PileOfIndex
type PileOfInterface struct{ ... }
    func NewPileOfInterface(name ID, items ...interface{}) *PileOfInterface
type PileOfIterable struct{ ... }
    func NewPileOfIterable(name ID, items ...Iterable) *PileOfIterable
type PileOfPair struct{ ... }
    func NewPileOfPair(name ID, items ...Pair) *PileOfPair
type PileOfPile struct{ ... }
    func NewPileOfPile(name ID, items ...Pile) *PileOfPile
type PileOfType struct{ ... }
    func NewPileOfType(name ID, items ...Type) *PileOfType
type PileS []Pile
type Tail = core.Tail
type Type = core.Type
type TypeS []Type
