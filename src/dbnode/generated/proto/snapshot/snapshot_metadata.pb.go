// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/dbnode/generated/proto/snapshot/snapshot_metadata.proto

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
	Package snapshot is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3/src/dbnode/generated/proto/snapshot/snapshot_metadata.proto

	It has these top-level messages:
		Metadata
*/
package snapshot

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Metadata struct {
	SnapshotIndex int64  `protobuf:"varint,1,opt,name=snapshotIndex,proto3" json:"snapshotIndex,omitempty"`
	SnapshotUUID  []byte `protobuf:"bytes,2,opt,name=snapshotUUID,proto3" json:"snapshotUUID,omitempty"`
	CommitlogID   []byte `protobuf:"bytes,3,opt,name=commitlogID,proto3" json:"commitlogID,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorSnapshotMetadata, []int{0} }

func (m *Metadata) GetSnapshotIndex() int64 {
	if m != nil {
		return m.SnapshotIndex
	}
	return 0
}

func (m *Metadata) GetSnapshotUUID() []byte {
	if m != nil {
		return m.SnapshotUUID
	}
	return nil
}

func (m *Metadata) GetCommitlogID() []byte {
	if m != nil {
		return m.CommitlogID
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "snapshot.Metadata")
}
func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SnapshotIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSnapshotMetadata(dAtA, i, uint64(m.SnapshotIndex))
	}
	if len(m.SnapshotUUID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSnapshotMetadata(dAtA, i, uint64(len(m.SnapshotUUID)))
		i += copy(dAtA[i:], m.SnapshotUUID)
	}
	if len(m.CommitlogID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSnapshotMetadata(dAtA, i, uint64(len(m.CommitlogID)))
		i += copy(dAtA[i:], m.CommitlogID)
	}
	return i, nil
}

func encodeVarintSnapshotMetadata(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Metadata) Size() (n int) {
	var l int
	_ = l
	if m.SnapshotIndex != 0 {
		n += 1 + sovSnapshotMetadata(uint64(m.SnapshotIndex))
	}
	l = len(m.SnapshotUUID)
	if l > 0 {
		n += 1 + l + sovSnapshotMetadata(uint64(l))
	}
	l = len(m.CommitlogID)
	if l > 0 {
		n += 1 + l + sovSnapshotMetadata(uint64(l))
	}
	return n
}

func sovSnapshotMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSnapshotMetadata(x uint64) (n int) {
	return sovSnapshotMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshotMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotIndex", wireType)
			}
			m.SnapshotIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshotMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SnapshotIndex |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotUUID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshotMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSnapshotMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotUUID = append(m.SnapshotUUID[:0], dAtA[iNdEx:postIndex]...)
			if m.SnapshotUUID == nil {
				m.SnapshotUUID = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitlogID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshotMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSnapshotMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitlogID = append(m.CommitlogID[:0], dAtA[iNdEx:postIndex]...)
			if m.CommitlogID == nil {
				m.CommitlogID = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshotMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshotMetadata
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
func skipSnapshotMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshotMetadata
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
					return 0, ErrIntOverflowSnapshotMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSnapshotMetadata
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSnapshotMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSnapshotMetadata
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSnapshotMetadata(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSnapshotMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshotMetadata   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/dbnode/generated/proto/snapshot/snapshot_metadata.proto", fileDescriptorSnapshotMetadata)
}

var fileDescriptorSnapshotMetadata = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x35, 0x4e, 0x49, 0xd2, 0xcf, 0x35, 0xd6, 0x2f,
	0x2e, 0x4a, 0xd6, 0x4f, 0x49, 0xca, 0xcb, 0x4f, 0x49, 0xd5, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a,
	0x2c, 0x49, 0x4d, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0xce, 0x4b, 0x2c, 0x28, 0xce,
	0xc8, 0x2f, 0x81, 0x33, 0xe2, 0x73, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0xf5, 0xc0, 0x0a,
	0x84, 0x38, 0x60, 0x12, 0x4a, 0x65, 0x5c, 0x1c, 0xbe, 0x50, 0x39, 0x21, 0x15, 0x2e, 0x5e, 0x98,
	0xb8, 0x67, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0xaa, 0xa0, 0x90,
	0x12, 0x17, 0x0f, 0x4c, 0x20, 0x34, 0xd4, 0xd3, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08,
	0x45, 0x4c, 0x48, 0x81, 0x8b, 0x3b, 0x39, 0x3f, 0x37, 0x37, 0xb3, 0x24, 0x27, 0x3f, 0xdd, 0xd3,
	0x45, 0x82, 0x19, 0xac, 0x04, 0x59, 0xc8, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x34, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xc3, 0xc8, 0x99, 0xec, 0x00, 0x00, 0x00,
}
