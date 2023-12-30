// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: gommerce/social/v1beta/invitations.proto

package social_v1beta

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Invitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviterId string                  `protobuf:"bytes,1,opt,name=inviter_id,json=inviterId,proto3" json:"inviter_id,omitempty"`
	InviteeId string                  `protobuf:"bytes,2,opt,name=invitee_id,json=inviteeId,proto3" json:"invitee_id,omitempty"`
	ChannelId *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	CreatedAt *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Invitation) Reset() {
	*x = Invitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invitation) ProtoMessage() {}

func (x *Invitation) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invitation.ProtoReflect.Descriptor instead.
func (*Invitation) Descriptor() ([]byte, []int) {
	return file_gommerce_social_v1beta_invitations_proto_rawDescGZIP(), []int{0}
}

func (x *Invitation) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

func (x *Invitation) GetInviteeId() string {
	if x != nil {
		return x.InviteeId
	}
	return ""
}

func (x *Invitation) GetChannelId() *wrapperspb.StringValue {
	if x != nil {
		return x.ChannelId
	}
	return nil
}

func (x *Invitation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviterId string                  `protobuf:"bytes,1,opt,name=inviter_id,json=inviterId,proto3" json:"inviter_id,omitempty"`
	InviteeId string                  `protobuf:"bytes,2,opt,name=invitee_id,json=inviteeId,proto3" json:"invitee_id,omitempty"`
	ChannelId *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *CreateInvitationRequest) Reset() {
	*x = CreateInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvitationRequest) ProtoMessage() {}

func (x *CreateInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvitationRequest.ProtoReflect.Descriptor instead.
func (*CreateInvitationRequest) Descriptor() ([]byte, []int) {
	return file_gommerce_social_v1beta_invitations_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInvitationRequest) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

func (x *CreateInvitationRequest) GetInviteeId() string {
	if x != nil {
		return x.InviteeId
	}
	return ""
}

func (x *CreateInvitationRequest) GetChannelId() *wrapperspb.StringValue {
	if x != nil {
		return x.ChannelId
	}
	return nil
}

type CreateInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateInvitationResponse) Reset() {
	*x = CreateInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvitationResponse) ProtoMessage() {}

func (x *CreateInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvitationResponse.ProtoReflect.Descriptor instead.
func (*CreateInvitationResponse) Descriptor() ([]byte, []int) {
	return file_gommerce_social_v1beta_invitations_proto_rawDescGZIP(), []int{2}
}

type ListInvitationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Sort string `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ListInvitationsRequest) Reset() {
	*x = ListInvitationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvitationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvitationsRequest) ProtoMessage() {}

func (x *ListInvitationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvitationsRequest.ProtoReflect.Descriptor instead.
func (*ListInvitationsRequest) Descriptor() ([]byte, []int) {
	return file_gommerce_social_v1beta_invitations_proto_rawDescGZIP(), []int{3}
}

func (x *ListInvitationsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListInvitationsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListInvitationsRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type ListInvitationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Invitation `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Page  int32         `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size  int32         `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Total int64         `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListInvitationsResponse) Reset() {
	*x = ListInvitationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvitationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvitationsResponse) ProtoMessage() {}

func (x *ListInvitationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gommerce_social_v1beta_invitations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvitationsResponse.ProtoReflect.Descriptor instead.
func (*ListInvitationsResponse) Descriptor() ([]byte, []int) {
	return file_gommerce_social_v1beta_invitations_proto_rawDescGZIP(), []int{4}
}

func (x *ListInvitationsResponse) GetItems() []*Invitation {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListInvitationsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListInvitationsResponse) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListInvitationsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_gommerce_social_v1beta_invitations_proto protoreflect.FileDescriptor

var file_gommerce_social_v1beta_invitations_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x49, 0x64,
	0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65,
	0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22,
	0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x22, 0x91, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x82, 0x02, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2f, 0x2e, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x61, 0x6c, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gommerce_social_v1beta_invitations_proto_rawDescOnce sync.Once
	file_gommerce_social_v1beta_invitations_proto_rawDescData = file_gommerce_social_v1beta_invitations_proto_rawDesc
)

func file_gommerce_social_v1beta_invitations_proto_rawDescGZIP() []byte {
	file_gommerce_social_v1beta_invitations_proto_rawDescOnce.Do(func() {
		file_gommerce_social_v1beta_invitations_proto_rawDescData = protoimpl.X.CompressGZIP(file_gommerce_social_v1beta_invitations_proto_rawDescData)
	})
	return file_gommerce_social_v1beta_invitations_proto_rawDescData
}

var file_gommerce_social_v1beta_invitations_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gommerce_social_v1beta_invitations_proto_goTypes = []interface{}{
	(*Invitation)(nil),               // 0: gommerce.social.v1beta.Invitation
	(*CreateInvitationRequest)(nil),  // 1: gommerce.social.v1beta.CreateInvitationRequest
	(*CreateInvitationResponse)(nil), // 2: gommerce.social.v1beta.CreateInvitationResponse
	(*ListInvitationsRequest)(nil),   // 3: gommerce.social.v1beta.ListInvitationsRequest
	(*ListInvitationsResponse)(nil),  // 4: gommerce.social.v1beta.ListInvitationsResponse
	(*wrapperspb.StringValue)(nil),   // 5: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
}
var file_gommerce_social_v1beta_invitations_proto_depIdxs = []int32{
	5, // 0: gommerce.social.v1beta.Invitation.channel_id:type_name -> google.protobuf.StringValue
	6, // 1: gommerce.social.v1beta.Invitation.created_at:type_name -> google.protobuf.Timestamp
	5, // 2: gommerce.social.v1beta.CreateInvitationRequest.channel_id:type_name -> google.protobuf.StringValue
	0, // 3: gommerce.social.v1beta.ListInvitationsResponse.items:type_name -> gommerce.social.v1beta.Invitation
	1, // 4: gommerce.social.v1beta.InvitationService.CreateInvitation:input_type -> gommerce.social.v1beta.CreateInvitationRequest
	3, // 5: gommerce.social.v1beta.InvitationService.ListInvitations:input_type -> gommerce.social.v1beta.ListInvitationsRequest
	2, // 6: gommerce.social.v1beta.InvitationService.CreateInvitation:output_type -> gommerce.social.v1beta.CreateInvitationResponse
	4, // 7: gommerce.social.v1beta.InvitationService.ListInvitations:output_type -> gommerce.social.v1beta.ListInvitationsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gommerce_social_v1beta_invitations_proto_init() }
func file_gommerce_social_v1beta_invitations_proto_init() {
	if File_gommerce_social_v1beta_invitations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gommerce_social_v1beta_invitations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invitation); i {
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
		file_gommerce_social_v1beta_invitations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvitationRequest); i {
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
		file_gommerce_social_v1beta_invitations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvitationResponse); i {
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
		file_gommerce_social_v1beta_invitations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvitationsRequest); i {
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
		file_gommerce_social_v1beta_invitations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvitationsResponse); i {
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
			RawDescriptor: file_gommerce_social_v1beta_invitations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gommerce_social_v1beta_invitations_proto_goTypes,
		DependencyIndexes: file_gommerce_social_v1beta_invitations_proto_depIdxs,
		MessageInfos:      file_gommerce_social_v1beta_invitations_proto_msgTypes,
	}.Build()
	File_gommerce_social_v1beta_invitations_proto = out.File
	file_gommerce_social_v1beta_invitations_proto_rawDesc = nil
	file_gommerce_social_v1beta_invitations_proto_goTypes = nil
	file_gommerce_social_v1beta_invitations_proto_depIdxs = nil
}
