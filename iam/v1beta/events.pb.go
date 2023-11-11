// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: gommerce/iam/v1beta/events.proto

package iam_v1beta

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

type UserCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm  string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserCreatedEvent) Reset() {
	*x = UserCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_iam_v1beta_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreatedEvent) ProtoMessage() {}

func (x *UserCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_iam_v1beta_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreatedEvent.ProtoReflect.Descriptor instead.
func (*UserCreatedEvent) Descriptor() ([]byte, []int) {
	return file_gommerce_iam_v1beta_events_proto_rawDescGZIP(), []int{0}
}

func (x *UserCreatedEvent) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *UserCreatedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm  string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserUpdatedEvent) Reset() {
	*x = UserUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_iam_v1beta_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdatedEvent) ProtoMessage() {}

func (x *UserUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_iam_v1beta_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdatedEvent.ProtoReflect.Descriptor instead.
func (*UserUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_gommerce_iam_v1beta_events_proto_rawDescGZIP(), []int{1}
}

func (x *UserUpdatedEvent) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *UserUpdatedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserActivatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm  string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserActivatedEvent) Reset() {
	*x = UserActivatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_iam_v1beta_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserActivatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserActivatedEvent) ProtoMessage() {}

func (x *UserActivatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_iam_v1beta_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserActivatedEvent.ProtoReflect.Descriptor instead.
func (*UserActivatedEvent) Descriptor() ([]byte, []int) {
	return file_gommerce_iam_v1beta_events_proto_rawDescGZIP(), []int{2}
}

func (x *UserActivatedEvent) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *UserActivatedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserLoggedInEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm      string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Provider   string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	Identifier string `protobuf:"bytes,4,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *UserLoggedInEvent) Reset() {
	*x = UserLoggedInEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_iam_v1beta_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoggedInEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoggedInEvent) ProtoMessage() {}

func (x *UserLoggedInEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_iam_v1beta_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoggedInEvent.ProtoReflect.Descriptor instead.
func (*UserLoggedInEvent) Descriptor() ([]byte, []int) {
	return file_gommerce_iam_v1beta_events_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoggedInEvent) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *UserLoggedInEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLoggedInEvent) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *UserLoggedInEvent) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

var File_gommerce_iam_v1beta_events_proto protoreflect.FileDescriptor

var file_gommerce_iam_v1beta_events_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x22, 0x41, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a,
	0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x7e, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x49, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x6f, 0x72, 0x61, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x69, 0x61, 0x6d, 0x5f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gommerce_iam_v1beta_events_proto_rawDescOnce sync.Once
	file_gommerce_iam_v1beta_events_proto_rawDescData = file_gommerce_iam_v1beta_events_proto_rawDesc
)

func file_gommerce_iam_v1beta_events_proto_rawDescGZIP() []byte {
	file_gommerce_iam_v1beta_events_proto_rawDescOnce.Do(func() {
		file_gommerce_iam_v1beta_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_gommerce_iam_v1beta_events_proto_rawDescData)
	})
	return file_gommerce_iam_v1beta_events_proto_rawDescData
}

var file_gommerce_iam_v1beta_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gommerce_iam_v1beta_events_proto_goTypes = []interface{}{
	(*UserCreatedEvent)(nil),   // 0: gommerce.iam.v1beta.UserCreatedEvent
	(*UserUpdatedEvent)(nil),   // 1: gommerce.iam.v1beta.UserUpdatedEvent
	(*UserActivatedEvent)(nil), // 2: gommerce.iam.v1beta.UserActivatedEvent
	(*UserLoggedInEvent)(nil),  // 3: gommerce.iam.v1beta.UserLoggedInEvent
}
var file_gommerce_iam_v1beta_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gommerce_iam_v1beta_events_proto_init() }
func file_gommerce_iam_v1beta_events_proto_init() {
	if File_gommerce_iam_v1beta_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gommerce_iam_v1beta_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreatedEvent); i {
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
		file_gommerce_iam_v1beta_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdatedEvent); i {
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
		file_gommerce_iam_v1beta_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserActivatedEvent); i {
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
		file_gommerce_iam_v1beta_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoggedInEvent); i {
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
			RawDescriptor: file_gommerce_iam_v1beta_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gommerce_iam_v1beta_events_proto_goTypes,
		DependencyIndexes: file_gommerce_iam_v1beta_events_proto_depIdxs,
		MessageInfos:      file_gommerce_iam_v1beta_events_proto_msgTypes,
	}.Build()
	File_gommerce_iam_v1beta_events_proto = out.File
	file_gommerce_iam_v1beta_events_proto_rawDesc = nil
	file_gommerce_iam_v1beta_events_proto_goTypes = nil
	file_gommerce_iam_v1beta_events_proto_depIdxs = nil
}
