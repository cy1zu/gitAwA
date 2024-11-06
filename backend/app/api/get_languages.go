package api

import (
	"backend/app/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func GetLanguageListHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "page must be a number",
		})
		return
	}
	lang := c.Query("lang")
	githubToken := c.GetHeader("Authorization")
	if strings.HasPrefix(githubToken, "Bearer ") {
		githubToken = githubToken[7:]
	} else {
		githubToken = ""
	}
	data, err := services.GetLanguageListServices(lang, page)
	if err != nil {
		zap.L().Error("GetLanguageListServices failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "GetLanguageList failed",
		})
		return
	}
	if data != nil {
		c.JSON(http.StatusOK, data)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "language not found",
	})
}
