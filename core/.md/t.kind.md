package core // import "github.com/GoLangsam/anda/twos/core"

type kind struct {
	ID
	Type
}
    kind implements a named type


func KindOfPair(a Pair) kind
func (a kind) Both() (aten, apep interface{})
func (a kind) Length() Cardinality
func (a kind) Name() ID
func (a kind) Of(index Index) Head
func (a kind) String() string
func (a kind) Tail() Tail
