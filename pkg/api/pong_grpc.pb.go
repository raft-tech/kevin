// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/pong.proto

package api

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

// PongServiceClient is the client API for PongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PongServiceClient interface {
	SayPong(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pong, error)
	StreamPong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (PongService_StreamPongClient, error)
	WritePong(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pong, error)
}

type pongServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPongServiceClient(cc grpc.ClientConnInterface) PongServiceClient {
	return &pongServiceClient{cc}
}

func (c *pongServiceClient) SayPong(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/pingpong.PongService/SayPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pongServiceClient) StreamPong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (PongService_StreamPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &PongService_ServiceDesc.Streams[0], "/pingpong.PongService/StreamPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &pongServiceStreamPongClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PongService_StreamPongClient interface {
	Recv() (*Pong, error)
	grpc.ClientStream
}

type pongServiceStreamPongClient struct {
	grpc.ClientStream
}

func (x *pongServiceStreamPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pongServiceClient) WritePong(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/pingpong.PongService/WritePong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PongServiceServer is the server API for PongService service.
// All implementations must embed UnimplementedPongServiceServer
// for forward compatibility
type PongServiceServer interface {
	SayPong(context.Context, *emptypb.Empty) (*Pong, error)
	StreamPong(*Ping, PongService_StreamPongServer) error
	WritePong(context.Context, *emptypb.Empty) (*Pong, error)
	mustEmbedUnimplementedPongServiceServer()
}

// UnimplementedPongServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPongServiceServer struct {
}

func (UnimplementedPongServiceServer) SayPong(context.Context, *emptypb.Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayPong not implemented")
}
func (UnimplementedPongServiceServer) StreamPong(*Ping, PongService_StreamPongServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPong not implemented")
}
func (UnimplementedPongServiceServer) WritePong(context.Context, *emptypb.Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WritePong not implemented")
}
func (UnimplementedPongServiceServer) mustEmbedUnimplementedPongServiceServer() {}

// UnsafePongServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PongServiceServer will
// result in compilation errors.
type UnsafePongServiceServer interface {
	mustEmbedUnimplementedPongServiceServer()
}

func RegisterPongServiceServer(s grpc.ServiceRegistrar, srv PongServiceServer) {
	s.RegisterService(&PongService_ServiceDesc, srv)
}

func _PongService_SayPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PongServiceServer).SayPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.PongService/SayPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PongServiceServer).SayPong(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PongService_StreamPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Ping)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PongServiceServer).StreamPong(m, &pongServiceStreamPongServer{stream})
}

type PongService_StreamPongServer interface {
	Send(*Pong) error
	grpc.ServerStream
}

type pongServiceStreamPongServer struct {
	grpc.ServerStream
}

func (x *pongServiceStreamPongServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

func _PongService_WritePong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PongServiceServer).WritePong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.PongService/WritePong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PongServiceServer).WritePong(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PongService_ServiceDesc is the grpc.ServiceDesc for PongService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PongService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PongService",
	HandlerType: (*PongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayPong",
			Handler:    _PongService_SayPong_Handler,
		},
		{
			MethodName: "WritePong",
			Handler:    _PongService_WritePong_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPong",
			Handler:       _PongService_StreamPong_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/pong.proto",
}
