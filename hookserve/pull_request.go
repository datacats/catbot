package hookserve

import "github.com/bmatsuo/go-jsontree"

func parsePullRequst(event *Event, request *jsontree.JsonTree) (err error) {
	// Fill in values
	event.Owner, err = request.Get("pull_request").Get("head").Get("repo").Get("owner").Get("login").String()
	event.Branch, err = request.Get("pull_request").Get("head").Get("ref").String()
	event.Commit, err = request.Get("pull_request").Get("head").Get("sha").String()
	event.BaseOwner, err = request.Get("pull_request").Get("base").Get("repo").Get("owner").Get("login").String()
	event.BaseRepo, err = request.Get("pull_request").Get("base").Get("repo").Get("name").String()
	event.BaseBranch, err = request.Get("pull_request").Get("base").Get("ref").String()

	return nil
}
