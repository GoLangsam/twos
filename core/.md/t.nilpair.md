package core // import "github.com/GoLangsam/twos/core"

type nilPair struct{}

func (a nilPair) Both() (aten, apep interface{})
func (a nilPair) Contains(item interface{}) (contains bool)
func (a nilPair) Length() Cardinality
func (a nilPair) Name() string
func (a nilPair) Of(index Index) Head
func (a nilPair) String() string
func (a nilPair) Tail() Tail
