// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: load.proto

package proto

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

// LoadClient is the client API for Load service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoadClient interface {
	Work(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error)
}

type loadClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadClient(cc grpc.ClientConnInterface) LoadClient {
	return &loadClient{cc}
}

func (c *loadClient) Work(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error) {
	out := new(LoadResponse)
	err := c.cc.Invoke(ctx, "/tempotaletl.Load/Work", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadServer is the server API for Load service.
// All implementations must embed UnimplementedLoadServer
// for forward compatibility
type LoadServer interface {
	Work(context.Context, *LoadRequest) (*LoadResponse, error)
	mustEmbedUnimplementedLoadServer()
}

// UnimplementedLoadServer must be embedded to have forward compatible implementations.
type UnimplementedLoadServer struct {
}

func (UnimplementedLoadServer) Work(context.Context, *LoadRequest) (*LoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Work not implemented")
}
func (UnimplementedLoadServer) mustEmbedUnimplementedLoadServer() {}

// UnsafeLoadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoadServer will
// result in compilation errors.
type UnsafeLoadServer interface {
	mustEmbedUnimplementedLoadServer()
}

func RegisterLoadServer(s grpc.ServiceRegistrar, srv LoadServer) {
	s.RegisterService(&Load_ServiceDesc, srv)
}

func _Load_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tempotaletl.Load/Work",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadServer).Work(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Load_ServiceDesc is the grpc.ServiceDesc for Load service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Load_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tempotaletl.Load",
	HandlerType: (*LoadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Work",
			Handler:    _Load_Work_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "load.proto",
}
