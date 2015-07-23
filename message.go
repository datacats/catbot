package main

import (
	"fmt"

	"github.com/phayes/hookserve/hookserve"
)

var MssgTemplatePullRequest = "%s repo: @%v made a pull request: '%s'! w00t w00t!"
var MssgTemplatePush = "%v repo: @%v pushed '%v' commit to %v branch"

func fmtEventMessage(ev hookserve.Event) string {

	if ev.Type == "pull_request" {
		return fmt.Sprintf(MssgTemplatePullRequest, ev.Repo, ev.BaseOwner, ev.Branch)
	}
	return fmt.Sprintf(MssgTemplatePush, ev.Repo, ev.Owner, ev.Commit, ev.Branch)
}
