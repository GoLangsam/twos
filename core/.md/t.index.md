package core // import "github.com/GoLangsam/twos/core"

type Index = ordinalNumber
    Index represents an ordinal number and may be used to enumerate a collention
    or access some item. Intentionally, the first item has Index = 1. This is
    more intuitive for users. (Well, programmers prefer offsets over ordinals).


func At(i int) Index
func (a Index) AsOffset() int
func (a Index) Both() (aten, apep interface{})
func (a Index) Cmp(b Index) (r int)
func (a Index) Contains(item interface{}) (contains bool)
func (a Index) Name() string
func (a Index) Of(index Index) Head
func (a Index) Size() Cardinality
func (a Index) String() string
func (a Index) Tail() Tail
func (a *Index) Unit() Index
