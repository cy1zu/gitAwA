package fetchers

import (
	"backend/app/models"
)

type CommentItems struct {
	Comments []Comment `json:"items"`
}
type Comment struct {
	User        models.MiniDeveloper `json:"user"`
	TimelineUrl string               `json:"timeline_url"`
}
type CommentLines struct {
	Event string               `json:"event" comment:"should be commented"`
	User  models.MiniDeveloper `json:"user"`
	Body  string               `json:"body"`
}

type DeveloperFull struct {
	Login       string      `json:"login"`
	Id          int64       `json:"id"`
	StarredUrl  string      `json:"starred_url"`
	ReposUrl    string      `json:"repos_url"`
	Type        string      `json:"type"`
	Name        string      `json:"name"`
	Company     string      `json:"company"`
	Blog        string      `json:"blog"`
	Location    string      `json:"location"`
	Email       string      `json:"email"`
	Bio         string      `json:"bio"`
	PublicRepos int         `json:"public_repos"`
	AllRepos    []ReposFull `json:"all_repos"`
}

type ReposFull struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login        string `json:"login"`
		Id           int    `json:"id"`
		Type         string `json:"type"`
		UserViewType string `json:"user_view_type"`
	} `json:"owner"`
	Fork            bool              `json:"fork"`
	LanguagesUrl    string            `json:"languages_url"`
	StargazersUrl   string            `json:"stargazers_url"`
	ContributorsUrl string            `json:"contributors_url"`
	Size            int               `json:"size"`
	StargazersCount int               `json:"stargazers_count"`
	Language        string            `json:"language"`
	Visibility      string            `json:"visibility"`
	Parent          *ReposDetailsFull `json:"parent"`
}

type ReposDetailsFull struct {
	Id       int64  `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login string `json:"login"`
		Id    int64  `json:"id"`
		Type  string `json:"type"`
	} `json:"owner"`

	Fork            bool   `json:"fork"`
	StatusesUrl     string `json:"statuses_url"`
	LanguagesUrl    string `json:"languages_url"`
	StargazersUrl   string `json:"stargazers_url"`
	ContributorsUrl string `json:"contributors_url"`
	Size            int64  `json:"size"`
	StargazersCount int64  `json:"stargazers_count"`
	WatchersCount   int    `json:"watchers_count"`
	Language        string `json:"language"`
	Parent          struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login string `json:"login"`
			Id    int64  `json:"id"`
			Type  string `json:"type"`
		} `json:"owner"`
		Fork            bool   `json:"fork"`
		LanguagesUrl    string `json:"languages_url"`
		StargazersUrl   string `json:"stargazers_url"`
		ContributorsUrl string `json:"contributors_url"`
		Size            int64  `json:"size"`
		StargazersCount int64  `json:"stargazers_count"`
		Language        string `json:"language"`
	} `json:"parent"`
	Source struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			UserViewType      string `json:"user_view_type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		Fork            bool   `json:"fork"`
		LanguagesUrl    string `json:"languages_url"`
		StargazersUrl   string `json:"stargazers_url"`
		ContributorsUrl string `json:"contributors_url"`
		Size            int    `json:"size"`
		StargazersCount int    `json:"stargazers_count"`
		Language        string `json:"language"`
	} `json:"source"`
}
