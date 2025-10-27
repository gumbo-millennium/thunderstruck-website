package main

import (
	"fmt"

	"github.com/gumbo-millennium/thunderstruck-website/migrations"
	_ "github.com/lib/pq"
)

func main() {
	migrations.Execute()
	fmt.Println("Hello, World!")
}
