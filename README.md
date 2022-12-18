# `enum` in Go

I provide a package to simulate Rust `enum` feature in Go.
Also, you could see examples in the `example` folder.

But, it still has some problems:

* If defines two types with the same interface (e.g. `SumTypeZero`), then the type assertion would be ambiguous.

```go
type Type1 enum.SumTypeZero
type Type2 enum.SumTypeZero

func CheckType(typ enum.SumType) {
    if _, ok := typ.(Type1); ok {
        ...
        return 
    }
    if _, ok := typ.(Type2); ok {
        ...
        return
    }
}
```

Because both of `Type1` and `Type2` implemented `enum.SumTypeZero`, so the `CheckType` cannot distinguish both of them.

* When defines two different types with same interface and with same type parameter(s), it still has the same issue.

```go
type Type1 enum.SumTypeOne[int]
type Type2 enum.SumTypeOne[int]

func CheckType(typ enum.SumType) {
    if _, ok := typ.(Type1); ok {
        ...
        return 
    }
    if _, ok := typ.(Type2); ok {
        ...
        return
    }
}
```

In this implementation, `enum.SumType` is `interface{}`. As we know, any types in Go could be treated as `interface{}`.

IMO, we might could solve this problems, if we could: 

* Treat `named empty interface` and `unamed empty interface` different. 

```go
// (O) compiled
// (X) not compiled

// named empty interface
type NamedEmpty interface {}

// unamed empty interface 
interface{}

func AcceptNameEmpty(NamedEmpty) {}
func AcceptUnamedEmpty(interface{}) {}

type implNameEmpty struct{
    NamedEmpty
    ...
}

type notImplNamedEmpty struct { 
    ... 
}

AcceptNameEmpty(implNameEmpty{})       // Case 1 (O)
AcceptNameEmpty(&implNameEmpty{})      // Case 2 (O)
AcceptNameEmpty(notImplNamedEmpty{})   // Case 3 (X)
AcceptUnamedEmpty(implNameEmpty{})     // Case 4 (O) backward compatible
AcceptUnamedEmpty(notImplNamedEmpty{}) // Case 5 (O)
```

* Treat `named empty interface`s different.

```go
// (O) compiled
// (X) not compiled

type NamedEmpty1 interface{}

type NamedEmpty2 interface{}

func AcceptNamedEmpty1(NamedEmpty1) {}

func AcceptNamedEmpty2(NamedEmpty2) {}

type implNamedEmpty1 struct{
    NamedEmpty1
    ...
}

type implNamedEmpty2 struct{
    NamedEmpty2
    ...
}

type implBothNamedEmpty12 struct {
    NamedEmpty1
    NamedEmpty2
    ...
}

AcceptNamedEmpty1(implNamedEmpty1{})       // Case 1  (O)
AcceptNamedEmpty1(&implNamedEmpty1{})      // Case 2  (O)
AcceptNamedEmpty1(implNamedEmpty2{})       // Case 3  (X)
AcceptNamedEmpty1(&implNamedEmpty2{})      // Case 4  (X)
AcceptNamedEmpty1(implBothNamedEmpty12{})  // Case 5  (O)
AcceptNamedEmpty1(&implBothNamedEmpty12{}) // Case 6  (O)

AcceptNamedEmpty2(implNamedEmpty1{})       // Case 7  (X)
AcceptNamedEmpty2(&implNamedEmpty1{})      // Case 8  (X)
AcceptNamedEmpty2(implNamedEmpty2{})       // Case 9  (O)
AcceptNamedEmpty2(&implNamedEmpty2{})      // Case 10 (O)
AcceptNamedEmpty2(implBothNamedEmpty12{})  // Case 11 (O)
AcceptNamedEmpty2(&implBothNamedEmpty12{}) // Case 12 (O)
```

* Treat `named empty interface` and `type alias interface` different.
```go
// (O) compiled
// (X) not compiled

type NamedEmpty1 interface{}

type NamedEmpty2 NamedEmpty1

func AcceptNamedEmpty1(NamedEmpty1) {}

func AcceptNamedEmpty2(NamedEmpty2) {}

type implNamedEmpty1 struct{
    NamedEmpty1
    ...
}

type implNamedEmpty2 struct{
    NamedEmpty2
    ...
}

type implBothNamedEmpty12 struct {
    NamedEmpty1
    NamedEmpty2
    ...
}

AcceptNamedEmpty1(implNamedEmpty1{})       // Case 1  (O)
AcceptNamedEmpty1(&implNamedEmpty1{})      // Case 2  (O)
AcceptNamedEmpty1(implNamedEmpty2{})       // Case 3  (X)
AcceptNamedEmpty1(&implNamedEmpty2{})      // Case 4  (X)
AcceptNamedEmpty1(implBothNamedEmpty12{})  // Case 5  (O)
AcceptNamedEmpty1(&implBothNamedEmpty12{}) // Case 6  (O)

AcceptNamedEmpty2(implNamedEmpty1{})       // Case 7  (X)
AcceptNamedEmpty2(&implNamedEmpty1{})      // Case 8  (X)
AcceptNamedEmpty2(implNamedEmpty2{})       // Case 9  (O)
AcceptNamedEmpty2(&implNamedEmpty2{})      // Case 10 (O)
AcceptNamedEmpty2(implBothNamedEmpty12{})  // Case 11 (O)
AcceptNamedEmpty2(&implBothNamedEmpty12{}) // Case 12 (O)
```

In this proposal, which we didn't discuss:
1. `tilde`(~T) which I'm so sure would it be impacted? 
