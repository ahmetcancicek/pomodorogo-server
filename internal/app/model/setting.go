package model

// Setting ...
type Setting struct {
	ID           uint  `json:"id"`
	WorkDuration uint8 `json:"workDuration"`
	ShortBreak   int8  `json:"shortBreak"`
	LongBreak    int8  `json:"longBreak"`
	Rounds       int8  `json:"rounds"`
	UserID       uint
}
