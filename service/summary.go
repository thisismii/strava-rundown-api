package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strava-rundown-api/model"
)

type Activities []model.ActivityResponse

func GetWeeklyRunningTotal() float64 {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/activities", nil)
	req.Header.Set("Authorization", "JWT")
	res, _ := client.Do(req)

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bodyString := string(bodyBytes)

	var activities Activities

	json.Unmarshal([]byte(bodyString), &activities)

	for i, s := range activities {
		fmt.Println(i, s.Distance, s.StartDate, s.StartDateLocal)
	}

	return 0
}
