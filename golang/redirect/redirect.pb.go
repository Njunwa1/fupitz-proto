// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: redirect/redirect.proto

package redirect

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

type RedirectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
}

func (x *RedirectRequest) Reset() {
	*x = RedirectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redirect_redirect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectRequest) ProtoMessage() {}

func (x *RedirectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redirect_redirect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectRequest.ProtoReflect.Descriptor instead.
func (*RedirectRequest) Descriptor() ([]byte, []int) {
	return file_redirect_redirect_proto_rawDescGZIP(), []int{0}
}

func (x *RedirectRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type RedirectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShortUrl    string `protobuf:"bytes,2,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
	WebUrl      string `protobuf:"bytes,3,opt,name=webUrl,proto3" json:"webUrl,omitempty"`
	IosUrl      string `protobuf:"bytes,4,opt,name=iosUrl,proto3" json:"iosUrl,omitempty"`
	AndroidUrl  string `protobuf:"bytes,5,opt,name=androidUrl,proto3" json:"androidUrl,omitempty"`
	Password    string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	ExpiryAt    string `protobuf:"bytes,7,opt,name=expiryAt,proto3" json:"expiryAt,omitempty"`
	CustomAlias string `protobuf:"bytes,8,opt,name=customAlias,proto3" json:"customAlias,omitempty"`
	Type        string `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	QrcodeUrl   string `protobuf:"bytes,10,opt,name=qrcodeUrl,proto3" json:"qrcodeUrl,omitempty"`
}

func (x *RedirectResponse) Reset() {
	*x = RedirectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redirect_redirect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectResponse) ProtoMessage() {}

func (x *RedirectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redirect_redirect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectResponse.ProtoReflect.Descriptor instead.
func (*RedirectResponse) Descriptor() ([]byte, []int) {
	return file_redirect_redirect_proto_rawDescGZIP(), []int{1}
}

func (x *RedirectResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RedirectResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *RedirectResponse) GetWebUrl() string {
	if x != nil {
		return x.WebUrl
	}
	return ""
}

func (x *RedirectResponse) GetIosUrl() string {
	if x != nil {
		return x.IosUrl
	}
	return ""
}

func (x *RedirectResponse) GetAndroidUrl() string {
	if x != nil {
		return x.AndroidUrl
	}
	return ""
}

func (x *RedirectResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RedirectResponse) GetExpiryAt() string {
	if x != nil {
		return x.ExpiryAt
	}
	return ""
}

func (x *RedirectResponse) GetCustomAlias() string {
	if x != nil {
		return x.CustomAlias
	}
	return ""
}

func (x *RedirectResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RedirectResponse) GetQrcodeUrl() string {
	if x != nil {
		return x.QrcodeUrl
	}
	return ""
}

var File_redirect_redirect_proto protoreflect.FileDescriptor

var file_redirect_redirect_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x9a, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x62, 0x55, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6f, 0x73, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6f, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x41, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x55, 0x72, 0x6c, 0x32, 0x5f, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12,
	0x53, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x10, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x72, 0x6c, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x7d, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x6a, 0x75, 0x6e, 0x77, 0x61, 0x31, 0x2f, 0x66, 0x75, 0x70, 0x69, 0x2e,
	0x74, 0x7a, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_redirect_redirect_proto_rawDescOnce sync.Once
	file_redirect_redirect_proto_rawDescData = file_redirect_redirect_proto_rawDesc
)

func file_redirect_redirect_proto_rawDescGZIP() []byte {
	file_redirect_redirect_proto_rawDescOnce.Do(func() {
		file_redirect_redirect_proto_rawDescData = protoimpl.X.CompressGZIP(file_redirect_redirect_proto_rawDescData)
	})
	return file_redirect_redirect_proto_rawDescData
}

var file_redirect_redirect_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_redirect_redirect_proto_goTypes = []interface{}{
	(*RedirectRequest)(nil),  // 0: RedirectRequest
	(*RedirectResponse)(nil), // 1: RedirectResponse
}
var file_redirect_redirect_proto_depIdxs = []int32{
	0, // 0: Redirect.Redirect:input_type -> RedirectRequest
	1, // 1: Redirect.Redirect:output_type -> RedirectResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_redirect_redirect_proto_init() }
func file_redirect_redirect_proto_init() {
	if File_redirect_redirect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_redirect_redirect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectRequest); i {
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
		file_redirect_redirect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectResponse); i {
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
			RawDescriptor: file_redirect_redirect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_redirect_redirect_proto_goTypes,
		DependencyIndexes: file_redirect_redirect_proto_depIdxs,
		MessageInfos:      file_redirect_redirect_proto_msgTypes,
	}.Build()
	File_redirect_redirect_proto = out.File
	file_redirect_redirect_proto_rawDesc = nil
	file_redirect_redirect_proto_goTypes = nil
	file_redirect_redirect_proto_depIdxs = nil
}
