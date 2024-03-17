package svc

import (
	"cloud_disk3/core/internal/config"
	"cloud_disk3/core/internal/middleware"
	"cloud_disk3/core/models"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Engine: models.InitDB(c.Mysql.Datasource),
		RDB:    models.InitRedisConnection(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
