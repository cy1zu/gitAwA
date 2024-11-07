package postgres

import (
	"backend/app/models"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InsertContributions(githubId int64, login string, fork bool, repoId int64, fullName string, stars int64, cons float64, talent float64) error {
	con := models.ContributionsStored{
		DeveloperGithubId: githubId,
		DeveloperLogin:    login,
		Fork:              fork,
		RepoGithubId:      repoId,
		RepoFullName:      fullName,
		StargazersCount:   stars,
		Contributions:     cons,
		TalentScore:       talent,
	}
	res := pdb.Create(&con)
	if res.Error != nil {
		zap.L().Error("insert contributions failed", zap.Error(res.Error))
		zap.L().Debug("insert contributions failed", zap.Error(res.Error),
			zap.String("github_login", login),
			zap.String("repo_full_name", fullName))
		return res.Error
	}
	return nil
}

func GetContributionsByDeveloper(githubLogin string) ([]models.ContributionsStored, error) {
	var cons []models.ContributionsStored
	res := pdb.Order("talent_score desc").Find(&cons, "developer_login = ?", githubLogin)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		zap.L().Error("get contributions by developer failed", zap.Error(res.Error))
		return []models.ContributionsStored{}, res.Error
	}
	return cons, nil
}
