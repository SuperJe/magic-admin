package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/pkg/errors"
	"go-admin/app/admin/service/dto"
)

type Dashboard struct {
	service.Service
}

// All 获取所有看板
func (d *Dashboard) All(ctx context.Context, name string) (*dto.AllDashboardRsp, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid name")
	}
	campProgression, err := d.GetCampProgression(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "d.GetCampProgressions err")
	}
	bs, _ := json.Marshal(campProgression)
	fmt.Println("progressions:", string(bs))
	return &dto.AllDashboardRsp{CampProgressions: campProgression}, nil
	// return &dto.AllDashboardRsp{CampProgressions: &dto.CampaignProgression{
	// 	Dungeon: &dto.Progression{
	// 		Done:       20,
	// 		Unfinished: 80,
	// 		Total:      100,
	// 	},
	// 	Forest: &dto.Progression{
	// 		Done:       40,
	// 		Unfinished: 60,
	// 		Total:      100,
	// 	},
	// 	Desert: &dto.Progression{
	// 		Done:       50,
	// 		Unfinished: 50,
	// 		Total:      100,
	// 	},
	// 	Mountain: &dto.Progression{
	// 		Done:       80,
	// 		Unfinished: 20,
	// 		Total:      100,
	// 	},
	// 	Glacier: &dto.Progression{
	// 		Done:       100,
	// 		Unfinished: 0,
	// 		Total:      100,
	// 	},
	// }}, nil
}

func (d *Dashboard) GetCampProgression(ctx context.Context, name string) (*dto.CampaignProgression, error) {
	cli, err := mongo.NewCocoClient2()
	if err != nil {
		return nil, errors.Wrap(err, "NewCocoClient2 err")
	}
	counts, err := cli.CountLevels(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cli.CountLevels err")
	}
	levels, err := cli.GetCompletedLevels(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "cli.GetCompletedLevels err")
	}
	completed, err := cli.GroupLevelByCampaign(ctx, levels)
	if err != nil {
		return nil, errors.Wrap(err, "cli.GroupLevelByCampaign err")
	}

	progressions := &dto.CampaignProgression{}
	progressions.Dungeon = buildProgression("Dungeon", completed, counts)
	progressions.Forest = buildProgression("Forest", completed, counts)
	progressions.Desert = buildProgression("Desert", completed, counts)
	progressions.Mountain = buildProgression("Mountain", completed, counts)
	progressions.Glacier = buildProgression("Glacier", completed, counts)
	return progressions, nil
}

func buildProgression(campaign string, completed map[string][]string, counts map[string]int32) *dto.Progression {
	return &dto.Progression{
		Done:       int32(len(completed[campaign])),
		Unfinished: counts[campaign] - int32(len(completed[campaign])),
		Total:      counts[campaign],
	}
}
