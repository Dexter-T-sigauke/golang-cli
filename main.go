package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// string flag for the user's name
	name := flag.String("name", "Dex", "The name to greet.")

	// command-line arguments
	flag.Parse()

	// Greeting the user by name
	fmt.Printf("Hello, %s!\n", *name)

	// Checking for additional arguments after the flag
	if len(os.Args) > 2 {
		fmt.Println("Warning: Unexpected arguments provided.")
	}
}
