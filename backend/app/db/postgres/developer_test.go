package postgres

import (
	"backend/app/models"
	"backend/config"
	"encoding/json"
	"os"
	"testing"
)

func TestInsertUser(t *testing.T) {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	Init("test")
	dev := new(models.Developer)
	file, err := os.ReadFile("./mock/lightvector_dev.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, dev)
	if err != nil {
		panic(err)
	}
	err = InsertDeveloper(dev)
	if err != nil {
		t.Error(err)
	}

}
