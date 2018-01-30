package methods

import (
	"basebot/telegram/types"
	"encoding/json"
)

// 获取更新
type getUpdates struct {
	Timeout uint32 `json:"timeout"` // 超时时间
}

// 获取更新响应
type getUpdatesResonpe struct {
	OK     bool            `json:"ok"`               // 是否成功
	Result []*types.Update `json:"result,omitempty"` // 更新列表
}

// GetUpdates 获取更新
func (bot *BotExt) GetUpdates(timeout uint32) ([]*types.Update, error) {
	request := getUpdates{
		Timeout: timeout,
	}
	data, err := bot.Call("getUpdates", &request)
	if err != nil {
		return nil, err
	}

	res := getUpdatesResonpe{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
