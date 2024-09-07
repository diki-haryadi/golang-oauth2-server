package constants

type ChatbotChannelEnum int

const (
	ChatbotChannelTelegram ChatbotChannelEnum = 1
)

type ChatbotCommandMessage string

const (
	ChatbotCommandMessageStart         ChatbotCommandMessage = "/start"
	ChatbotCommandMessageEnterRole     ChatbotCommandMessage = "/enter_as"
	ChatbotCommandMessageCreateSession ChatbotCommandMessage = "/create_session"
	ChatbotCommandMessageQueueTrack    ChatbotCommandMessage = "/q"
)

type ChatbotUserRole string

const (
	ChatbotUserRoleGuest ChatbotUserRole = "guest"
	ChatbotUserRoleHost  ChatbotUserRole = "host"
)
