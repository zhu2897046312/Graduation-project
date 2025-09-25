package sp

import (
	"server/models/common"
	"time"
)

type SpCategory struct {
	ID             common.MyID      `gorm:"primaryKey;autoIncrement" json:"id"`
	Pid            common.MyID      `gorm:"not null;default:0" json:"pid"`
	Title          string           `gorm:"size:200;not null;default:''" json:"title"`
	Code           string           `gorm:"size:200;not null;default:''" json:"code"`
	State          common.MyState   `gorm:"not null;default:1" json:"state"`
	Icon           string           `gorm:"size:200;not null;default:''" json:"icon"`
	Picture        string           `gorm:"size:200;not null;default:''" json:"picture"`
	Description    string           `gorm:"size:200;not null;default:''" json:"description"`
	SortNum        common.MySortNum `gorm:"not null;default:0" json:"sort_num"`
	SeoTitle       string           `gorm:"size:200;not null;default:''" json:"seo_title"`
	SeoKeyword     string           `gorm:"size:200;not null;default:''" json:"seo_keyword"`
	SeoDescription string           `gorm:"size:200;not null;default:''" json:"seo_description"`
	CreatedTime    time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime    time.Time        `gorm:"autoUpdateTime default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime    *time.Time       `json:"deleted_time"`
	//Children      []SpCategory `gorm:"-" json:"children"` // 添加Children字段，gorm:"-"表示忽略数据库映射
}

func (SpCategory) TableName() string {
	return "sp_category"
}
