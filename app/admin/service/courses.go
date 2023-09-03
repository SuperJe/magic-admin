package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/pkg/errors"
	"go-admin/app/admin/service/dto"
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
	if err := c.Orm.Table(course.TableName()).Select("total_lesson_hours").
		Where("course_type = ?", req.CourseType).First(&course).Error; err != nil {
		return nil, errors.Wrap(err, "select total_lesson_hours err")
	}
	// 查上课记录
	lr := &dto.LessonRecord{UserID: req.UserID}
	tb := lr.TableName()
	sql := fmt.Sprintf(createLessonRecordSQL, tb)
	if err := c.Orm.Exec(sql).Error; err != nil {
		return nil, errors.Wrap(err, "Exec err")
	}
	records := make([]*dto.LessonRecord, 0)
	tx := c.Orm.Table(tb).Find(&records, &dto.LessonRecord{CourseType: req.CourseType, UserID: req.UserID})
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "Find err")
	}
	fmt.Printf("\n\ntotal hourse:%d, leanred hours:%d\n", course.TotalLessonHours, tx.RowsAffected)
	bs, _ := json.Marshal(records)
	fmt.Println("records:", string(bs))
	return &dto.GetLearnedRsp{
		TotalLessonHours: course.TotalLessonHours,
		Records:          records,
	}, nil
}
