package sub

import (
	"encoding/json"
	"github.com/StepanShevelev/l0/pkg/api"
	cfg "github.com/StepanShevelev/l0/pkg/config"
	mydb "github.com/StepanShevelev/l0/pkg/db"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
)

//var valid *validator.Validate

func Connect(valid *validator.Validate) {
	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		logrus.Fatal(err)
	}

	var order mydb.Order

	nc, err := nats.Connect(config.Url)
	if err != nil {
		logrus.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(config.Cluster, config.Client,
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logrus.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		logrus.Printf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, config.Url)
	}
	logrus.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", config.Url, config.Cluster, config.Client)

	sub, err := sc.Subscribe("test", func(msg *stan.Msg) {

		ProcessMessage(order, msg, valid)

	}, stan.StartWithLastReceived())

	// Close connection
	defer sc.Close()

	// Unsubscribe
	defer sub.Unsubscribe()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			logrus.Info("Received an interrupt, unsubscribing and closing connection...\n\n")

			os.Exit(0)

		}
	}()

}

func ProcessMessage(order mydb.Order, msg *stan.Msg, valid *validator.Validate) {

	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		logrus.Error("wrong data")

	}

	err = valid.Struct(&order)
	if err != nil {
		log.Println("wrong data")
		return
	}

	api.Caching.SetCache(order.OrderUID, order)

	logrus.Info("message: ", msg)
	logrus.Info("order: ", order)

	result := mydb.Database.Db.Create(&order)
	if result.Error != nil {

		logrus.Error("Could not create order", result.Error)

	}

}
