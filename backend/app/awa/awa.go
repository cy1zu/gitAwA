package awa

import (
	"backend/app/awa/fetchers"
	"backend/app/awa/processors"
	"backend/app/models"
	"go.uber.org/zap"
)

func FetchUser(githubId string, githubToken *string) (*models.Developer, error) {
	res, err := fetchers.GetUserInfo(githubId, githubToken)
	if err != nil {
		zap.L().Error("GetUserInfo failed", zap.Error(err))
		return &models.Developer{}, err
	}
	data, err := processors.ParseDevelopersData(res, githubToken)
	if err != nil {
		zap.L().Error("ParseDevelopersData failed", zap.Error(err))
		return &models.Developer{}, err
	}
	developers, err := processors.FinalDevelopers(data)
	if err != nil {
		zap.L().Error("FinalDevelopersData failed", zap.Error(err))
		return &models.Developer{}, err
	}
	return &developers, nil
}

func ValueUserNation(ways string) {

}
