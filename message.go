package main

import (
	"fmt"
	"strings"

	"github.com/dborzov/catbot/hookserve"
)

var MssgTemplatePullRequest = "%s repo: @%v made a pull request: '%s'! w00t w00t!"
var MssgTemplatePush = "%v repo: @%v pushed '%v' commit to %v branch"
var MssgUnsupportedEvent = "%v repo: whoah, there was an event with type \"%s\", I dunno what that means :P"

func fmtEventMessage(ev hookserve.Event) (msg string) {
	if ev.Type == "ping" {
		msg += fmt.Sprintf("repo *<%v|%v>* tickles me! I feel special :blush: \n ", ev.Repo.Url, ev.Repo.FullName)
		msg += fmt.Sprintf("@<%v|%v> set up a webhook for that repo. \n", ev.Sender.Url, ev.Sender.FullName)
		return
	}
	if ev.Type == "watch" {
		msg += fmt.Sprintf("repo *<%v|%v>* was *%v* ", ev.Repo.Url, ev.Repo.FullName, ev.Action)
		msg += fmt.Sprintf(" by @<%v|%v>! \n", ev.Sender.Url, ev.Sender.FullName)
		return
	}
	msg = fmt.Sprintf("*<%v|%v>*:*%v* (%v) \n", ev.Repo.Url, ev.Repo.FullName, ev.Branch, niceTypeMessage(ev.Type))
	for _, commit := range ev.Commits {
		// escaping msg as documented at https://api.slack.com/docs/formatting
		commit.Message = strings.Replace(commit.Message, "&", `&amp;`, -1)
		commit.Message = strings.Replace(commit.Message, ">", `&gt;`, -1)
		commit.Message = strings.Replace(commit.Message, "<", `&lt;`, -1)
		msg += fmt.Sprintf("   *[<%v|%v>]* <%v|%v>\n", commit.Url, commit.Author.Username, commit.Url, commit.Message)
	}
	if ev.Branch != "" {
		msg += fmt.Sprintf("The branch is %v \n", ev.Branch)
	}
	if ev.Sender.FullName != "" {
		msg += fmt.Sprintf("Authored by @<%v|%v>. \n", ev.Sender.Url, ev.Sender.FullName)
	}
	return
}

func niceTypeMessage(msgType string) string {
	switch msgType {
	case "pull_request":
		return "Pull Request Posted!"
	case "push":
		return "*New Commits*"
	case "ping":
		return "is pinging me. I feel special :blush:"
	}
	return "Event with status'" + msgType + "' happened"
}

func slackLink(title, href string) string {
	return fmt.Sprintf("<%v|%v>", href, title)
}
