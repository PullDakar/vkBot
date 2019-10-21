package api

import (
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"vkBot/keyboards"
)

type MessageSendQuery interface {
	UserId(string) *MessageSendQueryBuilder
	Keyboard(string) *MessageSendQueryBuilder
	Message(string) *MessageSendQueryBuilder
}

type MessageSendQueryBuilder struct {
	Vk *VkQueryBuilder
}

func (queryBuilder *MessageSendQueryBuilder) UserId(userId string) *MessageSendQueryBuilder {
	queryBuilder.Vk.Params["user_id"] = userId
	queryBuilder.Vk.Params["random_id"] = strconv.FormatInt(rand.Int63n(math.MaxInt64), 10)

	return queryBuilder
}

func (queryBuilder *MessageSendQueryBuilder) Keyboard(pathToJson string) *MessageSendQueryBuilder {
	queryBuilder.Vk.Params["keyboard"] = url.QueryEscape(keyboards.ParseJsonFileToString(pathToJson))
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

	b, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Panic("Error while reading response: ", bodyErr)
	}

	log.Println("Response: ", string(b))
	return resp
}
