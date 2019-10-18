package api

const apiAddress string = "https://api.vk.com/method/"

type IVkQuery interface {
	Execute() string
}

type IVkQueryBuilder interface {
	Messages() Messages
}

type VkQueryBuilder struct {
	Method string
	Params map[string]string
}

func NewVkRequest() IVkQueryBuilder {
	return &VkQueryBuilder{}
}

func (vkQueryBuilder *VkQueryBuilder) Messages() Messages {
	return &MessagesBuilder{Vk: vkQueryBuilder}
}

func (vkQueryBuilder VkQueryBuilder) FormRequest() string {
	res := apiAddress
	res += vkQueryBuilder.Method + "?"

	for k, v := range vkQueryBuilder.Params {
		res += k + "=" + v + "&"
	}

	return res[:len(res)-1]
}
