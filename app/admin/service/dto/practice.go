package dto

import (
	"time"
)

type CPPPractice struct {
	ID         int64     `json:"id" gorm:"column:id"`
	ProblemID  int64     `json:"problem_id" gorm:"column:p_id"`
	UserID     int64     `json:"user_id" gorm:"column:user_id"`
	Code       string    `json:"code" gorm:"column:code"`
	LastStatus int32     `json:"last_status" gorm:"column:last_status"`
	Updated    time.Time `json:"updated" gorm:"column:updated"`
	Created    time.Time `json:"created" gorm:"column:created"`
}

type Questions struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Options       string `json:"options"`
	CorrectOption int64  `json:"correct_option"`
	Tag           string `json:"tag"`
	Score         int64  `json:"score"`
}

func (cp *CPPPractice) TableName() string {
	return "practice_cpp"
}

func (cp *CPPPractice) LastStatusMsg() string {
	if cp == nil {
		return ""
	}
	if cp.LastStatus > 0 {
		return "此份代码已通过。"
	}
	return "此份代码未通过。"
}

type GetPracticeCodeReq struct {
	IDList   string `form:"ids" json:"ids"`
	UserID   int64  `json:"user_id"`
	Username string `form:"user_name" json:"user_name"`
}

type LastSubmitDetail struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
type QuestionDetail struct {
	Title         string `json:"title"`
	Option        string `json:"option"`
	CorrectOption int64  `json:"correct_option"`
	Tag           string `json:"tag"`
	Score         int64  `json:"score"`
}

type GetPracticeCodeRsp struct {
	BaseRsp
	Details map[int64]*LastSubmitDetail `json:"details"`
}

type SubmitPracticeCodeReq struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	UserID   int64  `json:"user_ id"`
	Username string `form:"username" json:"username"`
	Lang     string `json:"lang"`
}

type SubmitPracticeCodeRsp struct {
	BaseRsp
	Accept bool `json:"accept"`
}

type GetQuestionsReq struct {
	IDList string `form:"ids" json:"ids"`
}

type GetQuestionsRsp struct {
	BaseRsp
	Questions map[int64]*QuestionDetail `json:"questions"`
}

type QuestionSubmitReq struct {
	Title         string `json:"title"`
	Options       string `json:"options"`
	CorrectOption int64  `json:"correct_option"`
	Tag           string `json:"tag"`
	Score         int64  `json:"score"`
}

type QuestionSubmitRsp struct {
	BaseRsp
}

type GetTestReq struct {
	Id int64 `form:"id" json:"id"`
}

type GetTestRsp struct {
	BaseRsp
	Title string `json:"title"`
	Pids  string `json:"pids"`
}

type AddCodeProblemReq struct {
	PID           int64  `json:"p_id"`
	Score         int64  `json:"score"`
	Title         string `json:"title"`
	Detail        string `json:"detail"`
	ExampleInput  string `json:"example_input"`
	ExampleOutput string `json:"example_output"`
	Tag           string `json:"tag"`
}

type AddCodeProblemRsp struct {
	BaseRsp
}

type CodeProblem struct {
	ID            int64  `gorm:"column:id" json:"id" form:"id"`
	PID           int64  `gorm:"column:p_id" json:"p_id" form:"p_id"`
	Score         int64  `gorm:"column:score" json:"score" form:"score"`
	Title         string `gorm:"column:title" json:"title" form:"title"`
	Detail        string `gorm:"column:detail" json:"detail" form:"detail"`
	Tag           string `gorm:"column:tag" json:"tag" form:"tag"`
	ExampleInput  string `gorm:"column:example_input" json:"example_input" form:"example_input"`
	ExampleOutput string `gorm:"column:example_output" json:"example_output" form:"example_output"`
}

func (c *CodeProblem) TableName() string {
	return "code_problem"
}

func (c *CodeProblem) GetID() int64 {
	if c == nil {
		return 0
	}
	return c.ID
}

type CodeProblems []*CodeProblem

func (x CodeProblems) Len() int           { return len(x) }
func (x CodeProblems) Less(i, j int) bool { return x[i].ID < x[j].ID }
func (x CodeProblems) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
