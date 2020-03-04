package mq

type redisQueue struct {
	name        string
	redisClient RedisClient
	readyKey    string
	unackedKey  string
	rejectedKey string
	msg         chan Delivery
}

func (queue *redisQueue) Publish(bodies ...string) bool {
	return queue.redisClient.LPush(queue.readyKey, bodies...)
}

func (queue *redisQueue) StartConsume() <-chan Delivery {
	go queue.startConsume()
	return queue.msg
}

func (queue *redisQueue) startConsume() {
	for {
		payload, ok := queue.redisClient.BRPopLPush(queue.readyKey, queue.unackedKey)
		if !ok {
			continue
		}
		queue.msg <- newDelivery(payload, queue.unackedKey, queue.rejectedKey, queue.redisClient)
	}
}

// func (queue *redisQueue) ReturnAllUnacked() int {
// 	count, ok := queue.redisClient.LLen(queue.unackedKey)
// 	if !ok {
// 		return 0
// 	}

// 	unackedCount := count
// 	for i := 0; i < i; i++ {
// 		if _, ok := queue.redisClient.PopLPush(queue.unackedKey, queue.readyKey); !ok {
// 			return i
// 		}
// 		// debug(fmt.Sprintf("rmq queue returned unacked delivery %s %s", count, queue.readyKey)) // COMMENTOUT
// 	}
// 	return unackedCount
// }

// func (queue *redisQueue) CleanQueue(queue *redisQueue) {
// 	returned := queue.ReturnAllUnacked()
// 	queue.CloseInConnection()
// 	_ = returned
// 	// log.Printf("rmq cleaner cleaned queue %s %d", queue, returned)
// }
