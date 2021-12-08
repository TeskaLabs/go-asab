package asab

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-zookeeper/zk"
)

type ZookeeperService struct {
	Service

	Address    []string
	BasePath   string
	Connection *zk.Conn

	logger *log.Logger
}

func (svc *ZookeeperService) Initialize(app *Application) {
	svc.Service.Initialize(app)
	var err error

	cfgsec, err := app.Config.GetSection("zookeeper")
	if err != nil {
		panic(err)
	}

	svc.BasePath = cfgsec.Key("path").String()
	svc.Address = cfgsec.Key("address").Strings(" ")

	svc.logger = log.StandardLogger()

	svc.connect()
}

func (svc *ZookeeperService) connect() {

	if svc.Connection != nil {
		// Already connected
		return
	}

	// Connect to Zookeeper
	zk_conn, zk_chan, err := zk.Connect(
		svc.Address,
		3*time.Second,
		zk.WithLogger(svc),
	)
	if err != nil {
		log.Warn("Failed to connect to Zookeeper", err)
		return
	}
	svc.Connection = zk_conn

	svc.handleEvent(zk_chan)
}

func (svc *ZookeeperService) handleEvent(zk_chan <-chan zk.Event) {
	for event := range zk_chan {
		switch event.Type {
		case zk.EventSession:
			switch event.State {

			// This event happens when a Zookeeper client is connected/reconnected
			case zk.StateHasSession:
				svc.App.PubSub.Publish(PubSubMessage{
					Name: "Zookeeper/connected!",
					I:    svc,
				})

				// After we are connected, switch to goroutine
				go svc.handleEvent(zk_chan)
				return
			}

		default:
			log.Warn("Unhandled event:", event, event.Type)
		}
	}
}

func (svc *ZookeeperService) Printf(message string, v ...interface{}) {
	svc.logger.Printf("ZookeeperService "+message, v...)
}
