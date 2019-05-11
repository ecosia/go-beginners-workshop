# Go beginners workshop

## Setting up your environment

### Download go
To get started, go to the [Go download page](https://golang.org/dl/) and download the binary distribution according to your operation system.

// TODO: set up for each environment + check installation

### Text editor suggestions
// TODO: brief description

[Visual Studio Code](https://code.visualstudio.com/)

[Atom](https://atom.io/)

## Language mechanics

### Variables
// TODO

### Data types
// TODO: strings, integers

#### Structs
Structs are a collection of fields. Fields can be accessed using a dot; there's no need to implement getters and setters. Note that only exported fields (capitalized) can be accessed outside a package.

```go
// This `festival` struct type has name, latitude, longitude and date fields
type festival struct {
    name string
    lat float64
    lng float64
    date time.Time
}

// Creates and initializes a variable of type `festival`
primaveraSound := festival{
    name: "Primavera sound",
    lat: 41.385063,
    lng: 2.173404,
    date: time.Now()
}
```

Go does not provide inheritance as in common objected oriented languages, but it is possible to use the notion of composition instead. You can "borrow" pieces of a certain implementation and compose them in a struct/interface. For example, we can create a new `location` struct with `lat, lng` fields and embed it on the `festival` struct:

```go
type location struct {
    lat float64
    lng float64
}

type festival struct {
    location
    name string
    date time.Time
}
```

[Playground](https://play.golang.org/p/u2F8GUVU4vI)

### Functions and scopes
// TODO

### Data structures

#### Arrays
In Go, an array is a collection of elements that have the same type. This collection has a fixed size: once you declare an array, its size cannot increase or decrease.

```go
numbers := [5]int{5, 4, 3, 2, 1}
```
Above we have declared an array with 5 elements, in which each element is an integer.

[Playground](https://play.golang.org/p/k0-S_16N6tV)

#### Slices
Slices are much more used than arrays in Go and they are more flexible, since unlike arrays, they can be resized. They also provide a powerful interface for managing collections and support built-ins, like `append` and `make`. Slices are declared just like arrays, with the only difference being that you don't need to specify the length:

```go
letters := []string{"a", "b", "c", "d"}
```

The same slice can be created using the built-in function `make`, that creates an array in memory and returns a slice corresponding to that array:

```go
newLetters := make([]string, len(letters))
```

We can also copy a slice into a new one with the same length:

```go
copy(newLetters, letters)
```

It's also possible to create an slice from an array:

```go
numbers := [5]int{5, 4, 3, 2, 1}
numbersSlice := numbers[:]
```

[Playground](https://play.golang.org/p/4fNoR5hMm4h)


### Concurrency
