// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders/customer.proto

package orders

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Customer struct {
	// The Shopify Store to which customer record belongs to.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// A unique identifier for the customer.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// The date and time when the customer was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The date and time when the customer information was last updated.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The unique email address of the customer.
	Email string `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	// The customer's first name.
	FirstName string `protobuf:"bytes,11,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// The customer's last name.
	LastName string `protobuf:"bytes,12,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// The state of the customer's account with a shop. Default value: disabled.
	// Valid values:
	// disabled: The customer doesn't have an active account.
	//   Customer accounts can be disabled from the Shopify admin at any time.
	// invited: The customer has received an email invite to create an account.
	// enabled: The customer has created an account.
	// declined: The customer declined the email invite to create an account.
	State string `protobuf:"bytes,13,opt,name=state,proto3" json:"state,omitempty"`
	// A note about the customer.
	Note string `protobuf:"bytes,14,opt,name=note,proto3" json:"note,omitempty"`
	// The unique phone number (E.164 format) for this customer.
	Phone string `protobuf:"bytes,15,opt,name=phone,proto3" json:"phone,omitempty"`
	// The default address for the customer.
	DefaultAddress       *Address `protobuf:"bytes,16,opt,name=default_address,json=defaultAddress,proto3" json:"default_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b051bee154e0880, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *Customer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Customer) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Customer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Customer) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Customer) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Customer) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Customer) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *Customer) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Customer) GetDefaultAddress() *Address {
	if m != nil {
		return m.DefaultAddress
	}
	return nil
}

type ListCustomersRequest struct {
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The Shopify Store.
	ShopId               int64    `protobuf:"varint,3,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCustomersRequest) Reset()         { *m = ListCustomersRequest{} }
func (m *ListCustomersRequest) String() string { return proto.CompactTextString(m) }
func (*ListCustomersRequest) ProtoMessage()    {}
func (*ListCustomersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b051bee154e0880, []int{1}
}

func (m *ListCustomersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCustomersRequest.Unmarshal(m, b)
}
func (m *ListCustomersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCustomersRequest.Marshal(b, m, deterministic)
}
func (m *ListCustomersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCustomersRequest.Merge(m, src)
}
func (m *ListCustomersRequest) XXX_Size() int {
	return xxx_messageInfo_ListCustomersRequest.Size(m)
}
func (m *ListCustomersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCustomersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCustomersRequest proto.InternalMessageInfo

func (m *ListCustomersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCustomersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCustomersRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

type ListCustomersResponse struct {
	// There will be a maximum number of items returned based on the page_size field in the request.
	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCustomersResponse) Reset()         { *m = ListCustomersResponse{} }
func (m *ListCustomersResponse) String() string { return proto.CompactTextString(m) }
func (*ListCustomersResponse) ProtoMessage()    {}
func (*ListCustomersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b051bee154e0880, []int{2}
}

func (m *ListCustomersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCustomersResponse.Unmarshal(m, b)
}
func (m *ListCustomersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCustomersResponse.Marshal(b, m, deterministic)
}
func (m *ListCustomersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCustomersResponse.Merge(m, src)
}
func (m *ListCustomersResponse) XXX_Size() int {
	return xxx_messageInfo_ListCustomersResponse.Size(m)
}
func (m *ListCustomersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCustomersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCustomersResponse proto.InternalMessageInfo

func (m *ListCustomersResponse) GetCustomers() []*Customer {
	if m != nil {
		return m.Customers
	}
	return nil
}

func (m *ListCustomersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetCustomerRequest struct {
	// The Shopify Store.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// The field will contain id of the resource requested.
	CustomerId           int64    `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerRequest) Reset()         { *m = GetCustomerRequest{} }
func (m *GetCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerRequest) ProtoMessage()    {}
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b051bee154e0880, []int{3}
}

func (m *GetCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerRequest.Unmarshal(m, b)
}
func (m *GetCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerRequest.Merge(m, src)
}
func (m *GetCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerRequest.Size(m)
}
func (m *GetCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerRequest proto.InternalMessageInfo

func (m *GetCustomerRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *GetCustomerRequest) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

type CreateCustomerRequest struct {
	// The customer resource to create.
	Customer             *Customer `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCustomerRequest) Reset()         { *m = CreateCustomerRequest{} }
func (m *CreateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerRequest) ProtoMessage()    {}
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b051bee154e0880, []int{4}
}

func (m *CreateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerRequest.Unmarshal(m, b)
}
func (m *CreateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *CreateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerRequest.Merge(m, src)
}
func (m *CreateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerRequest.Size(m)
}
func (m *CreateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerRequest proto.InternalMessageInfo

func (m *CreateCustomerRequest) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type UpdateCustomerRequest struct {
	// The customer resource which replaces the resource on the server.
	Customer *Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCustomerRequest) Reset()         { *m = UpdateCustomerRequest{} }
func (m *UpdateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCustomerRequest) ProtoMessage()    {}
func (*UpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b051bee154e0880, []int{5}
}

func (m *UpdateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCustomerRequest.Unmarshal(m, b)
}
func (m *UpdateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCustomerRequest.Merge(m, src)
}
func (m *UpdateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCustomerRequest.Size(m)
}
func (m *UpdateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCustomerRequest proto.InternalMessageInfo

func (m *UpdateCustomerRequest) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

func (m *UpdateCustomerRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteCustomerRequest struct {
	// The Shopify Store.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// The resource id of the customer to be deleted.
	CustomerId           int64    `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCustomerRequest) Reset()         { *m = DeleteCustomerRequest{} }
func (m *DeleteCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCustomerRequest) ProtoMessage()    {}
func (*DeleteCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b051bee154e0880, []int{6}
}

func (m *DeleteCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCustomerRequest.Unmarshal(m, b)
}
func (m *DeleteCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCustomerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCustomerRequest.Merge(m, src)
}
func (m *DeleteCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCustomerRequest.Size(m)
}
func (m *DeleteCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCustomerRequest proto.InternalMessageInfo

func (m *DeleteCustomerRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *DeleteCustomerRequest) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func init() {
	proto.RegisterType((*Customer)(nil), "Customer")
	proto.RegisterType((*ListCustomersRequest)(nil), "ListCustomersRequest")
	proto.RegisterType((*ListCustomersResponse)(nil), "ListCustomersResponse")
	proto.RegisterType((*GetCustomerRequest)(nil), "GetCustomerRequest")
	proto.RegisterType((*CreateCustomerRequest)(nil), "CreateCustomerRequest")
	proto.RegisterType((*UpdateCustomerRequest)(nil), "UpdateCustomerRequest")
	proto.RegisterType((*DeleteCustomerRequest)(nil), "DeleteCustomerRequest")
}

func init() { proto.RegisterFile("orders/customer.proto", fileDescriptor_4b051bee154e0880) }

var fileDescriptor_4b051bee154e0880 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0x26, 0xa9, 0x89, 0x77, 0x13, 0x9b, 0xc8, 0x92, 0xe0, 0x51, 0x91, 0x86, 0x03, 0x35, 0x4f,
	0x17, 0xac, 0x4f, 0x22, 0x08, 0xb1, 0xa2, 0x14, 0xb4, 0xe8, 0x59, 0x5f, 0x7c, 0x39, 0xb6, 0xdd,
	0x49, 0xb2, 0xe4, 0xee, 0xf6, 0xbc, 0xdd, 0x03, 0xa9, 0xff, 0xd9, 0xdf, 0x20, 0x3b, 0xb7, 0x1b,
	0x6b, 0x42, 0x15, 0x7c, 0xdb, 0xf9, 0xe6, 0x9b, 0xfd, 0x66, 0x66, 0xbf, 0x85, 0x89, 0xaa, 0x05,
	0xd6, 0x7a, 0x7e, 0xd5, 0x68, 0xa3, 0x0a, 0xac, 0x93, 0xaa, 0x56, 0x46, 0x1d, 0x4d, 0x57, 0x4a,
	0xad, 0x72, 0x9c, 0x53, 0x74, 0xd9, 0x2c, 0xe7, 0x4b, 0x89, 0xb9, 0xc8, 0x0a, 0xae, 0x37, 0x8e,
	0x71, 0xbc, 0xcb, 0x30, 0xb2, 0x40, 0x6d, 0x78, 0x51, 0x39, 0xc2, 0xd8, 0xdd, 0xcc, 0x85, 0xa8,
	0x51, 0xeb, 0x16, 0x8d, 0x7f, 0x76, 0x21, 0x38, 0x75, 0x5a, 0xec, 0x01, 0xdc, 0xd5, 0x6b, 0x55,
	0x65, 0x52, 0x44, 0x9d, 0x69, 0x67, 0x76, 0x90, 0xf6, 0x6d, 0x78, 0x26, 0xd8, 0x10, 0xba, 0x52,
	0x44, 0x5d, 0xc2, 0xba, 0x52, 0xb0, 0x17, 0x00, 0x57, 0x35, 0x72, 0x83, 0x22, 0xe3, 0x26, 0x0a,
	0xa6, 0x9d, 0xd9, 0xe0, 0xe4, 0x28, 0x69, 0x3b, 0x48, 0x7c, 0x07, 0xc9, 0x85, 0xef, 0x20, 0x0d,
	0x1d, 0x7b, 0x61, 0x6c, 0x69, 0x53, 0x09, 0x5f, 0x1a, 0xfe, 0xbb, 0xd4, 0xb1, 0x17, 0x86, 0x8d,
	0xa1, 0x87, 0x05, 0x97, 0x79, 0x04, 0xd3, 0xce, 0x2c, 0x4c, 0xdb, 0x80, 0x3d, 0x02, 0x58, 0xca,
	0x5a, 0x9b, 0xac, 0xe4, 0x05, 0x46, 0x03, 0x4a, 0x85, 0x84, 0x9c, 0xf3, 0x02, 0xd9, 0x43, 0x08,
	0x73, 0xee, 0xb3, 0xf7, 0x28, 0x1b, 0x58, 0x80, 0x92, 0x63, 0xe8, 0x69, 0xc3, 0x0d, 0x46, 0x87,
	0xed, 0x8d, 0x14, 0x30, 0x06, 0x77, 0x4a, 0x65, 0x30, 0x1a, 0x12, 0x48, 0x67, 0xcb, 0xac, 0xd6,
	0xaa, 0xc4, 0x68, 0xd4, 0x32, 0x29, 0x60, 0xcf, 0x60, 0x24, 0x70, 0xc9, 0x9b, 0xdc, 0x64, 0x6e,
	0xad, 0xd1, 0x7d, 0x9a, 0x28, 0x48, 0x16, 0x6d, 0x9c, 0x0e, 0x1d, 0xc1, 0xc5, 0xf1, 0x06, 0xc6,
	0xef, 0xa5, 0x36, 0x7e, 0xe7, 0x3a, 0xc5, 0x6f, 0x0d, 0x6a, 0x63, 0xfb, 0xac, 0xf8, 0x0a, 0x33,
	0x2d, 0xaf, 0x91, 0xb6, 0xdf, 0x4b, 0x03, 0x0b, 0x7c, 0x96, 0xd7, 0x68, 0x67, 0xa4, 0xa4, 0x51,
	0x1b, 0x2c, 0xe9, 0x1d, 0xc2, 0x94, 0xe8, 0x17, 0x16, 0xb8, 0xf9, 0x6e, 0x07, 0x37, 0xdf, 0x2d,
	0x5e, 0xc3, 0x64, 0x47, 0x4c, 0x57, 0xaa, 0xd4, 0xc8, 0x9e, 0x42, 0xe8, 0x1d, 0xa6, 0xa3, 0xce,
	0xf4, 0x60, 0x36, 0x38, 0x09, 0x13, 0x4f, 0x4b, 0x7f, 0xe7, 0xd8, 0x13, 0x18, 0x95, 0xf8, 0xdd,
	0x64, 0x7b, 0xf2, 0x87, 0x16, 0xfe, 0xe8, 0x5b, 0x88, 0xcf, 0x81, 0xbd, 0xc3, 0xad, 0x90, 0x1f,
	0xea, 0x56, 0x43, 0x1d, 0xc3, 0xc0, 0x6b, 0x64, 0x5b, 0x67, 0x81, 0x87, 0xce, 0x44, 0xfc, 0x0a,
	0x26, 0xa7, 0xe4, 0x99, 0xdd, 0x2b, 0x1f, 0x43, 0xe0, 0x69, 0x34, 0xec, 0x1f, 0x8d, 0x6f, 0x53,
	0xf1, 0x0f, 0x98, 0x7c, 0x21, 0xe3, 0xfc, 0xad, 0xbe, 0x7b, 0x6b, 0x3d, 0x7b, 0x09, 0x83, 0xd6,
	0x78, 0xf4, 0xc7, 0x9c, 0xd2, 0xbe, 0x4f, 0xdf, 0xda, 0x6f, 0xf8, 0x81, 0xeb, 0x4d, 0xea, 0x5c,
	0x6d, 0xcf, 0xf1, 0x27, 0x98, 0xbc, 0xc1, 0x1c, 0xf7, 0xc5, 0xff, 0x7b, 0x1f, 0xaf, 0x83, 0xaf,
	0xfd, 0xf6, 0xff, 0x5e, 0xf6, 0x49, 0xfc, 0xf9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xf0,
	0x49, 0x4c, 0x2a, 0x04, 0x00, 0x00,
}
