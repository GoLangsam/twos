package pile // import "github.com/GoLangsam/twos/pile"

type PileS []Pile

func (a PileS) Both() (aten, apep interface{})
func (a PileS) Sort(less func(i, j int) bool) PileS
func (a PileS) String() string
