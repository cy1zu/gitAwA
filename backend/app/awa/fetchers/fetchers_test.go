package fetchers

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	res, _ := GetUserInfo("lightvector")
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposDetail(t *testing.T) {
	res, _ := GetReposDetail("lightvector/KataGo")
	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(str))
}

func TestGetReposContributors(t *testing.T) {
	res, _ := GetReposContributors("lightvector/KataGo")
	str, _ := json.MarshalIndent(*res, "", "\t")
	fmt.Println(string(str))
}
