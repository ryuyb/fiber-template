package presenter

import "time"

type SaveUserReq struct {
	Id       *uint32 `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

type UserResp struct {
	Id         uint32    `json:"id"`
	Username   string    `json:"username"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
