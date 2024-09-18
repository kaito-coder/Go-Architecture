package initialize

import (
	"fmt"
	"time"

	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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
	genTableNames()
	//MigrateTables()
}
func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	CheckErrorPanics(err, "Failed to set pool")
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func genTableNames() {
	g := gen.NewGenerator(gen.Config{
    OutPath: "./internal/model",
    Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
  })

  // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
  g.UseDB(global.Mdb) // reuse your gorm db
  g.GenerateModel("crm_user")

  // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`

  // Generate the code
  g.Execute()
}

func MigrateTables() {
	err := global.Mdb.AutoMigrate(&model.CrmUserV2{})
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
		panic(err)
	}
}
