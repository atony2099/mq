package main

import (
	"fmt"
	"mq"
	"time"
)

func main() {

	connect := mq.OpenConnection("connection2", &mq.ConnectOptions{Addr: "127.0.0.1:6379", DB: 5})
	queue := connect.OpenQueue("tang")
	var deliveries = time.Now().Unix()
	val := fmt.Sprintf("%d", deliveries)
	queue.Publish(val)

}
