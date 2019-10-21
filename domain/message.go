package domain

import (
	"encoding/json"
	"log"
	"strconv"
)

// plain old для сообщения
type Message struct {
	Id       string
	Type     MessageType
	AuthorId string
	Text     string
}

func (m *Message) UnmarshalJSON(b []byte) error {
	log.Println("Input json: ", string(b))

	var plainJsonStruct interface{}
	unmarshalErr := json.Unmarshal(b, &plainJsonStruct)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	messageStruct := plainJsonStruct.(map[string]interface{})
	msgObject := messageStruct["object"].(map[string]interface{})

	m.Type = GetMessageTypeByStrCode(messageStruct["type"].(string))

	if id, exist := msgObject["id"]; exist {
		m.Id = strconv.FormatFloat(id.(float64), 'f', -1, 64)
	}

	if authorId, exist := msgObject["from_id"]; exist {
		m.AuthorId = strconv.FormatFloat(authorId.(float64), 'f', -1, 64)
	}

	if text, exist := msgObject["text"]; exist {
		m.Text = text.(string)
	}

	// TODO fix STUB
	return nil
}
