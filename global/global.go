package global

import (
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/logger"
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
)