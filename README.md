## Overview
mq is  a message queue system written in Go and backed by Redis.

## example

publish message
```go

connect := mq.OpenConnection("connection2", &mq.ConnectOptions{Addr: "127.0.0.1:6379", DB: 5})
queue := connect.OpenQueue("tang")
var deliveries = time.Now().Unix()
val := fmt.Sprintf("%d", deliveries)
queue.Publish(val)

```



receive messages
```go
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
    




```