package pile // import "github.com/GoLangsam/twos/pile"

type onesOfPile struct {
	ID
	Apep Pile
}

func (a onesOfPile) Both() (aten, apep interface{})
func (a onesOfPile) Contains(item interface{}) (contains bool)
func (a onesOfPile) Of(index Index) Head
func (a onesOfPile) Size() Cardinality
func (a onesOfPile) String() string
func (a onesOfPile) Tail() Tail
