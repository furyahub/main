// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/house/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// MsgDeposit defines a SDK message for performing a deposit of coins to become part of the house corresponding to a sport event.
type MsgDeposit struct {
	DepositorAddress string     `protobuf:"bytes,1,opt,name=depositor_address,json=depositorAddress,proto3" json:"depositor_address,omitempty" yaml:"depositor_address"`
	SportEventUID    string     `protobuf:"bytes,2,opt,name=sport_event_uid,json=sportEventUID,proto3" json:"sportEventUID"`
	Amount           types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgDeposit) Reset()         { *m = MsgDeposit{} }
func (m *MsgDeposit) String() string { return proto.CompactTextString(m) }
func (*MsgDeposit) ProtoMessage()    {}
func (*MsgDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3891d05e499977f, []int{0}
}
func (m *MsgDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeposit.Merge(m, src)
}
func (m *MsgDeposit) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeposit proto.InternalMessageInfo

// MsgDepositResponse defines the Msg/Deposit response type.
type MsgDepositResponse struct {
	SportEventUID string `protobuf:"bytes,1,opt,name=sport_event_uid,json=sportEventUID,proto3" json:"sportEventUID"`
	ParticipantID uint64 `protobuf:"varint,2,opt,name=participant_id,json=participantID,proto3" json:"participantID"`
}

func (m *MsgDepositResponse) Reset()         { *m = MsgDepositResponse{} }
func (m *MsgDepositResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDepositResponse) ProtoMessage()    {}
func (*MsgDepositResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3891d05e499977f, []int{1}
}
func (m *MsgDepositResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositResponse.Merge(m, src)
}
func (m *MsgDepositResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositResponse proto.InternalMessageInfo

func (m *MsgDepositResponse) GetSportEventUID() string {
	if m != nil {
		return m.SportEventUID
	}
	return ""
}

func (m *MsgDepositResponse) GetParticipantID() uint64 {
	if m != nil {
		return m.ParticipantID
	}
	return 0
}

// MsgWithdraw defines a SDK message for performing a withdrawal of coins of unused amount corresponding to a deposit.
type MsgWithdraw struct {
	DepositorAddress string         `protobuf:"bytes,1,opt,name=depositor_address,json=depositorAddress,proto3" json:"depositor_address,omitempty" yaml:"depositor_address"`
	SportEventUID    string         `protobuf:"bytes,2,opt,name=sport_event_uid,json=sportEventUID,proto3" json:"sportEventUID"`
	ParticipantID    uint64         `protobuf:"varint,3,opt,name=participant_id,json=participantID,proto3" json:"participantID"`
	Mode             WithdrawalMode `protobuf:"varint,4,opt,name=mode,proto3,enum=sgenetwork.sge.house.WithdrawalMode" json:"mode,omitempty" yaml:"mode"`
	Amount           types.Coin     `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgWithdraw) Reset()         { *m = MsgWithdraw{} }
func (m *MsgWithdraw) String() string { return proto.CompactTextString(m) }
func (*MsgWithdraw) ProtoMessage()    {}
func (*MsgWithdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3891d05e499977f, []int{2}
}
func (m *MsgWithdraw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdraw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdraw.Merge(m, src)
}
func (m *MsgWithdraw) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdraw.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdraw proto.InternalMessageInfo

// MsgWithdrawResponse defines the Msg/Withdraw response type.
type MsgWithdrawResponse struct {
	SportEventUID    string `protobuf:"bytes,1,opt,name=sport_event_uid,json=sportEventUID,proto3" json:"sportEventUID"`
	ParticipantID    uint64 `protobuf:"varint,2,opt,name=participant_id,json=participantID,proto3" json:"participantID"`
	WithdrawalNumber uint64 `protobuf:"varint,3,opt,name=withdrawal_number,json=withdrawalNumber,proto3" json:"withdrawal_number,omitempty" yaml:"withdrawal_number"`
}

func (m *MsgWithdrawResponse) Reset()         { *m = MsgWithdrawResponse{} }
func (m *MsgWithdrawResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawResponse) ProtoMessage()    {}
func (*MsgWithdrawResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3891d05e499977f, []int{3}
}
func (m *MsgWithdrawResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawResponse.Merge(m, src)
}
func (m *MsgWithdrawResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawResponse proto.InternalMessageInfo

func (m *MsgWithdrawResponse) GetSportEventUID() string {
	if m != nil {
		return m.SportEventUID
	}
	return ""
}

func (m *MsgWithdrawResponse) GetParticipantID() uint64 {
	if m != nil {
		return m.ParticipantID
	}
	return 0
}

func (m *MsgWithdrawResponse) GetWithdrawalNumber() uint64 {
	if m != nil {
		return m.WithdrawalNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgDeposit)(nil), "sgenetwork.sge.house.MsgDeposit")
	proto.RegisterType((*MsgDepositResponse)(nil), "sgenetwork.sge.house.MsgDepositResponse")
	proto.RegisterType((*MsgWithdraw)(nil), "sgenetwork.sge.house.MsgWithdraw")
	proto.RegisterType((*MsgWithdrawResponse)(nil), "sgenetwork.sge.house.MsgWithdrawResponse")
}

func init() { proto.RegisterFile("sge/house/tx.proto", fileDescriptor_d3891d05e499977f) }

var fileDescriptor_d3891d05e499977f = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x18, 0xf5, 0x35, 0xa1, 0x94, 0x8b, 0xfa, 0x03, 0x53, 0xa4, 0x10, 0x21, 0x3b, 0x58, 0x0c, 0x61,
	0xe0, 0xac, 0x86, 0x01, 0xa9, 0x1b, 0xa6, 0x0c, 0x41, 0x0a, 0x42, 0x46, 0x15, 0x88, 0xc5, 0xb2,
	0xe3, 0xd3, 0xc5, 0xa2, 0xf6, 0x59, 0xfe, 0xce, 0x4d, 0xfb, 0x1f, 0x30, 0xf2, 0x27, 0x74, 0x67,
	0xe7, 0x6f, 0xe8, 0xd8, 0x91, 0xc9, 0x42, 0x89, 0x84, 0x10, 0x63, 0x17, 0x26, 0x24, 0xe4, 0x5f,
	0xb1, 0xab, 0x04, 0x55, 0xca, 0xd4, 0xc5, 0xba, 0x7b, 0xdf, 0xbb, 0xa7, 0x7b, 0xdf, 0x3d, 0x7f,
	0x58, 0x06, 0x46, 0xf5, 0x31, 0x8f, 0x81, 0xea, 0xe2, 0x84, 0x84, 0x11, 0x17, 0x5c, 0xde, 0x05,
	0x46, 0x03, 0x2a, 0x26, 0x3c, 0xfa, 0x44, 0x80, 0x51, 0x92, 0x95, 0x3b, 0xbb, 0x8c, 0x33, 0x9e,
	0x11, 0xf4, 0x74, 0x95, 0x73, 0x3b, 0xca, 0x88, 0x83, 0xcf, 0x41, 0x77, 0x6c, 0xa0, 0xfa, 0xf1,
	0x9e, 0x43, 0x85, 0xbd, 0xa7, 0x8f, 0xb8, 0x17, 0x14, 0xf5, 0xfb, 0x95, 0x7e, 0xf6, 0xcd, 0x61,
	0xed, 0x27, 0xc2, 0x78, 0x08, 0xec, 0x80, 0x86, 0x1c, 0x3c, 0x21, 0x0f, 0xf0, 0x5d, 0x37, 0x5f,
	0xf2, 0xc8, 0xb2, 0x5d, 0x37, 0xa2, 0x00, 0x6d, 0xd4, 0x45, 0xbd, 0x3b, 0xc6, 0xc3, 0xcb, 0x44,
	0x6d, 0x9f, 0xda, 0xfe, 0xd1, 0xbe, 0xb6, 0x40, 0xd1, 0xcc, 0x9d, 0x39, 0xf6, 0x22, 0x87, 0xe4,
	0xd7, 0x78, 0x1b, 0x42, 0x1e, 0x09, 0x8b, 0x1e, 0xd3, 0x40, 0x58, 0xb1, 0xe7, 0xb6, 0xd7, 0x32,
	0x21, 0x6d, 0x9a, 0xa8, 0x9b, 0xef, 0xd2, 0xd2, 0xab, 0xb4, 0x72, 0x38, 0x38, 0xf8, 0x9d, 0xa8,
	0x9b, 0x50, 0x07, 0xcc, 0xab, 0x5b, 0xf9, 0x39, 0x5e, 0xb7, 0x7d, 0x1e, 0x07, 0xa2, 0xdd, 0xe8,
	0xa2, 0x5e, 0xab, 0xff, 0x80, 0xe4, 0x6e, 0x49, 0xea, 0x96, 0x14, 0x6e, 0xc9, 0x4b, 0xee, 0x05,
	0x46, 0xf3, 0x3c, 0x51, 0x25, 0xb3, 0xa0, 0xef, 0x6f, 0x7c, 0x3e, 0x53, 0xa5, 0x5f, 0x67, 0xaa,
	0xa4, 0x7d, 0x45, 0x58, 0xae, 0x8c, 0x9a, 0x14, 0x42, 0x1e, 0x00, 0x5d, 0x76, 0x4b, 0xb4, 0xea,
	0x2d, 0x07, 0x78, 0x2b, 0xb4, 0x23, 0xe1, 0x8d, 0xbc, 0xd0, 0x0e, 0x84, 0x55, 0x18, 0x6e, 0xe6,
	0x52, 0x6f, 0xab, 0x4a, 0x2e, 0x15, 0xd6, 0x01, 0xf3, 0xea, 0x56, 0xfb, 0xbb, 0x86, 0x5b, 0x43,
	0x60, 0xef, 0x3d, 0x31, 0x76, 0x23, 0x7b, 0x72, 0x53, 0xdf, 0x65, 0xd1, 0x71, 0x63, 0x45, 0xc7,
	0xf2, 0x00, 0x37, 0x7d, 0xee, 0xd2, 0x76, 0xb3, 0x8b, 0x7a, 0x5b, 0xfd, 0xc7, 0x64, 0x59, 0xf4,
	0x49, 0xd9, 0x0f, 0xfb, 0x68, 0xc8, 0x5d, 0x6a, 0x6c, 0x5f, 0x26, 0x6a, 0x2b, 0xb7, 0x9e, 0x9e,
	0xd5, 0xcc, 0x4c, 0xa2, 0x96, 0x96, 0x5b, 0xab, 0xa6, 0xe5, 0x0f, 0xc2, 0xf7, 0x6a, 0xfd, 0xbf,
	0xe1, 0x71, 0x49, 0xe3, 0x31, 0x99, 0xb7, 0xc6, 0x0a, 0x62, 0xdf, 0xa1, 0x51, 0xf1, 0x14, 0xb5,
	0x78, 0x2c, 0x50, 0x34, 0x73, 0xa7, 0xc2, 0xde, 0x64, 0x50, 0xff, 0x1b, 0xc2, 0x8d, 0x21, 0x30,
	0xf9, 0x10, 0xdf, 0x2e, 0x87, 0x42, 0x77, 0xf9, 0x63, 0x54, 0x7f, 0x53, 0xa7, 0x77, 0x1d, 0x63,
	0xde, 0xc0, 0x0f, 0x78, 0x63, 0x1e, 0xea, 0x47, 0xff, 0x3d, 0x55, 0x52, 0x3a, 0x4f, 0xae, 0xa5,
	0x94, 0xca, 0x86, 0x71, 0x3e, 0x55, 0xd0, 0xc5, 0x54, 0x41, 0x3f, 0xa6, 0x0a, 0xfa, 0x32, 0x53,
	0xa4, 0x8b, 0x99, 0x22, 0x7d, 0x9f, 0x29, 0xd2, 0xc7, 0x1e, 0xf3, 0xc4, 0x38, 0x76, 0xc8, 0x88,
	0xfb, 0x3a, 0x30, 0xfa, 0xb4, 0xd0, 0x4b, 0xd7, 0xfa, 0x49, 0x39, 0x73, 0x4f, 0x43, 0x0a, 0xce,
	0x7a, 0x36, 0x14, 0x9f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x17, 0x29, 0x62, 0x8d, 0x05,
	0x00, 0x00,
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
	// Deposit defines a method for performing a deposit of coins to become part of the house corresponding to a sport event.
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error)
	// Withdraw defines a method for performing a withdrawal of coins of unused amount corresponding to a deposit.
	Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, "/sgenetwork.sge.house.Msg/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error) {
	out := new(MsgWithdrawResponse)
	err := c.cc.Invoke(ctx, "/sgenetwork.sge.house.Msg/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Deposit defines a method for performing a deposit of coins to become part of the house corresponding to a sport event.
	Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error)
	// Withdraw defines a method for performing a withdrawal of coins of unused amount corresponding to a deposit.
	Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Deposit(ctx context.Context, req *MsgDeposit) (*MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (*UnimplementedMsgServer) Withdraw(ctx context.Context, req *MsgWithdraw) (*MsgWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgenetwork.sge.house.Msg/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdraw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgenetwork.sge.house.Msg/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Withdraw(ctx, req.(*MsgWithdraw))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sgenetwork.sge.house.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Msg_Withdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sge/house/tx.proto",
}

func (m *MsgDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.SportEventUID) > 0 {
		i -= len(m.SportEventUID)
		copy(dAtA[i:], m.SportEventUID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SportEventUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DepositorAddress) > 0 {
		i -= len(m.DepositorAddress)
		copy(dAtA[i:], m.DepositorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DepositorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDepositResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ParticipantID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ParticipantID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SportEventUID) > 0 {
		i -= len(m.SportEventUID)
		copy(dAtA[i:], m.SportEventUID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SportEventUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdraw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdraw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdraw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Mode != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x20
	}
	if m.ParticipantID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ParticipantID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.SportEventUID) > 0 {
		i -= len(m.SportEventUID)
		copy(dAtA[i:], m.SportEventUID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SportEventUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DepositorAddress) > 0 {
		i -= len(m.DepositorAddress)
		copy(dAtA[i:], m.DepositorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DepositorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithdrawalNumber != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.WithdrawalNumber))
		i--
		dAtA[i] = 0x18
	}
	if m.ParticipantID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ParticipantID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SportEventUID) > 0 {
		i -= len(m.SportEventUID)
		copy(dAtA[i:], m.SportEventUID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SportEventUID)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepositorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SportEventUID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgDepositResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SportEventUID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ParticipantID != 0 {
		n += 1 + sovTx(uint64(m.ParticipantID))
	}
	return n
}

func (m *MsgWithdraw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepositorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SportEventUID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ParticipantID != 0 {
		n += 1 + sovTx(uint64(m.ParticipantID))
	}
	if m.Mode != 0 {
		n += 1 + sovTx(uint64(m.Mode))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgWithdrawResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SportEventUID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ParticipantID != 0 {
		n += 1 + sovTx(uint64(m.ParticipantID))
	}
	if m.WithdrawalNumber != 0 {
		n += 1 + sovTx(uint64(m.WithdrawalNumber))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDeposit) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositorAddress", wireType)
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
			m.DepositorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SportEventUID", wireType)
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
			m.SportEventUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgDepositResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDepositResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SportEventUID", wireType)
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
			m.SportEventUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipantID", wireType)
			}
			m.ParticipantID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipantID |= uint64(b&0x7F) << shift
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
func (m *MsgWithdraw) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdraw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdraw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositorAddress", wireType)
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
			m.DepositorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SportEventUID", wireType)
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
			m.SportEventUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipantID", wireType)
			}
			m.ParticipantID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipantID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= WithdrawalMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgWithdrawResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SportEventUID", wireType)
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
			m.SportEventUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipantID", wireType)
			}
			m.ParticipantID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipantID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalNumber", wireType)
			}
			m.WithdrawalNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawalNumber |= uint64(b&0x7F) << shift
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
