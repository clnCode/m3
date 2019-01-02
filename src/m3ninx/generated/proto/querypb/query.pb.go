// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/m3ninx/generated/proto/querypb/query.proto

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
	Package querypb is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3/src/m3ninx/generated/proto/querypb/query.proto

	It has these top-level messages:
		TermQuery
		RegexpQuery
		NegationQuery
		ConjunctionQuery
		DisjunctionQuery
		Query
*/
package querypb

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

type TermQuery struct {
	Field []byte `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Term  []byte `protobuf:"bytes,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (m *TermQuery) Reset()                    { *m = TermQuery{} }
func (m *TermQuery) String() string            { return proto.CompactTextString(m) }
func (*TermQuery) ProtoMessage()               {}
func (*TermQuery) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{0} }

func (m *TermQuery) GetField() []byte {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *TermQuery) GetTerm() []byte {
	if m != nil {
		return m.Term
	}
	return nil
}

type RegexpQuery struct {
	Field  []byte `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Regexp []byte `protobuf:"bytes,2,opt,name=regexp,proto3" json:"regexp,omitempty"`
}

func (m *RegexpQuery) Reset()                    { *m = RegexpQuery{} }
func (m *RegexpQuery) String() string            { return proto.CompactTextString(m) }
func (*RegexpQuery) ProtoMessage()               {}
func (*RegexpQuery) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{1} }

func (m *RegexpQuery) GetField() []byte {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *RegexpQuery) GetRegexp() []byte {
	if m != nil {
		return m.Regexp
	}
	return nil
}

type NegationQuery struct {
	Query *Query `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *NegationQuery) Reset()                    { *m = NegationQuery{} }
func (m *NegationQuery) String() string            { return proto.CompactTextString(m) }
func (*NegationQuery) ProtoMessage()               {}
func (*NegationQuery) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{2} }

func (m *NegationQuery) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type ConjunctionQuery struct {
	Queries []*Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
}

func (m *ConjunctionQuery) Reset()                    { *m = ConjunctionQuery{} }
func (m *ConjunctionQuery) String() string            { return proto.CompactTextString(m) }
func (*ConjunctionQuery) ProtoMessage()               {}
func (*ConjunctionQuery) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{3} }

func (m *ConjunctionQuery) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type DisjunctionQuery struct {
	Queries []*Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
}

func (m *DisjunctionQuery) Reset()                    { *m = DisjunctionQuery{} }
func (m *DisjunctionQuery) String() string            { return proto.CompactTextString(m) }
func (*DisjunctionQuery) ProtoMessage()               {}
func (*DisjunctionQuery) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{4} }

func (m *DisjunctionQuery) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Query struct {
	// Types that are valid to be assigned to Query:
	//	*Query_Term
	//	*Query_Regexp
	//	*Query_Negation
	//	*Query_Conjunction
	//	*Query_Disjunction
	Query isQuery_Query `protobuf_oneof:"query"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptorQuery, []int{5} }

type isQuery_Query interface {
	isQuery_Query()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Query_Term struct {
	Term *TermQuery `protobuf:"bytes,1,opt,name=term,oneof"`
}
type Query_Regexp struct {
	Regexp *RegexpQuery `protobuf:"bytes,2,opt,name=regexp,oneof"`
}
type Query_Negation struct {
	Negation *NegationQuery `protobuf:"bytes,3,opt,name=negation,oneof"`
}
type Query_Conjunction struct {
	Conjunction *ConjunctionQuery `protobuf:"bytes,4,opt,name=conjunction,oneof"`
}
type Query_Disjunction struct {
	Disjunction *DisjunctionQuery `protobuf:"bytes,5,opt,name=disjunction,oneof"`
}

func (*Query_Term) isQuery_Query()        {}
func (*Query_Regexp) isQuery_Query()      {}
func (*Query_Negation) isQuery_Query()    {}
func (*Query_Conjunction) isQuery_Query() {}
func (*Query_Disjunction) isQuery_Query() {}

func (m *Query) GetQuery() isQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *Query) GetTerm() *TermQuery {
	if x, ok := m.GetQuery().(*Query_Term); ok {
		return x.Term
	}
	return nil
}

func (m *Query) GetRegexp() *RegexpQuery {
	if x, ok := m.GetQuery().(*Query_Regexp); ok {
		return x.Regexp
	}
	return nil
}

func (m *Query) GetNegation() *NegationQuery {
	if x, ok := m.GetQuery().(*Query_Negation); ok {
		return x.Negation
	}
	return nil
}

func (m *Query) GetConjunction() *ConjunctionQuery {
	if x, ok := m.GetQuery().(*Query_Conjunction); ok {
		return x.Conjunction
	}
	return nil
}

func (m *Query) GetDisjunction() *DisjunctionQuery {
	if x, ok := m.GetQuery().(*Query_Disjunction); ok {
		return x.Disjunction
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Query) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Query_OneofMarshaler, _Query_OneofUnmarshaler, _Query_OneofSizer, []interface{}{
		(*Query_Term)(nil),
		(*Query_Regexp)(nil),
		(*Query_Negation)(nil),
		(*Query_Conjunction)(nil),
		(*Query_Disjunction)(nil),
	}
}

func _Query_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Query)
	// query
	switch x := m.Query.(type) {
	case *Query_Term:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Term); err != nil {
			return err
		}
	case *Query_Regexp:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Regexp); err != nil {
			return err
		}
	case *Query_Negation:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Negation); err != nil {
			return err
		}
	case *Query_Conjunction:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Conjunction); err != nil {
			return err
		}
	case *Query_Disjunction:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Disjunction); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Query.Query has unexpected type %T", x)
	}
	return nil
}

func _Query_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Query)
	switch tag {
	case 1: // query.term
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TermQuery)
		err := b.DecodeMessage(msg)
		m.Query = &Query_Term{msg}
		return true, err
	case 2: // query.regexp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegexpQuery)
		err := b.DecodeMessage(msg)
		m.Query = &Query_Regexp{msg}
		return true, err
	case 3: // query.negation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NegationQuery)
		err := b.DecodeMessage(msg)
		m.Query = &Query_Negation{msg}
		return true, err
	case 4: // query.conjunction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConjunctionQuery)
		err := b.DecodeMessage(msg)
		m.Query = &Query_Conjunction{msg}
		return true, err
	case 5: // query.disjunction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DisjunctionQuery)
		err := b.DecodeMessage(msg)
		m.Query = &Query_Disjunction{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Query_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Query)
	// query
	switch x := m.Query.(type) {
	case *Query_Term:
		s := proto.Size(x.Term)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_Regexp:
		s := proto.Size(x.Regexp)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_Negation:
		s := proto.Size(x.Negation)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_Conjunction:
		s := proto.Size(x.Conjunction)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_Disjunction:
		s := proto.Size(x.Disjunction)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*TermQuery)(nil), "query.TermQuery")
	proto.RegisterType((*RegexpQuery)(nil), "query.RegexpQuery")
	proto.RegisterType((*NegationQuery)(nil), "query.NegationQuery")
	proto.RegisterType((*ConjunctionQuery)(nil), "query.ConjunctionQuery")
	proto.RegisterType((*DisjunctionQuery)(nil), "query.DisjunctionQuery")
	proto.RegisterType((*Query)(nil), "query.Query")
}
func (m *TermQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TermQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Field) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Field)))
		i += copy(dAtA[i:], m.Field)
	}
	if len(m.Term) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Term)))
		i += copy(dAtA[i:], m.Term)
	}
	return i, nil
}

func (m *RegexpQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegexpQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Field) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Field)))
		i += copy(dAtA[i:], m.Field)
	}
	if len(m.Regexp) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Regexp)))
		i += copy(dAtA[i:], m.Regexp)
	}
	return i, nil
}

func (m *NegationQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NegationQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Query != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Query.Size()))
		n1, err := m.Query.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ConjunctionQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConjunctionQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, msg := range m.Queries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintQuery(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DisjunctionQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisjunctionQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, msg := range m.Queries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintQuery(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Query) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Query) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Query != nil {
		nn2, err := m.Query.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	return i, nil
}

func (m *Query_Term) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Term != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Term.Size()))
		n3, err := m.Term.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *Query_Regexp) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Regexp != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Regexp.Size()))
		n4, err := m.Regexp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *Query_Negation) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Negation != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Negation.Size()))
		n5, err := m.Negation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *Query_Conjunction) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Conjunction != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Conjunction.Size()))
		n6, err := m.Conjunction.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *Query_Disjunction) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Disjunction != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Disjunction.Size()))
		n7, err := m.Disjunction.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TermQuery) Size() (n int) {
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Term)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *RegexpQuery) Size() (n int) {
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Regexp)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *NegationQuery) Size() (n int) {
	var l int
	_ = l
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ConjunctionQuery) Size() (n int) {
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *DisjunctionQuery) Size() (n int) {
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *Query) Size() (n int) {
	var l int
	_ = l
	if m.Query != nil {
		n += m.Query.Size()
	}
	return n
}

func (m *Query_Term) Size() (n int) {
	var l int
	_ = l
	if m.Term != nil {
		l = m.Term.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}
func (m *Query_Regexp) Size() (n int) {
	var l int
	_ = l
	if m.Regexp != nil {
		l = m.Regexp.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}
func (m *Query_Negation) Size() (n int) {
	var l int
	_ = l
	if m.Negation != nil {
		l = m.Negation.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}
func (m *Query_Conjunction) Size() (n int) {
	var l int
	_ = l
	if m.Conjunction != nil {
		l = m.Conjunction.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}
func (m *Query_Disjunction) Size() (n int) {
	var l int
	_ = l
	if m.Disjunction != nil {
		l = m.Disjunction.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TermQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: TermQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TermQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = append(m.Field[:0], dAtA[iNdEx:postIndex]...)
			if m.Field == nil {
				m.Field = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Term = append(m.Term[:0], dAtA[iNdEx:postIndex]...)
			if m.Term == nil {
				m.Term = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
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
func (m *RegexpQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: RegexpQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegexpQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = append(m.Field[:0], dAtA[iNdEx:postIndex]...)
			if m.Field == nil {
				m.Field = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regexp", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Regexp = append(m.Regexp[:0], dAtA[iNdEx:postIndex]...)
			if m.Regexp == nil {
				m.Regexp = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
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
func (m *NegationQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: NegationQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NegationQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Query == nil {
				m.Query = &Query{}
			}
			if err := m.Query.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
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
func (m *ConjunctionQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ConjunctionQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConjunctionQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, &Query{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
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
func (m *DisjunctionQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: DisjunctionQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DisjunctionQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, &Query{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
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
func (m *Query) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: Query: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Query: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TermQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Query = &Query_Term{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regexp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RegexpQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Query = &Query_Regexp{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Negation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NegationQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Query = &Query_Negation{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conjunction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ConjunctionQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Query = &Query_Conjunction{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disjunction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DisjunctionQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Query = &Query_Disjunction{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQuery
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
				next, err := skipQuery(dAtA[start:])
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
	ErrInvalidLengthQuery = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/m3ninx/generated/proto/querypb/query.proto", fileDescriptorQuery)
}

var fileDescriptorQuery = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0xb6, 0x69, 0xe1, 0xa6, 0x48, 0x95, 0x55, 0x81, 0x59, 0xa2, 0x2a, 0x03, 0x62,
	0x40, 0xb5, 0xd4, 0x88, 0x85, 0x4e, 0x14, 0x86, 0x4c, 0x48, 0x44, 0x4c, 0x6c, 0x4d, 0x62, 0x82,
	0x11, 0x76, 0x8a, 0x9b, 0x4a, 0xe5, 0x2d, 0x18, 0x78, 0x28, 0x46, 0x1e, 0x01, 0x95, 0x17, 0x41,
	0xb1, 0xdd, 0x5f, 0x24, 0x06, 0xa6, 0xe4, 0xfa, 0x9e, 0x4f, 0x3a, 0xf9, 0x62, 0xb8, 0xcc, 0x79,
	0xf9, 0x38, 0x4b, 0xfa, 0x69, 0x21, 0xa8, 0x08, 0xb3, 0x84, 0x8a, 0x90, 0x4e, 0x55, 0x4a, 0x45,
	0x28, 0xb9, 0x9c, 0xd3, 0x9c, 0x49, 0xa6, 0xc6, 0x25, 0xcb, 0xe8, 0x44, 0x15, 0x65, 0x41, 0x5f,
	0x66, 0x4c, 0xbd, 0x4e, 0x12, 0xf3, 0xec, 0xeb, 0x33, 0xec, 0xea, 0x21, 0x38, 0x87, 0xfd, 0x3b,
	0xa6, 0xc4, 0x6d, 0x35, 0xe0, 0x2e, 0xb8, 0x0f, 0x9c, 0x3d, 0x67, 0x04, 0xf5, 0xd0, 0x69, 0x3b,
	0x36, 0x03, 0xc6, 0xd0, 0x28, 0x99, 0x12, 0xa4, 0xa6, 0x0f, 0xf5, 0x7b, 0x30, 0x04, 0x2f, 0x66,
	0x39, 0x9b, 0x4f, 0xfe, 0x02, 0x0f, 0xa1, 0xa9, 0x74, 0xc8, 0xa2, 0x76, 0x0a, 0x42, 0x38, 0xb8,
	0x61, 0xf9, 0xb8, 0xe4, 0x85, 0x34, 0x78, 0x00, 0xa6, 0x8d, 0xc6, 0xbd, 0x41, 0xbb, 0x6f, 0x8a,
	0xea, 0x65, 0x6c, 0x8b, 0x5e, 0x40, 0xe7, 0xaa, 0x90, 0x4f, 0x33, 0x99, 0xae, 0xb9, 0x13, 0x68,
	0x55, 0x4b, 0xce, 0xa6, 0x04, 0xf5, 0xea, 0xbf, 0xc8, 0xe5, 0xb2, 0x62, 0xaf, 0xf9, 0xf4, 0x7f,
	0xec, 0x7b, 0x0d, 0xdc, 0x25, 0x61, 0x3c, 0x98, 0x92, 0x1d, 0x1b, 0x5f, 0xd9, 0x8b, 0x1c, 0xe3,
	0x06, 0x9f, 0x6d, 0x7d, 0xb6, 0x37, 0xc0, 0x36, 0xb9, 0x21, 0x2c, 0x72, 0x96, 0x32, 0xf0, 0x00,
	0xf6, 0xa4, 0x95, 0x41, 0xea, 0x3a, 0xdf, 0xb5, 0xf9, 0x2d, 0x47, 0x91, 0x13, 0xaf, 0x72, 0x78,
	0x08, 0x5e, 0xba, 0x76, 0x41, 0x1a, 0x1a, 0x3b, 0xb2, 0xd8, 0xae, 0xa5, 0xc8, 0x89, 0x37, 0xd3,
	0x15, 0x9c, 0xad, 0x65, 0x10, 0x77, 0x0b, 0xde, 0xd5, 0x54, 0xc1, 0x1b, 0xe9, 0x51, 0xcb, 0xfe,
	0xa9, 0xd1, 0xf1, 0xc7, 0xc2, 0x47, 0x9f, 0x0b, 0x1f, 0x7d, 0x2d, 0x7c, 0xf4, 0xf6, 0xed, 0x3b,
	0xf7, 0x2d, 0x7b, 0xcb, 0x92, 0xa6, 0xbe, 0x60, 0xe1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87,
	0xea, 0x95, 0x8a, 0xa5, 0x02, 0x00, 0x00,
}
