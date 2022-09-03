package main

import (
	"fmt"

	"github.com/takumi2786/sample-go-mod/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys v1.0.0")
	fmt.Println(message)
}
