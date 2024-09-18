package model

const TableNameCrmUser = "crm_user"

// CrmUser Account
type CrmUser struct {
	UserID           int32  `gorm:"column:user_id;primaryKey;autoIncrement:true;comment:Account ID" json:"user_id"` // Account ID
	UsrEmail         string `gorm:"column:usr_email;not null;);comment:Email" json:"usr_email"`
	UsrPassword      string `gorm:"column:usr_password;not null;comment:Password" json:"usr_password"`                         // Password
	UsrUsername      string `gorm:"column:usr_username;not null;comment:Username" json:"usr_username"`                         // Username
	UsrPhone         string `gorm:"column:usr_phone;not null;comment:Phone Number" json:"usr_phone"`                           // Phone Number
	UsrCreateAt      int32  `gorm:"column:usr_create_at;not null;comment:Creation Time IP" json:"usr_create_at"`               // Creation Time IP
	UsrCreatedIPAt   int32  `gorm:"column:usr_created_ip_at;not null;comment:Creation Time" json:"usr_created_ip_at"`          // Creation Time
	UsrUpdatedAt     int32  `gorm:"column:usr_updated_at;not null;comment:Update Time" json:"usr_updated_at"`                  // Update Time
	UsrLastLoginAt   int32  `gorm:"column:usr_last_login_at;not null;comment:Last Login Time" json:"usr_last_login_at"`        // Last Login Time
	UsrLastLoginIPAt int32  `gorm:"column:usr_last_login_ip_at;not null;comment:Last Login IP" json:"usr_last_login_ip_at"`    // Last Login IP
	UsrStatus        bool   `gorm:"column:usr_status;not null;comment:Status 1:enable 0:disable -1:deleted" json:"usr_status"` // Status 1:enable 0:disable -1:deleted
}

// TableName CrmUser's table name
func (*CrmUser) TableName() string {
	return TableNameCrmUser
}
