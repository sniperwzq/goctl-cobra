package config

import (
    "github.com/gogf/gf/v2/container/gmap"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

const (
	defaultName = "gz.config"
)

var instances = gmap.NewStrAnyMap(true)

type Config struct {
	zrpc.RpcServerConf
}

func LoadCfg(path string) *Config {
	return instances.GetOrSetFuncLock(defaultName, func() interface{} {
		var c Config
		conf.MustLoad(path, &c)
		return &c
	}).(*Config)
}

func Cfg() *Config {
	return instances.Get(defaultName).(*Config)
}
