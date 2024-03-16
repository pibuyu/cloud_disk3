package svc

import (
	"cloud_disk3/core/internal/config"
	"cloud_disk3/core/models"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.InitDB(c.Mysql.Datasource),
		RDB:    models.InitRedisConnection(c),
	}
}
