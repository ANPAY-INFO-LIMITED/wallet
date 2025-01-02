// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: proxy.proto

package litrpc

import (
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

type BakeSuperMacaroonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The root key ID suffix is the 4-byte suffix of the root key ID that will
	// be used to create the macaroon.
	RootKeyIdSuffix uint32 `protobuf:"varint,1,opt,name=root_key_id_suffix,json=rootKeyIdSuffix,proto3" json:"root_key_id_suffix,omitempty"`
	// Whether the macaroon should only contain read permissions.
	ReadOnly bool `protobuf:"varint,2,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
}

func (x *BakeSuperMacaroonRequest) Reset() {
	*x = BakeSuperMacaroonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BakeSuperMacaroonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BakeSuperMacaroonRequest) ProtoMessage() {}

func (x *BakeSuperMacaroonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BakeSuperMacaroonRequest.ProtoReflect.Descriptor instead.
func (*BakeSuperMacaroonRequest) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *BakeSuperMacaroonRequest) GetRootKeyIdSuffix() uint32 {
	if x != nil {
		return x.RootKeyIdSuffix
	}
	return 0
}

func (x *BakeSuperMacaroonRequest) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

type BakeSuperMacaroonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hex encoded macaroon.
	Macaroon string `protobuf:"bytes,1,opt,name=macaroon,proto3" json:"macaroon,omitempty"`
}

func (x *BakeSuperMacaroonResponse) Reset() {
	*x = BakeSuperMacaroonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BakeSuperMacaroonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BakeSuperMacaroonResponse) ProtoMessage() {}

func (x *BakeSuperMacaroonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BakeSuperMacaroonResponse.ProtoReflect.Descriptor instead.
func (*BakeSuperMacaroonResponse) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *BakeSuperMacaroonResponse) GetMacaroon() string {
	if x != nil {
		return x.Macaroon
	}
	return ""
}

type StopDaemonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopDaemonRequest) Reset() {
	*x = StopDaemonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopDaemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopDaemonRequest) ProtoMessage() {}

func (x *StopDaemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopDaemonRequest.ProtoReflect.Descriptor instead.
func (*StopDaemonRequest) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{2}
}

type StopDaemonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopDaemonResponse) Reset() {
	*x = StopDaemonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopDaemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopDaemonResponse) ProtoMessage() {}

func (x *StopDaemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopDaemonResponse.ProtoReflect.Descriptor instead.
func (*StopDaemonResponse) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{3}
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{4}
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the LiTd software that the node is running.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{5}
}

func (x *GetInfoResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_proxy_proto protoreflect.FileDescriptor

var file_proxy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c,
	0x69, 0x74, 0x72, 0x70, 0x63, 0x22, 0x64, 0x0a, 0x18, 0x42, 0x61, 0x6b, 0x65, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x4d, 0x61, 0x63, 0x61, 0x72, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64,
	0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72,
	0x6f, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x37, 0x0a, 0x19, 0x42,
	0x61, 0x6b, 0x65, 0x53, 0x75, 0x70, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x61, 0x72, 0x6f, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x61,
	0x72, 0x6f, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x61,
	0x72, 0x6f, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x74, 0x6f,
	0x70, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xe2,
	0x01, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x69,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x42, 0x61, 0x6b,
	0x65, 0x53, 0x75, 0x70, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x61, 0x72, 0x6f, 0x6f, 0x6e, 0x12, 0x20,
	0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x4d, 0x61, 0x63, 0x61, 0x72, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x53, 0x75,
	0x70, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x61, 0x72, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proxy_proto_rawDescOnce sync.Once
	file_proxy_proto_rawDescData = file_proxy_proto_rawDesc
)

func file_proxy_proto_rawDescGZIP() []byte {
	file_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_proto_rawDescData)
	})
	return file_proxy_proto_rawDescData
}

var file_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proxy_proto_goTypes = []any{
	(*BakeSuperMacaroonRequest)(nil),  // 0: litrpc.BakeSuperMacaroonRequest
	(*BakeSuperMacaroonResponse)(nil), // 1: litrpc.BakeSuperMacaroonResponse
	(*StopDaemonRequest)(nil),         // 2: litrpc.StopDaemonRequest
	(*StopDaemonResponse)(nil),        // 3: litrpc.StopDaemonResponse
	(*GetInfoRequest)(nil),            // 4: litrpc.GetInfoRequest
	(*GetInfoResponse)(nil),           // 5: litrpc.GetInfoResponse
}
var file_proxy_proto_depIdxs = []int32{
	4, // 0: litrpc.Proxy.GetInfo:input_type -> litrpc.GetInfoRequest
	2, // 1: litrpc.Proxy.StopDaemon:input_type -> litrpc.StopDaemonRequest
	0, // 2: litrpc.Proxy.BakeSuperMacaroon:input_type -> litrpc.BakeSuperMacaroonRequest
	5, // 3: litrpc.Proxy.GetInfo:output_type -> litrpc.GetInfoResponse
	3, // 4: litrpc.Proxy.StopDaemon:output_type -> litrpc.StopDaemonResponse
	1, // 5: litrpc.Proxy.BakeSuperMacaroon:output_type -> litrpc.BakeSuperMacaroonResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proxy_proto_init() }
func file_proxy_proto_init() {
	if File_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BakeSuperMacaroonRequest); i {
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
		file_proxy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BakeSuperMacaroonResponse); i {
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
		file_proxy_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StopDaemonRequest); i {
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
		file_proxy_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StopDaemonResponse); i {
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
		file_proxy_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetInfoRequest); i {
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
		file_proxy_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetInfoResponse); i {
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
			RawDescriptor: file_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_proto_depIdxs,
		MessageInfos:      file_proxy_proto_msgTypes,
	}.Build()
	File_proxy_proto = out.File
	file_proxy_proto_rawDesc = nil
	file_proxy_proto_goTypes = nil
	file_proxy_proto_depIdxs = nil
}
