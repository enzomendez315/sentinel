package main

import (
	"fmt"

	"github.com/enzomendez315/sentinel/internal/server"
)

func main() {
	err := server.Run()
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
