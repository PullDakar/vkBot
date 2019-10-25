package cache

import (
	"github.com/go-redis/redis"
	"log"
)

var possibleStates = [...]DialogState{
	DialogNotStartedState{},
	DialogStartedState{},
	IdeaBranchStartedState{},
	IdeaDescribedState{},
	AnalystTypeChosenState{},
	SpecialistCountSelectedState{},
	TeamFormedState{},
	DeveloperTypeChosenState{},
	TesterTypeChosenState{},
	BackDeveloperChosenState{},
	FrontDeveloperChosenState{},
	MobileDeveloperChosenState{},
	AddingDreamerStartedState{},
	AddedDreamerLinkState{},
	AddedDreamerMailState{},
	NewDreamerSpecializationSelectedState{},
}

// FIXME переделать чление полей из yml
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6380",
})

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
		log.Panic("Error while setting value to redis cache by key "+key+" with ", err)
	}
}

func indexOf(value DialogState) (int, error) {
	for index, state := range possibleStates {
		if state == value {
			return index, nil
		}
	}

	// FIXME добавить передачу в ошибку конкретный state
	return -1, &stateNotFound{"Dialog state not found "}
}
