package fetchers

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	githubToken := "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"
	res, _ := GetDeveloperInfo("lightvector", &githubToken)
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposDetail(t *testing.T) {
	githubToken := "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"
	res, _ := GetReposDetail("lightvector/KataGo", &githubToken)
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposContributors(t *testing.T) {
	githubToken := "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"
	res, _ := GetReposContributors("lightvector/KataGo", &githubToken)
	str, _ := json.MarshalIndent(*res, "", "\t")
	fmt.Println(string(str))
}

func TestGetDeveloperComments(t *testing.T) {
	githubToken := "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"
	res, _ := GetDeveloperComments("Reinhare", &githubToken)
	str, _ := json.MarshalIndent(*res, "", "\t")
	fmt.Println(string(str))
}
