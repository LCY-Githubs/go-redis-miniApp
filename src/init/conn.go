package init

import (
	"fmt"

	"github.com/go-redis/redis"
)

//定义一个全局的rdb变量
var rdb *redis.Client

//初始化链接
func initClient() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return nil
}
func Prints() {
	fmt.Println("Hello I am init")
}
