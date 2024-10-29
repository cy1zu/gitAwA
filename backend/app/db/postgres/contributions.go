package postgres

import (
	"backend/app/models"
	"go.uber.org/zap"
)

func InsertContributions(githubId int, login string, repoId int, fullName string, cons float64, talent float64) error {
	con := models.ContributionsStored{
		DeveloperGithubId: githubId,
		DeveloperLogin:    login,
		RepoGithubId:      repoId,
		RepoFullName:      fullName,
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
