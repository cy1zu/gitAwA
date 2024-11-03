package awa

import (
	"backend/app/awa/fetchers"
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

func ValueUserNation(ways string) {

}

type dingdan struct {
	id         int
	user       string
	itemInfoId string
	pingid     string
	nums       int
	pic        string
	wuliuid    string
	wuliu      string
	sstatus    string
	finalprice string
}

type shop struct {
	id       int
	name     string
	pingtype string
	price    int
}
