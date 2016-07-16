package main

import (
	"fmt"
	"github.com/kataras/iris"
	"gopkg.in/redis.v4"
)

var redisStorage = redis.NewClient(&redis.Options{
	Addr: "redis:6379",
})

type UrlAPI struct {
	*iris.Context
}

func (u UrlAPI) Get() {
	u.Write("Hello world")
}

func (u UrlAPI) GetBy(id string) {
	url, err := redisStorage.Get(fmt.Sprintf("URL_%s", id)).Result()

	if err != nil {
		u.JSON(iris.StatusNotFound, iris.Map{
			"error": fmt.Sprintf("Not found: %s", id),
		})
		return
	}

	u.JSON(iris.StatusOK, iris.Map{
		"location": url,
	})
}

func (u UrlAPI) Post() {
	url := u.FormValue("url")
	key := getNextKey()

	err := redisStorage.Set(fmt.Sprintf("URL_%d", key), url, 0).Err()

	if err != nil {
		u.JSON(iris.StatusInternalServerError, iris.Map{
			"error": "Failed to save",
		})
	}

	u.JSON(iris.StatusCreated, iris.Map{
		"location": url,
		"key":      key,
	})
}

func getNextKey() int64 {
	if err := redisStorage.Incr("key").Err(); err != nil {
		panic(err)
	}

	key, _ := redisStorage.Get("key").Int64()

	return key
}

func main() {
	defer redisStorage.Close()

	iris.API("/", UrlAPI{})
	iris.Listen("0.0.0.0:3001")
}
