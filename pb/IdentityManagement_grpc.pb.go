// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: proto/IdentityManagement.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IdentityManagementService_Login_FullMethodName            = "/hestia.api.idm.IdentityManagementService/Login"
	IdentityManagementService_Register_FullMethodName         = "/hestia.api.idm.IdentityManagementService/Register"
	IdentityManagementService_Alive_FullMethodName            = "/hestia.api.idm.IdentityManagementService/Alive"
	IdentityManagementService_Logout_FullMethodName           = "/hestia.api.idm.IdentityManagementService/Logout"
	IdentityManagementService_AddUserToCompany_FullMethodName = "/hestia.api.idm.IdentityManagementService/AddUserToCompany"
)

// IdentityManagementServiceClient is the client API for IdentityManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityManagementServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Alive(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Logout(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddUserToCompany(ctx context.Context, in *AddUserToCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type identityManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityManagementServiceClient(cc grpc.ClientConnInterface) IdentityManagementServiceClient {
	return &identityManagementServiceClient{cc}
}

func (c *identityManagementServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, IdentityManagementService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityManagementService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementServiceClient) Alive(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityManagementService_Alive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementServiceClient) Logout(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityManagementService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementServiceClient) AddUserToCompany(ctx context.Context, in *AddUserToCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityManagementService_AddUserToCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityManagementServiceServer is the server API for IdentityManagementService service.
// All implementations must embed UnimplementedIdentityManagementServiceServer
// for forward compatibility
type IdentityManagementServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*emptypb.Empty, error)
	Alive(context.Context, *TokenRequest) (*emptypb.Empty, error)
	Logout(context.Context, *TokenRequest) (*emptypb.Empty, error)
	AddUserToCompany(context.Context, *AddUserToCompanyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedIdentityManagementServiceServer()
}

// UnimplementedIdentityManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityManagementServiceServer struct {
}

func (UnimplementedIdentityManagementServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedIdentityManagementServiceServer) Register(context.Context, *RegisterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedIdentityManagementServiceServer) Alive(context.Context, *TokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}
func (UnimplementedIdentityManagementServiceServer) Logout(context.Context, *TokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedIdentityManagementServiceServer) AddUserToCompany(context.Context, *AddUserToCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToCompany not implemented")
}
func (UnimplementedIdentityManagementServiceServer) mustEmbedUnimplementedIdentityManagementServiceServer() {
}

// UnsafeIdentityManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityManagementServiceServer will
// result in compilation errors.
type UnsafeIdentityManagementServiceServer interface {
	mustEmbedUnimplementedIdentityManagementServiceServer()
}

func RegisterIdentityManagementServiceServer(s grpc.ServiceRegistrar, srv IdentityManagementServiceServer) {
	s.RegisterService(&IdentityManagementService_ServiceDesc, srv)
}

func _IdentityManagementService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagementService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagementService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagementService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagementService_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServiceServer).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagementService_Alive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServiceServer).Alive(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagementService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagementService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServiceServer).Logout(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagementService_AddUserToCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServiceServer).AddUserToCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagementService_AddUserToCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServiceServer).AddUserToCompany(ctx, req.(*AddUserToCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityManagementService_ServiceDesc is the grpc.ServiceDesc for IdentityManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hestia.api.idm.IdentityManagementService",
	HandlerType: (*IdentityManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _IdentityManagementService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _IdentityManagementService_Register_Handler,
		},
		{
			MethodName: "Alive",
			Handler:    _IdentityManagementService_Alive_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _IdentityManagementService_Logout_Handler,
		},
		{
			MethodName: "AddUserToCompany",
			Handler:    _IdentityManagementService_AddUserToCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/IdentityManagement.proto",
}
