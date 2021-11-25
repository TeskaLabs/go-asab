package main

import (
	"fmt"

	"github.com/teskalabs/go-asab/asab"
)

type MyApplication struct {
	asab.Application
}

func (app *MyApplication) Initialize() {
	app.Application.Initialize()
}

func main() {
	MyApp := new(MyApplication)

	MyApp.Initialize()
	defer MyApp.Finalize()

	fmt.Println("Hello world!")
	fmt.Println("Press Ctrl-C for an exit.")
	MyApp.Run()
}
