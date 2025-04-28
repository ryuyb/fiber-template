package presenter

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	AccessToken string `json:"accessToken"`
}
