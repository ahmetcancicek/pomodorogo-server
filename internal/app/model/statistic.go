package model

import "time"

// Statistics ...
type Statistic struct {
	ID         uint      `json:"id"`
	StartTime  time.Time `json:"startTime"`
	FinishTime time.Time `json:"finishTime"`
	TagID      uint
}
