package main

import "fmt"

func main() {
	app := NewApp()
	port := "3000"
	app.Run(fmt.Sprint(":", port))
}
