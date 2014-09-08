package main

import "fmt"
import ".."

func main() {
	println("need your name!")
	first := prompt.StringRequired("first")
	last := prompt.StringRequired("last")
	fmt.Printf("\nHello %s %s\n", first, last)
}
