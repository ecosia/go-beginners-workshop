<img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GOPHER_AVATARS.jpg" width="950" title="Avatar Gophers">

# Go beginners workshop

## Contents
- [Go beginners workshop](#go-beginners-workshop)
  - [Contents](#contents)
  - [Setting up your environment](#setting-up-your-environment)
    - [Download go](#download-go)
    - [Workspace suggestions](#workspace-suggestions)
  - [Language mechanics](#language-mechanics)
    - [Variables](#variables)
    - [Functions](#functions)
    - [Scopes](#scopes)
    - [Types](#types)
      - [Basic types](#basic-types)
      - [Pointers](#pointers)
      - [Structs](#structs)
    - [Loops](#loops)
    - [Data structures](#data-structures)
      - [Arrays](#arrays)
      - [Slices](#slices)
    - [Concurrency](#concurrency)
      - [Goroutines](#goroutines)
      - [Channels](#channels)
  - [References and resources](#references-and-resources)
    
<img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GopherSpaceMentor.png" width="250" title="Space mentor">

## Setting up your environment

### Download go
To get started, go to the [Go download page](https://golang.org/dl/) and download the binary distribution according to your operation system.

 #### Mac OS
Download the OS X installer from https://golang.org/dl/. Double tap to start the installation. Follow the prompts and this should have Golang installed in /usr/local/go and would have also added the folder /usr/local/go/bin to your PATH environment variable.

 #### Windows
Download the MSI installer from https://golang.org/dl/. Double tap to start the installation and follow the prompts. This will install Golang in location c:\Go and will also add the directory c:\Go\bin to your path environment variable.

 #### Linux
Download the tar file from https://golang.org/dl/ and unzip it to /usr/local.

 Add /usr/local/go/bin to the PATH environment variable. This should install go in linux.

Please note the above text is taken from this [tutorial](https://golangbot.com/golang-tutorial-part-1-introduction-and-installation/)

### Workspace suggestions
When writing code you can use any text editing application, such as TextEdit or Notepad. However there are IDE (Integrated Development Environment) applications designed specifically for development, offering helpful tools such as syntax highlighting, auto correction, and terminals where you can run the code. To understand more about the difference between text editors and IDEs we recommend this [article](https://www.thoughtco.com/beginners-guide-to-using-an-ide-versus-a-text-editor-2034114).

Almost every developer you ask has an IDE or text editor they prefer to use. When you are starting out it can be difficult to choose which one works best for you. Below are a couple we recommend. Visual Studio Code, although a Microsoft product, is open source and it has a range of plugins. It's used by the majority of the developers at Ecosia.

Goland is an IDE developed specifically for working with GO and has a number of language specific tools. Atom is another highly used and praised IDE. Whichever one you choose to work with, it will take a while to get used to all shortcuts and plugins but it is well worth taking some time to get to know it and set it up the way that works most efficiently for you.

[Visual Studio Code](https://code.visualstudio.com/)

[Goland](https://www.jetbrains.com/go/)

[Atom](https://atom.io/)

<img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GO_BUILD.png" width="250" title="Go Build">

## Language mechanics

### Variables
When building your application, you will need at some point to store data in memory, at a specific location. To do that, you will need a variable - that is a memory location where a value of a specific type is stored. A variable can receive a value with the condition that the value has its same type.

```go
// Declare a variable of type integer
var year int
// Variable assignment
year = 1990
// New variable assignment of a value with the same integer type
year = 2000

// Short variable declaration: you can use it when inside a function.
currentYear := 2019
```

[Playground](https://play.golang.org/p/iLIXqjD8ayW)

### Functions
A function is a group of statements that together perform a task. Every Go program has at least one function, which is main(). You can divide your code into separate functions. How you divide your code among different functions is up to you, but logically, the division should be such that each function performs a specific task. [Tutorial point](https://www.tutorialspoint.com/go/go_functions.htm)

[Playground](https://play.golang.org/p/AyD4N4zgdJx)

### Scopes
The places in your Go program where you are allowed to use a variable is called scope. In Go, the scope of the variable is connected to the closest pair of curly braces `{` and `}`. There are different types of scopes in a Go program, as mentioned in the [official docs](https://golang.org/ref/spec#Declarations_and_scope).

[Playground](https://play.golang.org/p/0bcUSGtKeXo)

### Types
We will introduce here some of the most used types in Go, necessary to have a basic understand of the language.

#### Basic types
As described in [Tour of Go](https://tour.golang.org/basics/11), those are Go's basic types:

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

[Playground - strings](https://play.golang.org/p/nMrOiDk8uXX) |
[Playground - integers, floats](https://play.golang.org/p/xAN0-OV78lB)
[Playground - booleans](https://play.golang.org/p/DigM6AI7ogJ)

#### Pointers
[Pointers](https://en.wikipedia.org/wiki/Pointer_(computer_programming)) in Go are also a type and they store a memory address. In Go, everything is passed by value: when passing data to a function, the function gets an entire copy of what's being passed. If you pass a `string` value, Go makes a copy of the `string` value. When passing a pointer value, Go makes a copy of the pointer. The correct usage of pointers is very important in Go: pointers are a way to efficiently share data, without having to make a copy of the entire value. The official [Go documentation](https://golang.org/doc/faq#Pointers) is a great resource to understand the usage of pointers and its advantages.

To create a pointer to a variable `x`, we use the `&` syntax:

```go
// Declares a variable x
x := 42
// Creates a variable y of type pointer that creates a pointer to x
y := &x
```

To read the value of the new created variable `y`, we use the `*` syntax. This is also called `dereference`:

```go
// Gets value of y variable
value := *y
// Changes the value of the y variable
value = 22
// Updates the y variable
y = &value
// Changes the value of y directly, without the need of assigning to a new variable
*y = 23
```

[Playground](https://play.golang.org/p/3eRREFtFfJB)


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

[Playground](https://play.golang.org/p/M3ed_-cUdWH)

### Loops
Loops are used to execute blocks of code repeatedly given a certain condition. Unlike other languages, Go has only a `for` loop available:

```go
for i := 1; i <= 10; i++ {
    fmt.Printf("%d ",i)
}
```

You can use a `break` statement to terminate the `for` loop. Go will execute the next line of code outside the `for` loop.

[Playground](https://play.golang.org/p/8nCbcAoEnXe)

### Data structures

#### Arrays
In Go, an array is a collection of elements that have the same type. This collection has a fixed size: once you declare an array, its size cannot increase or decrease.

```go
numbers := [5]int{5, 4, 3, 2, 1}
```
Above we have declared an array with 5 elements, in which each element is an integer.

[Playground](https://play.golang.org/p/4WalWG0buHs)

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

[Playground](https://play.golang.org/p/aPmGlu2SIzk)

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


<img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GopherInclusion.png" width="250" title="Community Gopher">

## References and resources

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

[Go 101](go101.org)

<img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GoCommunity.png" width="950" title="Community Gopher">

