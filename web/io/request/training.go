package request

type List struct {
	LID       int64  `json:"LID"`
	UID       string `json:"UID"`
	Title     string `json:"Title"`
	StartTime int64  `json:"StartTime"`
}

type ListProblem struct {
	LID int64  `json:"LID"`
	PID string `json:"PID"`
}

type ListAll struct {
	LID       int64  `json:"LID"`
	UID       string `json:"UID"`
	PID       string `json:"PID"`
	Title     string `json:"Title"`
	StartTime int64  `json:"StartTime"`
}
type TrainingListReq GetListReq
