# FAQ - Frequently asked questions

## How to make a `Pair`

- How to make a `Pair`

> You better don't.

- Please excuse me???

> You better don't. There's not much fun having a single, lonely pair only.
>
> You'd better make a pile - a pile of similar pairs.

- Like `NewPile().Add(1, "a")`?

> No. Not like this.

- Please excuse me???

> There is no meaning in neither this digit `1` nor in this character `a`.
>
> You'd better be meaningful - express Your semantics up front.

- Please excuse me???

> Try `NewKind("Professor", "Tom")`

- And so???

> You get a _Kind_ - a named type.

- Ahh! So _now_ I have a `Pair` _{Professor|Tom}_, do I not?

> No. No, You don't. Sorry.
>
> You have a kind named "Professor" of type `string`, as that is the type of Your example argument "Tom".

- But I still do not have my `Pair` - do I?

> Yes. You do not have Your `Pair` yet.
>
> Just: You have another `Pair`, a pair which represents Your string-type named "Professor".

- How awfull! How awkward! How ugly!!!

> Well - there's a shortcut! A sweet and simple shortcut.

- Please show me!!! I love shortcuts!

> Try `profs := NewPile("Professor", "Tom", "Andreas", "Dr. Knuth", "Albert Einstein" )`

- And ... ???

> Now You have a pile of four pairs - each of the same type, and with a common semantic meaning.

- And ... ???

> You may extend it. Just be sure to use the same type.
>
> Try `profs = profs.Append( "Stephen Hawking" )`

- And ... ???

> Now You have a pile of five ...

- What is this thing with the type all about?

> Well, `profs.Append( 2011 )` will panic - 2011 is not a `string`. It's some integer constant.

- Ahh! I see.

> You like type safety, do You not? Avoids a plethora of stupid, unintended errors.

- Yes, I do! Sure.

> Fine `:-)`
