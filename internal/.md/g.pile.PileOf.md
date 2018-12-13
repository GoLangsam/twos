package core // import "github.com/GoLangsam/twos/internal"

type PileOfanyType struct {
	ID
	anyTypeS

	lookUpanyType
	duplicates lookUpanyType
}
    PileOfanyType holds different items of some kind, accessible in order of
    arrival and by index as well as via reverse look-up: Idx(item) returns the
    index (if any). It may be seen as a finite ordered indexed immutable set.

    Intentionally there is no removal, neither are add nor append exported.


func NewPileOfanyType(name string, items ...anyType) *PileOfanyType
func (a PileOfanyType) At(idx Index) anyType
func (a PileOfanyType) Both() (aten, apep interface{})
func (a PileOfanyType) Contains(item interface{}) (contains bool)
func (a PileOfanyType) Duplicates() map[anyType]Index
func (a PileOfanyType) Fmap(f func(anyType) anyType) *PileOfanyType
func (a PileOfanyType) Idx(item anyType) (idx Index, found bool)
func (a PileOfanyType) Len() int
func (a PileOfanyType) Of(idx Index) Head
func (a PileOfanyType) Random() <-chan anyType
func (a PileOfanyType) Range() <-chan anyType
func (a PileOfanyType) S() []anyType
func (a PileOfanyType) Size() Cardinality
func (a PileOfanyType) Sort(less func(i, j int) bool) *PileOfanyType
func (a PileOfanyType) String() string
func (a PileOfanyType) Tail() Tail
func (a *PileOfanyType) add(item anyType) *PileOfanyType
func (a *PileOfanyType) append(items ...anyType) *PileOfanyType
func (a PileOfanyType) head(idx int) Head
func (a *PileOfanyType) put(item anyType, idx Index)
func (a PileOfanyType) tail(idx int) Tail
