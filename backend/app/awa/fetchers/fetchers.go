package fetchers

import (
	"context"
	"github.com/carlmjohnson/requests"
)

func GetUserInfo(githubId string) (*DeveloperFull, error) {
	var data *DeveloperFull
	err := requests.
		URL("https://api.github.com/users/" + githubId).
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}
	data.AllRepos, err = getUserPublicRepos(githubId, data.PublicRepos)
	if err != nil {
		data.AllRepos = nil
		return data, err
	}

	return data, nil
}

func getUserPublicRepos(githubId string, lens int) (*[]ReposFull, error) {
	data := make([]ReposFull, 0, lens)
	err := requests.URL("https://api.github.com/users/" + githubId + "/repos").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}
	return &data, nil
}
