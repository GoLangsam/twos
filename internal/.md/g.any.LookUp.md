package core // import "github.com/GoLangsam/anda/twos/internal"

type lookUpanyType struct{ look map[anyType]Index }

func (a lookUpanyType) Idx(item anyType) (idx Index, found bool)
func (a lookUpanyType) Len() int
func (a lookUpanyType) Length() Cardinality
func (a lookUpanyType) Random() <-chan anyType
func (a *lookUpanyType) put(item anyType, idx Index)
