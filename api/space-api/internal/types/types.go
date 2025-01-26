// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AddSpaceMemberReq struct {
	UserId    int64  `json:"userId"`
	SpaceRole string `json:"spaceRole"` // viewer/editor/admin
}

type AddSpaceMemberResp struct {
	Success bool `json:"success"`
}

type CreateSpaceReq struct {
	SpaceName  string `json:"spaceName"`
	SpaceDesc  string `json:"spaceDesc,optional"`
	SpaceType  string `json:"spaceType"`  // private/team
	SpaceLevel string `json:"spaceLevel"` // normal/pro/premium
}

type CreateSpaceResp struct {
	Id int64 `json:"id"`
}

type GetSpaceReq struct {
	Id int64 `path:"id, optional"` // 从 URL 路径中获取空间 ID
}

type GetSpaceResp struct {
	Id         int64  `json:"id"`
	SpaceName  string `json:"spaceName"`
	SpaceType  string `json:"spaceType"`  // private/team
	SpaceLevel string `json:"spaceLevel"` // normal/pro/premium
	MaxSize    int64  `json:"maxSize"`    // 最大总大小
	MaxCount   int64  `json:"maxCount"`   // 最大数量
	TotalSize  int64  `json:"totalSize"`  // 当前总大小
	TotalCount int64  `json:"totalCount"` // 当前数量
	UserId     int64  `json:"userId"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type ListSpaceMembersResp struct {
	Members []SpaceMemberInfo `json:"members"`
}

type SpaceMemberInfo struct {
	Id         int64  `json:"id"`
	SpaceId    int64  `json:"spaceId"`
	UserId     int64  `json:"userId"`
	UserName   string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
	SpaceRole  string `json:"spaceRole"`
	CreateTime string `json:"createTime"`
}

type UpdateSpaceUsageReq struct {
	Size      int64  `json:"size"`
	Operation string `json:"operation"` // add/subtract
}

type UpdateSpaceUsageResp struct {
	Success       bool  `json:"success"`
	RemainingSize int64 `json:"remainingSize"`
}
