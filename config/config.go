package config

import (
	"github.com/BurntSushi/toml"
	"github.com/feitianlove/golib/common/logger"
	golibVar "github.com/feitianlove/golib/config"
	"path/filepath"
)

type Config struct {
	App       *AppConf
	Redis     *RedisConf
	CtrlLog   *logger.LogConf
	AccessLog *logger.LogConf
}

type AppConf struct {
	HttpPort      int32
	WebSocketPort int32
	RpcPort       int32
	HttpUrl       string
	WebSocketUrl  string
}
type RedisConf struct {
	Addr        string
	Password    string
	DB          int
	PoolSize    int
	MinIdleConn int
}

func defaultConfig() *Config {
	return &Config{
		App: &AppConf{
			HttpPort:      0,
			WebSocketPort: 0,
			RpcPort:       0,
			HttpUrl:       "",
			WebSocketUrl:  "",
		},
		Redis: &RedisConf{
			Addr:        "",
			Password:    "",
			DB:          0,
			PoolSize:    0,
			MinIdleConn: 0,
		},
		AccessLog: &logger.LogConf{
			LogLevel:      "info",
			LogPath:       filepath.Join(golibVar.LaunchDir, "../log/access.log"),
			LogReserveDay: 90,
			ReportCaller:  false,
		},
		CtrlLog: &logger.LogConf{
			LogLevel:      "info",
			LogPath:       filepath.Join(golibVar.LaunchDir, "../log/ctrl.log"),
			LogReserveDay: 90,
			ReportCaller:  false,
		},
	}
}

func InitConfig(filePath string) (*Config, error) {
	cfg := defaultConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
