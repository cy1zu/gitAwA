package processors

import (
	"backend/app/awa/fetchers"
	"encoding/json"
	"fmt"
	"testing"
)

func TestFinalDevelopers(t *testing.T) {
	githubToken := new(string)
	*githubToken = "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"
	res, err := fetchers.GetDeveloperInfo("Reinhare", githubToken)
	if err != nil {
		t.Error(err)
	}
	data, err := ParseDevelopersData(res, githubToken)
	if err != nil {
		t.Error(err)
	}
	developers, err := FinalDevelopers(data)
	if err != nil {
		t.Error(err)
	}
	str, _ := json.MarshalIndent(developers, "", "\t")
	fmt.Println(string(str))
}

func TestParseDevelopersData(t *testing.T) {
	githubToken := new(string)
	res, err := fetchers.GetDeveloperInfo("lightvector", githubToken)
	if err != nil {
		t.Error(err)
	}
	data, err := ParseDevelopersData(res, githubToken)
	if err != nil {
		t.Error(err)
	}
	str, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(str))
}
