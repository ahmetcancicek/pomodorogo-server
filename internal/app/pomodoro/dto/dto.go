package dto

import (
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"time"
)

// PomodoroDTO ...
type PomodoroDTO struct {
	ID         uint      `json:"id"`
	StartTime  time.Time `json:"startTime"`
	FinishTime time.Time `json:"finishTime"`
	TagID      uint      `json:"tagId"`
}

// ToPomodoro ...
func ToPomodoro(pomodoroDTO *PomodoroDTO) *model.Pomodoro {
	return &model.Pomodoro{
		ID:         pomodoroDTO.ID,
		StartTime:  pomodoroDTO.StartTime,
		FinishTime: pomodoroDTO.FinishTime,
		TagID:      pomodoroDTO.TagID,
	}
}

// ToPomodoroDTO
func ToPomodoroDTO(pomodoro *model.Pomodoro) *PomodoroDTO {
	return &PomodoroDTO{
		ID:         pomodoro.ID,
		StartTime:  pomodoro.StartTime,
		FinishTime: pomodoro.FinishTime,
		TagID:      pomodoro.TagID,
	}
}
