left-pad
========
Go package to left pad a string with a character.

### Install

```
go get github.com/keltia/leftpad 
```

### Usage

```Go
import 'github.com/keltia/leftpad'

leftpad.Leftpad("foo", 5)       # => "  foo"
leftpad.Leftpad("foobar", 8)    # => "  foobar"
leftpad.Leftpad("foobar", 6)    # => "foobar"
leftpad.Leftpad(1, 10, 0)       # => "0000000001"
```