// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: url/url.proto

package url

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
	Url_CreateShortUrl_FullMethodName = "/Url/CreateShortUrl"
	Url_GetUrlByKey_FullMethodName    = "/Url/GetUrlByKey"
	Url_GetAllUserUrls_FullMethodName = "/Url/GetAllUserUrls"
)

// UrlClient is the client API for Url service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlClient interface {
	CreateShortUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error)
	GetUrlByKey(ctx context.Context, in *UrlByKeyRequest, opts ...grpc.CallOption) (*UrlResponse, error)
	GetAllUserUrls(ctx context.Context, in *UserUrlsRequest, opts ...grpc.CallOption) (*UserUrlsResponse, error)
}

type urlClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlClient(cc grpc.ClientConnInterface) UrlClient {
	return &urlClient{cc}
}

func (c *urlClient) CreateShortUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, Url_CreateShortUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlClient) GetUrlByKey(ctx context.Context, in *UrlByKeyRequest, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, Url_GetUrlByKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlClient) GetAllUserUrls(ctx context.Context, in *UserUrlsRequest, opts ...grpc.CallOption) (*UserUrlsResponse, error) {
	out := new(UserUrlsResponse)
	err := c.cc.Invoke(ctx, Url_GetAllUserUrls_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlServer is the server API for Url service.
// All implementations must embed UnimplementedUrlServer
// for forward compatibility
type UrlServer interface {
	CreateShortUrl(context.Context, *UrlRequest) (*UrlResponse, error)
	GetUrlByKey(context.Context, *UrlByKeyRequest) (*UrlResponse, error)
	GetAllUserUrls(context.Context, *UserUrlsRequest) (*UserUrlsResponse, error)
	mustEmbedUnimplementedUrlServer()
}

// UnimplementedUrlServer must be embedded to have forward compatible implementations.
type UnimplementedUrlServer struct {
}

func (UnimplementedUrlServer) CreateShortUrl(context.Context, *UrlRequest) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortUrl not implemented")
}
func (UnimplementedUrlServer) GetUrlByKey(context.Context, *UrlByKeyRequest) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrlByKey not implemented")
}
func (UnimplementedUrlServer) GetAllUserUrls(context.Context, *UserUrlsRequest) (*UserUrlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserUrls not implemented")
}
func (UnimplementedUrlServer) mustEmbedUnimplementedUrlServer() {}

// UnsafeUrlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlServer will
// result in compilation errors.
type UnsafeUrlServer interface {
	mustEmbedUnimplementedUrlServer()
}

func RegisterUrlServer(s grpc.ServiceRegistrar, srv UrlServer) {
	s.RegisterService(&Url_ServiceDesc, srv)
}

func _Url_CreateShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServer).CreateShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Url_CreateShortUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServer).CreateShortUrl(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Url_GetUrlByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServer).GetUrlByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Url_GetUrlByKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServer).GetUrlByKey(ctx, req.(*UrlByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Url_GetAllUserUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServer).GetAllUserUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Url_GetAllUserUrls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServer).GetAllUserUrls(ctx, req.(*UserUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Url_ServiceDesc is the grpc.ServiceDesc for Url service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Url_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Url",
	HandlerType: (*UrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortUrl",
			Handler:    _Url_CreateShortUrl_Handler,
		},
		{
			MethodName: "GetUrlByKey",
			Handler:    _Url_GetUrlByKey_Handler,
		},
		{
			MethodName: "GetAllUserUrls",
			Handler:    _Url_GetAllUserUrls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "url/url.proto",
}
