type onesOfanyType struct {
	ID
	Apep anyType
}

func (a onesOfanyType) Both() (aten, apep interface{})
func (a onesOfanyType) Contains(item interface{}) (contains bool)
func (a onesOfanyType) Of(index Index) Head
func (a onesOfanyType) Size() Cardinality
func (a onesOfanyType) String() string
func (a onesOfanyType) Tail() Tail
