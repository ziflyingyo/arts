package main

import (
	"bytes"
	"fmt"

	"github.com/streadway/amqp"
)

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Println("msg", err, msg)
		panic("error")
	}
}

func mqConnect() {
	var err error
	conn, err = amqp.Dial(mqurl)
	failOnErr(err, "connect")

	channel, err = conn.Channel()
	failOnErr(err, "channel")
}

func close() {
	channel.Close()
	conn.Close()
}

func push() {
	if channel == nil {
		mqConnect()
	}

	mgsConnect := "hello world"
	err := channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	failOnErr(err, "ExchangeDeclare")

	_, err = channel.QueueDeclare(queueName, true, false,
		false, false, nil)
	failOnErr(err, "QueueDeclare")

	err = channel.QueueBind(queueName, "info", exchange, false, nil)
	failOnErr(err, "QueueBind")

	err = channel.Publish(exchange, "info", false, false, amqp.Publishing{
		ContentType: "text/plain", Body: []byte(mgsConnect),
	})
	failOnErr(err, "Publish")

	fmt.Println("push ok")
}

func receive() {
	if channel == nil {
		mqConnect()
	}

	msg, ok, err := channel.Get(queueName, false)
	failOnErr(err, "")
	if !ok {
		fmt.Println("do not get msg")
		return
	}

	err = channel.Ack(msg.DeliveryTag, false)
	failOnErr(err, "")

	s := BytesToString(&(msg.Body))
	fmt.Printf("receve msg is :%s\n", *s)
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

var conn *amqp.Connection
var channel *amqp.Channel

const (
	queueName = "hello1"
	exchange  = "exchange1"
	mqurl     = "amqp://guest:guest@127.0.0.1:5672/"
)

func main() {
	push()
	receive()
	fmt.Println("end")
	close()
}
