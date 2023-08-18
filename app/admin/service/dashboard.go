package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-admin/app/admin/models"
	"io/ioutil"
	"net/http"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/pkg/errors"
	"go-admin/app/admin/service/dto"
)

type Dashboard struct {
	service.Service
}

// All 获取所有看板
func (d *Dashboard) All(ctx context.Context, name string, id int) (*dto.AllDashboardRsp, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid name")
	}
	campProgression, err := d.GetCampProgression(name)
	if err != nil {
		return nil, errors.Wrap(err, "d.GetCampProgressions err")
	}
	// 查评语
	var data models.SysUser
	remark := ""
	if err := d.Orm.First(&data, id).Error; err != nil {
		remark = "N/A"
	} else {
		remark = data.Remark
	}
	if remark == "" {
		remark = "N/A"
	}
	return &dto.AllDashboardRsp{CampProgressions: campProgression, Remark: remark}, nil
}

func (d *Dashboard) GetCampProgression(name string) (*dto.CampaignProgression, error) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:7777/user_progression", nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest err")
	}
	params := req.URL.Query()
	params.Add("name", name)
	req.URL.RawQuery = params.Encode()
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "cli.Do err")
	}
	defer func() {
		if err := rsp.Body.Close(); err != nil {
			_ = rsp.Body.Close()
		}
	}()
	bs, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ReadAll err")
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http err code:%d", rsp.StatusCode)
	}
	data := &model.GetUserProgressionRsp{}
	if err := json.Unmarshal(bs, data); err != nil {
		return nil, errors.Wrap(err, "unmarshal err")
	}
	return dto.NewCampaignProgression(data.CampaignProgression), nil
}
