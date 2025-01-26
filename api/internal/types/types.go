// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type LoginReq struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

type LoginResp struct {
	Id          int64  `json:"id"`
	UserAccount string `json:"userAccount"`
	UserName    string `json:"userName"`
	UserAvatar  string `json:"userAvatar,optional"`
	UserProfile string `json:"userProfile,optional"`
	UserRole    string `json:"userRole"`
	AccessToken string `json:"accessToken"`
}

type RegisterReq struct {
	UserAccount   string `json:"userAccount"`
	UserPassword  string `json:"userPassword"`
	CheckPassword string `json:"checkPassword"`
}

type RegisterResp struct {
	Id int64 `json:"id"`
}
