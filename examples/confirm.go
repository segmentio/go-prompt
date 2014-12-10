package main

import ".."

func main() {
	if ok := prompt.Confirm("launch new instance?"); ok {
		println("launching")
	} else {
		println("not launching")
	}
}
