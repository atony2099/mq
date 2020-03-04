package mq

type Delivery interface {
	Payload() string
	Ack() bool
	Reject() bool
}

type wrapDelivery struct {
	payload     string
	redisClient RedisClient
	unackedKey  string
	rejectedKey string
}

func newDelivery(payload, unackedKey, rejectedKey string, redisClient RedisClient) *wrapDelivery {
	return &wrapDelivery{
		payload:     payload,
		unackedKey:  unackedKey,
		rejectedKey: rejectedKey,
		redisClient: redisClient,
	}
}

func (d *wrapDelivery) Payload() string {
	return d.payload
}

func (d *wrapDelivery) Ack() bool {
	count, ok := d.redisClient.LRem(d.unackedKey, -1, d.payload)
	return ok && count == 1
}

func (d *wrapDelivery) Reject() bool {
	return d.move(d.rejectedKey)
}

func (d *wrapDelivery) move(key string) bool {
	if ok := d.redisClient.LPush(key, d.payload); !ok {
		return false
	}
	if _, ok := d.redisClient.LRem(d.unackedKey, 1, d.payload); !ok {
		return false
	}

	return true
}
