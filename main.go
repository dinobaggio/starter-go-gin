package main

import "fmt"

func main() {
	app := App{}
	app.Initialize()

	port := "3000"
	app.Run(fmt.Sprint(":", port))
}
