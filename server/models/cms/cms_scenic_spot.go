package cms

import (
	"time"
)
type CmsScenicSpot struct {
    ID               int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    AssociatedPlaceID int64     `gorm:"not null;default:0" json:"associatedPlaceId"`
    Title            string    `gorm:"size:255;not null;default:''" json:"title"`
    State            int8      `gorm:"not null;default:0" json:"state"`
    Code             string    `gorm:"size:255;not null;default:''" json:"code"`
    FullPinyin       string    `gorm:"size:255" json:"fullPinyin"`
    InitialPinyin    string    `gorm:"size:255" json:"initialPinyin"`
    ThumbImg         string    `gorm:"size:200;not null;default:''" json:"thumbImg"`
    ThumbVideo       string    `gorm:"size:200;not null;default:''" json:"thumbVideo"`
    DocumentTotal    int       `gorm:"not null;default:0" json:"documentTotal"`
    ReadNum          int       `gorm:"not null;default:0" json:"readNum"`
    Description      string    `gorm:"type:text" json:"description"`
    Content          string    `gorm:"type:text" json:"content"`
    Score            int       `gorm:"not null;default:0" json:"score"`
    CreatedTime      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}