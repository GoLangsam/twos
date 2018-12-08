package core // import "github.com/GoLangsam/twos/internal"

Package core provides generic implementations such as

    - bool.go for boolean operations (And, Or, Not) on predicates
    - fmap.go for Fmap... and Join... for slices and read-only channels
    - pile.go for PileOf...

as well as

    - pile_test.go for basic tests of PileOf... functionalities

This package is not imported anwhere. See genny.go elsewhere for it's usage.

Note: anyType_test.go provides some test data - any client package must
provide similar data, appropriate for the target type, if pile_test.go is
generated into such package.

var NilTail = core.NilTail ...
var isanyTypeFalse = func(a anyType) bool { ... }
var isanyTypeTrue = func(a anyType) bool { ... }
func FmapanyTypeRoC(f func(anyType) anyType, RoCs ...<-chan anyType) <-chan anyType
func assertPileOfanyTypeInterfaces()
func boolanyType()
func fmapanyType()
func pileanyType()
type Cardinality = core.Cardinality
type Head = core.Head
type ID = core.ID
type Index = core.Index
type Pair = core.Pair
type PileOfanyType struct{ ... }
    func NewPileOfanyType(name ID, items ...anyType) *PileOfanyType
type Tail = core.Tail
type anyType generic.Type
    func FmapanyTypes(f func(anyType) anyType, anyTypes ...anyType) (anyTypeS []anyType)
    func JoinanyTypeS(anyTypeSs ...[]anyType) (anyTypeS []anyType)
type anyTypeIs func(anyType) bool
type anyTypeS []anyType
type lookUpanyType struct{ ... }
type lookerOfanyType interface{ ... }
type onesOfanyType struct{ ... }
type pairOfanyType interface{ ... }
type pilerOfanyType interface{ ... }
