package postgres

import (
	"backend/app/models"
	"go.uber.org/zap"
)

func InsertContributions(githubId int64, login string, repoId int64, fullName string, cons float64, talent float64) error {
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

func GetContributionsByDeveloper(githubLogin string, offset int, limit int) ([]models.ContributionsStored, error) {
	cons := make([]models.ContributionsStored, 0)
	res := pdb.Limit(limit).Offset(offset).Order("talent_score").Find(&cons, "developer_github_id = ?", githubLogin)
	if res.Error != nil {
		zap.L().Error("get contributions by developer failed", zap.Error(res.Error))
		return []models.ContributionsStored{}, res.Error
	}
	return cons, nil
}
