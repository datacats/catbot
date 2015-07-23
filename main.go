package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dborzov/catbot/hookserve"
)

var githubSecret = os.Getenv("GITHUB_SECRET")
var slackWebHookURL = os.Getenv("SLACK_HOOK")

func main() {
	fmt.Println("Catbot listens to the whispers of your heart...")
	postToSlack("I am back on. *❤* you guys :blush:")
	server := hookserve.NewServer()
	server.Secret = githubSecret
	server.GoListenAndServe()

	for {
		select {
		case event := <-server.Events:
			fmt.Println("Received event: ", event)
			go postToSlack(fmtEventMessage(event))
		default:
			time.Sleep(100)
		}
	}
}

func postToSlack(message string) {
	type responseFormat struct {
		Message string `json:"text"`
	}

	payload := responseFormat{
		Message: message,
	}

	pb, err := json.Marshal(payload)
	fmt.Printf("----> sending a message: \n %v\n---->  to slack...", message)
	resp, err := http.Post(slackWebHookURL, "application/json", bytes.NewBuffer(pb))
	if err == nil {
		fmt.Printf(" done\n")
	}
	fmt.Printf("\n----> Slack complains: %v, %v\n", resp, err)
}
