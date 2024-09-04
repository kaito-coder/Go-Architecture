package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID int64 `gorm:"column:id; type:int; not null, auto_increment"`
	RoleName string `gorm:"column:role_name; type:varchar(255)"`
}

func (r *Role) TableName() string {
	return "db_role"
}