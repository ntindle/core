// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: protocols/transmit/v1/transmit.proto

package transmit

import (
	common "github.com/sonr-io/core/types/go/common"
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

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction    common.Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=common.Direction" json:"direction,omitempty"`
	From         *common.Peer     `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To           *common.Peer     `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Payload      *common.Payload  `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	LastUpdated  int64            `protobuf:"varint,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Items        []*SessionItem   `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	CurrentIndex int32            `protobuf:"varint,7,opt,name=current_index,json=currentIndex,proto3" json:"current_index,omitempty"`
	Results      map[int32]bool   `protobuf:"bytes,8,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_transmit_v1_transmit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_transmit_v1_transmit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_protocols_transmit_v1_transmit_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetDirection() common.Direction {
	if x != nil {
		return x.Direction
	}
	return common.Direction(0)
}

func (x *Session) GetFrom() *common.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Session) GetTo() *common.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Session) GetPayload() *common.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Session) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

func (x *Session) GetItems() []*SessionItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Session) GetCurrentIndex() int32 {
	if x != nil {
		return x.CurrentIndex
	}
	return 0
}

func (x *Session) GetResults() map[int32]bool {
	if x != nil {
		return x.Results
	}
	return nil
}

type SessionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     int32            `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Count     int32            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Item      *common.FileItem `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	Written   int64            `protobuf:"varint,4,opt,name=written,proto3" json:"written,omitempty"`
	Size      int64            `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	TotalSize int64            `protobuf:"varint,6,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	Direction common.Direction `protobuf:"varint,7,opt,name=direction,proto3,enum=common.Direction" json:"direction,omitempty"`
	Path      string           `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *SessionItem) Reset() {
	*x = SessionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_transmit_v1_transmit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionItem) ProtoMessage() {}

func (x *SessionItem) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_transmit_v1_transmit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionItem.ProtoReflect.Descriptor instead.
func (*SessionItem) Descriptor() ([]byte, []int) {
	return file_protocols_transmit_v1_transmit_proto_rawDescGZIP(), []int{1}
}

func (x *SessionItem) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SessionItem) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SessionItem) GetItem() *common.FileItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *SessionItem) GetWritten() int64 {
	if x != nil {
		return x.Written
	}
	return 0
}

func (x *SessionItem) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SessionItem) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *SessionItem) GetDirection() common.Direction {
	if x != nil {
		return x.Direction
	}
	return common.Direction(0)
}

func (x *SessionItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type SessionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload   *common.Payload  `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Direction common.Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=common.Direction" json:"direction,omitempty"`
}

func (x *SessionPayload) Reset() {
	*x = SessionPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_transmit_v1_transmit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionPayload) ProtoMessage() {}

func (x *SessionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_transmit_v1_transmit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionPayload.ProtoReflect.Descriptor instead.
func (*SessionPayload) Descriptor() ([]byte, []int) {
	return file_protocols_transmit_v1_transmit_proto_rawDescGZIP(), []int{2}
}

func (x *SessionPayload) GetPayload() *common.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SessionPayload) GetDirection() common.Direction {
	if x != nil {
		return x.Direction
	}
	return common.Direction(0)
}

var File_protocols_transmit_v1_transmit_proto protoreflect.FileDescriptor

var file_protocols_transmit_v1_transmit_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x29, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x38,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x45, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xf1, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x2f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x6c, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x2f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocols_transmit_v1_transmit_proto_rawDescOnce sync.Once
	file_protocols_transmit_v1_transmit_proto_rawDescData = file_protocols_transmit_v1_transmit_proto_rawDesc
)

func file_protocols_transmit_v1_transmit_proto_rawDescGZIP() []byte {
	file_protocols_transmit_v1_transmit_proto_rawDescOnce.Do(func() {
		file_protocols_transmit_v1_transmit_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocols_transmit_v1_transmit_proto_rawDescData)
	})
	return file_protocols_transmit_v1_transmit_proto_rawDescData
}

var file_protocols_transmit_v1_transmit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocols_transmit_v1_transmit_proto_goTypes = []interface{}{
	(*Session)(nil),         // 0: protocols.transmit.v1.Session
	(*SessionItem)(nil),     // 1: protocols.transmit.v1.SessionItem
	(*SessionPayload)(nil),  // 2: protocols.transmit.v1.SessionPayload
	nil,                     // 3: protocols.transmit.v1.Session.ResultsEntry
	(common.Direction)(0),   // 4: common.Direction
	(*common.Peer)(nil),     // 5: common.Peer
	(*common.Payload)(nil),  // 6: common.Payload
	(*common.FileItem)(nil), // 7: common.FileItem
}
var file_protocols_transmit_v1_transmit_proto_depIdxs = []int32{
	4,  // 0: protocols.transmit.v1.Session.direction:type_name -> common.Direction
	5,  // 1: protocols.transmit.v1.Session.from:type_name -> common.Peer
	5,  // 2: protocols.transmit.v1.Session.to:type_name -> common.Peer
	6,  // 3: protocols.transmit.v1.Session.payload:type_name -> common.Payload
	1,  // 4: protocols.transmit.v1.Session.items:type_name -> protocols.transmit.v1.SessionItem
	3,  // 5: protocols.transmit.v1.Session.results:type_name -> protocols.transmit.v1.Session.ResultsEntry
	7,  // 6: protocols.transmit.v1.SessionItem.item:type_name -> common.FileItem
	4,  // 7: protocols.transmit.v1.SessionItem.direction:type_name -> common.Direction
	6,  // 8: protocols.transmit.v1.SessionPayload.payload:type_name -> common.Payload
	4,  // 9: protocols.transmit.v1.SessionPayload.direction:type_name -> common.Direction
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_protocols_transmit_v1_transmit_proto_init() }
func file_protocols_transmit_v1_transmit_proto_init() {
	if File_protocols_transmit_v1_transmit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocols_transmit_v1_transmit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_protocols_transmit_v1_transmit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionItem); i {
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
		file_protocols_transmit_v1_transmit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionPayload); i {
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
			RawDescriptor: file_protocols_transmit_v1_transmit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocols_transmit_v1_transmit_proto_goTypes,
		DependencyIndexes: file_protocols_transmit_v1_transmit_proto_depIdxs,
		MessageInfos:      file_protocols_transmit_v1_transmit_proto_msgTypes,
	}.Build()
	File_protocols_transmit_v1_transmit_proto = out.File
	file_protocols_transmit_v1_transmit_proto_rawDesc = nil
	file_protocols_transmit_v1_transmit_proto_goTypes = nil
	file_protocols_transmit_v1_transmit_proto_depIdxs = nil
}
