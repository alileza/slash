// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1/memo_service.proto

package apiv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Memo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId   int32                  `protobuf:"varint,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	RowStatus   RowStatus              `protobuf:"varint,5,opt,name=row_status,json=rowStatus,proto3,enum=slash.api.v1.RowStatus" json:"row_status,omitempty"`
	Name        string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Title       string                 `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Content     string                 `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	Tags        []string               `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	Visibility  Visibility             `protobuf:"varint,10,opt,name=visibility,proto3,enum=slash.api.v1.Visibility" json:"visibility,omitempty"`
}

func (x *Memo) Reset() {
	*x = Memo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memo) ProtoMessage() {}

func (x *Memo) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memo.ProtoReflect.Descriptor instead.
func (*Memo) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{0}
}

func (x *Memo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Memo) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Memo) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Memo) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Memo) GetRowStatus() RowStatus {
	if x != nil {
		return x.RowStatus
	}
	return RowStatus_ROW_STATUS_UNSPECIFIED
}

func (x *Memo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Memo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Memo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Memo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Memo) GetVisibility() Visibility {
	if x != nil {
		return x.Visibility
	}
	return Visibility_VISIBILITY_UNSPECIFIED
}

type ListMemosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMemosRequest) Reset() {
	*x = ListMemosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMemosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemosRequest) ProtoMessage() {}

func (x *ListMemosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemosRequest.ProtoReflect.Descriptor instead.
func (*ListMemosRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{1}
}

type ListMemosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memos []*Memo `protobuf:"bytes,1,rep,name=memos,proto3" json:"memos,omitempty"`
}

func (x *ListMemosResponse) Reset() {
	*x = ListMemosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMemosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemosResponse) ProtoMessage() {}

func (x *ListMemosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemosResponse.ProtoReflect.Descriptor instead.
func (*ListMemosResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListMemosResponse) GetMemos() []*Memo {
	if x != nil {
		return x.Memos
	}
	return nil
}

type GetMemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMemoRequest) Reset() {
	*x = GetMemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoRequest) ProtoMessage() {}

func (x *GetMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoRequest.ProtoReflect.Descriptor instead.
func (*GetMemoRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetMemoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memo *Memo `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *GetMemoResponse) Reset() {
	*x = GetMemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoResponse) ProtoMessage() {}

func (x *GetMemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoResponse.ProtoReflect.Descriptor instead.
func (*GetMemoResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetMemoResponse) GetMemo() *Memo {
	if x != nil {
		return x.Memo
	}
	return nil
}

type CreateMemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memo *Memo `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *CreateMemoRequest) Reset() {
	*x = CreateMemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemoRequest) ProtoMessage() {}

func (x *CreateMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemoRequest.ProtoReflect.Descriptor instead.
func (*CreateMemoRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateMemoRequest) GetMemo() *Memo {
	if x != nil {
		return x.Memo
	}
	return nil
}

type CreateMemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memo *Memo `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *CreateMemoResponse) Reset() {
	*x = CreateMemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemoResponse) ProtoMessage() {}

func (x *CreateMemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemoResponse.ProtoReflect.Descriptor instead.
func (*CreateMemoResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateMemoResponse) GetMemo() *Memo {
	if x != nil {
		return x.Memo
	}
	return nil
}

type UpdateMemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memo       *Memo                  `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateMemoRequest) Reset() {
	*x = UpdateMemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemoRequest) ProtoMessage() {}

func (x *UpdateMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemoRequest.ProtoReflect.Descriptor instead.
func (*UpdateMemoRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateMemoRequest) GetMemo() *Memo {
	if x != nil {
		return x.Memo
	}
	return nil
}

func (x *UpdateMemoRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateMemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memo *Memo `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *UpdateMemoResponse) Reset() {
	*x = UpdateMemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemoResponse) ProtoMessage() {}

func (x *UpdateMemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemoResponse.ProtoReflect.Descriptor instead.
func (*UpdateMemoResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateMemoResponse) GetMemo() *Memo {
	if x != nil {
		return x.Memo
	}
	return nil
}

type DeleteMemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMemoRequest) Reset() {
	*x = DeleteMemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemoRequest) ProtoMessage() {}

func (x *DeleteMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemoRequest.ProtoReflect.Descriptor instead.
func (*DeleteMemoRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMemoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteMemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMemoResponse) Reset() {
	*x = DeleteMemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_memo_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemoResponse) ProtoMessage() {}

func (x *DeleteMemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemoResponse.ProtoReflect.Descriptor instead.
func (*DeleteMemoResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_service_proto_rawDescGZIP(), []int{10}
}

var File_api_v1_memo_service_proto protoreflect.FileDescriptor

var file_api_v1_memo_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x02, 0x0a, 0x04, 0x4d, 0x65, 0x6d,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36,
	0x0a, 0x0a, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x6f, 0x77,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x38,
	0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x6d, 0x6f, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x3b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x22, 0x78, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x3c, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc7, 0x04, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6d, 0x6f, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x12, 0x67, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x1c, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0xda, 0x41, 0x02, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x6f, 0x12, 0x1f, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x12, 0x1f, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0xda, 0x41, 0x10, 0x6d, 0x65, 0x6d, 0x6f, 0x2c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x1a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x69, 0x64, 0x7d, 0x12, 0x70,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x1f, 0x2e, 0x73,
	0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0xda, 0x41, 0x02, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0xae, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x4d, 0x65, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x6c, 0x66, 0x68, 0x6f,
	0x73, 0x74, 0x65, 0x64, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x2e,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c, 0x41,
	0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_memo_service_proto_rawDescOnce sync.Once
	file_api_v1_memo_service_proto_rawDescData = file_api_v1_memo_service_proto_rawDesc
)

func file_api_v1_memo_service_proto_rawDescGZIP() []byte {
	file_api_v1_memo_service_proto_rawDescOnce.Do(func() {
		file_api_v1_memo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_memo_service_proto_rawDescData)
	})
	return file_api_v1_memo_service_proto_rawDescData
}

var file_api_v1_memo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1_memo_service_proto_goTypes = []interface{}{
	(*Memo)(nil),                  // 0: slash.api.v1.Memo
	(*ListMemosRequest)(nil),      // 1: slash.api.v1.ListMemosRequest
	(*ListMemosResponse)(nil),     // 2: slash.api.v1.ListMemosResponse
	(*GetMemoRequest)(nil),        // 3: slash.api.v1.GetMemoRequest
	(*GetMemoResponse)(nil),       // 4: slash.api.v1.GetMemoResponse
	(*CreateMemoRequest)(nil),     // 5: slash.api.v1.CreateMemoRequest
	(*CreateMemoResponse)(nil),    // 6: slash.api.v1.CreateMemoResponse
	(*UpdateMemoRequest)(nil),     // 7: slash.api.v1.UpdateMemoRequest
	(*UpdateMemoResponse)(nil),    // 8: slash.api.v1.UpdateMemoResponse
	(*DeleteMemoRequest)(nil),     // 9: slash.api.v1.DeleteMemoRequest
	(*DeleteMemoResponse)(nil),    // 10: slash.api.v1.DeleteMemoResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(RowStatus)(0),                // 12: slash.api.v1.RowStatus
	(Visibility)(0),               // 13: slash.api.v1.Visibility
	(*fieldmaskpb.FieldMask)(nil), // 14: google.protobuf.FieldMask
}
var file_api_v1_memo_service_proto_depIdxs = []int32{
	11, // 0: slash.api.v1.Memo.created_time:type_name -> google.protobuf.Timestamp
	11, // 1: slash.api.v1.Memo.updated_time:type_name -> google.protobuf.Timestamp
	12, // 2: slash.api.v1.Memo.row_status:type_name -> slash.api.v1.RowStatus
	13, // 3: slash.api.v1.Memo.visibility:type_name -> slash.api.v1.Visibility
	0,  // 4: slash.api.v1.ListMemosResponse.memos:type_name -> slash.api.v1.Memo
	0,  // 5: slash.api.v1.GetMemoResponse.memo:type_name -> slash.api.v1.Memo
	0,  // 6: slash.api.v1.CreateMemoRequest.memo:type_name -> slash.api.v1.Memo
	0,  // 7: slash.api.v1.CreateMemoResponse.memo:type_name -> slash.api.v1.Memo
	0,  // 8: slash.api.v1.UpdateMemoRequest.memo:type_name -> slash.api.v1.Memo
	14, // 9: slash.api.v1.UpdateMemoRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 10: slash.api.v1.UpdateMemoResponse.memo:type_name -> slash.api.v1.Memo
	1,  // 11: slash.api.v1.MemoService.ListMemos:input_type -> slash.api.v1.ListMemosRequest
	3,  // 12: slash.api.v1.MemoService.GetMemo:input_type -> slash.api.v1.GetMemoRequest
	5,  // 13: slash.api.v1.MemoService.CreateMemo:input_type -> slash.api.v1.CreateMemoRequest
	7,  // 14: slash.api.v1.MemoService.UpdateMemo:input_type -> slash.api.v1.UpdateMemoRequest
	9,  // 15: slash.api.v1.MemoService.DeleteMemo:input_type -> slash.api.v1.DeleteMemoRequest
	2,  // 16: slash.api.v1.MemoService.ListMemos:output_type -> slash.api.v1.ListMemosResponse
	4,  // 17: slash.api.v1.MemoService.GetMemo:output_type -> slash.api.v1.GetMemoResponse
	6,  // 18: slash.api.v1.MemoService.CreateMemo:output_type -> slash.api.v1.CreateMemoResponse
	8,  // 19: slash.api.v1.MemoService.UpdateMemo:output_type -> slash.api.v1.UpdateMemoResponse
	10, // 20: slash.api.v1.MemoService.DeleteMemo:output_type -> slash.api.v1.DeleteMemoResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_v1_memo_service_proto_init() }
func file_api_v1_memo_service_proto_init() {
	if File_api_v1_memo_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_memo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memo); i {
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
		file_api_v1_memo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMemosRequest); i {
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
		file_api_v1_memo_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMemosResponse); i {
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
		file_api_v1_memo_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemoRequest); i {
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
		file_api_v1_memo_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemoResponse); i {
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
		file_api_v1_memo_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMemoRequest); i {
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
		file_api_v1_memo_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMemoResponse); i {
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
		file_api_v1_memo_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemoRequest); i {
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
		file_api_v1_memo_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemoResponse); i {
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
		file_api_v1_memo_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemoRequest); i {
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
		file_api_v1_memo_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemoResponse); i {
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
			RawDescriptor: file_api_v1_memo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_memo_service_proto_goTypes,
		DependencyIndexes: file_api_v1_memo_service_proto_depIdxs,
		MessageInfos:      file_api_v1_memo_service_proto_msgTypes,
	}.Build()
	File_api_v1_memo_service_proto = out.File
	file_api_v1_memo_service_proto_rawDesc = nil
	file_api_v1_memo_service_proto_goTypes = nil
	file_api_v1_memo_service_proto_depIdxs = nil
}
