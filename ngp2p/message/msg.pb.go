// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: msg.proto

package message

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	ngtypes "github.com/ngchain/ngcore/ngtypes"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// enums
type MessageType int32

const (
	MessageType_INVALID MessageType = 0
	// basic messages
	MessageType_PING     MessageType = 1
	MessageType_PONG     MessageType = 2
	MessageType_REJECT   MessageType = 3
	MessageType_NOTFOUND MessageType = 4
	// chain
	MessageType_GETCHAIN MessageType = 10
	MessageType_CHAIN    MessageType = 11
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0:  "INVALID",
		1:  "PING",
		2:  "PONG",
		3:  "REJECT",
		4:  "NOTFOUND",
		10: "GETCHAIN",
		11: "CHAIN",
	}
	MessageType_value = map[string]int32{
		"INVALID":  0,
		"PING":     1,
		"PONG":     2,
		"REJECT":   3,
		"NOTFOUND": 4,
		"GETCHAIN": 10,
		"CHAIN":    11,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type ChainType int32

const (
	ChainType_Invalid ChainType = 0
	ChainType_Blocks  ChainType = 1
	ChainType_Headers ChainType = 2
	ChainType_Hashes  ChainType = 3
)

// Enum value maps for ChainType.
var (
	ChainType_name = map[int32]string{
		0: "Invalid",
		1: "Blocks",
		2: "Headers",
		3: "Hashes",
	}
	ChainType_value = map[string]int32{
		"Invalid": 0,
		"Blocks":  1,
		"Headers": 2,
		"Hashes":  3,
	}
)

func (x ChainType) Enum() *ChainType {
	p := new(ChainType)
	*p = x
	return p
}

func (x ChainType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChainType) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[1].Descriptor()
}

func (ChainType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[1]
}

func (x ChainType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChainType.Descriptor instead.
func (ChainType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

// common
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network     ngtypes.NetworkType `protobuf:"varint,1,opt,name=network,proto3,enum=ngtypes.NetworkType" json:"network,omitempty"`
	MessageId   []byte              `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	MessageType MessageType         `protobuf:"varint,3,opt,name=message_type,json=messageType,proto3,enum=pb.MessageType" json:"message_type,omitempty"`
	Timestamp   int64               `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PeerKey     []byte              `protobuf:"bytes,5,opt,name=peer_key,json=peerKey,proto3" json:"peer_key,omitempty"` // act as the node ID
	Sign        []byte              `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetNetwork() ngtypes.NetworkType {
	if x != nil {
		return x.Network
	}
	return ngtypes.NetworkType_ZERONET
}

func (x *Header) GetMessageId() []byte {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *Header) GetMessageType() MessageType {
	if x != nil {
		return x.MessageType
	}
	return MessageType_INVALID
}

func (x *Header) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetPeerKey() []byte {
	if x != nil {
		return x.PeerKey
	}
	return nil
}

func (x *Header) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload []byte  `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Message) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// wired
type PingPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin               uint64 `protobuf:"varint,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Latest               uint64 `protobuf:"varint,2,opt,name=latest,proto3" json:"latest,omitempty"`
	CheckpointHash       []byte `protobuf:"bytes,3,opt,name=checkpoint_hash,json=checkpointHash,proto3" json:"checkpoint_hash,omitempty"`
	CheckpointActualDiff []byte `protobuf:"bytes,4,opt,name=checkpoint_actual_diff,json=checkpointActualDiff,proto3" json:"checkpoint_actual_diff,omitempty"`
}

func (x *PingPayload) Reset() {
	*x = PingPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPayload) ProtoMessage() {}

func (x *PingPayload) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPayload.ProtoReflect.Descriptor instead.
func (*PingPayload) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *PingPayload) GetOrigin() uint64 {
	if x != nil {
		return x.Origin
	}
	return 0
}

func (x *PingPayload) GetLatest() uint64 {
	if x != nil {
		return x.Latest
	}
	return 0
}

func (x *PingPayload) GetCheckpointHash() []byte {
	if x != nil {
		return x.CheckpointHash
	}
	return nil
}

func (x *PingPayload) GetCheckpointActualDiff() []byte {
	if x != nil {
		return x.CheckpointActualDiff
	}
	return nil
}

type PongPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin               uint64 `protobuf:"varint,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Latest               uint64 `protobuf:"varint,2,opt,name=latest,proto3" json:"latest,omitempty"`
	CheckpointHash       []byte `protobuf:"bytes,3,opt,name=checkpoint_hash,json=checkpointHash,proto3" json:"checkpoint_hash,omitempty"`
	CheckpointActualDiff []byte `protobuf:"bytes,4,opt,name=checkpoint_actual_diff,json=checkpointActualDiff,proto3" json:"checkpoint_actual_diff,omitempty"`
}

func (x *PongPayload) Reset() {
	*x = PongPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongPayload) ProtoMessage() {}

func (x *PongPayload) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongPayload.ProtoReflect.Descriptor instead.
func (*PongPayload) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *PongPayload) GetOrigin() uint64 {
	if x != nil {
		return x.Origin
	}
	return 0
}

func (x *PongPayload) GetLatest() uint64 {
	if x != nil {
		return x.Latest
	}
	return 0
}

func (x *PongPayload) GetCheckpointHash() []byte {
	if x != nil {
		return x.CheckpointHash
	}
	return nil
}

func (x *PongPayload) GetCheckpointActualDiff() []byte {
	if x != nil {
		return x.CheckpointActualDiff
	}
	return nil
}

type GetChainPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ChainType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.ChainType" json:"type,omitempty"`
	From [][]byte  `protobuf:"bytes,2,rep,name=from,proto3" json:"from,omitempty"` // beginning hashes
	To   []byte    `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`     // ending hash
}

func (x *GetChainPayload) Reset() {
	*x = GetChainPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChainPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChainPayload) ProtoMessage() {}

func (x *GetChainPayload) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChainPayload.ProtoReflect.Descriptor instead.
func (*GetChainPayload) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{4}
}

func (x *GetChainPayload) GetType() ChainType {
	if x != nil {
		return x.Type
	}
	return ChainType_Invalid
}

func (x *GetChainPayload) GetFrom() [][]byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *GetChainPayload) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

type ChainPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes [][]byte         `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
	Blocks []*ngtypes.Block `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *ChainPayload) Reset() {
	*x = ChainPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainPayload) ProtoMessage() {}

func (x *ChainPayload) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainPayload.ProtoReflect.Descriptor instead.
func (*ChainPayload) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{5}
}

func (x *ChainPayload) GetHashes() [][]byte {
	if x != nil {
		return x.Hashes
	}
	return nil
}

func (x *ChainPayload) GetBlocks() []*ngtypes.Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

// fast-sync for state
// support checkpoint height only
// when fast-sync: 1. sync to latest checkpoint 2. sync to latest checkpoint state 3. sync the remaining blocks 4. update local state
type GetSheetPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckpointHeight uint64 `protobuf:"varint,1,opt,name=checkpoint_height,json=checkpointHeight,proto3" json:"checkpoint_height,omitempty"` // required
	CheckpointHash   []byte `protobuf:"bytes,2,opt,name=checkpoint_hash,json=checkpointHash,proto3" json:"checkpoint_hash,omitempty"`        // required
}

func (x *GetSheetPayload) Reset() {
	*x = GetSheetPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSheetPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSheetPayload) ProtoMessage() {}

func (x *GetSheetPayload) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSheetPayload.ProtoReflect.Descriptor instead.
func (*GetSheetPayload) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{6}
}

func (x *GetSheetPayload) GetCheckpointHeight() uint64 {
	if x != nil {
		return x.CheckpointHeight
	}
	return 0
}

func (x *GetSheetPayload) GetCheckpointHash() []byte {
	if x != nil {
		return x.CheckpointHash
	}
	return nil
}

type SheetPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sheet *ngtypes.Sheet `protobuf:"bytes,2,opt,name=sheet,proto3" json:"sheet,omitempty"`
}

func (x *SheetPayload) Reset() {
	*x = SheetPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetPayload) ProtoMessage() {}

func (x *SheetPayload) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetPayload.ProtoReflect.Descriptor instead.
func (*SheetPayload) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{7}
}

func (x *SheetPayload) GetSheet() *ngtypes.Sheet {
	if x != nil {
		return x.Sheet
	}
	return nil
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x13, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x2e, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x32,
	0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x65, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22,
	0x47, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x50, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x63, 0x74,
	0x75, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x6e, 0x67,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x34, 0x0a, 0x16, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61,
	0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x75,
	0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x22, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f,
	0x22, 0x4e, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x22, 0x67, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0x34, 0x0a, 0x0c, 0x53, 0x68, 0x65,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2a,
	0x61, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4e,
	0x4f, 0x54, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x45, 0x54,
	0x43, 0x48, 0x41, 0x49, 0x4e, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x48, 0x41, 0x49, 0x4e,
	0x10, 0x0b, 0x2a, 0x3d, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x10,
	0x03, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x67, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x67, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e,
	0x67, 0x70, 0x32, 0x70, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_msg_proto_goTypes = []interface{}{
	(MessageType)(0),         // 0: pb.MessageType
	(ChainType)(0),           // 1: pb.ChainType
	(*Header)(nil),           // 2: pb.Header
	(*Message)(nil),          // 3: pb.Message
	(*PingPayload)(nil),      // 4: pb.PingPayload
	(*PongPayload)(nil),      // 5: pb.PongPayload
	(*GetChainPayload)(nil),  // 6: pb.GetChainPayload
	(*ChainPayload)(nil),     // 7: pb.ChainPayload
	(*GetSheetPayload)(nil),  // 8: pb.GetSheetPayload
	(*SheetPayload)(nil),     // 9: pb.SheetPayload
	(ngtypes.NetworkType)(0), // 10: ngtypes.NetworkType
	(*ngtypes.Block)(nil),    // 11: ngtypes.Block
	(*ngtypes.Sheet)(nil),    // 12: ngtypes.Sheet
}
var file_msg_proto_depIdxs = []int32{
	10, // 0: pb.Header.network:type_name -> ngtypes.NetworkType
	0,  // 1: pb.Header.message_type:type_name -> pb.MessageType
	2,  // 2: pb.Message.header:type_name -> pb.Header
	1,  // 3: pb.GetChainPayload.type:type_name -> pb.ChainType
	11, // 4: pb.ChainPayload.blocks:type_name -> ngtypes.Block
	12, // 5: pb.SheetPayload.sheet:type_name -> ngtypes.Sheet
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChainPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSheetPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
