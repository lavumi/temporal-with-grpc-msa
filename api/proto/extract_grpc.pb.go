// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: extract.proto

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

// ExtractClient is the client API for Extract service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtractClient interface {
	Work(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error)
}

type extractClient struct {
	cc grpc.ClientConnInterface
}

func NewExtractClient(cc grpc.ClientConnInterface) ExtractClient {
	return &extractClient{cc}
}

func (c *extractClient) Work(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error) {
	out := new(ExtractResponse)
	err := c.cc.Invoke(ctx, "/tempotaletl.Extract/Work", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtractServer is the server API for Extract service.
// All implementations must embed UnimplementedExtractServer
// for forward compatibility
type ExtractServer interface {
	Work(context.Context, *ExtractRequest) (*ExtractResponse, error)
	mustEmbedUnimplementedExtractServer()
}

// UnimplementedExtractServer must be embedded to have forward compatible implementations.
type UnimplementedExtractServer struct {
}

func (UnimplementedExtractServer) Work(context.Context, *ExtractRequest) (*ExtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Work not implemented")
}
func (UnimplementedExtractServer) mustEmbedUnimplementedExtractServer() {}

// UnsafeExtractServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtractServer will
// result in compilation errors.
type UnsafeExtractServer interface {
	mustEmbedUnimplementedExtractServer()
}

func RegisterExtractServer(s grpc.ServiceRegistrar, srv ExtractServer) {
	s.RegisterService(&Extract_ServiceDesc, srv)
}

func _Extract_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtractServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tempotaletl.Extract/Work",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtractServer).Work(ctx, req.(*ExtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Extract_ServiceDesc is the grpc.ServiceDesc for Extract service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Extract_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tempotaletl.Extract",
	HandlerType: (*ExtractServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Work",
			Handler:    _Extract_Work_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "extract.proto",
}
