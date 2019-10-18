package api

import (
	"math"
	"math/rand"
	"strconv"
	"vkBot/keyboards"
)

type MessageSendQuery interface {
	UserId(int64) *MessageSendQueryBuilder
	Keyboard(string) *MessageSendQueryBuilder
	Message(string) *MessageSendQueryBuilder
}

type MessageSendQueryBuilder struct {
	Vk *VkQueryBuilder
}

func (queryBuilder *MessageSendQueryBuilder) UserId(userId int64) *MessageSendQueryBuilder {
	queryBuilder.Vk.Params["user_id"] = strconv.FormatInt(userId, 10)
	queryBuilder.Vk.Params["random_id"] = strconv.FormatInt(rand.Int63n(math.MaxInt64), 10)

	return queryBuilder
}

func (queryBuilder *MessageSendQueryBuilder) Keyboard(pathToJson string) *MessageSendQueryBuilder {
	queryBuilder.Vk.Params["keyboard"] = keyboards.ParseJsonFileToString(pathToJson)
	return queryBuilder
}

func (queryBuilder *MessageSendQueryBuilder) Message(message string) *MessageSendQueryBuilder {
	queryBuilder.Vk.Params["message"] = message
	return queryBuilder
}

func (queryBuilder MessageSendQueryBuilder) Execute() string {
	return queryBuilder.Vk.FormRequest()
}
