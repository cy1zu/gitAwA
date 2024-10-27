package fetchers

import (
	"backend/app/models"
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
	data.AllRepos, err = GetUserPublicRepos(githubId, data.PublicRepos)
	if err != nil {
		data.AllRepos = nil
		return data, err
	}

	return data, nil
}

func GetUserPublicRepos(githubId string, lens int) (*[]ReposFull, error) {
	data := make([]ReposFull, 0, lens)
	err := requests.URL("https://api.github.com/users/" + githubId + "/repos").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetReposDetail(reposFullName string) (*ReposDetailsFull, error) {
	data := new(ReposDetailsFull)
	err := requests.URL("https://api.github.com/repos/" + reposFullName).
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		panic(err)
	}
	return data, nil
}

func GetReposLanguages(reposFullName string) (map[string]int64, error) {
	var data map[string]int64
	err := requests.URL("https://api.github.com/repos/" + reposFullName + "/languages").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetReposContributors(reposFullName string) (*[]models.MiniDeveloper, error) {
	data := make([]models.MiniDeveloper, 0, 16)
	err := requests.URL("https://api.github.com/repos/" + reposFullName + "/contributors").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		panic(err)
	}
	return &data, nil
}
