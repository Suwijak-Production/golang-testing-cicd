package main

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/app"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app.Run()
}
