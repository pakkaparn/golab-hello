package main

import (
	"fmt"

	greeting "github.com/pakkaparn/golab-greeting"
)

func main() {
	message, err := greeting.Hello("pan")

	if err != nil {
		fmt.Printf("Oh! no one here.")
	}

	fmt.Printf(message)
}
