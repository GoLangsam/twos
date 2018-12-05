package core // import "github.com/GoLangsam/twos/internal"

Package core provides generic implementations such as

    - bool.go for boolean operations (And, Or, Not) on attribute functions
    - fmap.go for Fmap... and Join... for slices and read-only channels
    - pile.go for PileOf...

as well as

    - pile_test.go for basic tests of PileOf... functionalities

This package is not imported anwhere. See genny.go elsewhere for it's usage.

Note: anyType_test.go provides some test data - any client package must
provide similar data, appropriate for the target type, if pile_test.go is
generated into such package.

func FmapanyTypeRoC(f func(anyType) anyType, RoCs ...<-chan anyType) <-chan anyType
func StringOfOnes(a interface{}) string
func StringOfPair(a Pair) string
func StringOfTwos(a, b interface{}) string
func FmapanyTypes(f func(anyType) anyType, anyTypes ...anyType) (anyTypeS []anyType)
func JoinanyTypeS(anyTypeSS [][]anyType) (anyTypeS []anyType)
type Cardinality = core.Cardinality
type Container interface{ ... }
type Head = core.Head
type ID = core.ID
type Index = core.Index
type Indexed interface{ ... }
type Iterable interface{ ... }
type LookeranyType interface{ ... }
type Named interface{ ... }
type Pair = core.Pair
type Pairs <-chan Pair
type Pile interface{ ... }
type PileOfanyType struct{ ... }
    func NewPileOfanyType(name ID, items ...anyType) *PileOfanyType
type PilerOfanyType interface{ ... }
type Tail = core.Tail
    func NilTail() Tail
type Type = core.Type
