# golang-rabbitmq-example
golang rabbitmq example

1. 宣告 exchange
2. 宣告 queue (Routing key)
3. 綁定 Exchange 和 Queue
4. 發布信息到 exchange 且指定 routing key, routing key 會將訊息送至指定的 queue

run:
1. go build main.go
2. ./main  -> 會發送消息至 exchange_1

package:
1.  go get github.com/streadway/amqp