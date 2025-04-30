package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// function debug
func Debug(obj any) {
	raw, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(raw))
}

// function defie local time
func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

// function convert time string to type time
func ConvertStringTimeToTime(t string) time.Time {
	layout := "2006-01-02T15:04:05.999 -0700 MST"

	result, err := time.Parse(layout, t)
	if err != nil {
		log.Printf("Error: Parse time failed: %s", err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")
	return result.In(loc)
}
