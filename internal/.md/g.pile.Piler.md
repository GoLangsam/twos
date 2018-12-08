package core // import "github.com/GoLangsam/twos/internal"

type pilerOfanyType interface {
	Pair
	Tail() Tail                // Iterable
	Name() string              // Named
	String() string            // Stringer
	Of(Index) Head             // Indexed
	Contains(interface{}) bool // Container

	At(Index) anyType
	Duplicates() map[anyType]Index
	Fmap(f func(anyType) anyType) *PileOfanyType
	Range() <-chan anyType
	S() []anyType
	Sort(less func(i, j int) bool) *PileOfanyType

	lookerOfanyType // Pile: Length()

	add(item anyType) *PileOfanyType
	append(items ...anyType) *PileOfanyType
}
    pilerOfanyType is implemented by *PileOfanyType

