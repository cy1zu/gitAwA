package api

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	//业务路由
	v1 := r.Group("/api")
	v1.GET("health", GetHealthHandler)

	{
		// v1.GET("developers", GetDeveloperListHandler)
		v1.GET("developers/:developerLogin", GetDeveloperHandler)
	}

	return r
}
