package pile // import "github.com/GoLangsam/twos/pile"

type lookUpPile struct{ look map[Pile]Index }

func (a lookUpPile) Idx(item Pile) (idx Index, found bool)
func (a lookUpPile) Len() int
func (a lookUpPile) Length() Cardinality
func (a lookUpPile) Random() <-chan Pile
func (a *lookUpPile) put(item Pile, idx Index)
