package api

import "github.com/gin-gonic/gin"

func Install(eng *gin.Engine) {
	eng.GET("/api/tab_videos/1", hello)
}

func hello(c *gin.Context) {
	c.JSON(1, "ok")
}
