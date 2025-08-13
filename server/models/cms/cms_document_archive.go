package cms

import (
	"encoding/json"
)
type CmsDocumentArchive struct {
    DocumentID    int64           `gorm:"primaryKey" json:"documentId"`
    Cont          string          `gorm:"type:mediumtext" json:"cont"`
    DownloadFiles json.RawMessage `gorm:"type:json" json:"downloadFiles"`
}