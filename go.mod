module tgif-pipeline

go 1.17

require (
	github.com/AlnsV/go-amqp v0.0.0-20200705021909-f58fe1c148d2
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/pkg/errors v0.9.1
)

require github.com/streadway/amqp v1.0.0 // indirect

require (
	github.com/AlnsV/go-crypto-ws-gateway v0.0.2-alpha
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)
