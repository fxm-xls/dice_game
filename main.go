package main

import (
	"dice/common/db"
	"dice/common/ico"
	"dice/controller/check"
	"dice/controller/dice"
	"github.com/gin-gonic/gin"
)

func setup() {
	//1. 初始化数据库链接
	db.DBSetUp("root:123456@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local")
}

func main() {
	setup()
	engine := gin.Default()
	r := engine.Group("/dices")
	{
		r.POST("/check", ico.Handler(check.Check{}))
		r.POST("/throw", ico.Handler(dice.Throw{}))
		r.POST("/replay", ico.Handler(dice.Replay{}))
	}
	engine.Run()
}
