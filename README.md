todo
====

todo is a simple package that helps you keep your TODOs under control. Oftentimes developers write a TODO and forget about it. The todo package helps you remember those TODOs by allowing you to set an expiration date with the message. If that expiration date is hit, the program will exit and print the message, along with the location of the TODO.

[![wercker status](https://app.wercker.com/status/5c46f31b4c8b2ba071426379ef692894/m "wercker status")](https://app.wercker.com/project/bykey/5c46f31b4c8b2ba071426379ef692894)

##Usage

Usage of todo is very simple. Here is a complete example:

```go
package main

import "github.com/stretchr/todo"

func main() {
  to.Do("2014-May-21","Write the main function")
}
```

## Enabling

By default, todo is not enabled. This is to prevent accidental termination of the application in a production environment.

When disabled, the function it calls is empty. There is still a bit of overhead in passing the arguments by value, however.

In order to enable the todo package, set `TODO_ENABLED=1` as an environment variable. As a result of this requirement, the tests also depend on this environment variable being set. Use `TODO_ENABLED=1 go test` to test. Simlarly, to run the benchmarks, use `TODO_ENABLED=1 go test -bench=.`.

## Behavior

The `Do(by, msg string)` function takes a "short form" date (`YYYY-Month-DD`) as the first argument. Internally, the date is parsed and used for comparisons. If the date fails to parse, the function panics to inform the programmer of the error.

