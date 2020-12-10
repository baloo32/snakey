# snakey
[![GoDoc](https://godoc.org/github.com/baloo32/snakey?status.svg)](https://godoc.org/github.com/baloo32/snakey)

This is a helper class to convert text to snake_case, removing any non-alphanumeric characters and spaces.

To use, import the package:

```go
import (
  github.com/baloo32/snakey
)
```
Use method `TextToSnake` passing a string to receive a **snake_case** response:

```go
s := "This is my string-example!!!"
res := snakey.TextToSnake(s)
```
Output
```
this_is_my_string_example
````