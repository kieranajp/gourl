package main

import (
	"github.com/kataras/iris"
	"gopkg.in/redis.v3"
)

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis",
	})
}

func main() {
	redisStorage := initRedis()
	defer redisStorage.Close()

	iris.Get("/", func(c *iris.Context) {
		c.JSON(iris.StatusOK, iris.Map{
			"Hello": "World",
		})
	})
	iris.Listen("0.0.0.0:3001")
}
