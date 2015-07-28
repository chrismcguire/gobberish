Gobberish
=========
Building random strings to fuzz test unicode handling.

[![Build Status](https://travis-ci.org/chrismcguire/gobberish.svg?branch=master)](https://travis-ci.org/chrismcguire/gobberish)
[![GoDoc](https://godoc.org/github.com/chrismcguire/gobberish?status.svg)](https://godoc.org/github.com/chrismcguire/gobberish)

Usage
-----
Importing
```Go
import "github.com/chrismcguire/gobberish"
import "fmt"
```

Generating a random utf-8 string of a specified length
```Go
testString := gobberish.GenerateString(15)
fmt.Println(testString)

>> "티냮絥䯩얻橸禸䨃ȲСᣫ흳乃!*
```

You can also specify unicode ranges. For instance, to generate a 5 character string composed of Greek and Latin code points
```Go
greekAndLatin := gobberish.GenerateStringInRange(5, unicode.Greek, unicode.Latin)
fmt.Println(greekAndLatin)

>> "ƈꝨĶῈΊ"
```
A full list of character ranges can be found at http://golang.org/pkg/unicode/#pkg-variables
