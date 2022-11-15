package request

import "ahutoj/web/io/constanct"

type AddSubmitReq struct {
	PID        string         `json:"PID"`
	UID        string         `json:"UID"`
	CID        int64          `json:"CID"`
	Source     string         `json:"Source"`
	Lang       constanct.LANG `json:"Lang"`
	SubmitTime int64          `json:"SubmitTime"`
}

type RejudgeSubmitReq struct {
	SID *int64  `json:"SID"`
	PID *string `json:"PID"`
	UID *string `json:"UID"`
	CID *int64  `json:"CID"`
}

type GetSubmitReq struct {
	SID int64 `param:"SID"`
}

type SubmitListReq struct {
	PID    *string             `query:"PID"`
	UID    *string             `query:"UID"`
	CID    *int64              `query:"CID"`
	Result *constanct.OJResult `query:"Result"`
	Lang   *constanct.LANG     `query:"Lang"`
	GetListReq
}

type GetCodeReq struct {
	SID int64 `json:"SID"`
}
