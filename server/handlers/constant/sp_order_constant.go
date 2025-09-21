package constant

import (
    "encoding/json"
    "errors"
)


// OrderStatus 表示订单状态的枚举类型
// 使用 int 类型作为基础类型，便于存储和比较
type OrderStatus int

// 订单状态枚举值定义
const (
    // PENDING_PAYMENT 表示待付款状态
    PENDING_PAYMENT OrderStatus = iota + 1 // 从1开始

    // ON_DELIVERY 表示待发货状态
    ON_DELIVERY

    // SHIPPED 表示已发货状态
    SHIPPED

    // COMPLETED 表示已完成状态
    COMPLETED

    // CLOSED 表示已关闭状态
    CLOSED

    // INVALID 表示无效订单状态
    INVALID

    // PART_REFUND 表示部分退款状态
    PART_REFUND

    // ALL_REFUND 表示全部退款状态
    ALL_REFUND
)

// OrderStatusDesc 是订单状态到描述文本的映射
// 用于将枚举值转换为可读性更好的文本
var OrderStatusDesc = map[OrderStatus]string{
    PENDING_PAYMENT: "待付款",
    ON_DELIVERY:     "待发货",
    SHIPPED:        "已发货",
    COMPLETED:      "已完成",
    CLOSED:         "已关闭",
    INVALID:        "无效订单",
    PART_REFUND:    "部分退款",
    ALL_REFUND:     "全部退款",
}

// String 实现 fmt.Stringer 接口
// 返回订单状态的描述文本
func (s OrderStatus) String() string {
    if desc, ok := OrderStatusDesc[s]; ok {
        return desc
    }
    return "未知状态"
}

// MarshalJSON 实现 json.Marshaler 接口
// 将订单状态序列化为描述文本
func (s OrderStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
// 将描述文本反序列化为订单状态
func (s *OrderStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range OrderStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的订单状态")
}

// SpOrderSourceType 表示订单来源类型的枚举类型
type SpOrderSourceType int

const (
    // SP_ORDER_SOURCE_PC 表示PC订单
    SP_ORDER_SOURCE_PC SpOrderSourceType = iota + 1
    
    // SP_ORDER_SOURCE_MOBILE 表示移动端订单
    SP_ORDER_SOURCE_MOBILE
)

// SpOrderSourceTypeDesc 是订单来源类型到描述文本的映射
var SpOrderSourceTypeDesc = map[SpOrderSourceType]string{
    SP_ORDER_SOURCE_PC:     "订单来源:PC订单:#108ee9",
    SP_ORDER_SOURCE_MOBILE: "订单来源:移动端订单:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (t SpOrderSourceType) String() string {
    if desc, ok := SpOrderSourceTypeDesc[t]; ok {
        return desc
    }
    return "未知来源类型"
}

// MarshalJSON 实现 json.Marshaler 接口
func (t SpOrderSourceType) MarshalJSON() ([]byte, error) {
    return json.Marshal(t.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (t *SpOrderSourceType) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpOrderSourceTypeDesc {
        if v == desc {
            *t = k
            return nil
        }
    }
    return errors.New("未知的订单来源类型")
}

// SpOrderPayType 表示订单支付类型的枚举类型
type SpOrderPayType int

const (
    // SP_ORDER_PAY_ON_DELIVERY 表示货到付款
    SP_ORDER_PAY_ON_DELIVERY SpOrderPayType = iota + 1
)

// SpOrderPayTypeDesc 是订单支付类型到描述文本的映射
var SpOrderPayTypeDesc = map[SpOrderPayType]string{
    SP_ORDER_PAY_ON_DELIVERY: "支付方式:货到付款:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (t SpOrderPayType) String() string {
    if desc, ok := SpOrderPayTypeDesc[t]; ok {
        return desc
    }
    return "未知支付类型"
}

// MarshalJSON 实现 json.Marshaler 接口
func (t SpOrderPayType) MarshalJSON() ([]byte, error) {
    return json.Marshal(t.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (t *SpOrderPayType) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpOrderPayTypeDesc {
        if v == desc {
            *t = k
            return nil
        }
    }
    return errors.New("未知的订单支付类型")
}

// SpOrderDisputeStatus 表示订单争议状态的枚举类型
type SpOrderDisputeStatus int

const (
    // SP_ORDER_NO_DISPUTE 表示无争议
    SP_ORDER_NO_DISPUTE SpOrderDisputeStatus = iota
    
    // SP_ORDER_BUYER_INITIATED 表示买家发起争议
    SP_ORDER_BUYER_INITIATED
    
    // SP_ORDER_SELLER_PROCESSING 表示卖家处理中
    SP_ORDER_SELLER_PROCESSING
    
    // SP_ORDER_PLATFORM_ARBITRATION 表示平台仲裁中
    SP_ORDER_PLATFORM_ARBITRATION
    
    // SP_ORDER_RESOLVED 表示已解决
    SP_ORDER_RESOLVED
)

// SpOrderDisputeStatusDesc 是订单争议状态到描述文本的映射
var SpOrderDisputeStatusDesc = map[SpOrderDisputeStatus]string{
    SP_ORDER_NO_DISPUTE:           "无争议",
    SP_ORDER_BUYER_INITIATED:      "买家发起争议",
    SP_ORDER_SELLER_PROCESSING:    "卖家处理中",
    SP_ORDER_PLATFORM_ARBITRATION: "平台仲裁中",
    SP_ORDER_RESOLVED:             "已解决",
}

// String 实现 fmt.Stringer 接口
func (s SpOrderDisputeStatus) String() string {
    if desc, ok := SpOrderDisputeStatusDesc[s]; ok {
        return desc
    }
    return "未知争议状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s SpOrderDisputeStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *SpOrderDisputeStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpOrderDisputeStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的订单争议状态")
}

// SpOrderState 表示订单状态的枚举类型
type SpOrderState int

const (
    // SP_ORDER_PENDING_PAYMENT 表示待付款状态
    SP_ORDER_PENDING_PAYMENT SpOrderState = iota + 1
    
    // SP_ORDER_ON_DELIVERY 表示待发货状态
    SP_ORDER_ON_DELIVERY
    
    // SP_ORDER_SHIPPED 表示已发货状态
    SP_ORDER_SHIPPED
    
    // SP_ORDER_COMPLETED 表示已完成状态
    SP_ORDER_COMPLETED
    
    // SP_ORDER_CLOSED 表示已关闭状态
    SP_ORDER_CLOSED
    
    // SP_ORDER_INVALID 表示无效订单状态
    SP_ORDER_INVALID
    
    // SP_ORDER_PART_REFUND 表示部分退款状态
    SP_ORDER_PART_REFUND
    
    // SP_ORDER_ALL_REFUND 表示全部退款状态
    SP_ORDER_ALL_REFUND
)

// SpOrderStateDesc 是订单状态到描述文本的映射
var SpOrderStateDesc = map[SpOrderState]string{
    SP_ORDER_PENDING_PAYMENT: "待付款",
    SP_ORDER_ON_DELIVERY:     "待发货",
    SP_ORDER_SHIPPED:         "已发货",
    SP_ORDER_COMPLETED:       "已完成",
    SP_ORDER_CLOSED:          "已关闭",
    SP_ORDER_INVALID:         "无效订单",
    SP_ORDER_PART_REFUND:     "部分退款",
    SP_ORDER_ALL_REFUND:      "全部退款",
}

// String 实现 fmt.Stringer 接口
func (s SpOrderState) String() string {
    if desc, ok := SpOrderStateDesc[s]; ok {
        return desc
    }
    return "未知状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s SpOrderState) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *SpOrderState) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpOrderStateDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的订单状态")
}