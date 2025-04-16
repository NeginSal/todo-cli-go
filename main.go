package main

import (
	"fmt"
)

type Task struct {
	ID          int    `json:"id"`
	title       string `json:"title"`
	description string `json:"description"`
	status      bool   `json:"status"`
}

func main() {
	fmt.Println("hello world")
}
