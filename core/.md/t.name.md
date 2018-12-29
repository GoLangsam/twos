package core // import "github.com/GoLangsam/twos/core"

type name string
    name as a unique identifier - unique among its Kind.


func (a name) Both() (aten, apep interface{})
func (a name) Contains(item interface{}) (contains bool)
func (a name) IsEq() Predicate
func (a name) IsLess() Predicate
func (*name) IsNul() Predicate
func (a name) Lift(as ...name) []name
func (a name) Name() string
func (a name) Of(index Index) Head
func (a name) Size() Cardinality
func (a name) String() string
func (a name) Tail() Tail
func (*name) Unit() name
