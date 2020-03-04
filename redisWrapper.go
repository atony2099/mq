package mq

import (
	"gopkg.in/redis.v3"
)

type redisWrapper struct {
	redisClient *redis.Client
}

func (wrapper *redisWrapper) SAdd(key, value string) bool {
	_, err := wrapper.redisClient.SAdd(key, value).Result()
	if err != nil {
		return false
	}

	return true

}

func (wrapper *redisWrapper) LPush(key string, value ...string) bool {
	_, err := wrapper.redisClient.LPush(key, value...).Result()
	return err == nil
}

func (wrapper *redisWrapper) RPOPLPUSH(source string, destination string) (string, bool) {
	value, err := wrapper.redisClient.RPopLPush(source, destination).Result()

	if err == redis.Nil {
		return "", false
	}

	return value, true
}

func (wrapper *redisWrapper) BRPopLPush(source, destination string) (string, bool) {
	value, err := wrapper.redisClient.BRPopLPush(source, destination, 0).Result()
	if err != nil {
		return "", false
	}
	return value, true
}

func (wrapper *redisWrapper) RPopLPush(source, destination string) (string, bool) {
	value, err := wrapper.redisClient.RPopLPush(source, destination).Result()
	if err != nil {
		return "", false
	}
	return value, true

}

func (wrapper *redisWrapper) LLen(key string) (int, bool) {
	count, err := wrapper.redisClient.LLen(key).Result()
	if err != nil {
		return 0, false
	}
	return int(count), true

}

func (wrapper *redisWrapper) LRem(key string, count int64, value string) (affected int64, ok bool) {
	count, err := wrapper.redisClient.LRem(key, count, value).Result()

	if err != nil {
		return 0, false
	}

	return count, true
}
