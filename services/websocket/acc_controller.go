package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/feitianlove/golib/common/ecode"
	"github.com/feitianlove/wechat/logger"
	"github.com/feitianlove/wechat/models"
	"github.com/feitianlove/wechat/services/redis_cache"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// 用户登录
type login struct {
	AppId  uint32
	UserId string
	Client *Client
}
type DisposeFunc func(client *Client, seq string, message []byte) (code ecode.ECode, msg string, data interface{})

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutex sync.RWMutex
	serverIp        string
	serverPort      string
)

// 注册
func Register(key string, value DisposeFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value
	data, _ := json.Marshal(handlers)
	logger.CtrlLog.WithFields(logrus.Fields{
		"websocket handlers": string(data),
	}).Info("Register web socket router")
}

// 用户登陆
func LoginController(client *Client, seq string, message []byte) (code ecode.ECode, msg string, data interface{}) {
	code = ecode.OK
	currentTime := time.Now().Unix()
	request := &models.Login{}
	if err := json.Unmarshal(message, request); err != nil {
		code = ecode.ParamError
		logger.AccessLog.WithFields(logrus.Fields{
			"seq": seq,
			"err": err,
			"Err": string(message),
		}).Error("Login error")
		return
	}
	logger.AccessLog.WithFields(logrus.Fields{
		"seq":     seq,
		"request": request.ServiceToken,
	}).Info("Login success")

	// TODO::进行用户权限认证，一般是客户端传入TOKEN，然后检验TOKEN是否合法，通过TOKEN解析出来用户ID
	// 本项目只是演示，所以直接过去客户端传入的用户ID
	if request.UserId == "" || len(request.UserId) >= 20 {
		code = ecode.ErrError
		logger.AccessLog.WithFields(logrus.Fields{
			"seq":    seq,
			"UserId": request.UserId,
			"Err":    "用户登录非法的用户",
		}).Error("Login error")
		return
	}

	// 判断平台
	if !InAppIds(request.AppId) {
		code = ecode.ErrError
		logger.AccessLog.WithFields(logrus.Fields{
			"seq":    seq,
			"UserId": request.AppId,
			"Err":    "非法登陆平台",
		}).Error("Login error")
		return
	}
	//判断用户是否登陆
	if client.IsLogin() {
		fmt.Println("用户登录 用户已经登录", client.AppId, client.UserId, seq)
		code = ecode.ErrError
		logger.AccessLog.WithFields(logrus.Fields{
			"seq":    seq,
			"UserId": request.UserId,
			"Err":    "用户已经登陆",
		}).Error("Login error")
		return
	}
	client.Login(request.AppId, request.UserId, uint64(currentTime))
	//存储用户登陆信息
	userOnline := models.UserLogin(serverIp, serverPort, request.AppId, request.UserId, client.Addr, uint64(currentTime))

	err := redis_cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = ecode.ErrError
		logger.AccessLog.WithFields(logrus.Fields{
			"seq": seq,
			"Err": err,
		}).Info("Login error")
		return
	}
	login := &login{
		AppId:  request.AppId,
		UserId: request.UserId,
		Client: client,
	}
	clientManager.Login <- login

	fmt.Println("用户登录 成功", seq, client.Addr, request.UserId)

	logger.AccessLog.WithFields(logrus.Fields{
		"seq":    seq,
		"UserId": request.UserId,
	}).Info("Login success")
	return
}

func IsLocal(server *models.Server) (isLocal bool) {
	if server.Ip == serverIp && server.Port == serverPort {
		isLocal = true
	}
	return
}
