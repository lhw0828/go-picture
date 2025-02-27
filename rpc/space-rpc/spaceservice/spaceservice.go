// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: space.proto

package spaceservice

import (
	"context"

	"picture/rpc/space-rpc/space"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseResponse                = space.BaseResponse
	CategoryCount               = space.CategoryCount
	CreateSpaceRequest          = space.CreateSpaceRequest
	DeleteSpaceRequest          = space.DeleteSpaceRequest
	GetSpaceAnalysisRequest     = space.GetSpaceAnalysisRequest
	GetSpacePermissionsRequest  = space.GetSpacePermissionsRequest
	GetSpacePermissionsResponse = space.GetSpacePermissionsResponse
	GetSpaceRequest             = space.GetSpaceRequest
	GetSpaceVORequest           = space.GetSpaceVORequest
	ListSpaceMembersRequest     = space.ListSpaceMembersRequest
	ListSpaceMembersResponse    = space.ListSpaceMembersResponse
	ListSpaceRequest            = space.ListSpaceRequest
	ListSpaceResponse           = space.ListSpaceResponse
	SizeCount                   = space.SizeCount
	SpaceAnalysis               = space.SpaceAnalysis
	SpaceInfo                   = space.SpaceInfo
	SpaceMember                 = space.SpaceMember
	SpaceUsage                  = space.SpaceUsage
	SpaceUsageTrend             = space.SpaceUsageTrend
	SpaceVO                     = space.SpaceVO
	TagCount                    = space.TagCount
	UpdateSpaceRequest          = space.UpdateSpaceRequest
	UserInfo                    = space.UserInfo

	SpaceService interface {
		// 创建空间
		CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*SpaceInfo, error)
		// 获取空间信息
		GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*SpaceInfo, error)
		// 更新空间信息
		UpdateSpace(ctx context.Context, in *UpdateSpaceRequest, opts ...grpc.CallOption) (*BaseResponse, error)
		// 删除空间
		DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*BaseResponse, error)
		// 获取空间列表
		ListSpace(ctx context.Context, in *ListSpaceRequest, opts ...grpc.CallOption) (*ListSpaceResponse, error)
		// 获取空间成员列表
		ListSpaceMembers(ctx context.Context, in *ListSpaceMembersRequest, opts ...grpc.CallOption) (*ListSpaceMembersResponse, error)
		// 获取空间分析数据
		GetSpaceAnalysis(ctx context.Context, in *GetSpaceAnalysisRequest, opts ...grpc.CallOption) (*SpaceAnalysis, error)
		// 获取空间权限
		GetSpacePermissions(ctx context.Context, in *GetSpacePermissionsRequest, opts ...grpc.CallOption) (*GetSpacePermissionsResponse, error)
		GetSpaceVO(ctx context.Context, in *GetSpaceVORequest, opts ...grpc.CallOption) (*SpaceVO, error)
	}

	defaultSpaceService struct {
		cli zrpc.Client
	}
)

func NewSpaceService(cli zrpc.Client) SpaceService {
	return &defaultSpaceService{
		cli: cli,
	}
}

// 创建空间
func (m *defaultSpaceService) CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*SpaceInfo, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.CreateSpace(ctx, in, opts...)
}

// 获取空间信息
func (m *defaultSpaceService) GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*SpaceInfo, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.GetSpace(ctx, in, opts...)
}

// 更新空间信息
func (m *defaultSpaceService) UpdateSpace(ctx context.Context, in *UpdateSpaceRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.UpdateSpace(ctx, in, opts...)
}

// 删除空间
func (m *defaultSpaceService) DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.DeleteSpace(ctx, in, opts...)
}

// 获取空间列表
func (m *defaultSpaceService) ListSpace(ctx context.Context, in *ListSpaceRequest, opts ...grpc.CallOption) (*ListSpaceResponse, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.ListSpace(ctx, in, opts...)
}

// 获取空间成员列表
func (m *defaultSpaceService) ListSpaceMembers(ctx context.Context, in *ListSpaceMembersRequest, opts ...grpc.CallOption) (*ListSpaceMembersResponse, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.ListSpaceMembers(ctx, in, opts...)
}

// 获取空间分析数据
func (m *defaultSpaceService) GetSpaceAnalysis(ctx context.Context, in *GetSpaceAnalysisRequest, opts ...grpc.CallOption) (*SpaceAnalysis, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.GetSpaceAnalysis(ctx, in, opts...)
}

// 获取空间权限
func (m *defaultSpaceService) GetSpacePermissions(ctx context.Context, in *GetSpacePermissionsRequest, opts ...grpc.CallOption) (*GetSpacePermissionsResponse, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.GetSpacePermissions(ctx, in, opts...)
}

func (m *defaultSpaceService) GetSpaceVO(ctx context.Context, in *GetSpaceVORequest, opts ...grpc.CallOption) (*SpaceVO, error) {
	client := space.NewSpaceServiceClient(m.cli.Conn())
	return client.GetSpaceVO(ctx, in, opts...)
}
