package initialize

import (
	"fmt"
	"time"

	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CheckErrorPanics(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	CheckErrorPanics(err, "Failed to connect to database")
	global.Logger.Info("Database init successful")
	global.Mdb = db
	SetPool()
	MigrateTables()
}
func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	CheckErrorPanics(err, "Failed to set pool")
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func MigrateTables() {
	err := global.Mdb.AutoMigrate(&po.User{},
		&po.Role{})
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
		panic(err)
	}
}
