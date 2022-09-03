package main

import (
	"fmt"

	"github.com/takumi2786/sample-go-mod/echo/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
