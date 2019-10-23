package conversation

import (
	"github.com/go-redis/redis"
	"log"
	"vkBot/api"
	"vkBot/api/actors"
	"vkBot/conversation/domain"
)

var group = actors.GroupActor{
	GroupId:     187421915,
	AccessToken: "dc2665238b736d270d6314240e62affceca6e8560c16d8f661cba5c58d83e030ff54a20236240141bf3e0",
}

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var possibleStates = [...]DialogState{
	DialogNotStartedState{},
	DialogStartedState{},
}

func GetDialogState(authorId string) DialogState {
	stateIndex, _ := redisClient.Get(authorId).Int()
	return possibleStates[stateIndex]
}

func setDialogState(key string, value DialogState) {
	index, notFoundErr := indexOf(value)
	if notFoundErr != nil {
		log.Panic(notFoundErr)
	}

	_, err := redisClient.Set(key, index, 0).Result()
	if err != nil {
		log.Panic("Error while setting value to redis cache by key", key, err)
	}
}

func indexOf(value DialogState) (int, error) {
	for index, state := range possibleStates {
		if state == value {
			return index, nil
		}
	}

	// TODO добавить передачу в ошибку конкретный state
	return -1, &stateNotFound{"Dialog state not found "}
}

type stateNotFound struct {
	error string
}

func (err *stateNotFound) Error() string {
	return err.error
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

type DialogNotStartedState struct{}

func (state DialogNotStartedState) Reply(ctx *DialogContext) {
	if ctx.InputMessage.Text == "Начать" {
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Привет! Данный бот поможет тебе найти нужных специалистов для реализации твоей бизнес идеи! " +
				"В случае, если идея будет доработана до минимально жизнеспособного продукта (MVP), у тебя появится " +
				"возможность получить свои первые инвестиции на развитие проекта! Удачи!").Keyboard(
			"./keyboards/init.json")
		setDialogState(ctx.InputMessage.AuthorId, DialogStartedState{})
	}
}

type DialogStartedState struct{}

func (state DialogStartedState) Reply(ctx *DialogContext) {

}

type DialogContext struct {
	InputMessage *domain.Message
	CurrentState DialogState
}

func (ctx *DialogContext) Reply() {
	ctx.CurrentState.Reply(ctx)
}
