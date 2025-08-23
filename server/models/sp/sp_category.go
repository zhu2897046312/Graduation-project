package sp

import (
	"time"
)

type SpCategory struct {
	ID             uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Pid            uint       `gorm:"not null;default:0" json:"pid"`
	Title          string     `gorm:"size:200;not null;default:''" json:"title"`
	Code           string     `gorm:"size:200;not null;default:''" json:"code"`
	State          uint8      `gorm:"not null;default:1" json:"state"`
	Icon           string     `gorm:"size:200;not null;default:''" json:"icon"`
	Picture        string     `gorm:"size:200;not null;default:''" json:"picture"`
	Description    string     `gorm:"size:200;not null;default:''" json:"description"`
	SortNum        uint16     `gorm:"not null;default:0" json:"sort_num"`
	SeoTitle       string     `gorm:"size:200;not null;default:''" json:"seo_title"`
	SeoKeyword     string     `gorm:"size:200;not null;default:''" json:"seo_keyword"`
	SeoDescription string     `gorm:"size:200;not null;default:''" json:"seo_description"`
	CreatedTime    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime    *time.Time `json:"deleted_time"`
	//Children      []SpCategory `gorm:"-" json:"children"` // 添加Children字段，gorm:"-"表示忽略数据库映射
}

func (SpCategory) TableName() string {
	return "sp_category"
}
