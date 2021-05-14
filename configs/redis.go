package configs

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

var (
	redisConn redis.Conn
	redisErr  error
)

func InitRedis() {
	var port = os.Getenv("REDIS_PORT")
	redisConn, redisErr = redis.Dial("tcp", "localhost:"+port)
	if redisErr != nil {
		fmt.Println("Redis Connection Error:", err.Error())
		return
	}
}

func GetRedis() redis.Conn {
	return redisConn
}

func GetRedisDb() string {
	return os.Getenv("REDIS_DB")
}