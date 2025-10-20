# Variables in Go (Golang)

Go provides multiple ways to declare and use variables. This guide summarizes the key concepts and syntax.

---

## 1. Declaring Variables

### Long Form (with type)

```go
var x int = 10
var name string = "Go"
```

### Type Inference

```go
var x = 10        // Go infers int
var name = "Go"   // Go infers string
```

### Short Declaration (inside functions only)

```go
x := 10
name := "Go"
```

---

## 2. Multiple Variable Declaration

```go
var a, b, c int
var x, y = 10, "Hello"
a, b := 1, 2
```

---

## 3. Constants (Immutable)

```go
const pi = 3.14
const greeting = "Hello"
```

---

## 4. Zero Values

If a variable is declared but not assigned, Go gives it a zero value:

| Type    | Zero Value |
| ------- | ---------- |
| int     | 0          |
| float   | 0.0        |
| string  | ""         |
| bool    | false      |
| pointer | nil        |

---

## 5. Exported vs Unexported Variables

Variables starting with a **capital letter** are exported (public).

```go
var MyName = "Public"
var myName = "Private"
```

---

## 6. Example Program

```go
package main

import "fmt"

var x int = 100
const pi = 3.14

func main() {
    y := 20
    fmt.Println(x, y, pi)
}
```

---

## Summary

* Use `var` for package-level variables or when you want explicit types.
* Use `:=` inside functions for quick declarations.
* Constants use `const` and cannot be modified.
* Go assigns zero values to uninitialized variables.

---

