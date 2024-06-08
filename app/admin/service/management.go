package service

import (
	"context"
	"fmt"
	"sort"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/service/dto"
)

type Management struct {
	service.Service
}

func (m *Management) GetCodeProblem(ctx context.Context, offset, limit uint32, isReverse bool) (*dto.GetCodeProblemRsp, error) {
	records := make(dto.CodeProblems, 0)
	condID := "id > ?"
	condSelect := "MAX(id)"
	condOrder := "id ASC"
	if isReverse {
		condID = "id < ?"
		condSelect = "MIN(id)"
		condOrder = "id DESC"
	}
	if err := m.Orm.WithContext(ctx).Where(condID, offset).Order(condOrder).Limit(int(limit)).Find(&records).Error; err != nil {
		return nil, err
	}
	sort.Sort(records)

	count := int64(0)
	if err := m.Orm.Model(&dto.CodeProblem{}).WithContext(ctx).Count(&count).Error; err != nil {
		return nil, err
	}
	ID := 0
	if err := m.Orm.Model(&dto.CodeProblem{}).WithContext(ctx).Select(condSelect).Scan(&ID).Error; err != nil {
		return nil, err
	}
	hasMore := true
	if len(records) > 0 {
		if isReverse {
			hasMore = true
		} else {
			hasMore = records[len(records)-1].ID != int64(ID)
		}
	}
	return &dto.GetCodeProblemRsp{
		BaseRsp:  dto.BaseRsp{},
		Problems: records,
		Total:    uint32(count),
		HasMore:  hasMore,
	}, nil
}

func (m *Management) UpdateCodeProblem(ctx context.Context, record *dto.CodeProblem) error {
	if record.GetID() == 0 {
		return fmt.Errorf("id invalid")
	}
	if err := m.Orm.WithContext(ctx).Where("id = ?", record.GetID()).Updates(record).Error; err != nil {
		return err
	}
	return nil
}
