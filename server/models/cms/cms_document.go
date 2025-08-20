package cms

import (
	"time"
)
type CmsDocument struct {
    ID                int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Title             string    `gorm:"size:200;not null;default:''" json:"title"`
    CategoryID        int64     `gorm:"not null;default:0" json:"categoryId"`
    AssociatedPlaceID int64     `gorm:"not null;default:0" json:"associatedPlaceId"`
    ContTpl           string    `gorm:"size:200;not null;default:''" json:"contTpl"`
    Keyword           string    `gorm:"size:200;not null;default:''" json:"keyword"`
    Description       string    `gorm:"size:200;not null;default:''" json:"description"`
    Thumb             string    `gorm:"size:200;not null;default:''" json:"thumb"`
    State             int8      `gorm:"not null;default:0" json:"state"`
    LinkType          int8      `gorm:"not null;default:0" json:"linkType"`
    DocumentType      int8      `gorm:"not null;default:0" json:"documentType"`
    VideoDuration     int       `gorm:"not null;default:0" json:"videoDuration"`
    SendTime          *time.Time `json:"sendTime"`
    Author            string    `gorm:"size:200;not null;default:''" json:"author"`
    Source            string    `gorm:"size:200;not null;default:''" json:"source"`
    UserID            int64     `gorm:"not null;default:0" json:"userId"`
    AdminID           int64     `gorm:"not null;default:0" json:"adminId"`
    ReadNum           int       `gorm:"not null;default:0" json:"readNum"`
    LikeNum           int       `gorm:"not null;default:0" json:"likeNum"`
    SortNum           int       `gorm:"not null;default:0" json:"sortNum"`
    CreatedTime       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime       *time.Time `json:"deletedTime"`
    Code              string    `gorm:"size:512" json:"code"`
}

func (CmsDocument) TableName() string {
	return "cms_document"
}