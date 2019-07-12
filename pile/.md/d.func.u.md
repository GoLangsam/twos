func pileCardinality()
func pileID()
func pileIndex()
func pileInterface()
func pileIterable()
func pilePair()
func pilePile()
func pileType()
    func NewCardinalityPile(name string, items <-chan Cardinality) *PileOfCardinality
    func NewPileOfCardinality(name string, items ...Cardinality) *PileOfCardinality
    func NewIDPile(name string, items <-chan ID) *PileOfID
    func NewPileOfID(name string, items ...ID) *PileOfID
    func NewIndexPile(name string, items <-chan Index) *PileOfIndex
    func NewPileOfIndex(name string, items ...Index) *PileOfIndex
    func NewInterfacePile(name string, items <-chan interface{}) *PileOfInterface
    func NewPileOfInterface(name string, items ...interface{}) *PileOfInterface
    func NewIterablePile(name string, items <-chan Iterable) *PileOfIterable
    func NewPileOfIterable(name string, items ...Iterable) *PileOfIterable
    func NewPairPile(name string, items <-chan Pair) *PileOfPair
    func NewPileOfPair(name string, items ...Pair) *PileOfPair
    func NewPileOfPile(name string, items ...Pile) *PileOfPile
    func NewPilePile(name string, items <-chan Pile) *PileOfPile
    func NewPileOfType(name string, items ...Type) *PileOfType
    func NewTypePile(name string, items <-chan Type) *PileOfType
