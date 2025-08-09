package main

import (
	"api/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	app.Run()
}
