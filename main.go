package main

import (
	_ "books/docs"
	"books/routers"
)

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}
