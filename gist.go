package gistore

import "time"

// Gist represents a gist, which includes files.
type Gist struct {
	URL         string              `json:"url"`
	ForksURL    string              `json:"forks_url"`
	CommitsURL  string              `json:"commits_url"`
	ID          string              `json:"id"`
	Description string              `json:"description"`
	Public      bool                `json:"public"`
	Owner       GithubUser          `json:"owner"`
	User        GithubUser          `json:"user"`
	Files       map[string]GistFile `json:"files"`
	Truncated   bool                `json:"truncated"`
	Comments    int                 `json:"comments"`
	CommentsURL string              `json:"comments_url"`
	HTMLURL     string              `json:"html_url"`
	GitPullURL  string              `json:"git_pull_url"`
	GitPushURL  string              `json:"git_push_url"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	Forks       []Fork              `json:"forks"`
	History     []History           `json:"history"`
}

// GistFile is a file with content.
type GistFile struct {
	Size      int    `json:"size"`
	RawURL    string `json:"raw_url"`
	Type      string `json:"type"`
	Language  string `json:"language"`
	Truncated bool   `json:"truncated"`
	Content   string `json:"content"`
}

type GithubUser struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Fork struct {
	User      GithubUser `json:"user"`
	URL       string     `json:"url"`
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type History struct {
	URL          string       `json:"url"`
	Version      string       `json:"version"`
	User         GithubUser   `json:"user"`
	ChangeStatus ChangeStatus `json:"change_status"`
	CommittedAt  time.Time    `json:"committed_at"`
}

type ChangeStatus struct {
	Deletions int `json:"deletions"`
	Additions int `json:"additions"`
	Total     int `json:"total"`
}
