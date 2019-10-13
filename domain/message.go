package domain

import (
	"encoding/json"
)

// plain old для сообщения
type Message struct {
	Id       int64
	Type     MessageType
	AuthorId int64
	Text     string
}

func (m *Message) UnmarshalJSON(b []byte) error {
	var plainJsonStruct interface{}
	unmarshalErr := json.Unmarshal(b, &plainJsonStruct)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	messageStruct := plainJsonStruct.(map[string]interface{})
	msgObject := messageStruct["object"].(map[string]interface{})

	m.Type = GetMessageTypeByStrCode(messageStruct["type"].(string))

	if id, exist := msgObject["id"]; exist {
		m.Id = int64(id.(float64))
	}

	if authorId, exist := msgObject["from_id"]; exist {
		m.AuthorId = int64(authorId.(float64))
	}

	if text, exist := msgObject["text"]; exist {
		m.Text = text.(string)
	}

	// TODO fix STUB
	return nil
}
