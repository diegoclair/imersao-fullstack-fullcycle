// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PixServiceClient is the client API for PixService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PixServiceClient interface {
	AddPixKey(ctx context.Context, in *AddPixKeyRequest, opts ...grpc.CallOption) (*AddPixKeyResponse, error)
	FindPixKeyByID(ctx context.Context, in *FindPixKeyByIDRequest, opts ...grpc.CallOption) (*FindPixKeyByIDResponse, error)
}

type pixServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPixServiceClient(cc grpc.ClientConnInterface) PixServiceClient {
	return &pixServiceClient{cc}
}

func (c *pixServiceClient) AddPixKey(ctx context.Context, in *AddPixKeyRequest, opts ...grpc.CallOption) (*AddPixKeyResponse, error) {
	out := new(AddPixKeyResponse)
	err := c.cc.Invoke(ctx, "/github.com.diegoclair.codepix.PixService/AddPixKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pixServiceClient) FindPixKeyByID(ctx context.Context, in *FindPixKeyByIDRequest, opts ...grpc.CallOption) (*FindPixKeyByIDResponse, error) {
	out := new(FindPixKeyByIDResponse)
	err := c.cc.Invoke(ctx, "/github.com.diegoclair.codepix.PixService/FindPixKeyByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PixServiceServer is the server API for PixService service.
// All implementations must embed UnimplementedPixServiceServer
// for forward compatibility
type PixServiceServer interface {
	AddPixKey(context.Context, *AddPixKeyRequest) (*AddPixKeyResponse, error)
	FindPixKeyByID(context.Context, *FindPixKeyByIDRequest) (*FindPixKeyByIDResponse, error)
	mustEmbedUnimplementedPixServiceServer()
}

// UnimplementedPixServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPixServiceServer struct {
}

func (UnimplementedPixServiceServer) AddPixKey(context.Context, *AddPixKeyRequest) (*AddPixKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPixKey not implemented")
}
func (UnimplementedPixServiceServer) FindPixKeyByID(context.Context, *FindPixKeyByIDRequest) (*FindPixKeyByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPixKeyByID not implemented")
}
func (UnimplementedPixServiceServer) mustEmbedUnimplementedPixServiceServer() {}

// UnsafePixServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PixServiceServer will
// result in compilation errors.
type UnsafePixServiceServer interface {
	mustEmbedUnimplementedPixServiceServer()
}

func RegisterPixServiceServer(s grpc.ServiceRegistrar, srv PixServiceServer) {
	s.RegisterService(&PixService_ServiceDesc, srv)
}

func _PixService_AddPixKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPixKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PixServiceServer).AddPixKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.diegoclair.codepix.PixService/AddPixKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PixServiceServer).AddPixKey(ctx, req.(*AddPixKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PixService_FindPixKeyByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPixKeyByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PixServiceServer).FindPixKeyByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.diegoclair.codepix.PixService/FindPixKeyByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PixServiceServer).FindPixKeyByID(ctx, req.(*FindPixKeyByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PixService_ServiceDesc is the grpc.ServiceDesc for PixService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PixService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.diegoclair.codepix.PixService",
	HandlerType: (*PixServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPixKey",
			Handler:    _PixService_AddPixKey_Handler,
		},
		{
			MethodName: "FindPixKeyByID",
			Handler:    _PixService_FindPixKeyByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pix.proto",
}
