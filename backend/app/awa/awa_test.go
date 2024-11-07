package awa

import (
	"backend/app/awa/fetchers"
	"backend/app/awa/processors"
	"backend/config"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGuessNationByInfo(t *testing.T) {
	githubToken := new(string)
	*githubToken = config.Conf.TestGithubAccessToken
	res, err := fetchers.GetDeveloperInfo("tiger1103", githubToken)
	if err != nil {
		t.Error(err)
	}
	data, err := processors.ParseDevelopersData(res, githubToken)
	if err != nil {
		t.Error(err)
	}
	developers, err := processors.FinalDevelopers(data)
	developers.Nation = GuessNationByInfo(res, githubToken)
	if err != nil {
		t.Error(err)
	}
	str, _ := json.MarshalIndent(developers, "", "\t")
	fmt.Println(string(str))
}
