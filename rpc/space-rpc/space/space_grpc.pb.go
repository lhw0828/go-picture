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
	SpaceService_CreateSpace_FullMethodName         = "/space.SpaceService/CreateSpace"
	SpaceService_GetSpace_FullMethodName            = "/space.SpaceService/GetSpace"
	SpaceService_UpdateSpace_FullMethodName         = "/space.SpaceService/UpdateSpace"
	SpaceService_DeleteSpace_FullMethodName         = "/space.SpaceService/DeleteSpace"
	SpaceService_ListSpace_FullMethodName           = "/space.SpaceService/ListSpace"
	SpaceService_ListSpaceMembers_FullMethodName    = "/space.SpaceService/ListSpaceMembers"
	SpaceService_GetSpaceAnalysis_FullMethodName    = "/space.SpaceService/GetSpaceAnalysis"
	SpaceService_GetSpacePermissions_FullMethodName = "/space.SpaceService/GetSpacePermissions"
	SpaceService_GetSpaceVO_FullMethodName          = "/space.SpaceService/GetSpaceVO"
)

// SpaceServiceClient is the client API for SpaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceServiceClient interface {
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

type spaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceServiceClient(cc grpc.ClientConnInterface) SpaceServiceClient {
	return &spaceServiceClient{cc}
}

func (c *spaceServiceClient) CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*SpaceInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SpaceInfo)
	err := c.cc.Invoke(ctx, SpaceService_CreateSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*SpaceInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SpaceInfo)
	err := c.cc.Invoke(ctx, SpaceService_GetSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) UpdateSpace(ctx context.Context, in *UpdateSpaceRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SpaceService_UpdateSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SpaceService_DeleteSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) ListSpace(ctx context.Context, in *ListSpaceRequest, opts ...grpc.CallOption) (*ListSpaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_ListSpace_FullMethodName, in, out, cOpts...)
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

func (c *spaceServiceClient) GetSpaceAnalysis(ctx context.Context, in *GetSpaceAnalysisRequest, opts ...grpc.CallOption) (*SpaceAnalysis, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SpaceAnalysis)
	err := c.cc.Invoke(ctx, SpaceService_GetSpaceAnalysis_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetSpacePermissions(ctx context.Context, in *GetSpacePermissionsRequest, opts ...grpc.CallOption) (*GetSpacePermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSpacePermissionsResponse)
	err := c.cc.Invoke(ctx, SpaceService_GetSpacePermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetSpaceVO(ctx context.Context, in *GetSpaceVORequest, opts ...grpc.CallOption) (*SpaceVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SpaceVO)
	err := c.cc.Invoke(ctx, SpaceService_GetSpaceVO_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServiceServer is the server API for SpaceService service.
// All implementations must embed UnimplementedSpaceServiceServer
// for forward compatibility.
type SpaceServiceServer interface {
	// 创建空间
	CreateSpace(context.Context, *CreateSpaceRequest) (*SpaceInfo, error)
	// 获取空间信息
	GetSpace(context.Context, *GetSpaceRequest) (*SpaceInfo, error)
	// 更新空间信息
	UpdateSpace(context.Context, *UpdateSpaceRequest) (*BaseResponse, error)
	// 删除空间
	DeleteSpace(context.Context, *DeleteSpaceRequest) (*BaseResponse, error)
	// 获取空间列表
	ListSpace(context.Context, *ListSpaceRequest) (*ListSpaceResponse, error)
	// 获取空间成员列表
	ListSpaceMembers(context.Context, *ListSpaceMembersRequest) (*ListSpaceMembersResponse, error)
	// 获取空间分析数据
	GetSpaceAnalysis(context.Context, *GetSpaceAnalysisRequest) (*SpaceAnalysis, error)
	// 获取空间权限
	GetSpacePermissions(context.Context, *GetSpacePermissionsRequest) (*GetSpacePermissionsResponse, error)
	GetSpaceVO(context.Context, *GetSpaceVORequest) (*SpaceVO, error)
	mustEmbedUnimplementedSpaceServiceServer()
}

// UnimplementedSpaceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSpaceServiceServer struct{}

func (UnimplementedSpaceServiceServer) CreateSpace(context.Context, *CreateSpaceRequest) (*SpaceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (UnimplementedSpaceServiceServer) GetSpace(context.Context, *GetSpaceRequest) (*SpaceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpace not implemented")
}
func (UnimplementedSpaceServiceServer) UpdateSpace(context.Context, *UpdateSpaceRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpace not implemented")
}
func (UnimplementedSpaceServiceServer) DeleteSpace(context.Context, *DeleteSpaceRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpace not implemented")
}
func (UnimplementedSpaceServiceServer) ListSpace(context.Context, *ListSpaceRequest) (*ListSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpace not implemented")
}
func (UnimplementedSpaceServiceServer) ListSpaceMembers(context.Context, *ListSpaceMembersRequest) (*ListSpaceMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpaceMembers not implemented")
}
func (UnimplementedSpaceServiceServer) GetSpaceAnalysis(context.Context, *GetSpaceAnalysisRequest) (*SpaceAnalysis, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpaceAnalysis not implemented")
}
func (UnimplementedSpaceServiceServer) GetSpacePermissions(context.Context, *GetSpacePermissionsRequest) (*GetSpacePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpacePermissions not implemented")
}
func (UnimplementedSpaceServiceServer) GetSpaceVO(context.Context, *GetSpaceVORequest) (*SpaceVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpaceVO not implemented")
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

func _SpaceService_UpdateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).UpdateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_UpdateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).UpdateSpace(ctx, req.(*UpdateSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_DeleteSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).DeleteSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_DeleteSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).DeleteSpace(ctx, req.(*DeleteSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_ListSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).ListSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_ListSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).ListSpace(ctx, req.(*ListSpaceRequest))
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

func _SpaceService_GetSpacePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpacePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetSpacePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_GetSpacePermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetSpacePermissions(ctx, req.(*GetSpacePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_GetSpaceVO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpaceVORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetSpaceVO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_GetSpaceVO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetSpaceVO(ctx, req.(*GetSpaceVORequest))
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
			MethodName: "UpdateSpace",
			Handler:    _SpaceService_UpdateSpace_Handler,
		},
		{
			MethodName: "DeleteSpace",
			Handler:    _SpaceService_DeleteSpace_Handler,
		},
		{
			MethodName: "ListSpace",
			Handler:    _SpaceService_ListSpace_Handler,
		},
		{
			MethodName: "ListSpaceMembers",
			Handler:    _SpaceService_ListSpaceMembers_Handler,
		},
		{
			MethodName: "GetSpaceAnalysis",
			Handler:    _SpaceService_GetSpaceAnalysis_Handler,
		},
		{
			MethodName: "GetSpacePermissions",
			Handler:    _SpaceService_GetSpacePermissions_Handler,
		},
		{
			MethodName: "GetSpaceVO",
			Handler:    _SpaceService_GetSpaceVO_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "space.proto",
}
