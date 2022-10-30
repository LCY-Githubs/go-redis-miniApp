package conn

import (
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//定义一个全局的rdb变量
var (
	rdb *redis.Client
	//定义一个全局的orm变量
	db *gorm.DB
)

// InitClient 初始化链接
func InitClient() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return nil
}

func InitMysqlClientConnection() {
	var dsn = "root:000000@tcp(127.0.0.1:3306)/small_web?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // 账号密码地址端口
	}), &gorm.Config{})
	fmt.Println(db)
}

func GetDB() *gorm.DB {
	var dsn = "root:000000@tcp(127.0.0.1:3306)/small_web?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // 账号密码地址端口
	}), &gorm.Config{})
	return db
}
