// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: space.proto

package space

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SpaceService_CreateSpace_FullMethodName      = "/space.SpaceService/CreateSpace"
	SpaceService_GetSpace_FullMethodName         = "/space.SpaceService/GetSpace"
	SpaceService_AddSpaceMember_FullMethodName   = "/space.SpaceService/AddSpaceMember"
	SpaceService_ListSpaceMembers_FullMethodName = "/space.SpaceService/ListSpaceMembers"
	SpaceService_UpdateSpaceUsage_FullMethodName = "/space.SpaceService/UpdateSpaceUsage"
	SpaceService_GetSpaceAnalysis_FullMethodName = "/space.SpaceService/GetSpaceAnalysis"
)

// SpaceServiceClient is the client API for SpaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 空间服务
type SpaceServiceClient interface {
	// 创建空间
	CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*CreateSpaceResponse, error)
	// 获取空间信息
	GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error)
	// 添加空间成员
	AddSpaceMember(ctx context.Context, in *AddSpaceMemberRequest, opts ...grpc.CallOption) (*AddSpaceMemberResponse, error)
	// 获取空间成员列表
	ListSpaceMembers(ctx context.Context, in *ListSpaceMembersRequest, opts ...grpc.CallOption) (*ListSpaceMembersResponse, error)
	// 更新空间使用容量
	UpdateSpaceUsage(ctx context.Context, in *UpdateSpaceUsageRequest, opts ...grpc.CallOption) (*UpdateSpaceUsageResponse, error)
	// 获取空间分析数据
	GetSpaceAnalysis(ctx context.Context, in *GetSpaceAnalysisRequest, opts ...grpc.CallOption) (*GetSpaceAnalysisResponse, error)
}

type spaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceServiceClient(cc grpc.ClientConnInterface) SpaceServiceClient {
	return &spaceServiceClient{cc}
}

func (c *spaceServiceClient) CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*CreateSpaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_CreateSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_GetSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) AddSpaceMember(ctx context.Context, in *AddSpaceMemberRequest, opts ...grpc.CallOption) (*AddSpaceMemberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddSpaceMemberResponse)
	err := c.cc.Invoke(ctx, SpaceService_AddSpaceMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) ListSpaceMembers(ctx context.Context, in *ListSpaceMembersRequest, opts ...grpc.CallOption) (*ListSpaceMembersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSpaceMembersResponse)
	err := c.cc.Invoke(ctx, SpaceService_ListSpaceMembers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) UpdateSpaceUsage(ctx context.Context, in *UpdateSpaceUsageRequest, opts ...grpc.CallOption) (*UpdateSpaceUsageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSpaceUsageResponse)
	err := c.cc.Invoke(ctx, SpaceService_UpdateSpaceUsage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetSpaceAnalysis(ctx context.Context, in *GetSpaceAnalysisRequest, opts ...grpc.CallOption) (*GetSpaceAnalysisResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSpaceAnalysisResponse)
	err := c.cc.Invoke(ctx, SpaceService_GetSpaceAnalysis_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServiceServer is the server API for SpaceService service.
// All implementations must embed UnimplementedSpaceServiceServer
// for forward compatibility.
//
// 空间服务
type SpaceServiceServer interface {
	// 创建空间
	CreateSpace(context.Context, *CreateSpaceRequest) (*CreateSpaceResponse, error)
	// 获取空间信息
	GetSpace(context.Context, *GetSpaceRequest) (*GetSpaceResponse, error)
	// 添加空间成员
	AddSpaceMember(context.Context, *AddSpaceMemberRequest) (*AddSpaceMemberResponse, error)
	// 获取空间成员列表
	ListSpaceMembers(context.Context, *ListSpaceMembersRequest) (*ListSpaceMembersResponse, error)
	// 更新空间使用容量
	UpdateSpaceUsage(context.Context, *UpdateSpaceUsageRequest) (*UpdateSpaceUsageResponse, error)
	// 获取空间分析数据
	GetSpaceAnalysis(context.Context, *GetSpaceAnalysisRequest) (*GetSpaceAnalysisResponse, error)
	mustEmbedUnimplementedSpaceServiceServer()
}

// UnimplementedSpaceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSpaceServiceServer struct{}

func (UnimplementedSpaceServiceServer) CreateSpace(context.Context, *CreateSpaceRequest) (*CreateSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (UnimplementedSpaceServiceServer) GetSpace(context.Context, *GetSpaceRequest) (*GetSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpace not implemented")
}
func (UnimplementedSpaceServiceServer) AddSpaceMember(context.Context, *AddSpaceMemberRequest) (*AddSpaceMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSpaceMember not implemented")
}
func (UnimplementedSpaceServiceServer) ListSpaceMembers(context.Context, *ListSpaceMembersRequest) (*ListSpaceMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpaceMembers not implemented")
}
func (UnimplementedSpaceServiceServer) UpdateSpaceUsage(context.Context, *UpdateSpaceUsageRequest) (*UpdateSpaceUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpaceUsage not implemented")
}
func (UnimplementedSpaceServiceServer) GetSpaceAnalysis(context.Context, *GetSpaceAnalysisRequest) (*GetSpaceAnalysisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpaceAnalysis not implemented")
}
func (UnimplementedSpaceServiceServer) mustEmbedUnimplementedSpaceServiceServer() {}
func (UnimplementedSpaceServiceServer) testEmbeddedByValue()                      {}

// UnsafeSpaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceServiceServer will
// result in compilation errors.
type UnsafeSpaceServiceServer interface {
	mustEmbedUnimplementedSpaceServiceServer()
}

func RegisterSpaceServiceServer(s grpc.ServiceRegistrar, srv SpaceServiceServer) {
	// If the following call pancis, it indicates UnimplementedSpaceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SpaceService_ServiceDesc, srv)
}

func _SpaceService_CreateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).CreateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_CreateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).CreateSpace(ctx, req.(*CreateSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_GetSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_GetSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetSpace(ctx, req.(*GetSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_AddSpaceMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSpaceMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).AddSpaceMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_AddSpaceMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).AddSpaceMember(ctx, req.(*AddSpaceMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_ListSpaceMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSpaceMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).ListSpaceMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_ListSpaceMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).ListSpaceMembers(ctx, req.(*ListSpaceMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_UpdateSpaceUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSpaceUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).UpdateSpaceUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_UpdateSpaceUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).UpdateSpaceUsage(ctx, req.(*UpdateSpaceUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_GetSpaceAnalysis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpaceAnalysisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetSpaceAnalysis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_GetSpaceAnalysis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetSpaceAnalysis(ctx, req.(*GetSpaceAnalysisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpaceService_ServiceDesc is the grpc.ServiceDesc for SpaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "space.SpaceService",
	HandlerType: (*SpaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSpace",
			Handler:    _SpaceService_CreateSpace_Handler,
		},
		{
			MethodName: "GetSpace",
			Handler:    _SpaceService_GetSpace_Handler,
		},
		{
			MethodName: "AddSpaceMember",
			Handler:    _SpaceService_AddSpaceMember_Handler,
		},
		{
			MethodName: "ListSpaceMembers",
			Handler:    _SpaceService_ListSpaceMembers_Handler,
		},
		{
			MethodName: "UpdateSpaceUsage",
			Handler:    _SpaceService_UpdateSpaceUsage_Handler,
		},
		{
			MethodName: "GetSpaceAnalysis",
			Handler:    _SpaceService_GetSpaceAnalysis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "space.proto",
}
