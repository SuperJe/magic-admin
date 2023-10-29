package dto

import "time"

type CPPPractice struct {
	ID        int64     `json:"id" gorm:"column:id"`
	ProblemID int64     `json:"problem_id" gorm:"column:p_id"`
	UserID    int64     `json:"user_id" gorm:"column:user_id"`
	Code      string    `json:"code" gorm:"column:code"`
	Updated   time.Time `json:"updated" gorm:"column:updated"`
	Created   time.Time `json:"created" gorm:"column:created"`
}

func (cp *CPPPractice) TableName() string {
	return "practice_cpp"
}

type GetPracticeCodeReq struct {
	IDList string `form:"ids" json:"ids"`
}

type GetPracticeCodeRsp struct {
	BaseRsp
	Codes map[int64]string `json:"codes"`
}

type SubmitPracticeCodeReq struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
}

type SubmitPracticeCodeRsp struct {
	BaseRsp
	Accept bool `json:"accept"`
}
