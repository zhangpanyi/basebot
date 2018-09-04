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

// 输入消息内容接口
type InputMessageContent interface{}

// 输入文本消息内容
type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`             // 消息文本
	ParseMode             string `json:"parse_mode"`               // 解析模式
	DisableWebPagePreview bool   `json:"disable_web_page_preview"` // 禁用页面预览
}

// 文章结果
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`                   // 结果类型
	ID                  string                `json:"id"`                     // 结果ID
	Title               string                `json:"title"`                  // 标题
	InputMessageContent InputMessageContent   `json:"input_message_content"`  // 消息内容
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Reply Markup
	Description         string                `json:"description,omitempty"`  // 描述
	URL                 string                `json:"url,omitempty"`          // 地址
	HideURL             bool                  `json:"hide_url,omitempty"`     // 隐藏地址
	ThumbURL            string                `json:"thumb_url,omitempty"`    // 缩略图地址
	ThumbWidth          int32                 `json:"thumb_width,omitempty"`  // 缩略图宽度
	ThumbHeight         int32                 `json:"thumb_height,omitempty"` // 缩略图高度
}

// 设置类型
func (result *InlineQueryResultArticle) SetType() {
	result.Type = "article"
}

// 应答内联查询
func (bot *BotExt) AnswerInlineQuery(query *types.InlineQuery, offset *int32, cacheTime int32,
	results []InlineQueryResult) error {

	nextOffset := ""
	if offset != nil {
		strconv.FormatInt(int64(*offset), 10)
	}

	request := answerInlineQuery{
		InlineQueryID: query.ID,
		CacheTime:     cacheTime,
		IsPersonal:    true,
		NextOffset:    nextOffset,
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
