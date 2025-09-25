package cms

import (
	"server/models/common"
	"time"
)

type CmsDocument struct {
	ID                common.MyID      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title             string           `gorm:"size:200;not null;default:''" json:"title"`
	CategoryID        common.MyID      `gorm:"not null;default:0" json:"category_Id"`
	AssociatedPlaceID common.MyID      `gorm:"not null;default:0" json:"associated_PlaceId"`
	UserID            common.MyID      `gorm:"not null;default:0" json:"user_Id"`
	AdminID           common.MyID      `gorm:"not null;default:0" json:"admin_Id"`
	ContTpl           string           `gorm:"size:200;not null;default:''" json:"cont_Tpl"`
	Keyword           string           `gorm:"size:200;not null;default:''" json:"keyword"`
	Description       string           `gorm:"size:200;not null;default:''" json:"description"`
	Thumb             string           `gorm:"size:200;not null;default:''" json:"thumb"`
	State             common.MyState   `gorm:"not null;default:0" json:"state"`
	LinkType          common.MyType    `gorm:"not null;default:0" json:"link_type"`
	DocumentType      common.MyType    `gorm:"not null;default:0" json:"document_type"`
	VideoDuration     int              `gorm:"not null;default:0" json:"video_duration"`
	SendTime          *time.Time       `json:"send_time"`
	Author            string           `gorm:"size:200;not null;default:''" json:"author"`
	Source            string           `gorm:"size:200;not null;default:''" json:"source"`
	ReadNum           common.MyNumber  `gorm:"not null;default:0" json:"read_num"`
	LikeNum           common.MyNumber  `gorm:"not null;default:0" json:"like_num"`
	SortNum           common.MySortNum `gorm:"not null;default:0" json:"sort_num"`
	CreatedTime       time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime       time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime       *time.Time       `json:"deleted_time"`
	Code              string           `gorm:"size:512" json:"code"`
}

func (CmsDocument) TableName() string {
	return "cms_document"
}
