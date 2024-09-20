package global

import (
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/logger"
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb *redis.Client
	Mdb *gorm.DB
)