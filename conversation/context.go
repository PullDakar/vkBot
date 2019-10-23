package conversation

import (
	"vkBot/api"
	"vkBot/api/actors"
	"vkBot/conversation/cache"
	"vkBot/conversation/domain"
)

var group = actors.GroupActor{
	GroupId:     187421915,
	AccessToken: "dc2665238b736d270d6314240e62affceca6e8560c16d8f661cba5c58d83e030ff54a20236240141bf3e0",
}

type State int

const (
	XDialogNotStarted State = iota
	DialogStarted           // отправил кнопку "Начать"

	IdeaDescribed // описал суть идеи

	ChooseDreamerType // выбрал тип участника

	ChooseAnalystType // выбрал аналитика

	ChooseDevType       //  выбрал разработчика
	ChooseBackDevType   //  выбрал back-end разработчика
	ChooseFrontDevType  // выбрал front-end разработчика
	ChooseMobileDevType // выбрал мобильного разработчика

	ChooseTesterType // выбрал тестировщика

	ChooseDreamerCount // выбрал количество участников
)

type DialogState interface {
	Reply(ctx *DialogContext)
}

type DialogNotStartedState struct {
	stateId int8
}

func (state DialogNotStartedState) Reply(ctx *DialogContext) {
	if ctx.InputMessage.Text == "Начать" {
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Привет! Данный бот поможет тебе найти нужных специалистов для реализации твоей бизнес идеи! " +
				"В случае, если идея будет доработана до минимально жизнеспособного продукта (MVP), у тебя появится " +
				"возможность получить свои первые инвестиции на развитие проекта! Удачи!").Keyboard(
			"./keyboards/init.json")
		cache.SetDialogState(ctx.InputMessage.AuthorId, DialogStartedState{stateId: 1})
	}
}

type DialogStartedState struct {
	stateId int8
}

func (state DialogStartedState) Reply(ctx *DialogContext) {

}

type DialogContext struct {
	InputMessage *domain.Message
	CurrentState DialogState
}

func (ctx *DialogContext) Reply() {
	ctx.CurrentState.Reply(ctx)
}
