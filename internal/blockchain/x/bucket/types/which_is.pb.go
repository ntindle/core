// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bucket/which_is.proto

package types

import (
	fmt "fmt"
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

type WhichIs struct {
	// DID is the DID of the bucket.
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Creator is the Account that created the bucket.
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	// Bucket is the document of the bucket.
	Bucket *BucketDoc `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Timestamp is the time of the last update of the DID Document
	Timestamp int64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// IsActive is the status of the DID Document
	IsActive bool `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (m *WhichIs) Reset()         { *m = WhichIs{} }
func (m *WhichIs) String() string { return proto.CompactTextString(m) }
func (*WhichIs) ProtoMessage()    {}
func (*WhichIs) Descriptor() ([]byte, []int) {
	return fileDescriptor_b12711a1b047ec75, []int{0}
}
func (m *WhichIs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhichIs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhichIs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhichIs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhichIs.Merge(m, src)
}
func (m *WhichIs) XXX_Size() int {
	return m.Size()
}
func (m *WhichIs) XXX_DiscardUnknown() {
	xxx_messageInfo_WhichIs.DiscardUnknown(m)
}

var xxx_messageInfo_WhichIs proto.InternalMessageInfo

func (m *WhichIs) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *WhichIs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *WhichIs) GetBucket() *BucketDoc {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *WhichIs) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *WhichIs) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func init() {
	proto.RegisterType((*WhichIs)(nil), "sonrio.sonr.bucket.WhichIs")
}

func init() { proto.RegisterFile("bucket/which_is.proto", fileDescriptor_b12711a1b047ec75) }

var fileDescriptor_b12711a1b047ec75 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2a, 0x4d, 0xce,
	0x4e, 0x2d, 0xd1, 0x2f, 0xcf, 0xc8, 0x4c, 0xce, 0x88, 0xcf, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x2a, 0xce, 0xcf, 0x2b, 0xca, 0xcc, 0xd7, 0x03, 0x51, 0x7a, 0x10, 0x25, 0x52,
	0xc2, 0x50, 0xa5, 0x10, 0x0a, 0xa2, 0x50, 0x69, 0x39, 0x23, 0x17, 0x7b, 0x38, 0x48, 0xaf, 0x67,
	0xb1, 0x90, 0x00, 0x17, 0x73, 0x4a, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88,
	0x29, 0x24, 0xc1, 0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0x0c, 0x16, 0x85,
	0x71, 0x85, 0x4c, 0xb9, 0xd8, 0x20, 0xe6, 0x48, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x1b, 0xc9, 0xea,
	0x61, 0xda, 0xa8, 0xe7, 0x04, 0xa6, 0x5c, 0xf2, 0x93, 0x83, 0xa0, 0x8a, 0x85, 0x64, 0xb8, 0x38,
	0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x24, 0x58, 0x15, 0x18, 0x35, 0x98, 0x83,
	0x10, 0x02, 0x42, 0xd2, 0x5c, 0x9c, 0x99, 0xc5, 0xf1, 0x89, 0xc9, 0x25, 0x99, 0x65, 0xa9, 0x12,
	0x6c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x1c, 0x99, 0xc5, 0x8e, 0x60, 0xbe, 0x93, 0xdb, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0x83, 0xac, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x4e,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xaf, 0x80, 0xfa, 0x58, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0xec, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xdf, 0xb6, 0xed, 0x3a, 0x01,
	0x00, 0x00,
}

func (m *WhichIs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhichIs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhichIs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Timestamp != 0 {
		i = encodeVarintWhichIs(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x28
	}
	if m.Bucket != nil {
		{
			size, err := m.Bucket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWhichIs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWhichIs(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintWhichIs(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhichIs(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhichIs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhichIs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovWhichIs(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWhichIs(uint64(l))
	}
	if m.Bucket != nil {
		l = m.Bucket.Size()
		n += 1 + l + sovWhichIs(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovWhichIs(uint64(m.Timestamp))
	}
	if m.IsActive {
		n += 2
	}
	return n
}

func sovWhichIs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhichIs(x uint64) (n int) {
	return sovWhichIs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhichIs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhichIs
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
			return fmt.Errorf("proto: WhichIs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhichIs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhichIs
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
				return ErrInvalidLengthWhichIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhichIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhichIs
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
				return ErrInvalidLengthWhichIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhichIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bucket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhichIs
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
				return ErrInvalidLengthWhichIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhichIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bucket == nil {
				m.Bucket = &BucketDoc{}
			}
			if err := m.Bucket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhichIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhichIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipWhichIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhichIs
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
func skipWhichIs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhichIs
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
					return 0, ErrIntOverflowWhichIs
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
					return 0, ErrIntOverflowWhichIs
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
				return 0, ErrInvalidLengthWhichIs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhichIs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhichIs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhichIs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhichIs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhichIs = fmt.Errorf("proto: unexpected end of group")
)