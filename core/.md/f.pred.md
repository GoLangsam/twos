package core // import "github.com/GoLangsam/twos/core"

type Predicate func(Pair) bool
    Predicate is the type of a Pair predicate.


func Is(a func(Pair) bool) Predicate
