type anyTypeS []anyType

func (a anyTypeS) Both() (aten, apep interface{})
func (a anyTypeS) Len() int
func (a anyTypeS) Size() Cardinality
func (a anyTypeS) Sort(less func(i, j int) bool) anyTypeS
