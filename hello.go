package main

import (
	"fmt"

	"github.com/pakkaparn/greeting"
)

func main() {
	message, err := greeting.Hello("pan")

	if err != nil {
		fmt.Printf("Oh! no one here.")
	}

	fmt.Printf(message)
}
