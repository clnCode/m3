// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/cluster/generated/proto/changesetpb/changeset.proto

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
	Package changesetpb is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3/src/cluster/generated/proto/changesetpb/changeset.proto

	It has these top-level messages:
		ChangeSet
*/
package changesetpb

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

// ChangeSetState tracks the stateof a changeset
type ChangeSetState int32

const (
	ChangeSetState_UNKNOWN ChangeSetState = 0
	ChangeSetState_OPEN    ChangeSetState = 1
	ChangeSetState_CLOSED  ChangeSetState = 2
)

var ChangeSetState_name = map[int32]string{
	0: "UNKNOWN",
	1: "OPEN",
	2: "CLOSED",
}
var ChangeSetState_value = map[string]int32{
	"UNKNOWN": 0,
	"OPEN":    1,
	"CLOSED":  2,
}

func (x ChangeSetState) String() string {
	return proto.EnumName(ChangeSetState_name, int32(x))
}
func (ChangeSetState) EnumDescriptor() ([]byte, []int) { return fileDescriptorChangeset, []int{0} }

// A ChangeSet is a set of changes that are applied together.  The exact
// format of the changes is up to the application; the ChangeSet simply
// tracks the state of application
type ChangeSet struct {
	// for_version is the version of configuration on which this ChangeSet is built
	ForVersion int32 `protobuf:"varint,1,opt,name=for_version,json=forVersion,proto3" json:"for_version,omitempty"`
	// state is the state of the ChangeSet
	State ChangeSetState `protobuf:"varint,2,opt,name=state,proto3,enum=changesetpb.ChangeSetState" json:"state,omitempty"`
	// changes are the marshalled form of the changes
	Changes []byte `protobuf:"bytes,3,opt,name=changes,proto3" json:"changes,omitempty"`
}

func (m *ChangeSet) Reset()                    { *m = ChangeSet{} }
func (m *ChangeSet) String() string            { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()               {}
func (*ChangeSet) Descriptor() ([]byte, []int) { return fileDescriptorChangeset, []int{0} }

func (m *ChangeSet) GetForVersion() int32 {
	if m != nil {
		return m.ForVersion
	}
	return 0
}

func (m *ChangeSet) GetState() ChangeSetState {
	if m != nil {
		return m.State
	}
	return ChangeSetState_UNKNOWN
}

func (m *ChangeSet) GetChanges() []byte {
	if m != nil {
		return m.Changes
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeSet)(nil), "changesetpb.ChangeSet")
	proto.RegisterEnum("changesetpb.ChangeSetState", ChangeSetState_name, ChangeSetState_value)
}
func (m *ChangeSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ForVersion != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintChangeset(dAtA, i, uint64(m.ForVersion))
	}
	if m.State != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintChangeset(dAtA, i, uint64(m.State))
	}
	if len(m.Changes) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintChangeset(dAtA, i, uint64(len(m.Changes)))
		i += copy(dAtA[i:], m.Changes)
	}
	return i, nil
}

func encodeVarintChangeset(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ChangeSet) Size() (n int) {
	var l int
	_ = l
	if m.ForVersion != 0 {
		n += 1 + sovChangeset(uint64(m.ForVersion))
	}
	if m.State != 0 {
		n += 1 + sovChangeset(uint64(m.State))
	}
	l = len(m.Changes)
	if l > 0 {
		n += 1 + l + sovChangeset(uint64(l))
	}
	return n
}

func sovChangeset(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozChangeset(x uint64) (n int) {
	return sovChangeset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChangeSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChangeset
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
			return fmt.Errorf("proto: ChangeSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForVersion", wireType)
			}
			m.ForVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChangeset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ForVersion |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChangeset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (ChangeSetState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChangeset
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
				return ErrInvalidLengthChangeset
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes[:0], dAtA[iNdEx:postIndex]...)
			if m.Changes == nil {
				m.Changes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChangeset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChangeset
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
func skipChangeset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChangeset
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
					return 0, ErrIntOverflowChangeset
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
					return 0, ErrIntOverflowChangeset
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
				return 0, ErrInvalidLengthChangeset
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowChangeset
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
				next, err := skipChangeset(dAtA[start:])
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
	ErrInvalidLengthChangeset = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChangeset   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/cluster/generated/proto/changesetpb/changeset.proto", fileDescriptorChangeset)
}

var fileDescriptorChangeset = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x35, 0x4e, 0x49, 0xd2, 0xcf, 0x35, 0xd6, 0x2f,
	0x2e, 0x4a, 0xd6, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0xd2, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d,
	0x4a, 0x2c, 0x49, 0x4d, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xce, 0x48, 0xcc, 0x4b,
	0x4f, 0x2d, 0x4e, 0x2d, 0x29, 0x48, 0x42, 0xb0, 0xf5, 0xc0, 0x72, 0x42, 0xdc, 0x48, 0x92, 0x4a,
	0x95, 0x5c, 0x9c, 0xce, 0x60, 0x6e, 0x70, 0x6a, 0x89, 0x90, 0x3c, 0x17, 0x77, 0x5a, 0x7e, 0x51,
	0x7c, 0x59, 0x6a, 0x51, 0x71, 0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x57,
	0x5a, 0x7e, 0x51, 0x18, 0x44, 0x44, 0xc8, 0x90, 0x8b, 0xb5, 0xb8, 0x24, 0xb1, 0x24, 0x55, 0x82,
	0x49, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x5a, 0x0f, 0xc9, 0x28, 0x3d, 0xb8, 0x39, 0xc1, 0x20, 0x25,
	0x41, 0x10, 0x95, 0x42, 0x12, 0x5c, 0xec, 0x50, 0x45, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x3c, 0x41,
	0x30, 0xae, 0x96, 0x31, 0x17, 0x1f, 0xaa, 0x16, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f,
	0xff, 0x70, 0x3f, 0x01, 0x06, 0x21, 0x0e, 0x2e, 0x16, 0xff, 0x00, 0x57, 0x3f, 0x01, 0x46, 0x21,
	0x2e, 0x2e, 0x36, 0x67, 0x1f, 0xff, 0x60, 0x57, 0x17, 0x01, 0x26, 0x27, 0x81, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x24, 0x36,
	0xb0, 0xaf, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xed, 0x57, 0x5c, 0x23, 0x01, 0x00,
	0x00,
}
