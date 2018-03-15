// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamer.proto

/*
Package streamer is a generated protocol buffer package.

It is generated from these files:
	streamer.proto

It has these top-level messages:
	BTCTransaction
	AddSpOut
	ReqDeleteSpOut
	AddressToDelete
	WatchAddress
	Empty
	RawTx
	AddressToResync
	UsersData
	ReplyInfo
	CliRequest
	SrvReply
*/
package streamer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BTCTransaction struct {
	UserID        string                         `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	TxID          string                         `protobuf:"bytes,2,opt,name=txID" json:"txID,omitempty"`
	TxHash        string                         `protobuf:"bytes,3,opt,name=txHash" json:"txHash,omitempty"`
	TxOutScript   string                         `protobuf:"bytes,4,opt,name=txOutScript" json:"txOutScript,omitempty"`
	TxAddress     []string                       `protobuf:"bytes,5,rep,name=txAddress" json:"txAddress,omitempty"`
	TxStatus      int32                          `protobuf:"varint,6,opt,name=txStatus" json:"txStatus,omitempty"`
	TxOutAmount   int64                          `protobuf:"varint,7,opt,name=txOutAmount" json:"txOutAmount,omitempty"`
	BlockTime     int64                          `protobuf:"varint,8,opt,name=blockTime" json:"blockTime,omitempty"`
	BlockHeight   int64                          `protobuf:"varint,9,opt,name=blockHeight" json:"blockHeight,omitempty"`
	Confirmations int32                          `protobuf:"varint,10,opt,name=confirmations" json:"confirmations,omitempty"`
	TxFee         int64                          `protobuf:"varint,11,opt,name=txFee" json:"txFee,omitempty"`
	MempoolTime   int64                          `protobuf:"varint,12,opt,name=mempoolTime" json:"mempoolTime,omitempty"`
	TxInputs      []*BTCTransaction_AddresAmount `protobuf:"bytes,13,rep,name=txInputs" json:"txInputs,omitempty"`
	TxOutputs     []*BTCTransaction_AddresAmount `protobuf:"bytes,14,rep,name=TxOutputs" json:"TxOutputs,omitempty"`
}

func (m *BTCTransaction) Reset()                    { *m = BTCTransaction{} }
func (m *BTCTransaction) String() string            { return proto.CompactTextString(m) }
func (*BTCTransaction) ProtoMessage()               {}
func (*BTCTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BTCTransaction) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *BTCTransaction) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *BTCTransaction) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *BTCTransaction) GetTxOutScript() string {
	if m != nil {
		return m.TxOutScript
	}
	return ""
}

func (m *BTCTransaction) GetTxAddress() []string {
	if m != nil {
		return m.TxAddress
	}
	return nil
}

func (m *BTCTransaction) GetTxStatus() int32 {
	if m != nil {
		return m.TxStatus
	}
	return 0
}

func (m *BTCTransaction) GetTxOutAmount() int64 {
	if m != nil {
		return m.TxOutAmount
	}
	return 0
}

func (m *BTCTransaction) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *BTCTransaction) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *BTCTransaction) GetConfirmations() int32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *BTCTransaction) GetTxFee() int64 {
	if m != nil {
		return m.TxFee
	}
	return 0
}

func (m *BTCTransaction) GetMempoolTime() int64 {
	if m != nil {
		return m.MempoolTime
	}
	return 0
}

func (m *BTCTransaction) GetTxInputs() []*BTCTransaction_AddresAmount {
	if m != nil {
		return m.TxInputs
	}
	return nil
}

func (m *BTCTransaction) GetTxOutputs() []*BTCTransaction_AddresAmount {
	if m != nil {
		return m.TxOutputs
	}
	return nil
}

type BTCTransaction_AddresAmount struct {
	Address string `protobuf:"bytes,1,opt,name=Address" json:"Address,omitempty"`
	Amount  int64  `protobuf:"varint,2,opt,name=Amount" json:"Amount,omitempty"`
}

func (m *BTCTransaction_AddresAmount) Reset()                    { *m = BTCTransaction_AddresAmount{} }
func (m *BTCTransaction_AddresAmount) String() string            { return proto.CompactTextString(m) }
func (*BTCTransaction_AddresAmount) ProtoMessage()               {}
func (*BTCTransaction_AddresAmount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *BTCTransaction_AddresAmount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BTCTransaction_AddresAmount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type AddSpOut struct {
	TxID        string `protobuf:"bytes,1,opt,name=txID" json:"txID,omitempty"`
	TxOutID     int32  `protobuf:"varint,2,opt,name=txOutID" json:"txOutID,omitempty"`
	TxOutAmount int64  `protobuf:"varint,3,opt,name=txOutAmount" json:"txOutAmount,omitempty"`
	TxOutScript string `protobuf:"bytes,4,opt,name=txOutScript" json:"txOutScript,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	UserID      string `protobuf:"bytes,6,opt,name=userID" json:"userID,omitempty"`
	TxStatus    int32  `protobuf:"varint,7,opt,name=txStatus" json:"txStatus,omitempty"`
}

func (m *AddSpOut) Reset()                    { *m = AddSpOut{} }
func (m *AddSpOut) String() string            { return proto.CompactTextString(m) }
func (*AddSpOut) ProtoMessage()               {}
func (*AddSpOut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddSpOut) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *AddSpOut) GetTxOutID() int32 {
	if m != nil {
		return m.TxOutID
	}
	return 0
}

func (m *AddSpOut) GetTxOutAmount() int64 {
	if m != nil {
		return m.TxOutAmount
	}
	return 0
}

func (m *AddSpOut) GetTxOutScript() string {
	if m != nil {
		return m.TxOutScript
	}
	return ""
}

func (m *AddSpOut) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddSpOut) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AddSpOut) GetTxStatus() int32 {
	if m != nil {
		return m.TxStatus
	}
	return 0
}

type ReqDeleteSpOut struct {
	UserID  string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	TxID    string `protobuf:"bytes,2,opt,name=txID" json:"txID,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
}

func (m *ReqDeleteSpOut) Reset()                    { *m = ReqDeleteSpOut{} }
func (m *ReqDeleteSpOut) String() string            { return proto.CompactTextString(m) }
func (*ReqDeleteSpOut) ProtoMessage()               {}
func (*ReqDeleteSpOut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReqDeleteSpOut) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ReqDeleteSpOut) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

func (m *ReqDeleteSpOut) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AddressToDelete struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *AddressToDelete) Reset()                    { *m = AddressToDelete{} }
func (m *AddressToDelete) String() string            { return proto.CompactTextString(m) }
func (*AddressToDelete) ProtoMessage()               {}
func (*AddressToDelete) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddressToDelete) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type WatchAddress struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *WatchAddress) Reset()                    { *m = WatchAddress{} }
func (m *WatchAddress) String() string            { return proto.CompactTextString(m) }
func (*WatchAddress) ProtoMessage()               {}
func (*WatchAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WatchAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type RawTx struct {
	Transaction string `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
}

func (m *RawTx) Reset()                    { *m = RawTx{} }
func (m *RawTx) String() string            { return proto.CompactTextString(m) }
func (*RawTx) ProtoMessage()               {}
func (*RawTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RawTx) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

type AddressToResync struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *AddressToResync) Reset()                    { *m = AddressToResync{} }
func (m *AddressToResync) String() string            { return proto.CompactTextString(m) }
func (*AddressToResync) ProtoMessage()               {}
func (*AddressToResync) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddressToResync) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type UsersData struct {
	Map map[string]string `protobuf:"bytes,1,rep,name=map" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UsersData) Reset()                    { *m = UsersData{} }
func (m *UsersData) String() string            { return proto.CompactTextString(m) }
func (*UsersData) ProtoMessage()               {}
func (*UsersData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UsersData) GetMap() map[string]string {
	if m != nil {
		return m.Map
	}
	return nil
}

type ReplyInfo struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ReplyInfo) Reset()                    { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()               {}
func (*ReplyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReplyInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CliRequest struct {
	Request string `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *CliRequest) Reset()                    { *m = CliRequest{} }
func (m *CliRequest) String() string            { return proto.CompactTextString(m) }
func (*CliRequest) ProtoMessage()               {}
func (*CliRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CliRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

type SrvReply struct {
	Reply string `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
}

func (m *SrvReply) Reset()                    { *m = SrvReply{} }
func (m *SrvReply) String() string            { return proto.CompactTextString(m) }
func (*SrvReply) ProtoMessage()               {}
func (*SrvReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SrvReply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*BTCTransaction)(nil), "BTCTransaction")
	proto.RegisterType((*BTCTransaction_AddresAmount)(nil), "BTCTransaction.AddresAmount")
	proto.RegisterType((*AddSpOut)(nil), "AddSpOut")
	proto.RegisterType((*ReqDeleteSpOut)(nil), "ReqDeleteSpOut")
	proto.RegisterType((*AddressToDelete)(nil), "AddressToDelete")
	proto.RegisterType((*WatchAddress)(nil), "WatchAddress")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*RawTx)(nil), "RawTx")
	proto.RegisterType((*AddressToResync)(nil), "AddressToResync")
	proto.RegisterType((*UsersData)(nil), "UsersData")
	proto.RegisterType((*ReplyInfo)(nil), "ReplyInfo")
	proto.RegisterType((*CliRequest)(nil), "CliRequest")
	proto.RegisterType((*SrvReply)(nil), "SrvReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeCommuunications service

type NodeCommuunicationsClient interface {
	LotsOfReplies(ctx context.Context, in *CliRequest, opts ...grpc.CallOption) (NodeCommuunications_LotsOfRepliesClient, error)
	EventInitialAdd(ctx context.Context, in *UsersData, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventGetAllMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventResyncAddress(ctx context.Context, in *AddressToResync, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventSendRawTx(ctx context.Context, in *RawTx, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventAddNewAddress(ctx context.Context, in *WatchAddress, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventDeleteMempool(ctx context.Context, in *AddressToDelete, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventDeleteSpendableOut(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventDeleteSpendableOutClient, error)
	EventAddSpendableOut(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventAddSpendableOutClient, error)
	NewTx(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_NewTxClient, error)
}

type nodeCommuunicationsClient struct {
	cc *grpc.ClientConn
}

func NewNodeCommuunicationsClient(cc *grpc.ClientConn) NodeCommuunicationsClient {
	return &nodeCommuunicationsClient{cc}
}

func (c *nodeCommuunicationsClient) LotsOfReplies(ctx context.Context, in *CliRequest, opts ...grpc.CallOption) (NodeCommuunications_LotsOfRepliesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[0], c.cc, "/NodeCommuunications/LotsOfReplies", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsLotsOfRepliesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_LotsOfRepliesClient interface {
	Recv() (*SrvReply, error)
	grpc.ClientStream
}

type nodeCommuunicationsLotsOfRepliesClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsLotsOfRepliesClient) Recv() (*SrvReply, error) {
	m := new(SrvReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventInitialAdd(ctx context.Context, in *UsersData, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/NodeCommuunications/EventInitialAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventGetAllMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/NodeCommuunications/EventGetAllMempool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventResyncAddress(ctx context.Context, in *AddressToResync, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/NodeCommuunications/EventResyncAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventSendRawTx(ctx context.Context, in *RawTx, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/NodeCommuunications/EventSendRawTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventAddNewAddress(ctx context.Context, in *WatchAddress, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/NodeCommuunications/EventAddNewAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventDeleteMempool(ctx context.Context, in *AddressToDelete, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/NodeCommuunications/EventDeleteMempool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventDeleteSpendableOut(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventDeleteSpendableOutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[1], c.cc, "/NodeCommuunications/EventDeleteSpendableOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventDeleteSpendableOutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventDeleteSpendableOutClient interface {
	Recv() (*ReqDeleteSpOut, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventDeleteSpendableOutClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventDeleteSpendableOutClient) Recv() (*ReqDeleteSpOut, error) {
	m := new(ReqDeleteSpOut)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventAddSpendableOut(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventAddSpendableOutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[2], c.cc, "/NodeCommuunications/EventAddSpendableOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventAddSpendableOutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventAddSpendableOutClient interface {
	Recv() (*AddSpOut, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventAddSpendableOutClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventAddSpendableOutClient) Recv() (*AddSpOut, error) {
	m := new(AddSpOut)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) NewTx(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_NewTxClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[3], c.cc, "/NodeCommuunications/NewTx", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsNewTxClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_NewTxClient interface {
	Recv() (*BTCTransaction, error)
	grpc.ClientStream
}

type nodeCommuunicationsNewTxClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsNewTxClient) Recv() (*BTCTransaction, error) {
	m := new(BTCTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for NodeCommuunications service

type NodeCommuunicationsServer interface {
	LotsOfReplies(*CliRequest, NodeCommuunications_LotsOfRepliesServer) error
	EventInitialAdd(context.Context, *UsersData) (*ReplyInfo, error)
	EventGetAllMempool(context.Context, *Empty) (*ReplyInfo, error)
	EventResyncAddress(context.Context, *AddressToResync) (*ReplyInfo, error)
	EventSendRawTx(context.Context, *RawTx) (*ReplyInfo, error)
	EventAddNewAddress(context.Context, *WatchAddress) (*ReplyInfo, error)
	EventDeleteMempool(context.Context, *AddressToDelete) (*ReplyInfo, error)
	EventDeleteSpendableOut(*Empty, NodeCommuunications_EventDeleteSpendableOutServer) error
	EventAddSpendableOut(*Empty, NodeCommuunications_EventAddSpendableOutServer) error
	NewTx(*Empty, NodeCommuunications_NewTxServer) error
}

func RegisterNodeCommuunicationsServer(s *grpc.Server, srv NodeCommuunicationsServer) {
	s.RegisterService(&_NodeCommuunications_serviceDesc, srv)
}

func _NodeCommuunications_LotsOfReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CliRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).LotsOfReplies(m, &nodeCommuunicationsLotsOfRepliesServer{stream})
}

type NodeCommuunications_LotsOfRepliesServer interface {
	Send(*SrvReply) error
	grpc.ServerStream
}

type nodeCommuunicationsLotsOfRepliesServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsLotsOfRepliesServer) Send(m *SrvReply) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventInitialAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventInitialAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeCommuunications/EventInitialAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventInitialAdd(ctx, req.(*UsersData))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventGetAllMempool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventGetAllMempool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeCommuunications/EventGetAllMempool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventGetAllMempool(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventResyncAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressToResync)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventResyncAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeCommuunications/EventResyncAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventResyncAddress(ctx, req.(*AddressToResync))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventSendRawTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventSendRawTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeCommuunications/EventSendRawTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventSendRawTx(ctx, req.(*RawTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventAddNewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventAddNewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeCommuunications/EventAddNewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventAddNewAddress(ctx, req.(*WatchAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventDeleteMempool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressToDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventDeleteMempool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeCommuunications/EventDeleteMempool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventDeleteMempool(ctx, req.(*AddressToDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventDeleteSpendableOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventDeleteSpendableOut(m, &nodeCommuunicationsEventDeleteSpendableOutServer{stream})
}

type NodeCommuunications_EventDeleteSpendableOutServer interface {
	Send(*ReqDeleteSpOut) error
	grpc.ServerStream
}

type nodeCommuunicationsEventDeleteSpendableOutServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventDeleteSpendableOutServer) Send(m *ReqDeleteSpOut) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventAddSpendableOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventAddSpendableOut(m, &nodeCommuunicationsEventAddSpendableOutServer{stream})
}

type NodeCommuunications_EventAddSpendableOutServer interface {
	Send(*AddSpOut) error
	grpc.ServerStream
}

type nodeCommuunicationsEventAddSpendableOutServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventAddSpendableOutServer) Send(m *AddSpOut) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_NewTx_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).NewTx(m, &nodeCommuunicationsNewTxServer{stream})
}

type NodeCommuunications_NewTxServer interface {
	Send(*BTCTransaction) error
	grpc.ServerStream
}

type nodeCommuunicationsNewTxServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsNewTxServer) Send(m *BTCTransaction) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeCommuunications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NodeCommuunications",
	HandlerType: (*NodeCommuunicationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventInitialAdd",
			Handler:    _NodeCommuunications_EventInitialAdd_Handler,
		},
		{
			MethodName: "EventGetAllMempool",
			Handler:    _NodeCommuunications_EventGetAllMempool_Handler,
		},
		{
			MethodName: "EventResyncAddress",
			Handler:    _NodeCommuunications_EventResyncAddress_Handler,
		},
		{
			MethodName: "EventSendRawTx",
			Handler:    _NodeCommuunications_EventSendRawTx_Handler,
		},
		{
			MethodName: "EventAddNewAddress",
			Handler:    _NodeCommuunications_EventAddNewAddress_Handler,
		},
		{
			MethodName: "EventDeleteMempool",
			Handler:    _NodeCommuunications_EventDeleteMempool_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LotsOfReplies",
			Handler:       _NodeCommuunications_LotsOfReplies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventDeleteSpendableOut",
			Handler:       _NodeCommuunications_EventDeleteSpendableOut_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventAddSpendableOut",
			Handler:       _NodeCommuunications_EventAddSpendableOut_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NewTx",
			Handler:       _NodeCommuunications_NewTx_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "streamer.proto",
}

func init() { proto.RegisterFile("streamer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x6d, 0x4f, 0xe3, 0x46,
	0x10, 0xc6, 0x98, 0xbc, 0x78, 0x42, 0x02, 0x5a, 0x50, 0xbb, 0x8a, 0xf8, 0x10, 0x59, 0x80, 0xd2,
	0xa2, 0x5a, 0x2d, 0xad, 0x10, 0xe2, 0x53, 0x53, 0x42, 0x4b, 0xa4, 0x02, 0x92, 0x93, 0xf6, 0x3e,
	0x2f, 0xf1, 0x00, 0x3e, 0xfc, 0x86, 0x77, 0x0d, 0xc9, 0x3f, 0xba, 0xdf, 0x72, 0x7f, 0xe6, 0xfe,
	0xc2, 0x69, 0x77, 0xed, 0xc4, 0xce, 0x1d, 0xf7, 0xf2, 0x29, 0x3b, 0x33, 0xcf, 0xcc, 0x3e, 0x33,
	0xf3, 0xac, 0x03, 0x1d, 0x2e, 0x52, 0x64, 0x21, 0xa6, 0x4e, 0x92, 0xc6, 0x22, 0xb6, 0xdf, 0x6d,
	0x40, 0xe7, 0xaf, 0xc9, 0xf9, 0x24, 0x65, 0x11, 0x67, 0x53, 0xe1, 0xc7, 0x11, 0xf9, 0x01, 0xea,
	0x19, 0xc7, 0x74, 0x34, 0xa4, 0x46, 0xcf, 0xe8, 0x5b, 0x6e, 0x6e, 0x11, 0x02, 0x1b, 0x62, 0x36,
	0x1a, 0xd2, 0x75, 0xe5, 0x55, 0x67, 0x89, 0x15, 0xb3, 0x4b, 0xc6, 0x1f, 0xa8, 0xa9, 0xb1, 0xda,
	0x22, 0x3d, 0x68, 0x89, 0xd9, 0x4d, 0x26, 0xc6, 0xd3, 0xd4, 0x4f, 0x04, 0xdd, 0x50, 0xc1, 0xb2,
	0x8b, 0xec, 0x81, 0x25, 0x66, 0x03, 0xcf, 0x4b, 0x91, 0x73, 0x5a, 0xeb, 0x99, 0x7d, 0xcb, 0x5d,
	0x3a, 0x48, 0x17, 0x9a, 0x62, 0x36, 0x16, 0x4c, 0x64, 0x9c, 0xd6, 0x7b, 0x46, 0xbf, 0xe6, 0x2e,
	0xec, 0x45, 0xed, 0x41, 0x18, 0x67, 0x91, 0xa0, 0x8d, 0x9e, 0xd1, 0x37, 0xdd, 0xb2, 0x4b, 0xd6,
	0xbe, 0x0d, 0xe2, 0xe9, 0xe3, 0xc4, 0x0f, 0x91, 0x36, 0x55, 0x7c, 0xe9, 0x90, 0xf9, 0xca, 0xb8,
	0x44, 0xff, 0xfe, 0x41, 0x50, 0x4b, 0xe7, 0x97, 0x5c, 0x64, 0x1f, 0xda, 0xd3, 0x38, 0xba, 0xf3,
	0xd3, 0x90, 0xc9, 0x89, 0x70, 0x0a, 0x8a, 0x42, 0xd5, 0x49, 0x76, 0xa1, 0x26, 0x66, 0x7f, 0x23,
	0xd2, 0x96, 0xaa, 0xa0, 0x0d, 0x59, 0x3d, 0xc4, 0x30, 0x89, 0xe3, 0x40, 0xdd, 0xbe, 0xa9, 0xab,
	0x97, 0x5c, 0xe4, 0x54, 0xf6, 0x36, 0x8a, 0x92, 0x4c, 0x70, 0xda, 0xee, 0x99, 0xfd, 0xd6, 0xf1,
	0x9e, 0x53, 0x5d, 0x81, 0xa3, 0xc7, 0xa0, 0xbb, 0x71, 0x17, 0x68, 0x72, 0x06, 0xd6, 0x44, 0xb6,
	0xa9, 0x52, 0x3b, 0xdf, 0x90, 0xba, 0x84, 0x77, 0xff, 0x84, 0xcd, 0x72, 0x88, 0x50, 0x68, 0x14,
	0xd3, 0xd7, 0x6b, 0x2e, 0x4c, 0xb9, 0xd3, 0x7c, 0xb4, 0xeb, 0x8a, 0x7c, 0x6e, 0xd9, 0xef, 0x0d,
	0x68, 0x0e, 0x3c, 0x6f, 0x9c, 0xdc, 0x64, 0x62, 0x21, 0x06, 0xa3, 0x24, 0x06, 0x0a, 0x0d, 0xb5,
	0x85, 0x5c, 0x23, 0x35, 0xb7, 0x30, 0x57, 0x57, 0x66, 0x7e, 0xba, 0xb2, 0xaf, 0x0b, 0x86, 0x42,
	0x83, 0x2d, 0xe4, 0xa2, 0x08, 0xb3, 0x25, 0xe1, 0x5c, 0xb0, 0xf5, 0x8a, 0x60, 0xcb, 0x22, 0x6a,
	0x54, 0x45, 0x64, 0xff, 0x0f, 0x1d, 0x17, 0x9f, 0x86, 0x18, 0xa0, 0x40, 0xdd, 0xd1, 0xf7, 0xc8,
	0xbe, 0xc4, 0xc5, 0xac, 0x70, 0xb1, 0x8f, 0x60, 0x2b, 0x9f, 0xe3, 0x24, 0xd6, 0xd5, 0xcb, 0x60,
	0xa3, 0x0a, 0xee, 0xc3, 0xe6, 0x1b, 0x26, 0xa6, 0x0f, 0xc5, 0xe4, 0x5f, 0x47, 0x36, 0xa0, 0x76,
	0x11, 0x26, 0x62, 0x6e, 0xff, 0x04, 0x35, 0x97, 0xbd, 0x4c, 0x66, 0x6a, 0x60, 0xcb, 0xb5, 0xe7,
	0xf8, 0xb2, 0xab, 0x42, 0xc5, 0x45, 0x3e, 0x8f, 0xa6, 0x5f, 0xb8, 0xe0, 0x2d, 0x58, 0xff, 0x71,
	0x4c, 0xf9, 0x90, 0x09, 0x46, 0x0e, 0xc0, 0x0c, 0x59, 0x42, 0x0d, 0xa5, 0xb0, 0x1d, 0x67, 0x11,
	0x70, 0xae, 0x58, 0x72, 0x11, 0x89, 0x74, 0xee, 0xca, 0x78, 0xf7, 0x04, 0x9a, 0x85, 0x83, 0x6c,
	0x83, 0xf9, 0x88, 0xf3, 0xbc, 0xaa, 0x3c, 0xca, 0xe7, 0xf1, 0xcc, 0x82, 0x0c, 0xf3, 0xc1, 0x69,
	0xe3, 0x6c, 0xfd, 0xd4, 0xb0, 0x0f, 0xc0, 0x72, 0x31, 0x09, 0xe6, 0xa3, 0xe8, 0x2e, 0x96, 0x94,
	0x42, 0xe4, 0x9c, 0xdd, 0x63, 0x41, 0x29, 0x37, 0xed, 0x43, 0x80, 0xf3, 0xc0, 0x77, 0xf1, 0x29,
	0x43, 0xae, 0xd6, 0x9f, 0xea, 0x63, 0x81, 0xcb, 0x4d, 0xbb, 0x07, 0xcd, 0x71, 0xfa, 0xac, 0x2a,
	0xca, 0x4b, 0x53, 0x79, 0xc8, 0x31, 0xda, 0x38, 0xfe, 0x60, 0xc2, 0xce, 0x75, 0xec, 0xe1, 0x79,
	0x1c, 0x86, 0x59, 0x16, 0xf9, 0xd3, 0xfc, 0x05, 0x1f, 0x41, 0xfb, 0xdf, 0x58, 0xf0, 0x9b, 0x3b,
	0x99, 0xec, 0x23, 0x27, 0x2d, 0x67, 0x79, 0x63, 0xd7, 0x72, 0x8a, 0xb2, 0xf6, 0xda, 0xaf, 0x06,
	0x39, 0x82, 0xad, 0x8b, 0x67, 0x8c, 0xc4, 0x28, 0xf2, 0x85, 0xcf, 0x82, 0x81, 0xe7, 0x11, 0x58,
	0x8e, 0xa6, 0x0b, 0xce, 0xa2, 0x27, 0x7b, 0x8d, 0xfc, 0x0c, 0x44, 0x81, 0xff, 0x41, 0x31, 0x08,
	0x82, 0x2b, 0xfd, 0xfa, 0x49, 0xdd, 0x51, 0x4b, 0x5c, 0xc1, 0xfe, 0x91, 0x63, 0xf5, 0x8e, 0x0a,
	0x2d, 0x6c, 0x3b, 0x2b, 0xcb, 0x5b, 0xc9, 0x3a, 0x84, 0x8e, 0xca, 0x1a, 0x63, 0xe4, 0x69, 0x45,
	0xd4, 0x1d, 0xf5, 0xbb, 0x82, 0xfb, 0x2d, 0xaf, 0x3e, 0xf0, 0xbc, 0x6b, 0x7c, 0x29, 0xaa, 0xb7,
	0x9d, 0xb2, 0xf0, 0x5e, 0x21, 0xa4, 0xf5, 0x5b, 0x90, 0x2f, 0x11, 0xd2, 0x81, 0x95, 0xac, 0x13,
	0xf8, 0xb1, 0x94, 0x35, 0x4e, 0x30, 0xf2, 0xd8, 0x6d, 0x80, 0xf2, 0x69, 0x15, 0x7d, 0x6f, 0x39,
	0xd5, 0x37, 0xa7, 0xe6, 0xfa, 0x0b, 0xec, 0x16, 0x04, 0x3f, 0x9b, 0x64, 0x39, 0xc5, 0x47, 0x47,
	0xc1, 0xf7, 0xa1, 0x76, 0x8d, 0xba, 0xdd, 0xa2, 0x68, 0xf5, 0x0b, 0x28, 0x51, 0xb7, 0x75, 0xf5,
	0xef, 0xf6, 0xfb, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0xee, 0x9b, 0xc5, 0xef, 0x06, 0x00,
	0x00,
}