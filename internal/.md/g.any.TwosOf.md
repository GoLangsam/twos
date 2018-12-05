package core // import "github.com/GoLangsam/twos/internal"

type twosOfanyType struct {
	Aten anyType
	Apep anyType
}

func (a twosOfanyType) Both() (aten, apep interface{})
func (a twosOfanyType) Length() Cardinality
func (a twosOfanyType) Of(index Index) Head
func (a twosOfanyType) String() string
func (a twosOfanyType) Tail() Tail
