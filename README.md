# Go Programming Language - Complete Reference Guide

![Go Logo](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)

## Table of Contents

1. [Introduction to Go](#introduction-to-go)
2. [Installation and Setup](#installation-and-setup)
3. [Variables and Data Types](#variables-and-data-types)
4. [Operators](#operators)
5. [Packages and Imports](#packages-and-imports)
6. [Arrays and Slices](#arrays-and-slices)
7. [Control Structures](#control-structures)
8. [Functions](#functions)
9. [Structs](#structs)
10. [Pointers](#pointers)
11. [Methods and Interfaces](#methods-and-interfaces)
12. [Advanced Data Structures](#advanced-data-structures)
13. [Concurrency](#concurrency)
14. [Error Handling](#error-handling)
15. [File I/O](#file-io)
16. [JSON and XML Processing](#json-and-xml-processing)
17. [HTTP and Web Development](#http-and-web-development)
18. [Testing and Benchmarking](#testing-and-benchmarking)
19. [Best Practices](#best-practices)

---

## Introduction to Go

Go (also called Golang) is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. First released in 2009, Go is known for its:

- **Simplicity**: Clean, readable syntax
- **Performance**: Compiled to machine code
- **Concurrency**: Built-in goroutines and channels
- **Fast Compilation**: Quick build times
- **Static Typing**: Type safety with inference
- **Garbage Collection**: Automatic memory management
- **Standard Library**: Rich set of built-in packages

### Key Features

- Simple and easy to learn
- Fast compilation and execution
- Built-in concurrency support
- Strong standard library
- Cross-platform support
- Excellent tooling (`go fmt`, `go test`, `go mod`)

---

## Installation and Setup

### Installing Go

**Windows:**

1. Download from https://go.dev/dl/
2. Run the MSI installer
3. Verify: `go version`

**macOS:**

```bash
brew install go
```

**Linux:**

```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### Setting up GOPATH

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### First Go Program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Run:**

```bash
go run main.go
```

**Build:**

```bash
go build main.go
./main
```

---

## Variables and Data Types

### Variable Declaration

In Go, once a variable's data type is assigned, it **cannot be changed** (statically typed).

#### Three Ways to Declare Variables:

```go
// Method 1: Full declaration with type
var name string = "test"

// Method 2: Type inference with var
var name2 = "test"

// Method 3: Short declaration (only inside functions)
name := "test"
```

**Important Notes:**

- Short declaration (`:=`) can only be used **inside functions**, not at the package level
- Once declared, you cannot change the type

### String Type

```go
var name string = "test"
var name2 = "test"
name := "test" // shorthand, function-scope only

// String operations
var greeting = "Hello, " + "World!" // concatenation
length := len("Hello") // 5

// Multi-line strings
var multiline = `This is a
multi-line
string`
```

### Integer Types

```go
var age int = 10
var age = 10
age := 10 // will be assigned default int type
```

#### Integer Type Specifications:

| Type           | Bits | Signed Range                                            | Unsigned Range                  |
| -------------- | ---- | ------------------------------------------------------- | ------------------------------- |
| `int8`         | 8    | -128 to 127                                             | -                               |
| `uint8` (byte) | 8    | -                                                       | 0 to 255                        |
| `int16`        | 16   | -32,768 to 32,767                                       | -                               |
| `uint16`       | 16   | -                                                       | 0 to 65,535                     |
| `int32` (rune) | 32   | -2,147,483,648 to 2,147,483,647                         | -                               |
| `uint32`       | 32   | -                                                       | 0 to 4,294,967,295              |
| `int64`        | 64   | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 | -                               |
| `uint64`       | 64   | -                                                       | 0 to 18,446,744,073,709,551,615 |

### Unsigned Integer Types (Uint)

Unsigned integers **cannot hold negative values**.

```go
var count uint = 100
var smallNum uint8 = 255 // range: 0 to 255
```

### Float, Boolean, Complex Types

```go
// Float
var price float32 = 19.99
var precise float64 = 3.14159265359

// Boolean
var isActive bool = true
var flag = true

// Complex numbers
var c complex64 = 1 + 2i
var c2 complex128 = complex(3, 4)
```

### Constants

```go
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
)

// iota - constant generator
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
)
```

### Type Conversion

```go
var i int = 42
var f float64 = float64(i)

// String conversion
import "strconv"
str := strconv.Itoa(42)
num, err := strconv.Atoi("42")
```

---

## Operators

### Arithmetic Operators

```go
a + b    // Addition
a - b    // Subtraction
a * b    // Multiplication
a / b    // Division
a % b    // Modulus
a++      // Increment
b--      // Decrement
```

### Comparison Operators

```go
a == b   // Equal
a != b   // Not equal
a < b    // Less than
a > b    // Greater than
a <= b   // Less than or equal
a >= b   // Greater than or equal
```

### Logical Operators

```go
a && b   // AND
a || b   // OR
!a       // NOT
```

### Bitwise Operators

```go
a & b    // AND
a | b    // OR
a ^ b    // XOR
a &^ b   // AND NOT
a << 2   // Left shift
a >> 2   // Right shift
```

---

## Packages and Imports

### The `fmt` Package

```go
import "fmt"

fmt.Print("text")           // prints without newline
fmt.Println("text")         // prints with newline
fmt.Printf("formatted %v", value) // prints formatted string

// Sprintf stores the formatted string
str := fmt.Sprintf("Value: %v", 42)
```

#### Format Specifiers:

| Specifier | Purpose                 | Example               |
| --------- | ----------------------- | --------------------- |
| `%v`      | Default format          | `{Alice 30}`          |
| `%+v`     | Struct with field names | `{Name:Alice Age:30}` |
| `%T`      | Type of value           | `int`                 |
| `%d`      | Decimal integer         | `42`                  |
| `%f`      | Float                   | `3.140000`            |
| `%.2f`    | Float (2 decimals)      | `3.14`                |
| `%s`      | String                  | `hello`               |
| `%q`      | Quoted string           | `"hello"`             |
| `%p`      | Pointer address         | `0xc000010200`        |

### Other Common Packages

```go
import (
    "strings"  // string manipulation
    "sort"     // sorting algorithms
    "math"     // mathematical functions
    "time"     // time and date
    "os"       // OS functionality
)

// strings package
upper := strings.ToUpper("hello")
contains := strings.Contains("hello", "ll")

// sort package
nums := []int{3, 1, 4}
sort.Ints(nums)
```

---

## Arrays and Slices

### Arrays (Fixed Size)

```go
// Array declaration
var arr [5]int = [5]int{1, 2, 3, 4, 5}
arr := [5]int{1, 2, 3, 4, 5}
arr := [...]int{1, 2, 3, 4, 5} // compiler counts

// Access and modify
second := arr[1]
arr[0] = 10

// Iterate
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

for index, value := range arr {
    fmt.Printf("%d: %d\n", index, value)
}
```

### Slices (Dynamic Arrays)

```go
// Slice declaration
slice := []int{1, 2, 3, 4, 5}

// Make function
slice := make([]int, 5)       // length 5, capacity 5
slice := make([]int, 5, 10)   // length 5, capacity 10

// Append elements
slice = append(slice, 6)
slice = append(slice, 7, 8, 9)

// Slicing operations
sub := slice[1:4]   // index 1, 2, 3 (4 excluded)
sub2 := slice[2:]   // from index 2 to end
sub3 := slice[:3]   // from start to index 2
sub4 := slice[:]    // entire slice

// Copy slices
dst := make([]int, len(src))
copy(dst, src)

// Length and capacity
length := len(slice)
capacity := cap(slice)
```

---

## Control Structures

### If-Else

```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x > 5 {
    fmt.Println("x is greater than 5")
} else {
    fmt.Println("x is 5 or less")
}

// If with short statement
if num := 9; num < 0 {
    fmt.Println("Negative")
} else {
    fmt.Println("Positive")
}
```

### Switch

```go
// Basic switch
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("End of week")
default:
    fmt.Println("Middle of week")
}

// Multiple conditions
switch grade {
case "A", "B":
    fmt.Println("Excellent")
default:
    fmt.Println("Needs improvement")
}

// Switch without expression
switch {
case age < 18:
    fmt.Println("Minor")
case age < 65:
    fmt.Println("Adult")
default:
    fmt.Println("Senior")
}
```

### Loops

```go
// Traditional for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-style loop
x := 0
for x < 5 {
    fmt.Println(x)
    x++
}

// Infinite loop
for {
    // runs forever, use break to exit
}

// Range loop
names := []string{"Alice", "Bob", "Charlie"}

for index, value := range names {
    fmt.Printf("Index %d: %s\n", index, value)
}

// Only value
for _, value := range names {
    fmt.Println(value)
}

// Break and Continue
for i := 0; i < 10; i++ {
    if i == 5 {
        break    // exit loop
    }
    if i%2 == 0 {
        continue // skip to next iteration
    }
    fmt.Println(i)
}
```

---

## Functions

### Basic Functions

```go
// Function with return value
func add(x, y int) int {
    return x + y
}

// Multiple parameters (same type)
func add(x, y int) int {
    return x + y
}

// Multiple return values
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return x / y, nil
}

// Named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // naked return
}
```

### Variadic Functions

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

result := sum(1, 2, 3, 4, 5)
```

### Anonymous Functions & Closures

```go
// Anonymous function
func() {
    fmt.Println("Hello")
}()

// Assigned to variable
add := func(x, y int) int {
    return x + y
}

// Closure
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

### Defer

```go
func readFile() {
    file, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer file.Close() // executed when function returns

    // work with file
}

// Multiple defers (LIFO - Last In First Out)
func example() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
}
// Output: Third, Second, First
```

---

## Structs

### Basic Struct

```go
// Define struct
type Person struct {
    Name  string
    Age   int
    Email string
}

// Create instances
var p1 Person
p1.Name = "Alice"

// Struct literal
p2 := Person{
    Name:  "Bob",
    Age:   25,
    Email: "bob@example.com",
}

// Pointer to struct
p3 := &Person{Name: "Charlie", Age: 30}
```

### Nested Structs

```go
type Address struct {
    Street  string
    City    string
    ZipCode string
}

type Employee struct {
    Name    string
    Age     int
    Address Address
}

emp := Employee{
    Name: "John",
    Address: Address{
        Street: "123 Main St",
        City:   "New York",
    },
}
```

### Embedded Structs

```go
type Contact struct {
    Email string
    Phone string
}

type User struct {
    Name    string
    Contact // embedded
}

user := User{
    Name: "Alice",
    Contact: Contact{Email: "alice@example.com"},
}

// Access embedded fields directly
fmt.Println(user.Email)
```

### Struct Tags

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age,omitempty"`
}
```

---

## Pointers

### Pointer Basics

```go
// Declare pointer
var ptr *int

// Get address (&)
x := 10
ptr = &x

// Dereference (*)
value := *ptr // value = 10

// Modify through pointer
*ptr = 20
fmt.Println(x) // 20

// Nil pointer
var p *int
fmt.Println(p == nil) // true
```

### Pointers with Structs

```go
type Person struct {
    Name string
    Age  int
}

// Pointer to struct
p := &Person{Name: "Alice", Age: 30}

// Access fields (auto-dereference)
fmt.Println(p.Name)
p.Age = 31
```

### Pointers in Functions

```go
// Pass by value (copy)
func incrementValue(x int) {
    x++
}

// Pass by pointer (reference)
func incrementPointer(x *int) {
    *x++
}

num := 10
incrementValue(num)
fmt.Println(num) // 10 (unchanged)

incrementPointer(&num)
fmt.Println(num) // 11 (changed)
```

### New Function

```go
ptr := new(int)
*ptr = 100
```

---

## Methods and Interfaces

### Methods

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()
rect.Scale(2)
```

### Interfaces

```go
// Define interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Using interface
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

rect := Rectangle{10, 5}
printShapeInfo(rect)
```

### Empty Interface

```go
var anything interface{}

anything = 42
anything = "hello"
anything = []int{1, 2, 3}

// Type assertion
str, ok := anything.(string)
if ok {
    fmt.Println("It's a string:", str)
}

// Type switch
switch v := anything.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

---

## Advanced Data Structures

### Maps (Hash Tables)

```go
// Create map
ages := make(map[string]int)

// Map literal
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
}

// Set value
ages["Alice"] = 25

// Get value
age := ages["Alice"]

// Check if key exists
age, exists := ages["Charlie"]
if exists {
    fmt.Println("Found:", age)
}

// Delete key
delete(ages, "Alice")

// Iterate
for key, value := range scores {
    fmt.Printf("%s: %d\n", key, value)
}

// Length
length := len(scores)
```

### Sets (Using Maps)

```go
// Set using map[T]struct{}
set := make(map[string]struct{})

// Add elements
set["apple"] = struct{}{}
set["banana"] = struct{}{}

// Check membership
if _, exists := set["apple"]; exists {
    fmt.Println("apple is in set")
}

// Remove
delete(set, "banana")

// Iterate
for item := range set {
    fmt.Println(item)
}
```

### Linked Lists

```go
import "container/list"

// Create list
l := list.New()

// Add elements
l.PushBack(1)
l.PushBack(2)
l.PushFront(0)

// Iterate
for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}

// Remove
e := l.Front()
l.Remove(e)
```

### Stack Implementation

```go
type Stack []interface{}

func (s *Stack) Push(v interface{}) {
    *s = append(*s, v)
}

func (s *Stack) Pop() (interface{}, bool) {
    if len(*s) == 0 {
        return nil, false
    }
    index := len(*s) - 1
    value := (*s)[index]
    *s = (*s)[:index]
    return value, true
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}
```

### Queue Implementation

```go
type Queue []interface{}

func (q *Queue) Enqueue(v interface{}) {
    *q = append(*q, v)
}

func (q *Queue) Dequeue() (interface{}, bool) {
    if len(*q) == 0 {
        return nil, false
    }
    value := (*q)[0]
    *q = (*q)[1:]
    return value, true
}
```

### Priority Queue (Heap)

```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Usage
h := &IntHeap{2, 1, 5}
heap.Init(h)
heap.Push(h, 3)
min := heap.Pop(h) // 1
```

### Binary Tree

```go
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

func (t *TreeNode) Insert(value int) {
    if value < t.Value {
        if t.Left == nil {
            t.Left = &TreeNode{Value: value}
        } else {
            t.Left.Insert(value)
        }
    } else {
        if t.Right == nil {
            t.Right = &TreeNode{Value: value}
        } else {
            t.Right.Insert(value)
        }
    }
}

func (t *TreeNode) Search(value int) bool {
    if t == nil {
        return false
    }
    if t.Value == value {
        return true
    }
    if value < t.Value {
        return t.Left.Search(value)
    }
    return t.Right.Search(value)
}

// Traversals
func (t *TreeNode) InOrder() {
    if t != nil {
        t.Left.InOrder()
        fmt.Printf("%d ", t.Value)
        t.Right.InOrder()
    }
}
```

### Graph

```go
type Graph struct {
    vertices map[int][]int
}

func NewGraph() *Graph {
    return &Graph{vertices: make(map[int][]int)}
}

func (g *Graph) AddEdge(v1, v2 int) {
    g.vertices[v1] = append(g.vertices[v1], v2)
}

// BFS
func (g *Graph) BFS(start int) {
    visited := make(map[int]bool)
    queue := []int{start}
    visited[start] = true

    for len(queue) > 0 {
        vertex := queue[0]
        queue = queue[1:]
        fmt.Printf("%d ", vertex)

        for _, neighbor := range g.vertices[vertex] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}

// DFS
func (g *Graph) DFS(start int, visited map[int]bool) {
    visited[start] = true
    fmt.Printf("%d ", start)

    for _, neighbor := range g.vertices[start] {
        if !visited[neighbor] {
            g.DFS(neighbor, visited)
        }
    }
}
```

---

## Concurrency

### Goroutines

```go
// Basic goroutine
go func() {
    fmt.Println("Running in goroutine")
}()

// With parameters
go printMessage("Hello")

// Anonymous with parameters
message := "Hello"
go func(msg string) {
    fmt.Println(msg)
}(message)
```

### WaitGroup

```go
import "sync"

var wg sync.WaitGroup

wg.Add(3)

go func() {
    defer wg.Done()
    fmt.Println("Goroutine 1")
}()

go func() {
    defer wg.Done()
    fmt.Println("Goroutine 2")
}()

go func() {
    defer wg.Done()
    fmt.Println("Goroutine 3")
}()

wg.Wait() // Wait for all goroutines
```

### Channels

```go
// Create channel
ch := make(chan int)

// Send to channel
go func() {
    ch <- 42
}()

// Receive from channel
value := <-ch

// Buffered channel
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3

// Close channel
close(ch)

// Range over channel
ch := make(chan int)
go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}()

for value := range ch {
    fmt.Println(value)
}
```

### Select Statement

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    ch1 <- "from channel 1"
}()

go func() {
    ch2 <- "from channel 2"
}()

select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
case <-time.After(time.Second):
    fmt.Println("timeout")
}
```

### Mutex

```go
import "sync"

type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

### Worker Pool Pattern

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```

### Context

```go
import "context"

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-time.After(3 * time.Second):
    fmt.Println("completed")
case <-ctx.Done():
    fmt.Println("timeout")
}

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
go func() {
    time.Sleep(1 * time.Second)
    cancel()
}()

// With values
ctx := context.WithValue(context.Background(), "userID", 12345)
userID := ctx.Value("userID").(int)
```

---

## Error Handling

### Basic Errors

```go
import "errors"

// Create error
err := errors.New("something went wrong")

// Return error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Handle error
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}

// Formatted error
err := fmt.Errorf("invalid id: %d", id)
```

### Custom Errors

```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething() error {
    return &MyError{Code: 404, Message: "not found"}
}
```

### Error Wrapping (Go 1.13+)

```go
import "errors"

// Wrap error
func readConfig() error {
    err := readFile("config.txt")
    if err != nil {
        return fmt.Errorf("read config: %w", err)
    }
    return nil
}

// Unwrap
unwrapped := errors.Unwrap(err)

// Check error type
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("Path error")
}

// Check error value
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File not found")
}
```

### Panic and Recover

```go
// Panic
func mayCrash(x int) {
    if x == 0 {
        panic("x cannot be zero")
    }
}

// Recover
func safeExecute() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    mayCrash(0)
}
```

---

## File I/O

### Reading Files

```go
import (
    "io/ioutil"
    "os"
    "bufio"
)

// Read entire file
data, err := ioutil.ReadFile("file.txt")
if err != nil {
    panic(err)
}
fmt.Println(string(data))

// Read line by line
file, err := os.Open("file.txt")
if err != nil {
    panic(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

### Writing Files

```go
// Write entire file
data := []byte("Hello, World!")
err := ioutil.WriteFile("output.txt", data, 0644)

// Write with os
file, err := os.Create("output.txt")
if err != nil {
    panic(err)
}
defer file.Close()

file.WriteString("Hello, World!\n")

// Buffered writer
writer := bufio.NewWriter(file)
writer.WriteString("Buffered write\n")
writer.Flush()
```

### File Operations

```go
// Check if exists
if _, err := os.Stat("file.txt"); err == nil {
    fmt.Println("File exists")
}

// Get file info
info, err := os.Stat("file.txt")
fmt.Println("Name:", info.Name())
fmt.Println("Size:", info.Size())
fmt.Println("Modified:", info.ModTime())

// Delete file
os.Remove("file.txt")

// Rename file
os.Rename("old.txt", "new.txt")
```

### Directory Operations

```go
// Create directory
os.Mkdir("mydir", 0755)

// Create directory tree
os.MkdirAll("path/to/mydir", 0755)

// Remove directory
os.Remove("mydir")
os.RemoveAll("path/to/mydir")

// Read directory
entries, err := ioutil.ReadDir(".")
for _, entry := range entries {
    fmt.Printf("%s (%v)\n", entry.Name(), entry.IsDir())
}

// Walk directory tree
import "path/filepath"

filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    fmt.Printf("%s\n", path)
    return nil
})
```

---

## JSON and XML Processing

### JSON Encoding/Decoding

```go
import "encoding/json"

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

// Marshal (struct to JSON)
person := Person{Name: "Alice", Age: 30}
jsonData, err := json.Marshal(person)

// MarshalIndent (pretty print)
jsonData, err := json.MarshalIndent(person, "", "  ")

// Unmarshal (JSON to struct)
jsonStr := `{"name":"Bob","age":25}`
var decoded Person
json.Unmarshal([]byte(jsonStr), &decoded)

// Working with maps
data := map[string]interface{}{
    "name": "Alice",
    "age":  30,
}
jsonData, _ := json.Marshal(data)
```

### XML Encoding/Decoding

```go
import "encoding/xml"

type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
}

// Marshal
person := Person{Name: "Alice", Age: 30}
xmlData, err := xml.Marshal(person)

// MarshalIndent
xmlData, err := xml.MarshalIndent(person, "", "  ")

// Unmarshal
var decoded Person
xml.Unmarshal(xmlData, &decoded)
```

---

## HTTP and Web Development

### HTTP Client

```go
import "net/http"

// GET request
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    panic(err)
}
defer resp.Body.Close()

body, _ := ioutil.ReadAll(resp.Body)

// POST request
data := []byte(`{"name":"Alice"}`)
resp, err := http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(data),
)

// Custom request
req, _ := http.NewRequest("PUT", "https://api.example.com/users/1", nil)
req.Header.Set("Authorization", "Bearer token")
client := &http.Client{}
resp, err := client.Do(req)
```

### HTTP Server

```go
import "net/http"

// Simple handler
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

// JSON response
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    data := map[string]string{"message": "success"}
    json.NewEncoder(w).Encode(data)
}

// Different methods
func apiHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        fmt.Fprintf(w, "GET request")
    case "POST":
        body, _ := ioutil.ReadAll(r.Body)
        fmt.Fprintf(w, "POST: %s", body)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

// Middleware
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("%s %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```

---

## Testing and Benchmarking

### Unit Testing

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

// Table-driven tests
func TestAddTable(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

### Benchmarking

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

// Run: go test -bench=.
// Run with memory: go test -bench=. -benchmem
```

### Test Coverage

```bash
# Run tests with coverage
go test -cover

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## Best Practices

### Code Organization

```
myproject/
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── config/
│   ├── database/
│   └── handlers/
├── pkg/
│   └── utils/
├── go.mod
└── README.md
```

### Error Handling

```go
// Good: Handle errors immediately
f, err := os.Open("file.txt")
if err != nil {
    return fmt.Errorf("failed to open: %w", err)
}
defer f.Close()

// Bad: Ignoring errors
f, _ := os.Open("file.txt")
```

### Naming Conventions

```go
// Variables: camelCase (private), PascalCase (exported)
var userName string      // private
var UserName string      // exported

// Functions: PascalCase (exported), camelCase (private)
func GetUser() {}        // exported
func getUserByID() {}    // private

// Interfaces: -er suffix
type Reader interface {}
type Writer interface {}

// Acronyms: Keep case
var userID int           // not userId
var httpServer Server    // not httpServer
```

### Comments

```go
// Package documentation
// Package math provides basic math operations.
package math

// Function documentation
// Add returns the sum of a and b.
func Add(a, b int) int {
    return a + b
}
```

### Dependency Management

```bash
# Initialize module
go mod init github.com/username/project

# Add dependency
go get github.com/pkg/errors

# Update dependencies
go get -u ./...

# Tidy up
go mod tidy
```

### Performance Tips

```go
// Pre-allocate slices
slice := make([]int, 0, 100)

// Use strings.Builder
var builder strings.Builder
for i := 0; i < 100; i++ {
    builder.WriteString("text")
}

// Avoid allocations in loops
var buf bytes.Buffer
for i := 0; i < n; i++ {
    buf.Reset()
    fmt.Fprintf(&buf, "item %d", i)
}
```

---

## Useful Commands

```bash
# Build and run
go run main.go
go build
go build -o myapp

# Cross compilation
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build

# Testing
go test
go test -v
go test -cover
go test ./...

# Benchmarking
go test -bench=.
go test -bench=. -benchmem

# Module management
go mod init
go mod tidy
go mod vendor

# Formatting
go fmt ./...
go vet ./...

# Documentation
go doc fmt.Println
godoc -http=:6060

# Race detection
go run -race main.go
go test -race
```

---

## Conclusion

This comprehensive guide covers:

✅ **Basics**: Variables, types, operators  
✅ **Control Flow**: Loops, conditionals, switches  
✅ **Functions**: Declaration, closures, defer  
✅ **Data Structures**: Arrays, slices, maps, structs  
✅ **Advanced Topics**: Pointers, interfaces, methods  
✅ **Data Structures**: Sets, stacks, queues, trees, graphs  
✅ **Concurrency**: Goroutines, channels, mutexes  
✅ **Error Handling**: Custom errors, panic/recover  
✅ **File I/O**: Reading, writing, directories  
✅ **Web Development**: HTTP client/server  
✅ **Testing**: Unit tests, benchmarks, coverage  
✅ **Best Practices**: Code organization, patterns

### Learning Resources

- **Official Documentation**: https://go.dev/doc/
- **Go by Example**: https://gobyexample.com/
- **Effective Go**: https://go.dev/doc/effective_go
- **Go Tour**: https://go.dev/tour/
- **Go Playground**: https://go.dev/play/

### Next Steps

1. Practice with small projects
2. Read the standard library source code
3. Contribute to open-source Go projects
4. Build production applications
5. Explore advanced topics (reflection, unsafe, cgo)

---

**Happy Coding with Go! 🚀**
