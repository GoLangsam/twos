package core // import "github.com/GoLangsam/anda/twos/core"

type nest struct {
	Aten Pair
	Apep Pair
}
    nest represents a nested pair, a pair of pairs.


func (a nest) Both() (aten, apep interface{})
func (a nest) Contains(item interface{}) (contains bool)
func (a nest) Length() Cardinality
func (a nest) Name() ID
func (a nest) Of(index Index) Head
func (a nest) String() string
func (a nest) Tail() Tail
