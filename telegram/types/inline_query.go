package types

import (
	"encoding/json"
)

// 内联回调
type InlineQuery struct {
	ID     string `json:"id"`     // 查询唯一ID
	From   *User  `json:"from"`   // 发送者
	Query  string `json:"query"`  // 查询数据
	Offset string `json:"offset"` // 结果偏移
}

// 转换为JSON
func (inlineQuery *InlineQuery) ToJSON() ([]byte, error) {
	return json.Marshal(inlineQuery)
}

// 从JSON反序列化
func (inlineQuery *InlineQuery) FromJSON(jsb []byte) error {
	return json.Unmarshal(jsb, inlineQuery)
}
