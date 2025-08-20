package cms

import (
	"time"
)

type CmsComment struct {
    ID                 int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    DocumentID         int64     `gorm:"not null;default:0" json:"documentId"`
    ReplyID           int64     `gorm:"not null;default:0" json:"replyId"`
    ReplyToReplyID    int64     `gorm:"not null;default:0" json:"replyToReplyId"`
    ReplyToUsername   string    `gorm:"size:200;not null;default:''" json:"replyToUsername"`
    ReplyToAvatar     string    `gorm:"size:200;not null;default:''" json:"replyToAvatar"`
    ReplyToUserid     int64     `gorm:"not null;default:0" json:"replyToUserid"`
    Text              string    `gorm:"type:text;not null" json:"text"`
    State             int8      `gorm:"not null;default:0" json:"state"`
    UserID            int64     `gorm:"not null;default:0" json:"userId"`
    Username          string    `gorm:"size:200;not null;default:''" json:"username"`
    Avatar            string    `gorm:"size:200;not null;default:''" json:"avatar"`
    IP                string    `gorm:"size:200;not null;default:''" json:"ip"`
    CommentReplyTotal int       `gorm:"not null;default:0" json:"commentReplyTotal"`
    LikeNum           int       `gorm:"not null;default:0" json:"likeNum"`
    BadNum            int       `gorm:"not null;default:0" json:"badNum"`
    Remark            string    `gorm:"size:200;not null;default:''" json:"remark"`
    CreatedTime       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime       *time.Time `json:"deletedTime"`
}

func (CmsComment) TableName() string {
	return "cms_comment"
}