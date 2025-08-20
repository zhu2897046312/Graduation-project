package cms

import (
	"time"
)
type CmsDocumentTag struct {
    ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    DocumentID  int64     `gorm:"not null;default:0" json:"documentId"`
    TagID       int64     `gorm:"not null;default:0" json:"tagId"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (CmsDocumentTag) TableName() string {
	return "cms_document_tag"
}