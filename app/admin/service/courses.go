package service

import (
	"context"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/pkg/errors"
	"go-admin/app/admin/service/dto"
	"time"
)

const createLessonRecordSQL = "CREATE TABLE IF NOT EXISTS `%s` ( " +
	" `id` bigint NOT NULL AUTO_INCREMENT, " +
	"`user_id` int NOT NULL," +
	"`course_type` tinyint(4) NOT NULL," +
	"`teacher` varchar(255) NOT NULL," +
	"`tags` varchar(255) DEFAULT NULL," +
	"`remark` text DEFAULT NULL," +
	"`created` datetime DEFAULT CURRENT_TIMESTAMP," +
	"`updated` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, " +
	"`is_deleted` tinyint(1) NOT NULL DEFAULT 0, " +
	"PRIMARY KEY (`id`)," +
	"KEY `idx_user_course` (`user_id`, `course_type`)" +
	") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci"

type Courses struct {
	service.Service
}

func (c *Courses) GetCourseDetail(ctx context.Context, courseType int32) (*dto.GetCoursesDetailRsp, error) {
	course := &dto.Course{}
	cond := dto.CourseDetail{CourseType: courseType}
	if err := c.Orm.First(course, &dto.Course{CourseDetail: cond}).Error; err != nil {
		return nil, errors.Wrap(err, "c.Orm.First err")
	}
	return &dto.GetCoursesDetailRsp{CourseDetail: course.CourseDetail}, nil
}

func (c *Courses) GetLearnedLessons(ctx context.Context, req *dto.GetLearnedReq) (*dto.GetLearnedRsp, error) {
	// 查总课时
	course := &dto.Course{}
	if req.CourseType != 0 {
		if err := c.Orm.Table(course.TableName()).Select("total_lesson_hours").
			Where("course_type = ?", req.CourseType).First(&course).Error; err != nil {
			return nil, errors.Wrap(err, "select total_lesson_hours err")
		}
	}
	// 查上课记录
	lr := &dto.LessonRecord{UserID: req.UserID}
	tb := lr.TableName()
	sql := fmt.Sprintf(createLessonRecordSQL, tb)
	if err := c.Orm.Exec(sql).Error; err != nil {
		return nil, errors.Wrap(err, "Exec err")
	}
	records := make([]*dto.LessonRecord, 0)
	tx := c.Orm.Table(tb).Find(
		&records, &dto.LessonRecord{CourseType: req.CourseType, UserID: req.UserID})
	if req.CourseType == 0 {
		tx = c.Orm.Table(tb).Where("is_deleted = ?", 0).Find(
			&records, &dto.LessonRecord{UserID: req.UserID})
	}
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "Find err")
	}
	for i := range records {
		records[i].LearnedTime = records[i].Created.Format("2006-01-02 15:04:05")
	}
	return &dto.GetLearnedRsp{
		TotalLessonHours: course.TotalLessonHours,
		Records:          records,
	}, nil
}

func (c *Courses) hasSigned(userID int64, courseType int32) (bool, int64, error) {
	lr := &dto.LessonRecord{UserID: userID}
	tb := lr.TableName()
	count := int64(0)
	err := c.Orm.Table(tb).Where("user_id = ? and course_type = ?", userID, courseType).Count(&count).Error
	if err != nil {
		return false, 0, errors.Wrap(err, "Count err")
	}
	// 没有记录 说明没有签到过
	if count == 0 {
		return false, 0, nil
	}
	// 查最近一条记录
	err = c.Orm.Table(tb).Last(lr, &dto.LessonRecord{UserID: userID, CourseType: courseType}).Error
	if err != nil {
		return false, 0, errors.Wrap(err, "Last err")
	}

	nowStr := time.Now().Format("2006-01-02")
	createStr := lr.Created.Format("2006-01-02")
	return nowStr == createStr, count, nil
}

func (c *Courses) SignLesson(ctx context.Context, req *dto.SignLessonReq) (*dto.SignLessonRsp, error) {
	// 查数据库里有没有在今天上过课的
	signed, count, err := c.hasSigned(req.UserID, req.CourseType)
	if err != nil {
		return nil, errors.Wrap(err, "c.hasSigned err")
	}
	// 已签到，直接返回
	if signed {
		c.Log.Warnf("user %d has signed course %d today", req.UserID, req.CourseType)
		return &dto.SignLessonRsp{LearnedLessons: int32(count)}, nil
	}

	// 未签到，创建一条新记录
	lr := &dto.LessonRecord{
		UserID:        req.UserID,
		CourseType:    req.CourseType,
		KnowledgeTags: "待确认",
		Teacher:       "同步中",
		Remark:        "",
		Created:       time.Now(),
		Updated:       time.Now(),
	}
	err = c.Orm.Table(lr.TableName()).Create(lr).Error
	if err != nil {
		return nil, errors.Wrap(err, "Create err")
	}
	return &dto.SignLessonRsp{LearnedLessons: int32(count) + 1}, nil
}

func (c *Courses) AddLessonRecord(ctx context.Context, req *dto.AddLessonRecordReq) (*dto.AddLessonRecordRsp, error) {
	newRecord := &dto.LessonRecord{
		UserID:        int64(req.UserID),
		CourseType:    req.CourseType,
		KnowledgeTags: req.KnowledgeTags,
		Teacher:       req.Teacher,
		Remark:        req.Remark,
		Created:       req.Created,
		Updated:       time.Now(),
	}
	// TODO: 创建表
	lr := &dto.LessonRecord{UserID: int64(req.UserID)}
	tb := lr.TableName()
	sql := fmt.Sprintf(createLessonRecordSQL, tb)
	if err := c.Orm.Exec(sql).Error; err != nil {
		return nil, errors.Wrap(err, "Exec err")
	}
	err := c.Orm.Table(tb).Create(&newRecord).Error
	if err != nil {
		return nil, errors.Wrap(err, "Create err")
	}
	return &dto.AddLessonRecordRsp{Code: 0, Msg: "ok", ID: newRecord.ID}, nil
}

func (c *Courses) UpdateCourse(ctx context.Context, req *dto.UpdateRecordReq) (*dto.UpdateRecordRsp, error) {
	lr := &dto.LessonRecord{UserID: int64(req.UserID)}
	tb := lr.TableName()
	err := c.Orm.Table(tb).Where("id = ?", req.ID).Updates(dto.LessonRecord{
		CourseType:    req.CourseType,
		KnowledgeTags: req.KnowledgeTags,
		Teacher:       req.Teacher,
		Remark:        req.Remark,
		Created:       req.Created,
		Updated:       time.Now(),
	}).Error
	if err != nil {
		return nil, errors.Wrap(err, "Update err")
	}
	return &dto.UpdateRecordRsp{Code: 0, Msg: "ok", ID: req.ID}, nil
}

func (c *Courses) DeleteCourse(ctx context.Context, req *dto.DeleteRecordReq) (*dto.DeleteRecordRsp, error) {
	lr := &dto.LessonRecord{UserID: int64(req.UserID)}
	tb := lr.TableName()
	err := c.Orm.Table(tb).Where("id = ?", req.ID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, errors.Wrap(err, "Delete err")
	}
	return &dto.DeleteRecordRsp{Code: 0, Msg: "ok", ID: req.ID}, nil
}
