package fetchers

import (
	"backend/app/models"
	"context"
	"github.com/carlmjohnson/requests"
	"go.uber.org/zap"
)

func GetUserInfo(githubId string) (*DeveloperFull, error) {
	var data *DeveloperFull
	err := requests.
		URL("https://api.github.com/users/" + githubId).
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		zap.L().Error("Error fetching developer info", zap.Error(err))
		return nil, err
	}
	data.AllRepos, err = GetUserPublicRepos(githubId, data.PublicRepos)
	if err != nil {
		zap.L().Error("GetUserPublicRepos failed", zap.Error(err))
		zap.L().Debug("GetUserPublicRepos failed", zap.Error(err),
			zap.String("githubId", githubId))
		allRepos := make([]ReposFull, 0)
		data.AllRepos = allRepos
		return data, err
	}
	return data, nil
}

func GetUserPublicRepos(githubId string, lens int) ([]ReposFull, error) {
	data := make([]ReposFull, 0, lens)
	err := requests.URL("https://api.github.com/users/" + githubId + "/repos").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetReposDetail(reposFullName string) (*ReposDetailsFull, error) {
	data := new(ReposDetailsFull)
	err := requests.URL("https://api.github.com/repos/" + reposFullName).
		ToJSON(data).
		Fetch(context.Background())
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetReposLanguages(reposFullName string) (map[string]int64, error) {
	var data map[string]int64
	err := requests.URL("https://api.github.com/repos/" + reposFullName + "/languages").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetReposContributors(reposFullName string) (*[]models.MiniDeveloper, error) {
	data := make([]models.MiniDeveloper, 0, 16)
	err := requests.URL("https://api.github.com/repos/" + reposFullName + "/contributors").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return &data, err
	}
	return &data, nil
}
