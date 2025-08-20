package cms

type CmsDocumentVideo struct {
    DocumentID int64  `gorm:"primaryKey" json:"documentId"`
    VideoPath  string `gorm:"size:200;not null;default:''" json:"videoPath"`
}

func (CmsDocumentVideo) TableName() string {
	return "cms_document_video"
}