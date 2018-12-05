    - bool.go for boolean operations (And, Or, Not) on attribute functions
    - pile_test.go for basic tests of PileOf... functionalities
func FmapanyTypeRoC(f func(anyType) anyType, RoCs ...<-chan anyType) <-chan anyType
func StringOfOnes(a interface{}) string
func StringOfPair(a Pair) string
func StringOfTwos(a, b interface{}) string
func FmapanyTypes(f func(anyType) anyType, anyTypes ...anyType) (anyTypeS []anyType)
func JoinanyTypeS(anyTypeSS [][]anyType) (anyTypeS []anyType)
    func NewPileOfanyType(name ID, items ...anyType) *PileOfanyType
    func NilTail() Tail
