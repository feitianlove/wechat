package models

import "github.com/feitianlove/golib/common/ecode"

const (
	MessageTypeText = "text"

	MessageCmdMsg   = "msg"
	MessageCmdEnter = "enter"
	MessageCmdExit  = "exit"
)

// 消息的定义
type Message struct {
	Target string `json:"target"` // 目标
	Type   string `json:"type"`   // 消息类型 text/img/
	Msg    string `json:"msg"`    // 消息内容
	From   string `json:"from"`   // 发送者
}

// 文本消息
func GetMsgData(uuId, msgId, cmd, message string) string {
	return getTextMsgData(cmd, uuId, msgId, message)
}

func getTextMsgData(cmd, uuId, msgId, message string) string {
	textMsg := NewTestMsg(uuId, message)
	head := NewResponseHead(msgId, cmd, ecode.OK, "Ok", textMsg)

	return head.String()
}

func NewTestMsg(from string, Msg string) (message *Message) {

	message = &Message{
		Type: MessageTypeText,
		From: from,
		Msg:  Msg,
	}

	return
}
