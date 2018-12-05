package core // import "github.com/GoLangsam/anda/twos/internal"

type onesOfanyType struct {
	ID
	Apep anyType
}

func (a onesOfanyType) Both() (aten, apep interface{})
func (a onesOfanyType) Length() Cardinality
func (a onesOfanyType) Of(index Index) Head
func (a onesOfanyType) String() string
func (a onesOfanyType) Tail() Tail
