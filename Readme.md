
forked from segmentio/go-prompt

# go-prompt

 Terminal prompts for Go.

 View the [docs](http://godoc.org/pkg/github.com/segmentio/go-prompt).

## How to use:

```go
package main

import "github.com/riteshpradhan/go-prompt"

var langs = []string{
  "c",
  "c++",
  "lua",
  "go",
  "js",
  "ruby",
  "python",
}

func main() {
  i := prompt.Choose("What's your favorite language?", langs)
  println("picked: " + langs[i])
}
```

## Functions:
* String
* StringRequired
* StringDefault
* Stringln
* Integer
* IntegerRequired
* Confirm
* Choose
* ChooseInterface
* Password
* PasswordMasked
