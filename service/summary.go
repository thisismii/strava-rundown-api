package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strava-rundown-api/model"
	"time"
)

type Activities []model.ActivityResponse

func GetRunningTotal(duration string, authToken string) float64 {

	activities := callStravaActivites(authToken)
	totalDistance := 0.0

	startTime := retrieveStartDate(duration)

	fmt.Println(startTime)

	for _, s := range activities {
		if s.StartDate.After(startTime) {
			totalDistance += s.Distance
		} else {
			break
		}
	}

	fmt.Println(totalDistance)

	return totalDistance
}

func callStravaActivites(authToken string) Activities {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/activities", nil)
	req.Header.Set("Authorization", authToken)
	res, _ := client.Do(req)

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bodyString := string(bodyBytes)

	var activities Activities

	json.Unmarshal([]byte(bodyString), &activities)

	return activities
}

func retrieveStartDate(duration string) time.Time {
	currentTime := time.Now().UTC()
	currentTime = currentTime.Truncate(24 * time.Hour)
	startTime := currentTime
	
	if duration == "WEEK" {
		// first day of the training week
		weekday := int(currentTime.Weekday())
		startTime = currentTime.AddDate(0, 0, -(weekday))
	} else if duration == "MONTH" {
		// first day of the current month
		days := currentTime.Day() - 1
		startTime = currentTime.AddDate(0, 0, -(days))
	} else if duration == "YEAR" {
		// first day of the current year
		days := currentTime.YearDay() - 1
		startTime = currentTime.AddDate(0, 0, -(days))
	}

	return startTime
}
