package core // import "github.com/GoLangsam/twos/core"

type Cardinality = cardinalNumber
    Cardinality represents a cardinal number such as the #-of items in a Pile.


func SizeOfPair(a Pair) (size Cardinality)
func (a Cardinality) Both() (aten, apep interface{})
func (a Cardinality) Cmp(b Cardinality) (r int)
func (a Cardinality) Contains(item interface{}) (contains bool)
func (a Cardinality) Lift(as ...Cardinality) []Cardinality
func (a Cardinality) Name() string
func (a Cardinality) Of(index Index) Head
func (a Cardinality) Size() Cardinality
func (a Cardinality) String() string
func (a Cardinality) Tail() Tail
func (a *Cardinality) Unit() Cardinality
