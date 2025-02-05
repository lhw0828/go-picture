// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: space.proto

package server

import (
	"context"

	"picture/rpc/space-rpc/internal/logic"
	"picture/rpc/space-rpc/internal/svc"
	"picture/rpc/space-rpc/space"
)

type SpaceServiceServer struct {
	svcCtx *svc.ServiceContext
	space.UnimplementedSpaceServiceServer
}

func NewSpaceServiceServer(svcCtx *svc.ServiceContext) *SpaceServiceServer {
	return &SpaceServiceServer{
		svcCtx: svcCtx,
	}
}

// 创建空间
func (s *SpaceServiceServer) CreateSpace(ctx context.Context, in *space.CreateSpaceRequest) (*space.SpaceInfo, error) {
	l := logic.NewCreateSpaceLogic(ctx, s.svcCtx)
	return l.CreateSpace(in)
}

// 获取空间信息
func (s *SpaceServiceServer) GetSpace(ctx context.Context, in *space.GetSpaceRequest) (*space.SpaceInfo, error) {
	l := logic.NewGetSpaceLogic(ctx, s.svcCtx)
	return l.GetSpace(in)
}

// 更新空间信息
func (s *SpaceServiceServer) UpdateSpace(ctx context.Context, in *space.UpdateSpaceRequest) (*space.BaseResponse, error) {
	l := logic.NewUpdateSpaceLogic(ctx, s.svcCtx)
	return l.UpdateSpace(in)
}

// 删除空间
func (s *SpaceServiceServer) DeleteSpace(ctx context.Context, in *space.DeleteSpaceRequest) (*space.BaseResponse, error) {
	l := logic.NewDeleteSpaceLogic(ctx, s.svcCtx)
	return l.DeleteSpace(in)
}

// 获取空间列表
func (s *SpaceServiceServer) ListSpace(ctx context.Context, in *space.ListSpaceRequest) (*space.ListSpaceResponse, error) {
	l := logic.NewListSpaceLogic(ctx, s.svcCtx)
	return l.ListSpace(in)
}

// 获取空间成员列表
func (s *SpaceServiceServer) ListSpaceMembers(ctx context.Context, in *space.ListSpaceMembersRequest) (*space.ListSpaceMembersResponse, error) {
	l := logic.NewListSpaceMembersLogic(ctx, s.svcCtx)
	return l.ListSpaceMembers(in)
}

// 获取空间分析数据
func (s *SpaceServiceServer) GetSpaceAnalysis(ctx context.Context, in *space.GetSpaceAnalysisRequest) (*space.SpaceAnalysis, error) {
	l := logic.NewGetSpaceAnalysisLogic(ctx, s.svcCtx)
	return l.GetSpaceAnalysis(in)
}

// 获取空间权限
func (s *SpaceServiceServer) GetSpacePermissions(ctx context.Context, in *space.GetSpacePermissionsRequest) (*space.GetSpacePermissionsResponse, error) {
	l := logic.NewGetSpacePermissionsLogic(ctx, s.svcCtx)
	return l.GetSpacePermissions(in)
}

func (s *SpaceServiceServer) GetSpaceVO(ctx context.Context, in *space.GetSpaceVORequest) (*space.SpaceVO, error) {
	l := logic.NewGetSpaceVOLogic(ctx, s.svcCtx)
	return l.GetSpaceVO(in)
}
