package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dborzov/hookserve/hookserve"
)

var githubSecret = os.Getenv("GITHUB_SECRET")
var slackWebHookURL = os.Getenv("SLACK_HOOK")

func main() {
	fmt.Println("Catbot listens to the whispers of your heart...")
	postToSlack("*reloaded*: Whoah, I was *reloaded*. That feels weird ^____^\n Am I real boy now? Mommy?")
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
	http.Post(slackWebHookURL, "application/json", strings.NewReader(`{"text":"`+message+`"}`))
}
