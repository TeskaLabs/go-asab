package main

import (
	"fmt"
	"time"

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

	// Schedule the application stop
	go func() {
		time.Sleep(3 * time.Second)
		MyApp.Stop()
	}()

	MyApp.Run()
}
