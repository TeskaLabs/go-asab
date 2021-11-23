package asab

import (
	"math/rand"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

func init() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())
}

type Application struct {
	_InterruptSignal chan os.Signal
	ReturnCode       int

	// Periodic 1-second ticker
	_Ticker *time.Ticker

	// Service registry
	// Note that the service doesn't have to be registered
	_RegisteredServices map[string]ServiceI

	Hostname string
	Config   *ini.File

	PubSub PubSub
}

func (app *Application) Initialize() {

	app.Config = _LoadConfig()
	app._InterruptSignal = make(chan os.Signal, 1)
	app._RegisteredServices = make(map[string]ServiceI)
	app.Hostname, _ = os.Hostname()
	app.ReturnCode = 0

	app._Ticker = time.NewTicker(1 * time.Second)

	// Prepare PubSub
	app.PubSub.Initialize()

	// Install SIGINT handler
	signal.Notify(app._InterruptSignal, os.Interrupt)
}

func (app *Application) Finalize() {
	// Uninstall SIGINT handler
	signal.Stop(app._InterruptSignal)

	if len(app._RegisteredServices) != 0 {
		log.Warnln("Application is exiting with registered services:")
		for name, svc := range app._RegisteredServices {
			log.Warnln(" *", name, svc)
		}
	}
}

func (app *Application) Run() {
	cycle_no := 0

	log.Info("Application is ready.")
	for {
		select {

		// Handling of SIGING signal
		case <-app._InterruptSignal:
			println("") // To put ^C on the dedicated line
			log.Println("Terminating.")
			return

		// Handling of the periodic ticking
		case <-app._Ticker.C:
			app.PubSub.Publish("Application.tick!")
			cycle_no += 1
			if (cycle_no % 5) == 0 {
				app.PubSub.Publish("Application.tick/5!")
			}
			if (cycle_no % 10) == 0 {
				app.PubSub.Publish("Application.tick/10!")
			}
			if (cycle_no % 60) == 0 {
				app.PubSub.Publish("Application.tick/60!")
			}
			if (cycle_no % 300) == 0 {
				app.PubSub.Publish("Application.tick/300!")
			}
			if (cycle_no % 600) == 0 {
				app.PubSub.Publish("Application.tick/600!")
			}
			if (cycle_no % 1800) == 0 {
				app.PubSub.Publish("Application.tick/1800!")
			}
			if (cycle_no % 3600) == 0 {
				app.PubSub.Publish("Application.tick/3600!")
			}
			if (cycle_no % 43200) == 0 {
				app.PubSub.Publish("Application.tick/43200!")
			}
			if (cycle_no % 86400) == 0 {
				app.PubSub.Publish("Application.tick/86400!")
			}
		}

	}
}

// Service registry

func (app *Application) RegisterService(svc ServiceI, service_name string) {
	app._RegisteredServices[service_name] = svc
}

func (app *Application) UnregisterService(svc ServiceI) {
	for iname, isvc := range app._RegisteredServices {
		if svc == isvc {
			delete(app._RegisteredServices, iname)
			return
		}
	}
	log.Warn("Cannot find service to unregister", svc)
}

func (app *Application) LocateService(service_name string) ServiceI {
	return app._RegisteredServices[service_name]
}
