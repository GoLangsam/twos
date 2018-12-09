package pile // import "github.com/GoLangsam/twos/pile"

type PileOfPile struct {
	ID
	PileS

	lookUpPile
	duplicates lookUpPile
}
    PileOfPile holds different items of some kind, accessible in order of
    arrival and by index as well as via reverse look-up: Idx(item) returns the
    index (if any). It may be seen as a finite ordered indexed immutable set.

    Intentionally there is no removal, neither are add nor append exported.


func NewPileOfPile(name string, items ...Pile) *PileOfPile
func (a PileOfPile) At(idx Index) Pile
func (a PileOfPile) Both() (aten, apep interface{})
func (a PileOfPile) Contains(item interface{}) (contains bool)
func (a PileOfPile) Duplicates() map[Pile]Index
func (a PileOfPile) Fmap(f func(Pile) Pile) *PileOfPile
func (a PileOfPile) Idx(item Pile) (idx Index, found bool)
func (a PileOfPile) Len() int
func (a PileOfPile) Length() Cardinality
func (a PileOfPile) Of(idx Index) Head
func (a PileOfPile) Random() <-chan Pile
func (a PileOfPile) Range() <-chan Pile
func (a PileOfPile) S() []Pile
func (a PileOfPile) Sort(less func(i, j int) bool) *PileOfPile
func (a PileOfPile) String() string
func (a PileOfPile) Tail() Tail
func (a *PileOfPile) add(item Pile) *PileOfPile
func (a *PileOfPile) append(items ...Pile) *PileOfPile
func (a PileOfPile) head(idx int) Head
func (a *PileOfPile) put(item Pile, idx Index)
func (a PileOfPile) tail(idx int) Tail
