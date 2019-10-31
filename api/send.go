package api

import (
	"encoding/binary"
	"encoding/json"
	"github.com/spaolacci/murmur3"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"vkBot/keyboard"
)

type MessageSendQuery interface {
	UserId(string) *MessageSendQueryBuilder
	KeyboardByPath(string) *MessageSendQueryBuilder
	Message(string) *MessageSendQueryBuilder
}

type MessageSendQueryBuilder struct {
	Vk *VkQueryBuilder
}

func (queryBuilder *MessageSendQueryBuilder) UserId(userId string) *MessageSendQueryBuilder {
	queryBuilder.Vk.Params["user_id"] = userId
	queryBuilder.Vk.Params["random_id"] = strconv.FormatUint(randomId(userId), 10)

	return queryBuilder
}

func randomId(userId string) uint64 {
	h := murmur3.New64()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().UnixNano()))

	_, timeErr := h.Write(b)
	if timeErr != nil {
		log.Panic(timeErr)
	}

	_, idErr := h.Write([]byte(userId))
	if idErr != nil {
		log.Panic(idErr)
	}

	return h.Sum64()
}

func (queryBuilder *MessageSendQueryBuilder) KeyboardByPath(pathToJson string) *MessageSendQueryBuilder {
	queryBuilder.Vk.Params["keyboard"] = url.QueryEscape(keyboard.ParseJsonFileToString(pathToJson))
	return queryBuilder
}

func (queryBuilder *MessageSendQueryBuilder) Keyboard(k keyboard.Keyboard) *MessageSendQueryBuilder {
	b, err := json.Marshal(k)
	if err != nil {
		log.Panic("Error while marshaling keyboard to json. Keyboard: "+string(b)+". Error: ", err)
	}

	queryBuilder.Vk.Params["keyboard"] = url.QueryEscape(string(b))
	return queryBuilder
}

func (queryBuilder *MessageSendQueryBuilder) Message(message string) *MessageSendQueryBuilder {
	queryBuilder.Vk.Params["message"] = url.QueryEscape(message)
	return queryBuilder
}

func (queryBuilder MessageSendQueryBuilder) Execute() *http.Response {
	reqUrl := queryBuilder.Vk.FormRequest()
	log.Println("Formed request: ", reqUrl)

	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Panic("Error while sending post request: ", resp)
	}

	log.Println("Response: ", resp)
	return resp
}
