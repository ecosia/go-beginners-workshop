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
// TODO

#### Strings
// TODO

#### Integers
// TODO

#### Booleans
// TODO

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

### Loops
// TODO

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

Concurrency means that multiple functions/tasks make progress at the same time and run independently. Go has a scheduler that works directly with the operation system, scheduling functions that are created as goroutines.

#### Goroutines

Goroutines, as also called `lightweight threads`, are functions that run independently and are scheduled by the Go scheduler. A Go application starts with only one goroutine, which we call main goroutine. A goroutine can create new goroutines.

To call a function in a goroutine, simply use `go` in front of it:

```go
go myFunction()
```

#### Channels

Channels synchronize goroutines and make them communicate with each other. When declaring a channel, you'll need to specify what's the type of data you want your channel to be. To create a channel, you can make use of the `make` built-in:

```go
// Create a channel of strings
messages := make(chan string, 10)
```

You can send a value to a channel using `channel <-` syntax. To retrieve a value from a channel, the syntax `<-channel` is used.
```go
// Send a string to the channel, from a new goroutine
go func() { messages <- "Hey you" }()

// Get a string from the messages channel
message := <-messages
```

### Resources

To continue your learning journey with Go, here are some useful and informative resources to help you get started:

[Tour of Go](https://tour.golang.org/list)

[Women Who Go Berlin](https://github.com/wwgberlin)

[Go Study Group - WWG Berlin](https://github.com/wwgberlin/GoStudyGroup)

[Tutorials - WWG Berlin](https://github.com/wwgberlin/tutorials)

[Effective Go](https://golang.org/doc/effective_go.html)

[Go by example](https://gobyexample.com/)

[Awesome Go](https://awesome-go.com/)

[Gophercises](https://gophercises.com/)

[Go training - Ardan Labs](https://github.com/ardanlabs/gotraining/tree/master/topics)
