// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: users/v1/users.proto

package users_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UsersService_CreateUsers_FullMethodName  = "/users.v1.UsersService/CreateUsers"
	UsersService_GetUsersById_FullMethodName = "/users.v1.UsersService/GetUsersById"
)

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	CreateUsers(ctx context.Context, in *CreateUsersRequest, opts ...grpc.CallOption) (*CreateUsersResponse, error)
	GetUsersById(ctx context.Context, in *GetUsersByIdRequest, opts ...grpc.CallOption) (*GetUsersByIdResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) CreateUsers(ctx context.Context, in *CreateUsersRequest, opts ...grpc.CallOption) (*CreateUsersResponse, error) {
	out := new(CreateUsersResponse)
	err := c.cc.Invoke(ctx, UsersService_CreateUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUsersById(ctx context.Context, in *GetUsersByIdRequest, opts ...grpc.CallOption) (*GetUsersByIdResponse, error) {
	out := new(GetUsersByIdResponse)
	err := c.cc.Invoke(ctx, UsersService_GetUsersById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations should embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	CreateUsers(context.Context, *CreateUsersRequest) (*CreateUsersResponse, error)
	GetUsersById(context.Context, *GetUsersByIdRequest) (*GetUsersByIdResponse, error)
}

// UnimplementedUsersServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) CreateUsers(context.Context, *CreateUsersRequest) (*CreateUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUsers not implemented")
}
func (UnimplementedUsersServiceServer) GetUsersById(context.Context, *GetUsersByIdRequest) (*GetUsersByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersById not implemented")
}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_CreateUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).CreateUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_CreateUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).CreateUsers(ctx, req.(*CreateUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUsersById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUsersById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_GetUsersById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUsersById(ctx, req.(*GetUsersByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.v1.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUsers",
			Handler:    _UsersService_CreateUsers_Handler,
		},
		{
			MethodName: "GetUsersById",
			Handler:    _UsersService_GetUsersById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/v1/users.proto",
}
