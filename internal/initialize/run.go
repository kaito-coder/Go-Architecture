package initialize

import (
	"fmt"

	"github.com/kaito-coder/go-ecommerce-architecture/global"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Printf("User: %s, Host: %s, Port: %s, Password: %s\n", m.User, m.Host, m.Port, m.Password)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8080")
}
