package cms

import (
	"time"
)
type CmsFile struct {
    ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    FileName    string    `gorm:"size:200;not null;default:''" json:"fileName"`
    FilePath    string    `gorm:"size:200;not null;default:''" json:"filePath"`
    FileSize    int       `gorm:"not null;default:0" json:"fileSize"`
    FileType    string    `gorm:"size:50;not null;default:''" json:"fileType"`
    FileMd5     string    `gorm:"size:200;not null;default:''" json:"fileMd5"`
    FileExt     string    `gorm:"size:50;not null;default:''" json:"fileExt"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}