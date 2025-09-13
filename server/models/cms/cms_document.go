package cms

import (
	"server/models/common"
	"time"
)

type CmsDocument struct {
	ID                common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	Title             string      `gorm:"size:200;not null;default:''" json:"title"`
	CategoryID        common.MyID `gorm:"not null;default:0" json:"categoryId"`
	AssociatedPlaceID common.MyID `gorm:"not null;default:0" json:"associatedPlaceId"`
    UserID            common.MyID `gorm:"not null;default:0" json:"userId"`
	AdminID           common.MyID `gorm:"not null;default:0" json:"adminId"`
	ContTpl           string      `gorm:"size:200;not null;default:''" json:"contTpl"`
	Keyword           string      `gorm:"size:200;not null;default:''" json:"keyword"`
	Description       string      `gorm:"size:200;not null;default:''" json:"description"`
	Thumb             string      `gorm:"size:200;not null;default:''" json:"thumb"`
	State             int8        `gorm:"not null;default:0" json:"state"`
	LinkType          int8        `gorm:"not null;default:0" json:"linkType"`
	DocumentType      int8        `gorm:"not null;default:0" json:"documentType"`
	VideoDuration     int         `gorm:"not null;default:0" json:"videoDuration"`
	SendTime          *time.Time  `json:"sendTime"`
	Author            string      `gorm:"size:200;not null;default:''" json:"author"`
	Source            string      `gorm:"size:200;not null;default:''" json:"source"`
	ReadNum           int         `gorm:"not null;default:0" json:"readNum"`
	LikeNum           int         `gorm:"not null;default:0" json:"likeNum"`
	SortNum           int         `gorm:"not null;default:0" json:"sortNum"`
	CreatedTime       time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
	UpdatedTime       time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
	DeletedTime       *time.Time  `json:"deletedTime"`
	Code              string      `gorm:"size:512" json:"code"`
}

func (CmsDocument) TableName() string {
	return "cms_document"
}
