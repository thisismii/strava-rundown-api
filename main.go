package main

import (
	"github.com/gin-gonic/gin"

	"net/http"

	service "strava-rundown-api/service"
)

type DistanceResponse struct {
	Duration  string  `json:"duration"`
	Distance  float64 `json:"distance"`
}

func main() {
	router := gin.Default()
	router.GET("/distance", getDistance)
	router.Run("localhost:3000")
}

func getDistance(context *gin.Context) {
	authToken := context.Request.Header.Get("Authorization")
	duration := context.Request.URL.Query().Get("training-duration")
	distance := service.GetRunningTotal(duration, authToken)

	resp := DistanceResponse{Duration: duration, Distance: distance}

	context.IndentedJSON(http.StatusOK, resp)
}