package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	natsURL := os.Getenv("NATS_URL")
	user := os.Getenv("NATS_USERNAME")
	psw := os.Getenv("NATS_PASSWORD")
	subject := os.Getenv("NATS_SUBJECT")
	queue := os.Getenv("NATS_QUEUE")

	nc, err := nats.Connect(natsURL, nats.UserInfo(user, psw))
	if err != nil {
		log.Fatal(err)
	}

	_, err = nc.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		log.Printf("get data  %q \n", msg.Data)
	})
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	<-c
}
