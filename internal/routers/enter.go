package routers

import (
	"github.com/kaito-coder/go-ecommerce-architecture/internal/routers/admin"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Admin admin.AdminRouterGroup
}
var RouterGroupApp = new(RouterGroup)