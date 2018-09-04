package methods

import (
	"encoding/json"

	"github.com/zhangpanyi/basebot/telegram/types"
)

type editMessageText struct {
	ChatID                int64       `json:"chat_id,omitempty"`                  // 聊天ID
	MessageID             int32       `json:"message_id,omitempty"`               // 消息ID
	InlineMessageID       string      `json:"inline_message_id,omitempty"`        // 内联消息ID
	Text                  string      `json:"text"`                               // 消息文本
	ParseMode             string      `json:"parse_mode,omitempty"`               // 解析模式
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"` // 禁用网页预览
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`             // Reply Markup
}

// 标记消息回复标记
func (bot *BotExt) editReplyMarkup(request *editMessageText) (*types.Message, error) {
	res, err := bot.Call("editMessageText", request)
	if err != nil {
		return nil, err
	}

	if len(request.InlineMessageID) > 0 {
		return nil, nil
	}

	resonpe := SendMessageResonpe{}
	err = json.Unmarshal(res, &resonpe)
	if err != nil {
		return nil, err
	}
	return resonpe.Result, nil
}

// 编辑消息回复标记
func (bot *BotExt) EditReplyMarkup(chatID int64, messageID int32, text string, mdMode bool,
	markup *InlineKeyboardMarkup) (*types.Message, error) {

	request := editMessageText{
		ChatID:    chatID,
		MessageID: messageID,
		Text:      text,
	}

	if mdMode {
		request.ParseMode = ParseModeMarkdown
	}
	if markup != nil {
		request.ReplyMarkup = markup
	}
	return bot.editReplyMarkup(&request)
}

// 编辑消息回复标记
func (bot *BotExt) EditReplyMarkupByInlineMessageID(inlineMessageID, text string, mdMode bool,
	markup *InlineKeyboardMarkup) (*types.Message, error) {
	request := editMessageText{
		InlineMessageID: inlineMessageID,
		Text:            text,
	}

	if mdMode {
		request.ParseMode = ParseModeMarkdown
	}
	if markup != nil {
		request.ReplyMarkup = markup
	}
	return bot.editReplyMarkup(&request)
}

// 标记消息回复标记
func (bot *BotExt) EditMessageReplyMarkup(message *types.Message, text string, mdMode bool,
	markup *InlineKeyboardMarkup) (*types.Message, error) {
	return bot.EditReplyMarkup(message.Chat.ID, message.MessageID, text, mdMode, markup)
}

// 标记消息回复标记并禁用网页预览
func (bot *BotExt) EditReplyMarkupDisableWebPagePreview(chatID int64, messageID int32, text string, mdMode bool,
	markup *InlineKeyboardMarkup) (*types.Message, error) {

	request := editMessageText{
		ChatID:    chatID,
		MessageID: messageID,
		Text:      text,
		DisableWebPagePreview: true,
	}

	if mdMode {
		request.ParseMode = ParseModeMarkdown
	}
	if markup != nil {
		request.ReplyMarkup = markup
	}
	return bot.editReplyMarkup(&request)
}

// 标记消息回复标记并禁用网页预览
func (bot *BotExt) EditMessageReplyMarkupDisableWebPagePreview(message *types.Message, text string, mdMode bool,
	markup *InlineKeyboardMarkup) (*types.Message, error) {
	return bot.EditReplyMarkupDisableWebPagePreview(message.Chat.ID, message.MessageID, text, mdMode, markup)
}
