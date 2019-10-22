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

	IdeaSuggest        // отправил кнопку "Предложить идею"
	IdeaDescribed      // описал суть идеи
	ChooseDreamerType  // выбрал тип участника
	ChooseAnalystType  // выбрал аналитика
	ChooseDreamerCount // выбрал количество участников
	FormedTeam         // сформировал команду

	AddDreamer     // отправил кнопку "Добавить участника"
	AddDreamerName // добавил имя участника
	AddDreamerMail // добавил почту участника
	AddDreamerRole // добавил роль участника
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
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Keyboard(
				"./keyboards/analysis/dreamers_an_specialization.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseAnalystType), 0)
		}

		if state == int(ChooseAnalystType) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Keyboard(
				"./keyboards/dreamers_count.json").Execute()
			redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		}

		if state == int(ChooseDreamerCount) {
			api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Keyboard(
				"./keyboards/dreamers_type.json").Execute()
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
		Addr: "localhost:6379",
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
