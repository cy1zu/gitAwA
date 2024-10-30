package api

import (
	"backend/app/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDeveloperHandler(c *gin.Context) {
	developerLogin := c.Param("developerLogin")
	githubToken := c.Request.Header.Get("x-github-token")
	if developerLogin == "" {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	data, err := services.GetDeveloperServices(developerLogin, &githubToken)
	if errors.Is(err, services.ErrorDataNeedFetch) || errors.Is(err, services.ErrorDataProcessing) {
		c.JSON(http.StatusOK, err)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetDeveloperListHandler(c *gin.Context) {

}
