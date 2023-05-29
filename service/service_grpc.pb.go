//
//service.proto 수정시 아래 명령어 실행
//
//protoc -I=. \
//--go_out . --go_opt paths=source_relative \
//--go-grpc_out . --go-grpc_opt paths=source_relative \
//service/service.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: service/service.proto

package service

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
	ServiceInterface_InsertData_FullMethodName = "/service.ServiceInterface/InsertData"
	ServiceInterface_GetData_FullMethodName    = "/service.ServiceInterface/GetData"
)

// ServiceInterfaceClient is the client API for ServiceInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceInterfaceClient interface {
	// Server insert message into DB
	InsertData(ctx context.Context, in *InsertMsg, opts ...grpc.CallOption) (*InsertReponse, error)
	// Server get message from DB
	GetData(ctx context.Context, in *GetMsg, opts ...grpc.CallOption) (*GetResponse, error)
}

type serviceInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceInterfaceClient(cc grpc.ClientConnInterface) ServiceInterfaceClient {
	return &serviceInterfaceClient{cc}
}

func (c *serviceInterfaceClient) InsertData(ctx context.Context, in *InsertMsg, opts ...grpc.CallOption) (*InsertReponse, error) {
	out := new(InsertReponse)
	err := c.cc.Invoke(ctx, ServiceInterface_InsertData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) GetData(ctx context.Context, in *GetMsg, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, ServiceInterface_GetData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceInterfaceServer is the server API for ServiceInterface service.
// All implementations must embed UnimplementedServiceInterfaceServer
// for forward compatibility
type ServiceInterfaceServer interface {
	// Server insert message into DB
	InsertData(context.Context, *InsertMsg) (*InsertReponse, error)
	// Server get message from DB
	GetData(context.Context, *GetMsg) (*GetResponse, error)
	mustEmbedUnimplementedServiceInterfaceServer()
}

// UnimplementedServiceInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceInterfaceServer struct {
}

func (UnimplementedServiceInterfaceServer) InsertData(context.Context, *InsertMsg) (*InsertReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertData not implemented")
}
func (UnimplementedServiceInterfaceServer) GetData(context.Context, *GetMsg) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedServiceInterfaceServer) mustEmbedUnimplementedServiceInterfaceServer() {}

// UnsafeServiceInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceInterfaceServer will
// result in compilation errors.
type UnsafeServiceInterfaceServer interface {
	mustEmbedUnimplementedServiceInterfaceServer()
}

func RegisterServiceInterfaceServer(s grpc.ServiceRegistrar, srv ServiceInterfaceServer) {
	s.RegisterService(&ServiceInterface_ServiceDesc, srv)
}

func _ServiceInterface_InsertData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).InsertData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_InsertData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).InsertData(ctx, req.(*InsertMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).GetData(ctx, req.(*GetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceInterface_ServiceDesc is the grpc.ServiceDesc for ServiceInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ServiceInterface",
	HandlerType: (*ServiceInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertData",
			Handler:    _ServiceInterface_InsertData_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _ServiceInterface_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}
