type Head func() Pair
    Head is a thunk which evaluates to a Pair.

func (a Head) Both() (aten, apep interface{})
func (a Head) Contains(item interface{}) (contains bool)
func (a Head) Fmap(f func(Head) Head) Head
func (a Head) FmapPair(f func(Pair) Pair) Head
func (a Head) Lift(as ...Head) []Head
func (a Head) Name() string
func (a Head) Of(index Index) Head
func (a Head) Size() Cardinality
func (a Head) String() string
func (a Head) Tail() Tail
