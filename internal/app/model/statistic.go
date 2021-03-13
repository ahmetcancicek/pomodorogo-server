package model

import "time"

// Statistics ...
type Statistic struct {
	ID         uint      `json:"id"`
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
	TagID      uint
}
