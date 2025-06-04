 🧠 Go Language Concepts (Intermediate)

---

## 🔁 **Closures**

**Definition**
A closure is a function value that captures variables from its surrounding lexical scope, even after that scope has exited.

**Practical Use Cases**

* Stateful functions (e.g., counters)
* Encapsulation (hide implementation details)
* Callbacks (e.g., async handlers, event listeners)

**Usefulness**

* Encapsulation of logic and data
* Flexibility in design
* Enhances code readability

**Considerations**

* Memory usage (captured variables may prevent garbage collection)
* Concurrency issues (shared mutable state)

**Best Practices**

* Limit the scope of closures
* Avoid excessive use for clarity and maintainability

---

## 🔃 **Recursion**

**Practical Use Cases**

* Mathematical algorithms (factorial, Fibonacci)
* Tree/graph traversal (DFS, BFS)
* Divide and conquer (merge sort, quicksort)

**Benefits**

* Simplicity in logic
* Clarity of code
* Flexibility to model self-repeating problems

**Considerations**

* Performance (stack usage, depth limit)
* Base case must be well-defined to prevent infinite recursion

**Best Practices**

* Always test base and edge cases
* Consider optimization (e.g., memoization, tail recursion)
* Clearly define recursive and termination conditions

---

## 📍 **Pointers**

**Definition**
A pointer stores the memory address of another variable.

**Use Cases**

* Modify variables indirectly (pass-by-reference)
* Efficiently pass large structs or arrays
* Direct memory management for performance-critical applications

**Declaration & Initialization**

```go
var ptr *int       // declaration
var a int = 10
ptr = &a           // initialization
```

---

## 🧵 **Strings**

**Definition**
A string is a sequence of bytes in Go; strings are immutable and contain UTF-8 encoded characters (runes).

**String Types**

* **Raw strings**: use backticks (`` ` ``)
* **Interpreted strings**: use double quotes with escape sequences

**Escape Sequences**

* `\n`: newline
* `\t`: tab
* `\r`: carriage return

**Characteristics**

* Strings are arrays of runes (Unicode code points)
* Accessing characters returns the rune (ASCII if in that range)

**Operations & Methods**

* `len(s)`: returns byte length
* `utf8.RuneCountInString(s)`: counts actual characters

---

## 🖨️ **fmt Package**

**Features**

* **Printing functions**: `Print`, `Printf`, `Println`
* **Formatting functions**: `Sprintf`, `Errorf`
* **Input functions**: `Scan`, `Scanf`, `Scanln`
* **Error formatting**: embed context into errors

---

## 🧱 **Structs**

**Definition**
A struct is a composite type that groups related variables (fields) under one name.

**Features**

* Supports **anonymous structs** for inline definitions
* Can define **methods** on struct types

---

## 🔧 **Methods**

Methods are functions with a receiver (value or pointer) and are bound to types (usually structs).

---

## 🎭 **Interfaces**

**Definition**
Interfaces define method sets and represent behavior. Similar to abstract classes in other languages but without inheritance.

**Key Concepts**

* No need for `implements` keyword — a type satisfies an interface by implementing its methods
* Enables **polymorphism** and **decoupled design**

**Exporting Rule in Go**

* Exported names must **start with an uppercase letter**

---

## 🧬 **Struct Embedding**

**Definition**
Embedding allows one struct to include another, achieving a form of composition (similar to inheritance).

**Best Practices**

* Prefer **composition over inheritance**
* Avoid the **diamond problem** (common in multiple inheritance)
* Prioritize clarity and simplicity
* Explicitly initialize embedded structs when needed

---

## 🧰 **Generics**

**Definition**
Generics allow writing functions, types, and data structures that work with any type — without sacrificing type safety.

**Benefits**

* Reusable and DRY (Don't Repeat Yourself) code
* Compile-time type safety
* No performance penalty compared to interface{} + type assertions

**Considerations**

* Use **type constraints** (`~int`, `constraints.Ordered`, etc.)
* Ensure proper documentation for generic types
* Test with multiple types to avoid hidden bugs

**Example**

```go
type Number interface {
    ~int | ~float64
}

func Add[T Number](a, b T) T {
    return a + b
}
```
## 🧰 **Errors**

### 🔍 **Definition**

In Go, errors are represented by the built-in `error` interface. It is the standard way to signal that an operation did not complete successfully.

```go
type error interface {
    Error() string
}
```

### 🛠️ **Implementation**

* You can **create errors** using the `errors` package:

  ```go
  import "errors"

  err := errors.New("something went wrong")
  ```

* Or using `fmt.Errorf` for formatted error messages:

  ```go
  err := fmt.Errorf("failed to load user: %v", userID)
  ```

* You can also **implement the `error` interface** with custom types:

  ```go
  type MyError struct {
      Message string
  }

  func (e *MyError) Error() string {
      return e.Message
  }
  ```

### 🚨 **Best Practices**

* Use `error` returns for **expected or recoverable errors**.
* Use `panic` only for **unrecoverable conditions** (e.g., corrupted program state).
* Always **check and handle errors** returned by functions.
* Prefer wrapping errors with context to make debugging easier.


## 📍 **Custom errors**

**Definition**
In go you can create custom errors.

### 🛠️ **Implementation**
You create a custom error by creating a struct which would have the message of the error you want to pass and then create an error function that returns the an error type with the neccessary struct.

# strings functions


**Definition**
Strings in go are a sequence of bytes.In go you can create 

**Implementation**

str  := "Hello Go!"

# text templates
## actions
- variable insertion
- pipelines
- control structures
- iteration
## advanced features
- nested templates
- functions
- custom delimiters
- error handling
## use cases
- Html generation
- email templates
- code generation
- document generation
## best practices
- Separation of concern
- efficieny
- security

# Regex
escape sequences

# Random
seed:starting point for generating a sequence of random numbers
```
rand.Intn(n)
rand.Float64()
```
considerations:

deterministic nature
thread safety
cyrptographic security
need to use mutex
to generate randoms for cryptographic use make use of the crypto packagae instead of math

# Number Parsing
make use of this package strconv