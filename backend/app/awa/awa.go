package awa

import (
	"backend/app/awa/fetchers"
	"backend/app/awa/processors"
	"backend/app/db/postgres"
	"go.uber.org/zap"
)

func FetchDeveloper(githubId string, githubToken *string) {
	_, ok := postgres.CacheDevelopersSet.Load(githubId)
	if ok {
		return
	}
	postgres.CacheDevelopersSet.Store(githubId, postgres.DataProcessing)
	res, err := fetchers.GetUserInfo(githubId, githubToken)
	if err != nil {
		zap.L().Error("GetUserInfo failed", zap.Error(err))
		postgres.CacheDevelopersSet.Delete(githubId)
		return
	}
	data, err := processors.ParseDevelopersData(res, githubToken)
	if err != nil {
		zap.L().Error("ParseDevelopersData failed", zap.Error(err))
		postgres.CacheDevelopersSet.Delete(githubId)
		return
	}
	developers, err := processors.FinalDevelopers(data)
	if err != nil {
		zap.L().Error("FinalDevelopersData failed", zap.Error(err))
		postgres.CacheDevelopersSet.Delete(githubId)
		return
	}
	err = postgres.InsertDeveloper(&developers)
	if err != nil {
		zap.L().Error("postgres.InsertDeveloper failed", zap.Error(err))
		postgres.CacheDevelopersSet.Delete(githubId)
		return
	}
}

func ValueUserNation(ways string) {

}
