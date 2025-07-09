# Go Language Concepts

Go (or Golang) is an open-source programming language designed for simplicity, reliability, and efficiency. Below are some core concepts:

## 1. Basics

- **Variables:** Declared using `var` or `:=`.
- **Constants:** Declared with `const`.
- **Functions:** Defined with `func` keyword.
- **Packages:** Code is organized in packages; entry point is `main` package.

## 2. Data Types

- **Primitive:** `int`, `float64`, `string`, `bool`
- **Composite:** `array`, `slice`, `map`, `struct`

## 3. Control Structures

- **If/Else:** Standard conditional logic.
- **Switch:** Multi-way branch.
- **For:** Only looping construct (no `while`).

## 4. Pointers

- Variables that store memory addresses, using `*` and `&`.

## 5. Structs and Interfaces

- **Structs:** Custom data types.
- **Interfaces:** Define method sets for polymorphism.

## 6. Concurrency

- **Goroutines:** Lightweight threads managed by Go.
- **Channels:** Communicate between goroutines.

## 7. Error Handling

- Errors are values, handled explicitly (no exceptions).

## 8. Standard Library

- Rich set of packages for I/O, networking, and more.

## Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

## Resources

- [Go Official Site](https://golang.org/)
- [Tour of Go](https://tour.golang.org/)
