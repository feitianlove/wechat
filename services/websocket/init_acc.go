package websocket

import (
	"fmt"
	"github.com/feitianlove/wechat/config"
	"github.com/feitianlove/wechat/lib"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

var (
	appIds        = []uint32{101, 102, 103, 104} // 全部的平台
	clientManager = NewClientManager()
)

func InAppIds(appId uint32) (inAppId bool) {
	for _, value := range appIds {
		if value == appId {
			inAppId = true
			return
		}
	}
	return
}

// 启动程序
func StartWebSocket(config *config.Config) {

	serverIp = lib.GetServerIp()
	serverPort := config.App.RpcPort

	http.HandleFunc("/acc", wsPage)

	// 添加处理程序
	go clientManager.start()
	fmt.Println("WebSocket 启动程序成功", serverIp, serverPort)

	func() {
		_ = http.ListenAndServe(":"+strconv.Itoa(int(config.App.WebSocketPort)), nil)
	}()

}

func wsPage(w http.ResponseWriter, req *http.Request) {

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client
}
