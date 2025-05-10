# 🐹 Go Language (Basics)

---

## 📌 About Go

* **Statically typed, compiled** language by Google (Robert Griesemer, 2009)
* Similar syntax to C, but safer and more modern
* Features: memory safety, garbage collection (GC), concurrency support
* Competitor to Rust

---

## ❓ Why Go

* **Performance**: Compiles to machine code
* **Concurrency**: Goroutines make concurrent programming easy and lightweight
* **Readability**: Simple and clear syntax
* **Maintainability**: Easy to manage large codebases
* **Security**: Strong typing, memory safety
* **Versatility**: Suitable for backend, cloud services, DevOps, etc.

---

## 💬 Community & Tools

* **Discord**: [Go Discord](https://discord.com/invite/aVJPs76B6M)

* **Playground**: Search "Go Playground" on Google

* **Extensions** (VS Code):

  * Go
  * Code Runner
  * Error Lens
  * TODO Tree

* **Version Control**: Git

---

## 🧱 Project Structure & Compilation

### Basic Commands

* Run: `go run filename.go`
* Build: `go build filename.go` → Produces executable binary

### Compiler & Runtime

* Go compiler handles memory allocation, GC, and concurrency support
* Go includes a lightweight runtime unlike C/C++

### Entry Point

* `package main` → Required for executable programs
* `func main()` → Program entry point
* `fmt` package used for I/O

---

## 📚 Standard Library & Imports

* Use: `import "fmt"`
* Tree shaking: Unused imports are removed at compile time
* Multiple imports:

  ```go
  import (
      "fmt"
      "net/http"
  )
  ```
* Named import:

  ```go
  import (
      foo "net/http"
  )
  ```

---

## 🧮 Data Types

* Integers, Floating Point, Complex Numbers
* Booleans, Strings
* Constants, Arrays, Structs, Pointers
* Maps, Slices, Functions, Channels
* JSON, HTML/Text templates

---

## 🔤 Variables

```go
var age int = 10
age := 10 // shorthand, can't be used with var
```

### Default Values

* Numeric: `0`
* Boolean: `false`
* String: `""`
* Pointers/maps/slices/etc.: `nil`

---

## 🔠 Naming Conventions

| Element            | Convention   |
| ------------------ | ------------ |
| Structs/Interfaces | `PascalCase` |
| Variables/Files    | `snake_case` |
| Constants          | `UPPERCASE`  |
| Booleans/Libraries | `mixedCase`  |

---

## ➕ Operators & Flow

### Arithmetic

* Overflow: Exceeds max data type value
* Underflow: Below min value (loss of precision)

### Logical

* `!` NOT, `||` OR, `&&` AND

### `if`/`else`

Standard `if`, `else if`, and `else` syntax

---

## 🔁 Loops

### `for` Loop

```go
for i := 0; i < 10; i++ {
    // loop logic
}
```

* `break`: exit loop
* `continue`: skip current iteration

---

## 🧮 Arrays & Slices

### Arrays

```go
var arr [3]int = [3]int{1, 2, 3}
```

### Multi-dimensional

```go
var matrix [3][3]int = [3][3]int{
    {1,2,3},
    {4,5,6},
    {7,8,9},
}
```

### Slices

* Dynamic arrays
* Operations: `copy`, `append`, `make`, `cap`, `slices.Equal`

---

## 🗺️ Maps

```go
var m map[string]int = make(map[string]int)
m = map[string]int{"key": 1}
delete(m, "key")
_, exists := m["key"]
```

---

## 🔁 `range` Keyword

Used for iteration:

* Arrays/Slices/Strings: sequential
* Maps: unordered
* Channels: until closed

Use `_` to ignore values

---

## 🧭 Functions

### Basic Syntax

```go
func add(a int, b int) int {
    return a + b
}
```

* **First-class citizens**
* **Exported functions**: Uppercase names
* **Private/internal**: Lowercase names

### Multiple Returns

```go
func compare(a, b int) (string, error) {
    // logic
}
```

### Variadic

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

---

## 🧼 `defer`

Postpones execution until surrounding function returns

### Use Cases

* Resource cleanup
* Unlock mutexes
* Logging

**Best Practices**

* Keep it short
* Understand evaluation timing
* Avoid complex control flow

---

## 😱 `panic` and `recover`

### `panic(interface{})`

* Abruptly stops normal execution

### `recover()`

* Regains control of panic
* Only works within `defer`

**Use Cases**

* Graceful error recovery
* Logging
* Prevent full crash

**Best Practices**

* Don’t abuse
* Avoid silent recovery
* Combine with `defer`

---

## 🛑 Exit

* Used to terminate unrecoverable programs
* Avoid deferred calls
* Use proper status codes

---

## ⚙️ `init()` Function

* Special function executed **once** per package
* Runs before `main()`

**Use Cases**

* Setup configs
* Register components
* Initialize databases

**Best Practices**

* Avoid side effects
* Understand execution order
