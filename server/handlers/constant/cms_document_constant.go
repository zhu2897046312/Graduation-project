package constant

import (
    "encoding/json"
    "errors"
)

// CmsDocumentStatus 表示CMS文档状态的枚举类型
type CmsDocumentStatus int

const (
    // CMS_DOCUMENT_PUBLISHED 表示已发布状态
    CMS_DOCUMENT_PUBLISHED CmsDocumentStatus = iota + 1
    
    // CMS_DOCUMENT_UN_PUBLISHED 表示未发布状态
    CMS_DOCUMENT_UN_PUBLISHED
)

// CmsDocumentStatusDesc 是CMS文档状态到描述文本的映射
var CmsDocumentStatusDesc = map[CmsDocumentStatus]string{
    CMS_DOCUMENT_PUBLISHED:    "状态:已发布:#108ee9",
    CMS_DOCUMENT_UN_PUBLISHED: "状态:未发布:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s CmsDocumentStatus) String() string {
    if desc, ok := CmsDocumentStatusDesc[s]; ok {
        return desc
    }
    return "未知状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s CmsDocumentStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *CmsDocumentStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range CmsDocumentStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的CMS文档状态")
}

// CmsDocumentLinkType 表示CMS文档链接类型的枚举类型
type CmsDocumentLinkType int

const (
    // CMS_DOCUMENT_INNER 表示内部文档类型
    CMS_DOCUMENT_INNER CmsDocumentLinkType = iota + 1
    
    // CMS_DOCUMENT_LINK 表示外部链接类型
    CMS_DOCUMENT_LINK
)

// CmsDocumentLinkTypeDesc 是CMS文档链接类型到描述文本的映射
var CmsDocumentLinkTypeDesc = map[CmsDocumentLinkType]string{
    CMS_DOCUMENT_INNER: "链接类型:内部文档:#108ee9",
    CMS_DOCUMENT_LINK:  "链接类型:外部链接:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (t CmsDocumentLinkType) String() string {
    if desc, ok := CmsDocumentLinkTypeDesc[t]; ok {
        return desc
    }
    return "未知类型"
}

// MarshalJSON 实现 json.Marshaler 接口
func (t CmsDocumentLinkType) MarshalJSON() ([]byte, error) {
    return json.Marshal(t.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (t *CmsDocumentLinkType) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range CmsDocumentLinkTypeDesc {
        if v == desc {
            *t = k
            return nil
        }
    }
    return errors.New("未知的CMS文档链接类型")
}