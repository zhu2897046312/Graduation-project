package constant

import (
    "encoding/json"
    "errors"
)

// SpCategoryStatus 表示商品类目状态的枚举类型
type SpCategoryStatus int

const (
    // SP_CATEGORY_ABLE 表示正常状态
    SP_CATEGORY_ABLE SpCategoryStatus = iota + 1
    
    // SP_CATEGORY_DISABLE 表示停用状态
    SP_CATEGORY_DISABLE
)

// SpCategoryStatusDesc 是商品类目状态到描述文本的映射
var SpCategoryStatusDesc = map[SpCategoryStatus]string{
    SP_CATEGORY_ABLE:   "商品类目状态:正常:#108ee9",
    SP_CATEGORY_DISABLE: "商品类目状态:停用:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s SpCategoryStatus) String() string {
    if desc, ok := SpCategoryStatusDesc[s]; ok {
        return desc
    }
    return "未知状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s SpCategoryStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *SpCategoryStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpCategoryStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的商品类目状态")
}