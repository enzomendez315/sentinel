package main

import "github.com/enzomendez315/sentinel/internal/server"

func main() {
	err := server.Run()
	if err != nil {
		return
	}
}
