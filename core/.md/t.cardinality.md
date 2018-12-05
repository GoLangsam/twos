package core // import "github.com/GoLangsam/anda/twos/core"

type Cardinality = cardinalNumber
    Cardinality represents a cardinal number such as the #-of items in a Pile.


func LengthOfPair(a Pair) (length Cardinality)
func (a Cardinality) Both() (aten, apep interface{})
func (a Cardinality) Length() Cardinality
func (a Cardinality) Name() ID
func (a Cardinality) Of(index Index) Head
func (a Cardinality) String() string
func (a Cardinality) Tail() Tail
