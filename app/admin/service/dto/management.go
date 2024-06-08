package dto

type GetCodeProblemReq struct {
	Offset    uint32 `json:"offset" form:"offset"`
	Limit     uint32 `json:"limit" form:"limit"`
	IsReverse bool   `json:"reverse" form:"reverse"`
}

type GetCodeProblemRsp struct {
	BaseRsp
	Problems []*CodeProblem `json:"problems"`
	Total    uint32         `json:"total"`
	HasMore  bool           `json:"has_more"`
}

type UpdateCodeProblemReq struct {
	ID            int64  `json:"id"`
	PID           int64  `json:"p_id"`
	Score         int64  `json:"score"`
	Title         string `json:"title"`
	Detail        string `json:"detail"`
	ExampleInput  string `json:"example_input"`
	ExampleOutput string `json:"example_output"`
	Tag           string `json:"tag"`
}

func (up *UpdateCodeProblemReq) ToCodeProblem() *CodeProblem {
	if up == nil {
		return nil
	}
	return &CodeProblem{
		ID:            up.ID,
		PID:           up.PID,
		Score:         up.Score,
		Title:         up.Title,
		Detail:        up.Detail,
		Tag:           up.Tag,
		ExampleInput:  up.ExampleInput,
		ExampleOutput: up.ExampleOutput,
	}
}

type UpdateCodeProblemRsp struct {
	BaseRsp
}
