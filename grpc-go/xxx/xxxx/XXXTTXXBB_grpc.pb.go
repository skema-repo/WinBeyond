// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.0
// source: XXXTTXXBB.proto

package xxxx

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

// XXXTTXXBBClient is the client API for XXXTTXXBB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XXXTTXXBBClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type xXXTTXXBBClient struct {
	cc grpc.ClientConnInterface
}

func NewXXXTTXXBBClient(cc grpc.ClientConnInterface) XXXTTXXBBClient {
	return &xXXTTXXBBClient{cc}
}

func (c *xXXTTXXBBClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/xxx.xxxx.XXXTTXXBB/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XXXTTXXBBServer is the server API for XXXTTXXBB service.
// All implementations must embed UnimplementedXXXTTXXBBServer
// for forward compatibility
type XXXTTXXBBServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedXXXTTXXBBServer()
}

// UnimplementedXXXTTXXBBServer must be embedded to have forward compatible implementations.
type UnimplementedXXXTTXXBBServer struct {
}

func (UnimplementedXXXTTXXBBServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedXXXTTXXBBServer) mustEmbedUnimplementedXXXTTXXBBServer() {}

// UnsafeXXXTTXXBBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XXXTTXXBBServer will
// result in compilation errors.
type UnsafeXXXTTXXBBServer interface {
	mustEmbedUnimplementedXXXTTXXBBServer()
}

func RegisterXXXTTXXBBServer(s grpc.ServiceRegistrar, srv XXXTTXXBBServer) {
	s.RegisterService(&XXXTTXXBB_ServiceDesc, srv)
}

func _XXXTTXXBB_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XXXTTXXBBServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxx.xxxx.XXXTTXXBB/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XXXTTXXBBServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// XXXTTXXBB_ServiceDesc is the grpc.ServiceDesc for XXXTTXXBB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XXXTTXXBB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xxx.xxxx.XXXTTXXBB",
	HandlerType: (*XXXTTXXBBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _XXXTTXXBB_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "XXXTTXXBB.proto",
}
