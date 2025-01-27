// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: common/core.proto

// Package common defines commonly used types agnostic to the node role on the Sonr network.

package common

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

// Internet Connection Type
type Connection int32

const (
	Connection_CONNECTION_UNSPECIFIED Connection = 0
	// ConnectionWifi is used for WiFi connections.
	Connection_CONNECTION_WIFI Connection = 1
	// ConnectionEthernet is used for Ethernet connections.
	Connection_CONNECTION_ETHERNET Connection = 2
	// ConnectionMobile is used for mobile connections.
	Connection_CONNECTION_MOBILE Connection = 3
	// CONNECTION_OFFLINE
	Connection_CONNECTION_OFFLINE Connection = 4 // No Internet Connection
)

// Enum value maps for Connection.
var (
	Connection_name = map[int32]string{
		0: "CONNECTION_UNSPECIFIED",
		1: "CONNECTION_WIFI",
		2: "CONNECTION_ETHERNET",
		3: "CONNECTION_MOBILE",
		4: "CONNECTION_OFFLINE",
	}
	Connection_value = map[string]int32{
		"CONNECTION_UNSPECIFIED": 0,
		"CONNECTION_WIFI":        1,
		"CONNECTION_ETHERNET":    2,
		"CONNECTION_MOBILE":      3,
		"CONNECTION_OFFLINE":     4,
	}
)

func (x Connection) Enum() *Connection {
	p := new(Connection)
	*p = x
	return p
}

func (x Connection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Connection) Descriptor() protoreflect.EnumDescriptor {
	return file_common_core_proto_enumTypes[0].Descriptor()
}

func (Connection) Type() protoreflect.EnumType {
	return &file_common_core_proto_enumTypes[0]
}

func (x Connection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Connection.Descriptor instead.
func (Connection) EnumDescriptor() ([]byte, []int) {
	return file_common_core_proto_rawDescGZIP(), []int{0}
}

// File Content Type
type MIME_Type int32

const (
	// Other File Type - If cannot derive from Subtype
	MIME_TYPE_UNSPECIFIED MIME_Type = 0
	// Sound, Audio Files
	MIME_TYPE_AUDIO MIME_Type = 1
	// Document Files - PDF, Word, Excel, etc.
	MIME_TYPE_DOCUMENT MIME_Type = 2
	// Image Files
	MIME_TYPE_IMAGE MIME_Type = 3
	// Text Based Files
	MIME_TYPE_TEXT MIME_Type = 4
	// Video Files
	MIME_TYPE_VIDEO MIME_Type = 5
	// URL Links
	MIME_TYPE_URL MIME_Type = 6
	// Crypto Files
	MIME_TYPE_CRYPTO MIME_Type = 7
)

// Enum value maps for MIME_Type.
var (
	MIME_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_AUDIO",
		2: "TYPE_DOCUMENT",
		3: "TYPE_IMAGE",
		4: "TYPE_TEXT",
		5: "TYPE_VIDEO",
		6: "TYPE_URL",
		7: "TYPE_CRYPTO",
	}
	MIME_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_AUDIO":       1,
		"TYPE_DOCUMENT":    2,
		"TYPE_IMAGE":       3,
		"TYPE_TEXT":        4,
		"TYPE_VIDEO":       5,
		"TYPE_URL":         6,
		"TYPE_CRYPTO":      7,
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
	return file_common_core_proto_enumTypes[1].Descriptor()
}

func (MIME_Type) Type() protoreflect.EnumType {
	return &file_common_core_proto_enumTypes[1]
}

func (x MIME_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MIME_Type.Descriptor instead.
func (MIME_Type) EnumDescriptor() ([]byte, []int) {
	return file_common_core_proto_rawDescGZIP(), []int{2, 0}
}

// Peers Active Status
type Peer_Status int32

const (
	Peer_STATUS_UNSPECIFIED Peer_Status = 0 // Offline - Not Online or Not a Full Node
	Peer_STATUS_ONLINE      Peer_Status = 1 // Online - Full Node Available
	Peer_STATUS_AWAY        Peer_Status = 2 // Away - Not Online, but has a full node
	Peer_STATUS_BUSY        Peer_Status = 3 // Busy - Online, but busy with Transfer
)

// Enum value maps for Peer_Status.
var (
	Peer_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_ONLINE",
		2: "STATUS_AWAY",
		3: "STATUS_BUSY",
	}
	Peer_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_ONLINE":      1,
		"STATUS_AWAY":        2,
		"STATUS_BUSY":        3,
	}
)

func (x Peer_Status) Enum() *Peer_Status {
	p := new(Peer_Status)
	*p = x
	return p
}

func (x Peer_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Peer_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_common_core_proto_enumTypes[2].Descriptor()
}

func (Peer_Status) Type() protoreflect.EnumType {
	return &file_common_core_proto_enumTypes[2]
}

func (x Peer_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Peer_Status.Descriptor instead.
func (Peer_Status) EnumDescriptor() ([]byte, []int) {
	return file_common_core_proto_rawDescGZIP(), []int{3, 0}
}

// Location from GeoIP and OLC information
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Location Latitude
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Location Longitude
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Location Placemark Information - Generated
	Placemark *Location_Placemark `protobuf:"bytes,3,opt,name=placemark,proto3" json:"placemark,omitempty"`
	// Last Updated Time
	LastModified int64 `protobuf:"varint,4,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_common_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_common_core_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Location) GetPlacemark() *Location_Placemark {
	if x != nil {
		return x.Placemark
	}
	return nil
}

func (x *Location) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

// Shared Metadata for Messages on all Protocols
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Node ID
	NodeId string `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Signature of the message
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// Public Key of the message sender
	PublicKey []byte `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_common_core_proto_msgTypes[1]
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
	return file_common_core_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Metadata) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Metadata) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Metadata) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Standard MIME with Additional Extensions
type MIME struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of File
	Type MIME_Type `protobuf:"varint,1,opt,name=type,proto3,enum=common.MIME_Type" json:"type,omitempty"`
	// Extension of File
	Subtype string `protobuf:"bytes,2,opt,name=subtype,proto3" json:"subtype,omitempty"`
	// Type/Subtype i.e. (image/jpeg)
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MIME) Reset() {
	*x = MIME{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MIME) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MIME) ProtoMessage() {}

func (x *MIME) ProtoReflect() protoreflect.Message {
	mi := &file_common_core_proto_msgTypes[2]
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
	return file_common_core_proto_rawDescGZIP(), []int{2}
}

func (x *MIME) GetType() MIME_Type {
	if x != nil {
		return x.Type
	}
	return MIME_TYPE_UNSPECIFIED
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

// Basic Info Sent to Peers to Establish Connections
type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName        string       `protobuf:"bytes,1,opt,name=s_name,json=sName,proto3" json:"s_name,omitempty"`                       // User Sonr Domain
	Status       Peer_Status  `protobuf:"varint,2,opt,name=status,proto3,enum=common.Peer_Status" json:"status,omitempty"`         // Peer Status
	Device       *Peer_Device `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`                                  // Peer Device Info
	Profile      *Profile     `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`                                // Peers General Information
	PublicKey    []byte       `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`           // Public Key of the Peer
	PeerId       string       `protobuf:"bytes,6,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`                    // Peer ID
	LastModified int64        `protobuf:"varint,7,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"` // Last Modified Timestamp
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_common_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_common_core_proto_rawDescGZIP(), []int{3}
}

func (x *Peer) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *Peer) GetStatus() Peer_Status {
	if x != nil {
		return x.Status
	}
	return Peer_STATUS_UNSPECIFIED
}

func (x *Peer) GetDevice() *Peer_Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *Peer) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *Peer) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Peer) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *Peer) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

// General Information about Peer passed during Authentication
type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName        string `protobuf:"bytes,1,opt,name=s_name,json=sName,proto3" json:"s_name,omitempty"`                       // Sonr Based Username
	FirstName    string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`           // General Info
	LastName     string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`              // General Info
	Picture      []byte `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`                                // Profile Picture
	Bio          string `protobuf:"bytes,6,opt,name=bio,proto3" json:"bio,omitempty"`                                        // User Biography
	LastModified int64  `protobuf:"varint,7,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"` // Last Modified Timestamp
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_common_core_proto_msgTypes[4]
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
	return file_common_core_proto_rawDescGZIP(), []int{4}
}

func (x *Profile) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *Profile) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Profile) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Profile) GetPicture() []byte {
	if x != nil {
		return x.Picture
	}
	return nil
}

func (x *Profile) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *Profile) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

// Contains detailed placemark information.
type Location_Placemark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name associated with the placemark.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The street associated with the placemark.
	Street string `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
	// The abbreviated country name, according to the two letter (alpha-2) [ISO
	// standard](https://www.iso.org/iso-3166-country-codes.html).
	IsoCountryCode string `protobuf:"bytes,3,opt,name=iso_country_code,json=isoCountryCode,proto3" json:"iso_country_code,omitempty"`
	// The name of the country associated with the placemark.
	Country string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	// The postal code associated with the placemark.
	PostalCode string `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	// The name of the state or province associated with the placemark.
	AdministrativeArea string `protobuf:"bytes,6,opt,name=administrative_area,json=administrativeArea,proto3" json:"administrative_area,omitempty"`
	// Additional administrative area information for the placemark.
	SubAdministrativeArea string `protobuf:"bytes,7,opt,name=sub_administrative_area,json=subAdministrativeArea,proto3" json:"sub_administrative_area,omitempty"`
	// The name of the city associated with the placemark.
	Locality string `protobuf:"bytes,8,opt,name=locality,proto3" json:"locality,omitempty"`
	// Additional city-level information for the placemark.
	SubLocality string `protobuf:"bytes,9,opt,name=sub_locality,json=subLocality,proto3" json:"sub_locality,omitempty"`
	// The street address associated with the placemark.
	Thoroughfare string `protobuf:"bytes,10,opt,name=thoroughfare,proto3" json:"thoroughfare,omitempty"`
	// Additional street address information for the placemark.
	SubThoroughfare string `protobuf:"bytes,11,opt,name=sub_thoroughfare,json=subThoroughfare,proto3" json:"sub_thoroughfare,omitempty"`
}

func (x *Location_Placemark) Reset() {
	*x = Location_Placemark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location_Placemark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location_Placemark) ProtoMessage() {}

func (x *Location_Placemark) ProtoReflect() protoreflect.Message {
	mi := &file_common_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location_Placemark.ProtoReflect.Descriptor instead.
func (*Location_Placemark) Descriptor() ([]byte, []int) {
	return file_common_core_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Location_Placemark) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location_Placemark) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Location_Placemark) GetIsoCountryCode() string {
	if x != nil {
		return x.IsoCountryCode
	}
	return ""
}

func (x *Location_Placemark) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Location_Placemark) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Location_Placemark) GetAdministrativeArea() string {
	if x != nil {
		return x.AdministrativeArea
	}
	return ""
}

func (x *Location_Placemark) GetSubAdministrativeArea() string {
	if x != nil {
		return x.SubAdministrativeArea
	}
	return ""
}

func (x *Location_Placemark) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *Location_Placemark) GetSubLocality() string {
	if x != nil {
		return x.SubLocality
	}
	return ""
}

func (x *Location_Placemark) GetThoroughfare() string {
	if x != nil {
		return x.Thoroughfare
	}
	return ""
}

func (x *Location_Placemark) GetSubThoroughfare() string {
	if x != nil {
		return x.SubThoroughfare
	}
	return ""
}

// Peer Info for Device
type Peer_Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                             // Peer Device ID
	HostName string `protobuf:"bytes,2,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"` // Peer Host Name
	Os       string `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`                             // Peer Operating System
	Arch     string `protobuf:"bytes,4,opt,name=arch,proto3" json:"arch,omitempty"`                         // Peer Architecture
	Model    string `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`                       // Peers Device Model
}

func (x *Peer_Device) Reset() {
	*x = Peer_Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer_Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer_Device) ProtoMessage() {}

func (x *Peer_Device) ProtoReflect() protoreflect.Message {
	mi := &file_common_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer_Device.ProtoReflect.Descriptor instead.
func (*Peer_Device) Descriptor() ([]byte, []int) {
	return file_common_core_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Peer_Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Peer_Device) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *Peer_Device) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *Peer_Device) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *Peer_Device) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

var File_common_core_proto protoreflect.FileDescriptor

var file_common_core_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xb9, 0x04, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x1a, 0x93, 0x03, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x69,
	0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x73, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x2f, 0x0a, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x41, 0x72, 0x65,
	0x61, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x75, 0x62, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x73, 0x75, 0x62, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x6f, 0x72,
	0x6f, 0x75, 0x67, 0x68, 0x66, 0x61, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x68, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x66, 0x61, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x73, 0x75, 0x62, 0x5f, 0x74, 0x68, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x66, 0x61, 0x72, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x54, 0x68, 0x6f, 0x72, 0x6f,
	0x75, 0x67, 0x68, 0x66, 0x61, 0x72, 0x65, 0x22, 0x7e, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0xed, 0x01, 0x0a, 0x04, 0x4d, 0x49, 0x4d, 0x45,
	0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x49, 0x4d, 0x45, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x55, 0x44, 0x49, 0x4f, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x52, 0x4c, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x52, 0x59, 0x50, 0x54, 0x4f, 0x10, 0x07, 0x22, 0xc7, 0x03, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x1a, 0x6f, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x55, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x57, 0x41, 0x59, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10,
	0x03, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x2a, 0x85, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x49, 0x46, 0x49, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x45, 0x54, 0x48, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10,
	0x03, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x04, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_core_proto_rawDescOnce sync.Once
	file_common_core_proto_rawDescData = file_common_core_proto_rawDesc
)

func file_common_core_proto_rawDescGZIP() []byte {
	file_common_core_proto_rawDescOnce.Do(func() {
		file_common_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_core_proto_rawDescData)
	})
	return file_common_core_proto_rawDescData
}

var file_common_core_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_core_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_core_proto_goTypes = []interface{}{
	(Connection)(0),            // 0: common.Connection
	(MIME_Type)(0),             // 1: common.MIME.Type
	(Peer_Status)(0),           // 2: common.Peer.Status
	(*Location)(nil),           // 3: common.Location
	(*Metadata)(nil),           // 4: common.Metadata
	(*MIME)(nil),               // 5: common.MIME
	(*Peer)(nil),               // 6: common.Peer
	(*Profile)(nil),            // 7: common.Profile
	(*Location_Placemark)(nil), // 8: common.Location.Placemark
	(*Peer_Device)(nil),        // 9: common.Peer.Device
}
var file_common_core_proto_depIdxs = []int32{
	8, // 0: common.Location.placemark:type_name -> common.Location.Placemark
	1, // 1: common.MIME.type:type_name -> common.MIME.Type
	2, // 2: common.Peer.status:type_name -> common.Peer.Status
	9, // 3: common.Peer.device:type_name -> common.Peer.Device
	7, // 4: common.Peer.profile:type_name -> common.Profile
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_core_proto_init() }
func file_common_core_proto_init() {
	if File_common_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_common_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
		file_common_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location_Placemark); i {
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
		file_common_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer_Device); i {
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
			RawDescriptor: file_common_core_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_core_proto_goTypes,
		DependencyIndexes: file_common_core_proto_depIdxs,
		EnumInfos:         file_common_core_proto_enumTypes,
		MessageInfos:      file_common_core_proto_msgTypes,
	}.Build()
	File_common_core_proto = out.File
	file_common_core_proto_rawDesc = nil
	file_common_core_proto_goTypes = nil
	file_common_core_proto_depIdxs = nil
}
