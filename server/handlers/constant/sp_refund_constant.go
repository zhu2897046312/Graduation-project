package constant

import (
    "encoding/json"
    "errors"
)

// SpRefundDisputeStatus 表示退款争议状态的枚举类型
type SpRefundDisputeStatus int

const (
    // SP_REFUND_PROCESSING 表示处理中状态
    SP_REFUND_PROCESSING SpRefundDisputeStatus = iota + 2
    
    // SP_REFUND_COMPLETED 表示已退款状态
    SP_REFUND_COMPLETED
    
    // SP_REFUND_FAILED 表示退款失败状态
    SP_REFUND_FAILED
)

// SpRefundDisputeStatusDesc 是退款争议状态到描述文本的映射
var SpRefundDisputeStatusDesc = map[SpRefundDisputeStatus]string{
    SP_REFUND_PROCESSING: "处理中",
    SP_REFUND_COMPLETED:  "已退款",
    SP_REFUND_FAILED:     "退款失败",
}

// String 实现 fmt.Stringer 接口
func (s SpRefundDisputeStatus) String() string {
    if desc, ok := SpRefundDisputeStatusDesc[s]; ok {
        return desc
    }
    return "未知争议状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s SpRefundDisputeStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *SpRefundDisputeStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range SpRefundDisputeStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的退款争议状态")
}