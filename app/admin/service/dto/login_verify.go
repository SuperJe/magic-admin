package dto

type LoginVerifyReq struct {
	ID     int64  `json:"id" gorm:"column:id"`
	UserID int64  `json:"user_id" gorm:"column:user_id"`
	Psw    string `json:"psw" gorm:"column:psw"`
}

type LoginVerifyRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
