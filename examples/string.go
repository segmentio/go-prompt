package main

import (
	"fmt"

	"github.com/riteshpradhan/go-prompt"
)

func main() {
	println("need your name!")
	first := prompt.String("first")
	last := prompt.String("last")
	def := prompt.StringDefault("Want default", "defaultValue")
	fmt.Printf("\nHello %s %s \n", first, last)
	fmt.Println("Default value is : ", def)
}
