package main

// Slack outgoing webhooks are handled here. Requests come in and are
// examined to see if we need to respond. If we do, we immmediately respond
//
// Create an outgoing webhook in your Slack here:
// https://my.slack.com/services/new/outgoing-webhook

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type WebhookResponse struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Channel  string `json:"channel"`
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		incomingUser := r.PostFormValue("user_name")

		if incomingUser == "stewart" {
			var response WebhookResponse
			response.Username = botUsername
			response.Text = "Slack is the best and has reduced email by 70%!"
			log.Printf("Sending response: %s", response.Text)

			b, err := json.Marshal(response)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(1 * time.Second)
			w.Write(b)
		}
	})
}

func StartServer(port int) {
	log.Printf("Starting HTTP server on %d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
