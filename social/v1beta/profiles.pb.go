// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: gommerce/social/v1beta/profiles.proto

package social_v1beta

import (
	date "github.com/choral-io/gommerce-protobuf-go/types/v1/date"
	gender "github.com/choral-io/gommerce-protobuf-go/types/v1/gender"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string                  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DisplayName  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	AvatarUrl    *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Gender       gender.Gender           `protobuf:"varint,4,opt,name=gender,proto3,enum=gommerce.types.v1.Gender" json:"gender,omitempty"`
	Birthdate    *date.Date              `protobuf:"bytes,5,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Introduction *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=introduction,proto3" json:"introduction,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_social_v1beta_profiles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_social_v1beta_profiles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_gommerce_social_v1beta_profiles_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Profile) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *Profile) GetAvatarUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.AvatarUrl
	}
	return nil
}

func (x *Profile) GetGender() gender.Gender {
	if x != nil {
		return x.Gender
	}
	return gender.Gender(0)
}

func (x *Profile) GetBirthdate() *date.Date {
	if x != nil {
		return x.Birthdate
	}
	return nil
}

func (x *Profile) GetIntroduction() *wrapperspb.StringValue {
	if x != nil {
		return x.Introduction
	}
	return nil
}

var File_gommerce_social_v1beta_profiles_proto protoreflect.FileDescriptor

var file_gommerce_social_v1beta_profiles_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a,
	0x1c, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c,
	0x12, 0x31, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x69, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x47, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x61,
	0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gommerce_social_v1beta_profiles_proto_rawDescOnce sync.Once
	file_gommerce_social_v1beta_profiles_proto_rawDescData = file_gommerce_social_v1beta_profiles_proto_rawDesc
)

func file_gommerce_social_v1beta_profiles_proto_rawDescGZIP() []byte {
	file_gommerce_social_v1beta_profiles_proto_rawDescOnce.Do(func() {
		file_gommerce_social_v1beta_profiles_proto_rawDescData = protoimpl.X.CompressGZIP(file_gommerce_social_v1beta_profiles_proto_rawDescData)
	})
	return file_gommerce_social_v1beta_profiles_proto_rawDescData
}

var file_gommerce_social_v1beta_profiles_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gommerce_social_v1beta_profiles_proto_goTypes = []interface{}{
	(*Profile)(nil),                // 0: gommerce.social.v1beta.Profile
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(gender.Gender)(0),             // 2: gommerce.types.v1.Gender
	(*date.Date)(nil),              // 3: gommerce.types.v1.Date
}
var file_gommerce_social_v1beta_profiles_proto_depIdxs = []int32{
	1, // 0: gommerce.social.v1beta.Profile.display_name:type_name -> google.protobuf.StringValue
	1, // 1: gommerce.social.v1beta.Profile.avatar_url:type_name -> google.protobuf.StringValue
	2, // 2: gommerce.social.v1beta.Profile.gender:type_name -> gommerce.types.v1.Gender
	3, // 3: gommerce.social.v1beta.Profile.birthdate:type_name -> gommerce.types.v1.Date
	1, // 4: gommerce.social.v1beta.Profile.introduction:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_gommerce_social_v1beta_profiles_proto_init() }
func file_gommerce_social_v1beta_profiles_proto_init() {
	if File_gommerce_social_v1beta_profiles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gommerce_social_v1beta_profiles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
			RawDescriptor: file_gommerce_social_v1beta_profiles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gommerce_social_v1beta_profiles_proto_goTypes,
		DependencyIndexes: file_gommerce_social_v1beta_profiles_proto_depIdxs,
		MessageInfos:      file_gommerce_social_v1beta_profiles_proto_msgTypes,
	}.Build()
	File_gommerce_social_v1beta_profiles_proto = out.File
	file_gommerce_social_v1beta_profiles_proto_rawDesc = nil
	file_gommerce_social_v1beta_profiles_proto_goTypes = nil
	file_gommerce_social_v1beta_profiles_proto_depIdxs = nil
}
