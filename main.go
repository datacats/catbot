package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/phayes/hookserve/hookserve"
)

var githubSecret = os.Getenv("GITHUB_SECRET")
var slackWebHookURL = os.Getenv("SLACK_HOOK")

func main() {
	fmt.Println("Catbot listens to the whispers of your heart...")
	postToSlack("Hey everybody! This is me, your lovable catbot integration ^____^")
	server := hookserve.NewServer()
	server.Port = 5000
	server.Secret = githubSecret
	server.GoListenAndServe()

	for {
		select {
		case event := <-server.Events:
			go postToSlack(`<@` + event.Owner + " " + event.Repo + " " + event.Branch + " " + event.Commit)
		default:
			time.Sleep(100)
		}
	}
}

func postToSlack(message string) {
	http.Post(slackWebHookURL, "application/json", strings.NewReader(`{"text":"`+message+`"}`))
}
