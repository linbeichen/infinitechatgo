package dto

type TwoTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ModifyFriendApplicationResponse struct {
	UserId      string `json:"userId"`
	SessionId   string `json:"sessionId"`
	SessionType int32  `json:"sessionType"`
	SessionName string `json:"sessionName"`
	Avatar      string `json:"avatar"`
}
