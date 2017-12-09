package redisBus

import (
	"github.com/go-redis/redis"
)

// var redisAddress = "localhost"

os.Getenv("REDIS_ADDRESS")
// var redisPort = "6379"

os.Getenv("REDIS_PORT")

//Client ...
// Interface for interacting with redis
var Client = &redis.Client{}
