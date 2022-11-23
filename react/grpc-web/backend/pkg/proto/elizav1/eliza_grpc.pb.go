// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: eliza.proto

package elizav1

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

// ElizaServiceClient is the client API for ElizaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElizaServiceClient interface {
	// Say is a unary request demo. This method should allow for a one sentence
	// response given a one sentence request.
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
	// Converse is a bi-directional streaming request demo. This method should allow for
	// many requests and many responses.
	Converse(ctx context.Context, opts ...grpc.CallOption) (ElizaService_ConverseClient, error)
	// Introduce is a server-streaming request demo.  This method allows for a single request that will return a series
	// of responses
	Introduce(ctx context.Context, in *IntroduceRequest, opts ...grpc.CallOption) (ElizaService_IntroduceClient, error)
}

type elizaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewElizaServiceClient(cc grpc.ClientConnInterface) ElizaServiceClient {
	return &elizaServiceClient{cc}
}

func (c *elizaServiceClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, "/buf.connect.demo.eliza.v1.ElizaService/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elizaServiceClient) Converse(ctx context.Context, opts ...grpc.CallOption) (ElizaService_ConverseClient, error) {
	stream, err := c.cc.NewStream(ctx, &ElizaService_ServiceDesc.Streams[0], "/buf.connect.demo.eliza.v1.ElizaService/Converse", opts...)
	if err != nil {
		return nil, err
	}
	x := &elizaServiceConverseClient{stream}
	return x, nil
}

type ElizaService_ConverseClient interface {
	Send(*ConverseRequest) error
	Recv() (*ConverseResponse, error)
	grpc.ClientStream
}

type elizaServiceConverseClient struct {
	grpc.ClientStream
}

func (x *elizaServiceConverseClient) Send(m *ConverseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *elizaServiceConverseClient) Recv() (*ConverseResponse, error) {
	m := new(ConverseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *elizaServiceClient) Introduce(ctx context.Context, in *IntroduceRequest, opts ...grpc.CallOption) (ElizaService_IntroduceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ElizaService_ServiceDesc.Streams[1], "/buf.connect.demo.eliza.v1.ElizaService/Introduce", opts...)
	if err != nil {
		return nil, err
	}
	x := &elizaServiceIntroduceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ElizaService_IntroduceClient interface {
	Recv() (*IntroduceResponse, error)
	grpc.ClientStream
}

type elizaServiceIntroduceClient struct {
	grpc.ClientStream
}

func (x *elizaServiceIntroduceClient) Recv() (*IntroduceResponse, error) {
	m := new(IntroduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ElizaServiceServer is the server API for ElizaService service.
// All implementations must embed UnimplementedElizaServiceServer
// for forward compatibility
type ElizaServiceServer interface {
	// Say is a unary request demo. This method should allow for a one sentence
	// response given a one sentence request.
	Say(context.Context, *SayRequest) (*SayResponse, error)
	// Converse is a bi-directional streaming request demo. This method should allow for
	// many requests and many responses.
	Converse(ElizaService_ConverseServer) error
	// Introduce is a server-streaming request demo.  This method allows for a single request that will return a series
	// of responses
	Introduce(*IntroduceRequest, ElizaService_IntroduceServer) error
	mustEmbedUnimplementedElizaServiceServer()
}

// UnimplementedElizaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedElizaServiceServer struct {
}

func (UnimplementedElizaServiceServer) Say(context.Context, *SayRequest) (*SayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}
func (UnimplementedElizaServiceServer) Converse(ElizaService_ConverseServer) error {
	return status.Errorf(codes.Unimplemented, "method Converse not implemented")
}
func (UnimplementedElizaServiceServer) Introduce(*IntroduceRequest, ElizaService_IntroduceServer) error {
	return status.Errorf(codes.Unimplemented, "method Introduce not implemented")
}
func (UnimplementedElizaServiceServer) mustEmbedUnimplementedElizaServiceServer() {}

// UnsafeElizaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElizaServiceServer will
// result in compilation errors.
type UnsafeElizaServiceServer interface {
	mustEmbedUnimplementedElizaServiceServer()
}

func RegisterElizaServiceServer(s grpc.ServiceRegistrar, srv ElizaServiceServer) {
	s.RegisterService(&ElizaService_ServiceDesc, srv)
}

func _ElizaService_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElizaServiceServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.connect.demo.eliza.v1.ElizaService/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElizaServiceServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElizaService_Converse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ElizaServiceServer).Converse(&elizaServiceConverseServer{stream})
}

type ElizaService_ConverseServer interface {
	Send(*ConverseResponse) error
	Recv() (*ConverseRequest, error)
	grpc.ServerStream
}

type elizaServiceConverseServer struct {
	grpc.ServerStream
}

func (x *elizaServiceConverseServer) Send(m *ConverseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *elizaServiceConverseServer) Recv() (*ConverseRequest, error) {
	m := new(ConverseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ElizaService_Introduce_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IntroduceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ElizaServiceServer).Introduce(m, &elizaServiceIntroduceServer{stream})
}

type ElizaService_IntroduceServer interface {
	Send(*IntroduceResponse) error
	grpc.ServerStream
}

type elizaServiceIntroduceServer struct {
	grpc.ServerStream
}

func (x *elizaServiceIntroduceServer) Send(m *IntroduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ElizaService_ServiceDesc is the grpc.ServiceDesc for ElizaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElizaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.connect.demo.eliza.v1.ElizaService",
	HandlerType: (*ElizaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _ElizaService_Say_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Converse",
			Handler:       _ElizaService_Converse_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Introduce",
			Handler:       _ElizaService_Introduce_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eliza.proto",
}