package api

import "vkBot/api/actions"

type VkApiClient struct {
	transportClient TransportClient
}

const apiAddress = "https://api.vk.com/method/"

func (vk VkApiClient) messages() actions.Messages {
	return actions.Messages{}
}
