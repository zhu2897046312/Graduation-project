package cms

import (
	"encoding/json"
	"time"
)

type CmsDocumentArchive struct {
	DocumentID     int64           `gorm:"primaryKey" json:"documentId"`
	Cont           string          `gorm:"type:mediumtext" json:"cont"`
	DownloadFiles  json.RawMessage `gorm:"type:json" json:"downloadFiles"`
	SeoTitle       string          `gorm:"type:varchar(255)" json:"seo_title"`
	SeoKeyword     string          `gorm:"type:varchar(255)" json:"seo_keyword"`
	SeoDescription string          `gorm:"type:varchar(255)" json:"seo_description"`
	DeletedTime       *time.Time `json:"deletedTime"`
}

func (CmsDocumentArchive) TableName() string {
	return "cms_document_archive"
}
