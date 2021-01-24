// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package statistico

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StrategyServiceClient is the client API for StrategyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrategyServiceClient interface {
	StrategyTradeSearch(ctx context.Context, in *StrategyTradeSearchRequest, opts ...grpc.CallOption) (StrategyService_StrategyTradeSearchClient, error)
}

type strategyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategyServiceClient(cc grpc.ClientConnInterface) StrategyServiceClient {
	return &strategyServiceClient{cc}
}

func (c *strategyServiceClient) StrategyTradeSearch(ctx context.Context, in *StrategyTradeSearchRequest, opts ...grpc.CallOption) (StrategyService_StrategyTradeSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StrategyService_serviceDesc.Streams[0], "/statistico.StrategyService/StrategyTradeSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &strategyServiceStrategyTradeSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StrategyService_StrategyTradeSearchClient interface {
	Recv() (*StrategyTrade, error)
	grpc.ClientStream
}

type strategyServiceStrategyTradeSearchClient struct {
	grpc.ClientStream
}

func (x *strategyServiceStrategyTradeSearchClient) Recv() (*StrategyTrade, error) {
	m := new(StrategyTrade)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StrategyServiceServer is the server API for StrategyService service.
// All implementations must embed UnimplementedStrategyServiceServer
// for forward compatibility
type StrategyServiceServer interface {
	StrategyTradeSearch(*StrategyTradeSearchRequest, StrategyService_StrategyTradeSearchServer) error
	mustEmbedUnimplementedStrategyServiceServer()
}

// UnimplementedStrategyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStrategyServiceServer struct {
}

func (UnimplementedStrategyServiceServer) StrategyTradeSearch(*StrategyTradeSearchRequest, StrategyService_StrategyTradeSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method StrategyTradeSearch not implemented")
}
func (UnimplementedStrategyServiceServer) mustEmbedUnimplementedStrategyServiceServer() {}

// UnsafeStrategyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrategyServiceServer will
// result in compilation errors.
type UnsafeStrategyServiceServer interface {
	mustEmbedUnimplementedStrategyServiceServer()
}

func RegisterStrategyServiceServer(s grpc.ServiceRegistrar, srv StrategyServiceServer) {
	s.RegisterService(&_StrategyService_serviceDesc, srv)
}

func _StrategyService_StrategyTradeSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StrategyTradeSearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StrategyServiceServer).StrategyTradeSearch(m, &strategyServiceStrategyTradeSearchServer{stream})
}

type StrategyService_StrategyTradeSearchServer interface {
	Send(*StrategyTrade) error
	grpc.ServerStream
}

type strategyServiceStrategyTradeSearchServer struct {
	grpc.ServerStream
}

func (x *strategyServiceStrategyTradeSearchServer) Send(m *StrategyTrade) error {
	return x.ServerStream.SendMsg(m)
}

var _StrategyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "statistico.StrategyService",
	HandlerType: (*StrategyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StrategyTradeSearch",
			Handler:       _StrategyService_StrategyTradeSearch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "strategy.proto",
}
