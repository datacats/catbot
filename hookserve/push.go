package hookserve

import "github.com/bmatsuo/go-jsontree"

func parsePush(event *Event, request *jsontree.JsonTree) error {
	rawRef, err := request.Get("ref").String()
	if err != nil {
		return err
	}

	event.Branch = rawRef[11:]
	event.Commit, err = request.Get("head_commit").Get("id").String()
	event.Owner, err = request.Get("repository").Get("owner").Get("name").String()

	return nil
}
