package api

import (
	"strconv"
	"vkBot/api/actors"
)

type Messages interface {
	Send(actors.GroupActor) MessageSendQuery
}

type MessagesBuilder struct {
	Vk *VkQueryBuilder
}

func (mb *MessagesBuilder) Send(actor actors.GroupActor) MessageSendQuery {
	mb.Vk.Method = "messages.send"

	mb.Vk.Params = make(map[string]string)
	mb.Vk.Params["group_id"] = strconv.FormatInt(actor.GroupId, 10)
	mb.Vk.Params["access_token"] = actor.AccessToken
	mb.Vk.Params["v"] = "5.102"

	return &MessageSendQueryBuilder{Vk: mb.Vk}
}
