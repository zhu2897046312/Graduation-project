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