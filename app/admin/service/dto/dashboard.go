package dto

import (
	"encoding/json"
	"fmt"
	"github.com/SuperJe/coco/app/data_proxy/model"
)

type AllDashboardReq struct {
}

type Progression struct {
	Done       int32 `json:"done"`
	Unfinished int32 `json:"unfinished"`
	Total      int32 `json:"total"`
}

type CampaignProgression struct {
	Dungeon  *Progression `json:"dungeon"`
	Forest   *Progression `json:"forrest"`
	Desert   *Progression `json:"desert"`
	Mountain *Progression `json:"mountain"`
	Glacier  *Progression `json:"glacier"`
}

func NewCampaignProgression(cp *model.CampaignProgression) *CampaignProgression {
	bs, _ := json.Marshal(cp)
	data := &CampaignProgression{}
	if err := json.Unmarshal(bs, data); err != nil {
		fmt.Println("unmarshal err:", err.Error())
		return nil
	}
	return data
}

type AllDashboardRsp struct {
	CampProgressions *CampaignProgression `json:"camp_progressions"`
	Remark           string               `json:"remark"`
}
