package core

import (
	"time"
)
type CoreAdminRoleIndex struct {
    ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    RoleID      int64     `gorm:"not null;default:0" json:"roleId"`
    AdminID     int64     `gorm:"not null;default:0" json:"adminId"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}