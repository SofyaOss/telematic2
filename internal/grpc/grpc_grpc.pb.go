// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: grpc/grpc.proto

package __

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

// GRPCServiceClient is the client API for GRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCServiceClient interface {
	GetCarsByDate(ctx context.Context, in *CarsByDateRequest, opts ...grpc.CallOption) (*CarsByDateResponse, error)
	GetLastCars(ctx context.Context, in *LastCarsRequest, opts ...grpc.CallOption) (*LastCarsResponse, error)
}

type gRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCServiceClient(cc grpc.ClientConnInterface) GRPCServiceClient {
	return &gRPCServiceClient{cc}
}

func (c *gRPCServiceClient) GetCarsByDate(ctx context.Context, in *CarsByDateRequest, opts ...grpc.CallOption) (*CarsByDateResponse, error) {
	out := new(CarsByDateResponse)
	err := c.cc.Invoke(ctx, "/grpc.GRPCService/GetCarsByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCServiceClient) GetLastCars(ctx context.Context, in *LastCarsRequest, opts ...grpc.CallOption) (*LastCarsResponse, error) {
	out := new(LastCarsResponse)
	err := c.cc.Invoke(ctx, "/grpc.GRPCService/GetLastCars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCServiceServer is the server API for GRPCService service.
// All implementations must embed UnimplementedGRPCServiceServer
// for forward compatibility
type GRPCServiceServer interface {
	GetCarsByDate(context.Context, *CarsByDateRequest) (*CarsByDateResponse, error)
	GetLastCars(context.Context, *LastCarsRequest) (*LastCarsResponse, error)
	mustEmbedUnimplementedGRPCServiceServer()
}

// UnimplementedGRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCServiceServer struct {
}

func (UnimplementedGRPCServiceServer) GetCarsByDate(context.Context, *CarsByDateRequest) (*CarsByDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarsByDate not implemented")
}
func (UnimplementedGRPCServiceServer) GetLastCars(context.Context, *LastCarsRequest) (*LastCarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastCars not implemented")
}
func (UnimplementedGRPCServiceServer) mustEmbedUnimplementedGRPCServiceServer() {}

// UnsafeGRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCServiceServer will
// result in compilation errors.
type UnsafeGRPCServiceServer interface {
	mustEmbedUnimplementedGRPCServiceServer()
}

func RegisterGRPCServiceServer(s grpc.ServiceRegistrar, srv GRPCServiceServer) {
	s.RegisterService(&GRPCService_ServiceDesc, srv)
}

func _GRPCService_GetCarsByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarsByDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServiceServer).GetCarsByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GRPCService/GetCarsByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServiceServer).GetCarsByDate(ctx, req.(*CarsByDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCService_GetLastCars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastCarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServiceServer).GetLastCars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GRPCService/GetLastCars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServiceServer).GetLastCars(ctx, req.(*LastCarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCService_ServiceDesc is the grpc.ServiceDesc for GRPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.GRPCService",
	HandlerType: (*GRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCarsByDate",
			Handler:    _GRPCService_GetCarsByDate_Handler,
		},
		{
			MethodName: "GetLastCars",
			Handler:    _GRPCService_GetLastCars_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/grpc.proto",
}
