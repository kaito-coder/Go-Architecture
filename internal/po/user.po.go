package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID uuid.UUID `gorm:"column:uuid; type:char(36); index:idx_uuid; not null"`
	Username string `gorm:"column:username; type:varchar(255)"`
	IsActive bool `gorm:"column:is_active; type:boolean; default:true"`
	Roles []Role `gorm:"many2many:user_roles;"`
}

func (u *User) TableName() string {
	return "db_users"
}