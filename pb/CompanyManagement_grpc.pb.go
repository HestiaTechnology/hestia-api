// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: proto/CompanyManagement.proto

package pb

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
	CompanyManagementService_CreateCompany_FullMethodName = "/hestia.api.com.CompanyManagementService/CreateCompany"
	CompanyManagementService_GetCompany_FullMethodName    = "/hestia.api.com.CompanyManagementService/GetCompany"
	CompanyManagementService_UpdateCompany_FullMethodName = "/hestia.api.com.CompanyManagementService/UpdateCompany"
)

// CompanyManagementServiceClient is the client API for CompanyManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyManagementServiceClient interface {
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error)
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error)
}

type companyManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyManagementServiceClient(cc grpc.ClientConnInterface) CompanyManagementServiceClient {
	return &companyManagementServiceClient{cc}
}

func (c *companyManagementServiceClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error) {
	out := new(CreateCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyManagementService_CreateCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyManagementServiceClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyManagementService_GetCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyManagementServiceClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error) {
	out := new(UpdateCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyManagementService_UpdateCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyManagementServiceServer is the server API for CompanyManagementService service.
// All implementations must embed UnimplementedCompanyManagementServiceServer
// for forward compatibility
type CompanyManagementServiceServer interface {
	CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error)
	GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error)
	mustEmbedUnimplementedCompanyManagementServiceServer()
}

// UnimplementedCompanyManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyManagementServiceServer struct {
}

func (UnimplementedCompanyManagementServiceServer) CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedCompanyManagementServiceServer) GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedCompanyManagementServiceServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCompanyManagementServiceServer) mustEmbedUnimplementedCompanyManagementServiceServer() {
}

// UnsafeCompanyManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyManagementServiceServer will
// result in compilation errors.
type UnsafeCompanyManagementServiceServer interface {
	mustEmbedUnimplementedCompanyManagementServiceServer()
}

func RegisterCompanyManagementServiceServer(s grpc.ServiceRegistrar, srv CompanyManagementServiceServer) {
	s.RegisterService(&CompanyManagementService_ServiceDesc, srv)
}

func _CompanyManagementService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyManagementServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyManagementService_CreateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyManagementServiceServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyManagementService_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyManagementServiceServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyManagementService_GetCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyManagementServiceServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyManagementService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyManagementServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyManagementService_UpdateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyManagementServiceServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyManagementService_ServiceDesc is the grpc.ServiceDesc for CompanyManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hestia.api.com.CompanyManagementService",
	HandlerType: (*CompanyManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyManagementService_CreateCompany_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _CompanyManagementService_GetCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyManagementService_UpdateCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/CompanyManagement.proto",
}
