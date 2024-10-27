package fetchers

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	res, _ := GetUserInfo("fuhaoda")
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposDetail(t *testing.T) {
	res, _ := GetReposDetail("fuhaoda/KataGo")
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposContributors(t *testing.T) {
	res, _ := GetReposContributors("fuhaoda/KataGo")
	str, _ := json.MarshalIndent(*res, "", "\t")
	fmt.Println(string(str))
}
