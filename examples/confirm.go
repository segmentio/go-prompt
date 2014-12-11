package main

import ".."

func main() {
	if ok := prompt.Confirm("launch %s?", "something"); ok {
		println("launching")
	} else {
		println("not launching")
	}
}
