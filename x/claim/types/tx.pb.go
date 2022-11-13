// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/claim/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgClaim defines a SDK message for claiming claimable amount.
type MsgClaim struct {
	// airdrop_id specifies index of the airdrop
	AirdropId uint64 `protobuf:"varint,1,opt,name=airdrop_id,json=airdropId,proto3" json:"airdrop_id,omitempty"`
	// recipient specifies the bech32-encoded address that is eligible to claim airdrop
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// condition_type specifies the condition type
	ConditionType ConditionType `protobuf:"varint,3,opt,name=condition_type,json=conditionType,proto3,enum=crescent.claim.v1beta1.ConditionType" json:"condition_type,omitempty"`
}

func (m *MsgClaim) Reset()         { *m = MsgClaim{} }
func (m *MsgClaim) String() string { return proto.CompactTextString(m) }
func (*MsgClaim) ProtoMessage()    {}
func (*MsgClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcf8d68258dc7e31, []int{0}
}
func (m *MsgClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaim.Merge(m, src)
}
func (m *MsgClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaim proto.InternalMessageInfo

type MsgClaimResponse struct {
}

func (m *MsgClaimResponse) Reset()         { *m = MsgClaimResponse{} }
func (m *MsgClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimResponse) ProtoMessage()    {}
func (*MsgClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcf8d68258dc7e31, []int{1}
}
func (m *MsgClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimResponse.Merge(m, src)
}
func (m *MsgClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgClaim)(nil), "crescent.claim.v1beta1.MsgClaim")
	proto.RegisterType((*MsgClaimResponse)(nil), "crescent.claim.v1beta1.MsgClaimResponse")
}

func init() { proto.RegisterFile("crescent/claim/v1beta1/tx.proto", fileDescriptor_fcf8d68258dc7e31) }

var fileDescriptor_fcf8d68258dc7e31 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x27, 0x5f, 0x3f, 0xc5, 0x06, 0x2c, 0x12, 0x44, 0x4a, 0xd1, 0x38, 0x14, 0x84, 0xd9,
	0x98, 0xd0, 0x16, 0x6f, 0xc0, 0xae, 0x04, 0xbb, 0x99, 0xba, 0xea, 0xa6, 0xb4, 0x99, 0x30, 0x06,
	0x6d, 0x12, 0x92, 0x58, 0xdb, 0xbb, 0x70, 0xe3, 0x3d, 0x75, 0xd9, 0xa5, 0x4b, 0x9d, 0xb9, 0x11,
	0x99, 0x9f, 0xa8, 0x0b, 0x8b, 0xbb, 0xe4, 0xe5, 0x39, 0x0f, 0xef, 0x39, 0xf0, 0x9c, 0x19, 0x6e,
	0x19, 0x97, 0x8e, 0xb2, 0xc7, 0x99, 0x58, 0xd0, 0x65, 0x6f, 0xce, 0xdd, 0xac, 0x47, 0xdd, 0x8a,
	0x68, 0xa3, 0x9c, 0x42, 0x27, 0x1e, 0x20, 0x25, 0x40, 0x6a, 0xa0, 0x73, 0x9c, 0xaa, 0x54, 0x95,
	0x08, 0x2d, 0x5e, 0x15, 0xdd, 0xe9, 0xee, 0xd0, 0x55, 0xb3, 0x25, 0xd3, 0x7d, 0x05, 0xf0, 0x60,
	0x64, 0xd3, 0x61, 0x11, 0xa1, 0x33, 0x08, 0x67, 0xc2, 0x24, 0x46, 0xe9, 0xa9, 0x48, 0xda, 0x20,
	0x04, 0xd1, 0xff, 0xb8, 0x59, 0x27, 0x37, 0x09, 0x3a, 0x85, 0x4d, 0xc3, 0x99, 0xd0, 0x82, 0x4b,
	0xd7, 0xfe, 0x17, 0x82, 0xa8, 0x19, 0x7f, 0x07, 0xe8, 0x16, 0xb6, 0x98, 0x92, 0x89, 0x70, 0x42,
	0xc9, 0xa9, 0x5b, 0x6b, 0xde, 0x6e, 0x84, 0x20, 0x6a, 0xf5, 0x2f, 0xc8, 0xef, 0xa5, 0xc9, 0xd0,
	0xd3, 0x77, 0x6b, 0xcd, 0xe3, 0x43, 0xf6, 0xf3, 0xdb, 0x45, 0xf0, 0xc8, 0xd7, 0x8a, 0xb9, 0xd5,
	0x4a, 0x5a, 0xde, 0x9f, 0xc0, 0xc6, 0xc8, 0xa6, 0x68, 0x0c, 0xf7, 0xaa, 0xba, 0xe1, 0x2e, 0xb3,
	0x9f, 0xec, 0x44, 0x7f, 0x11, 0xde, 0x7d, 0x3d, 0xde, 0x7c, 0xe0, 0x60, 0x93, 0x61, 0xb0, 0xcd,
	0x30, 0x78, 0xcf, 0x30, 0x78, 0xc9, 0x71, 0xb0, 0xcd, 0x71, 0xf0, 0x96, 0xe3, 0x60, 0x72, 0x95,
	0x0a, 0x77, 0xff, 0x34, 0x27, 0x4c, 0x2d, 0xa8, 0x37, 0x5e, 0x4a, 0xee, 0x9e, 0x95, 0x79, 0xf8,
	0x0a, 0xe8, 0x72, 0x40, 0x57, 0xf5, 0xa9, 0x8b, 0xfd, 0xed, 0x7c, 0xbf, 0xbc, 0xf1, 0xe0, 0x33,
	0x00, 0x00, 0xff, 0xff, 0x36, 0xe2, 0xf3, 0x25, 0xd8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Claim(ctx context.Context, in *MsgClaim, opts ...grpc.CallOption) (*MsgClaimResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Claim(ctx context.Context, in *MsgClaim, opts ...grpc.CallOption) (*MsgClaimResponse, error) {
	out := new(MsgClaimResponse)
	err := c.cc.Invoke(ctx, "/crescent.claim.v1beta1.Msg/Claim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Claim(context.Context, *MsgClaim) (*MsgClaimResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Claim(ctx context.Context, req *MsgClaim) (*MsgClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claim not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crescent.claim.v1beta1.Msg/Claim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Claim(ctx, req.(*MsgClaim))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crescent.claim.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Claim",
			Handler:    _Msg_Claim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crescent/claim/v1beta1/tx.proto",
}

func (m *MsgClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConditionType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ConditionType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if m.AirdropId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AirdropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AirdropId != 0 {
		n += 1 + sovTx(uint64(m.AirdropId))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ConditionType != 0 {
		n += 1 + sovTx(uint64(m.ConditionType))
	}
	return n
}

func (m *MsgClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropId", wireType)
			}
			m.AirdropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AirdropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConditionType", wireType)
			}
			m.ConditionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConditionType |= ConditionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
