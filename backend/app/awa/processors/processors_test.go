package processors

import (
	"backend/app/awa/fetchers"
	"encoding/json"
	"fmt"
	"testing"
)

func TestFinalDevelopers(t *testing.T) {
	res, err := fetchers.GetUserInfo("lightvector")
	if err != nil {
		t.Error(err)
	}
	data, err := ParseDevelopersData(res)
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
	res, err := fetchers.GetUserInfo("lightvector")
	if err != nil {
		t.Error(err)
	}
	data, err := ParseDevelopersData(res)
	if err != nil {
		t.Error(err)
	}
	str, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(str))
}
