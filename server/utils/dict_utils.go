package utils

import (
	"encoding/json"
	"time"
)

// EnumDictResponse 枚举字典响应结构
type EnumDictResponse struct {
	Code    int           `json:"code"`
	Result  []DictResult  `json:"result"`
	Error   interface{}   `json:"error"`
	Message interface{}   `json:"message"`
	Time    int64         `json:"time"`
}

// DictResult 字典结果
type DictResult struct {
	Code  string     `json:"code"`
	Items []DictItem `json:"items"`
}

// DictItem 字典项
type DictItem struct {
	Value interface{} `json:"value"`
	Label string      `json:"label"`
	Color string      `json:"color"`
}

// GetEnumDict 获取枚举字典
func GetEnumDict() EnumDictResponse {
	// 创建响应对象
	response := EnumDictResponse{
		Code:    0,
		Error:   nil,
		Message: nil,
		Time:    time.Now().UnixNano() / int64(time.Millisecond),
	}

	// 添加所有枚举字典
	response.Result = []DictResult{
		{
			Code: "CareAdminConstant$Status",
			Items: []DictItem{
				{Value: 1, Label: "启用", Color: "#108ee9"},
				{Value: 2, Label: "停用", Color: "#d9d9d9"},
			},
		},
		{
			Code: "CoreConfigConstant$Key",
			Items: []DictItem{
				{Value: 0, Label: "未知", Color: ""},
				{Value: 1, Label: "云支付", Color: ""},
				{Value: 2, Label: "免费体验", Color: ""},
				{Value: 3, Label: "累计发放奖励", Color: ""},
				{Value: 4, Label: "对话视频配置", Color: ""},
				{Value: 5, Label: "AI二创视频相识度", Color: ""},
				{Value: 8, Label: "积分购买配置", Color: ""},
				{Value: 9, Label: "AI生图每日免费数量", Color: ""},
				{Value: 10, Label: "推送通道配置", Color: ""},
			},
		},
		{
			Code: "CoreMenuConstant$HideStatus",
			Items: []DictItem{
				{Value: 0, Label: "未知", Color: ""},
				{Value: 1, Label: "隐藏", Color: "#87d068"},
				{Value: 2, Label: "不隐藏", Color: "#d9d9d9"},
			},
		},
		{
			Code: "CoreRoleConstant$RoleStatus",
			Items: []DictItem{
				{Value: 0, Label: "未知", Color: ""},
				{Value: 1, Label: "启用", Color: "#108ee9"},
				{Value: 2, Label: "停用", Color: "#d9d9d9"},
			},
		},
		{
			Code: "CommonConstant$Sex",
			Items: []DictItem{
				{Value: 0, Label: "未知", Color: ""},
				{Value: 1, Label: "男", Color: "#108ee9"},
				{Value: 2, Label: "女", Color: "#d9d9d9"},
			},
		},
		{
			Code: "CommonConstant$State",
			Items: []DictItem{
				{Value: 0, Label: "未知", Color: ""},
				{Value: 1, Label: "启用", Color: "#108ee9"},
				{Value: 2, Label: "停用", Color: "#d9d9d9"},
			},
		},
		{
			Code: "LogRecordConstant$Source",
			Items: []DictItem{
				{Value: 1, Label: "后台", Color: ""},
				{Value: 2, Label: "用户端", Color: ""},
			},
		},
		{
			Code: "CmsDocumentConstant$LinkType",
			Items: []DictItem{
				{Value: 1, Label: "内部文档", Color: "#108ee9"},
				{Value: 2, Label: "外部链接", Color: "#108ee9"},
			},
		},
		{
			Code: "CmsDocumentConstant$Status",
			Items: []DictItem{
				{Value: 1, Label: "已发布", Color: "#108ee9"},
				{Value: 2, Label: "未发布", Color: "#108ee9"},
			},
		},
		{
			Code: "CmsRecommendConstant$Status",
			Items: []DictItem{
				{Value: 1, Label: "已发布", Color: "#108ee9"},
				{Value: 2, Label: "未发布", Color: "#108ee9"},
			},
		},
		{
			Code: "CtTaskConstant$Platform",
			Items: []DictItem{
				{Value: 1, Label: "淘宝", Color: "#108ee9"},
			},
		},
		{
			Code: "CtTaskConstant$State",
			Items: []DictItem{
				{Value: 1, Label: "开启", Color: "#108ee9"},
				{Value: 2, Label: "关闭", Color: "#108ee9"},
			},
		},
		{
			Code: "CtTaskConstant$Type",
			Items: []DictItem{
				{Value: 1, Label: "关键词", Color: "#108ee9"},
				{Value: 2, Label: "详情", Color: "#108ee9"},
			},
		},
		{
			Code: "CtTaskLogConstant$State",
			Items: []DictItem{
				{Value: 1, Label: "执行中", Color: "#108ee9"},
				{Value: 2, Label: "执行成功", Color: "#108ee9"},
				{Value: 3, Label: "执行失败", Color: "#108ee9"},
			},
		},
		{
			Code: "MpUserConstant$DefaultAddress",
			Items: []DictItem{
				{Value: 1, Label: "是", Color: "#108ee9"},
				{Value: 2, Label: "否", Color: "#108ee9"},
			},
		},
		{
			Code: "MpUserConstant$SEX",
			Items: []DictItem{
				{Value: 1, Label: "男", Color: "#108ee9"},
				{Value: 2, Label: "女", Color: "#108ee9"},
				{Value: 3, Label: "未知", Color: "#108ee9"},
			},
		},
		{
			Code: "MpUserConstant$ValidateStatus",
			Items: []DictItem{
				{Value: 1, Label: "未验证", Color: "#108ee9"},
				{Value: 2, Label: "已验证", Color: "#108ee9"},
			},
		},
		{
			Code: "SpCategoryConstant$Status",
			Items: []DictItem{
				{Value: 1, Label: "正常", Color: "#108ee9"},
				{Value: 2, Label: "停用", Color: "#108ee9"},
			},
		},
		{
			Code: "SpOrderConstant$PayType",
			Items: []DictItem{
				{Value: 1, Label: "货到付款", Color: "#108ee9"},
			},
		},
		{
			Code: "SpOrderConstant$SourceType",
			Items: []DictItem{
				{Value: 1, Label: "PC订单", Color: "#108ee9"},
				{Value: 2, Label: "移动端订单", Color: "#108ee9"},
			},
		},
		{
			Code: "SpOrderConstant$State",
			Items: []DictItem{
				{Value: 1, Label: "待付款", Color: ""},
				{Value: 2, Label: "待发货", Color: ""},
				{Value: 3, Label: "已发货", Color: ""},
				{Value: 4, Label: "已完成", Color: ""},
				{Value: 5, Label: "已关闭", Color: ""},
				{Value: 6, Label: "无效订单", Color: ""},
				{Value: 7, Label: "部分退款", Color: ""},
				{Value: 8, Label: "全部退款", Color: ""},
			},
		},
		{
			Code: "SpOrderConstant$disputeStatus",
			Items: []DictItem{
				{Value: 0, Label: "无争议", Color: ""},
				{Value: 1, Label: "买家发起争议", Color: ""},
				{Value: 2, Label: "卖家处理中", Color: ""},
				{Value: 3, Label: "平台仲裁中", Color: ""},
				{Value: 4, Label: "已解决", Color: ""},
			},
		},
		{
			Code: "SpProductConstant$Hot",
			Items: []DictItem{
				{Value: 1, Label: "是", Color: "#108ee9"},
				{Value: 2, Label: "否", Color: "#108ee9"},
			},
		},
		{
			Code: "SpProductConstant$OpenSku",
			Items: []DictItem{
				{Value: 1, Label: "是", Color: "#108ee9"},
				{Value: 2, Label: "否", Color: "#108ee9"},
			},
		},
		{
			Code: "SpProductConstant$SkuDefaultShow",
			Items: []DictItem{
				{Value: 1, Label: "是", Color: "#108ee9"},
				{Value: 2, Label: "否", Color: "#108ee9"},
			},
		},
		{
			Code: "SpProductConstant$SkuStatus",
			Items: []DictItem{
				{Value: 1, Label: "正常", Color: "#108ee9"},
				{Value: 2, Label: "停用", Color: "#108ee9"},
			},
		},
		{
			Code: "SpProductConstant$Status",
			Items: []DictItem{
				{Value: 1, Label: "上线", Color: "#108ee9"},
				{Value: 2, Label: "下线", Color: "#108ee9"},
				{Value: 3, Label: "待审核", Color: "#108ee9"},
				{Value: 4, Label: "审核不通过", Color: "#108ee9"},
			},
		},
		{
			Code: "SpRefundConstant$disputeStatus",
			Items: []DictItem{
				{Value: 2, Label: "处理中", Color: ""},
				{Value: 3, Label: "已退款", Color: ""},
				{Value: 4, Label: "退款失败", Color: ""},
			},
		},
		{
			Code: "SpTagConstant$Status",
			Items: []DictItem{
				{Value: 1, Label: "已发布", Color: "#108ee9"},
				{Value: 2, Label: "未发布", Color: "#108ee9"},
			},
		},
	}

	return response
}

// GetEnumDictJSON 获取枚举字典的JSON字符串
func GetEnumDictJSON() (string, error) {
	response := GetEnumDict()
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}