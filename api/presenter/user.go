package presenter

type SaveUserReq struct {
	Id       *uint32 `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}
