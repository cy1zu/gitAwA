package fetchers

import (
	"backend/config"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	githubToken := config.Conf.TestGithubAccessToken
	res, _ := GetDeveloperInfo("lightvector", &githubToken)
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposDetail(t *testing.T) {
	githubToken := config.Conf.TestGithubAccessToken
	res, _ := GetReposDetail("lightvector/KataGo", &githubToken)
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposContributors(t *testing.T) {
	githubToken := config.Conf.TestGithubAccessToken
	res, _ := GetReposContributors("lightvector/KataGo", &githubToken)
	str, _ := json.MarshalIndent(*res, "", "\t")
	fmt.Println(string(str))
}

func TestGetDeveloperComments(t *testing.T) {
	githubToken := config.Conf.TestGithubAccessToken
	res, _ := GetDeveloperComments("tiger1103", &githubToken)
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}
