package core // import "github.com/GoLangsam/anda/twos/core"

type name string
    name as a unique identifier - unique among its Kind.


func (a name) Both() (aten, apep interface{})
func (a name) Length() Cardinality
func (a name) Name() ID
func (a name) Of(index Index) Head
func (a name) String() string
func (a name) Tail() Tail
