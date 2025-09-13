package cms

import (
	"encoding/json"
	"server/models/common"
	"time"
)

type CmsDocumentArchive struct {
	DocumentID     common.MyID     `gorm:"primaryKey" json:"document_id"`
	Cont           string          `gorm:"type:mediumtext" json:"cont"`
	DownloadFiles  json.RawMessage `gorm:"type:json" json:"download_files"`
	SeoTitle       string          `gorm:"type:varchar(255)" json:"seo_title"`
	SeoKeyword     string          `gorm:"type:varchar(255)" json:"seo_keyword"`
	SeoDescription string          `gorm:"type:varchar(255)" json:"seo_description"`
	DeletedTime    *time.Time      `json:"deletedTime"`
}

func (CmsDocumentArchive) TableName() string {
	return "cms_document_archive"
}
