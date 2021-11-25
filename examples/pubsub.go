package main

import (
	"github.com/teskalabs/go-asab/asab"
	"fmt"
)

type MyApplication struct {
	asab.Application

	TickCounter int
}

func (app *MyApplication) OnTick() {
	fmt.Printf("On tick %#v!\n", app.TickCounter)
	app.TickCounter += 1
}

func (app *MyApplication) Initialize() {
	app.Application.Initialize()

	app.TickCounter = 1
	app.PubSub.Subscribe("Application.tick!", app.OnTick)
}

func main() {
	MyApp := new(MyApplication)
	
	MyApp.Initialize()
	defer MyApp.Finalize()

	MyApp.Run()
}
