left-pad
========
Go package to left pad a string with a character.

Inspired by the "left-pad" NPM package (and the fiasco that happened after its removal).

### Build status


Branch: master — [![master|Build Status](https://travis-ci.org/keltia/leftpad.svg?branch=master)](http://travis-ci.org/keltia/leftpad)

### Install

```
go get github.com/keltia/leftpad 
```

### Usage

[![GoDoc](https://godoc.org/github.com/keltia/leftpad?status.svg)](https://godoc.org/github.com/keltia/leftpad)

```Go
import 'github.com/keltia/leftpad'

leftpad.LeftPad("foo", 5)       # => "  foo"
leftpad.LeftPad("foobar", 8)    # => "  foobar"
leftpad.LeftPad("foobar", 6)    # => "foobar"

leftpad.LeftPadStr("foo", 5, "X") # => "XXfoo"
```
