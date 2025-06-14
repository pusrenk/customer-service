// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: internal/protobuf/customer.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Customer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy     string                 `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy     string                 `protobuf:"bytes,8,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Customer) Reset() {
	*x = Customer{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Customer) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Customer) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Customer) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Customer) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

type CreateCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedBy     string                 `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCustomerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateCustomerRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateCustomerRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type CreateCustomerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customer      *Customer              `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCustomerResponse) Reset() {
	*x = CreateCustomerResponse{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerResponse) ProtoMessage() {}

func (x *CreateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type GetCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{3}
}

func (x *GetCustomerRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCustomerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customer      *Customer              `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerResponse) Reset() {
	*x = GetCustomerResponse{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerResponse) ProtoMessage() {}

func (x *GetCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{4}
}

func (x *GetCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type ListCustomersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCustomersRequest) Reset() {
	*x = ListCustomersRequest{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomersRequest) ProtoMessage() {}

func (x *ListCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomersRequest.ProtoReflect.Descriptor instead.
func (*ListCustomersRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{5}
}

type ListCustomersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customers     []*Customer            `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCustomersResponse) Reset() {
	*x = ListCustomersResponse{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomersResponse) ProtoMessage() {}

func (x *ListCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomersResponse.ProtoReflect.Descriptor instead.
func (*ListCustomersResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{6}
}

func (x *ListCustomersResponse) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

type UpdateCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	UpdatedBy     string                 `protobuf:"bytes,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCustomerRequest) Reset() {
	*x = UpdateCustomerRequest{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerRequest) ProtoMessage() {}

func (x *UpdateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerRequest.ProtoReflect.Descriptor instead.
func (*UpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCustomerRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCustomerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateCustomerRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateCustomerRequest) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

type UpdateCustomerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customer      *Customer              `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCustomerResponse) Reset() {
	*x = UpdateCustomerResponse{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerResponse) ProtoMessage() {}

func (x *UpdateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerResponse.ProtoReflect.Descriptor instead.
func (*UpdateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type DeleteCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCustomerRequest) Reset() {
	*x = DeleteCustomerRequest{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCustomerRequest) ProtoMessage() {}

func (x *DeleteCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCustomerRequest.ProtoReflect.Descriptor instead.
func (*DeleteCustomerRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCustomerRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCustomerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCustomerResponse) Reset() {
	*x = DeleteCustomerResponse{}
	mi := &file_internal_protobuf_customer_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCustomerResponse) ProtoMessage() {}

func (x *DeleteCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_customer_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCustomerResponse.ProtoReflect.Descriptor instead.
func (*DeleteCustomerResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_customer_proto_rawDescGZIP(), []int{10}
}

var File_internal_protobuf_customer_proto protoreflect.FileDescriptor

const file_internal_protobuf_customer_proto_rawDesc = "" +
	"\n" +
	" internal/protobuf/customer.proto\x12\bcustomer\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8e\x02\n" +
	"\bCustomer\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"created_by\x18\x06 \x01(\tR\tcreatedBy\x129\n" +
	"\n" +
	"updated_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x1d\n" +
	"\n" +
	"updated_by\x18\b \x01(\tR\tupdatedBy\"v\n" +
	"\x15CreateCustomerRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x1d\n" +
	"\n" +
	"created_by\x18\x04 \x01(\tR\tcreatedBy\"H\n" +
	"\x16CreateCustomerResponse\x12.\n" +
	"\bcustomer\x18\x01 \x01(\v2\x12.customer.CustomerR\bcustomer\"$\n" +
	"\x12GetCustomerRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"E\n" +
	"\x13GetCustomerResponse\x12.\n" +
	"\bcustomer\x18\x01 \x01(\v2\x12.customer.CustomerR\bcustomer\"\x16\n" +
	"\x14ListCustomersRequest\"I\n" +
	"\x15ListCustomersResponse\x120\n" +
	"\tcustomers\x18\x01 \x03(\v2\x12.customer.CustomerR\tcustomers\"\x86\x01\n" +
	"\x15UpdateCustomerRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12\x1d\n" +
	"\n" +
	"updated_by\x18\x05 \x01(\tR\tupdatedBy\"H\n" +
	"\x16UpdateCustomerResponse\x12.\n" +
	"\bcustomer\x18\x01 \x01(\v2\x12.customer.CustomerR\bcustomer\"'\n" +
	"\x15DeleteCustomerRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"\x18\n" +
	"\x16DeleteCustomerResponse2\xae\x03\n" +
	"\x0fCustomerService\x12S\n" +
	"\x0eCreateCustomer\x12\x1f.customer.CreateCustomerRequest\x1a .customer.CreateCustomerResponse\x12J\n" +
	"\vGetCustomer\x12\x1c.customer.GetCustomerRequest\x1a\x1d.customer.GetCustomerResponse\x12P\n" +
	"\rListCustomers\x12\x1e.customer.ListCustomersRequest\x1a\x1f.customer.ListCustomersResponse\x12S\n" +
	"\x0eUpdateCustomer\x12\x1f.customer.UpdateCustomerRequest\x1a .customer.UpdateCustomerResponse\x12S\n" +
	"\x0eDeleteCustomer\x12\x1f.customer.DeleteCustomerRequest\x1a .customer.DeleteCustomerResponseB?Z=github.com/pusrenk/customer-service/internal/customers/modelsb\x06proto3"

var (
	file_internal_protobuf_customer_proto_rawDescOnce sync.Once
	file_internal_protobuf_customer_proto_rawDescData []byte
)

func file_internal_protobuf_customer_proto_rawDescGZIP() []byte {
	file_internal_protobuf_customer_proto_rawDescOnce.Do(func() {
		file_internal_protobuf_customer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_protobuf_customer_proto_rawDesc), len(file_internal_protobuf_customer_proto_rawDesc)))
	})
	return file_internal_protobuf_customer_proto_rawDescData
}

var file_internal_protobuf_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_protobuf_customer_proto_goTypes = []any{
	(*Customer)(nil),               // 0: customer.Customer
	(*CreateCustomerRequest)(nil),  // 1: customer.CreateCustomerRequest
	(*CreateCustomerResponse)(nil), // 2: customer.CreateCustomerResponse
	(*GetCustomerRequest)(nil),     // 3: customer.GetCustomerRequest
	(*GetCustomerResponse)(nil),    // 4: customer.GetCustomerResponse
	(*ListCustomersRequest)(nil),   // 5: customer.ListCustomersRequest
	(*ListCustomersResponse)(nil),  // 6: customer.ListCustomersResponse
	(*UpdateCustomerRequest)(nil),  // 7: customer.UpdateCustomerRequest
	(*UpdateCustomerResponse)(nil), // 8: customer.UpdateCustomerResponse
	(*DeleteCustomerRequest)(nil),  // 9: customer.DeleteCustomerRequest
	(*DeleteCustomerResponse)(nil), // 10: customer.DeleteCustomerResponse
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
}
var file_internal_protobuf_customer_proto_depIdxs = []int32{
	11, // 0: customer.Customer.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: customer.Customer.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: customer.CreateCustomerResponse.customer:type_name -> customer.Customer
	0,  // 3: customer.GetCustomerResponse.customer:type_name -> customer.Customer
	0,  // 4: customer.ListCustomersResponse.customers:type_name -> customer.Customer
	0,  // 5: customer.UpdateCustomerResponse.customer:type_name -> customer.Customer
	1,  // 6: customer.CustomerService.CreateCustomer:input_type -> customer.CreateCustomerRequest
	3,  // 7: customer.CustomerService.GetCustomer:input_type -> customer.GetCustomerRequest
	5,  // 8: customer.CustomerService.ListCustomers:input_type -> customer.ListCustomersRequest
	7,  // 9: customer.CustomerService.UpdateCustomer:input_type -> customer.UpdateCustomerRequest
	9,  // 10: customer.CustomerService.DeleteCustomer:input_type -> customer.DeleteCustomerRequest
	2,  // 11: customer.CustomerService.CreateCustomer:output_type -> customer.CreateCustomerResponse
	4,  // 12: customer.CustomerService.GetCustomer:output_type -> customer.GetCustomerResponse
	6,  // 13: customer.CustomerService.ListCustomers:output_type -> customer.ListCustomersResponse
	8,  // 14: customer.CustomerService.UpdateCustomer:output_type -> customer.UpdateCustomerResponse
	10, // 15: customer.CustomerService.DeleteCustomer:output_type -> customer.DeleteCustomerResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_internal_protobuf_customer_proto_init() }
func file_internal_protobuf_customer_proto_init() {
	if File_internal_protobuf_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_protobuf_customer_proto_rawDesc), len(file_internal_protobuf_customer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_customer_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_customer_proto_depIdxs,
		MessageInfos:      file_internal_protobuf_customer_proto_msgTypes,
	}.Build()
	File_internal_protobuf_customer_proto = out.File
	file_internal_protobuf_customer_proto_goTypes = nil
	file_internal_protobuf_customer_proto_depIdxs = nil
}
