package routers

import (
	"github.com/feitianlove/wechat/services/websocket"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.LoadHTMLGlob("views/**/*")
}

// web socket 的初始化
func WebSocketInit() {
	websocket.Register("login", websocket.LoginController)
}
