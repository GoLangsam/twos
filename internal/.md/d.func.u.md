    - bool.go for boolean operations (And, Or, Not) on attribute functions
    - pile_test.go for basic tests of PileOf... functionalities
var isanyTypeFalse = func(a anyType) bool { ... }
var isanyTypeTrue = func(a anyType) bool { ... }
func FmapanyTypeRoC(f func(anyType) anyType, RoCs ...<-chan anyType) <-chan anyType
func StringOfOnes(a interface{}) string
func StringOfPair(a Pair) string
func StringOfTwos(a, b interface{}) string
func assertPileOfanyTypeInterfaces()
func boolanyType()
func fmapanyType()
func pileanyType()
    func NewPileOfanyType(name ID, items ...anyType) *PileOfanyType
    func NilTail() Tail
    func FmapanyTypes(f func(anyType) anyType, anyTypes ...anyType) (anyTypeS []anyType)
    func JoinanyTypeS(anyTypeSS [][]anyType) (anyTypeS []anyType)
type anyTypeIs func(anyType) bool
