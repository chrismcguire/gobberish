Gobberish
=========
Building random strings to help test unic�de issues.

Usage
-----
Importing
```Go
import "github.com/chrismcguire/gobberish"
```

Generating a random utf-8 string of a specified length
```Go
myString := gobberish.GenerateString(17)
```

You can also specify unicode ranges. For instance, to generate a 5 character string composed of Greek and Latin code points
```Go
gobberish.GenerateStringInRange(5, unicode.Greek, unicode.Latin)
```
A full list of character ranges can be found at http://golang.org/pkg/unicode/#pkg-variables
