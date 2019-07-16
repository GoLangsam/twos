# Twos

Package `twos` is all about pairs - ordered pairs, to be specific.

And pairs of pairs. And pairs of pairs of pairs. And pairs of pairs of pairs of pairs...

Plenty of pairs. Piles of pairs. And piles of piles of pairs. And piles of piles of piles of pairs...

You know about that nice old lady, who -when asked what the universe rests upon- said:
"Turtles! Turtles all the way down!", do You not? Well, relax: there are no turtles in here.

## `Pair`

Any `Pair` has two sides, here called `Aten` & `Apep` in respect to ancient mythology.

> Many other terms come to mind: _First/Second Left/Right Up/Down One/Two ..._
> each coming with particular (sometimes subconscious) and likely confusing connotations.
> 
> `Aten` & `Apep` aim to avoid subconscious confusion by confusing the reader explicitly ;-)

Note: _Inti_ is what the ancient Inca of America call the ancient Egypt `Aten`.
And `Aten` is sometimes spelled _Aton_, but that may easily be confused with the Greek term _Atom_.
Thus: This spelling is not used here.

Together, `Aten` & `Apep` may be understood to represent a very fundamental _polarity_,
arosen upon the burst of _Anda_, the cosmic egg.

Ancient mythology `:-)`

### Characteristic: Identity

The characteristic (or defining property) of ordered pairs is:

Two Pairs `p1 := {a1|b1}` and `p2 := {a2|b2}` are identical iff both respective sides are:

```go
	p1 == p2   <==>   a1 == a2 && b1 == b2
```

### Ordered pair

Order matters:

```go
	{a|b} == {b|a}   <==>   a == b.
```

### No pairs here

Just: there are no pairs here! (And no turtles.) Nowhere. Pairs do not exist.

Alas: Plenty of the many things which exist here know how to behave as a pair!

( Or to camouflage - if You wish. To pretend. )

And something behaves as a pair iff it has a simple method named `.Both()`.

A method which returns the two sides -untyped- as `aten, apep interface{}`.

Thus: `Pair` is an interface. Anything which satisfies it is considered a `Pair`.

```go 
type Pair interface {
	Both() (aten, apep interface{})
}
```

### Implementations

Different implementations of `Pair` contain (and return) different types.

( Technically, the two underlying types may be understood as the signature of the pair,
  a pair of `reflect.Type`. )

Certain combinations are predefined - be it for ease of use, be it for useful discriminations.

"Pair of Pairs" is among them as is "pair of `reflect.Type`".

By convention, a Pair implemented by a `struct` exposes its sides as fields `Aten` & `Apep`,
thus permitting type-safe access of each side.

---
## Tail - a Pair-Sequence

A `Tail` represents a sequence of `Pair`.

A `Tail` is a function which returns a `Head` and (another) `Tail`, where a `Head` is a function which returns a `Pair`.

Sooner or later, some `Tail` shall return `nil` for both functions, for `Head` and for `Tail`.
Note: A `Tail` must never return one nil only (and some non-nil function).

A `Tail` allows to iterate thru such collection.
Let `tail` be a `Tail`:
```go
	for head, tail := tail(); head != nil; head, tail = tail() {
		// do sth with head, e.g. obtain its pair
		myPair := head()
		_ = myPair // do sth with myPair
	}
```

> Note: Heavy use of functions is a necessity as it allows to circumvent stupid constraints imposed by the underlying language `Go`.

### `Iterable` is what provides a `Tail` function

Note: Most predefined pairs implement `Iterable`: a method `Tail()` which returns a `Tail` function.

Thus: Any such `Pair` may be iterated as it manifests a single-element sequence as well.

### `.Range()` for easy iteration and filtering 

`Tail` has a method `.Range() <-chan Pair` for easier iteration.
Let `tail` be a `Tail`:
```go
	for pair := range tail.Range() {
		// do sth with pair
		_ = pair
	}
```

`Range` is variadic - it accepts any number of predicates - unary functions `func(Pair) bool` (see below).

Used this way, `Range` acts as an `and`-filter: any `Pair` must satisfy each and every predicate.

Note: The beauty and versatility of `Range` comes at a price: a channel and a go routine are involved.

Thus: Prefer the `for head, tail := tail() ...` whenever speed is an issue.

### A taylor with a needle: Stitch & Thread - a metaphore

If You like to see Yourself as a taylor busy sewing and weaving a fabric of growing beauty and complexity, then
You may like to think of `Tail` as a thread and `head` as something which enables You to stitch:

Whenever You apply it, You get an applicable stitch in one hand a the remaining thread in the other.
Just apply the stitch and keep sewing...

---
## noteworthy Pairs - combinations

### `twoTwo` - a Pair of Pairs, Two Pairs, a Twos

An important combination is the pair of pairs, which allows to nest pairs to arbitrary depth.

`IsNested(a Pair)` discriminates iff the Pair is nested on one side at least.

Note: boolean convenience functions (predicates) such as IsNested, IsAtomAten, IsAtomApep facilitate further discriminations.

### `typTyp` - a Pair of Types - the signature of the `Pair`

---
## Some non-pair `Pair`

Some non-structs, typically certain function types, also love to mimic `Pair`, even so there are no fields named `Aten` & `Apep`.

And: They also implement `Iterable` ... Thus: By construction: They are very versatile. 

### `Tail`
Apply `.Both()` to receive both its `head` and `tail` functions.

### `Head`
Apply `.Both()` to receive both sides of the `Pair` it represents.

Note: `Head` represents a `Pair` (the `Pair` within) and also behaves as an iterable `Pair` by itself.

---
## Discriminate

### unary boolean - Predicate

```go
type PairIs = func (Pair) bool
```

Predefined	- discriminates iff:
- IsNested	- the `Pair` is nested on one side at least.
 - IsAtomApep	- the `Pair` ain't nested on its apep side.
 - IsAtomAten	- the `Pair` ain't nested on its aten side.
- IsSymmetric	- the `Pair` has the same `reflect.Type` on `.Both()` sides
- IsUnit	- `.Both()` sides are equal (implies: have same Type)

### binary boolean - Relation

```go
type PairsAre = func (a, b Pair) bool
```

Predefined:
- SameSame - two Pairs of same type TODO
- AreEqual - SameSame && Aten == Apep

---
