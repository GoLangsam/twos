type kind struct {
	ID
	Type
}
    kind implements a named type

func KindOfPair(a Pair) kind
func (a kind) Both() (aten, apep interface{})
func (a kind) Lift(as ...kind) []kind
func (a kind) Name() string
func (a kind) Of(index Index) Head
func (a kind) Size() Cardinality
func (a kind) String() string
func (a kind) Tail() Tail
