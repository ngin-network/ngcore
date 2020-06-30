// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: types.proto

package ngtypes

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type NetworkType int32

const (
	NetworkType_ZERONET NetworkType = 0 // zero is INVALID: TODO: Add rule to block/tx check
	NetworkType_TESTNET NetworkType = 1 // TESTNET uses odd number
	NetworkType_MAINNET NetworkType = 2 // MAINNET uses even number
)

// Enum value maps for NetworkType.
var (
	NetworkType_name = map[int32]string{
		0: "ZERONET",
		1: "TESTNET",
		2: "MAINNET",
	}
	NetworkType_value = map[string]int32{
		"ZERONET": 0,
		"TESTNET": 1,
		"MAINNET": 2,
	}
)

func (x NetworkType) Enum() *NetworkType {
	p := new(NetworkType)
	*p = x
	return p
}

func (x NetworkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[0].Descriptor()
}

func (NetworkType) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[0]
}

func (x NetworkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkType.Descriptor instead.
func (NetworkType) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

type TxType int32

const (
	TxType_INVALID     TxType = 0
	TxType_GENERATE    TxType = 1
	TxType_REGISTER    TxType = 2
	TxType_LOGOUT      TxType = 3
	TxType_TRANSACTION TxType = 10
	TxType_ASSIGN      TxType = 20
	TxType_APPEND      TxType = 21
)

// Enum value maps for TxType.
var (
	TxType_name = map[int32]string{
		0:  "INVALID",
		1:  "GENERATE",
		2:  "REGISTER",
		3:  "LOGOUT",
		10: "TRANSACTION",
		20: "ASSIGN",
		21: "APPEND",
	}
	TxType_value = map[string]int32{
		"INVALID":     0,
		"GENERATE":    1,
		"REGISTER":    2,
		"LOGOUT":      3,
		"TRANSACTION": 10,
		"ASSIGN":      20,
		"APPEND":      21,
	}
)

func (x TxType) Enum() *TxType {
	p := new(TxType)
	*p = x
	return p
}

func (x TxType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[1].Descriptor()
}

func (TxType) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[1]
}

func (x TxType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxType.Descriptor instead.
func (TxType) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num      uint64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Owner    []byte `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`       // address -> balance in anonymous fields
	Contract []byte `protobuf:"bytes,5,opt,name=contract,proto3" json:"contract,omitempty"` // separate the code and state
	Context  []byte `protobuf:"bytes,6,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Account) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Account) GetContract() []byte {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *Account) GetContext() []byte {
	if x != nil {
		return x.Context
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network       NetworkType `protobuf:"varint,1,opt,name=network,proto3,enum=ngtypes.NetworkType" json:"network,omitempty"`
	Height        uint64      `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp     int64       `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevBlockHash []byte      `protobuf:"bytes,4,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"`
	TrieHash      []byte      `protobuf:"bytes,5,opt,name=trie_hash,json=trieHash,proto3" json:"trie_hash,omitempty"` // bytes prev_sheet_hash = 6; // the hash of sheet before block txs
	Difficulty    []byte      `protobuf:"bytes,7,opt,name=difficulty,proto3" json:"difficulty,omitempty"`             // diff = maxTarget / Big(header_hash)
	Nonce         []byte      `protobuf:"bytes,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Sheet prev_sheet = 9; // can be nil, the sheet before block txs
	Txs []*Tx `protobuf:"bytes,10,rep,name=txs,proto3" json:"txs,omitempty"` // miner should build the generate tx by themselves
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetNetwork() NetworkType {
	if x != nil {
		return x.Network
	}
	return NetworkType_ZERONET
}

func (x *Block) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Block) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Block) GetPrevBlockHash() []byte {
	if x != nil {
		return x.PrevBlockHash
	}
	return nil
}

func (x *Block) GetTrieHash() []byte {
	if x != nil {
		return x.TrieHash
	}
	return nil
}

func (x *Block) GetDifficulty() []byte {
	if x != nil {
		return x.Difficulty
	}
	return nil
}

func (x *Block) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Block) GetTxs() []*Tx {
	if x != nil {
		return x.Txs
	}
	return nil
}

// Or you can call it Op
type Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network       NetworkType `protobuf:"varint,1,opt,name=network,proto3,enum=ngtypes.NetworkType" json:"network,omitempty"`
	Type          TxType      `protobuf:"varint,2,opt,name=type,proto3,enum=ngtypes.TxType" json:"type,omitempty"`
	PrevBlockHash []byte      `protobuf:"bytes,3,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"` // only available within one block
	Convener      uint64      `protobuf:"varint,4,opt,name=convener,proto3" json:"convener,omitempty"`                                 // account num required
	Participants  [][]byte    `protobuf:"bytes,5,rep,name=participants,proto3" json:"participants,omitempty"`
	Fee           []byte      `protobuf:"bytes,6,opt,name=fee,proto3" json:"fee,omitempty"`
	Values        [][]byte    `protobuf:"bytes,7,rep,name=values,proto3" json:"values,omitempty"`
	// extension
	Extra []byte `protobuf:"bytes,8,opt,name=extra,proto3" json:"extra,omitempty"`
	Sign  []byte `protobuf:"bytes,9,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

func (x *Tx) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tx.ProtoReflect.Descriptor instead.
func (*Tx) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *Tx) GetNetwork() NetworkType {
	if x != nil {
		return x.Network
	}
	return NetworkType_ZERONET
}

func (x *Tx) GetType() TxType {
	if x != nil {
		return x.Type
	}
	return TxType_INVALID
}

func (x *Tx) GetPrevBlockHash() []byte {
	if x != nil {
		return x.PrevBlockHash
	}
	return nil
}

func (x *Tx) GetConvener() uint64 {
	if x != nil {
		return x.Convener
	}
	return 0
}

func (x *Tx) GetParticipants() [][]byte {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *Tx) GetFee() []byte {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *Tx) GetValues() [][]byte {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Tx) GetExtra() []byte {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *Tx) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

// Sheet is a simplified States
type Sheet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrevBlockHash []byte              `protobuf:"bytes,1,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`                                                                                 // use prevBlockHash to replace height to keep uniqueness
	Anonymous     map[string][]byte   `protobuf:"bytes,2,rep,name=anonymous,proto3" json:"anonymous,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // bs58_address(aka anonymous recipient) -> balance
	Accounts      map[uint64]*Account `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Sheet) Reset() {
	*x = Sheet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sheet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sheet) ProtoMessage() {}

func (x *Sheet) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sheet.ProtoReflect.Descriptor instead.
func (*Sheet) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *Sheet) GetPrevBlockHash() []byte {
	if x != nil {
		return x.PrevBlockHash
	}
	return nil
}

func (x *Sheet) GetAnonymous() map[string][]byte {
	if x != nil {
		return x.Anonymous
	}
	return nil
}

func (x *Sheet) GetAccounts() map[uint64]*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e,
	0x67, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x67, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x87, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6e, 0x67, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x26, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x69, 0x65, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x72, 0x69, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x74, 0x78,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x54, 0x78, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0x95, 0x02, 0x0a, 0x02, 0x54, 0x78,
	0x12, 0x2e, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x22, 0xb1, 0x02, 0x0a, 0x05, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70,
	0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x3b, 0x0a, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x12, 0x38,
	0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x41, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x67, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x34, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x5a, 0x45, 0x52, 0x4f, 0x4e, 0x45, 0x54, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53, 0x54, 0x4e, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x4d, 0x41, 0x49, 0x4e, 0x4e, 0x45, 0x54, 0x10, 0x02, 0x2a, 0x66, 0x0a, 0x06, 0x54,
	0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x53, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x14, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x50, 0x50, 0x45, 0x4e,
	0x44, 0x10, 0x15, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x67, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_types_proto_goTypes = []interface{}{
	(NetworkType)(0), // 0: ngtypes.NetworkType
	(TxType)(0),      // 1: ngtypes.TxType
	(*Account)(nil),  // 2: ngtypes.Account
	(*Block)(nil),    // 3: ngtypes.Block
	(*Tx)(nil),       // 4: ngtypes.Tx
	(*Sheet)(nil),    // 5: ngtypes.Sheet
	nil,              // 6: ngtypes.Sheet.AnonymousEntry
	nil,              // 7: ngtypes.Sheet.AccountsEntry
}
var file_types_proto_depIdxs = []int32{
	0, // 0: ngtypes.Block.network:type_name -> ngtypes.NetworkType
	4, // 1: ngtypes.Block.txs:type_name -> ngtypes.Tx
	0, // 2: ngtypes.Tx.network:type_name -> ngtypes.NetworkType
	1, // 3: ngtypes.Tx.type:type_name -> ngtypes.TxType
	6, // 4: ngtypes.Sheet.anonymous:type_name -> ngtypes.Sheet.AnonymousEntry
	7, // 5: ngtypes.Sheet.accounts:type_name -> ngtypes.Sheet.AccountsEntry
	2, // 6: ngtypes.Sheet.AccountsEntry.value:type_name -> ngtypes.Account
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tx); i {
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
		file_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sheet); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		EnumInfos:         file_types_proto_enumTypes,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
