package main

import (
	"fmt"

	"github.com/riteshpradhan/go-prompt"
)

// MyFriend is custom type
type MyFriend struct {
	Name string
	Age  int
}

func createListInterface() []interface{} {
	friend1 := MyFriend{Name: "Matt", Age: 33}
	friend2 := MyFriend{Name: "Jing", Age: 25}
	friend3 := MyFriend{Name: "Arya", Age: 20}

	return []interface{}{
		friend1,
		friend2,
		friend3,
	}
}

func main() {
	myList := createListInterface()
	i := prompt.ChooseInterface("What's your favorite friend?", myList)
	fmt.Printf("picked: %+v\n", myList[i])
}
