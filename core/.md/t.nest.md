type nest struct {
	Aten Pair
	Apep Pair
}
    nest represents a nested pair, a pair of pairs.

func (a nest) Both() (aten, apep interface{})
func (a nest) Contains(item interface{}) (contains bool)
func (a nest) Lift(as ...nest) []nest
func (a nest) Name() string
func (a nest) Of(index Index) Head
func (a nest) Size() Cardinality
func (a nest) String() string
func (a nest) Tail() Tail
