package api

import (
	"backend/app/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type ParamSearch struct {
	Login    string `form:"dev"`
	Language string `form:"lang"`
	Nation   string `form:"nation"`
}

func GetDeveloperListHandler(c *gin.Context) {
	var p ParamSearch
	if err := c.ShouldBindQuery(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	githubToken := c.GetHeader("Authorization")
	if strings.HasPrefix(githubToken, "Bearer ") {
		githubToken = githubToken[7:]
	} else {
		githubToken = ""
	}
	data, err := services.GetDeveloperListServices(p.Login, p.Language, p.Nation)
	if err != nil {
		zap.L().Error("GetLanguageListServices failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "GetLanguageList failed",
		})
		return
	}
	if data != nil {
		c.JSON(http.StatusOK, *data)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "no such developer",
	})
}
