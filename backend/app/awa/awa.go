package awa

import (
	"backend/app/awa/fetchers"
	"backend/app/awa/guessers"
	"backend/app/awa/processors"
	"backend/app/db/postgres"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func FetchDeveloper(githubId string, githubToken *string) {
	_, ok := postgres.CacheDevelopersSet.Load(githubId)
	if ok {
		return
	}
	postgres.CacheDevelopersSet.Store(githubId, postgres.DataProcessing)
	startTime := time.Now()
	fmt.Printf("awa %s now!\n", githubId)
	res, err := fetchers.GetDeveloperInfo(githubId, githubToken)
	if err != nil {
		zap.L().Error("GetDeveloperInfo failed", zap.Error(err))
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
	developers.Nation = GuessNationByInfo(res, githubToken)
	err = postgres.InsertDeveloper(&developers)
	if err != nil {
		zap.L().Error("postgres.InsertDeveloper failed", zap.Error(err))
		postgres.CacheDevelopersSet.Delete(githubId)
		return
	}
	// insert cache
	err = postgres.CacheInsertDeveloper(&developers)
	if err != nil {
		zap.L().Error("cache developer failed", zap.Error(err),
			zap.String("github_login", developers.Login))
		return
	}
	fmt.Printf("awa %s use %.2fs to done\n", githubId, float64(time.Since(startTime)/time.Millisecond)/1000)
}

func GuessNationByInfo(dev *fetchers.DeveloperFull, githubToken *string) string {
	comments, err := fetchers.GetDeveloperComments(dev.Login, githubToken)
	if err != nil {
		zap.L().Error("FetchDeveloperComments failed", zap.Error(err))
		return ""
	}
	query := map[string]interface{}{
		"url":      "https://github.com/" + dev.Login,
		"Name":     dev.Name,
		"Location": dev.Location,
		"Company":  dev.Company,
		"Blog":     dev.Blog,
		"Email":    dev.Email,
		"Comments": *comments,
	}
	head, err := guessers.Init()
	if err != nil {
		zap.L().Error("init guesser failed", zap.Error(err))
		return ""
	}

	// TODO: guess nation by info use llm
	res, err := guessers.GuessNation(head, query)
	if err != nil {
		zap.L().Error("guess nation failed", zap.Error(err))
		return ""
	}
	if res.Value < 0.5 {
		return "N/A"
	}
	return res.Nation
}

/*
	{
		"Name":     Wild Heart,
		Location: Budapest,
		Company:  Jimi System,
		Blog:     ,
		Email:    ,
	}
*/
