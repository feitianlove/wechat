package logger

import (
	goliblogger "github.com/feitianlove/golib/common/logger"
	"github.com/feitianlove/wechat/config"
	"github.com/sirupsen/logrus"
)

var (
	CtrlLog   *logrus.Logger
	AccessLog *logrus.Logger
)

func init() {
	CtrlLog = goliblogger.NewLoggerInstance()
	AccessLog = goliblogger.NewLoggerInstance()
}

func InitCtrlLog(conf *goliblogger.LogConf) error {
	l, err := goliblogger.InitLogger(conf)
	if err != nil {
		return err
	}
	CtrlLog = l
	return nil
}

func InitAccessLog(conf *goliblogger.LogConf) error {
	l, err := goliblogger.InitLogger(conf)
	if err != nil {
		return err
	}
	AccessLog = l
	return nil
}

func InitLog(conf *config.Config) error {
	if err := InitCtrlLog(conf.CtrlLog); err != nil {
		return err
	}
	if err := InitAccessLog(conf.AccessLog); err != nil {
		return err
	}
	return nil
}
