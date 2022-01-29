// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package connectv1alpha1

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

// WatchServiceClient is the client API for WatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchServiceClient interface {
	// Watch watches for session proposals for taking player connections.
	Watch(ctx context.Context, opts ...grpc.CallOption) (WatchService_WatchClient, error)
}

type watchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchServiceClient(cc grpc.ClientConnInterface) WatchServiceClient {
	return &watchServiceClient{cc}
}

func (c *watchServiceClient) Watch(ctx context.Context, opts ...grpc.CallOption) (WatchService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &WatchService_ServiceDesc.Streams[0], "/minekube.connect.v1alpha1.WatchService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceWatchClient{stream}
	return x, nil
}

type WatchService_WatchClient interface {
	Send(*WatchRequest) error
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type watchServiceWatchClient struct {
	grpc.ClientStream
}

func (x *watchServiceWatchClient) Send(m *WatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *watchServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchServiceServer is the server API for WatchService service.
// All implementations must embed UnimplementedWatchServiceServer
// for forward compatibility
type WatchServiceServer interface {
	// Watch watches for session proposals for taking player connections.
	Watch(WatchService_WatchServer) error
	mustEmbedUnimplementedWatchServiceServer()
}

// UnimplementedWatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWatchServiceServer struct {
}

func (UnimplementedWatchServiceServer) Watch(WatchService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedWatchServiceServer) mustEmbedUnimplementedWatchServiceServer() {}

// UnsafeWatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchServiceServer will
// result in compilation errors.
type UnsafeWatchServiceServer interface {
	mustEmbedUnimplementedWatchServiceServer()
}

func RegisterWatchServiceServer(s grpc.ServiceRegistrar, srv WatchServiceServer) {
	s.RegisterService(&WatchService_ServiceDesc, srv)
}

func _WatchService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WatchServiceServer).Watch(&watchServiceWatchServer{stream})
}

type WatchService_WatchServer interface {
	Send(*WatchResponse) error
	Recv() (*WatchRequest, error)
	grpc.ServerStream
}

type watchServiceWatchServer struct {
	grpc.ServerStream
}

func (x *watchServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *watchServiceWatchServer) Recv() (*WatchRequest, error) {
	m := new(WatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchService_ServiceDesc is the grpc.ServiceDesc for WatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "minekube.connect.v1alpha1.WatchService",
	HandlerType: (*WatchServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _WatchService_Watch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "minekube/connect/v1alpha1/watch_service.proto",
}
