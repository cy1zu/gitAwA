package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type v struct {
	Pong map[string]int64 `json:"pong"`
}

func TestHandler(c *gin.Context) {
	c.JSON(200, v{
		Pong: map[string]int64{"test": 222},
	})
	str, _ := json.Marshal(v{Pong: map[string]int64{"test": 222}})
	println(string(str))
}
