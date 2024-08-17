// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sedachain/data_proxy/v1/data_proxy.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Module parameters which can be changed through governance.
type Params struct {
	// min_fee_update_delay is the minimum number of blocks after which a fee
	// update comes into effect.
	MinFeeUpdateDelay uint32 `protobuf:"varint,1,opt,name=min_fee_update_delay,json=minFeeUpdateDelay,proto3" json:"min_fee_update_delay,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce367ef351b38f5e, []int{0}
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

func (m *Params) GetMinFeeUpdateDelay() uint32 {
	if m != nil {
		return m.MinFeeUpdateDelay
	}
	return 0
}

// ProxyConfig defines a data-proxy entry in the registry.
type ProxyConfig struct {
	// payout_address defines the address to which the data proxy fees should be
	// transferred.
	PayoutAddress string `protobuf:"bytes,1,opt,name=payout_address,json=payoutAddress,proto3" json:"payout_address,omitempty"`
	// fee defines the amount in aseda this data-proxy charges when utilised.
	Fee *types.Coin `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	// memo defines an optional string which is not used by the protocol.
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
	// only the admin address of a data proxy can submit config updates.
	AdminAddress string `protobuf:"bytes,4,opt,name=admin_address,json=adminAddress,proto3" json:"admin_address,omitempty"`
	// fee_update defines an upcoming fee change which will take effect at a
	// future height.
	FeeUpdate *FeeUpdate `protobuf:"bytes,5,opt,name=fee_update,json=feeUpdate,proto3" json:"fee_update,omitempty"`
}

func (m *ProxyConfig) Reset()         { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()    {}
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce367ef351b38f5e, []int{1}
}
func (m *ProxyConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProxyConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyConfig.Merge(m, src)
}
func (m *ProxyConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyConfig proto.InternalMessageInfo

func (m *ProxyConfig) GetPayoutAddress() string {
	if m != nil {
		return m.PayoutAddress
	}
	return ""
}

func (m *ProxyConfig) GetFee() *types.Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *ProxyConfig) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *ProxyConfig) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
	}
	return ""
}

func (m *ProxyConfig) GetFeeUpdate() *FeeUpdate {
	if m != nil {
		return m.FeeUpdate
	}
	return nil
}

// FeeUpdate defines a new fee amount and the height at which it will take
// effect.
type FeeUpdate struct {
	// new_fee defines the new fee for the data proxy.
	NewFee types.Coin `protobuf:"bytes,1,opt,name=new_fee,json=newFee,proto3" json:"new_fee"`
	// update_height defines the height after which the new fee comes into effect.
	UpdateHeight int64 `protobuf:"varint,2,opt,name=update_height,json=updateHeight,proto3" json:"update_height,omitempty"`
}

func (m *FeeUpdate) Reset()         { *m = FeeUpdate{} }
func (m *FeeUpdate) String() string { return proto.CompactTextString(m) }
func (*FeeUpdate) ProtoMessage()    {}
func (*FeeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce367ef351b38f5e, []int{2}
}
func (m *FeeUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeUpdate.Merge(m, src)
}
func (m *FeeUpdate) XXX_Size() int {
	return m.Size()
}
func (m *FeeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_FeeUpdate proto.InternalMessageInfo

func (m *FeeUpdate) GetNewFee() types.Coin {
	if m != nil {
		return m.NewFee
	}
	return types.Coin{}
}

func (m *FeeUpdate) GetUpdateHeight() int64 {
	if m != nil {
		return m.UpdateHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "sedachain.data_proxy.v1.Params")
	proto.RegisterType((*ProxyConfig)(nil), "sedachain.data_proxy.v1.ProxyConfig")
	proto.RegisterType((*FeeUpdate)(nil), "sedachain.data_proxy.v1.FeeUpdate")
}

func init() {
	proto.RegisterFile("sedachain/data_proxy/v1/data_proxy.proto", fileDescriptor_ce367ef351b38f5e)
}

var fileDescriptor_ce367ef351b38f5e = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x92, 0x10, 0x94, 0x6d, 0x83, 0xc4, 0x2a, 0x12, 0x6e, 0x0f, 0x6e, 0x15, 0x2e, 0x91,
	0x50, 0x76, 0x15, 0x10, 0x12, 0xe2, 0x52, 0x91, 0xa2, 0xc2, 0x8d, 0xca, 0x88, 0x0b, 0x17, 0x6b,
	0x63, 0x8f, 0xed, 0x45, 0xf5, 0xae, 0xe5, 0xdd, 0xa4, 0xf5, 0x5f, 0xf0, 0x09, 0x7c, 0x04, 0x1f,
	0xd1, 0x63, 0xc5, 0x89, 0x13, 0x42, 0xc9, 0x85, 0x9f, 0x40, 0x42, 0xbb, 0xeb, 0x84, 0x5e, 0xe8,
	0x6d, 0xe7, 0xbd, 0xe7, 0x99, 0xf7, 0xc6, 0x83, 0x27, 0x1a, 0x52, 0x9e, 0x14, 0x5c, 0x48, 0x96,
	0x72, 0xc3, 0xe3, 0xaa, 0x56, 0x57, 0x0d, 0x5b, 0xcd, 0x6e, 0x55, 0xb4, 0xaa, 0x95, 0x51, 0xe4,
	0xf1, 0x4e, 0x49, 0x6f, 0x71, 0xab, 0xd9, 0xe1, 0x28, 0x57, 0xb9, 0x72, 0x1a, 0x66, 0x5f, 0x5e,
	0x7e, 0x78, 0x90, 0x28, 0x5d, 0x2a, 0x1d, 0x7b, 0xc2, 0x17, 0x2d, 0x15, 0xfa, 0x8a, 0x2d, 0xb8,
	0x06, 0xb6, 0x9a, 0x2d, 0xc0, 0xf0, 0x19, 0x4b, 0x94, 0x90, 0x9e, 0x1f, 0x9f, 0xe0, 0xfe, 0x39,
	0xaf, 0x79, 0xa9, 0x09, 0xc3, 0xa3, 0x52, 0xc8, 0x38, 0x03, 0x88, 0x97, 0x55, 0xca, 0x0d, 0xc4,
	0x29, 0x5c, 0xf0, 0x26, 0x40, 0xc7, 0x68, 0x32, 0x8c, 0x1e, 0x95, 0x42, 0x9e, 0x01, 0x7c, 0x74,
	0xcc, 0x1b, 0x4b, 0xbc, 0xea, 0xfd, 0xfe, 0x7a, 0x84, 0xc6, 0x7f, 0x10, 0xde, 0x3b, 0xb7, 0xf6,
	0x4e, 0x95, 0xcc, 0x44, 0x4e, 0x4e, 0xf0, 0xc3, 0x8a, 0x37, 0x6a, 0x69, 0x62, 0x9e, 0xa6, 0x35,
	0x68, 0xed, 0x1a, 0x0c, 0xe6, 0xc1, 0xf7, 0x6f, 0xd3, 0x51, 0x6b, 0xed, 0xb5, 0x67, 0x3e, 0x98,
	0x5a, 0xc8, 0x3c, 0x1a, 0x7a, 0x7d, 0x0b, 0x92, 0xa7, 0xb8, 0x9b, 0x01, 0x04, 0xf7, 0x8e, 0xd1,
	0x64, 0xef, 0xd9, 0x01, 0x6d, 0x3f, 0xb1, 0xfe, 0x69, 0xeb, 0x9f, 0x9e, 0x2a, 0x21, 0x23, 0xab,
	0x22, 0x04, 0xf7, 0x4a, 0x28, 0x55, 0xd0, 0xb5, 0x33, 0x22, 0xf7, 0x26, 0x4f, 0xf0, 0x90, 0xa7,
	0x36, 0xca, 0xd6, 0x40, 0xcf, 0x91, 0xfb, 0x0e, 0xdc, 0x4e, 0x79, 0x8b, 0xf1, 0xbf, 0xa4, 0xc1,
	0x7d, 0x37, 0x6c, 0x4c, 0xff, 0xb3, 0x76, 0xba, 0x4b, 0x3e, 0xef, 0x5d, 0xff, 0x3c, 0x42, 0xd1,
	0x20, 0xdb, 0x02, 0xe3, 0xcf, 0x78, 0xb0, 0x63, 0xc9, 0x4b, 0xfc, 0x40, 0xc2, 0xa5, 0xdd, 0xa1,
	0x4b, 0x7d, 0x97, 0x7f, 0xd7, 0xa9, 0x13, 0xf5, 0x25, 0x5c, 0x9e, 0x01, 0x58, 0xd3, 0xed, 0xd6,
	0x0b, 0x10, 0x79, 0x61, 0x5c, 0xfe, 0x6e, 0xb4, 0xef, 0xc1, 0x77, 0x0e, 0x9b, 0xbf, 0xbf, 0x5e,
	0x87, 0xe8, 0x66, 0x1d, 0xa2, 0x5f, 0xeb, 0x10, 0x7d, 0xd9, 0x84, 0x9d, 0x9b, 0x4d, 0xd8, 0xf9,
	0xb1, 0x09, 0x3b, 0x9f, 0x5e, 0xe4, 0xc2, 0x14, 0xcb, 0x05, 0x4d, 0x54, 0xc9, 0x6c, 0x08, 0xf7,
	0x73, 0x13, 0x75, 0xe1, 0x8a, 0xa9, 0xbf, 0xb9, 0x2b, 0x77, 0x67, 0x53, 0x7f, 0x75, 0xa6, 0xa9,
	0x40, 0x2f, 0xfa, 0x4e, 0xf7, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xe3, 0x70, 0x28,
	0x9a, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MinFeeUpdateDelay != that1.MinFeeUpdateDelay {
		return false
	}
	return true
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
	if m.MinFeeUpdateDelay != 0 {
		i = encodeVarintDataProxy(dAtA, i, uint64(m.MinFeeUpdateDelay))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProxyConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProxyConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProxyConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeUpdate != nil {
		{
			size, err := m.FeeUpdate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDataProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AdminAddress) > 0 {
		i -= len(m.AdminAddress)
		copy(dAtA[i:], m.AdminAddress)
		i = encodeVarintDataProxy(dAtA, i, uint64(len(m.AdminAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintDataProxy(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDataProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.PayoutAddress) > 0 {
		i -= len(m.PayoutAddress)
		copy(dAtA[i:], m.PayoutAddress)
		i = encodeVarintDataProxy(dAtA, i, uint64(len(m.PayoutAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FeeUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdateHeight != 0 {
		i = encodeVarintDataProxy(dAtA, i, uint64(m.UpdateHeight))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.NewFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDataProxy(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDataProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinFeeUpdateDelay != 0 {
		n += 1 + sovDataProxy(uint64(m.MinFeeUpdateDelay))
	}
	return n
}

func (m *ProxyConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PayoutAddress)
	if l > 0 {
		n += 1 + l + sovDataProxy(uint64(l))
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovDataProxy(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovDataProxy(uint64(l))
	}
	l = len(m.AdminAddress)
	if l > 0 {
		n += 1 + l + sovDataProxy(uint64(l))
	}
	if m.FeeUpdate != nil {
		l = m.FeeUpdate.Size()
		n += 1 + l + sovDataProxy(uint64(l))
	}
	return n
}

func (m *FeeUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NewFee.Size()
	n += 1 + l + sovDataProxy(uint64(l))
	if m.UpdateHeight != 0 {
		n += 1 + sovDataProxy(uint64(m.UpdateHeight))
	}
	return n
}

func sovDataProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataProxy(x uint64) (n int) {
	return sovDataProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataProxy
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinFeeUpdateDelay", wireType)
			}
			m.MinFeeUpdateDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinFeeUpdateDelay |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDataProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataProxy
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
func (m *ProxyConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataProxy
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
			return fmt.Errorf("proto: ProxyConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProxyConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayoutAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataProxy
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
				return ErrInvalidLengthDataProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayoutAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataProxy
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
				return ErrInvalidLengthDataProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &types.Coin{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataProxy
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
				return ErrInvalidLengthDataProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataProxy
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
				return ErrInvalidLengthDataProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataProxy
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
				return ErrInvalidLengthDataProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeeUpdate == nil {
				m.FeeUpdate = &FeeUpdate{}
			}
			if err := m.FeeUpdate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataProxy
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
func (m *FeeUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataProxy
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
			return fmt.Errorf("proto: FeeUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataProxy
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
				return ErrInvalidLengthDataProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NewFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateHeight", wireType)
			}
			m.UpdateHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDataProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataProxy
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
func skipDataProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataProxy
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
					return 0, ErrIntOverflowDataProxy
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
					return 0, ErrIntOverflowDataProxy
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
				return 0, ErrInvalidLengthDataProxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataProxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataProxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataProxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataProxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataProxy = fmt.Errorf("proto: unexpected end of group")
)
