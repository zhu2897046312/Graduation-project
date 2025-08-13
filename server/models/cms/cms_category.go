package cms

import (
	"time"
)

type CmsCategory struct {
    ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Title        string    `gorm:"size:200;not null;default:''" json:"title"`
    Code         string    `gorm:"size:200;not null;default:''" json:"code"`
    Pid          int64     `gorm:"not null;default:0" json:"pid"`
    CategoryTpl  string    `gorm:"size:200;not null;default:''" json:"categoryTpl"`
    ContTpl      string    `gorm:"size:200;not null;default:''" json:"contTpl"`
    Keyword      string    `gorm:"size:200;not null;default:''" json:"keyword"`
    Description  string    `gorm:"size:200;not null;default:''" json:"description"`
    Thumb        string    `gorm:"size:200;not null;default:''" json:"thumb"`
    SortNum      int       `gorm:"not null;default:0" json:"sortNum"`
    CategoryType int8      `gorm:"not null;default:0" json:"categoryType"`
    CreatedTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}