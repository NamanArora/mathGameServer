package main

import "github.com/gin-gonic/gin"
import rtr "github.com/NamanArora/mathGameServer/router"

func main() {
	router := gin.Default()
	router.POST("/saveScore", rtr.SaveScore)
	router.POST("/findNearby", rtr.FindNearby)
	router.POST("/checkUsername", rtr.CheckUsername)
	router.POST("/getPercentile", rtr.GetPercentile)
	router.GET("/healthCheck", rtr.HealthCheck)
	router.Run()
}
