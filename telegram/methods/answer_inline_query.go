package methods

import (
	"reflect"
	"strconv"

	"github.com/zhangpanyi/basebot/telegram/types"
)

// 应答查询回调
type answerInlineQuery struct {
	InlineQueryID string              `json:"inline_query_id"`       // 内联查询唯一ID
	Results       []InlineQueryResult `json:"results"`               // 结果集
	CacheTime     int32               `json:"cache_time,omitempty"`  // 缓存时间
	IsPersonal    bool                `json:"is_personal,omitempty"` // 是否缓存结果到服务器
	NextOffset    string              `json:"next_offset,omitempty"` // 下次偏移量
}

// 内联查询结果接口
type InlineQueryResult interface {
	SetType()
}

// InlineQueryResultPhoto 照片结果
type InlineQueryResultPhoto struct {
	Type        string `json:"type"`                  // 结果类型
	ID          string `json:"id"`                    // 结果ID
	PhotoURL    string `json:"photo_url"`             // 照片地址
	ThumbURL    string `json:"thumb_url"`             // 缩略图地址
	Title       string `json:"title,omitempty"`       // 标题
	Description string `json:"description,omitempty"` // 描述
	Caption     string `json:"caption,omitempty"`     // 说明文字
	ParseMode   string `json:"parse_mode,omitempty"`  // 解析模式
}

func (result *InlineQueryResultPhoto) SetType() {
	result.Type = "photo"
}

// AnswerInlineQuery 应答内联查询
func (bot *BotExt) AnswerInlineQuery(query *types.InlineQuery, offset, cacheTime int32, results []InlineQueryResult) error {
	request := answerInlineQuery{
		InlineQueryID: query.ID,
		CacheTime:     cacheTime,
		IsPersonal:    false,
		NextOffset:    strconv.FormatInt(int64(offset), 10),
	}

	if results != nil && !reflect.ValueOf(results).IsNil() {
		request.Results = results
		for _, result := range request.Results {
			result.SetType()
		}
	}
	_, err := bot.Call("answerInlineQuery", &request)
	return err
}
