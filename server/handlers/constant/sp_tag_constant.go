package constant

// SpTagStatus 表示标签状态的枚举类型
type SpTagStatus int

const (
    // SP_TAG_PUBLISHED 表示已发布状态
    SP_TAG_PUBLISHED SpTagStatus = iota + 1
    
    // SP_TAG_UN_PUBLISHED 表示未发布状态
    SP_TAG_UN_PUBLISHED
)
