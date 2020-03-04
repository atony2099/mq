package mq

import (
	"fmt"
	"strings"

	"gopkg.in/redis.v3"
)

const (
	queueKey              = "mq:queue"
	queueReadyTemplate    = "mq:queue:[{queue}]:ready"
	queueRejectedTemplate = "rmq::queue::[{queue}]::rejected"

	connectionQueueUnackedTemplate = "rmq:connection:[{connection}]:queue::[{queue}]::unacked"

	phConnection = "{connection}" // connection name
	phQueue      = "{queue}"
)

type redisConnection struct {
	Name  string
	redis RedisClient
}

type ConnectOptions struct {
	// The network type, either tcp or unix.
	// Default is tcp.
	Network string

	// host:port address.
	Addr string
	// An optional password. Must match the password specified in the
	// requirepass server configuration option.
	Password string
	// A database to be selected after connecting to server.
	DB int64
	// The maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int
}

// OpenConnection return a new connection
func OpenConnection(tag string, optins *ConnectOptions) *redisConnection {
	// 1. create connection
	redisClient := redis.NewClient(&redis.Options{PoolSize: 2, Addr: optins.Addr, Password: optins.Password, DB: optins.DB, MaxRetries: optins.MaxRetries})
	wrapper := &redisWrapper{redisClient}

	name := fmt.Sprintf("%s-%s", tag, RandStringBytes(6))

	connection := &redisConnection{
		Name:  name,
		redis: wrapper,
	}
	return connection
}

func (connection *redisConnection) OpenQueue(queueName string) *redisQueue {
	return connection.newQueue(queueName, connection.Name)
}

func (connection *redisConnection) newQueue(queueName, connectionName string) *redisQueue {
	// 1.add  queue
	connection.redis.SAdd(queueKey, queueName)

	readyKey := strings.Replace(queueReadyTemplate, phQueue, queueName, 1)
	rejectedKey := strings.Replace(queueRejectedTemplate, phQueue, queueName, 1)

	unackedKey := strings.Replace(connectionQueueUnackedTemplate, phConnection, connectionName, 1)
	unackedKey = strings.Replace(connectionQueueUnackedTemplate, phQueue, queueName, 1)

	// 2. create
	msg := make(chan Delivery)

	queue := &redisQueue{queueName, connection.redis, readyKey, unackedKey, rejectedKey, msg}

	return queue
}
