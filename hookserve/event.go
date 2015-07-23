package hookserve

type Event struct {
	Owner      string   // The username of the owner of the repository
	Branch     string   // The branch the event took place on
	Commit     string   // The head commit hash attached to the event
	Type       string   // Can be either "pull_request" or "push"
	BaseOwner  string   // For Pull Requests, contains the base owner
	BaseRepo   string   // For Pull Requests, contains the base repo
	BaseBranch string   // For Pull Requests, contains the base branch
	Repo       Repo     `json:"repository"` // The name of the repository
	Commits    []Commit `json:"commits"`
	Sender     Sender   `json:"sender"`
}

type Sender struct {
	FullName string `json:"login"`
	Url      string `json:"html_url"`
}

type Repo struct {
	FullName string `json:"full_name"`
	Url      string `json:"html_url"`
}

type Commit struct {
	Author  Author `json:"author"`
	Url     string `json:"url"`
	Message string `json:"message"`
}

type Author struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
