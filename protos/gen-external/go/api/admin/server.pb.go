// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: api/admin/server.proto

package dndbotv1alpha1api

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

type GetRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRoomRequest) Reset() {
	*x = GetRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomRequest) ProtoMessage() {}

func (x *GetRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomRequest.ProtoReflect.Descriptor instead.
func (*GetRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{0}
}

func (x *GetRoomRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *GetRoomResponse) Reset() {
	*x = GetRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomResponse) ProtoMessage() {}

func (x *GetRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomResponse.ProtoReflect.Descriptor instead.
func (*GetRoomResponse) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type ListRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Limit     int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListRoomsRequest) Reset() {
	*x = ListRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsRequest) ProtoMessage() {}

func (x *ListRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsRequest.ProtoReflect.Descriptor instead.
func (*ListRoomsRequest) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{2}
}

func (x *ListRoomsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListRoomsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRoomsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms         []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	NextPageToken string  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListRoomsResponse) Reset() {
	*x = ListRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsResponse) ProtoMessage() {}

func (x *ListRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsResponse.ProtoReflect.Descriptor instead.
func (*ListRoomsResponse) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{3}
}

func (x *ListRoomsResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

func (x *ListRoomsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRoomRequest) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type UpdateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *UpdateRoomRequest) Reset() {
	*x = UpdateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomRequest) ProtoMessage() {}

func (x *UpdateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRoomRequest) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type UpdateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *UpdateRoomResponse) Reset() {
	*x = UpdateRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomResponse) ProtoMessage() {}

func (x *UpdateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomResponse.ProtoReflect.Descriptor instead.
func (*UpdateRoomResponse) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type DeleteRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRoomRequest) Reset() {
	*x = DeleteRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoomRequest) ProtoMessage() {}

func (x *DeleteRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoomRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRoomRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRoomResponse) Reset() {
	*x = DeleteRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoomResponse) ProtoMessage() {}

func (x *DeleteRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoomResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoomResponse) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRoomResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_api_admin_server_proto_rawDescGZIP(), []int{10}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_api_admin_server_proto protoreflect.FileDescriptor

var file_api_admin_server_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x20,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72,
	0x6f, 0x6f, 0x6d, 0x22, 0x5b, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x6d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x43, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x44, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62,
	0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x43, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22,
	0x44, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x4c, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xec,
	0x03, 0x0a, 0x07, 0x41, 0x6d, 0x69, 0x6e, 0x41, 0x50, 0x49, 0x12, 0x58, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x24, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x6e,
	0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x12, 0x26, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x6e, 0x64, 0x2e,
	0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x27, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x6e,
	0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x27, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x27, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62,
	0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x64, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a,
	0x26, 0x64, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x6e, 0x64, 0x62, 0x6f, 0x74, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_admin_server_proto_rawDescOnce sync.Once
	file_api_admin_server_proto_rawDescData = file_api_admin_server_proto_rawDesc
)

func file_api_admin_server_proto_rawDescGZIP() []byte {
	file_api_admin_server_proto_rawDescOnce.Do(func() {
		file_api_admin_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_admin_server_proto_rawDescData)
	})
	return file_api_admin_server_proto_rawDescData
}

var file_api_admin_server_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_admin_server_proto_goTypes = []interface{}{
	(*GetRoomRequest)(nil),     // 0: dnd.bot.api.v1alpha1.GetRoomRequest
	(*GetRoomResponse)(nil),    // 1: dnd.bot.api.v1alpha1.GetRoomResponse
	(*ListRoomsRequest)(nil),   // 2: dnd.bot.api.v1alpha1.ListRoomsRequest
	(*ListRoomsResponse)(nil),  // 3: dnd.bot.api.v1alpha1.ListRoomsResponse
	(*CreateRoomRequest)(nil),  // 4: dnd.bot.api.v1alpha1.CreateRoomRequest
	(*CreateRoomResponse)(nil), // 5: dnd.bot.api.v1alpha1.CreateRoomResponse
	(*UpdateRoomRequest)(nil),  // 6: dnd.bot.api.v1alpha1.UpdateRoomRequest
	(*UpdateRoomResponse)(nil), // 7: dnd.bot.api.v1alpha1.UpdateRoomResponse
	(*DeleteRoomRequest)(nil),  // 8: dnd.bot.api.v1alpha1.DeleteRoomRequest
	(*DeleteRoomResponse)(nil), // 9: dnd.bot.api.v1alpha1.DeleteRoomResponse
	(*Room)(nil),               // 10: dnd.bot.api.v1alpha1.Room
}
var file_api_admin_server_proto_depIdxs = []int32{
	10, // 0: dnd.bot.api.v1alpha1.GetRoomResponse.room:type_name -> dnd.bot.api.v1alpha1.Room
	10, // 1: dnd.bot.api.v1alpha1.ListRoomsResponse.rooms:type_name -> dnd.bot.api.v1alpha1.Room
	10, // 2: dnd.bot.api.v1alpha1.CreateRoomRequest.room:type_name -> dnd.bot.api.v1alpha1.Room
	10, // 3: dnd.bot.api.v1alpha1.CreateRoomResponse.room:type_name -> dnd.bot.api.v1alpha1.Room
	10, // 4: dnd.bot.api.v1alpha1.UpdateRoomRequest.room:type_name -> dnd.bot.api.v1alpha1.Room
	10, // 5: dnd.bot.api.v1alpha1.UpdateRoomResponse.room:type_name -> dnd.bot.api.v1alpha1.Room
	0,  // 6: dnd.bot.api.v1alpha1.AminAPI.GetRoom:input_type -> dnd.bot.api.v1alpha1.GetRoomRequest
	2,  // 7: dnd.bot.api.v1alpha1.AminAPI.ListRooms:input_type -> dnd.bot.api.v1alpha1.ListRoomsRequest
	4,  // 8: dnd.bot.api.v1alpha1.AminAPI.CreateRoom:input_type -> dnd.bot.api.v1alpha1.CreateRoomRequest
	6,  // 9: dnd.bot.api.v1alpha1.AminAPI.UpdateRoom:input_type -> dnd.bot.api.v1alpha1.UpdateRoomRequest
	8,  // 10: dnd.bot.api.v1alpha1.AminAPI.DeleteRoom:input_type -> dnd.bot.api.v1alpha1.DeleteRoomRequest
	1,  // 11: dnd.bot.api.v1alpha1.AminAPI.GetRoom:output_type -> dnd.bot.api.v1alpha1.GetRoomResponse
	3,  // 12: dnd.bot.api.v1alpha1.AminAPI.ListRooms:output_type -> dnd.bot.api.v1alpha1.ListRoomsResponse
	5,  // 13: dnd.bot.api.v1alpha1.AminAPI.CreateRoom:output_type -> dnd.bot.api.v1alpha1.CreateRoomResponse
	7,  // 14: dnd.bot.api.v1alpha1.AminAPI.UpdateRoom:output_type -> dnd.bot.api.v1alpha1.UpdateRoomResponse
	9,  // 15: dnd.bot.api.v1alpha1.AminAPI.DeleteRoom:output_type -> dnd.bot.api.v1alpha1.DeleteRoomResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_admin_server_proto_init() }
func file_api_admin_server_proto_init() {
	if File_api_admin_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_admin_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomRequest); i {
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
		file_api_admin_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomResponse); i {
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
		file_api_admin_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsRequest); i {
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
		file_api_admin_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsResponse); i {
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
		file_api_admin_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
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
		file_api_admin_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomResponse); i {
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
		file_api_admin_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoomRequest); i {
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
		file_api_admin_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoomResponse); i {
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
		file_api_admin_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoomRequest); i {
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
		file_api_admin_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoomResponse); i {
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
		file_api_admin_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
			RawDescriptor: file_api_admin_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_admin_server_proto_goTypes,
		DependencyIndexes: file_api_admin_server_proto_depIdxs,
		MessageInfos:      file_api_admin_server_proto_msgTypes,
	}.Build()
	File_api_admin_server_proto = out.File
	file_api_admin_server_proto_rawDesc = nil
	file_api_admin_server_proto_goTypes = nil
	file_api_admin_server_proto_depIdxs = nil
}