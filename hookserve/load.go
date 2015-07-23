package hookserve

import (
	"encoding/json"
	"fmt"

	"github.com/bmatsuo/go-jsontree"
)

// LoadEvent takes a JSON and parses it into an Event
func LoadEvent(jsonstring []byte, eventType string) (event *Event, err error) {
	// Parse the request and build the Event
	event = &Event{}
	err = json.Unmarshal(jsonstring, event)
	if err != nil {
		return event, fmt.Errorf("LoadEvent: cant preunmarshall tagged fields %#v, \n[ %v ]", err, jsonstring)
	}

	request := jsontree.New()
	err = request.UnmarshalJSON(jsonstring)
	event.Type = eventType
	event.Repo.FullName, err = request.Get("repository").Get("full_name").String()
	event.Repo.Url, err = request.Get("repository").Get("html_url").String()

	if err != nil {
		return event, fmt.Errorf("LoadEvent: I am not find the repo fields even -  %v", err)
	}

	switch eventType {
	case "push":
		err = parsePush(event, request)
	case "pull_request":
		err = parsePullRequst(event, request)
	}
	return event, err
}
