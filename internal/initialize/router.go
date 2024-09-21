package initialize

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/routers"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/model"
)
var rebates = []model.Rebate{
	{
		ID:                  1,
		ProductManufacturerID: 1001,
		Name:                "Holiday Discount",
		Code:                "HOLIDAY2024",
		Description:         "Get 10% off during the holiday season!",
		StartDate:           time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
		EndDate:             time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
		MaxUses:             newInt(100),
		MixedAvailability:   false,
		State:               2,
		ActiveTime:          ptrToTime(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)),
		IsCombinable:        1,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	},
	{
		ID:                  2,
		ProductManufacturerID: 1002,
		Name:                "Spring Sale",
		Code:                "SPRING2024",
		Description:         "Enjoy 15% off this spring!",
		StartDate:           time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
		EndDate:             time.Date(2024, 4, 20, 23, 59, 59, 0, time.UTC),
		MaxUses:             newInt(200),
		MixedAvailability:   true,
		State:               2,
		ActiveTime:          ptrToTime(time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)),
		IsCombinable:        0,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	},
	// Other rebates...
}
func newInt(i int) *int {
	return &i
}

func ptrToTime(t time.Time) *time.Time {
	return &t
}

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev"{
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	//r.Use(middleware.limiter())
	//r.Use(middleware.Cors())
	//r.Use(middleware.Logger())
	//r.Use(middleware.Recovery())
	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}
	{
		adminRouter.InitAdminRouter(MainGroup)
		adminRouter.InitUserRouter(MainGroup)
	}
	r.GET("/rebates", func(c *gin.Context) {
		c.JSON(200,rebates)
	})
	return r
}

