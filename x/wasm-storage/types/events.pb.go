// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sedachain/wasm_storage/v1/events.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// The msg for storing a data request wasm.
type EventStoreDataRequestWasm struct {
	Hash     string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Bytecode []byte `protobuf:"bytes,2,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
}

func (m *EventStoreDataRequestWasm) Reset()         { *m = EventStoreDataRequestWasm{} }
func (m *EventStoreDataRequestWasm) String() string { return proto.CompactTextString(m) }
func (*EventStoreDataRequestWasm) ProtoMessage()    {}
func (*EventStoreDataRequestWasm) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b0422555944a3a, []int{0}
}
func (m *EventStoreDataRequestWasm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventStoreDataRequestWasm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventStoreDataRequestWasm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventStoreDataRequestWasm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStoreDataRequestWasm.Merge(m, src)
}
func (m *EventStoreDataRequestWasm) XXX_Size() int {
	return m.Size()
}
func (m *EventStoreDataRequestWasm) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStoreDataRequestWasm.DiscardUnknown(m)
}

var xxx_messageInfo_EventStoreDataRequestWasm proto.InternalMessageInfo

func (m *EventStoreDataRequestWasm) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *EventStoreDataRequestWasm) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

// The msg for storing a overlay wasm(i.e. relayer or executor)
type EventStoreOverlayWasm struct {
	Hash     string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Bytecode []byte `protobuf:"bytes,2,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
}

func (m *EventStoreOverlayWasm) Reset()         { *m = EventStoreOverlayWasm{} }
func (m *EventStoreOverlayWasm) String() string { return proto.CompactTextString(m) }
func (*EventStoreOverlayWasm) ProtoMessage()    {}
func (*EventStoreOverlayWasm) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b0422555944a3a, []int{1}
}
func (m *EventStoreOverlayWasm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventStoreOverlayWasm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventStoreOverlayWasm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventStoreOverlayWasm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStoreOverlayWasm.Merge(m, src)
}
func (m *EventStoreOverlayWasm) XXX_Size() int {
	return m.Size()
}
func (m *EventStoreOverlayWasm) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStoreOverlayWasm.DiscardUnknown(m)
}

var xxx_messageInfo_EventStoreOverlayWasm proto.InternalMessageInfo

func (m *EventStoreOverlayWasm) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *EventStoreOverlayWasm) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

func init() {
	proto.RegisterType((*EventStoreDataRequestWasm)(nil), "sedachain.wasm_storage.v1.EventStoreDataRequestWasm")
	proto.RegisterType((*EventStoreOverlayWasm)(nil), "sedachain.wasm_storage.v1.EventStoreOverlayWasm")
}

func init() {
	proto.RegisterFile("sedachain/wasm_storage/v1/events.proto", fileDescriptor_90b0422555944a3a)
}

var fileDescriptor_90b0422555944a3a = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x4e, 0x4d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x4f, 0x2c, 0xce, 0x8d, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a,
	0x4c, 0x4f, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x84, 0xab, 0xd3, 0x43, 0x56, 0xa7, 0x57, 0x66, 0x28, 0xa5, 0x83, 0xdb,
	0x08, 0x14, 0xa5, 0x60, 0x83, 0x94, 0xbc, 0xb9, 0x24, 0x5d, 0x41, 0x06, 0x07, 0x97, 0xe4, 0x17,
	0xa5, 0xba, 0x24, 0x96, 0x24, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x84, 0x27, 0x16, 0xe7,
	0x0a, 0x09, 0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x42, 0x52, 0x5c, 0x1c, 0x49, 0x95, 0x25, 0xa9, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x70, 0xbe, 0x92, 0x3b, 0x97, 0x28, 0xc2, 0x30, 0xff, 0xb2, 0xd4, 0xa2,
	0x9c, 0xc4, 0x4a, 0x72, 0x0c, 0x72, 0x0a, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0xf3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x90, 0x47,
	0xc1, 0xbe, 0x48, 0xce, 0xcf, 0x01, 0x73, 0x74, 0x21, 0xde, 0xae, 0x00, 0x7b, 0x54, 0x17, 0xe6,
	0xf1, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x4a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xab, 0x7c, 0xd4, 0x59, 0x62, 0x01, 0x00, 0x00,
}

func (m *EventStoreDataRequestWasm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventStoreDataRequestWasm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventStoreDataRequestWasm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventStoreOverlayWasm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventStoreOverlayWasm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventStoreOverlayWasm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventStoreDataRequestWasm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventStoreOverlayWasm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventStoreDataRequestWasm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventStoreDataRequestWasm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventStoreDataRequestWasm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytecode = append(m.Bytecode[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytecode == nil {
				m.Bytecode = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventStoreOverlayWasm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventStoreOverlayWasm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventStoreOverlayWasm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytecode = append(m.Bytecode[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytecode == nil {
				m.Bytecode = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
