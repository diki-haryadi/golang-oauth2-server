package model

type TelegramInlineKeyboardData struct {
	Value string
	URL   string
	Text  string
}

type TelegramInlineKeyboardRow struct {
	Data []TelegramInlineKeyboardData
}
