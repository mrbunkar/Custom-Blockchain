// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/types.proto

package proto

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header      *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transaction []*Transaction `protobuf:"bytes,2,rep,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[0]
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
	return file_proto_types_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTransaction() []*Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version       string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Height        int32  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	PrevBlockHash []byte `protobuf:"bytes,3,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"` // Prievious block hash
	DataHash      []byte `protobuf:"bytes,4,opt,name=dataHash,proto3" json:"dataHash,omitempty"`
	Timestamp     int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce         int32  `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[1]
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
	return file_proto_types_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Header) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Header) GetPrevBlockHash() []byte {
	if x != nil {
		return x.PrevBlockHash
	}
	return nil
}

func (x *Header) GetDataHash() []byte {
	if x != nil {
		return x.DataHash
	}
	return nil
}

func (x *Header) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetNonce() int32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrevOutIdx  uint32 `protobuf:"varint,1,opt,name=prevOutIdx,proto3" json:"prevOutIdx,omitempty"`
	PrevOutHash []byte `protobuf:"bytes,2,opt,name=prevOutHash,proto3" json:"prevOutHash,omitempty"`
	PublicKey   []byte `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Signature   []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{2}
}

func (x *Input) GetPrevOutIdx() uint32 {
	if x != nil {
		return x.PrevOutIdx
	}
	return 0
}

func (x *Input) GetPrevOutHash() []byte {
	if x != nil {
		return x.PrevOutHash
	}
	return nil
}

func (x *Input) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Input) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   int64  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Reciever []byte `protobuf:"bytes,2,opt,name=reciever,proto3" json:"reciever,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{3}
}

func (x *Output) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Output) GetReciever() []byte {
	if x != nil {
		return x.Reciever
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Input   []*Input  `protobuf:"bytes,2,rep,name=input,proto3" json:"input,omitempty"`
	Output  []*Output `protobuf:"bytes,3,rep,name=output,proto3" json:"output,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[4]
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
	return file_proto_types_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetInput() []*Input {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Transaction) GetOutput() []*Output {
	if x != nil {
		return x.Output
	}
	return nil
}

type None struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *None) Reset() {
	*x = None{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use None.ProtoReflect.Descriptor instead.
func (*None) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{5}
}

type Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *Chain) Reset() {
	*x = Chain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain) ProtoMessage() {}

func (x *Chain) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain.ProtoReflect.Descriptor instead.
func (*Chain) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{6}
}

func (x *Chain) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Height     int32    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	ListenAddr string   `protobuf:"bytes,3,opt,name=listenAddr,proto3" json:"listenAddr,omitempty"`
	Peerlist   []string `protobuf:"bytes,4,rep,name=peerlist,proto3" json:"peerlist,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_proto_types_proto_rawDescGZIP(), []int{7}
}

func (x *Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Version) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Version) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *Version) GetPeerlist() []string {
	if x != nil {
		return x.Peerlist
	}
	return nil
}

var File_proto_types_proto protoreflect.FileDescriptor

var file_proto_types_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb0, 0x01,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72,
	0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x22, 0x85, 0x01, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x65, 0x76, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x70, 0x72, 0x65, 0x76, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x65, 0x76, 0x4f, 0x75, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x76, 0x4f, 0x75, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x3c, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x06,
	0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x22, 0x27, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x1e, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22,
	0x77, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x65, 0x65, 0x72, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x65, 0x65, 0x72, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x9e, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x06, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x06, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x12, 0x08, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x11, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x72, 0x62, 0x75, 0x6e, 0x6b, 0x61, 0x72,
	0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_proto_rawDescOnce sync.Once
	file_proto_types_proto_rawDescData = file_proto_types_proto_rawDesc
)

func file_proto_types_proto_rawDescGZIP() []byte {
	file_proto_types_proto_rawDescOnce.Do(func() {
		file_proto_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_proto_rawDescData)
	})
	return file_proto_types_proto_rawDescData
}

var file_proto_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_types_proto_goTypes = []any{
	(*Block)(nil),       // 0: Block
	(*Header)(nil),      // 1: Header
	(*Input)(nil),       // 2: Input
	(*Output)(nil),      // 3: Output
	(*Transaction)(nil), // 4: Transaction
	(*None)(nil),        // 5: None
	(*Chain)(nil),       // 6: Chain
	(*Version)(nil),     // 7: Version
}
var file_proto_types_proto_depIdxs = []int32{
	1, // 0: Block.header:type_name -> Header
	4, // 1: Block.transaction:type_name -> Transaction
	2, // 2: Transaction.input:type_name -> Input
	3, // 3: Transaction.output:type_name -> Output
	0, // 4: Chain.blocks:type_name -> Block
	6, // 5: Node.HandleChain:input_type -> Chain
	0, // 6: Node.HandleBlock:input_type -> Block
	7, // 7: Node.Handshake:input_type -> Version
	4, // 8: Node.HandleTransaction:input_type -> Transaction
	6, // 9: Node.HandleChain:output_type -> Chain
	0, // 10: Node.HandleBlock:output_type -> Block
	7, // 11: Node.Handshake:output_type -> Version
	4, // 12: Node.HandleTransaction:output_type -> Transaction
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_types_proto_init() }
func file_proto_types_proto_init() {
	if File_proto_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_proto_types_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_proto_types_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Input); i {
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
		file_proto_types_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Output); i {
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
		file_proto_types_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
		file_proto_types_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*None); i {
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
		file_proto_types_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Chain); i {
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
		file_proto_types_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Version); i {
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
			RawDescriptor: file_proto_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_types_proto_goTypes,
		DependencyIndexes: file_proto_types_proto_depIdxs,
		MessageInfos:      file_proto_types_proto_msgTypes,
	}.Build()
	File_proto_types_proto = out.File
	file_proto_types_proto_rawDesc = nil
	file_proto_types_proto_goTypes = nil
	file_proto_types_proto_depIdxs = nil
}
