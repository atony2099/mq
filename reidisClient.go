package mq

type RedisClient interface {
	SAdd(key, value string) bool
	LPush(key string, value ...string) bool
	BRPopLPush(source, destination string) (value string, succes bool)
	RPopLPush(source, destination string) (value string, success bool)
	LRem(key string, count int64, value string) (affected int64, ok bool)
	LLen(key string) (count int, ok bool)
}
