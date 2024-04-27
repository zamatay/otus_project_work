// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package monitor_v1

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

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitorClient interface {
	Connect(ctx context.Context, in *RequestConnect, opts ...grpc.CallOption) (Monitor_ConnectClient, error)
}

type monitorClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorClient(cc grpc.ClientConnInterface) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) Connect(ctx context.Context, in *RequestConnect, opts ...grpc.CallOption) (Monitor_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Monitor_ServiceDesc.Streams[0], "/monitor.Monitor/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Monitor_ConnectClient interface {
	Recv() (*AllResponse, error)
	grpc.ClientStream
}

type monitorConnectClient struct {
	grpc.ClientStream
}

func (x *monitorConnectClient) Recv() (*AllResponse, error) {
	m := new(AllResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorServer is the server API for Monitor service.
// All implementations must embed UnimplementedMonitorServer
// for forward compatibility
type MonitorServer interface {
	Connect(*RequestConnect, Monitor_ConnectServer) error
	mustEmbedUnimplementedMonitorServer()
}

// UnimplementedMonitorServer must be embedded to have forward compatible implementations.
type UnimplementedMonitorServer struct {
}

func (UnimplementedMonitorServer) Connect(*RequestConnect, Monitor_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedMonitorServer) mustEmbedUnimplementedMonitorServer() {}

// UnsafeMonitorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitorServer will
// result in compilation errors.
type UnsafeMonitorServer interface {
	mustEmbedUnimplementedMonitorServer()
}

func RegisterMonitorServer(s grpc.ServiceRegistrar, srv MonitorServer) {
	s.RegisterService(&Monitor_ServiceDesc, srv)
}

func _Monitor_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestConnect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorServer).Connect(m, &monitorConnectServer{stream})
}

type Monitor_ConnectServer interface {
	Send(*AllResponse) error
	grpc.ServerStream
}

type monitorConnectServer struct {
	grpc.ServerStream
}

func (x *monitorConnectServer) Send(m *AllResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Monitor_ServiceDesc is the grpc.ServiceDesc for Monitor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monitor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Monitor_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "monitor.proto",
}
