package redis_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feitianlove/wechat/logger"
	"github.com/feitianlove/wechat/models"
	"github.com/feitianlove/wechat/redis"
	"github.com/sirupsen/logrus"
)

const (
	userOnlinePrefix    = "acc:user:online:" // 用户在线状态
	userOnlineCacheTime = 24 * 60 * 60
)

func getUserOnlineKey(userKey string) (key string) {
	key = fmt.Sprintf("%s%s", userOnlinePrefix, userKey)
	return
}

// 设置用户在线
func SetUserOnlineInfo(userKey string, userOnline *models.UserOnline) error {
	redisClient := redis.GetRedisClient()
	key := getUserOnlineKey(userKey)
	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		logger.AccessLog.WithFields(logrus.Fields{
			"seq": userKey,
			"Err": err,
		}).Error("SetUserOnlineInfo error")
		return err
	}
	_, err = redisClient.Do(context.Background(), "setEx", key, userOnlineCacheTime, string(valueByte)).Result()
	if err != nil {
		logger.AccessLog.WithFields(logrus.Fields{
			"seq":   userKey,
			"key":   key,
			"value": string(valueByte),
			"Err":   err,
		}).Error("SetUserOnlineInfo error")
		return err
	}
	logger.AccessLog.WithFields(logrus.Fields{
		"seq":   userKey,
		"key":   key,
		"value": string(valueByte),
	}).Info("SetUserOnlineInfo success")
	return err
}
