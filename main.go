package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"vkBot/conversation"
	"vkBot/conversation/domain"
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
		state := conversation.GetDialogState(msg.AuthorId)
		ctx := conversation.DialogContext{
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
