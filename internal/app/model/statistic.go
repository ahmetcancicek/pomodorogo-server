package model

import "time"

// Pomodoro ...
type Pomodoro struct {
	ID         uint      `json:"id"`
	StartTime  time.Time `json:"startTime"`
	FinishTime time.Time `json:"finishTime"`
	TagID      uint
}
