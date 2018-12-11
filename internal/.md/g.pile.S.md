package core // import "github.com/GoLangsam/twos/internal"

type anyTypeS []anyType

func (a anyTypeS) Both() (aten, apep interface{})
func (a anyTypeS) Len() int
func (a anyTypeS) Length() Cardinality
func (a anyTypeS) Sort(less func(i, j int) bool) anyTypeS
func (a anyTypeS) String() string
