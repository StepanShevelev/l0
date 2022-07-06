package pub

import (
	cfg "github.com/StepanShevelev/l0/pkg/config"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func Publish() {

	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		logrus.Fatal(err)
	}

	sc, err := stan.Connect(config.Cluster, config.Client, stan.NatsURL(config.Url))
	logrus.Info("Pub connected")
	if err != nil {
		logrus.Fatal(err)
	}

	jsonFile, err := ioutil.ReadFile("pkg/pub/model.json")
	if err != nil {
		logrus.Fatal(err)
	}

	sc.Publish("test", jsonFile)
	sc.Close()

}
