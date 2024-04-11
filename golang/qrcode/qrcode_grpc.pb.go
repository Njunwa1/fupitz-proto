// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: qrcode/qrcode.proto

package qrcode

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
	QRCode_GetQRCodes_FullMethodName       = "/QRCode/GetQRCodes"
	QRCode_GenerateQRCode_FullMethodName   = "/QRCode/GenerateQRCode"
	QRCode_GetQRCode_FullMethodName        = "/QRCode/GetQRCode"
	QRCode_UpdateQRCode_FullMethodName     = "/QRCode/UpdateQRCode"
	QRCode_UploadQRCodeLogo_FullMethodName = "/QRCode/uploadQRCodeLogo"
)

// QRCodeClient is the client API for QRCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QRCodeClient interface {
	GetQRCodes(ctx context.Context, in *QRCodesRequest, opts ...grpc.CallOption) (*QRCodesResponse, error)
	GenerateQRCode(ctx context.Context, in *CreateQRCodeRequest, opts ...grpc.CallOption) (*QRCodeResponse, error)
	GetQRCode(ctx context.Context, in *GetQRCodeRequest, opts ...grpc.CallOption) (*QRCodeResponse, error)
	UpdateQRCode(ctx context.Context, in *UpdateQRCodeRequest, opts ...grpc.CallOption) (*QRCodeResponse, error)
	UploadQRCodeLogo(ctx context.Context, opts ...grpc.CallOption) (QRCode_UploadQRCodeLogoClient, error)
}

type qRCodeClient struct {
	cc grpc.ClientConnInterface
}

func NewQRCodeClient(cc grpc.ClientConnInterface) QRCodeClient {
	return &qRCodeClient{cc}
}

func (c *qRCodeClient) GetQRCodes(ctx context.Context, in *QRCodesRequest, opts ...grpc.CallOption) (*QRCodesResponse, error) {
	out := new(QRCodesResponse)
	err := c.cc.Invoke(ctx, QRCode_GetQRCodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qRCodeClient) GenerateQRCode(ctx context.Context, in *CreateQRCodeRequest, opts ...grpc.CallOption) (*QRCodeResponse, error) {
	out := new(QRCodeResponse)
	err := c.cc.Invoke(ctx, QRCode_GenerateQRCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qRCodeClient) GetQRCode(ctx context.Context, in *GetQRCodeRequest, opts ...grpc.CallOption) (*QRCodeResponse, error) {
	out := new(QRCodeResponse)
	err := c.cc.Invoke(ctx, QRCode_GetQRCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qRCodeClient) UpdateQRCode(ctx context.Context, in *UpdateQRCodeRequest, opts ...grpc.CallOption) (*QRCodeResponse, error) {
	out := new(QRCodeResponse)
	err := c.cc.Invoke(ctx, QRCode_UpdateQRCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qRCodeClient) UploadQRCodeLogo(ctx context.Context, opts ...grpc.CallOption) (QRCode_UploadQRCodeLogoClient, error) {
	stream, err := c.cc.NewStream(ctx, &QRCode_ServiceDesc.Streams[0], QRCode_UploadQRCodeLogo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &qRCodeUploadQRCodeLogoClient{stream}
	return x, nil
}

type QRCode_UploadQRCodeLogoClient interface {
	Send(*QRCodeLogoRequest) error
	CloseAndRecv() (*QRCodeLogoResponse, error)
	grpc.ClientStream
}

type qRCodeUploadQRCodeLogoClient struct {
	grpc.ClientStream
}

func (x *qRCodeUploadQRCodeLogoClient) Send(m *QRCodeLogoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *qRCodeUploadQRCodeLogoClient) CloseAndRecv() (*QRCodeLogoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(QRCodeLogoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QRCodeServer is the server API for QRCode service.
// All implementations must embed UnimplementedQRCodeServer
// for forward compatibility
type QRCodeServer interface {
	GetQRCodes(context.Context, *QRCodesRequest) (*QRCodesResponse, error)
	GenerateQRCode(context.Context, *CreateQRCodeRequest) (*QRCodeResponse, error)
	GetQRCode(context.Context, *GetQRCodeRequest) (*QRCodeResponse, error)
	UpdateQRCode(context.Context, *UpdateQRCodeRequest) (*QRCodeResponse, error)
	UploadQRCodeLogo(QRCode_UploadQRCodeLogoServer) error
	mustEmbedUnimplementedQRCodeServer()
}

// UnimplementedQRCodeServer must be embedded to have forward compatible implementations.
type UnimplementedQRCodeServer struct {
}

func (UnimplementedQRCodeServer) GetQRCodes(context.Context, *QRCodesRequest) (*QRCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQRCodes not implemented")
}
func (UnimplementedQRCodeServer) GenerateQRCode(context.Context, *CreateQRCodeRequest) (*QRCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateQRCode not implemented")
}
func (UnimplementedQRCodeServer) GetQRCode(context.Context, *GetQRCodeRequest) (*QRCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQRCode not implemented")
}
func (UnimplementedQRCodeServer) UpdateQRCode(context.Context, *UpdateQRCodeRequest) (*QRCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQRCode not implemented")
}
func (UnimplementedQRCodeServer) UploadQRCodeLogo(QRCode_UploadQRCodeLogoServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadQRCodeLogo not implemented")
}
func (UnimplementedQRCodeServer) mustEmbedUnimplementedQRCodeServer() {}

// UnsafeQRCodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QRCodeServer will
// result in compilation errors.
type UnsafeQRCodeServer interface {
	mustEmbedUnimplementedQRCodeServer()
}

func RegisterQRCodeServer(s grpc.ServiceRegistrar, srv QRCodeServer) {
	s.RegisterService(&QRCode_ServiceDesc, srv)
}

func _QRCode_GetQRCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QRCodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QRCodeServer).GetQRCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QRCode_GetQRCodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QRCodeServer).GetQRCodes(ctx, req.(*QRCodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QRCode_GenerateQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQRCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QRCodeServer).GenerateQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QRCode_GenerateQRCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QRCodeServer).GenerateQRCode(ctx, req.(*CreateQRCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QRCode_GetQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQRCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QRCodeServer).GetQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QRCode_GetQRCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QRCodeServer).GetQRCode(ctx, req.(*GetQRCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QRCode_UpdateQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQRCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QRCodeServer).UpdateQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QRCode_UpdateQRCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QRCodeServer).UpdateQRCode(ctx, req.(*UpdateQRCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QRCode_UploadQRCodeLogo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QRCodeServer).UploadQRCodeLogo(&qRCodeUploadQRCodeLogoServer{stream})
}

type QRCode_UploadQRCodeLogoServer interface {
	SendAndClose(*QRCodeLogoResponse) error
	Recv() (*QRCodeLogoRequest, error)
	grpc.ServerStream
}

type qRCodeUploadQRCodeLogoServer struct {
	grpc.ServerStream
}

func (x *qRCodeUploadQRCodeLogoServer) SendAndClose(m *QRCodeLogoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *qRCodeUploadQRCodeLogoServer) Recv() (*QRCodeLogoRequest, error) {
	m := new(QRCodeLogoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QRCode_ServiceDesc is the grpc.ServiceDesc for QRCode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QRCode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "QRCode",
	HandlerType: (*QRCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQRCodes",
			Handler:    _QRCode_GetQRCodes_Handler,
		},
		{
			MethodName: "GenerateQRCode",
			Handler:    _QRCode_GenerateQRCode_Handler,
		},
		{
			MethodName: "GetQRCode",
			Handler:    _QRCode_GetQRCode_Handler,
		},
		{
			MethodName: "UpdateQRCode",
			Handler:    _QRCode_UpdateQRCode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "uploadQRCodeLogo",
			Handler:       _QRCode_UploadQRCodeLogo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "qrcode/qrcode.proto",
}