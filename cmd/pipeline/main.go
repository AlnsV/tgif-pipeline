package main

import (
	"github.com/AlnsV/go-crypto-ws-gateway/wsgateway"
	"github.com/sirupsen/logrus"
	"tgif-pipeline/internal/config"
)

var (
	logger = logrus.New()
)

func handleMessage(msg map[string]interface{}) {
	logger.Infoln(msg)
}

func main() {
	cfg, _ := config.New()

	client, err := wsgateway.BuildWSClient(
		"FTX",
		cfg.FTXAPIKey,
		cfg.FTXAPISecret,
	)
	if err != nil {
		logger.Error(err)
	}

	err = client.Connect()
	if err != nil {
		logger.Error(err)
	}

	forever := make(chan bool)
	err = client.Listen([]string{"BTC-PERP", "SOL-PERP"}, handleMessage)
	if err != nil {
		logger.Error(err)
	}
	<-forever
}
