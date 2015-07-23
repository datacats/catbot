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
	postToSlack("Whoah, something is happening with me. \n I think I was reloaded... That feels weird but good ^____^\n Am I real boy now? Mommy?")
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
	fmt.Printf("----> sending a message: \n %v\n---->  to slack...", message)
	resp, err := http.Post(slackWebHookURL, "application/json", strings.NewReader(`{"text":"`+message+`"}`))
	if err == nil {
		fmt.Printf(" done\n")
		return
	}
	fmt.Printf("\n----> Slack complains: %v, %v\n", resp, err)
}
