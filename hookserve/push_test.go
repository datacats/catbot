package hookserve

import (
	"fmt"
	"testing"
)

var examplePushJSON = `{
  "ref": "refs/heads/master",
  "before": "57568bef36df8b0b4ef4e11087885dc2976dbd47",
  "after": "8e31ae47a079515c1c415ad259172185f15469fd",
  "created": false,
  "deleted": false,
  "forced": false,
  "base_ref": null,
  "compare": "https://github.com/dborzov/lsp/compare/57568bef36df...8e31ae47a079",
  "commits": [
    {
      "id": "8e31ae47a079515c1c415ad259172185f15469fd",
      "distinct": true,
      "message": "Update 1.0 TODO Roadmap",
      "timestamp": "2015-07-22T17:03:20-04:00",
      "url": "https://github.com/dborzov/lsp/commit/8e31ae47a079515c1c415ad259172185f15469fd",
      "author": {
        "name": "Dmitry Borzov",
        "email": "tihoutrom@gmail.com",
        "username": "dborzov"
      },
      "committer": {
        "name": "Dmitry Borzov",
        "email": "tihoutrom@gmail.com",
        "username": "dborzov"
      },
      "added": [

      ],
      "removed": [

      ],
      "modified": [
        "README.md"
      ]
    }
  ],
  "head_commit": {
    "id": "8e31ae47a079515c1c415ad259172185f15469fd",
    "distinct": true,
    "message": "Update 1.0 TODO Roadmap",
    "timestamp": "2015-07-22T17:03:20-04:00",
    "url": "https://github.com/dborzov/lsp/commit/8e31ae47a079515c1c415ad259172185f15469fd",
    "author": {
      "name": "Dmitry Borzov",
      "email": "tihoutrom@gmail.com",
      "username": "dborzov"
    },
    "committer": {
      "name": "Dmitry Borzov",
      "email": "tihoutrom@gmail.com",
      "username": "dborzov"
    },
    "added": [

    ],
    "removed": [

    ],
    "modified": [
      "README.md"
    ]
  },
  "repository": {
    "id": 22778688,
    "name": "lsp",
    "full_name": "dborzov/lsp",
    "owner": {
      "name": "dborzov",
      "email": "tihoutrom@gmail.com"
    },
    "private": false,
    "html_url": "https://github.com/dborzov/lsp",
    "description": "lsp is like ls command but more human-friendly",
    "fork": false,
    "url": "https://github.com/dborzov/lsp",
    "forks_url": "https://api.github.com/repos/dborzov/lsp/forks",
    "keys_url": "https://api.github.com/repos/dborzov/lsp/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/dborzov/lsp/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/dborzov/lsp/teams",
    "hooks_url": "https://api.github.com/repos/dborzov/lsp/hooks",
    "issue_events_url": "https://api.github.com/repos/dborzov/lsp/issues/events{/number}",
    "events_url": "https://api.github.com/repos/dborzov/lsp/events",
    "assignees_url": "https://api.github.com/repos/dborzov/lsp/assignees{/user}",
    "branches_url": "https://api.github.com/repos/dborzov/lsp/branches{/branch}",
    "tags_url": "https://api.github.com/repos/dborzov/lsp/tags",
    "blobs_url": "https://api.github.com/repos/dborzov/lsp/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/dborzov/lsp/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/dborzov/lsp/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/dborzov/lsp/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/dborzov/lsp/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/dborzov/lsp/languages",
    "stargazers_url": "https://api.github.com/repos/dborzov/lsp/stargazers",
    "contributors_url": "https://api.github.com/repos/dborzov/lsp/contributors",
    "subscribers_url": "https://api.github.com/repos/dborzov/lsp/subscribers",
    "subscription_url": "https://api.github.com/repos/dborzov/lsp/subscription",
    "commits_url": "https://api.github.com/repos/dborzov/lsp/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/dborzov/lsp/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/dborzov/lsp/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/dborzov/lsp/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/dborzov/lsp/contents/{+path}",
    "compare_url": "https://api.github.com/repos/dborzov/lsp/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/dborzov/lsp/merges",
    "archive_url": "https://api.github.com/repos/dborzov/lsp/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/dborzov/lsp/downloads",
    "issues_url": "https://api.github.com/repos/dborzov/lsp/issues{/number}",
    "pulls_url": "https://api.github.com/repos/dborzov/lsp/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/dborzov/lsp/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/dborzov/lsp/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/dborzov/lsp/labels{/name}",
    "releases_url": "https://api.github.com/repos/dborzov/lsp/releases{/id}",
    "created_at": 1407558118,
    "updated_at": "2015-07-22T10:05:15Z",
    "pushed_at": 1437599003,
    "git_url": "git://github.com/dborzov/lsp.git",
    "ssh_url": "git@github.com:dborzov/lsp.git",
    "clone_url": "https://github.com/dborzov/lsp.git",
    "svn_url": "https://github.com/dborzov/lsp",
    "homepage": "",
    "size": 1759,
    "stargazers_count": 295,
    "watchers_count": 295,
    "language": "Go",
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 9,
    "mirror_url": null,
    "open_issues_count": 9,
    "forks": 9,
    "open_issues": 9,
    "watchers": 295,
    "default_branch": "master",
    "stargazers": 295,
    "master_branch": "master"
  },
  "pusher": {
    "name": "dborzov",
    "email": "tihoutrom@gmail.com"
  },
  "sender": {
    "login": "dborzov",
    "id": 816895,
    "avatar_url": "https://avatars.githubusercontent.com/u/816895?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/dborzov",
    "html_url": "https://github.com/dborzov",
    "followers_url": "https://api.github.com/users/dborzov/followers",
    "following_url": "https://api.github.com/users/dborzov/following{/other_user}",
    "gists_url": "https://api.github.com/users/dborzov/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/dborzov/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/dborzov/subscriptions",
    "organizations_url": "https://api.github.com/users/dborzov/orgs",
    "repos_url": "https://api.github.com/users/dborzov/repos",
    "events_url": "https://api.github.com/users/dborzov/events{/privacy}",
    "received_events_url": "https://api.github.com/users/dborzov/received_events",
    "type": "User",
    "site_admin": false
  }
}
`

func TestPushParser(t *testing.T) {
	event, err := LoadEvent([]byte(examplePushJSON), "push")
	if err != nil {
		t.Error("Load Event failed: ", err)
	}
	fmt.Printf("The result is: %#v \n", event)
}
