
forked from segmentio/go-prompt

# go-prompt

 Terminal prompts for Go.

 View the [docs](http://godoc.org/pkg/github.com/segmentio/go-prompt).

## Example

1.
```go
package main

import "github.com/segmentio/go-prompt"

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

2.
```go
package main

import (
	"fmt"
	"os"

	"bufio"

	prompt "github.com/segmentio/go-prompt"
)

func main() {
	println("Getting Started!")
	scanner := bufio.NewScanner(os.Stdin)
	longLine := prompt.Stringln(scanner, "Enter a long senctence")
	singleWord := prompt.String("Single Word")
	fmt.Printf("\nHello \n long: %s \n short: %s\n", longLine, singleWord)
}
```
