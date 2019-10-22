package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"vkBot/api/actors"
	"vkBot/conversation"
	"vkBot/conversation/cache"
	"vkBot/conversation/domain"
)

var group = actors.GroupActor{
	GroupId:     187421915,
	AccessToken: "dc2665238b736d270d6314240e62affceca6e8560c16d8f661cba5c58d83e030ff54a20236240141bf3e0",
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
		state := cache.GetDialogState(msg.AuthorId)
		dialogState = &conversation.DialogNotStartedState{}

		//state := cache.GetDialogState(msg.AuthorId)
		//if msg.Text == "Начать" && state == int(cache.DialogNotStarted) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Keyboard(
		//		"./keyboards/init.json").Message(getFileContent("./patterns/init")).Execute()
		//	cache.SetDialogState(msg.AuthorId, int(cache.DialogStarted))
		//} else if msg.Text == "Предложить идею" && state == int(cache.DialogStarted) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(getFileContent("./patterns/idea/idea_suggest")).Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.IdeaDescribed), 0)
		//}
		//
		//if state == int(cache.IdeaDescribed) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Keyboard(
		//		"./keyboards/dreamers_type.json").Message(getFileContent("./patterns/idea/dreamers_choose")).Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseDreamerType), 0)
		//}
		//
		//if state == int(cache.ChooseDreamerType) && msg.Text == "Аналитик" {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message("Выберите специализацию аналитика:").Keyboard(
		//		"./keyboards/analysis/dreamers_an_specialization.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseAnalystType), 0)
		//}
		//
		//if state == int(cache.ChooseAnalystType) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите количество аналитиков для успешной реализации идеи:").Keyboard(
		//		"./keyboards/dreamers_count.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseDreamerCount), 0)
		//}
		//
		//if state == int(cache.ChooseDreamerCount) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message("Требуется ли Вам еще специалисты?").Keyboard(
		//		"./keyboards/dreamers_type.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseDreamerType), 0)
		//}
		//
		//if state == int(cache.ChooseDreamerType) && msg.Text == "Разработчик" {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите специализацию разработчика:").Keyboard(
		//		"./keyboards/dev/dreamers_subtype_dev.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseDevType), 0)
		//}
		//
		//if state == int(cache.ChooseDreamerType) && msg.Text == "Тестировщик" {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите специализацию тестировщика:").Keyboard(
		//		"./keyboards/tester/dreamers_test_specialization.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseTesterType), 0)
		//}
		//
		//if state == int(cache.ChooseTesterType) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите количество " + msg.Text + " тестировщиков для успешной реализации идеи:").Keyboard(
		//		"./keyboards/dreamers_count.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseDreamerCount), 0)
		//}
		//
		//if state == int(cache.ChooseDreamerType) && msg.Text == "Менеджер проектов" {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите количество менеджеров для успешной реализации идеи:").Keyboard(
		//		"./keyboards/dreamers_count.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseDreamerCount), 0)
		//}
		//
		//if state == int(cache.ChooseDevType) && msg.Text == "Бэк" {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите специализацию back-end разработчика:").Keyboard(
		//		"./keyboards/dev/dreamers_back_specialization.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(cache.ChooseBackDevType), 0)
		//}
		//
		//if state == int(ChooseDevType) && msg.Text == "Фронт" {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите специализацию front-end разработчика:").Keyboard(
		//		"./keyboards/dev/dreamers_front_specialization.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(ChooseFrontDevType), 0)
		//}
		//
		//if state == int(ChooseDevType) && msg.Text == "Мобильный" {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите специализацию мобильного разработчика:").Keyboard(
		//		"./keyboards/dev/dreamers_mobile_specialization.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(ChooseMobileDevType), 0)
		//}
		//
		//if state == int(ChooseBackDevType) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите количество " + msg.Text + " разработчиков для успешной реализации идеи:").Keyboard(
		//		"./keyboards/dreamers_count.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		//}
		//
		//if state == int(ChooseFrontDevType) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите количество " + msg.Text + " разработчиков для успешной реализации идеи:").Keyboard(
		//		"./keyboards/dreamers_count.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		//}
		//
		//if state == int(ChooseMobileDevType) {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message(
		//		"Выберите количество " + msg.Text + " разработчиков для успешной реализации идеи:").Keyboard(
		//		"./keyboards/dreamers_count.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(ChooseDreamerCount), 0)
		//}
		//
		//if state == int(ChooseDreamerType) && msg.Text == "Сформировать команду" {
		//	api.NewVkRequest().Messages().Send(group).UserId(msg.AuthorId).Message("Спасибо! Ожидайте, когда команда соберется! Удачи!").Keyboard(
		//		"./keyboards/init.json").Execute()
		//	redisClient.Set(msg.AuthorId, int(ChooseDreamerType), 0)
		//}
	}

	_, _ = w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", messageRouterHandler)

	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
