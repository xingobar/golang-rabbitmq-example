package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
)

const (
	queueName = "hello"
	exchange = "exchange1"
	mqUrl = "amqp://guest:guest@localhost:5672/"
)

var channel *amqp.Channel
var conn *amqp.Connection

func main() {
	push()
	receive()
	close()
}

func push() {
	if channel == nil {
		mqConnect()
	}

	mgsConnect := "hello world"
	err := channel.ExchangeDeclare(exchange, "direct", false, false, false, false, nil)
	failOnError(err, " exchange declare")

	_, err = channel.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, " queue declare")

	err = channel.QueueBind(queueName, "info", exchange, false, nil)
	failOnError(err, " queue bind")


	err = channel.Publish(exchange, "info", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(mgsConnect),
	})
	failOnError(err, " Publish")
	fmt.Println("push ok")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println("msg: ", err, msg)
		panic(msg)
	}
}

// rabbitmq 連線
func mqConnect() {
	// 連接 rabbitmq
	var err error
	conn, err = amqp.Dial(mqUrl)
	failOnError(err, "mqconnect")

	// 建立 channel
	channel, err = conn.Channel()
	failOnError(err, "channel")
}

// rabbitmq close
func close() {
	fmt.Println("rabbitmq close connection .....")
	channel.Close()
	conn.Close()
}

func receive() {
	if channel == nil {
		mqConnect()
	}
	msg, ok, err := channel.Get(queueName, false)

	failOnError(err, "")
	if !ok {
		fmt.Println("do not get msg")
		return
	}

	err = channel.Ack(msg.DeliveryTag, true)
	failOnError(err, "")

	fmt.Println(bytes.NewBuffer(msg.Body).String())
}