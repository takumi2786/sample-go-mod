package main

import (
	"fmt"

	"github.com/takumi2786/takumi2786/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
