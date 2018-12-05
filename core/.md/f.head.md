package core // import "github.com/GoLangsam/anda/twos/core"

type Head func() Pair
    Head is a thunk which evaluates to a Pair.


var nilHead Head = func() Pair { ... }
func (a Head) Both() (aten, apep interface{})
func (a Head) Fmap(f func(Head) Head) Head
func (a Head) FmapPair(f func(Pair) Pair) Head
func (a Head) Length() Cardinality
func (a Head) Name() ID
func (a Head) Of(index Index) Head
func (a Head) String() string
func (a Head) Tail() Tail
