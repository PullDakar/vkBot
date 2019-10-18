package conversation

import (
	"vkBot/domain"
)

var permissibleValuesModel = map[string]Validator{
	"Начать": {
		validMessages:  []string{"Предложить идею", "Добавить участника"},
		message:        "",
		pathToKeyboard: "",
	},
	"Предложить идею": {},
	"Аналитик":        {"Системный", "Бизнес"},
	"Системный":       {"1", "2", "3"},
	"Бизнес":          {"1", "2", "3"},

	"Разработчик": {"Бэк", "Фронт", "Мобильный"},

	"Бэк":    {"Java", "Python", "Golang"},
	"Java":   {"1", "2", "3"},
	"Python": {"1", "2", "3"},
	"Golang": {"1", "2", "3"},

	"Фронт":   {"Angular", "React"},
	"Angular": {"1", "2", "3"},
	"React":   {"1", "2", "3"},

	"Мобильный": {"Android", "IOS"},
	"Android":   {"1", "2", "3"},
	"IOS":       {"1", "2", "3"},

	"Тестировщик": {"Ручной", "Авто"},
	"Ручной":      {"1", "2", "3"},
	"Авто":        {"1", "2", "3"},

	"Менеджер проектов": {"1", "2", "3"},
}

type Validator struct {
	validMessages  []string
	message        string
	pathToKeyboard string
}

type Chain struct {
	prevMessageText string
	currMessageText string
}

type MessageService struct {
	Message domain.Message
}

func (ms MessageService) Process() {
	if isChain(ms.Message) {

	}
}

var chain = &Chain{}

func isChain(msg domain.Message) bool {
	if msg.Type == domain.MessageNew {
		chain.currMessageText = msg.Text
		if chain.prevMessageText == "" {
			if chain.currMessageText == "Начать" {
				chain.prevMessageText = chain.currMessageText
				return true
			}
		} else {
			if len(permissibleValuesModel[chain.prevMessageText]) == 0 {
				return true
			}

			if containsInArray(permissibleValuesModel[chain.prevMessageText], chain.currMessageText) {
				chain.prevMessageText = chain.currMessageText
				return true
			}
		}
	}

	return false
}

func containsInArray(arr []string, value string) bool {
	for e := range arr {
		if arr[e] == value {
			return true
		}
	}

	return false
}
