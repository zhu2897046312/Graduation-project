package constant

import (
    "encoding/json"
    "errors"
)

// SpProductStatus 表示商品状态的枚举类型
type SpProductStatus int

const (
    // SP_PRODUCT_ABLE 表示上线状态
    SP_PRODUCT_ABLE SpProductStatus = iota + 1
    
    // SP_PRODUCT_DISABLE 表示下线状态
    SP_PRODUCT_DISABLE
    
    // SP_PRODUCT_PENDING 表示待审核状态
    SP_PRODUCT_PENDING
    
    // SP_PRODUCT_FAILED 表示审核不通过状态
    SP_PRODUCT_FAILED
)

// SpProductStatusDesc 是商品状态到描述文本的映射
var SpProductStatusDesc = map[SpProductStatus]string{
    SP_PRODUCT_ABLE:    "状态:上线:#108ee9",
    SP_PRODUCT_DISABLE: "状态:下线:#108ee9",
    SP_PRODUCT_PENDING: "状态:待审核:#108ee9",
    SP_PRODUCT_FAILED:  "状态:审核不通过:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s SpProductStatus) String() string {
    if desc, ok := SpProductStatusDesc[s]; ok {
        return desc
    }
    return "未知状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s SpProductStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *SpProductStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpProductStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的商品状态")
}

// SpProductOpenSku 表示商品是否开启SKU的枚举类型
type SpProductOpenSku int

const (
    // SP_PRODUCT_OPEN_SKU_YES 表示开启SKU
    SP_PRODUCT_OPEN_SKU_YES SpProductOpenSku = iota + 1
    
    // SP_PRODUCT_OPEN_SKU_NO 表示不开启SKU
    SP_PRODUCT_OPEN_SKU_NO
)

// SpProductOpenSkuDesc 是商品SKU开启状态到描述文本的映射
var SpProductOpenSkuDesc = map[SpProductOpenSku]string{
    SP_PRODUCT_OPEN_SKU_YES: "是否开启SKU:是:#108ee9",
    SP_PRODUCT_OPEN_SKU_NO:  "是否开启SKU:否:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s SpProductOpenSku) String() string {
    if desc, ok := SpProductOpenSkuDesc[s]; ok {
        return desc
    }
    return "未知SKU状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s SpProductOpenSku) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *SpProductOpenSku) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpProductOpenSkuDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的商品SKU开启状态")
}

// SpProductSkuStatus 表示SKU状态的枚举类型
type SpProductSkuStatus int

const (
    // SP_PRODUCT_SKU_ABLE 表示正常状态
    SP_PRODUCT_SKU_ABLE SpProductSkuStatus = iota + 1
    
    // SP_PRODUCT_SKU_DISABLE 表示停用状态
    SP_PRODUCT_SKU_DISABLE
)

// SpProductSkuStatusDesc 是SKU状态到描述文本的映射
var SpProductSkuStatusDesc = map[SpProductSkuStatus]string{
    SP_PRODUCT_SKU_ABLE:   "SKU状态:正常:#108ee9",
    SP_PRODUCT_SKU_DISABLE: "SKU状态:停用:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s SpProductSkuStatus) String() string {
    if desc, ok := SpProductSkuStatusDesc[s]; ok {
        return desc
    }
    return "未知SKU状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s SpProductSkuStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *SpProductSkuStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpProductSkuStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的SKU状态")
}

// SpProductSkuDefaultShow 表示SKU是否默认显示的枚举类型
type SpProductSkuDefaultShow int

const (
    // SP_PRODUCT_SKU_DEFAULT_SHOW_YES 表示默认显示
    SP_PRODUCT_SKU_DEFAULT_SHOW_YES SpProductSkuDefaultShow = iota + 1
    
    // SP_PRODUCT_SKU_DEFAULT_SHOW_NO 表示不默认显示
    SP_PRODUCT_SKU_DEFAULT_SHOW_NO
)

// SpProductSkuDefaultShowDesc 是SKU默认显示状态到描述文本的映射
var SpProductSkuDefaultShowDesc = map[SpProductSkuDefaultShow]string{
    SP_PRODUCT_SKU_DEFAULT_SHOW_YES: "SKU是否默认显示:是:#108ee9",
    SP_PRODUCT_SKU_DEFAULT_SHOW_NO:  "SKU是否默认显示:否:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s SpProductSkuDefaultShow) String() string {
    if desc, ok := SpProductSkuDefaultShowDesc[s]; ok {
        return desc
    }
    return "未知SKU显示状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s SpProductSkuDefaultShow) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *SpProductSkuDefaultShow) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpProductSkuDefaultShowDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的SKU显示状态")
}

// SpProductHot 表示商品是否热门的枚举类型
type SpProductHot int

const (
    // SP_PRODUCT_HOT_YES 表示热门商品
    SP_PRODUCT_HOT_YES SpProductHot = iota + 1
    
    // SP_PRODUCT_HOT_NO 表示非热门商品
    SP_PRODUCT_HOT_NO
)

// SpProductHotDesc 是商品热门状态到描述文本的映射
var SpProductHotDesc = map[SpProductHot]string{
    SP_PRODUCT_HOT_YES: "商品是否热门:是:#108ee9",
    SP_PRODUCT_HOT_NO:  "商品是否热门:否:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (h SpProductHot) String() string {
    if desc, ok := SpProductHotDesc[h]; ok {
        return desc
    }
    return "未知热门状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (h SpProductHot) MarshalJSON() ([]byte, error) {
    return json.Marshal(h.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (h *SpProductHot) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpProductHotDesc {
        if v == desc {
            *h = k
            return nil
        }
    }
    return errors.New("未知的商品热门状态")
}