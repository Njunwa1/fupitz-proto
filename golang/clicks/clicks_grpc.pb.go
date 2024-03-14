// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: clicks/clicks.proto

package clicks

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
	UrlClicks_Redirect_FullMethodName = "/UrlClicks/Redirect"
)

// UrlClicksClient is the client API for UrlClicks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlClicksClient interface {
	Redirect(ctx context.Context, in *UrlClickRequest, opts ...grpc.CallOption) (*UrlClickResponse, error)
}

type urlClicksClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlClicksClient(cc grpc.ClientConnInterface) UrlClicksClient {
	return &urlClicksClient{cc}
}

func (c *urlClicksClient) Redirect(ctx context.Context, in *UrlClickRequest, opts ...grpc.CallOption) (*UrlClickResponse, error) {
	out := new(UrlClickResponse)
	err := c.cc.Invoke(ctx, UrlClicks_Redirect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlClicksServer is the server API for UrlClicks service.
// All implementations must embed UnimplementedUrlClicksServer
// for forward compatibility
type UrlClicksServer interface {
	Redirect(context.Context, *UrlClickRequest) (*UrlClickResponse, error)
	mustEmbedUnimplementedUrlClicksServer()
}

// UnimplementedUrlClicksServer must be embedded to have forward compatible implementations.
type UnimplementedUrlClicksServer struct {
}

func (UnimplementedUrlClicksServer) Redirect(context.Context, *UrlClickRequest) (*UrlClickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redirect not implemented")
}
func (UnimplementedUrlClicksServer) mustEmbedUnimplementedUrlClicksServer() {}

// UnsafeUrlClicksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlClicksServer will
// result in compilation errors.
type UnsafeUrlClicksServer interface {
	mustEmbedUnimplementedUrlClicksServer()
}

func RegisterUrlClicksServer(s grpc.ServiceRegistrar, srv UrlClicksServer) {
	s.RegisterService(&UrlClicks_ServiceDesc, srv)
}

func _UrlClicks_Redirect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlClickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlClicksServer).Redirect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlClicks_Redirect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlClicksServer).Redirect(ctx, req.(*UrlClickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlClicks_ServiceDesc is the grpc.ServiceDesc for UrlClicks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlClicks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UrlClicks",
	HandlerType: (*UrlClicksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Redirect",
			Handler:    _UrlClicks_Redirect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clicks/clicks.proto",
}
