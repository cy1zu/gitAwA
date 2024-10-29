package fetchers

import (
	"backend/app/models"
	"context"
	"github.com/carlmjohnson/requests"
	"go.uber.org/zap"
)

func GetUserInfo(githubId string, githubToken *string) (*DeveloperFull, error) {
	var data *DeveloperFull
	err := requests.
		URL("https://api.github.com/users/"+githubId).
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		zap.L().Error("Error fetching developer info", zap.Error(err))
		return nil, err
	}
	data.AllRepos, err = GetUserPublicRepos(githubId, data.PublicRepos, githubToken)
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

func GetUserPublicRepos(githubId string, lens int, githubToken *string) ([]ReposFull, error) {
	data := make([]ReposFull, 0, lens)
	err := requests.URL("https://api.github.com/users/"+githubId+"/repos").
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetReposDetail(reposFullName string, githubToken *string) (*ReposDetailsFull, error) {
	data := new(ReposDetailsFull)
	err := requests.URL("https://api.github.com/repos/"+reposFullName).
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(data).
		Fetch(context.Background())
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetReposLanguages(reposFullName string, githubToken *string) (map[string]int64, error) {
	var data map[string]int64
	err := requests.URL("https://api.github.com/repos/"+reposFullName+"/languages").
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetReposContributors(reposFullName string, githubToken *string) (*[]models.MiniDeveloper, error) {
	data := make([]models.MiniDeveloper, 0, 16)
	err := requests.URL("https://api.github.com/repos/"+reposFullName+"/contributors").
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return &data, err
	}
	return &data, nil
}
