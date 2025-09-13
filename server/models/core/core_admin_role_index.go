package core

import (
	"server/models/common"
	"time"
)

type CoreAdminRoleIndex struct {
	ID          common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleID      common.MyID `gorm:"not null;default:0" json:"role_id"`
	AdminID     common.MyID `gorm:"not null;default:0" json:"admin_id"`
	CreatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
}

func (CoreAdminRoleIndex) TableName() string {
	return "core_admin_role_index"
}
