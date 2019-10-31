package cache

import (
	"log"
	"vkBot/api"
	"vkBot/api/actors"
	"vkBot/cache/domain"
)

// FIXME переделать чтение полей из yml
var group = actors.GroupActor{
	GroupId:     187421915,
	AccessToken: "dc2665238b736d270d6314240e62affceca6e8560c16d8f661cba5c58d83e030ff54a20236240141bf3e0",
}

type DialogState interface {
	reply(ctx *DialogContext)
}

type DialogContext struct {
	InputMessage *domain.Message
	CurrentState DialogState
}

func (ctx *DialogContext) Reply() {
	ctx.CurrentState.reply(ctx)
}

type DialogNotStartedState struct{}

func (state DialogNotStartedState) reply(ctx *DialogContext) {
	if ctx.InputMessage.Text == "Начать" {
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Привет! Данный бот поможет тебе найти нужных специалистов для реализации твоей бизнес идеи! " +
				"В случае, если идея будет доработана до минимально жизнеспособного продукта (MVP), у тебя появится " +
				"возможность получить свои первые инвестиции на развитие проекта! Удачи!").KeyboardByPath(
			"./keyboard/init.json").Execute()
		setDialogState(ctx.InputMessage.AuthorId, DialogStartedState{})
	}
}

type DialogStartedState struct{}

func (state DialogStartedState) reply(ctx *DialogContext) {
	switch ctx.InputMessage.Text {
	case "Предложить идею":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Опиши, пожалуйста, свою идею понятным и доступным языком:").Execute()
		setDialogState(ctx.InputMessage.AuthorId, IdeaBranchStartedState{})
	case "Добавить участника":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Как зовут нового стартапера?").Execute()
		setDialogState(ctx.InputMessage.AuthorId, AddingDreamerStartedState{})
	}
}

type IdeaBranchStartedState struct{}

func (state IdeaBranchStartedState) reply(ctx *DialogContext) {
	idea := ctx.InputMessage.Text
	log.Println("Suggested idea: " + idea)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, специалистов, нужных тебе для реализации идеи").KeyboardByPath(
		"./keyboard/dreamers_type.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, IdeaDescribedState{})
}

type IdeaDescribedState struct{}

func (state IdeaDescribedState) reply(ctx *DialogContext) {
	switch ctx.InputMessage.Text {
	case "Аналитик":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Выбери, пожалуйста, тип аналитика").KeyboardByPath(
			"./keyboard/analysis/dreamers_an_specialization.json").Execute()
		setDialogState(ctx.InputMessage.AuthorId, AnalystTypeChosenState{})
	case "Разработчик":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Выбери, пожалуйста, тип разработчика").KeyboardByPath(
			"./keyboard/dev/dreamers_subtype_dev.json").Execute()
		setDialogState(ctx.InputMessage.AuthorId, DeveloperTypeChosenState{})
	case "Тестировщик":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Выбери, пожалуйста, тип тестировщика").KeyboardByPath(
			"./keyboard/tester/dreamers_test_specialization.json").Execute()
		setDialogState(ctx.InputMessage.AuthorId, TesterTypeChosenState{})
	case "Менеджер проектов":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Выбери, пожалуйста, количество менеджеров").KeyboardByPath(
			"./keyboard/dreamers_count.json").Execute()
		setDialogState(ctx.InputMessage.AuthorId, SpecialistCountSelectedState{})
	case "Сформировать команду":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Спасибо! Теперь потребуется немного подождать, пока мы найдем нужных специалистов").Execute()
		setDialogState(ctx.InputMessage.AuthorId, TeamFormedState{})
	}
}

type AnalystTypeChosenState struct{}

func (state AnalystTypeChosenState) reply(ctx *DialogContext) {
	analystType := ctx.InputMessage.Text
	log.Println("Analyst type: " + analystType)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, количество аналитиков").KeyboardByPath(
		"./keyboard/dreamers_count.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, SpecialistCountSelectedState{})
}

type DeveloperTypeChosenState struct{}

func (state DeveloperTypeChosenState) reply(ctx *DialogContext) {
	devType := ctx.InputMessage.Text
	log.Println("Developer type: " + devType)

	switch ctx.InputMessage.Text {
	case "Бэк":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Выбери, пожалуйста, тип разработчика").KeyboardByPath(
			"./keyboard/dev/dreamers_back_specialization.json").Execute()
		setDialogState(ctx.InputMessage.AuthorId, BackDeveloperChosenState{})
	case "Фронт":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Выбери, пожалуйста, тип разработчика").KeyboardByPath(
			"./keyboard/dev/dreamers_front_specialization.json").Execute()
		setDialogState(ctx.InputMessage.AuthorId, FrontDeveloperChosenState{})
	case "Мобильный":
		api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
			"Выбери, пожалуйста, тип разработчика").KeyboardByPath(
			"./keyboard/dev/dreamers_mobile_specialization.json").Execute()
		setDialogState(ctx.InputMessage.AuthorId, MobileDeveloperChosenState{})
	}
}

type TesterTypeChosenState struct{}

func (state TesterTypeChosenState) reply(ctx *DialogContext) {
	testerType := ctx.InputMessage.Text
	log.Println("Tester type: " + testerType)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, количество тестировщиков").KeyboardByPath(
		"./keyboard/dreamers_count.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, SpecialistCountSelectedState{})
}

type SpecialistCountSelectedState struct{}

func (state SpecialistCountSelectedState) reply(ctx *DialogContext) {
	count := ctx.InputMessage.Text
	log.Println("Selected specialist count: " + count)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, специалистов, нужных тебе для реализации идеи").KeyboardByPath(
		"./keyboard/dreamers_type.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, IdeaDescribedState{})
}

type TeamFormedState struct{}

func (state TeamFormedState) reply(ctx *DialogContext) {
	log.Println("Team has formed successfully")
}

type BackDeveloperChosenState struct{}

func (state BackDeveloperChosenState) reply(ctx *DialogContext) {
	backDevType := ctx.InputMessage.Text
	log.Println("Chosen back-end developer: " + backDevType)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, количество " + backDevType + " разработчиков").KeyboardByPath(
		"./keyboard/dreamers_count.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, SpecialistCountSelectedState{})
}

type FrontDeveloperChosenState struct{}

func (state FrontDeveloperChosenState) reply(ctx *DialogContext) {
	frontDevType := ctx.InputMessage.Text
	log.Println("Chosen back-end developer: " + frontDevType)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, количество " + frontDevType + " разработчиков").KeyboardByPath(
		"./keyboard/dreamers_count.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, SpecialistCountSelectedState{})
}

type MobileDeveloperChosenState struct{}

func (state MobileDeveloperChosenState) reply(ctx *DialogContext) {
	mobileDevType := ctx.InputMessage.Text
	log.Println("Chosen back-end developer: " + mobileDevType)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, количество " + mobileDevType + " разработчиков").KeyboardByPath(
		"./keyboard/dreamers_count.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, SpecialistCountSelectedState{})
}

type AddingDreamerStartedState struct{}

func (state AddingDreamerStartedState) reply(ctx *DialogContext) {
	dreamerName := ctx.InputMessage.Text
	log.Println("New member name: " + dreamerName)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Ссылка на его вк: ").Execute()
	setDialogState(ctx.InputMessage.AuthorId, AddedDreamerLinkState{})
}

type AddedDreamerLinkState struct{}

func (state AddedDreamerLinkState) reply(ctx *DialogContext) {
	dreamerLink := ctx.InputMessage.Text
	log.Println("New member link: " + dreamerLink)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Почта:").Execute()
	setDialogState(ctx.InputMessage.AuthorId, AddedDreamerMailState{})
}

type AddedDreamerMailState struct{}

func (state AddedDreamerMailState) reply(ctx *DialogContext) {
	dreamerMail := ctx.InputMessage.Text
	log.Println("New member mail: " + dreamerMail)

	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, его специализацию").KeyboardByPath(
		"./keyboard/dreamers_type.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, NewDreamerSpecializationSelectedState{})
}

// TODO спросить у hr о возможности добавления почты сотрудника
type NewDreamerSpecializationSelectedState struct{}

func (state NewDreamerSpecializationSelectedState) reply(ctx *DialogContext) {
	api.NewVkRequest().Messages().Send(group).UserId(ctx.InputMessage.AuthorId).Message(
		"Выбери, пожалуйста, специалистов, нужных тебе для реализации идеи").KeyboardByPath(
		"./keyboard/dreamers_type.json").Execute()
	setDialogState(ctx.InputMessage.AuthorId, IdeaDescribedState{})
}
