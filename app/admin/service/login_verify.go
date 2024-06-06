package service

import (
	"go-admin/app/admin/service/dto"
	"go-admin/common/service"
)

type LoginVerify struct {
	service.Service
}

func (l *LoginVerify) LoginVerify() (*dto.LoginVerifyReq, error) {

	return nil, nil
}
