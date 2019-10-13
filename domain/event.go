package domain

type MessageType string

/*
	Перечисление, покрывающее все типы сообщений
*/
const (
	MessageTyping MessageType = "message_typing_state"
	MessageNew                = "message_new"
	MessageReply              = "message_reply"
)

var messageMapper = map[string]MessageType{
	"message_typing_state": MessageTyping,
	"message_new":          MessageNew,
	"message_reply":        MessageReply,
}

func GetMessageTypeByStrCode(code string) MessageType {
	return messageMapper[code]
}
