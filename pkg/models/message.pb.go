// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: message.proto

package proto

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

// Define Subject
type AuthMessage_Subject int32

const (
	AuthMessage_NONE    AuthMessage_Subject = 0
	AuthMessage_REQUEST AuthMessage_Subject = 1
	AuthMessage_ACCEPT  AuthMessage_Subject = 2
	AuthMessage_DECLINE AuthMessage_Subject = 3
)

// Enum value maps for AuthMessage_Subject.
var (
	AuthMessage_Subject_name = map[int32]string{
		0: "NONE",
		1: "REQUEST",
		2: "ACCEPT",
		3: "DECLINE",
	}
	AuthMessage_Subject_value = map[string]int32{
		"NONE":    0,
		"REQUEST": 1,
		"ACCEPT":  2,
		"DECLINE": 3,
	}
)

func (x AuthMessage_Subject) Enum() *AuthMessage_Subject {
	p := new(AuthMessage_Subject)
	*p = x
	return p
}

func (x AuthMessage_Subject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthMessage_Subject) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (AuthMessage_Subject) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x AuthMessage_Subject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthMessage_Subject.Descriptor instead.
func (AuthMessage_Subject) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4, 0}
}

// [PLUGIN]
// Initial Request Message
type RequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Olc     string   `protobuf:"bytes,1,opt,name=olc,proto3" json:"olc,omitempty"`
	Device  string   `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	Contact *Contact `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *RequestMessage) Reset() {
	*x = RequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessage) ProtoMessage() {}

func (x *RequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessage.ProtoReflect.Descriptor instead.
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMessage) GetOlc() string {
	if x != nil {
		return x.Olc
	}
	return ""
}

func (x *RequestMessage) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *RequestMessage) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

// Message Sent when User is connected and calls GetUser()
type ConnectedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId  string   `protobuf:"bytes,1,opt,name=hostId,proto3" json:"hostId,omitempty"`
	Profile *Profile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	Contact *Contact `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *ConnectedMessage) Reset() {
	*x = ConnectedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedMessage) ProtoMessage() {}

func (x *ConnectedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedMessage.ProtoReflect.Descriptor instead.
func (*ConnectedMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectedMessage) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *ConnectedMessage) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *ConnectedMessage) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

// [CORE]
// Message Sent when peer has direction update to Lobby Topic
type LobbyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event  string    `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Sender string    `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Data   *PeerInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LobbyMessage) Reset() {
	*x = LobbyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyMessage) ProtoMessage() {}

func (x *LobbyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyMessage.ProtoReflect.Descriptor instead.
func (*LobbyMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *LobbyMessage) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *LobbyMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *LobbyMessage) GetData() *PeerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

// Message Sent when peers returned from lobby
type RefreshMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count          int32       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Olc            string      `protobuf:"bytes,2,opt,name=olc,proto3" json:"olc,omitempty"`
	AvailablePeers []*PeerInfo `protobuf:"bytes,3,rep,name=availablePeers,proto3" json:"availablePeers,omitempty"`
}

func (x *RefreshMessage) Reset() {
	*x = RefreshMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshMessage) ProtoMessage() {}

func (x *RefreshMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshMessage.ProtoReflect.Descriptor instead.
func (*RefreshMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshMessage) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RefreshMessage) GetOlc() string {
	if x != nil {
		return x.Olc
	}
	return ""
}

func (x *RefreshMessage) GetAvailablePeers() []*PeerInfo {
	if x != nil {
		return x.AvailablePeers
	}
	return nil
}

// [CORE]
// Authorization message sent in stream
type AuthMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define Root Message
	Peer     *PeerInfo           `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Subject  AuthMessage_Subject `protobuf:"varint,2,opt,name=subject,proto3,enum=AuthMessage_Subject" json:"subject,omitempty"`
	Metadata *Metadata           `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *AuthMessage) Reset() {
	*x = AuthMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMessage) ProtoMessage() {}

func (x *AuthMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMessage.ProtoReflect.Descriptor instead.
func (*AuthMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *AuthMessage) GetPeer() *PeerInfo {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *AuthMessage) GetSubject() AuthMessage_Subject {
	if x != nil {
		return x.Subject
	}
	return AuthMessage_NONE
}

func (x *AuthMessage) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x6c, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x72, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x5b, 0x0a, 0x0c, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x6c, 0x63, 0x12, 0x31, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x65, 0x65, 0x72, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x70, 0x65, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x07, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x43,
	0x4c, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(AuthMessage_Subject)(0), // 0: AuthMessage.Subject
	(*RequestMessage)(nil),   // 1: RequestMessage
	(*ConnectedMessage)(nil), // 2: ConnectedMessage
	(*LobbyMessage)(nil),     // 3: LobbyMessage
	(*RefreshMessage)(nil),   // 4: RefreshMessage
	(*AuthMessage)(nil),      // 5: AuthMessage
	(*Contact)(nil),          // 6: Contact
	(*Profile)(nil),          // 7: Profile
	(*PeerInfo)(nil),         // 8: PeerInfo
	(*Metadata)(nil),         // 9: Metadata
}
var file_message_proto_depIdxs = []int32{
	6, // 0: RequestMessage.contact:type_name -> Contact
	7, // 1: ConnectedMessage.profile:type_name -> Profile
	6, // 2: ConnectedMessage.contact:type_name -> Contact
	8, // 3: LobbyMessage.data:type_name -> PeerInfo
	8, // 4: RefreshMessage.availablePeers:type_name -> PeerInfo
	8, // 5: AuthMessage.peer:type_name -> PeerInfo
	0, // 6: AuthMessage.subject:type_name -> AuthMessage.Subject
	9, // 7: AuthMessage.metadata:type_name -> Metadata
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	file_data_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessage); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectedMessage); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyMessage); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshMessage); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMessage); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
