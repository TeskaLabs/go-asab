package asab

import (
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/lesismal/nbio/logging"
	"github.com/lesismal/nbio/nbhttp"
)

type WebService struct {
	Service

	Router *mux.Router
	Server *nbhttp.Server

	// Logging
	logger *log.Logger
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

	logging.SetLogger(svc)
	svc.logger = log.StandardLogger()

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

// Logging

func (svc *WebService) SetLevel(lvl int) {
	switch lvl {
	case logging.LevelAll:
		svc.logger.SetLevel(log.TraceLevel)

	case logging.LevelDebug:
		svc.logger.SetLevel(log.DebugLevel)

	case logging.LevelInfo:
		svc.logger.SetLevel(log.InfoLevel)

	case logging.LevelWarn:
		svc.logger.SetLevel(log.WarnLevel)

	case logging.LevelError:
		svc.logger.SetLevel(log.ErrorLevel)

	case logging.LevelNone:
		svc.logger.SetLevel(log.FatalLevel)

	default:
		log.Warnf("Invalid log level: %v", lvl)
	}
}

func (svc *WebService) Debug(format string, v ...interface{}) {
	svc.logger.Debugf("WebService "+format, v...)
}

func (svc *WebService) Info(format string, v ...interface{}) {
	svc.logger.Infof("WebService "+format, v...)
}

func (svc *WebService) Warn(format string, v ...interface{}) {
	svc.logger.Warnf("WebService "+format, v...)
}

func (svc *WebService) Error(format string, v ...interface{}) {
	svc.logger.Errorf("WebService "+format, v...)
}
