package initialize

import (
	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/logger"
)

func InitLogger()  {
	global.Logger = logger.NewLogger(global.Config.Logger)
}