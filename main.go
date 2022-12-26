package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

var parser = flags.NewParser(&struct{}{}, flags.Default)

func main() {
	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Println("success ✅")
}
