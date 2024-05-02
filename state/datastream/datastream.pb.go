// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: datastream.proto

package datastream

import (
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

type BookmarkType int32

const (
	BookmarkType_BOOKMARK_TYPE_UNSPECIFIED BookmarkType = 0
	BookmarkType_BOOKMARK_TYPE_BATCH       BookmarkType = 1
	BookmarkType_BOOKMARK_TYPE_L2_BLOCK    BookmarkType = 2
)

// Enum value maps for BookmarkType.
var (
	BookmarkType_name = map[int32]string{
		0: "BOOKMARK_TYPE_UNSPECIFIED",
		1: "BOOKMARK_TYPE_BATCH",
		2: "BOOKMARK_TYPE_L2_BLOCK",
	}
	BookmarkType_value = map[string]int32{
		"BOOKMARK_TYPE_UNSPECIFIED": 0,
		"BOOKMARK_TYPE_BATCH":       1,
		"BOOKMARK_TYPE_L2_BLOCK":    2,
	}
)

func (x BookmarkType) Enum() *BookmarkType {
	p := new(BookmarkType)
	*p = x
	return p
}

func (x BookmarkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BookmarkType) Descriptor() protoreflect.EnumDescriptor {
	return file_datastream_proto_enumTypes[0].Descriptor()
}

func (BookmarkType) Type() protoreflect.EnumType {
	return &file_datastream_proto_enumTypes[0]
}

func (x BookmarkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BookmarkType.Descriptor instead.
func (BookmarkType) EnumDescriptor() ([]byte, []int) {
	return file_datastream_proto_rawDescGZIP(), []int{0}
}

type EntryType int32

const (
	EntryType_ENTRY_TYPE_UNSPECIFIED EntryType = 0
	EntryType_ENTRY_TYPE_BATCH_START EntryType = 1
	EntryType_ENTRY_TYPE_L2_BLOCK    EntryType = 2
	EntryType_ENTRY_TYPE_TRANSACTION EntryType = 3
	EntryType_ENTRY_TYPE_BATCH_END   EntryType = 4
	EntryType_ENTRY_TYPE_UPDATE_GER  EntryType = 5
)

// Enum value maps for EntryType.
var (
	EntryType_name = map[int32]string{
		0: "ENTRY_TYPE_UNSPECIFIED",
		1: "ENTRY_TYPE_BATCH_START",
		2: "ENTRY_TYPE_L2_BLOCK",
		3: "ENTRY_TYPE_TRANSACTION",
		4: "ENTRY_TYPE_BATCH_END",
		5: "ENTRY_TYPE_UPDATE_GER",
	}
	EntryType_value = map[string]int32{
		"ENTRY_TYPE_UNSPECIFIED": 0,
		"ENTRY_TYPE_BATCH_START": 1,
		"ENTRY_TYPE_L2_BLOCK":    2,
		"ENTRY_TYPE_TRANSACTION": 3,
		"ENTRY_TYPE_BATCH_END":   4,
		"ENTRY_TYPE_UPDATE_GER":  5,
	}
)

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}

func (x EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_datastream_proto_enumTypes[1].Descriptor()
}

func (EntryType) Type() protoreflect.EnumType {
	return &file_datastream_proto_enumTypes[1]
}

func (x EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryType.Descriptor instead.
func (EntryType) EnumDescriptor() ([]byte, []int) {
	return file_datastream_proto_rawDescGZIP(), []int{1}
}

type BatchStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number  uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	ForkId  uint64 `protobuf:"varint,4,opt,name=fork_id,json=forkId,proto3" json:"fork_id,omitempty"`
	ChainId uint64 `protobuf:"varint,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *BatchStart) Reset() {
	*x = BatchStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchStart) ProtoMessage() {}

func (x *BatchStart) ProtoReflect() protoreflect.Message {
	mi := &file_datastream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchStart.ProtoReflect.Descriptor instead.
func (*BatchStart) Descriptor() ([]byte, []int) {
	return file_datastream_proto_rawDescGZIP(), []int{0}
}

func (x *BatchStart) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BatchStart) GetForkId() uint64 {
	if x != nil {
		return x.ForkId
	}
	return 0
}

func (x *BatchStart) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

type BatchEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number        uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	LocalExitRoot []byte `protobuf:"bytes,2,opt,name=local_exit_root,json=localExitRoot,proto3" json:"local_exit_root,omitempty"`
	StateRoot     []byte `protobuf:"bytes,3,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
}

func (x *BatchEnd) Reset() {
	*x = BatchEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchEnd) ProtoMessage() {}

func (x *BatchEnd) ProtoReflect() protoreflect.Message {
	mi := &file_datastream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchEnd.ProtoReflect.Descriptor instead.
func (*BatchEnd) Descriptor() ([]byte, []int) {
	return file_datastream_proto_rawDescGZIP(), []int{1}
}

func (x *BatchEnd) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BatchEnd) GetLocalExitRoot() []byte {
	if x != nil {
		return x.LocalExitRoot
	}
	return nil
}

func (x *BatchEnd) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

type L2Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number          uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	BatchNumber     uint64 `protobuf:"varint,2,opt,name=batch_number,json=batchNumber,proto3" json:"batch_number,omitempty"`
	Timestamp       uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DeltaTimestamp  uint32 `protobuf:"varint,4,opt,name=delta_timestamp,json=deltaTimestamp,proto3" json:"delta_timestamp,omitempty"`
	MinTimestamp    uint64 `protobuf:"varint,5,opt,name=min_timestamp,json=minTimestamp,proto3" json:"min_timestamp,omitempty"`
	L1Blockhash     []byte `protobuf:"bytes,6,opt,name=l1_blockhash,json=l1Blockhash,proto3" json:"l1_blockhash,omitempty"`
	L1InfotreeIndex uint32 `protobuf:"varint,7,opt,name=l1_infotree_index,json=l1InfotreeIndex,proto3" json:"l1_infotree_index,omitempty"`
	Hash            []byte `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	StateRoot       []byte `protobuf:"bytes,9,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	GlobalExitRoot  []byte `protobuf:"bytes,10,opt,name=global_exit_root,json=globalExitRoot,proto3" json:"global_exit_root,omitempty"`
	Coinbase        []byte `protobuf:"bytes,11,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
}

func (x *L2Block) Reset() {
	*x = L2Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L2Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L2Block) ProtoMessage() {}

func (x *L2Block) ProtoReflect() protoreflect.Message {
	mi := &file_datastream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L2Block.ProtoReflect.Descriptor instead.
func (*L2Block) Descriptor() ([]byte, []int) {
	return file_datastream_proto_rawDescGZIP(), []int{2}
}

func (x *L2Block) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *L2Block) GetBatchNumber() uint64 {
	if x != nil {
		return x.BatchNumber
	}
	return 0
}

func (x *L2Block) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *L2Block) GetDeltaTimestamp() uint32 {
	if x != nil {
		return x.DeltaTimestamp
	}
	return 0
}

func (x *L2Block) GetMinTimestamp() uint64 {
	if x != nil {
		return x.MinTimestamp
	}
	return 0
}

func (x *L2Block) GetL1Blockhash() []byte {
	if x != nil {
		return x.L1Blockhash
	}
	return nil
}

func (x *L2Block) GetL1InfotreeIndex() uint32 {
	if x != nil {
		return x.L1InfotreeIndex
	}
	return 0
}

func (x *L2Block) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *L2Block) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

func (x *L2Block) GetGlobalExitRoot() []byte {
	if x != nil {
		return x.GlobalExitRoot
	}
	return nil
}

func (x *L2Block) GetCoinbase() []byte {
	if x != nil {
		return x.Coinbase
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L2BlockNumber               uint64 `protobuf:"varint,1,opt,name=l2block_number,json=l2blockNumber,proto3" json:"l2block_number,omitempty"`
	IsValid                     bool   `protobuf:"varint,2,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	Encoded                     []byte `protobuf:"bytes,3,opt,name=encoded,proto3" json:"encoded,omitempty"`
	EffectiveGasPricePercentage uint32 `protobuf:"varint,4,opt,name=effective_gas_price_percentage,json=effectiveGasPricePercentage,proto3" json:"effective_gas_price_percentage,omitempty"`
	ImStateRoot                 []byte `protobuf:"bytes,5,opt,name=im_state_root,json=imStateRoot,proto3" json:"im_state_root,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_datastream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_datastream_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetL2BlockNumber() uint64 {
	if x != nil {
		return x.L2BlockNumber
	}
	return 0
}

func (x *Transaction) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *Transaction) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

func (x *Transaction) GetEffectiveGasPricePercentage() uint32 {
	if x != nil {
		return x.EffectiveGasPricePercentage
	}
	return 0
}

func (x *Transaction) GetImStateRoot() []byte {
	if x != nil {
		return x.ImStateRoot
	}
	return nil
}

type UpdateGER struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchNumber    uint64 `protobuf:"varint,1,opt,name=batch_number,json=batchNumber,proto3" json:"batch_number,omitempty"`
	Timestamp      uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GlobalExitRoot []byte `protobuf:"bytes,3,opt,name=global_exit_root,json=globalExitRoot,proto3" json:"global_exit_root,omitempty"`
	Coinbase       []byte `protobuf:"bytes,4,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	ForkId         uint64 `protobuf:"varint,5,opt,name=fork_id,json=forkId,proto3" json:"fork_id,omitempty"`
	ChainId        uint64 `protobuf:"varint,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	StateRoot      []byte `protobuf:"bytes,7,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
}

func (x *UpdateGER) Reset() {
	*x = UpdateGER{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGER) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGER) ProtoMessage() {}

func (x *UpdateGER) ProtoReflect() protoreflect.Message {
	mi := &file_datastream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGER.ProtoReflect.Descriptor instead.
func (*UpdateGER) Descriptor() ([]byte, []int) {
	return file_datastream_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGER) GetBatchNumber() uint64 {
	if x != nil {
		return x.BatchNumber
	}
	return 0
}

func (x *UpdateGER) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *UpdateGER) GetGlobalExitRoot() []byte {
	if x != nil {
		return x.GlobalExitRoot
	}
	return nil
}

func (x *UpdateGER) GetCoinbase() []byte {
	if x != nil {
		return x.Coinbase
	}
	return nil
}

func (x *UpdateGER) GetForkId() uint64 {
	if x != nil {
		return x.ForkId
	}
	return 0
}

func (x *UpdateGER) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *UpdateGER) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

type BookMark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  BookmarkType `protobuf:"varint,1,opt,name=type,proto3,enum=datastream.v1.BookmarkType" json:"type,omitempty"`
	Value uint64       `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BookMark) Reset() {
	*x = BookMark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookMark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookMark) ProtoMessage() {}

func (x *BookMark) ProtoReflect() protoreflect.Message {
	mi := &file_datastream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookMark.ProtoReflect.Descriptor instead.
func (*BookMark) Descriptor() ([]byte, []int) {
	return file_datastream_proto_rawDescGZIP(), []int{5}
}

func (x *BookMark) GetType() BookmarkType {
	if x != nil {
		return x.Type
	}
	return BookmarkType_BOOKMARK_TYPE_UNSPECIFIED
}

func (x *BookMark) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_datastream_proto protoreflect.FileDescriptor

var file_datastream_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x22, 0x58, 0x0a, 0x0a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6b, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x08, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x45,
	0x78, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0xf8, 0x02, 0x0a, 0x07, 0x4c, 0x32, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x69, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x31, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x6c, 0x31, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11,
	0x6c, 0x31, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x31, 0x49, 0x6e, 0x66, 0x6f, 0x74,
	0x72, 0x65, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x78, 0x69,
	0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x22, 0xd2, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x32, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6c, 0x32, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x43, 0x0a,
	0x1e, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0xe5, 0x01, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x45, 0x52, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x65, 0x78, 0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x78, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x6f,
	0x72, 0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0x51,
	0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2a, 0x62, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x4f, 0x4f, 0x4b, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x42, 0x4f, 0x4f, 0x4b, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x42, 0x41, 0x54, 0x43, 0x48, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x4f, 0x4f,
	0x4b, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x32, 0x5f, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x10, 0x02, 0x2a, 0xad, 0x01, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x45,
	0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x32, 0x5f, 0x42, 0x4c, 0x4f,
	0x43, 0x4b, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03,
	0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x4e,
	0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x47, 0x45, 0x52, 0x10, 0x05, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x30, 0x78, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x48, 0x65, 0x72,
	0x6d, 0x65, 0x7a, 0x2f, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x2d, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x2d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_datastream_proto_rawDescOnce sync.Once
	file_datastream_proto_rawDescData = file_datastream_proto_rawDesc
)

func file_datastream_proto_rawDescGZIP() []byte {
	file_datastream_proto_rawDescOnce.Do(func() {
		file_datastream_proto_rawDescData = protoimpl.X.CompressGZIP(file_datastream_proto_rawDescData)
	})
	return file_datastream_proto_rawDescData
}

var file_datastream_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_datastream_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_datastream_proto_goTypes = []interface{}{
	(BookmarkType)(0),   // 0: datastream.v1.BookmarkType
	(EntryType)(0),      // 1: datastream.v1.EntryType
	(*BatchStart)(nil),  // 2: datastream.v1.BatchStart
	(*BatchEnd)(nil),    // 3: datastream.v1.BatchEnd
	(*L2Block)(nil),     // 4: datastream.v1.L2Block
	(*Transaction)(nil), // 5: datastream.v1.Transaction
	(*UpdateGER)(nil),   // 6: datastream.v1.UpdateGER
	(*BookMark)(nil),    // 7: datastream.v1.BookMark
}
var file_datastream_proto_depIdxs = []int32{
	0, // 0: datastream.v1.BookMark.type:type_name -> datastream.v1.BookmarkType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datastream_proto_init() }
func file_datastream_proto_init() {
	if File_datastream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datastream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchStart); i {
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
		file_datastream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchEnd); i {
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
		file_datastream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L2Block); i {
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
		file_datastream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_datastream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGER); i {
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
		file_datastream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookMark); i {
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
			RawDescriptor: file_datastream_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datastream_proto_goTypes,
		DependencyIndexes: file_datastream_proto_depIdxs,
		EnumInfos:         file_datastream_proto_enumTypes,
		MessageInfos:      file_datastream_proto_msgTypes,
	}.Build()
	File_datastream_proto = out.File
	file_datastream_proto_rawDesc = nil
	file_datastream_proto_goTypes = nil
	file_datastream_proto_depIdxs = nil
}
