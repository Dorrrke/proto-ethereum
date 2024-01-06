// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.5
// source: ethereum_service.proto

package etheservicev1

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

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddr string `protobuf:"bytes,1,opt,name=ethereumAddr,proto3" json:"ethereumAddr,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_ethereum_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBalanceRequest) GetEthereumAddr() string {
	if x != nil {
		return x.EthereumAddr
	}
	return ""
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumBalance string `protobuf:"bytes,1,opt,name=ethereumBalance,proto3" json:"ethereumBalance,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_ethereum_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceResponse) GetEthereumBalance() string {
	if x != nil {
		return x.EthereumBalance
	}
	return ""
}

type GetLatestBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLatestBlockRequest) Reset() {
	*x = GetLatestBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockRequest) ProtoMessage() {}

func (x *GetLatestBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestBlockRequest.ProtoReflect.Descriptor instead.
func (*GetLatestBlockRequest) Descriptor() ([]byte, []int) {
	return file_ethereum_service_proto_rawDescGZIP(), []int{2}
}

type GetLatestBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber       string `protobuf:"bytes,1,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	TransactionsCount int64  `protobuf:"varint,2,opt,name=transactionsCount,proto3" json:"transactionsCount,omitempty"`
	BlockComplicacy   string `protobuf:"bytes,3,opt,name=blockComplicacy,proto3" json:"blockComplicacy,omitempty"`
	Date              string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetLatestBlockResponse) Reset() {
	*x = GetLatestBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockResponse) ProtoMessage() {}

func (x *GetLatestBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestBlockResponse.ProtoReflect.Descriptor instead.
func (*GetLatestBlockResponse) Descriptor() ([]byte, []int) {
	return file_ethereum_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLatestBlockResponse) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *GetLatestBlockResponse) GetTransactionsCount() int64 {
	if x != nil {
		return x.TransactionsCount
	}
	return 0
}

func (x *GetLatestBlockResponse) GetBlockComplicacy() string {
	if x != nil {
		return x.BlockComplicacy
	}
	return ""
}

func (x *GetLatestBlockResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type VerifyAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifyAddressRequest) Reset() {
	*x = VerifyAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAddressRequest) ProtoMessage() {}

func (x *VerifyAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAddressRequest.ProtoReflect.Descriptor instead.
func (*VerifyAddressRequest) Descriptor() ([]byte, []int) {
	return file_ethereum_service_proto_rawDescGZIP(), []int{4}
}

type VerifyAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberValid bool `protobuf:"varint,1,opt,name=numberValid,proto3" json:"numberValid,omitempty"`
}

func (x *VerifyAddressResponse) Reset() {
	*x = VerifyAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAddressResponse) ProtoMessage() {}

func (x *VerifyAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAddressResponse.ProtoReflect.Descriptor instead.
func (*VerifyAddressResponse) Descriptor() ([]byte, []int) {
	return file_ethereum_service_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyAddressResponse) GetNumberValid() bool {
	if x != nil {
		return x.NumberValid
	}
	return false
}

var File_ethereum_service_proto protoreflect.FileDescriptor

var file_ethereum_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64,
	0x64, 0x72, 0x22, 0x3e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa6, 0x01, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x15,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x32, 0xab, 0x02, 0x0a, 0x0f, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x61, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x65, 0x74, 0x68, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x3b, 0x65, 0x74, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethereum_service_proto_rawDescOnce sync.Once
	file_ethereum_service_proto_rawDescData = file_ethereum_service_proto_rawDesc
)

func file_ethereum_service_proto_rawDescGZIP() []byte {
	file_ethereum_service_proto_rawDescOnce.Do(func() {
		file_ethereum_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethereum_service_proto_rawDescData)
	})
	return file_ethereum_service_proto_rawDescData
}

var file_ethereum_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ethereum_service_proto_goTypes = []interface{}{
	(*GetBalanceRequest)(nil),      // 0: ethereumservice.GetBalanceRequest
	(*GetBalanceResponse)(nil),     // 1: ethereumservice.GetBalanceResponse
	(*GetLatestBlockRequest)(nil),  // 2: ethereumservice.GetLatestBlockRequest
	(*GetLatestBlockResponse)(nil), // 3: ethereumservice.GetLatestBlockResponse
	(*VerifyAddressRequest)(nil),   // 4: ethereumservice.VerifyAddressRequest
	(*VerifyAddressResponse)(nil),  // 5: ethereumservice.VerifyAddressResponse
}
var file_ethereum_service_proto_depIdxs = []int32{
	0, // 0: ethereumservice.EthereumService.GetBalance:input_type -> ethereumservice.GetBalanceRequest
	2, // 1: ethereumservice.EthereumService.GetLatestBlock:input_type -> ethereumservice.GetLatestBlockRequest
	4, // 2: ethereumservice.EthereumService.VerifyAddress:input_type -> ethereumservice.VerifyAddressRequest
	1, // 3: ethereumservice.EthereumService.GetBalance:output_type -> ethereumservice.GetBalanceResponse
	3, // 4: ethereumservice.EthereumService.GetLatestBlock:output_type -> ethereumservice.GetLatestBlockResponse
	5, // 5: ethereumservice.EthereumService.VerifyAddress:output_type -> ethereumservice.VerifyAddressResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ethereum_service_proto_init() }
func file_ethereum_service_proto_init() {
	if File_ethereum_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ethereum_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_ethereum_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
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
		file_ethereum_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestBlockRequest); i {
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
		file_ethereum_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestBlockResponse); i {
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
		file_ethereum_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAddressRequest); i {
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
		file_ethereum_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAddressResponse); i {
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
			RawDescriptor: file_ethereum_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ethereum_service_proto_goTypes,
		DependencyIndexes: file_ethereum_service_proto_depIdxs,
		MessageInfos:      file_ethereum_service_proto_msgTypes,
	}.Build()
	File_ethereum_service_proto = out.File
	file_ethereum_service_proto_rawDesc = nil
	file_ethereum_service_proto_goTypes = nil
	file_ethereum_service_proto_depIdxs = nil
}
