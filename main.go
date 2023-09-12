package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("✨ gofun ✨")
	if err := run(); err != nil {
		fmt.Printf("❌ error: [%T] %v\n", err, err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}
