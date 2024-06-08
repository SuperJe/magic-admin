package service

import (
	"context"
	"fmt"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/service/dto"
)

type Management struct {
	service.Service
}

func (m *Management) GetCodeProblem(ctx context.Context) ([]*dto.CodeProblem, error) {
	records := make([]*dto.CodeProblem, 0)
	if err := m.Orm.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
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
