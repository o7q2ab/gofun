package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("✨ gofun ✨")
	if err := run(); err != nil {
		fmt.Printf("❌ error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}
