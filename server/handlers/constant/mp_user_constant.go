package constant

import (
    "encoding/json"
    "errors"
)

// MpUserSex 表示用户性别的枚举类型
type MpUserSex int

const (
    // MP_USER_MAN 表示男性
    MP_USER_MAN MpUserSex = iota + 1
    
    // MP_USER_WOMAN 表示女性
    MP_USER_WOMAN
    
    // MP_USER_UNKNOWN 表示未知性别
    MP_USER_UNKNOWN
)

// MpUserSexDesc 是用户性别到描述文本的映射
var MpUserSexDesc = map[MpUserSex]string{
    MP_USER_MAN:     "性别:男:#108ee9",
    MP_USER_WOMAN:   "性别:女:#108ee9",
    MP_USER_UNKNOWN: "性别:未知:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s MpUserSex) String() string {
    if desc, ok := MpUserSexDesc[s]; ok {
        return desc
    }
    return "未知性别"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s MpUserSex) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *MpUserSex) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range MpUserSexDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的用户性别")
}

// MpUserValidateStatus 表示用户验证状态的枚举类型
type MpUserValidateStatus int

const (
    // MP_USER_VALIDATE_NO 表示未验证状态
    MP_USER_VALIDATE_NO MpUserValidateStatus = iota + 1
    
    // MP_USER_VALIDATE_YES 表示已验证状态
    MP_USER_VALIDATE_YES
)

// MpUserValidateStatusDesc 是用户验证状态到描述文本的映射
var MpUserValidateStatusDesc = map[MpUserValidateStatus]string{
    MP_USER_VALIDATE_NO:  "验证状态:未验证:#108ee9",
    MP_USER_VALIDATE_YES: "验证状态:已验证:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (s MpUserValidateStatus) String() string {
    if desc, ok := MpUserValidateStatusDesc[s]; ok {
        return desc
    }
    return "未知验证状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (s MpUserValidateStatus) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *MpUserValidateStatus) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range MpUserValidateStatusDesc {
        if v == desc {
            *s = k
            return nil
        }
    }
    return errors.New("未知的用户验证状态")
}

// MpUserDefaultAddress 表示用户默认地址的枚举类型
type MpUserDefaultAddress int

const (
    // MP_USER_DEFAULT_ADDRESS_YES 表示是默认地址
    MP_USER_DEFAULT_ADDRESS_YES MpUserDefaultAddress = iota + 1
    
    // MP_USER_DEFAULT_ADDRESS_NO 表示不是默认地址
    MP_USER_DEFAULT_ADDRESS_NO
)

// MpUserDefaultAddressDesc 是用户默认地址状态到描述文本的映射
var MpUserDefaultAddressDesc = map[MpUserDefaultAddress]string{
    MP_USER_DEFAULT_ADDRESS_YES: "是否默认收货地址:是:#108ee9",
    MP_USER_DEFAULT_ADDRESS_NO:  "是否默认收货地址:否:#108ee9",
}

// String 实现 fmt.Stringer 接口
func (a MpUserDefaultAddress) String() string {
    if desc, ok := MpUserDefaultAddressDesc[a]; ok {
        return desc
    }
    return "未知地址状态"
}

// MarshalJSON 实现 json.Marshaler 接口
func (a MpUserDefaultAddress) MarshalJSON() ([]byte, error) {
    return json.Marshal(a.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (a *MpUserDefaultAddress) UnmarshalJSON(data []byte) error {
    var desc string
    if err := json.Unmarshal(data, &desc); err != nil {
        return err
    }
    for k, v := range MpUserDefaultAddressDesc {
        if v == desc {
            *a = k
            return nil
        }
    }
    return errors.New("未知的用户默认地址状态")
}