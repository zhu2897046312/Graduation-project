package constant

import (
    "encoding/json"
    "errors"
)

// CmsRecommendStatus 表示CMS推荐状态的枚举类型
type CmsRecommendStatus int

const (
    // CMS_RECOMMEND_PUBLISHED 表示已发布状态
    CMS_RECOMMEND_PUBLISHED CmsRecommendStatus = iota + 1
    
    // CMS_RECOMMEND_UN_PUBLISHED 表示未发布状态
    CMS_RECOMMEND_UN_PUBLISHED
)

// CmsRecommendStatusDesc 是CMS推荐状态到描述文本的映射
var CmsRecommendStatusDesc = map[CmsRecommendStatus]string{
    CMS_RECOMMEND_PUBLISHED:    "状态:已发布:#108ee9",
    CMS_RECOMMEND_UN_PUBLISHED: "状态:未发布:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s CmsRecommendStatus) String() string {
    if desc, ok := CmsRecommendStatusDesc[s]; ok {
        return desc
    }
    return "未知状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s CmsRecommendStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *CmsRecommendStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range CmsRecommendStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的CMS推荐状态")
}