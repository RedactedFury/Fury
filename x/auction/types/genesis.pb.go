// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/auction/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type GenesisState struct {
	SurplusAuction     []SurplusAuction     `protobuf:"bytes,1,rep,name=surplusAuction,proto3" json:"surplusAuction" yaml:"surplusAuction"`
	DebtAuction        []DebtAuction        `protobuf:"bytes,2,rep,name=debtAuction,proto3" json:"debtAuction" yaml:"debtAuction"`
	DutchAuction       []DutchAuction       `protobuf:"bytes,3,rep,name=dutchAuction,proto3" json:"dutchAuction" yaml:"dutchAuction"`
	ProtocolStatistics []ProtocolStatistics `protobuf:"bytes,4,rep,name=protocolStatistics,proto3" json:"protocolStatistics" yaml:"protocolStatistics"`
	AuctionParams      []AuctionParams      `protobuf:"bytes,5,rep,name=auctionParams,proto3" json:"auctionParams" yaml:"auctionParams"`
	DutchLendAuction   []DutchAuction       `protobuf:"bytes,6,rep,name=dutchLendAuction,proto3" json:"dutchLendAuction" yaml:"dutchLendAuction"`
	Params             Params               `protobuf:"bytes,7,opt,name=params,proto3" json:"params"`
	UserBiddingID      uint64               `protobuf:"varint,8,opt,name=UserBiddingID,proto3" json:"UserBiddingID,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_49088f171dd3086d, []int{0}
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

func (m *GenesisState) GetSurplusAuction() []SurplusAuction {
	if m != nil {
		return m.SurplusAuction
	}
	return nil
}

func (m *GenesisState) GetDebtAuction() []DebtAuction {
	if m != nil {
		return m.DebtAuction
	}
	return nil
}

func (m *GenesisState) GetDutchAuction() []DutchAuction {
	if m != nil {
		return m.DutchAuction
	}
	return nil
}

func (m *GenesisState) GetProtocolStatistics() []ProtocolStatistics {
	if m != nil {
		return m.ProtocolStatistics
	}
	return nil
}

func (m *GenesisState) GetAuctionParams() []AuctionParams {
	if m != nil {
		return m.AuctionParams
	}
	return nil
}

func (m *GenesisState) GetDutchLendAuction() []DutchAuction {
	if m != nil {
		return m.DutchLendAuction
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetUserBiddingID() uint64 {
	if m != nil {
		return m.UserBiddingID
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fury.auction.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("fury/auction/v1beta1/genesis.proto", fileDescriptor_49088f171dd3086d)
}

var fileDescriptor_49088f171dd3086d = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x86, 0x1b, 0x56, 0x0a, 0x72, 0x37, 0x84, 0xcc, 0x80, 0x50, 0x20, 0x2d, 0x5e, 0x87, 0x2a,
	0x24, 0x12, 0x75, 0xdc, 0x10, 0x97, 0x45, 0x93, 0x10, 0x82, 0xc3, 0xe4, 0x89, 0x0b, 0x37, 0x27,
	0xf1, 0x32, 0x4b, 0x4d, 0x1c, 0x62, 0x07, 0xb1, 0x03, 0x77, 0x8e, 0x7c, 0xac, 0x1d, 0x77, 0xe4,
	0x34, 0xa1, 0xf6, 0x1b, 0xf0, 0x09, 0x50, 0x6c, 0x47, 0x4b, 0x68, 0x7d, 0xd8, 0x2d, 0x7f, 0x5e,
	0x3f, 0x4f, 0x7e, 0x6f, 0x6c, 0x30, 0x8d, 0x79, 0x96, 0xd0, 0xef, 0x01, 0xa9, 0x62, 0xc9, 0x78,
	0x1e, 0x7c, 0x9b, 0x47, 0x54, 0x92, 0x79, 0x90, 0xd2, 0x9c, 0x0a, 0x26, 0xfc, 0xa2, 0xe4, 0x92,
	0xc3, 0x47, 0x3a, 0xe5, 0x9b, 0x94, 0x6f, 0x52, 0xa3, 0xdd, 0x94, 0xa7, 0x5c, 0x45, 0x82, 0xfa,
	0x4a, 0xa7, 0x47, 0x7b, 0x16, 0x66, 0x41, 0x4a, 0x92, 0x19, 0xe4, 0xc8, 0x26, 0x6e, 0x14, 0x3a,
	0xb5, 0x6f, 0x49, 0x45, 0x2c, 0x49, 0x58, 0x9e, 0x1a, 0x18, 0xfa, 0x39, 0x00, 0xdb, 0xef, 0xf5,
	0x17, 0x9f, 0x48, 0x22, 0x29, 0xcc, 0xc0, 0x3d, 0x51, 0x95, 0xc5, 0xa2, 0x12, 0x87, 0x7a, 0xa5,
	0xeb, 0x4c, 0xb6, 0x66, 0xc3, 0x83, 0x97, 0xfe, 0xe6, 0x49, 0xfc, 0x93, 0x4e, 0x3a, 0x7c, 0x7e,
	0x71, 0x35, 0xee, 0xfd, 0xbd, 0x1a, 0x3f, 0x3c, 0x27, 0xd9, 0xe2, 0x2d, 0xea, 0xb2, 0x10, 0xfe,
	0x0f, 0x0e, 0x09, 0x18, 0x26, 0x34, 0x92, 0x8d, 0xeb, 0x96, 0x72, 0xed, 0xd9, 0x5c, 0x47, 0xd7,
	0xd1, 0x70, 0x64, 0x44, 0x50, 0x8b, 0x5a, 0x14, 0x84, 0xdb, 0x4c, 0x48, 0xc1, 0x76, 0x52, 0xc9,
	0xf8, 0xac, 0x71, 0x6c, 0x29, 0xc7, 0xd4, 0xea, 0x68, 0x65, 0xc3, 0xa7, 0x46, 0xf2, 0xc0, 0x48,
	0x5a, 0xef, 0x10, 0xee, 0x60, 0xe1, 0x0f, 0x00, 0x55, 0xa5, 0x31, 0x5f, 0xd4, 0x4d, 0x32, 0x21,
	0x59, 0x2c, 0xdc, 0xbe, 0x92, 0xbd, 0xb2, 0xc9, 0x8e, 0xd7, 0x56, 0x84, 0x2f, 0x8c, 0xf2, 0x89,
	0x56, 0xae, 0x33, 0x11, 0xde, 0x20, 0x82, 0x0c, 0xec, 0x18, 0xf8, 0xb1, 0xda, 0x2c, 0xee, 0x6d,
	0x65, 0xde, 0xb7, 0x99, 0x0f, 0xdb, 0xe1, 0xf0, 0x99, 0x91, 0xee, 0x6a, 0x69, 0x87, 0x84, 0x70,
	0x97, 0x0c, 0xbf, 0x82, 0xfb, 0x6a, 0xf2, 0x4f, 0x34, 0x4f, 0x9a, 0x52, 0x07, 0x37, 0x28, 0x75,
	0x6c, 0x64, 0x8f, 0x5b, 0xa5, 0xb6, 0x58, 0x08, 0xaf, 0xe1, 0xe1, 0x3b, 0x30, 0xd0, 0x67, 0xc0,
	0xbd, 0x33, 0x71, 0x66, 0xc3, 0x03, 0xcf, 0x5a, 0xa8, 0x9e, 0xa7, 0x5f, 0x2b, 0xb0, 0x59, 0x03,
	0xa7, 0x60, 0xe7, 0xb3, 0xa0, 0x65, 0xa8, 0xb7, 0xfe, 0x87, 0x23, 0xf7, 0xee, 0xc4, 0x99, 0xf5,
	0x71, 0xf7, 0x61, 0xf8, 0xf1, 0x62, 0xe9, 0x39, 0x97, 0x4b, 0xcf, 0xf9, 0xb3, 0xf4, 0x9c, 0x5f,
	0x2b, 0xaf, 0x77, 0xb9, 0xf2, 0x7a, 0xbf, 0x57, 0x5e, 0xef, 0xcb, 0x3c, 0x65, 0xf2, 0xac, 0x8a,
	0x6a, 0x67, 0xa0, 0xbd, 0xaf, 0xf9, 0xe9, 0x29, 0x8b, 0x19, 0x59, 0x98, 0xfb, 0xe0, 0xfa, 0xa0,
	0xc9, 0xf3, 0x82, 0x8a, 0x68, 0xa0, 0x7e, 0xd1, 0x9b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf3,
	0xf6, 0xde, 0xc2, 0x26, 0x04, 0x00, 0x00,
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
	if m.UserBiddingID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.UserBiddingID))
		i--
		dAtA[i] = 0x40
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
	dAtA[i] = 0x3a
	if len(m.DutchLendAuction) > 0 {
		for iNdEx := len(m.DutchLendAuction) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DutchLendAuction[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.AuctionParams) > 0 {
		for iNdEx := len(m.AuctionParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AuctionParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ProtocolStatistics) > 0 {
		for iNdEx := len(m.ProtocolStatistics) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProtocolStatistics[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DutchAuction) > 0 {
		for iNdEx := len(m.DutchAuction) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DutchAuction[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DebtAuction) > 0 {
		for iNdEx := len(m.DebtAuction) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DebtAuction[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.SurplusAuction) > 0 {
		for iNdEx := len(m.SurplusAuction) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SurplusAuction[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.SurplusAuction) > 0 {
		for _, e := range m.SurplusAuction {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DebtAuction) > 0 {
		for _, e := range m.DebtAuction {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DutchAuction) > 0 {
		for _, e := range m.DutchAuction {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProtocolStatistics) > 0 {
		for _, e := range m.ProtocolStatistics {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AuctionParams) > 0 {
		for _, e := range m.AuctionParams {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DutchLendAuction) > 0 {
		for _, e := range m.DutchLendAuction {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.UserBiddingID != 0 {
		n += 1 + sovGenesis(uint64(m.UserBiddingID))
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
				return fmt.Errorf("proto: wrong wireType = %d for field SurplusAuction", wireType)
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
			m.SurplusAuction = append(m.SurplusAuction, SurplusAuction{})
			if err := m.SurplusAuction[len(m.SurplusAuction)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtAuction", wireType)
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
			m.DebtAuction = append(m.DebtAuction, DebtAuction{})
			if err := m.DebtAuction[len(m.DebtAuction)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutchAuction", wireType)
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
			m.DutchAuction = append(m.DutchAuction, DutchAuction{})
			if err := m.DutchAuction[len(m.DutchAuction)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolStatistics", wireType)
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
			m.ProtocolStatistics = append(m.ProtocolStatistics, ProtocolStatistics{})
			if err := m.ProtocolStatistics[len(m.ProtocolStatistics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionParams", wireType)
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
			m.AuctionParams = append(m.AuctionParams, AuctionParams{})
			if err := m.AuctionParams[len(m.AuctionParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutchLendAuction", wireType)
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
			m.DutchLendAuction = append(m.DutchLendAuction, DutchAuction{})
			if err := m.DutchLendAuction[len(m.DutchLendAuction)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
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
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserBiddingID", wireType)
			}
			m.UserBiddingID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserBiddingID |= uint64(b&0x7F) << shift
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