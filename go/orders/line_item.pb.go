// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders/line_item.proto

package orders

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type LineItem struct {
	// The Shopify domain to which line_item record belongs to.
	DomainId int64 `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// The ID of the line item.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// The number of items that were purchased.
	Quantity int64 `protobuf:"varint,11,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// The amount available to fulfill, calculated as follows:
	// quantity - max(refunded_quantity, fulfilled_quantity) -
	// pending_fulfilled_quantity - open_fulfilled_quantity
	FulfillableQuantity int64 `protobuf:"varint,12,opt,name=fulfillable_quantity,json=fulfillableQuantity,proto3" json:"fulfillable_quantity,omitempty"`
	// Whether the item requires shipping.
	RequiresShipping bool `protobuf:"varint,13,opt,name=requires_shipping,json=requiresShipping,proto3" json:"requires_shipping,omitempty"`
	// The name of the product variant.
	Name string `protobuf:"bytes,14,opt,name=name,proto3" json:"name,omitempty"`
	// The title of the product.
	Title string `protobuf:"bytes,15,opt,name=title,proto3" json:"title,omitempty"`
	// The item's SKU (stock keeping unit).
	Sku string `protobuf:"bytes,16,opt,name=sku,proto3" json:"sku,omitempty"`
	// The service provider that's fulfilling the item. Valid values:
	// manual, or the name of the provider, such as amazon or shipwire.
	FulfillmentService   string   `protobuf:"bytes,17,opt,name=fulfillment_service,json=fulfillmentService,proto3" json:"fulfillment_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LineItem) Reset()         { *m = LineItem{} }
func (m *LineItem) String() string { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()    {}
func (*LineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1d935ba8fcfd10, []int{0}
}

func (m *LineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineItem.Unmarshal(m, b)
}
func (m *LineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineItem.Marshal(b, m, deterministic)
}
func (m *LineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineItem.Merge(m, src)
}
func (m *LineItem) XXX_Size() int {
	return xxx_messageInfo_LineItem.Size(m)
}
func (m *LineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LineItem.DiscardUnknown(m)
}

var xxx_messageInfo_LineItem proto.InternalMessageInfo

func (m *LineItem) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *LineItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LineItem) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *LineItem) GetFulfillableQuantity() int64 {
	if m != nil {
		return m.FulfillableQuantity
	}
	return 0
}

func (m *LineItem) GetRequiresShipping() bool {
	if m != nil {
		return m.RequiresShipping
	}
	return false
}

func (m *LineItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LineItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *LineItem) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *LineItem) GetFulfillmentService() string {
	if m != nil {
		return m.FulfillmentService
	}
	return ""
}

type ListLineItemsRequest struct {
	// The parent resource name, for example, "shelves/shelf1"
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLineItemsRequest) Reset()         { *m = ListLineItemsRequest{} }
func (m *ListLineItemsRequest) String() string { return proto.CompactTextString(m) }
func (*ListLineItemsRequest) ProtoMessage()    {}
func (*ListLineItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1d935ba8fcfd10, []int{1}
}

func (m *ListLineItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLineItemsRequest.Unmarshal(m, b)
}
func (m *ListLineItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLineItemsRequest.Marshal(b, m, deterministic)
}
func (m *ListLineItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLineItemsRequest.Merge(m, src)
}
func (m *ListLineItemsRequest) XXX_Size() int {
	return xxx_messageInfo_ListLineItemsRequest.Size(m)
}
func (m *ListLineItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLineItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListLineItemsRequest proto.InternalMessageInfo

func (m *ListLineItemsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListLineItemsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListLineItemsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListLineItemsResponse struct {
	// There will be a maximum number of items returned based on the page_size field in the request.
	LineItems []*LineItem `protobuf:"bytes,1,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLineItemsResponse) Reset()         { *m = ListLineItemsResponse{} }
func (m *ListLineItemsResponse) String() string { return proto.CompactTextString(m) }
func (*ListLineItemsResponse) ProtoMessage()    {}
func (*ListLineItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1d935ba8fcfd10, []int{2}
}

func (m *ListLineItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLineItemsResponse.Unmarshal(m, b)
}
func (m *ListLineItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLineItemsResponse.Marshal(b, m, deterministic)
}
func (m *ListLineItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLineItemsResponse.Merge(m, src)
}
func (m *ListLineItemsResponse) XXX_Size() int {
	return xxx_messageInfo_ListLineItemsResponse.Size(m)
}
func (m *ListLineItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLineItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLineItemsResponse proto.InternalMessageInfo

func (m *ListLineItemsResponse) GetLineItems() []*LineItem {
	if m != nil {
		return m.LineItems
	}
	return nil
}

func (m *ListLineItemsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetLineItemRequest struct {
	// The field will contain id of the resource requested.
	LineItemId           int64    `protobuf:"varint,1,opt,name=line_item_id,json=lineItemId,proto3" json:"line_item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLineItemRequest) Reset()         { *m = GetLineItemRequest{} }
func (m *GetLineItemRequest) String() string { return proto.CompactTextString(m) }
func (*GetLineItemRequest) ProtoMessage()    {}
func (*GetLineItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1d935ba8fcfd10, []int{3}
}

func (m *GetLineItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLineItemRequest.Unmarshal(m, b)
}
func (m *GetLineItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLineItemRequest.Marshal(b, m, deterministic)
}
func (m *GetLineItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLineItemRequest.Merge(m, src)
}
func (m *GetLineItemRequest) XXX_Size() int {
	return xxx_messageInfo_GetLineItemRequest.Size(m)
}
func (m *GetLineItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLineItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLineItemRequest proto.InternalMessageInfo

func (m *GetLineItemRequest) GetLineItemId() int64 {
	if m != nil {
		return m.LineItemId
	}
	return 0
}

type CreateLineItemRequest struct {
	// The line_item id to use for this line_item.
	LineItemId string `protobuf:"bytes,2,opt,name=line_item_id,json=lineItemId,proto3" json:"line_item_id,omitempty"`
	// The line_item resource to create.
	LineItem             *LineItem `protobuf:"bytes,3,opt,name=line_item,json=lineItem,proto3" json:"line_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateLineItemRequest) Reset()         { *m = CreateLineItemRequest{} }
func (m *CreateLineItemRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLineItemRequest) ProtoMessage()    {}
func (*CreateLineItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1d935ba8fcfd10, []int{4}
}

func (m *CreateLineItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLineItemRequest.Unmarshal(m, b)
}
func (m *CreateLineItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLineItemRequest.Marshal(b, m, deterministic)
}
func (m *CreateLineItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLineItemRequest.Merge(m, src)
}
func (m *CreateLineItemRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLineItemRequest.Size(m)
}
func (m *CreateLineItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLineItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLineItemRequest proto.InternalMessageInfo

func (m *CreateLineItemRequest) GetLineItemId() string {
	if m != nil {
		return m.LineItemId
	}
	return ""
}

func (m *CreateLineItemRequest) GetLineItem() *LineItem {
	if m != nil {
		return m.LineItem
	}
	return nil
}

type UpdateLineItemRequest struct {
	// The line_item resource which replaces the resource on the server.
	LineItem *LineItem `protobuf:"bytes,1,opt,name=line_item,json=lineItem,proto3" json:"line_item,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateLineItemRequest) Reset()         { *m = UpdateLineItemRequest{} }
func (m *UpdateLineItemRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateLineItemRequest) ProtoMessage()    {}
func (*UpdateLineItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1d935ba8fcfd10, []int{5}
}

func (m *UpdateLineItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateLineItemRequest.Unmarshal(m, b)
}
func (m *UpdateLineItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateLineItemRequest.Marshal(b, m, deterministic)
}
func (m *UpdateLineItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateLineItemRequest.Merge(m, src)
}
func (m *UpdateLineItemRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateLineItemRequest.Size(m)
}
func (m *UpdateLineItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateLineItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateLineItemRequest proto.InternalMessageInfo

func (m *UpdateLineItemRequest) GetLineItem() *LineItem {
	if m != nil {
		return m.LineItem
	}
	return nil
}

func (m *UpdateLineItemRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteLineItemRequest struct {
	// The resource id of the line_item to be deleted.
	LineItemId           int64    `protobuf:"varint,1,opt,name=line_item_id,json=lineItemId,proto3" json:"line_item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLineItemRequest) Reset()         { *m = DeleteLineItemRequest{} }
func (m *DeleteLineItemRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteLineItemRequest) ProtoMessage()    {}
func (*DeleteLineItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1d935ba8fcfd10, []int{6}
}

func (m *DeleteLineItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLineItemRequest.Unmarshal(m, b)
}
func (m *DeleteLineItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLineItemRequest.Marshal(b, m, deterministic)
}
func (m *DeleteLineItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLineItemRequest.Merge(m, src)
}
func (m *DeleteLineItemRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteLineItemRequest.Size(m)
}
func (m *DeleteLineItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLineItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLineItemRequest proto.InternalMessageInfo

func (m *DeleteLineItemRequest) GetLineItemId() int64 {
	if m != nil {
		return m.LineItemId
	}
	return 0
}

func init() {
	proto.RegisterType((*LineItem)(nil), "LineItem")
	proto.RegisterType((*ListLineItemsRequest)(nil), "ListLineItemsRequest")
	proto.RegisterType((*ListLineItemsResponse)(nil), "ListLineItemsResponse")
	proto.RegisterType((*GetLineItemRequest)(nil), "GetLineItemRequest")
	proto.RegisterType((*CreateLineItemRequest)(nil), "CreateLineItemRequest")
	proto.RegisterType((*UpdateLineItemRequest)(nil), "UpdateLineItemRequest")
	proto.RegisterType((*DeleteLineItemRequest)(nil), "DeleteLineItemRequest")
}

func init() { proto.RegisterFile("orders/line_item.proto", fileDescriptor_4e1d935ba8fcfd10) }

var fileDescriptor_4e1d935ba8fcfd10 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0x26, 0xad, 0x57, 0x92, 0xe9, 0xfd, 0xe8, 0xad, 0xed, 0xb1, 0x54, 0x84, 0x90, 0x87, 0xa3,
	0x20, 0xa4, 0x78, 0x82, 0x20, 0xbe, 0xa9, 0x28, 0x85, 0x13, 0x34, 0xd5, 0x17, 0x5f, 0xc2, 0xd6,
	0x4c, 0xeb, 0xd8, 0xcd, 0x26, 0xcd, 0x6e, 0x44, 0x0f, 0xff, 0x18, 0xff, 0x54, 0xc9, 0xe6, 0x36,
	0xf6, 0xd4, 0xe3, 0xde, 0x66, 0xbe, 0xf9, 0x66, 0xbe, 0xcc, 0x97, 0x1d, 0x38, 0x2b, 0xaa, 0x0c,
	0x2b, 0x3d, 0x97, 0xa4, 0x30, 0x25, 0x83, 0x79, 0x5c, 0x56, 0x85, 0x29, 0xa6, 0xe1, 0xa6, 0x28,
	0x36, 0x12, 0xe7, 0x36, 0x5b, 0xd5, 0xeb, 0xf9, 0x9a, 0x50, 0x66, 0x69, 0x2e, 0xf4, 0xb6, 0x65,
	0x44, 0xbf, 0x7a, 0xe0, 0x5f, 0x92, 0xc2, 0x85, 0xc1, 0x9c, 0x3d, 0x80, 0x20, 0x2b, 0x72, 0x41,
	0x2a, 0xa5, 0x8c, 0x7b, 0xa1, 0x37, 0xeb, 0x27, 0x7e, 0x0b, 0x2c, 0x32, 0x76, 0x0c, 0x3d, 0xca,
	0x78, 0xcf, 0xa2, 0x3d, 0xca, 0xd8, 0x14, 0xfc, 0x5d, 0x2d, 0x94, 0x21, 0xf3, 0x83, 0x0f, 0x5b,
	0xae, 0xcb, 0xd9, 0x63, 0x18, 0xaf, 0x6b, 0xb9, 0x26, 0x29, 0xc5, 0x4a, 0x62, 0xda, 0xf1, 0x0e,
	0x2d, 0xef, 0xfe, 0x5e, 0xed, 0xbd, 0x6b, 0x79, 0x04, 0xa7, 0x15, 0xee, 0x6a, 0xaa, 0x50, 0xa7,
	0xfa, 0x0b, 0x95, 0x25, 0xa9, 0x0d, 0x3f, 0x0a, 0xbd, 0x99, 0x9f, 0x8c, 0x5c, 0x61, 0x79, 0x8d,
	0x33, 0x06, 0xf7, 0x94, 0xc8, 0x91, 0x1f, 0x87, 0xde, 0x2c, 0x48, 0x6c, 0xcc, 0xc6, 0x70, 0x60,
	0xc8, 0x48, 0xe4, 0x27, 0x16, 0x6c, 0x13, 0x36, 0x82, 0xbe, 0xde, 0xd6, 0x7c, 0x64, 0xb1, 0x26,
	0x64, 0x73, 0x70, 0xfa, 0x39, 0x2a, 0x93, 0x6a, 0xac, 0xbe, 0xd1, 0x67, 0xe4, 0xa7, 0x96, 0xc1,
	0xf6, 0x4a, 0xcb, 0xb6, 0x12, 0x7d, 0x85, 0xf1, 0x25, 0x69, 0xe3, 0x5c, 0xd2, 0x09, 0xee, 0x6a,
	0xd4, 0x86, 0x9d, 0xc1, 0xa0, 0x14, 0x15, 0x2a, 0x63, 0xad, 0x0a, 0x92, 0xeb, 0xac, 0x71, 0xb1,
	0x14, 0x1b, 0x4c, 0x35, 0x5d, 0xa1, 0xf5, 0xeb, 0x20, 0xf1, 0x1b, 0x60, 0x49, 0x57, 0xc8, 0x1e,
	0x02, 0xd8, 0xa2, 0x29, 0xb6, 0xa8, 0x78, 0xdf, 0x36, 0x5a, 0xfa, 0x87, 0x06, 0x88, 0x08, 0x26,
	0x7f, 0x69, 0xe9, 0xb2, 0x50, 0x1a, 0xd9, 0x0c, 0xa0, 0xfb, 0xb9, 0x9a, 0x7b, 0x61, 0x7f, 0x36,
	0xbc, 0x08, 0x62, 0xc7, 0x4b, 0x02, 0xe9, 0x3a, 0xd8, 0x39, 0x9c, 0x28, 0xfc, 0x6e, 0xd2, 0x3d,
	0x99, 0x9e, 0x95, 0x39, 0x6a, 0xe0, 0x77, 0x9d, 0xd4, 0x53, 0x60, 0x6f, 0xb0, 0x53, 0x72, 0x4b,
	0x85, 0x70, 0xd8, 0xe9, 0xfc, 0x79, 0x05, 0xe0, 0xc6, 0x2f, 0xb2, 0x48, 0xc0, 0xe4, 0x65, 0x85,
	0xc2, 0xe0, 0x5d, 0xad, 0xad, 0xea, 0x5e, 0x2b, 0x3b, 0x87, 0xa0, 0x63, 0xd8, 0xdd, 0x6f, 0xec,
	0xe0, 0x3b, 0x66, 0xf4, 0x13, 0x26, 0x1f, 0xcb, 0xec, 0x3f, 0x12, 0x37, 0x06, 0x78, 0xb7, 0x0e,
	0x60, 0xcf, 0x61, 0x58, 0xdb, 0x01, 0xf6, 0xa9, 0xdb, 0x2f, 0x19, 0x5e, 0x4c, 0xe3, 0xf6, 0x1a,
	0x62, 0x77, 0x0d, 0xf1, 0xeb, 0xe6, 0x1a, 0xde, 0x0a, 0xbd, 0x4d, 0xa0, 0xa5, 0x37, 0x71, 0xf4,
	0x0c, 0x26, 0xaf, 0x50, 0xe2, 0xdd, 0x0b, 0xfe, 0xe3, 0xcd, 0x0b, 0xff, 0xd3, 0xa0, 0xbd, 0xc4,
	0xd5, 0xc0, 0x8a, 0x3c, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x27, 0xc5, 0xdc, 0x46, 0x9a, 0x03,
	0x00, 0x00,
}
