package main

import (
	"fmt"
	"os"

	"bufio"

	"github.com/riteshpradhan/go-prompt"
)

func main() {
	println("Getting Started!")
	scanner := bufio.NewScanner(os.Stdin)
	longLine := prompt.Stringln(scanner, "Enter a long senctence")
	singleWord := prompt.String("Single Word")
	fmt.Printf("\nHello \n long: %s \n short: %s\n", longLine, singleWord)
}
