package daemons

import (
	"github.com/kklab-com/gone-http/http/httpsession/redis"
	kkdaemon "github.com/kklab-com/goth-daemon"
	datastore "github.com/kklab-com/goth-kkdatastore"
	kklogger "github.com/kklab-com/goth-kklogger"
	"github.com/kklab-com/goth-scaffold/app/conf"
	redis2 "github.com/kklab-com/goth-scaffold/app/services/redis"
	"github.com/pkg/errors"
)

var DaemonSetupRedis = &SetupRedis{}

type SetupRedis struct {
	kkdaemon.DefaultDaemon
}

func (d *SetupRedis) Start() {
	redis.RedisName = conf.Config().DataStore.RedisName
	datastore.KKRedisIdleTimeout = 60000
	datastore.KKRedisMaxConnLifetime = 300000
	datastore.KKRedisMaxIdle = 10
	datastore.KKRedisMaxActive = 2000
	datastore.KKRedisWait = true

	if redis2.Master() == nil {
		panic(errors.Errorf("can't connect to master"))
	}

	if redis2.Slave() == nil {
		panic(errors.Errorf("can't connect to slave"))
	}

	if r := redis2.Master().Keys("*"); r.Error != nil {
		kklogger.ErrorJ("daemon.SetupRedis#Master", r.Error)
		panic(r.Error)
	}

	if r := redis2.Slave().Keys("*"); r.Error != nil {
		kklogger.ErrorJ("daemon.SetupRedis#Slave", r.Error)
		panic(r.Error)
	}
}
