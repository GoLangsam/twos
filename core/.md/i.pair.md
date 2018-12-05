package core // import "github.com/GoLangsam/anda/twos/core"

type Pair interface {
	Both() (aten, apep interface{}) // both sides - whatever each type is
}
    Pair has two sides: Aten & Apep. It may be atomic, or composed, nested.


func Apep(a Pair) Pair
func Aten(a Pair) Pair
func Join(a, b Pair) Pair
func Swap(a Pair) Pair
