# Expr – a tiny stack-based virtual machine written in Go

[![Build Status](https://travis-ci.com/regeda/expr.svg?branch=main)](https://travis-ci.com/regeda/expr)
[![codecov](https://codecov.io/gh/regeda/expr/branch/main/graph/badge.svg?token=99QXNC2IAO)](https://codecov.io/gh/regeda/expr)
[![Go Reference](https://pkg.go.dev/badge/gihub.com/regeda/expr.svg)](https://pkg.go.dev/github.com/regeda/expr)

The executor is designed to interpret **a simple expression language** and it's useful in delegating decision enforcement to user scripts.

User scripts can produce scalar or array variables.

You can easily embed the toolchain in your Go application.

The executor is blazingly fast and it makes no memory allocation.

You can add custom functions to the executor empowering your application.

## Syntax

### Data types

*integer* (64 bit)
```js
-9223372036854775808 .. 9223372036854775807
```

*string* (wrapped by double quote)
```js
"Hello Expr!"
"Welcome \"Alice\" and \"Bob\""
```

*boolean*
```js
true, false
```

*array* (a vector of elements)
```js
[1, true, "text"]
[["hello"], "world!"]
```

### Operators

The virtual machine supports basic math operators `+-*/`. A math expression might be surrounded by parentheses.
Examples:
```
1 + -1
1 * (2 + 3)
```

### Delegators

In general, delegators are functions implemented by the hosted application.

Helpfully, this toolchain is equipped with an own standard library to handle basic operation on its data types.

#### stdlib

###### `concat(string, ...)`
returns a concatenated string
```js
concat("a", "b", "c") // "abc"
```

###### `join(string, [string, ...])`
returns a concatenated string with a separator
```js
join(", ", ["a", "b", "c"]) // "a, b, c"
```

###### `equals(string, string)`
###### `equals(int64, int64)`
###### `equals([string, ...], [string, ...])`
###### `equals([int64, ...], [int64, ...])`
returns `true` if both arguments are equal
```js
equals(1, 1) // true
equals(1, 0) // false
equals("foo", "foo") // true
equals("foo", "bar") // false
equals(["foo", 1], ["foo", 1]) // true
equals(["foo"], ["bar"]) // false
```

###### `intersects([string, ...], [string, ...])`
###### `intersects([int64, ...], [int64, ...])`
returns `true` if both arrays share the same item
```js
intersects([1, 2, 3], [3, 4]) // true
intersects([1, 2, 3], [4, 5]) // false
```

###### `contains([string, ...], string)`
###### `contains([int64, ...], int64)`
returns `true` if the value exists in the array
```js
contains([1, 2, 3], 1) // true
contains([1, 2, 3], 4) // false
```

## Architecture

The architecture consists of 3 components:
1. Lexer
2. Compiler
3. Virtual Machine

**The lexer** parses the input text:
```
join(",", ["a", "b"])
```
and generates a syntax tree:
```
STR ","
STR "a"
STR "b"
ARR 2
INVOKE join 2
```
> The lexer is implemented using [Ragel State Machine Compiler](https://www.colm.net/open-source/ragel/).

**The compiler** makes a bytecode from the syntax tree to make it executable by **a stack-based virtual machine**.
> The bytecode is described by [Flatbuffers](https://google.github.io/flatbuffers/flatbuffers_guide_use_go.html) to achieve high-throughput with low memory consumption.

## Usage

Compilation:
```go
import (
    "github.com/regeda/expr/compiler"
    "github.com/regeda/expr/lexer"
)

code := `join(",", ["a", "b"])`

tokens, err := lexer.Parse([]byte(code))
if err != nil {
    panic(err)
}

bytecode := compiler.Compile(tokens)

// save `bytecode` to be executed by the virtual machine
```

Running:
```go
import (
    "github.com/regeda/expr/delegate"
    "github.com/regeda/expr/exec"
    "github.com/regeda/expr/stdlib"
)

bytecode := ... // read []byte

ex := exec.New(
    exec.WithRegistry(delegate.Import(stdlib.Compare, stdlib.Strings)),
)
addr, err := ex.Exec(bytecode)
if err != nil {
    panic(err)
}
// `addr` contains the result, see github.com/regeda/expr/memory.Addr
```
> `Exec` is **not designed** to be run in the concurrent environment. However, you can define a pool of executors to consume them in the safe mode.

## Benchmark

The benchmark executes a compiled bytecode of the following statement:
```js
equals("foo,bar,baz", join(",", ["foo", "bar", "baz"]))
```
```
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkExec
BenchmarkExec-8          1635091               746.7 ns/op             0 B/op          0 allocs/op
```
