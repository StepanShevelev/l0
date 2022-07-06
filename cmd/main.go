package main

import (
	"fmt"
	"github.com/StepanShevelev/l0/pkg/api"
	cfg "github.com/StepanShevelev/l0/pkg/config"
	mydb "github.com/StepanShevelev/l0/pkg/db"
	"github.com/StepanShevelev/l0/pkg/pub"
	"github.com/StepanShevelev/l0/pkg/sub"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		logrus.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	exitCh := make(chan struct{}, 1)

	api.InitBackendApi()
	logrus.Info("API connected")

	mydb.ConnectToDb(config)
	logrus.Info("DB connected")

	err := mydb.GetOrderByUid()
	logrus.Info("Cache uploaded")
	if err != nil {
		logrus.Error(err)
		return
	}

	pub.Publish()

	sub.Connect()

	http.ListenAndServe(":"+config.Port, nil)

	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
	exitCh <- struct{}{}

}
