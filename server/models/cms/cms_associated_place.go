package cms

import (
	"time"
)

// CmsAssociatedPlace 地点信息表模型
type CmsAssociatedPlace struct {
	ID            int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                       // 主键ID
	Title         string    `gorm:"column:title;not null;default:''" json:"title"`                      // 地点名称
	State         int8      `gorm:"column:state;not null;default:0" json:"state"`                       // 状态:1=已发布;2=未发布
	Code          string    `gorm:"column:code;not null;default:''" json:"code"`                        // 地点代码
	FullPinyin    string    `gorm:"column:full_pinyin" json:"fullPinyin"`                              // 地点名称全拼音
	InitialPinyin string    `gorm:"column:initial_pinyin" json:"initialPinyin"`                        // 地点名称首字母拼音
	ThumbImg      string    `gorm:"column:thumb_img;not null;default:''" json:"thumbImg"`               // 缩略图
	ThumbVideo    string    `gorm:"column:thumb_video;not null;default:''" json:"thumbVideo"`           // 封面视频
	Description   string    `gorm:"column:description;type:text" json:"description"`                    // 地点简介
	Content       string    `gorm:"column:content;type:text" json:"content"`                            // 地点内容
	Score         int       `gorm:"column:score;not null;default:0" json:"score"`                       // 评分
	CreatedTime   time.Time `gorm:"column:created_time;default:CURRENT_TIMESTAMP" json:"createdTime"`   // 创建时间
	UpdatedTime   time.Time `gorm:"column:updated_time;default:CURRENT_TIMESTAMP" json:"updatedTime"`   // 更新时间
}

// TableName 设置表名
func (CmsAssociatedPlace) TableName() string {
	return "cms_associated_place"
}