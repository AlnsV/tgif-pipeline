package main

import (
	rabbit "github.com/AlnsV/go-amqp"
	"github.com/AlnsV/go-crypto-ws-gateway"
	"github.com/AlnsV/go-crypto-ws-gateway/pkg/model"
	"github.com/sirupsen/logrus"
	"log"
	"tgif-pipeline/internal/config"
)

var (
	logger = logrus.New()
)

func main() {
	cfg, _ := config.New()
	consumerBuffer := make(chan *model.Trade)
	forever := make(chan bool)
	startConsumer(cfg, consumerBuffer)
	client, err := setUpRabbit(cfg)
	if err != nil {
		log.Fatal(err)
	}
	go dispatchTrades(consumerBuffer, client)
	<-forever
}

func startConsumer(cfg *config.Config, bufferOutput chan<- *model.Trade) {
	client, err := wsgateway.BuildWSClient(
		"FTX",
		cfg.FTXAPIKey,
		cfg.FTXAPISecret,
	)

	err = client.Connect()
	if err != nil {
		logger.Error(err)
	}

	err = client.Listen(
		[]string{"BTC-PERP", "SOL-PERP"},
		func(trade *model.Trade) {
			bufferOutput <- trade
		},
	)

	if err != nil {
		logger.Error(err)
	}
}

func setUpRabbit(cfg *config.Config) (*rabbit.AMQPClient, error) {
	client := &rabbit.AMQPClient{}
	if err := client.StartConnection(
		cfg.RabbitUser,
		cfg.RabbitPWD,
		cfg.RabbitAddress,
		cfg.RabbitPort,
	); err != nil {
		return nil, err
	}
	err := client.SetupDispatcher("trades", "topic", true, false)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func dispatchTrades(messagesBuffer chan *model.Trade, rabbit *rabbit.AMQPClient) {
	for trade := range messagesBuffer {
		err := rabbit.SendMessage("trades", trade.Market, trade)
		if err != nil {
			log.Println(err)
		}
	}
}
