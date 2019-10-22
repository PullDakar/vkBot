package main

import (
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"vkBot/api"
	"vkBot/api/actors"
	"vkBot/domain"
)

var redisClient *redis.Client
var group = actors.GroupActor{
	GroupId:     187421915,
	AccessToken: "dc2665238b736d270d6314240e62affceca6e8560c16d8f661cba5c58d83e030ff54a20236240141bf3e0",
}

type DialogState int

const (
	DialogBegins DialogState = iota // отправил кнопку "Начать"

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

type Sender struct {
	pathToMessage  string
	pathToKeyboard string
}

func getDialogState(authorId string) int {
	state, _ := redisClient.Get(authorId).Int()
	return state
}

func setDialogState(authorId string, state DialogState) {
	redisClient.Set(authorId, state, 0)
}

func messageRouterHandler(w http.ResponseWriter, r *http.Request) {
	b, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Panic("Error while reading response: ", readErr)
	}

	msg := &domain.Message{}
	unmarshalErr := msg.UnmarshalJSON(b)
	if unmarshalErr != nil {
		log.Panic("Error while unmarshal response: ", unmarshalErr)
	}

	if msg.Type == domain.MessageNew {
		state := getDialogState(msg.AuthorId)
		if msg.Text == "Начать" && state == int(DialogBegins) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Keyboard(
				"./keyboards/init.json").Message(getFileContent("./patterns/init")).Execute()
		} else if msg.Text == "Предложить идею" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(getFileContent("./patterns/idea/idea_suggest")).Execute()
			redisClient.Set(msg.AuthorId, int(IdeaDescribed), 0)
		}

		if state == int(IdeaDescribed) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Keyboard(
				"./keyboards/dreamers_type.json").Message(getFileContent("./patterns/idea/dreamers_choose")).Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerType), 0)
		}

		if state == int(ChooseDreamerType) && msg.Text == "Аналитик" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message("Выберите специализацию аналитика:").Keyboard(
				"./keyboards/analysis/dreamers_an_specialization.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseAnalystType), 0)
		}

		if state == int(ChooseAnalystType) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите количество аналитиков для успешной реализации идеи:").Keyboard(
				"./keyboards/dreamers_count.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		}

		if state == int(ChooseDreamerCount) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message("Требуется ли Вам еще специалисты?").Keyboard(
				"./keyboards/dreamers_type.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerType), 0)
		}

		if state == int(ChooseDreamerType) && msg.Text == "Разработчик" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите специализацию разработчика:").Keyboard(
				"./keyboards/dev/dreamers_subtype_dev.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDevType), 0)
		}

		if state == int(ChooseDreamerType) && msg.Text == "Тестировщик" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите специализацию тестировщика:").Keyboard(
				"./keyboards/tester/dreamers_test_specialization.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseTesterType), 0)
		}

		if state == int(ChooseTesterType) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите количество " + msg.Text + " тестировщиков для успешной реализации идеи:").Keyboard(
				"./keyboards/dreamers_count.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		}

		if state == int(ChooseDreamerType) && msg.Text == "Менеджер проектов" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите количество менеджеров для успешной реализации идеи:").Keyboard(
				"./keyboards/dreamers_count.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		}

		if state == int(ChooseDevType) && msg.Text == "Бэк" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите специализацию back-end разработчика:").Keyboard(
				"./keyboards/dev/dreamers_back_specialization.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseBackDevType), 0)
		}

		if state == int(ChooseDevType) && msg.Text == "Фронт" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите специализацию front-end разработчика:").Keyboard(
				"./keyboards/dev/dreamers_front_specialization.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseFrontDevType), 0)
		}

		if state == int(ChooseDevType) && msg.Text == "Мобильный" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите специализацию мобильного разработчика:").Keyboard(
				"./keyboards/dev/dreamers_mobile_specialization.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseMobileDevType), 0)
		}

		if state == int(ChooseBackDevType) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите количество " + msg.Text + " разработчиков для успешной реализации идеи:").Keyboard(
				"./keyboards/dreamers_count.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		}

		if state == int(ChooseFrontDevType) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите количество " + msg.Text + " разработчиков для успешной реализации идеи:").Keyboard(
				"./keyboards/dreamers_count.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		}

		if state == int(ChooseMobileDevType) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
				"Выберите количество " + msg.Text + " разработчиков для успешной реализации идеи:").Keyboard(
				"./keyboards/dreamers_count.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		}

		if state == int(ChooseDreamerType) && msg.Text == "Сформировать команду" {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message("Спасибо! Ожидайте, когда команда соберется! Удачи!").Keyboard(
				"./keyboards/init.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerType), 0)
		}
	}

	_, _ = w.Write([]byte("ok"))
}

func getFileContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic("Error while reading file: ", path, err)
	}

	return string(data)
}

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
	})

	_, pingErr := redisClient.Ping().Result()
	if pingErr != nil {
		log.Panic("Error while pinging redis redisClient: ", pingErr)
	}
	defer redisClient.Close()

	http.HandleFunc("/", messageRouterHandler)

	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
