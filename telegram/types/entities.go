package types

import (
	"encoding/json"
)

const (
	// 提及用户
	EntityMention = "mention "
	// 哈希标记
	EntityHashTag = "hashtag"
	// 机器人命令
	EntityBotCommand = "bot_command"
	// 超链接
	EntityURL = "url"
	// 电子邮箱
	EntityEmail = "email"
	// 加粗
	EntityBold = "bold"
	// 斜体
	EntityItalic = "italic"
	// 代码
	EntityCode = "code"
	// 预格式化文本
	EntityPre = "pre"
	// 文本链接
	EntityTextLink = "text_link"
	// 文本提及
	EntityTextMention = "text_mention "
)

// Entity信息
type MessageEntity struct {
	Type   string `json:"type"`           // Entity类型
	Offset int32  `json:"offset"`         // 偏移
	Length int32  `json:"length"`         // 长度
	URL    string `json:"url,omitempty"`  // 地址
	User   *User  `json:"user,omitempty"` // 用户信息
}

// 转换为JSON
func (entity *MessageEntity) ToJSON() ([]byte, error) {
	return json.Marshal(entity)
}

// 从JSON反序列化
func (entity *MessageEntity) FromJSON(jsb []byte) error {
	return json.Unmarshal(jsb, entity)
}
