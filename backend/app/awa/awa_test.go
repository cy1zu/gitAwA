package awa

import (
	"backend/app/awa/fetchers"
	"backend/app/awa/processors"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGuessNationByInfo(t *testing.T) {
	githubToken := new(string)
	*githubToken = "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"
	res, err := fetchers.GetDeveloperInfo("lvr1997", githubToken)
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
