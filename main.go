package main

import (
	service "strava-rundown-api/service"
)

func main() {
	authToken := "TOKEN"
	duration := "WEEK"
	service.GetRunningTotal(duration, authToken)
}
