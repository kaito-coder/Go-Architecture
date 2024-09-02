package global

import (
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/logger"
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)