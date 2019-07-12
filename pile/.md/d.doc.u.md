package pile // import "github.com/GoLangsam/twos/pile"

var TypeOf = core.TypeOf ...
var Cardinal = core.Cardinal
var Ordinal = core.Ordinal
func pileCardinality()
func pileID()
func pileIndex()
func pileInterface()
func pileIterable()
func pilePair()
func pilePile()
func pileType()
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
    func NewCardinalityPile(name string, items <-chan Cardinality) *PileOfCardinality
    func NewPileOfCardinality(name string, items ...Cardinality) *PileOfCardinality
type PileOfID struct{ ... }
    func NewIDPile(name string, items <-chan ID) *PileOfID
    func NewPileOfID(name string, items ...ID) *PileOfID
type PileOfIndex struct{ ... }
    func NewIndexPile(name string, items <-chan Index) *PileOfIndex
    func NewPileOfIndex(name string, items ...Index) *PileOfIndex
type PileOfInterface struct{ ... }
    func NewInterfacePile(name string, items <-chan interface{}) *PileOfInterface
    func NewPileOfInterface(name string, items ...interface{}) *PileOfInterface
type PileOfIterable struct{ ... }
    func NewIterablePile(name string, items <-chan Iterable) *PileOfIterable
    func NewPileOfIterable(name string, items ...Iterable) *PileOfIterable
type PileOfPair struct{ ... }
    func NewPairPile(name string, items <-chan Pair) *PileOfPair
    func NewPileOfPair(name string, items ...Pair) *PileOfPair
type PileOfPile struct{ ... }
    func NewPileOfPile(name string, items ...Pile) *PileOfPile
    func NewPilePile(name string, items <-chan Pile) *PileOfPile
type PileOfType struct{ ... }
    func NewPileOfType(name string, items ...Type) *PileOfType
    func NewTypePile(name string, items <-chan Type) *PileOfType
type PileS []Pile
type Tail = core.Tail
type Type = core.Type
type TypeS []Type
type interfaceS []interface{}
type lookUpCardinality struct{ ... }
type lookUpID struct{ ... }
type lookUpIndex struct{ ... }
type lookUpInterface struct{ ... }
type lookUpIterable struct{ ... }
type lookUpPair struct{ ... }
type lookUpPile struct{ ... }
type lookUpType struct{ ... }
type onesOfCardinality struct{ ... }
type onesOfID struct{ ... }
type onesOfIndex struct{ ... }
type onesOfInterface struct{ ... }
type onesOfIterable struct{ ... }
type onesOfPair struct{ ... }
type onesOfPile struct{ ... }
type onesOfType struct{ ... }
