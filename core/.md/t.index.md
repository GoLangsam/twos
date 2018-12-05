package core // import "github.com/GoLangsam/anda/twos/core"

type Index = ordinalNumber
    Index represents an ordinal number and may be used to enumerate a collention
    or access some item. Intentionally, the first item has Index = 1. This is
    more intuitive for users. (Well, programmers prefer offsets over ordinals).


func At(i int) Index
func (a Index) AsOffset() int
func (a Index) Both() (aten, apep interface{})
func (a Index) Length() Cardinality
func (a Index) Name() ID
func (a Index) Of(index Index) Head
func (a Index) String() string
func (a Index) Tail() Tail
