package core // import "github.com/GoLangsam/twos/internal"

type pilerOfanyType interface {
	Pair
	Tail() Tail                // Iterable
	Name() string              // Named
	String() string            // Stringer
	Of(Index) Head             // Indexed
	Contains(interface{}) bool // Container

	At(Index) anyType
	Range() <-chan anyType
	Sort(less func(i, j int) bool) *PileOfanyType

	lookerOfanyType // Pile: Length()
	append(items ...anyType) *PileOfanyType
	add(item anyType) *PileOfanyType
}
    pilerOfanyType is implemented by *PileOfanyType

