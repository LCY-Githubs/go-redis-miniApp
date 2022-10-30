package main

import (
	"src/app/routers"
)

// @version 1.0
// @title gin+gorm crud 测试swagger
// @description 这里是描述
func main() {
	router := routers.Router{}
	//gin init
	router.Init()
	//connect to mysql server
	//conn.InitMysqlClientConnection()
}
