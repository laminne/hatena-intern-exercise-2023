// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: welcome.proto

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
	Welcome_Greet_FullMethodName = "/welcome.Welcome/Greet"
)

// WelcomeClient is the client API for Welcome service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WelcomeClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetReply, error)
}

type welcomeClient struct {
	cc grpc.ClientConnInterface
}

func NewWelcomeClient(cc grpc.ClientConnInterface) WelcomeClient {
	return &welcomeClient{cc}
}

func (c *welcomeClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetReply, error) {
	out := new(GreetReply)
	err := c.cc.Invoke(ctx, Welcome_Greet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WelcomeServer is the server API for Welcome service.
// All implementations must embed UnimplementedWelcomeServer
// for forward compatibility
type WelcomeServer interface {
	Greet(context.Context, *GreetRequest) (*GreetReply, error)
	mustEmbedUnimplementedWelcomeServer()
}

// UnimplementedWelcomeServer must be embedded to have forward compatible implementations.
type UnimplementedWelcomeServer struct {
}

func (UnimplementedWelcomeServer) Greet(context.Context, *GreetRequest) (*GreetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedWelcomeServer) mustEmbedUnimplementedWelcomeServer() {}

// UnsafeWelcomeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WelcomeServer will
// result in compilation errors.
type UnsafeWelcomeServer interface {
	mustEmbedUnimplementedWelcomeServer()
}

func RegisterWelcomeServer(s grpc.ServiceRegistrar, srv WelcomeServer) {
	s.RegisterService(&Welcome_ServiceDesc, srv)
}

func _Welcome_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WelcomeServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Welcome_Greet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WelcomeServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Welcome_ServiceDesc is the grpc.ServiceDesc for Welcome service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Welcome_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "welcome.Welcome",
	HandlerType: (*WelcomeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Welcome_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "welcome.proto",
}