package cms

import (
	"time"
)
type CmsRecommendIndex struct {
    ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
    RecommendID  int       `gorm:"not null;default:0" json:"recommendId"`
    Title        string    `gorm:"size:200;not null;default:''" json:"title"`
    Thumb        string    `gorm:"size:200;not null;default:''" json:"thumb"`
    Link         string    `gorm:"size:200;not null;default:''" json:"link"`
    State        int8      `gorm:"not null;default:0" json:"state"`
    ProductID    int       `gorm:"not null;default:0" json:"productId"`
    DocumentID   int       `gorm:"not null;default:0" json:"documentId"`
    SortNum      int       `gorm:"not null;default:0" json:"sortNum"`
    CreatedTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (CmsRecommendIndex) TableName() string {
	return "cms_recommend_index"
}