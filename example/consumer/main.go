package main

import (
	"fmt"
	"mq"
)

func main() {
	connect := mq.OpenConnection("connection1", &mq.ConnectOptions{Addr: "127.0.0.1:6379", DB: 5})
	queue := connect.OpenQueue("tang")
	c := queue.StartConsume()
	go func() {
		for e := range c {
			fmt.Println("eleme is:", e)
			res := e.Ack()
			fmt.Println(res)
		}
	}()
	select {}
}
