package dto

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

type AllDashboardRsp struct {
	CampProgressions *CampaignProgression `json:"camp_progressions"`
}
