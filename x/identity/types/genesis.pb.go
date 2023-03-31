// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/identity/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the identity module's genesis state.
type GenesisState struct {
	Params               Params                     `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PrimaryIdentities    []DidDocument              `protobuf:"bytes,2,rep,name=primaryIdentities,proto3" json:"primaryIdentities"`
	BlockchainIdentities []DidDocument              `protobuf:"bytes,3,rep,name=blockchainIdentities,proto3" json:"blockchainIdentities"`
	ServiceList          []Service                  `protobuf:"bytes,4,rep,name=serviceList,proto3" json:"serviceList"`
	Relationships        []VerificationRelationship `protobuf:"bytes,5,rep,name=relationships,proto3" json:"relationships"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ee3e6e2aad889c, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPrimaryIdentities() []DidDocument {
	if m != nil {
		return m.PrimaryIdentities
	}
	return nil
}

func (m *GenesisState) GetBlockchainIdentities() []DidDocument {
	if m != nil {
		return m.BlockchainIdentities
	}
	return nil
}

func (m *GenesisState) GetServiceList() []Service {
	if m != nil {
		return m.ServiceList
	}
	return nil
}

func (m *GenesisState) GetRelationships() []VerificationRelationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

// Params defines the parameters for the module.
type Params struct {
	DidBaseContext                  string `protobuf:"bytes,1,opt,name=did_base_context,json=didBaseContext,proto3" json:"did_base_context,omitempty"`
	DidMethodContext                string `protobuf:"bytes,2,opt,name=did_method_context,json=didMethodContext,proto3" json:"did_method_context,omitempty"`
	DidMethodName                   string `protobuf:"bytes,3,opt,name=did_method_name,json=didMethodName,proto3" json:"did_method_name,omitempty"`
	DidMethodVersion                string `protobuf:"bytes,4,opt,name=did_method_version,json=didMethodVersion,proto3" json:"did_method_version,omitempty"`
	DidNetwork                      string `protobuf:"bytes,5,opt,name=did_network,json=didNetwork,proto3" json:"did_network,omitempty"`
	IpfsGateway                     string `protobuf:"bytes,6,opt,name=ipfs_gateway,json=ipfsGateway,proto3" json:"ipfs_gateway,omitempty"`
	IpfsApi                         string `protobuf:"bytes,7,opt,name=ipfs_api,json=ipfsApi,proto3" json:"ipfs_api,omitempty"`
	WebauthnAttestionPreference     string `protobuf:"bytes,8,opt,name=webauthn_attestion_preference,json=webauthnAttestionPreference,proto3" json:"webauthn_attestion_preference,omitempty"`
	WebauthnAuthenticatorAttachment string `protobuf:"bytes,9,opt,name=webauthn_authenticator_attachment,json=webauthnAuthenticatorAttachment,proto3" json:"webauthn_authenticator_attachment,omitempty"`
	WebauthnTimeout                 int32  `protobuf:"varint,10,opt,name=webauthn_timeout,json=webauthnTimeout,proto3" json:"webauthn_timeout,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ee3e6e2aad889c, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDidBaseContext() string {
	if m != nil {
		return m.DidBaseContext
	}
	return ""
}

func (m *Params) GetDidMethodContext() string {
	if m != nil {
		return m.DidMethodContext
	}
	return ""
}

func (m *Params) GetDidMethodName() string {
	if m != nil {
		return m.DidMethodName
	}
	return ""
}

func (m *Params) GetDidMethodVersion() string {
	if m != nil {
		return m.DidMethodVersion
	}
	return ""
}

func (m *Params) GetDidNetwork() string {
	if m != nil {
		return m.DidNetwork
	}
	return ""
}

func (m *Params) GetIpfsGateway() string {
	if m != nil {
		return m.IpfsGateway
	}
	return ""
}

func (m *Params) GetIpfsApi() string {
	if m != nil {
		return m.IpfsApi
	}
	return ""
}

func (m *Params) GetWebauthnAttestionPreference() string {
	if m != nil {
		return m.WebauthnAttestionPreference
	}
	return ""
}

func (m *Params) GetWebauthnAuthenticatorAttachment() string {
	if m != nil {
		return m.WebauthnAuthenticatorAttachment
	}
	return ""
}

func (m *Params) GetWebauthnTimeout() int32 {
	if m != nil {
		return m.WebauthnTimeout
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sonrhq.core.identity.GenesisState")
	proto.RegisterType((*Params)(nil), "sonrhq.core.identity.Params")
}

func init() { proto.RegisterFile("core/identity/genesis.proto", fileDescriptor_22ee3e6e2aad889c) }

var fileDescriptor_22ee3e6e2aad889c = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd3, 0x41, 0x4f, 0xdb, 0x3c,
	0x18, 0xc0, 0xf1, 0x86, 0x42, 0x01, 0x17, 0x5e, 0x78, 0x2d, 0xa4, 0x79, 0x30, 0x42, 0xcb, 0x61,
	0xeb, 0xa4, 0x29, 0x95, 0xd8, 0x6d, 0xb7, 0x76, 0x4c, 0x68, 0xd3, 0x86, 0x50, 0xd9, 0x38, 0xb0,
	0x43, 0xe5, 0x26, 0x4f, 0x9b, 0x47, 0x10, 0x3b, 0xb3, 0x5d, 0x4a, 0xbf, 0xc5, 0xbe, 0xc2, 0xbe,
	0x0d, 0xc7, 0x1e, 0x77, 0x9a, 0xa6, 0xf6, 0x8b, 0x4c, 0xb1, 0xdb, 0x50, 0x58, 0x2f, 0xbb, 0xb5,
	0x8f, 0x7f, 0xfe, 0x47, 0x4a, 0x6c, 0xb2, 0x17, 0x4a, 0x05, 0x75, 0x8c, 0x40, 0x18, 0x34, 0xc3,
	0x7a, 0x0f, 0x04, 0x68, 0xd4, 0x41, 0xaa, 0xa4, 0x91, 0x74, 0x47, 0x4b, 0xa1, 0xe2, 0x6f, 0x41,
	0x66, 0x82, 0x99, 0xd9, 0xdd, 0xe9, 0xc9, 0x9e, 0xb4, 0xa0, 0x9e, 0xfd, 0x72, 0x76, 0xf7, 0xc9,
	0xc3, 0x50, 0x84, 0x91, 0x5b, 0x38, 0xfc, 0x51, 0x24, 0x1b, 0x27, 0x2e, 0x7b, 0x6e, 0xb8, 0x01,
	0xfa, 0x86, 0x94, 0x52, 0xae, 0x78, 0xa2, 0x99, 0x57, 0xf1, 0x6a, 0xe5, 0xa3, 0x67, 0xc1, 0xa2,
	0xc7, 0x04, 0x67, 0xd6, 0x34, 0x97, 0xef, 0x7e, 0x1d, 0x14, 0x5a, 0xd3, 0x1d, 0xf4, 0x0b, 0xf9,
	0x3f, 0x55, 0x98, 0x70, 0x35, 0x7c, 0xef, 0x1c, 0x82, 0x66, 0x4b, 0x95, 0x62, 0xad, 0x7c, 0x54,
	0x5d, 0x9c, 0x39, 0xc6, 0xe8, 0x58, 0x86, 0xfd, 0x04, 0x84, 0x99, 0xb6, 0xfe, 0x2e, 0xd0, 0xaf,
	0x64, 0xa7, 0x73, 0x2d, 0xc3, 0xab, 0x30, 0xe6, 0x28, 0xe6, 0xca, 0xc5, 0x7f, 0x2b, 0x2f, 0x8c,
	0xd0, 0x77, 0xa4, 0xac, 0x41, 0xdd, 0x60, 0x08, 0x1f, 0x51, 0x1b, 0xb6, 0x6c, 0x9b, 0xfb, 0x8b,
	0x9b, 0xe7, 0x0e, 0x4e, 0x7b, 0xf3, 0xfb, 0xe8, 0x25, 0xd9, 0x54, 0x70, 0xcd, 0x0d, 0x4a, 0xa1,
	0x63, 0x4c, 0x35, 0x5b, 0xb1, 0xa1, 0x60, 0x71, 0xe8, 0x02, 0x14, 0x76, 0x31, 0xb4, 0xbc, 0x35,
	0xb7, 0x6d, 0x5a, 0x7e, 0x98, 0x3a, 0x1c, 0x15, 0x49, 0xc9, 0xbd, 0x6f, 0x5a, 0x23, 0xdb, 0x11,
	0x46, 0xed, 0x0e, 0xd7, 0xd0, 0x0e, 0xa5, 0x30, 0x70, 0x6b, 0xec, 0x77, 0x5a, 0x6f, 0xfd, 0x17,
	0x61, 0xd4, 0xe4, 0x1a, 0xde, 0xba, 0x29, 0x7d, 0x45, 0x68, 0x26, 0x13, 0x30, 0xb1, 0x8c, 0x72,
	0xbb, 0x64, 0x6d, 0xd6, 0xf8, 0x64, 0x17, 0x66, 0xfa, 0x39, 0xd9, 0x9a, 0xd3, 0x82, 0x27, 0xc0,
	0x8a, 0x96, 0x6e, 0xe6, 0xf4, 0x94, 0x27, 0xf0, 0xa8, 0x7a, 0x03, 0x4a, 0xa3, 0x14, 0x6c, 0xf9,
	0x51, 0xf5, 0xc2, 0xcd, 0xe9, 0x01, 0x29, 0x67, 0x5a, 0x80, 0x19, 0x48, 0x75, 0xc5, 0x56, 0x2c,
	0x23, 0x11, 0x46, 0xa7, 0x6e, 0x42, 0xab, 0x64, 0x03, 0xd3, 0xae, 0x6e, 0xf7, 0xb8, 0x81, 0x01,
	0x1f, 0xb2, 0x92, 0x15, 0xe5, 0x6c, 0x76, 0xe2, 0x46, 0xf4, 0x29, 0x59, 0xb3, 0x84, 0xa7, 0xc8,
	0x56, 0xed, 0xf2, 0x6a, 0xf6, 0xbf, 0x91, 0x22, 0x6d, 0x92, 0xfd, 0x01, 0x74, 0x78, 0xdf, 0xc4,
	0xa2, 0xcd, 0x8d, 0x01, 0x9d, 0xbd, 0xb2, 0x76, 0xaa, 0xa0, 0x0b, 0x0a, 0x44, 0x08, 0x6c, 0xcd,
	0xfa, 0xbd, 0x19, 0x6a, 0xcc, 0xcc, 0x59, 0x4e, 0xe8, 0x07, 0x52, 0xbd, 0x6f, 0xf4, 0x4d, 0x9c,
	0x7d, 0xa0, 0x90, 0x1b, 0xa9, 0xb2, 0x22, 0x0f, 0xe3, 0xec, 0xfc, 0xb0, 0x75, 0xdb, 0x39, 0xc8,
	0x3b, 0xf3, 0xae, 0x91, 0x33, 0xfa, 0x92, 0x6c, 0xe7, 0x2d, 0x83, 0x09, 0xc8, 0xbe, 0x61, 0xa4,
	0xe2, 0xd5, 0x56, 0x5a, 0x5b, 0xb3, 0xf9, 0x67, 0x37, 0x6e, 0x36, 0xee, 0xc6, 0xbe, 0x37, 0x1a,
	0xfb, 0xde, 0xef, 0xb1, 0xef, 0x7d, 0x9f, 0xf8, 0x85, 0xd1, 0xc4, 0x2f, 0xfc, 0x9c, 0xf8, 0x85,
	0xcb, 0x17, 0x3d, 0x34, 0x71, 0xbf, 0x13, 0x84, 0x32, 0xa9, 0xbb, 0xb3, 0x53, 0xb7, 0x77, 0xf7,
	0xf6, 0xfe, 0xf6, 0x9a, 0x61, 0x0a, 0xba, 0x53, 0xb2, 0x17, 0xf8, 0xf5, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x05, 0x25, 0x50, 0x98, 0x24, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relationships) > 0 {
		for iNdEx := len(m.Relationships) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Relationships[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ServiceList) > 0 {
		for iNdEx := len(m.ServiceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.BlockchainIdentities) > 0 {
		for iNdEx := len(m.BlockchainIdentities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlockchainIdentities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PrimaryIdentities) > 0 {
		for iNdEx := len(m.PrimaryIdentities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PrimaryIdentities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WebauthnTimeout != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.WebauthnTimeout))
		i--
		dAtA[i] = 0x50
	}
	if len(m.WebauthnAuthenticatorAttachment) > 0 {
		i -= len(m.WebauthnAuthenticatorAttachment)
		copy(dAtA[i:], m.WebauthnAuthenticatorAttachment)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.WebauthnAuthenticatorAttachment)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.WebauthnAttestionPreference) > 0 {
		i -= len(m.WebauthnAttestionPreference)
		copy(dAtA[i:], m.WebauthnAttestionPreference)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.WebauthnAttestionPreference)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.IpfsApi) > 0 {
		i -= len(m.IpfsApi)
		copy(dAtA[i:], m.IpfsApi)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IpfsApi)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.IpfsGateway) > 0 {
		i -= len(m.IpfsGateway)
		copy(dAtA[i:], m.IpfsGateway)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IpfsGateway)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DidNetwork) > 0 {
		i -= len(m.DidNetwork)
		copy(dAtA[i:], m.DidNetwork)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DidNetwork)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DidMethodVersion) > 0 {
		i -= len(m.DidMethodVersion)
		copy(dAtA[i:], m.DidMethodVersion)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DidMethodVersion)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DidMethodName) > 0 {
		i -= len(m.DidMethodName)
		copy(dAtA[i:], m.DidMethodName)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DidMethodName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DidMethodContext) > 0 {
		i -= len(m.DidMethodContext)
		copy(dAtA[i:], m.DidMethodContext)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DidMethodContext)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DidBaseContext) > 0 {
		i -= len(m.DidBaseContext)
		copy(dAtA[i:], m.DidBaseContext)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DidBaseContext)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PrimaryIdentities) > 0 {
		for _, e := range m.PrimaryIdentities {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlockchainIdentities) > 0 {
		for _, e := range m.BlockchainIdentities {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ServiceList) > 0 {
		for _, e := range m.ServiceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Relationships) > 0 {
		for _, e := range m.Relationships {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DidBaseContext)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.DidMethodContext)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.DidMethodName)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.DidMethodVersion)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.DidNetwork)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.IpfsGateway)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.IpfsApi)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.WebauthnAttestionPreference)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.WebauthnAuthenticatorAttachment)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.WebauthnTimeout != 0 {
		n += 1 + sovGenesis(uint64(m.WebauthnTimeout))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryIdentities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrimaryIdentities = append(m.PrimaryIdentities, DidDocument{})
			if err := m.PrimaryIdentities[len(m.PrimaryIdentities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockchainIdentities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockchainIdentities = append(m.BlockchainIdentities, DidDocument{})
			if err := m.BlockchainIdentities[len(m.BlockchainIdentities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceList = append(m.ServiceList, Service{})
			if err := m.ServiceList[len(m.ServiceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relationships", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relationships = append(m.Relationships, VerificationRelationship{})
			if err := m.Relationships[len(m.Relationships)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidBaseContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidBaseContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidMethodContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidMethodContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidMethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidMethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidMethodVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidMethodVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidNetwork", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidNetwork = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpfsGateway", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpfsGateway = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpfsApi", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpfsApi = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WebauthnAttestionPreference", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WebauthnAttestionPreference = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WebauthnAuthenticatorAttachment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WebauthnAuthenticatorAttachment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WebauthnTimeout", wireType)
			}
			m.WebauthnTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WebauthnTimeout |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)