package utils
import (
	"fmt"
	"strconv"
)
func ConvertToUint(value interface{}) uint {
	if value == nil {
		return 0
	}
	
	switch v := value.(type) {
	case string:
		if v == "" || v == "0" {
			return 0
		}
		if id, err := strconv.ParseUint(v, 10, 32); err == nil {
			return uint(id)
		}
	case float64: // JSON数字默认是float64
		if v == 0 {
			return 0
		}
		return uint(v)
	case int:
		return uint(v)
	case int64:
		return uint(v)
	case uint:
		return v
	case uint64:
		return uint(v)
	default:
		// 尝试转换为字符串再解析
		if str, ok := v.(fmt.Stringer); ok {
			return ConvertToUint(str.String())
		}
	}
	
	return 0
}

func ConvertToFloat64(value interface{}) float64 {
	if value == nil {
		return 0
	}

	switch v := value.(type) {
	case string:
		if v == "" {
			return 0
		}
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		if str, ok := value.(fmt.Stringer); ok {
			return ConvertToFloat64(str.String())
		}
	}
	return 0
}