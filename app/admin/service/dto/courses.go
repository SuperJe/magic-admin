package dto

import (
	"fmt"
	"time"
)

const (
	CourseType = iota + 1
	CourseTypePY
	CourseTypeCPP
)

type LessonRecord struct {
	ID            int64     `json:"id" gorm:"column:id"`
	UserID        int64     `json:"user_id" gorm:"column:user_id"`
	CourseType    int32     `json:"course_type" gorm:"course_type"`
	KnowledgeTags string    `json:"tags" gorm:"column:tags"`
	Teacher       string    `json:"teacher" gorm:"column:teacher"`
	Remark        string    `json:"remark" gorm:"column:remark"`
	Updated       time.Time `json:"updated" gorm:"column:updated"`
	Created       time.Time `json:"created" gorm:"column:created"`
	LearnedTime   string    `json:"date" gorm:"-"`
}

func (lr *LessonRecord) TableName() string {
	return fmt.Sprintf("lesson_record_%03d", lr.UserID%10)
}

type Course struct {
	ID int64
	CourseDetail
	Updated time.Time `gorm:"column:updated"`
	Created time.Time `gorm:"column:created"`
}

func (c *Course) TableName() string {
	return "courses"
}

type GetCoursesDetailReq struct {
	CourseType int32 `form:"course_type" json:"course_type"`
}

type CourseDetail struct {
	Name                 string `json:"name"`
	PreRequired          string `json:"pre_required"`
	Target               string `json:"target"`
	RecommendCompetition string `json:"recommend_competition"`
	RecommendPeriod      string `json:"recommend_period"`
	CourseType           int32  `json:"course_type"`
	TotalLessonHours     int32  `json:"total_lesson_hours"`
}

type GetCoursesDetailRsp struct {
	CourseDetail `json:"detail"`
}

type GetLearnedReq struct {
	CourseType int32 `form:"course_type" json:"course_type"`
	UserID     int64 `json:"user_id"`
}

type GetLearnedRsp struct {
	TotalLessonHours int32           `json:"total_lesson_hours"`
	Records          []*LessonRecord `json:"records"`
}

type SignLessonReq struct {
	CourseType int32 `json:"course_type"`
	UserID     int64 `json:"user_id"`
}

type SignLessonRsp struct {
	LearnedLessons int32 `json:"learned_lessons"`
}

type AddLessonRecordReq struct {
	CourseType    int32     `json:"course_type"`
	UserID        int       `json:"-"`
	Name          string    `json:"name"`
	KnowledgeTags string    `json:"tags"`
	Teacher       string    `json:"teacher"`
	Remark        string    `json:"remark"`
	Created       time.Time `json:"created"`
}

type AddLessonRecordRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	ID   int64  `json:"id"`
}

type GetStudentNameReq struct {
	UserID int
	Value  string
}
