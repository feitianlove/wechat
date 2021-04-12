package websocket

import (
	"fmt"
	"github.com/feitianlove/wechat/models"
	"github.com/feitianlove/wechat/services/grpcclient"
	"github.com/feitianlove/wechat/services/redis_cache"
	"time"
)

// 给全体用户发消息
func SendUserMessageAll(appId uint32, userId string, msgId, cmd, message string) (sendResults bool, err error) {
	sendResults = true

	currentTime := uint64(time.Now().Unix())
	servers, err := redis_cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		if IsLocal(server) {
			data := models.GetMsgData(userId, msgId, cmd, message)
			AllSendMessages(appId, userId, data)
		} else {
			grpcclient.SendMsgAll(server, msgId, appId, userId, cmd, message)
		}
	}

	return
}
