// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: pix.proto

package pb

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

type AddPixKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keytype   string `protobuf:"bytes,1,opt,name=keytype,proto3" json:"keytype,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	AccountID string `protobuf:"bytes,3,opt,name=accountID,proto3" json:"accountID,omitempty"`
}

func (x *AddPixKeyRequest) Reset() {
	*x = AddPixKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pix_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPixKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPixKeyRequest) ProtoMessage() {}

func (x *AddPixKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pix_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPixKeyRequest.ProtoReflect.Descriptor instead.
func (*AddPixKeyRequest) Descriptor() ([]byte, []int) {
	return file_pix_proto_rawDescGZIP(), []int{0}
}

func (x *AddPixKeyRequest) GetKeytype() string {
	if x != nil {
		return x.Keytype
	}
	return ""
}

func (x *AddPixKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AddPixKeyRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type FindPixKeyByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keytype string `protobuf:"bytes,1,opt,name=keytype,proto3" json:"keytype,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *FindPixKeyByIDRequest) Reset() {
	*x = FindPixKeyByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pix_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPixKeyByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPixKeyByIDRequest) ProtoMessage() {}

func (x *FindPixKeyByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pix_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPixKeyByIDRequest.ProtoReflect.Descriptor instead.
func (*FindPixKeyByIDRequest) Descriptor() ([]byte, []int) {
	return file_pix_proto_rawDescGZIP(), []int{1}
}

func (x *FindPixKeyByIDRequest) GetKeytype() string {
	if x != nil {
		return x.Keytype
	}
	return ""
}

func (x *FindPixKeyByIDRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID     string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	AccountNumber string `protobuf:"bytes,2,opt,name=accountNumber,proto3" json:"accountNumber,omitempty"`
	BankID        string `protobuf:"bytes,3,opt,name=bankID,proto3" json:"bankID,omitempty"`
	BankName      string `protobuf:"bytes,4,opt,name=bankName,proto3" json:"bankName,omitempty"`
	OwnerName     string `protobuf:"bytes,5,opt,name=ownerName,proto3" json:"ownerName,omitempty"`
	CreatedAt     string `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pix_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_pix_proto_msgTypes[2]
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
	return file_pix_proto_rawDescGZIP(), []int{2}
}

func (x *Account) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Account) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *Account) GetBankID() string {
	if x != nil {
		return x.BankID
	}
	return ""
}

func (x *Account) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *Account) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *Account) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type FindPixKeyByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Keytype   string   `protobuf:"bytes,2,opt,name=keytype,proto3" json:"keytype,omitempty"`
	Key       string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Account   *Account `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	CreatedAt string   `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *FindPixKeyByIDResponse) Reset() {
	*x = FindPixKeyByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pix_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPixKeyByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPixKeyByIDResponse) ProtoMessage() {}

func (x *FindPixKeyByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pix_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPixKeyByIDResponse.ProtoReflect.Descriptor instead.
func (*FindPixKeyByIDResponse) Descriptor() ([]byte, []int) {
	return file_pix_proto_rawDescGZIP(), []int{3}
}

func (x *FindPixKeyByIDResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FindPixKeyByIDResponse) GetKeytype() string {
	if x != nil {
		return x.Keytype
	}
	return ""
}

func (x *FindPixKeyByIDResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FindPixKeyByIDResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *FindPixKeyByIDResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type AddPixKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddPixKeyResponse) Reset() {
	*x = AddPixKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pix_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPixKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPixKeyResponse) ProtoMessage() {}

func (x *AddPixKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pix_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPixKeyResponse.ProtoReflect.Descriptor instead.
func (*AddPixKeyResponse) Descriptor() ([]byte, []int) {
	return file_pix_proto_rawDescGZIP(), []int{4}
}

func (x *AddPixKeyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddPixKeyResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AddPixKeyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pix_proto protoreflect.FileDescriptor

var file_pix_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x65, 0x67, 0x6f, 0x63, 0x6c, 0x61,
	0x69, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x69, 0x78, 0x22, 0x5c, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x43, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xbd, 0x01,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x61, 0x6e, 0x6b, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb4, 0x01,
	0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x65, 0x67, 0x6f, 0x63, 0x6c, 0x61, 0x69, 0x72, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x70, 0x69, 0x78, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x51, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x69, 0x78, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xff, 0x01, 0x0a, 0x0a, 0x50, 0x69, 0x78, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x69, 0x78,
	0x4b, 0x65, 0x79, 0x12, 0x2f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x65, 0x67, 0x6f, 0x63, 0x6c, 0x61, 0x69, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x70, 0x69, 0x78, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x65, 0x67, 0x6f, 0x63, 0x6c, 0x61, 0x69, 0x72, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x70, 0x69, 0x78, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x49, 0x44, 0x12, 0x34, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x65, 0x67, 0x6f, 0x63, 0x6c, 0x61,
	0x69, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x69, 0x78, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50,
	0x69, 0x78, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69,
	0x65, 0x67, 0x6f, 0x63, 0x6c, 0x61, 0x69, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x69, 0x78,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x69, 0x78, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pix_proto_rawDescOnce sync.Once
	file_pix_proto_rawDescData = file_pix_proto_rawDesc
)

func file_pix_proto_rawDescGZIP() []byte {
	file_pix_proto_rawDescOnce.Do(func() {
		file_pix_proto_rawDescData = protoimpl.X.CompressGZIP(file_pix_proto_rawDescData)
	})
	return file_pix_proto_rawDescData
}

var file_pix_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pix_proto_goTypes = []interface{}{
	(*AddPixKeyRequest)(nil),       // 0: github.com.diegoclair.codepix.AddPixKeyRequest
	(*FindPixKeyByIDRequest)(nil),  // 1: github.com.diegoclair.codepix.FindPixKeyByIDRequest
	(*Account)(nil),                // 2: github.com.diegoclair.codepix.Account
	(*FindPixKeyByIDResponse)(nil), // 3: github.com.diegoclair.codepix.FindPixKeyByIDResponse
	(*AddPixKeyResponse)(nil),      // 4: github.com.diegoclair.codepix.AddPixKeyResponse
}
var file_pix_proto_depIdxs = []int32{
	2, // 0: github.com.diegoclair.codepix.FindPixKeyByIDResponse.account:type_name -> github.com.diegoclair.codepix.Account
	0, // 1: github.com.diegoclair.codepix.PixService.AddPixKey:input_type -> github.com.diegoclair.codepix.AddPixKeyRequest
	1, // 2: github.com.diegoclair.codepix.PixService.FindPixKeyByID:input_type -> github.com.diegoclair.codepix.FindPixKeyByIDRequest
	4, // 3: github.com.diegoclair.codepix.PixService.AddPixKey:output_type -> github.com.diegoclair.codepix.AddPixKeyResponse
	3, // 4: github.com.diegoclair.codepix.PixService.FindPixKeyByID:output_type -> github.com.diegoclair.codepix.FindPixKeyByIDResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pix_proto_init() }
func file_pix_proto_init() {
	if File_pix_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pix_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPixKeyRequest); i {
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
		file_pix_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPixKeyByIDRequest); i {
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
		file_pix_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_pix_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPixKeyByIDResponse); i {
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
		file_pix_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPixKeyResponse); i {
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
			RawDescriptor: file_pix_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pix_proto_goTypes,
		DependencyIndexes: file_pix_proto_depIdxs,
		MessageInfos:      file_pix_proto_msgTypes,
	}.Build()
	File_pix_proto = out.File
	file_pix_proto_rawDesc = nil
	file_pix_proto_goTypes = nil
	file_pix_proto_depIdxs = nil
}
