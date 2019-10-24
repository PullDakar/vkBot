package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"vkBot/cache"
	"vkBot/cache/domain"
)

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
		ctx := cache.DialogContext{
			InputMessage: msg,
			CurrentState: state,
		}
		ctx.Reply()
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
