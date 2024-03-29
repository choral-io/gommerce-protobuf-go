// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: gommerce/iam/v1beta/tokens.proto

package iam_v1beta

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CreateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope    []string                `protobuf:"bytes,1,rep,name=scope,proto3" json:"scope,omitempty"`
	Realm    string                  `protobuf:"bytes,2,opt,name=realm,proto3" json:"realm,omitempty"`
	Provider string                  `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	Username *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	IdToken  *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
}

func (x *CreateTokenRequest) Reset() {
	*x = CreateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_iam_v1beta_tokens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRequest) ProtoMessage() {}

func (x *CreateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_iam_v1beta_tokens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return file_gommerce_iam_v1beta_tokens_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTokenRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *CreateTokenRequest) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *CreateTokenRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CreateTokenRequest) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *CreateTokenRequest) GetPassword() *wrapperspb.StringValue {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *CreateTokenRequest) GetIdToken() *wrapperspb.StringValue {
	if x != nil {
		return x.IdToken
	}
	return nil
}

type CreateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenType    string `protobuf:"bytes,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresIn    int32  `protobuf:"varint,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *CreateTokenResponse) Reset() {
	*x = CreateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_iam_v1beta_tokens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenResponse) ProtoMessage() {}

func (x *CreateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_iam_v1beta_tokens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return file_gommerce_iam_v1beta_tokens_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTokenResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *CreateTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CreateTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *CreateTokenResponse) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope        []string `protobuf:"bytes,1,rep,name=scope,proto3" json:"scope,omitempty"`
	Realm        string   `protobuf:"bytes,2,opt,name=realm,proto3" json:"realm,omitempty"`
	RefreshToken string   `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_iam_v1beta_tokens_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_iam_v1beta_tokens_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_gommerce_iam_v1beta_tokens_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshTokenRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *RefreshTokenRequest) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenType    string `protobuf:"bytes,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresIn    int32  `protobuf:"varint,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_iam_v1beta_tokens_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_iam_v1beta_tokens_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_gommerce_iam_v1beta_tokens_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshTokenResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *RefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshTokenResponse) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_gommerce_iam_v1beta_tokens_proto protoreflect.FileDescriptor

var file_gommerce_iam_v1beta_tokens_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9b, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x23, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x38, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9b,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x22, 0x66, 0x0a, 0x13,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x49, 0x6e, 0x32, 0xda, 0x01, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x2e, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0c, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x6f, 0x72, 0x61, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x69, 0x61, 0x6d, 0x5f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gommerce_iam_v1beta_tokens_proto_rawDescOnce sync.Once
	file_gommerce_iam_v1beta_tokens_proto_rawDescData = file_gommerce_iam_v1beta_tokens_proto_rawDesc
)

func file_gommerce_iam_v1beta_tokens_proto_rawDescGZIP() []byte {
	file_gommerce_iam_v1beta_tokens_proto_rawDescOnce.Do(func() {
		file_gommerce_iam_v1beta_tokens_proto_rawDescData = protoimpl.X.CompressGZIP(file_gommerce_iam_v1beta_tokens_proto_rawDescData)
	})
	return file_gommerce_iam_v1beta_tokens_proto_rawDescData
}

var file_gommerce_iam_v1beta_tokens_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gommerce_iam_v1beta_tokens_proto_goTypes = []interface{}{
	(*CreateTokenRequest)(nil),     // 0: gommerce.iam.v1beta.CreateTokenRequest
	(*CreateTokenResponse)(nil),    // 1: gommerce.iam.v1beta.CreateTokenResponse
	(*RefreshTokenRequest)(nil),    // 2: gommerce.iam.v1beta.RefreshTokenRequest
	(*RefreshTokenResponse)(nil),   // 3: gommerce.iam.v1beta.RefreshTokenResponse
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
}
var file_gommerce_iam_v1beta_tokens_proto_depIdxs = []int32{
	4, // 0: gommerce.iam.v1beta.CreateTokenRequest.username:type_name -> google.protobuf.StringValue
	4, // 1: gommerce.iam.v1beta.CreateTokenRequest.password:type_name -> google.protobuf.StringValue
	4, // 2: gommerce.iam.v1beta.CreateTokenRequest.id_token:type_name -> google.protobuf.StringValue
	0, // 3: gommerce.iam.v1beta.TokensService.CreateToken:input_type -> gommerce.iam.v1beta.CreateTokenRequest
	2, // 4: gommerce.iam.v1beta.TokensService.RefreshToken:input_type -> gommerce.iam.v1beta.RefreshTokenRequest
	1, // 5: gommerce.iam.v1beta.TokensService.CreateToken:output_type -> gommerce.iam.v1beta.CreateTokenResponse
	3, // 6: gommerce.iam.v1beta.TokensService.RefreshToken:output_type -> gommerce.iam.v1beta.RefreshTokenResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gommerce_iam_v1beta_tokens_proto_init() }
func file_gommerce_iam_v1beta_tokens_proto_init() {
	if File_gommerce_iam_v1beta_tokens_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gommerce_iam_v1beta_tokens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRequest); i {
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
		file_gommerce_iam_v1beta_tokens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenResponse); i {
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
		file_gommerce_iam_v1beta_tokens_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenRequest); i {
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
		file_gommerce_iam_v1beta_tokens_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenResponse); i {
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
			RawDescriptor: file_gommerce_iam_v1beta_tokens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gommerce_iam_v1beta_tokens_proto_goTypes,
		DependencyIndexes: file_gommerce_iam_v1beta_tokens_proto_depIdxs,
		MessageInfos:      file_gommerce_iam_v1beta_tokens_proto_msgTypes,
	}.Build()
	File_gommerce_iam_v1beta_tokens_proto = out.File
	file_gommerce_iam_v1beta_tokens_proto_rawDesc = nil
	file_gommerce_iam_v1beta_tokens_proto_goTypes = nil
	file_gommerce_iam_v1beta_tokens_proto_depIdxs = nil
}
