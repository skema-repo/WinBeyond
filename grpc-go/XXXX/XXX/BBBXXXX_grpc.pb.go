// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.0
// source: BBBXXXX.proto

package XXX

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

// BBBXXXXClient is the client API for BBBXXXX service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BBBXXXXClient interface {
	Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
	Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type bBBXXXXClient struct {
	cc grpc.ClientConnInterface
}

func NewBBBXXXXClient(cc grpc.ClientConnInterface) BBBXXXXClient {
	return &bBBXXXXClient{cc}
}

func (c *bBBXXXXClient) Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "/XXXX.XXX.BBBXXXX/Heathcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bBBXXXXClient) Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/XXXX.XXX.BBBXXXX/Helloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BBBXXXXServer is the server API for BBBXXXX service.
// All implementations must embed UnimplementedBBBXXXXServer
// for forward compatibility
type BBBXXXXServer interface {
	Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	Helloworld(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedBBBXXXXServer()
}

// UnimplementedBBBXXXXServer must be embedded to have forward compatible implementations.
type UnimplementedBBBXXXXServer struct {
}

func (UnimplementedBBBXXXXServer) Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedBBBXXXXServer) Helloworld(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloworld not implemented")
}
func (UnimplementedBBBXXXXServer) mustEmbedUnimplementedBBBXXXXServer() {}

// UnsafeBBBXXXXServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BBBXXXXServer will
// result in compilation errors.
type UnsafeBBBXXXXServer interface {
	mustEmbedUnimplementedBBBXXXXServer()
}

func RegisterBBBXXXXServer(s grpc.ServiceRegistrar, srv BBBXXXXServer) {
	s.RegisterService(&BBBXXXX_ServiceDesc, srv)
}

func _BBBXXXX_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BBBXXXXServer).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XXXX.XXX.BBBXXXX/Heathcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BBBXXXXServer).Heathcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BBBXXXX_Helloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BBBXXXXServer).Helloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XXXX.XXX.BBBXXXX/Helloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BBBXXXXServer).Helloworld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BBBXXXX_ServiceDesc is the grpc.ServiceDesc for BBBXXXX service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BBBXXXX_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "XXXX.XXX.BBBXXXX",
	HandlerType: (*BBBXXXXServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _BBBXXXX_Heathcheck_Handler,
		},
		{
			MethodName: "Helloworld",
			Handler:    _BBBXXXX_Helloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BBBXXXX.proto",
}
