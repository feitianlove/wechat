package main

import (
	"fmt"
	"github.com/feitianlove/wechat/config"
	"github.com/feitianlove/wechat/logger"
	"github.com/feitianlove/wechat/redis"
	"github.com/feitianlove/wechat/routers"
	"github.com/feitianlove/wechat/services/websocket"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	// 初始化config
	cfg, err := config.InitConfig("./etc/wechat.conf")
	if err != nil {
		panic(err)
	}
	// 初始化log
	err = logger.InitLog(cfg)
	if err != nil {
		panic(err)
	}
	// 初始化redis
	redisCli, err := redis.NewRedisClient(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(redisCli)
	//初始化 gin
	router := gin.Default()
	// 初始化websocket  router
	routers.WebSocketInit()
	// 初始化路由
	routers.Init(router)
	//启动websocket
	websocket.StartWebSocket(cfg)
	func() {
		_ = http.ListenAndServe(":"+strconv.Itoa(int(cfg.App.HttpPort)), router)
	}()

	http.ListenAndServe(":"+strconv.Itoa(int(cfg.App.HttpPort)), router)
}
