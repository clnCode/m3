// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/aggregator/generated/proto/flush/flush.proto

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
	Package flush is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3/src/aggregator/generated/proto/flush/flush.proto

	It has these top-level messages:
		ShardSetFlushTimes
		ShardFlushTimes
		ForwardedFlushTimesForResolution
*/
package flush

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

type ShardSetFlushTimes struct {
	ByShard map[uint32]*ShardFlushTimes `protobuf:"bytes,1,rep,name=by_shard,json=byShard" json:"by_shard,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ShardSetFlushTimes) Reset()                    { *m = ShardSetFlushTimes{} }
func (m *ShardSetFlushTimes) String() string            { return proto.CompactTextString(m) }
func (*ShardSetFlushTimes) ProtoMessage()               {}
func (*ShardSetFlushTimes) Descriptor() ([]byte, []int) { return fileDescriptorFlush, []int{0} }

func (m *ShardSetFlushTimes) GetByShard() map[uint32]*ShardFlushTimes {
	if m != nil {
		return m.ByShard
	}
	return nil
}

type ShardFlushTimes struct {
	StandardByResolution  map[int64]int64                             `protobuf:"bytes,1,rep,name=standard_by_resolution,json=standardByResolution" json:"standard_by_resolution,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Tombstoned            bool                                        `protobuf:"varint,2,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	ForwardedByResolution map[int64]*ForwardedFlushTimesForResolution `protobuf:"bytes,3,rep,name=forwarded_by_resolution,json=forwardedByResolution" json:"forwarded_by_resolution,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	TimedByResolution     map[int64]int64                             `protobuf:"bytes,4,rep,name=timed_by_resolution,json=timedByResolution" json:"timed_by_resolution,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *ShardFlushTimes) Reset()                    { *m = ShardFlushTimes{} }
func (m *ShardFlushTimes) String() string            { return proto.CompactTextString(m) }
func (*ShardFlushTimes) ProtoMessage()               {}
func (*ShardFlushTimes) Descriptor() ([]byte, []int) { return fileDescriptorFlush, []int{1} }

func (m *ShardFlushTimes) GetStandardByResolution() map[int64]int64 {
	if m != nil {
		return m.StandardByResolution
	}
	return nil
}

func (m *ShardFlushTimes) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *ShardFlushTimes) GetForwardedByResolution() map[int64]*ForwardedFlushTimesForResolution {
	if m != nil {
		return m.ForwardedByResolution
	}
	return nil
}

func (m *ShardFlushTimes) GetTimedByResolution() map[int64]int64 {
	if m != nil {
		return m.TimedByResolution
	}
	return nil
}

type ForwardedFlushTimesForResolution struct {
	ByNumForwardedTimes map[int32]int64 `protobuf:"bytes,1,rep,name=by_num_forwarded_times,json=byNumForwardedTimes" json:"by_num_forwarded_times,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *ForwardedFlushTimesForResolution) Reset()         { *m = ForwardedFlushTimesForResolution{} }
func (m *ForwardedFlushTimesForResolution) String() string { return proto.CompactTextString(m) }
func (*ForwardedFlushTimesForResolution) ProtoMessage()    {}
func (*ForwardedFlushTimesForResolution) Descriptor() ([]byte, []int) {
	return fileDescriptorFlush, []int{2}
}

func (m *ForwardedFlushTimesForResolution) GetByNumForwardedTimes() map[int32]int64 {
	if m != nil {
		return m.ByNumForwardedTimes
	}
	return nil
}

func init() {
	proto.RegisterType((*ShardSetFlushTimes)(nil), "ShardSetFlushTimes")
	proto.RegisterType((*ShardFlushTimes)(nil), "ShardFlushTimes")
	proto.RegisterType((*ForwardedFlushTimesForResolution)(nil), "ForwardedFlushTimesForResolution")
}
func (m *ShardSetFlushTimes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShardSetFlushTimes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ByShard) > 0 {
		for k, _ := range m.ByShard {
			dAtA[i] = 0xa
			i++
			v := m.ByShard[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovFlush(uint64(msgSize))
			}
			mapSize := 1 + sovFlush(uint64(k)) + msgSize
			i = encodeVarintFlush(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintFlush(dAtA, i, uint64(k))
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintFlush(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func (m *ShardFlushTimes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShardFlushTimes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StandardByResolution) > 0 {
		for k, _ := range m.StandardByResolution {
			dAtA[i] = 0xa
			i++
			v := m.StandardByResolution[k]
			mapSize := 1 + sovFlush(uint64(k)) + 1 + sovFlush(uint64(v))
			i = encodeVarintFlush(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintFlush(dAtA, i, uint64(k))
			dAtA[i] = 0x10
			i++
			i = encodeVarintFlush(dAtA, i, uint64(v))
		}
	}
	if m.Tombstoned {
		dAtA[i] = 0x10
		i++
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.ForwardedByResolution) > 0 {
		for k, _ := range m.ForwardedByResolution {
			dAtA[i] = 0x1a
			i++
			v := m.ForwardedByResolution[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovFlush(uint64(msgSize))
			}
			mapSize := 1 + sovFlush(uint64(k)) + msgSize
			i = encodeVarintFlush(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintFlush(dAtA, i, uint64(k))
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintFlush(dAtA, i, uint64(v.Size()))
				n2, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n2
			}
		}
	}
	if len(m.TimedByResolution) > 0 {
		for k, _ := range m.TimedByResolution {
			dAtA[i] = 0x22
			i++
			v := m.TimedByResolution[k]
			mapSize := 1 + sovFlush(uint64(k)) + 1 + sovFlush(uint64(v))
			i = encodeVarintFlush(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintFlush(dAtA, i, uint64(k))
			dAtA[i] = 0x10
			i++
			i = encodeVarintFlush(dAtA, i, uint64(v))
		}
	}
	return i, nil
}

func (m *ForwardedFlushTimesForResolution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardedFlushTimesForResolution) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ByNumForwardedTimes) > 0 {
		for k, _ := range m.ByNumForwardedTimes {
			dAtA[i] = 0xa
			i++
			v := m.ByNumForwardedTimes[k]
			mapSize := 1 + sovFlush(uint64(k)) + 1 + sovFlush(uint64(v))
			i = encodeVarintFlush(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintFlush(dAtA, i, uint64(k))
			dAtA[i] = 0x10
			i++
			i = encodeVarintFlush(dAtA, i, uint64(v))
		}
	}
	return i, nil
}

func encodeVarintFlush(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ShardSetFlushTimes) Size() (n int) {
	var l int
	_ = l
	if len(m.ByShard) > 0 {
		for k, v := range m.ByShard {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovFlush(uint64(l))
			}
			mapEntrySize := 1 + sovFlush(uint64(k)) + l
			n += mapEntrySize + 1 + sovFlush(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ShardFlushTimes) Size() (n int) {
	var l int
	_ = l
	if len(m.StandardByResolution) > 0 {
		for k, v := range m.StandardByResolution {
			_ = k
			_ = v
			mapEntrySize := 1 + sovFlush(uint64(k)) + 1 + sovFlush(uint64(v))
			n += mapEntrySize + 1 + sovFlush(uint64(mapEntrySize))
		}
	}
	if m.Tombstoned {
		n += 2
	}
	if len(m.ForwardedByResolution) > 0 {
		for k, v := range m.ForwardedByResolution {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovFlush(uint64(l))
			}
			mapEntrySize := 1 + sovFlush(uint64(k)) + l
			n += mapEntrySize + 1 + sovFlush(uint64(mapEntrySize))
		}
	}
	if len(m.TimedByResolution) > 0 {
		for k, v := range m.TimedByResolution {
			_ = k
			_ = v
			mapEntrySize := 1 + sovFlush(uint64(k)) + 1 + sovFlush(uint64(v))
			n += mapEntrySize + 1 + sovFlush(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ForwardedFlushTimesForResolution) Size() (n int) {
	var l int
	_ = l
	if len(m.ByNumForwardedTimes) > 0 {
		for k, v := range m.ByNumForwardedTimes {
			_ = k
			_ = v
			mapEntrySize := 1 + sovFlush(uint64(k)) + 1 + sovFlush(uint64(v))
			n += mapEntrySize + 1 + sovFlush(uint64(mapEntrySize))
		}
	}
	return n
}

func sovFlush(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFlush(x uint64) (n int) {
	return sovFlush(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ShardSetFlushTimes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlush
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
			return fmt.Errorf("proto: ShardSetFlushTimes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShardSetFlushTimes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByShard", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFlush
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ByShard == nil {
				m.ByShard = make(map[uint32]*ShardFlushTimes)
			}
			var mapkey uint32
			var mapvalue *ShardFlushTimes
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFlush
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthFlush
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthFlush
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ShardFlushTimes{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFlush(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthFlush
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ByShard[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlush
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
func (m *ShardFlushTimes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlush
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
			return fmt.Errorf("proto: ShardFlushTimes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShardFlushTimes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandardByResolution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFlush
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StandardByResolution == nil {
				m.StandardByResolution = make(map[int64]int64)
			}
			var mapkey int64
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFlush
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFlush(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthFlush
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.StandardByResolution[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForwardedByResolution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFlush
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ForwardedByResolution == nil {
				m.ForwardedByResolution = make(map[int64]*ForwardedFlushTimesForResolution)
			}
			var mapkey int64
			var mapvalue *ForwardedFlushTimesForResolution
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFlush
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthFlush
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthFlush
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ForwardedFlushTimesForResolution{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFlush(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthFlush
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ForwardedByResolution[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimedByResolution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFlush
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TimedByResolution == nil {
				m.TimedByResolution = make(map[int64]int64)
			}
			var mapkey int64
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFlush
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFlush(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthFlush
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.TimedByResolution[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlush
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
func (m *ForwardedFlushTimesForResolution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlush
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
			return fmt.Errorf("proto: ForwardedFlushTimesForResolution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardedFlushTimesForResolution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByNumForwardedTimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFlush
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ByNumForwardedTimes == nil {
				m.ByNumForwardedTimes = make(map[int32]int64)
			}
			var mapkey int32
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFlush
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFlush
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFlush(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthFlush
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ByNumForwardedTimes[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlush
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
func skipFlush(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlush
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
					return 0, ErrIntOverflowFlush
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
					return 0, ErrIntOverflowFlush
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
				return 0, ErrInvalidLengthFlush
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFlush
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
				next, err := skipFlush(dAtA[start:])
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
	ErrInvalidLengthFlush = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlush   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/aggregator/generated/proto/flush/flush.proto", fileDescriptorFlush)
}

var fileDescriptorFlush = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xc5, 0x9b, 0x5d, 0x58, 0xcd, 0x82, 0x28, 0xde, 0xa5, 0x94, 0x1c, 0xa2, 0xd0, 0x03, 0x54,
	0x20, 0x25, 0xd2, 0xf6, 0x00, 0x5a, 0x6e, 0x05, 0xc2, 0x05, 0x71, 0x48, 0x91, 0x38, 0x06, 0xbb,
	0x71, 0xd3, 0x68, 0xeb, 0x18, 0xd9, 0x0e, 0xc8, 0x7f, 0xc1, 0x1f, 0xf0, 0x3b, 0x5c, 0x90, 0xb8,
	0x73, 0x41, 0xe5, 0x47, 0x50, 0x92, 0xb2, 0x09, 0x49, 0xaa, 0x8a, 0x8b, 0x95, 0xcc, 0x7b, 0x79,
	0xf3, 0x66, 0x5e, 0x0c, 0x2f, 0x92, 0x54, 0xaf, 0x72, 0xea, 0x2d, 0x04, 0xf7, 0xf9, 0x34, 0xa6,
	0x3e, 0x9f, 0xfa, 0x4a, 0x2e, 0x7c, 0x92, 0x24, 0x92, 0x25, 0x44, 0x0b, 0xe9, 0x27, 0x2c, 0x63,
	0x92, 0x68, 0x16, 0xfb, 0x1f, 0xa5, 0xd0, 0xc2, 0x5f, 0xae, 0x73, 0xb5, 0xaa, 0x4e, 0xaf, 0xac,
	0x8c, 0xbf, 0x22, 0xc0, 0xf3, 0x15, 0x91, 0xf1, 0x9c, 0xe9, 0xa0, 0xa8, 0xbf, 0x4b, 0x39, 0x53,
	0xf8, 0x39, 0x1c, 0x53, 0x13, 0xa9, 0x02, 0x18, 0x21, 0xd7, 0x9a, 0x9c, 0x9c, 0xbb, 0x5e, 0x97,
	0xe6, 0xcd, 0x4c, 0x59, 0x7c, 0x95, 0x69, 0x69, 0xc2, 0x1b, 0xb4, 0x7a, 0xb3, 0xdf, 0xc0, 0xcd,
	0x26, 0x80, 0x07, 0x60, 0x5d, 0x32, 0x33, 0x42, 0x2e, 0x9a, 0xdc, 0x0a, 0x8b, 0x47, 0xfc, 0x10,
	0x8e, 0x3e, 0x91, 0x75, 0xce, 0x46, 0x07, 0x2e, 0x9a, 0x9c, 0x9c, 0x0f, 0x2a, 0xed, 0x5a, 0x38,
	0xac, 0xe0, 0x8b, 0x83, 0x67, 0x68, 0xfc, 0xfd, 0x10, 0x6e, 0xb7, 0x60, 0xfc, 0x01, 0x86, 0x4a,
	0x93, 0x2c, 0x26, 0x32, 0x8e, 0xa8, 0x89, 0x24, 0x53, 0x62, 0x9d, 0xeb, 0x54, 0x64, 0x5b, 0xb3,
	0x8f, 0xdb, 0x82, 0xde, 0x7c, 0x4b, 0x9f, 0x99, 0xf0, 0x8a, 0x5c, 0xd9, 0x3e, 0x53, 0x3d, 0x10,
	0x76, 0x00, 0xb4, 0xe0, 0x54, 0x69, 0x91, 0xb1, 0xb8, 0xb4, 0x79, 0x1c, 0x36, 0x2a, 0x78, 0x01,
	0xf7, 0x96, 0x42, 0x7e, 0x26, 0x32, 0x66, 0x6d, 0x0b, 0x56, 0x69, 0xe1, 0x49, 0xc7, 0x42, 0xf0,
	0x97, 0xdf, 0xf5, 0x70, 0x77, 0xd9, 0x87, 0xe1, 0xf7, 0x70, 0xaa, 0x53, 0xde, 0x69, 0x70, 0x58,
	0x36, 0x78, 0xd4, 0x69, 0x50, 0x9c, 0x3d, 0xe2, 0x77, 0x74, 0xbb, 0x6e, 0xbf, 0x86, 0xfb, 0x3b,
	0x17, 0xd2, 0x8c, 0xcb, 0xaa, 0xe2, 0x3a, 0x6b, 0xc6, 0x65, 0x35, 0xc2, 0xb1, 0x2f, 0xc1, 0xde,
	0x3d, 0x56, 0x8f, 0xd2, 0xd3, 0x7f, 0x83, 0x7f, 0x50, 0x2f, 0xa5, 0x9e, 0x23, 0x10, 0xb2, 0x16,
	0x6a, 0x36, 0x7b, 0x09, 0xc3, 0xfe, 0x11, 0xff, 0xc7, 0xf2, 0xf8, 0x27, 0x02, 0x77, 0x5f, 0x57,
	0x2c, 0x60, 0x48, 0x4d, 0x94, 0xe5, 0x3c, 0xaa, 0x53, 0x2e, 0xd6, 0xa8, 0xb6, 0x3f, 0xd8, 0xc5,
	0x5e, 0xe3, 0xde, 0xcc, 0xbc, 0xcd, 0xf9, 0x15, 0xab, 0x24, 0x54, 0x79, 0x9c, 0xd2, 0x2e, 0x62,
	0x07, 0x30, 0xda, 0xf5, 0x41, 0x73, 0xba, 0xa3, 0x3d, 0xd3, 0xcd, 0x06, 0xdf, 0x36, 0x0e, 0xfa,
	0xb1, 0x71, 0xd0, 0xaf, 0x8d, 0x83, 0xbe, 0xfc, 0x76, 0xae, 0xd1, 0xeb, 0xe5, 0x45, 0x9f, 0xfe,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x79, 0x43, 0x3d, 0x62, 0x2f, 0x04, 0x00, 0x00,
}
