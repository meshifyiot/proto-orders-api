// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders/fulfillment.proto

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

type Fulfillment struct {
	// The Shopify domain to which fulfillment record belongs to.
	DomainId int64 `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// The ID for the fulfillment.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// The unique numeric identifier for the order.
	OrderId int64 `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// The date and time when the fulfillment was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The date and time when the fulfillment was updated.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The status of the fulfillment. Valid values:
	// pending: The fulfillment is pending.
	// open: The fulfillment has been acknowledged by the service and is
	//  in processing.
	// success: The fulfillment was successful.
	// cancelled: The fulfillment was cancelled.
	// error: There was an error with the fulfillment request.
	// failure: The fulfillment request failed.
	Status string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	// The name of the tracking company.
	TrackingCompany string `protobuf:"bytes,11,opt,name=tracking_company,json=trackingCompany,proto3" json:"tracking_company,omitempty"`
	// A list of tracking numbers, provided by the shipping company.
	TrackingNumbers []string `protobuf:"bytes,12,rep,name=tracking_numbers,json=trackingNumbers,proto3" json:"tracking_numbers,omitempty"`
	// The URLs of tracking pages for the fulfillment.
	TrackingUrls []string `protobuf:"bytes,13,rep,name=tracking_urls,json=trackingUrls,proto3" json:"tracking_urls,omitempty"`
	// A historical record of each item in the fulfillment.
	LineItemIds          []int64  `protobuf:"varint,20,rep,packed,name=line_item_ids,json=lineItemIds,proto3" json:"line_item_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fulfillment) Reset()         { *m = Fulfillment{} }
func (m *Fulfillment) String() string { return proto.CompactTextString(m) }
func (*Fulfillment) ProtoMessage()    {}
func (*Fulfillment) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfc39d6414e919e, []int{0}
}

func (m *Fulfillment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fulfillment.Unmarshal(m, b)
}
func (m *Fulfillment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fulfillment.Marshal(b, m, deterministic)
}
func (m *Fulfillment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fulfillment.Merge(m, src)
}
func (m *Fulfillment) XXX_Size() int {
	return xxx_messageInfo_Fulfillment.Size(m)
}
func (m *Fulfillment) XXX_DiscardUnknown() {
	xxx_messageInfo_Fulfillment.DiscardUnknown(m)
}

var xxx_messageInfo_Fulfillment proto.InternalMessageInfo

func (m *Fulfillment) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *Fulfillment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Fulfillment) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Fulfillment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Fulfillment) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Fulfillment) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Fulfillment) GetTrackingCompany() string {
	if m != nil {
		return m.TrackingCompany
	}
	return ""
}

func (m *Fulfillment) GetTrackingNumbers() []string {
	if m != nil {
		return m.TrackingNumbers
	}
	return nil
}

func (m *Fulfillment) GetTrackingUrls() []string {
	if m != nil {
		return m.TrackingUrls
	}
	return nil
}

func (m *Fulfillment) GetLineItemIds() []int64 {
	if m != nil {
		return m.LineItemIds
	}
	return nil
}

// ListFulfillmentsRequest
//
// All filter fields applied with AND operator.
// Zero value in filter (by default) means no filtering.
//
type ListFulfillmentsRequest struct {
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The domain_id adds filtering by Shopify's domain.
	DomainId int64 `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// The folder_id adds filtering by folder.
	FolderId             int64    `protobuf:"varint,4,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFulfillmentsRequest) Reset()         { *m = ListFulfillmentsRequest{} }
func (m *ListFulfillmentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListFulfillmentsRequest) ProtoMessage()    {}
func (*ListFulfillmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfc39d6414e919e, []int{1}
}

func (m *ListFulfillmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFulfillmentsRequest.Unmarshal(m, b)
}
func (m *ListFulfillmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFulfillmentsRequest.Marshal(b, m, deterministic)
}
func (m *ListFulfillmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFulfillmentsRequest.Merge(m, src)
}
func (m *ListFulfillmentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListFulfillmentsRequest.Size(m)
}
func (m *ListFulfillmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFulfillmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFulfillmentsRequest proto.InternalMessageInfo

func (m *ListFulfillmentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListFulfillmentsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListFulfillmentsRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *ListFulfillmentsRequest) GetFolderId() int64 {
	if m != nil {
		return m.FolderId
	}
	return 0
}

type ListFulfillmentsResponse struct {
	// There will be a maximum number of items returned based on the page_size
	// field in the request.
	Fulfillments []*Fulfillment `protobuf:"bytes,1,rep,name=fulfillments,proto3" json:"fulfillments,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFulfillmentsResponse) Reset()         { *m = ListFulfillmentsResponse{} }
func (m *ListFulfillmentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListFulfillmentsResponse) ProtoMessage()    {}
func (*ListFulfillmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfc39d6414e919e, []int{2}
}

func (m *ListFulfillmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFulfillmentsResponse.Unmarshal(m, b)
}
func (m *ListFulfillmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFulfillmentsResponse.Marshal(b, m, deterministic)
}
func (m *ListFulfillmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFulfillmentsResponse.Merge(m, src)
}
func (m *ListFulfillmentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListFulfillmentsResponse.Size(m)
}
func (m *ListFulfillmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFulfillmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFulfillmentsResponse proto.InternalMessageInfo

func (m *ListFulfillmentsResponse) GetFulfillments() []*Fulfillment {
	if m != nil {
		return m.Fulfillments
	}
	return nil
}

func (m *ListFulfillmentsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetFulfillmentRequest struct {
	// The field will contain id of fulfillment.
	FulfillmentId        int64    `protobuf:"varint,1,opt,name=fulfillment_id,json=fulfillmentId,proto3" json:"fulfillment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFulfillmentRequest) Reset()         { *m = GetFulfillmentRequest{} }
func (m *GetFulfillmentRequest) String() string { return proto.CompactTextString(m) }
func (*GetFulfillmentRequest) ProtoMessage()    {}
func (*GetFulfillmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfc39d6414e919e, []int{3}
}

func (m *GetFulfillmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFulfillmentRequest.Unmarshal(m, b)
}
func (m *GetFulfillmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFulfillmentRequest.Marshal(b, m, deterministic)
}
func (m *GetFulfillmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFulfillmentRequest.Merge(m, src)
}
func (m *GetFulfillmentRequest) XXX_Size() int {
	return xxx_messageInfo_GetFulfillmentRequest.Size(m)
}
func (m *GetFulfillmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFulfillmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFulfillmentRequest proto.InternalMessageInfo

func (m *GetFulfillmentRequest) GetFulfillmentId() int64 {
	if m != nil {
		return m.FulfillmentId
	}
	return 0
}

type CreateFulfillmentRequest struct {
	// The fulfillment resource to create.
	Fulfillment          *Fulfillment `protobuf:"bytes,3,opt,name=fulfillment,proto3" json:"fulfillment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateFulfillmentRequest) Reset()         { *m = CreateFulfillmentRequest{} }
func (m *CreateFulfillmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFulfillmentRequest) ProtoMessage()    {}
func (*CreateFulfillmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfc39d6414e919e, []int{4}
}

func (m *CreateFulfillmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFulfillmentRequest.Unmarshal(m, b)
}
func (m *CreateFulfillmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFulfillmentRequest.Marshal(b, m, deterministic)
}
func (m *CreateFulfillmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFulfillmentRequest.Merge(m, src)
}
func (m *CreateFulfillmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFulfillmentRequest.Size(m)
}
func (m *CreateFulfillmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFulfillmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFulfillmentRequest proto.InternalMessageInfo

func (m *CreateFulfillmentRequest) GetFulfillment() *Fulfillment {
	if m != nil {
		return m.Fulfillment
	}
	return nil
}

type UpdateFulfillmentRequest struct {
	// The fulfillment resource which replaces the resource on the server.
	Fulfillment *Fulfillment `protobuf:"bytes,1,opt,name=fulfillment,proto3" json:"fulfillment,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateFulfillmentRequest) Reset()         { *m = UpdateFulfillmentRequest{} }
func (m *UpdateFulfillmentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFulfillmentRequest) ProtoMessage()    {}
func (*UpdateFulfillmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfc39d6414e919e, []int{5}
}

func (m *UpdateFulfillmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFulfillmentRequest.Unmarshal(m, b)
}
func (m *UpdateFulfillmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFulfillmentRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFulfillmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFulfillmentRequest.Merge(m, src)
}
func (m *UpdateFulfillmentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFulfillmentRequest.Size(m)
}
func (m *UpdateFulfillmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFulfillmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFulfillmentRequest proto.InternalMessageInfo

func (m *UpdateFulfillmentRequest) GetFulfillment() *Fulfillment {
	if m != nil {
		return m.Fulfillment
	}
	return nil
}

func (m *UpdateFulfillmentRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteFulfillmentRequest struct {
	// The resource id of the fulfillment to be deleted.
	FulfillmentId        int64    `protobuf:"varint,1,opt,name=fulfillment_id,json=fulfillmentId,proto3" json:"fulfillment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFulfillmentRequest) Reset()         { *m = DeleteFulfillmentRequest{} }
func (m *DeleteFulfillmentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFulfillmentRequest) ProtoMessage()    {}
func (*DeleteFulfillmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfc39d6414e919e, []int{6}
}

func (m *DeleteFulfillmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFulfillmentRequest.Unmarshal(m, b)
}
func (m *DeleteFulfillmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFulfillmentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFulfillmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFulfillmentRequest.Merge(m, src)
}
func (m *DeleteFulfillmentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFulfillmentRequest.Size(m)
}
func (m *DeleteFulfillmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFulfillmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFulfillmentRequest proto.InternalMessageInfo

func (m *DeleteFulfillmentRequest) GetFulfillmentId() int64 {
	if m != nil {
		return m.FulfillmentId
	}
	return 0
}

func init() {
	proto.RegisterType((*Fulfillment)(nil), "Fulfillment")
	proto.RegisterType((*ListFulfillmentsRequest)(nil), "ListFulfillmentsRequest")
	proto.RegisterType((*ListFulfillmentsResponse)(nil), "ListFulfillmentsResponse")
	proto.RegisterType((*GetFulfillmentRequest)(nil), "GetFulfillmentRequest")
	proto.RegisterType((*CreateFulfillmentRequest)(nil), "CreateFulfillmentRequest")
	proto.RegisterType((*UpdateFulfillmentRequest)(nil), "UpdateFulfillmentRequest")
	proto.RegisterType((*DeleteFulfillmentRequest)(nil), "DeleteFulfillmentRequest")
}

func init() { proto.RegisterFile("orders/fulfillment.proto", fileDescriptor_2cfc39d6414e919e) }

var fileDescriptor_2cfc39d6414e919e = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x49, 0xa3, 0xb5, 0x79, 0x69, 0x77, 0x65, 0xf0, 0xc7, 0xb8, 0x8b, 0x18, 0x22, 0x4a,
	0xbd, 0xa4, 0xb2, 0x9e, 0x44, 0x10, 0xea, 0xca, 0x4a, 0x44, 0x45, 0xe2, 0xee, 0xc5, 0x4b, 0x48,
	0x3b, 0x2f, 0x65, 0xe8, 0x24, 0x13, 0x33, 0x13, 0xd0, 0x3d, 0x7a, 0xf1, 0xe8, 0xbf, 0x2c, 0x99,
	0x24, 0x6e, 0xda, 0x2d, 0x74, 0x6f, 0x9d, 0xef, 0x7c, 0xe6, 0xcd, 0xb7, 0xf3, 0x7d, 0x2f, 0x40,
	0x65, 0xc9, 0xb0, 0x54, 0xb3, 0xb4, 0x12, 0x29, 0x17, 0x22, 0xc3, 0x5c, 0x07, 0x45, 0x29, 0xb5,
	0x3c, 0xf2, 0x56, 0x52, 0xae, 0x04, 0xce, 0xcc, 0x6a, 0x51, 0xa5, 0xb3, 0x94, 0xa3, 0x60, 0x71,
	0x96, 0xa8, 0x75, 0x4b, 0x3c, 0xd9, 0x26, 0x34, 0xcf, 0x50, 0xe9, 0x24, 0x2b, 0x1a, 0xc0, 0xff,
	0x6d, 0x83, 0x7b, 0x76, 0x55, 0x98, 0x1c, 0x83, 0xc3, 0x64, 0x96, 0xf0, 0x3c, 0xe6, 0x8c, 0x5a,
	0x9e, 0x35, 0xb5, 0xa3, 0x51, 0x23, 0x84, 0x8c, 0x1c, 0xc0, 0x80, 0x33, 0x3a, 0x30, 0xea, 0x80,
	0x33, 0xf2, 0x08, 0x46, 0xc6, 0x5b, 0xcd, 0xda, 0x46, 0xbd, 0x63, 0xd6, 0x21, 0x23, 0xaf, 0x01,
	0x96, 0x25, 0x26, 0x1a, 0x59, 0x9c, 0x68, 0x3a, 0xf2, 0xac, 0xa9, 0x7b, 0x72, 0x14, 0x34, 0x6e,
	0x82, 0xce, 0x4d, 0x70, 0xde, 0xb9, 0x89, 0x9c, 0x96, 0x9e, 0xeb, 0xfa, 0x68, 0x55, 0xb0, 0xee,
	0xa8, 0xb3, 0xff, 0x68, 0x4b, 0xcf, 0x35, 0x79, 0x00, 0x43, 0xa5, 0x13, 0x5d, 0x29, 0x0a, 0x9e,
	0x35, 0x75, 0xa2, 0x76, 0x45, 0x5e, 0xc0, 0x5d, 0x5d, 0x26, 0xcb, 0x35, 0xcf, 0x57, 0xf1, 0x52,
	0x66, 0x45, 0x92, 0xff, 0xa2, 0xae, 0x21, 0x0e, 0x3b, 0xfd, 0xb4, 0x91, 0x37, 0xd0, 0xbc, 0xca,
	0x16, 0x58, 0x2a, 0x3a, 0xf6, 0xec, 0x3e, 0xfa, 0xa5, 0x91, 0xc9, 0x53, 0x98, 0xfc, 0x47, 0xab,
	0x52, 0x28, 0x3a, 0x31, 0xdc, 0xb8, 0x13, 0x2f, 0x4a, 0xa1, 0x88, 0x0f, 0x13, 0xc1, 0x73, 0x8c,
	0xb9, 0xc6, 0x2c, 0xe6, 0x4c, 0xd1, 0x7b, 0x9e, 0x3d, 0xb5, 0x23, 0xb7, 0x16, 0x43, 0x8d, 0x59,
	0xc8, 0x94, 0xff, 0xd7, 0x82, 0x87, 0x9f, 0xb8, 0xd2, 0xbd, 0x20, 0x54, 0x84, 0x3f, 0x2a, 0x54,
	0x26, 0x90, 0x22, 0x59, 0x61, 0xac, 0xf8, 0x25, 0x9a, 0x40, 0x6e, 0x47, 0xa3, 0x5a, 0xf8, 0xc6,
	0x2f, 0x91, 0x3c, 0x06, 0x30, 0x9b, 0x5a, 0xae, 0x31, 0x37, 0xc1, 0x38, 0x91, 0xc1, 0xcf, 0x6b,
	0x61, 0x33, 0x4c, 0x7b, 0x2b, 0xcc, 0x63, 0x70, 0x52, 0x29, 0xda, 0xf4, 0x6e, 0x35, 0x9b, 0x8d,
	0x10, 0x32, 0x5f, 0x03, 0xbd, 0x6e, 0x48, 0x15, 0x32, 0x57, 0x48, 0x5e, 0xc2, 0xb8, 0xd7, 0x8a,
	0x8a, 0x5a, 0x9e, 0x3d, 0x75, 0x4f, 0xc6, 0x41, 0x0f, 0x8e, 0x36, 0x08, 0xf2, 0x1c, 0x0e, 0x73,
	0xfc, 0xa9, 0xe3, 0x6b, 0x5e, 0x27, 0xb5, 0xfc, 0xb5, 0xf3, 0xeb, 0xbf, 0x85, 0xfb, 0x1f, 0xb0,
	0x7f, 0x69, 0xf7, 0x08, 0xcf, 0xe0, 0xa0, 0x57, 0xf0, 0xaa, 0x35, 0x27, 0x3d, 0x35, 0x64, 0xfe,
	0x47, 0xa0, 0xa7, 0xa6, 0x8d, 0x76, 0x94, 0x08, 0xc0, 0xed, 0xc1, 0xe6, 0x35, 0xb6, 0x4d, 0xf7,
	0x01, 0xff, 0x8f, 0x05, 0xf4, 0xc2, 0x34, 0xd6, 0xfe, 0x62, 0xd6, 0x9e, 0x62, 0xe4, 0x0d, 0xb8,
	0x4d, 0x93, 0x9a, 0xd9, 0x34, 0x7f, 0x7e, 0x57, 0x4f, 0x9f, 0xd5, 0xe3, 0xfb, 0x39, 0x51, 0xeb,
	0xa8, 0x9d, 0x80, 0xfa, 0xb7, 0x3f, 0x07, 0xfa, 0x1e, 0x05, 0xee, 0x34, 0x72, 0xb3, 0x87, 0x79,
	0x37, 0xfa, 0x3e, 0x6c, 0x3e, 0x22, 0x8b, 0xa1, 0xb9, 0xec, 0xd5, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x42, 0x35, 0x4d, 0x80, 0x55, 0x04, 0x00, 0x00,
}