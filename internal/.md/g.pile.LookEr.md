package core // import "github.com/GoLangsam/twos/internal"

type lookerOfanyType interface {
	Idx(item anyType) (idx Index, found bool) // returns the index of item iff applicable
	Len() int                                 // current Length
	Length() Cardinality                      // current Length
	Random() <-chan anyType                   // a la Range, just: in random order
	put(item anyType, idx Index)              // may panic("Overflow")
}
    lookerOfanyType is implemented by lookUpanyType

