type Cardinality = cardinalNumber
    Cardinality represents a cardinal number such as the #-of items in a Pile.

func Cardinal(a int) Cardinality
func SizeOfPair(a Pair) (size Cardinality)
func (a Cardinality) Add(x, y Cardinality) Cardinality
func (a Cardinality) AsInt() int
func (a Cardinality) Both() (aten, apep interface{})
func (a Cardinality) Contains(item interface{}) (contains bool)
func (a Cardinality) IsEq() Predicate
func (a Cardinality) IsEqInt() func(int) bool
func (a Cardinality) IsLess() Predicate
func (*Cardinality) IsNul() Predicate
func (*Cardinality) IsOne() Predicate
func (a Cardinality) Lift(as ...Cardinality) []Cardinality
func (a Cardinality) Mul(x, y Cardinality) Cardinality
func (a Cardinality) Name() string
func (a Cardinality) Of(index Index) Head
func (a Cardinality) Size() Cardinality
func (a Cardinality) String() string
func (a Cardinality) Tail() Tail
func (*Cardinality) Unit() Cardinality
