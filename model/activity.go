package model

import "time"

type TrainingDuration int

const (
	Week TrainingDuration = iota
	Month
	Year
)

type ActivityResponse struct {
	ResourceState int `json:"resource_state"`
	Athlete       struct {
		ID            int `json:"id"`
		ResourceState int `json:"resource_state"`
	} `json:"athlete"`
	Name               string        `json:"name"`
	Distance           float64       `json:"distance"`
	MovingTime         int           `json:"moving_time"`
	ElapsedTime        int           `json:"elapsed_time"`
	TotalElevationGain int           `json:"total_elevation_gain"`
	Type               string        `json:"type"`
	WorkoutType        interface{}   `json:"workout_type,omitempty"`
	ID                 int64         `json:"id"`
	ExternalID         string        `json:"external_id"`
	UploadID           int64         `json:"upload_id"`
	StartDate          time.Time     `json:"start_date"`
	StartDateLocal     time.Time     `json:"start_date_local"`
	Timezone           string        `json:"timezone"`
	UtcOffset          int           `json:"utc_offset"`
	StartLatlng        []interface{} `json:"start_latlng"`
	EndLatlng          []interface{} `json:"end_latlng"`
	LocationCity       interface{}   `json:"location_city"`
	LocationState      interface{}   `json:"location_state"`
	LocationCountry    string        `json:"location_country"`
	StartLatitude      interface{}   `json:"start_latitude"`
	StartLongitude     interface{}   `json:"start_longitude"`
	AchievementCount   int           `json:"achievement_count"`
	KudosCount         int           `json:"kudos_count"`
	CommentCount       int           `json:"comment_count"`
	AthleteCount       int           `json:"athlete_count"`
	PhotoCount         int           `json:"photo_count"`
	Map                struct {
		ID              string      `json:"id"`
		SummaryPolyline interface{} `json:"summary_polyline"`
		ResourceState   int         `json:"resource_state"`
	} `json:"map"`
	Trainer                    bool    `json:"trainer"`
	Commute                    bool    `json:"commute"`
	Manual                     bool    `json:"manual"`
	Private                    bool    `json:"private"`
	Visibility                 string  `json:"visibility"`
	Flagged                    bool    `json:"flagged"`
	GearID                     string  `json:"gear_id"`
	FromAcceptedTag            bool    `json:"from_accepted_tag"`
	UploadIDStr                string  `json:"upload_id_str"`
	AverageSpeed               float64 `json:"average_speed"`
	MaxSpeed                   float64 `json:"max_speed"`
	AverageCadence             float64 `json:"average_cadence"`
	HasHeartrate               bool    `json:"has_heartrate"`
	AverageHeartrate           float64 `json:"average_heartrate"`
	MaxHeartrate               float64 `json:"max_heartrate"`
	HeartrateOptOut            bool    `json:"heartrate_opt_out"`
	DisplayHideHeartrateOption bool    `json:"display_hide_heartrate_option"`
	PrCount                    int     `json:"pr_count"`
	TotalPhotoCount            int     `json:"total_photo_count"`
	HasKudoed                  bool    `json:"has_kudoed"`
	ElevHigh                   float64 `json:"elev_high,omitempty"`
	ElevLow                    float64 `json:"elev_low,omitempty"`
}
