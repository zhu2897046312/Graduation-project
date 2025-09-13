package cms

import (
	"server/models/common"
	"time"
)

type CmsDocumentTag struct {
	ID          common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	DocumentID  common.MyID `gorm:"not null;default:0" json:"document_id"`
	TagID       common.MyID `gorm:"not null;default:0" json:"tag_id"`
	CreatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
}

func (CmsDocumentTag) TableName() string {
	return "cms_document_tag"
}
