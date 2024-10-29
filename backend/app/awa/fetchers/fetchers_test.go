package fetchers

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	githubToken := new(string)
	res, _ := GetUserInfo("lightvector", githubToken)
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposDetail(t *testing.T) {
	githubToken := new(string)
	res, _ := GetReposDetail("lightvector/KataGo", githubToken)
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposContributors(t *testing.T) {
	githubToken := new(string)
	res, _ := GetReposContributors("lightvector/KataGo", githubToken)
	str, _ := json.MarshalIndent(*res, "", "\t")
	fmt.Println(string(str))
}
