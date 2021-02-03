// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: data.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Define Payload: Define Payload Enum
type Payload int32

const (
	Payload_NONE    Payload = 0
	Payload_IMAGE   Payload = 1
	Payload_VIDEO   Payload = 2
	Payload_FILE    Payload = 3
	Payload_CONTACT Payload = 4
	Payload_URL     Payload = 5
	Payload_TEXT    Payload = 6
)

// Enum value maps for Payload.
var (
	Payload_name = map[int32]string{
		0: "NONE",
		1: "IMAGE",
		2: "VIDEO",
		3: "FILE",
		4: "CONTACT",
		5: "URL",
		6: "TEXT",
	}
	Payload_value = map[string]int32{
		"NONE":    0,
		"IMAGE":   1,
		"VIDEO":   2,
		"FILE":    3,
		"CONTACT": 4,
		"URL":     5,
		"TEXT":    6,
	}
)

func (x Payload) Enum() *Payload {
	p := new(Payload)
	*p = x
	return p
}

func (x Payload) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Payload) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[0].Descriptor()
}

func (Payload) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[0]
}

func (x Payload) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Payload.Descriptor instead.
func (Payload) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

type MIME_Type int32

const (
	MIME_application MIME_Type = 0
	MIME_audio       MIME_Type = 1
	MIME_image       MIME_Type = 2
	MIME_text        MIME_Type = 3
	MIME_video       MIME_Type = 4
)

// Enum value maps for MIME_Type.
var (
	MIME_Type_name = map[int32]string{
		0: "application",
		1: "audio",
		2: "image",
		3: "text",
		4: "video",
	}
	MIME_Type_value = map[string]int32{
		"application": 0,
		"audio":       1,
		"image":       2,
		"text":        3,
		"video":       4,
	}
)

func (x MIME_Type) Enum() *MIME_Type {
	p := new(MIME_Type)
	*p = x
	return p
}

func (x MIME_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MIME_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[1].Descriptor()
}

func (MIME_Type) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[1]
}

func (x MIME_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MIME_Type.Descriptor instead.
func (MIME_Type) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2, 0}
}

type LobbyEvent_Event int32

const (
	LobbyEvent_ENTER    LobbyEvent_Event = 0
	LobbyEvent_EXIT     LobbyEvent_Event = 1
	LobbyEvent_STANDBY  LobbyEvent_Event = 2
	LobbyEvent_EXCHANGE LobbyEvent_Event = 3
	LobbyEvent_UPDATE   LobbyEvent_Event = 4
)

// Enum value maps for LobbyEvent_Event.
var (
	LobbyEvent_Event_name = map[int32]string{
		0: "ENTER",
		1: "EXIT",
		2: "STANDBY",
		3: "EXCHANGE",
		4: "UPDATE",
	}
	LobbyEvent_Event_value = map[string]int32{
		"ENTER":    0,
		"EXIT":     1,
		"STANDBY":  2,
		"EXCHANGE": 3,
		"UPDATE":   4,
	}
)

func (x LobbyEvent_Event) Enum() *LobbyEvent_Event {
	p := new(LobbyEvent_Event)
	*p = x
	return p
}

func (x LobbyEvent_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LobbyEvent_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[2].Descriptor()
}

func (LobbyEvent_Event) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[2]
}

func (x LobbyEvent_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LobbyEvent_Event.Descriptor instead.
func (LobbyEvent_Event) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4, 0}
}

// Define Metadata Type: For Received Transfer
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path      string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Size      int32    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Mime      *MIME    `protobuf:"bytes,5,opt,name=mime,proto3" json:"mime,omitempty"`
	Thumbnail []byte   `protobuf:"bytes,6,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Received  int32    `protobuf:"varint,7,opt,name=received,proto3" json:"received,omitempty"`
	Owner     *Profile `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Metadata) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Metadata) GetMime() *MIME {
	if x != nil {
		return x.Mime
	}
	return nil
}

func (x *Metadata) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *Metadata) GetReceived() int32 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *Metadata) GetOwner() *Profile {
	if x != nil {
		return x.Owner
	}
	return nil
}

// Define Preview: For File Transfer Request
type Preview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime      *MIME  `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size      int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Thumbnail []byte `protobuf:"bytes,4,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Duration  int32  `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Preview) Reset() {
	*x = Preview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Preview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preview) ProtoMessage() {}

func (x *Preview) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preview.ProtoReflect.Descriptor instead.
func (*Preview) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

func (x *Preview) GetMime() *MIME {
	if x != nil {
		return x.Mime
	}
	return nil
}

func (x *Preview) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Preview) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Preview) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *Preview) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

// Define MIME: Protobuf Version of Mime
type MIME struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    MIME_Type `protobuf:"varint,1,opt,name=type,proto3,enum=MIME_Type" json:"type,omitempty"`
	Subtype string    `protobuf:"bytes,2,opt,name=subtype,proto3" json:"subtype,omitempty"`
	Value   string    `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MIME) Reset() {
	*x = MIME{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MIME) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MIME) ProtoMessage() {}

func (x *MIME) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MIME.ProtoReflect.Descriptor instead.
func (*MIME) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

func (x *MIME) GetType() MIME_Type {
	if x != nil {
		return x.Type
	}
	return MIME_application
}

func (x *MIME) GetSubtype() string {
	if x != nil {
		return x.Subtype
	}
	return ""
}

func (x *MIME) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Define Lobby Type: For Info about Lobby
type Lobby struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Olc   string           `protobuf:"bytes,1,opt,name=olc,proto3" json:"olc,omitempty"`
	Size  int32            `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Count int32            `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Peers map[string]*Peer `protobuf:"bytes,4,rep,name=peers,proto3" json:"peers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Lobby) Reset() {
	*x = Lobby{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lobby) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lobby) ProtoMessage() {}

func (x *Lobby) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lobby.ProtoReflect.Descriptor instead.
func (*Lobby) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

func (x *Lobby) GetOlc() string {
	if x != nil {
		return x.Olc
	}
	return ""
}

func (x *Lobby) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Lobby) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Lobby) GetPeers() map[string]*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

// [CORE]
// Message Sent when peer messages Lobby Topic
type LobbyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event LobbyEvent_Event `protobuf:"varint,2,opt,name=event,proto3,enum=LobbyEvent_Event" json:"event,omitempty"`
	Data  *Peer            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LobbyEvent) Reset() {
	*x = LobbyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyEvent) ProtoMessage() {}

func (x *LobbyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyEvent.ProtoReflect.Descriptor instead.
func (*LobbyEvent) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4}
}

func (x *LobbyEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LobbyEvent) GetEvent() LobbyEvent_Event {
	if x != nil {
		return x.Event
	}
	return LobbyEvent_ENTER
}

func (x *LobbyEvent) GetData() *Peer {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_data_proto protoreflect.FileDescriptor

var file_data_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x4d, 0x49, 0x4d, 0x45, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x19, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x4d, 0x49, 0x4d, 0x45, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x9a, 0x01, 0x0a, 0x04, 0x4d, 0x49, 0x4d, 0x45, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4d, 0x49, 0x4d, 0x45, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x04, 0x22, 0xad, 0x01, 0x0a,
	0x05, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x6c, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x1a, 0x3f, 0x0a, 0x0a, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa5, 0x01, 0x0a,
	0x0a, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x43, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e, 0x54, 0x45,
	0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x58, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x42, 0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x04, 0x2a, 0x53, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e,
	0x54, 0x41, 0x43, 0x54, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x10, 0x05, 0x12,
	0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x06, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_proto_rawDescOnce sync.Once
	file_data_proto_rawDescData = file_data_proto_rawDesc
)

func file_data_proto_rawDescGZIP() []byte {
	file_data_proto_rawDescOnce.Do(func() {
		file_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_proto_rawDescData)
	})
	return file_data_proto_rawDescData
}

var file_data_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_data_proto_goTypes = []interface{}{
	(Payload)(0),          // 0: Payload
	(MIME_Type)(0),        // 1: MIME.Type
	(LobbyEvent_Event)(0), // 2: LobbyEvent.Event
	(*Metadata)(nil),      // 3: Metadata
	(*Preview)(nil),       // 4: Preview
	(*MIME)(nil),          // 5: MIME
	(*Lobby)(nil),         // 6: Lobby
	(*LobbyEvent)(nil),    // 7: LobbyEvent
	nil,                   // 8: Lobby.PeersEntry
	(*Profile)(nil),       // 9: Profile
	(*Peer)(nil),          // 10: Peer
}
var file_data_proto_depIdxs = []int32{
	5,  // 0: Metadata.mime:type_name -> MIME
	9,  // 1: Metadata.owner:type_name -> Profile
	5,  // 2: Preview.mime:type_name -> MIME
	1,  // 3: MIME.type:type_name -> MIME.Type
	8,  // 4: Lobby.peers:type_name -> Lobby.PeersEntry
	2,  // 5: LobbyEvent.event:type_name -> LobbyEvent.Event
	10, // 6: LobbyEvent.data:type_name -> Peer
	10, // 7: Lobby.PeersEntry.value:type_name -> Peer
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_data_proto_init() }
func file_data_proto_init() {
	if File_data_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Preview); i {
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
		file_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MIME); i {
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
		file_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lobby); i {
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
		file_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyEvent); i {
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
			RawDescriptor: file_data_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_proto_goTypes,
		DependencyIndexes: file_data_proto_depIdxs,
		EnumInfos:         file_data_proto_enumTypes,
		MessageInfos:      file_data_proto_msgTypes,
	}.Build()
	File_data_proto = out.File
	file_data_proto_rawDesc = nil
	file_data_proto_goTypes = nil
	file_data_proto_depIdxs = nil
}
