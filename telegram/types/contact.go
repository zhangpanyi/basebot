package types

import (
	"encoding/json"
)

// 联系人
type Contact struct {
	PhoneNumber string `json:"phone_number"`        // 电话号码
	FirstName   string `json:"first_name"`          // name
	LastName    string `json:"last_name,omitempty"` // name
	UserID      int64  `json:"user_id,omitempty"`   // 用户唯一ID
}

// 转换为JSON
func (contact *Contact) ToJSON() ([]byte, error) {
	return json.Marshal(contact)
}

// 从JSON反序列化
func (contact *Contact) FromJSON(jsb []byte) error {
	return json.Unmarshal(jsb, contact)
}
