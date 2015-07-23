package main

import (
	"fmt"

	"github.com/dborzov/hookserve/hookserve"
)

var MssgTemplatePullRequest = "%s repo: @%v made a pull request: '%s'! w00t w00t!"
var MssgTemplatePush = "%v repo: @%v pushed '%v' commit to %v branch"
var MssgUnsupportedEvent = "%v repo: whoah, there was an event with type \"%s\", I dunno what that means :P"

func fmtEventMessage(ev hookserve.Event) (msg string) {
	msg = fmt.Sprintf("*<%v|%v>* : %v ^____^ \n", ev.Repo.Url, ev.Repo.FullName, niceTypeMessage(ev.Type))
	for _, commit := range ev.Commits {
		msg += fmt.Sprintf("   [Commit] *<%v|%v>*\n", commit.Url, commit.Message)
	}
	msg += fmt.Sprintf("The branch is %v \n", ev.Branch)
	msg += fmt.Sprintf("Cudos to <%v|%v>. That fella is my favourite :blush: :blush: \n", ev.Sender.Url, ev.Sender.FullName)
	return
}

func niceTypeMessage(msgType string) string {
	switch msgType {
	case "pull_request":
		return "Pull Request Posted!"
	case "push":
		return "Commits Pushed!"
	}
	return "Event with status'" + msgType + "' happened"

}
