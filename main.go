package main

import (
	"digester-bot/gptassistant"
	"digester-bot/tbot"
	"log"
	"net/http"
	"os"
)

func main() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
		http.ListenAndServe(":8080", nil)
	}()
	if os.Getenv("CHAT_GPT_KEY") == "" {
		log.Fatal("CHAT_GPT_KEY cannot be empty")
	}

	b, err := tbot.New(os.Getenv("TBOT_TOKEN"), gptassistant.NewGPTBuilder(os.Getenv("CHAT_GPT_KEY")))
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Start()
}
