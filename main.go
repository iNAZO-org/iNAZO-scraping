package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
)

var parser = flags.NewParser(&struct{}{}, flags.Default)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func main() {
	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Println("success ✅")
}
