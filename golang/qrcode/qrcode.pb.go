// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: qrcode/qrcode.proto

package qrcode

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateQRCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationUrl  string `protobuf:"bytes,1,opt,name=destinationUrl,proto3" json:"destinationUrl,omitempty"`
	ShortUrl        string `protobuf:"bytes,2,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
	BackgroundColor string `protobuf:"bytes,3,opt,name=backgroundColor,proto3" json:"backgroundColor,omitempty"`
	ForegroundColor string `protobuf:"bytes,4,opt,name=foregroundColor,proto3" json:"foregroundColor,omitempty"`
	LogoPath        string `protobuf:"bytes,5,opt,name=logoPath,proto3" json:"logoPath,omitempty"`
	FrameColor      string `protobuf:"bytes,6,opt,name=frameColor,proto3" json:"frameColor,omitempty"`
	FrameText       string `protobuf:"bytes,7,opt,name=frameText,proto3" json:"frameText,omitempty"`
	Branded         bool   `protobuf:"varint,8,opt,name=branded,proto3" json:"branded,omitempty"`
	UserId          string `protobuf:"bytes,9,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateQRCodeRequest) Reset() {
	*x = CreateQRCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_qrcode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQRCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQRCodeRequest) ProtoMessage() {}

func (x *CreateQRCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_qrcode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQRCodeRequest.ProtoReflect.Descriptor instead.
func (*CreateQRCodeRequest) Descriptor() ([]byte, []int) {
	return file_qrcode_qrcode_proto_rawDescGZIP(), []int{0}
}

func (x *CreateQRCodeRequest) GetDestinationUrl() string {
	if x != nil {
		return x.DestinationUrl
	}
	return ""
}

func (x *CreateQRCodeRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *CreateQRCodeRequest) GetBackgroundColor() string {
	if x != nil {
		return x.BackgroundColor
	}
	return ""
}

func (x *CreateQRCodeRequest) GetForegroundColor() string {
	if x != nil {
		return x.ForegroundColor
	}
	return ""
}

func (x *CreateQRCodeRequest) GetLogoPath() string {
	if x != nil {
		return x.LogoPath
	}
	return ""
}

func (x *CreateQRCodeRequest) GetFrameColor() string {
	if x != nil {
		return x.FrameColor
	}
	return ""
}

func (x *CreateQRCodeRequest) GetFrameText() string {
	if x != nil {
		return x.FrameText
	}
	return ""
}

func (x *CreateQRCodeRequest) GetBranded() bool {
	if x != nil {
		return x.Branded
	}
	return false
}

func (x *CreateQRCodeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateQRCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationUrl  string `protobuf:"bytes,1,opt,name=destinationUrl,proto3" json:"destinationUrl,omitempty"`
	ShortUrl        string `protobuf:"bytes,2,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
	BackgroundColor string `protobuf:"bytes,3,opt,name=backgroundColor,proto3" json:"backgroundColor,omitempty"`
	ForegroundColor string `protobuf:"bytes,4,opt,name=foregroundColor,proto3" json:"foregroundColor,omitempty"`
	LogoUrl         string `protobuf:"bytes,5,opt,name=logoUrl,proto3" json:"logoUrl,omitempty"`
	FrameColor      string `protobuf:"bytes,6,opt,name=frameColor,proto3" json:"frameColor,omitempty"`
	FrameText       string `protobuf:"bytes,7,opt,name=frameText,proto3" json:"frameText,omitempty"`
	Branded         bool   `protobuf:"varint,8,opt,name=branded,proto3" json:"branded,omitempty"`
	Id              string `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string `protobuf:"bytes,10,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpdateQRCodeRequest) Reset() {
	*x = UpdateQRCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_qrcode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQRCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQRCodeRequest) ProtoMessage() {}

func (x *UpdateQRCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_qrcode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQRCodeRequest.ProtoReflect.Descriptor instead.
func (*UpdateQRCodeRequest) Descriptor() ([]byte, []int) {
	return file_qrcode_qrcode_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateQRCodeRequest) GetDestinationUrl() string {
	if x != nil {
		return x.DestinationUrl
	}
	return ""
}

func (x *UpdateQRCodeRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *UpdateQRCodeRequest) GetBackgroundColor() string {
	if x != nil {
		return x.BackgroundColor
	}
	return ""
}

func (x *UpdateQRCodeRequest) GetForegroundColor() string {
	if x != nil {
		return x.ForegroundColor
	}
	return ""
}

func (x *UpdateQRCodeRequest) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *UpdateQRCodeRequest) GetFrameColor() string {
	if x != nil {
		return x.FrameColor
	}
	return ""
}

func (x *UpdateQRCodeRequest) GetFrameText() string {
	if x != nil {
		return x.FrameText
	}
	return ""
}

func (x *UpdateQRCodeRequest) GetBranded() bool {
	if x != nil {
		return x.Branded
	}
	return false
}

func (x *UpdateQRCodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateQRCodeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetQRCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetQRCodeRequest) Reset() {
	*x = GetQRCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_qrcode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQRCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQRCodeRequest) ProtoMessage() {}

func (x *GetQRCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_qrcode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQRCodeRequest.ProtoReflect.Descriptor instead.
func (*GetQRCodeRequest) Descriptor() ([]byte, []int) {
	return file_qrcode_qrcode_proto_rawDescGZIP(), []int{2}
}

func (x *GetQRCodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type QRCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationUrl  string `protobuf:"bytes,1,opt,name=destinationUrl,proto3" json:"destinationUrl,omitempty"`
	ShortUrl        string `protobuf:"bytes,2,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
	BackgroundColor string `protobuf:"bytes,3,opt,name=backgroundColor,proto3" json:"backgroundColor,omitempty"`
	ForegroundColor string `protobuf:"bytes,4,opt,name=foregroundColor,proto3" json:"foregroundColor,omitempty"`
	LogoPath        string `protobuf:"bytes,5,opt,name=logoPath,proto3" json:"logoPath,omitempty"`
	FrameColor      string `protobuf:"bytes,6,opt,name=frameColor,proto3" json:"frameColor,omitempty"`
	FrameText       string `protobuf:"bytes,7,opt,name=frameText,proto3" json:"frameText,omitempty"`
	Branded         bool   `protobuf:"varint,8,opt,name=branded,proto3" json:"branded,omitempty"`
	QrcodeUrl       string `protobuf:"bytes,9,opt,name=qrcodeUrl,proto3" json:"qrcodeUrl,omitempty"`
	Id              string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *QRCodeResponse) Reset() {
	*x = QRCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_qrcode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRCodeResponse) ProtoMessage() {}

func (x *QRCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_qrcode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRCodeResponse.ProtoReflect.Descriptor instead.
func (*QRCodeResponse) Descriptor() ([]byte, []int) {
	return file_qrcode_qrcode_proto_rawDescGZIP(), []int{3}
}

func (x *QRCodeResponse) GetDestinationUrl() string {
	if x != nil {
		return x.DestinationUrl
	}
	return ""
}

func (x *QRCodeResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *QRCodeResponse) GetBackgroundColor() string {
	if x != nil {
		return x.BackgroundColor
	}
	return ""
}

func (x *QRCodeResponse) GetForegroundColor() string {
	if x != nil {
		return x.ForegroundColor
	}
	return ""
}

func (x *QRCodeResponse) GetLogoPath() string {
	if x != nil {
		return x.LogoPath
	}
	return ""
}

func (x *QRCodeResponse) GetFrameColor() string {
	if x != nil {
		return x.FrameColor
	}
	return ""
}

func (x *QRCodeResponse) GetFrameText() string {
	if x != nil {
		return x.FrameText
	}
	return ""
}

func (x *QRCodeResponse) GetBranded() bool {
	if x != nil {
		return x.Branded
	}
	return false
}

func (x *QRCodeResponse) GetQrcodeUrl() string {
	if x != nil {
		return x.QrcodeUrl
	}
	return ""
}

func (x *QRCodeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type QRCodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QrCodes []*QRCodeResponse `protobuf:"bytes,1,rep,name=qrCodes,proto3" json:"qrCodes,omitempty"`
}

func (x *QRCodesResponse) Reset() {
	*x = QRCodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_qrcode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRCodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRCodesResponse) ProtoMessage() {}

func (x *QRCodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_qrcode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRCodesResponse.ProtoReflect.Descriptor instead.
func (*QRCodesResponse) Descriptor() ([]byte, []int) {
	return file_qrcode_qrcode_proto_rawDescGZIP(), []int{4}
}

func (x *QRCodesResponse) GetQrCodes() []*QRCodeResponse {
	if x != nil {
		return x.QrCodes
	}
	return nil
}

type QRCodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *QRCodesRequest) Reset() {
	*x = QRCodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_qrcode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRCodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRCodesRequest) ProtoMessage() {}

func (x *QRCodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_qrcode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRCodesRequest.ProtoReflect.Descriptor instead.
func (*QRCodesRequest) Descriptor() ([]byte, []int) {
	return file_qrcode_qrcode_proto_rawDescGZIP(), []int{5}
}

func (x *QRCodesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type QRCodeLogoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunkData,proto3" json:"chunkData,omitempty"`
}

func (x *QRCodeLogoRequest) Reset() {
	*x = QRCodeLogoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_qrcode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRCodeLogoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRCodeLogoRequest) ProtoMessage() {}

func (x *QRCodeLogoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_qrcode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRCodeLogoRequest.ProtoReflect.Descriptor instead.
func (*QRCodeLogoRequest) Descriptor() ([]byte, []int) {
	return file_qrcode_qrcode_proto_rawDescGZIP(), []int{6}
}

func (x *QRCodeLogoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QRCodeLogoRequest) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

type QRCodeLogoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogoUrl  string `protobuf:"bytes,1,opt,name=logoUrl,proto3" json:"logoUrl,omitempty"`
	LogoPath string `protobuf:"bytes,2,opt,name=logoPath,proto3" json:"logoPath,omitempty"`
}

func (x *QRCodeLogoResponse) Reset() {
	*x = QRCodeLogoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qrcode_qrcode_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRCodeLogoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRCodeLogoResponse) ProtoMessage() {}

func (x *QRCodeLogoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qrcode_qrcode_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRCodeLogoResponse.ProtoReflect.Descriptor instead.
func (*QRCodeLogoResponse) Descriptor() ([]byte, []int) {
	return file_qrcode_qrcode_proto_rawDescGZIP(), []int{7}
}

func (x *QRCodeLogoResponse) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *QRCodeLogoResponse) GetLogoPath() string {
	if x != nil {
		return x.LogoPath
	}
	return ""
}

var File_qrcode_qrcode_proto protoreflect.FileDescriptor

var file_qrcode_qrcode_proto_rawDesc = []byte{
	0x0a, 0x13, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x28, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x6f, 0x72,
	0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xc7, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x62,
	0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x65,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xca, 0x02,
	0x0a, 0x0e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62,
	0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x28,
	0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x6f,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0f, 0x51, 0x52,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x07, 0x71, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x71, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x51, 0x52, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x49, 0x0a, 0x11, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a,
	0x12, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x32, 0xa9, 0x03, 0x0a, 0x06, 0x51, 0x52,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x51, 0x52, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x0f, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x12, 0x57, 0x0a,
	0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01,
	0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x51, 0x52, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x58, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x1a, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x12,
	0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a,
	0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x6c,
	0x6f, 0x67, 0x6f, 0x28, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6a, 0x75, 0x6e, 0x77, 0x61, 0x31, 0x2f, 0x66, 0x75, 0x70, 0x69,
	0x2e, 0x74, 0x7a, 0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qrcode_qrcode_proto_rawDescOnce sync.Once
	file_qrcode_qrcode_proto_rawDescData = file_qrcode_qrcode_proto_rawDesc
)

func file_qrcode_qrcode_proto_rawDescGZIP() []byte {
	file_qrcode_qrcode_proto_rawDescOnce.Do(func() {
		file_qrcode_qrcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_qrcode_qrcode_proto_rawDescData)
	})
	return file_qrcode_qrcode_proto_rawDescData
}

var file_qrcode_qrcode_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_qrcode_qrcode_proto_goTypes = []interface{}{
	(*CreateQRCodeRequest)(nil), // 0: CreateQRCodeRequest
	(*UpdateQRCodeRequest)(nil), // 1: UpdateQRCodeRequest
	(*GetQRCodeRequest)(nil),    // 2: GetQRCodeRequest
	(*QRCodeResponse)(nil),      // 3: QRCodeResponse
	(*QRCodesResponse)(nil),     // 4: QRCodesResponse
	(*QRCodesRequest)(nil),      // 5: QRCodesRequest
	(*QRCodeLogoRequest)(nil),   // 6: QRCodeLogoRequest
	(*QRCodeLogoResponse)(nil),  // 7: QRCodeLogoResponse
}
var file_qrcode_qrcode_proto_depIdxs = []int32{
	3, // 0: QRCodesResponse.qrCodes:type_name -> QRCodeResponse
	5, // 1: QRCode.GetQRCodes:input_type -> QRCodesRequest
	0, // 2: QRCode.GenerateQRCode:input_type -> CreateQRCodeRequest
	2, // 3: QRCode.GetQRCode:input_type -> GetQRCodeRequest
	1, // 4: QRCode.UpdateQRCode:input_type -> UpdateQRCodeRequest
	6, // 5: QRCode.uploadQRCodeLogo:input_type -> QRCodeLogoRequest
	4, // 6: QRCode.GetQRCodes:output_type -> QRCodesResponse
	3, // 7: QRCode.GenerateQRCode:output_type -> QRCodeResponse
	3, // 8: QRCode.GetQRCode:output_type -> QRCodeResponse
	3, // 9: QRCode.UpdateQRCode:output_type -> QRCodeResponse
	7, // 10: QRCode.uploadQRCodeLogo:output_type -> QRCodeLogoResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_qrcode_qrcode_proto_init() }
func file_qrcode_qrcode_proto_init() {
	if File_qrcode_qrcode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qrcode_qrcode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQRCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_qrcode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQRCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_qrcode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQRCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_qrcode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_qrcode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRCodesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_qrcode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRCodesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_qrcode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRCodeLogoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_qrcode_qrcode_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRCodeLogoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qrcode_qrcode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qrcode_qrcode_proto_goTypes,
		DependencyIndexes: file_qrcode_qrcode_proto_depIdxs,
		MessageInfos:      file_qrcode_qrcode_proto_msgTypes,
	}.Build()
	File_qrcode_qrcode_proto = out.File
	file_qrcode_qrcode_proto_rawDesc = nil
	file_qrcode_qrcode_proto_goTypes = nil
	file_qrcode_qrcode_proto_depIdxs = nil
}
