package asab

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/lesismal/nbio/nbhttp"
)

type WebService struct {
	Service

	Router *mux.Router
	Server *nbhttp.Server
}

func init() {
	AddConfigDefaults("web", map[string]string{
		"listen": ":8891 [::]:8891",
	})
}

func (svc *WebService) Initialize(app *Application) {
	svc.Service.Initialize(app)

	cfgsec, err := app.Config.GetSection("web")
	if err != nil {
		panic(err)
	}

	svc.Router = mux.NewRouter()

	svc.Server = nbhttp.NewServer(nbhttp.Config{
		Network: "tcp",
		Addrs:   cfgsec.Key("listen").Strings(" "),
		Handler: svc.Router,
	})

	err = svc.Server.Start()
	if err != nil {
		log.Fatalln("WebServer failed to start:", err)
	}

}

func (svc *WebService) Finalize() {
	svc.Server.Stop()
	svc.Service.Finalize()
}
