// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: intertx/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
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

// MsgSend defines a SDK message for creating a new identifer.
type MsgRegisterAccount struct {
	// the port by which the packet will be sent
	SourcePort string `protobuf:"bytes,2,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,3,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,4,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,5,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	Owner            string `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *MsgRegisterAccount) Reset()         { *m = MsgRegisterAccount{} }
func (m *MsgRegisterAccount) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterAccount) ProtoMessage()    {}
func (*MsgRegisterAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf629f09c7954014, []int{0}
}
func (m *MsgRegisterAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterAccount.Merge(m, src)
}
func (m *MsgRegisterAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterAccount proto.InternalMessageInfo

type MsgRegisterAccountResponse struct {
}

func (m *MsgRegisterAccountResponse) Reset()         { *m = MsgRegisterAccountResponse{} }
func (m *MsgRegisterAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterAccountResponse) ProtoMessage()    {}
func (*MsgRegisterAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf629f09c7954014, []int{1}
}
func (m *MsgRegisterAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterAccountResponse.Merge(m, src)
}
func (m *MsgRegisterAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterAccount)(nil), "intertx.MsgRegisterAccount")
	proto.RegisterType((*MsgRegisterAccountResponse)(nil), "intertx.MsgRegisterAccountResponse")
}

func init() { proto.RegisterFile("intertx/tx.proto", fileDescriptor_cf629f09c7954014) }

var fileDescriptor_cf629f09c7954014 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0x93, 0xfe, 0xa3, 0xb8, 0x6a, 0x55, 0xac, 0x82, 0x42, 0x28, 0xc9, 0x29, 0x2c, 0xb7,
	0x60, 0xd3, 0x32, 0x20, 0x75, 0x82, 0xb0, 0x00, 0xa2, 0x12, 0x0a, 0x4c, 0x2c, 0x25, 0xb1, 0xac,
	0xc4, 0x52, 0x62, 0x47, 0xb6, 0x53, 0xae, 0xdf, 0x80, 0x91, 0x8f, 0x70, 0x1f, 0xe7, 0xc6, 0x1b,
	0x19, 0x50, 0x84, 0xee, 0x16, 0xe6, 0xfb, 0x04, 0x28, 0xb1, 0x03, 0x9c, 0x4e, 0xea, 0x94, 0xf7,
	0xfd, 0x3d, 0xcf, 0xf3, 0x46, 0x6f, 0xf2, 0x82, 0x63, 0xc6, 0x35, 0x95, 0x7a, 0x82, 0xf5, 0x04,
	0xd5, 0x52, 0x68, 0x01, 0xef, 0x58, 0xe2, 0x9f, 0xe4, 0x22, 0x17, 0x3d, 0xc3, 0x5d, 0x65, 0x64,
	0x3f, 0x64, 0x19, 0xc1, 0x44, 0x48, 0x8a, 0x49, 0xc9, 0x28, 0xd7, 0xf8, 0xfa, 0xcc, 0x56, 0xc6,
	0x10, 0xfd, 0xdc, 0x02, 0xf0, 0x52, 0xe5, 0x09, 0xcd, 0x99, 0xd2, 0x54, 0xbe, 0x22, 0x44, 0x34,
	0x5c, 0xc3, 0x17, 0xe0, 0x40, 0x89, 0x46, 0x12, 0x7a, 0x55, 0x0b, 0xa9, 0xbd, 0xad, 0x91, 0x3b,
	0xbe, 0x1b, 0x3f, 0x58, 0xb5, 0x21, 0xbc, 0x49, 0xab, 0xf2, 0x22, 0xfa, 0x4f, 0x8c, 0x12, 0x60,
	0xba, 0x0f, 0x42, 0x6a, 0xf8, 0x12, 0x1c, 0x59, 0x8d, 0x14, 0x29, 0xe7, 0xb4, 0xf4, 0xb6, 0xfb,
	0xec, 0xc3, 0x55, 0x1b, 0xde, 0x5f, 0xcb, 0x5a, 0x3d, 0x4a, 0x0e, 0x0d, 0x78, 0x6d, 0x7a, 0xf8,
	0x05, 0x1c, 0x69, 0x56, 0x51, 0xd1, 0xe8, 0xab, 0x82, 0xb2, 0xbc, 0xd0, 0xde, 0xce, 0xc8, 0x1d,
	0x1f, 0x9c, 0xfb, 0x88, 0x65, 0x04, 0x75, 0xbb, 0x20, 0xbb, 0xc1, 0xf5, 0x19, 0x7a, 0xd3, 0x3b,
	0xe2, 0xc7, 0xb3, 0x36, 0x74, 0xfe, 0xbd, 0x61, 0x3d, 0x1f, 0x25, 0x87, 0x16, 0x18, 0x37, 0x7c,
	0x0b, 0xee, 0x0d, 0x8e, 0xee, 0xa9, 0x74, 0x5a, 0xd5, 0xde, 0xee, 0xc8, 0x1d, 0xef, 0xc4, 0xa7,
	0xab, 0x36, 0xf4, 0xd6, 0x87, 0xfc, 0xb5, 0x44, 0xc9, 0xb1, 0x65, 0x9f, 0x06, 0x04, 0x4f, 0xc0,
	0xae, 0xf8, 0xca, 0xa9, 0xf4, 0xf6, 0xba, 0x2d, 0x13, 0xd3, 0x5c, 0xec, 0x7f, 0x9b, 0x86, 0xce,
	0xef, 0x69, 0xe8, 0x44, 0xa7, 0xc0, 0xdf, 0xfc, 0xba, 0x09, 0x55, 0xb5, 0xe0, 0x8a, 0x9e, 0x7f,
	0x04, 0xdb, 0x97, 0x2a, 0x87, 0xef, 0xc1, 0xfe, 0xe0, 0x80, 0x8f, 0x90, 0xfd, 0xa1, 0x68, 0x33,
	0xe7, 0x3f, 0xb9, 0x45, 0x1c, 0x86, 0xc6, 0xef, 0x66, 0x8b, 0xc0, 0x9d, 0x2f, 0x02, 0xf7, 0xd7,
	0x22, 0x70, 0xbf, 0x2f, 0x03, 0x67, 0xbe, 0x0c, 0x9c, 0x1f, 0xcb, 0xc0, 0xf9, 0xfc, 0x2c, 0x67,
	0xba, 0x68, 0x32, 0x44, 0x44, 0x85, 0xfb, 0x41, 0xa4, 0x48, 0x19, 0xcf, 0xa8, 0x2c, 0x19, 0xc7,
	0x8c, 0xa4, 0x78, 0x62, 0xf0, 0xd3, 0xee, 0xc0, 0x6e, 0x6a, 0xaa, 0xb2, 0xbd, 0xfe, 0x48, 0x9e,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x39, 0x6a, 0x3c, 0x78, 0x02, 0x00, 0x00,
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
	Register(ctx context.Context, in *MsgRegisterAccount, opts ...grpc.CallOption) (*MsgRegisterAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Register(ctx context.Context, in *MsgRegisterAccount, opts ...grpc.CallOption) (*MsgRegisterAccountResponse, error) {
	out := new(MsgRegisterAccountResponse)
	err := c.cc.Invoke(ctx, "/intertx.Msg/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Register(context.Context, *MsgRegisterAccount) (*MsgRegisterAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Register(ctx context.Context, req *MsgRegisterAccount) (*MsgRegisterAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intertx.Msg/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Register(ctx, req.(*MsgRegisterAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "intertx.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Msg_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "intertx/tx.proto",
}

func (m *MsgRegisterAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRegisterAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterAccountResponse) Size() (n int) {
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
func (m *MsgRegisterAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
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
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
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
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *MsgRegisterAccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
