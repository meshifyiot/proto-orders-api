// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders/order.proto

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

// An order is a customer's completed request to purchase one or more products
// from a shop. An order is created when a customer completes the checkout
// process, during which time they provide an email address or phone number,
// billing address and payment information.
type Order struct {
	// The Shopify Store to which order belongs to.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// The Shopify ID of the order.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Customer facing order ID, used by the shop owner and customer.
	OrderNumber int64 `protobuf:"varint,3,opt,name=order_number,json=orderNumber,proto3" json:"order_number,omitempty"`
	// The customer's email address.
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// The order's position in the shop's count of orders.
	// Numbers are sequential and start at 1.
	Number int64 `protobuf:"varint,6,opt,name=number,proto3" json:"number,omitempty"`
	// An optional note that a shop owner can attach to the order.
	Note string `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
	// A unique token for the order.
	Token string `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	// The autogenerated date and time when the order was created in Shopify.
	// The value for this property cannot be changed.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The date and time when an order was processed.
	// This value is the date that appears on your orders and that's used in
	// the analytic reports. By default, it matches the created_at value.
	ProcessedAt *timestamp.Timestamp `protobuf:"bytes,11,opt,name=processed_at,json=processedAt,proto3" json:"processed_at,omitempty"`
	// The date and time when the order was last modified.
	// Its value can change when no visible fields of an order have been updated.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The date and time when the order was closed.
	ClosedAt *timestamp.Timestamp `protobuf:"bytes,13,opt,name=closed_at,json=closedAt,proto3" json:"closed_at,omitempty"`
	// The date and time when the order was canceled.
	CancelledAt *timestamp.Timestamp `protobuf:"bytes,14,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
	// The order name, generated by combining the order_number property with the
	// order prefix and suffix that are set in the merchant's general settings.
	// This is different from the id property, which is the ID of the order used
	// by the Shopify API. This field can also be set by the Shopify API
	// to be any string value.
	Name string `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// The reason why the order was canceled. Valid values:
	// customer: The customer canceled the order.
	// fraud: The order was fraudulent.
	// inventory: Items in the order were not in inventory.
	// declined: The payment was declined.
	// other: A reason not in this list.
	CancelReason string `protobuf:"bytes,27,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
	// The ID of the Shopify app that created the order.
	AppId        int64  `protobuf:"varint,39,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ContactEmail string `protobuf:"bytes,53,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	// The mailing address to where the order will be shipped.
	// This address is optional and will not be available on orders that
	// do not require shipping.
	ShippingAddress *Address `protobuf:"bytes,65,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	// A list of line item objects, each containing information about an item
	// in the order.
	LineItemIds []int64 `protobuf:"varint,66,rep,packed,name=line_item_ids,json=lineItemIds,proto3" json:"line_item_ids,omitempty"`
	// Information about the customer. The order might not have a customer
	// and apps should not depend on the existence of a customer object.
	// This value might be null if the order was created through Shopify POS.
	CustomerId int64 `protobuf:"varint,68,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// ID of folder that order belongs to.
	FolderId int32 `protobuf:"varint,69,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Whether this is a test order.
	Test bool `protobuf:"varint,70,opt,name=test,proto3" json:"test,omitempty"`
	// The sum of all line item prices, discounts, shipping, taxes, and tips
	// in the shop currency (in cents). Must be positive.
	TotalPrice int64 `protobuf:"varint,71,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	// The price of the order in the shop currency (in cents) after discounts
	// but before shipping, taxes, and tips.
	SubtotalPrice int64 `protobuf:"varint,72,opt,name=subtotal_price,json=subtotalPrice,proto3" json:"subtotal_price,omitempty"`
	// The sum of all line item weights in grams.
	TotalWeight int64 `protobuf:"varint,73,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	// The sum of all the taxes applied to the order in th shop currency
	// (in cents). Must be positive.
	TotalTax int64 `protobuf:"varint,74,opt,name=total_tax,json=totalTax,proto3" json:"total_tax,omitempty"`
	// Whether taxes are included in the order subtotal.
	TaxesIncluded bool `protobuf:"varint,75,opt,name=taxes_included,json=taxesIncluded,proto3" json:"taxes_included,omitempty"`
	// The three-letter code (ISO 4217 format) for the shop currency.
	Currency string `protobuf:"bytes,76,opt,name=currency,proto3" json:"currency,omitempty"`
	// The status of payments associated with the order.
	// Can only be set when the order is created. Valid values:
	// pending: The payments are pending. Payment might fail in this state.
	//  Check again to confirm whether the payments have been paid successfully.
	// authorized: The payments have been authorized.
	// partially_paid: The order have been partially paid.
	// paid: The payments have been paid.
	// partially_refunded: The payments have been partially refunded.
	// refunded: The payments have been refunded.
	// voided: The payments have been voided.
	FinancialStatus string `protobuf:"bytes,77,opt,name=financial_status,json=financialStatus,proto3" json:"financial_status,omitempty"`
	Confirmed       bool   `protobuf:"varint,78,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	// The sum of all line item prices, discounts, shipping, taxes, and tips
	// in the in USD cents. Must be positive.
	TotalPriceUsd int64 `protobuf:"varint,79,opt,name=total_price_usd,json=totalPriceUsd,proto3" json:"total_price_usd,omitempty"`
	// Extra information that is added to the order. Appears in the Additional
	// details section of an order details page. Each array entry must contain
	// a hash with name and value keys.
	NoteAttributes       []*NoteAttribute `protobuf:"bytes,80,rep,name=note_attributes,json=noteAttributes,proto3" json:"note_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_d70621d7f7645dc1, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetOrderNumber() int64 {
	if m != nil {
		return m.OrderNumber
	}
	return 0
}

func (m *Order) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Order) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Order) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *Order) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Order) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Order) GetProcessedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ProcessedAt
	}
	return nil
}

func (m *Order) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Order) GetClosedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ClosedAt
	}
	return nil
}

func (m *Order) GetCancelledAt() *timestamp.Timestamp {
	if m != nil {
		return m.CancelledAt
	}
	return nil
}

func (m *Order) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order) GetCancelReason() string {
	if m != nil {
		return m.CancelReason
	}
	return ""
}

func (m *Order) GetAppId() int64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *Order) GetContactEmail() string {
	if m != nil {
		return m.ContactEmail
	}
	return ""
}

func (m *Order) GetShippingAddress() *Address {
	if m != nil {
		return m.ShippingAddress
	}
	return nil
}

func (m *Order) GetLineItemIds() []int64 {
	if m != nil {
		return m.LineItemIds
	}
	return nil
}

func (m *Order) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Order) GetFolderId() int32 {
	if m != nil {
		return m.FolderId
	}
	return 0
}

func (m *Order) GetTest() bool {
	if m != nil {
		return m.Test
	}
	return false
}

func (m *Order) GetTotalPrice() int64 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *Order) GetSubtotalPrice() int64 {
	if m != nil {
		return m.SubtotalPrice
	}
	return 0
}

func (m *Order) GetTotalWeight() int64 {
	if m != nil {
		return m.TotalWeight
	}
	return 0
}

func (m *Order) GetTotalTax() int64 {
	if m != nil {
		return m.TotalTax
	}
	return 0
}

func (m *Order) GetTaxesIncluded() bool {
	if m != nil {
		return m.TaxesIncluded
	}
	return false
}

func (m *Order) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Order) GetFinancialStatus() string {
	if m != nil {
		return m.FinancialStatus
	}
	return ""
}

func (m *Order) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *Order) GetTotalPriceUsd() int64 {
	if m != nil {
		return m.TotalPriceUsd
	}
	return 0
}

func (m *Order) GetNoteAttributes() []*NoteAttribute {
	if m != nil {
		return m.NoteAttributes
	}
	return nil
}

// ListOrdersRequest
//
// All filter fields applied with AND operator.
// Zero value in filter (by default) means no filtering.
//
type ListOrdersRequest struct {
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The shop_id adds filtering by Shopify's shop.
	ShopId int64 `protobuf:"varint,3,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// The folder_id adds filtering by folder. Optional.
	FolderId             int64    `protobuf:"varint,4,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOrdersRequest) Reset()         { *m = ListOrdersRequest{} }
func (m *ListOrdersRequest) String() string { return proto.CompactTextString(m) }
func (*ListOrdersRequest) ProtoMessage()    {}
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d70621d7f7645dc1, []int{1}
}

func (m *ListOrdersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrdersRequest.Unmarshal(m, b)
}
func (m *ListOrdersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrdersRequest.Marshal(b, m, deterministic)
}
func (m *ListOrdersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrdersRequest.Merge(m, src)
}
func (m *ListOrdersRequest) XXX_Size() int {
	return xxx_messageInfo_ListOrdersRequest.Size(m)
}
func (m *ListOrdersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrdersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrdersRequest proto.InternalMessageInfo

func (m *ListOrdersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListOrdersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListOrdersRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *ListOrdersRequest) GetFolderId() int64 {
	if m != nil {
		return m.FolderId
	}
	return 0
}

type ListOrdersResponse struct {
	// There will be a maximum number of items returned based on the page_size
	// field in the request.
	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOrdersResponse) Reset()         { *m = ListOrdersResponse{} }
func (m *ListOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*ListOrdersResponse) ProtoMessage()    {}
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d70621d7f7645dc1, []int{2}
}

func (m *ListOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrdersResponse.Unmarshal(m, b)
}
func (m *ListOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrdersResponse.Marshal(b, m, deterministic)
}
func (m *ListOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrdersResponse.Merge(m, src)
}
func (m *ListOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_ListOrdersResponse.Size(m)
}
func (m *ListOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrdersResponse proto.InternalMessageInfo

func (m *ListOrdersResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *ListOrdersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetOrderRequest struct {
	// The Shopify Store.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// The field will contain id of order.
	OrderId              int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderRequest) Reset()         { *m = GetOrderRequest{} }
func (m *GetOrderRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderRequest) ProtoMessage()    {}
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d70621d7f7645dc1, []int{3}
}

func (m *GetOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderRequest.Unmarshal(m, b)
}
func (m *GetOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderRequest.Merge(m, src)
}
func (m *GetOrderRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderRequest.Size(m)
}
func (m *GetOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderRequest proto.InternalMessageInfo

func (m *GetOrderRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *GetOrderRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type CreateOrderRequest struct {
	// The order resource to create.
	Order                *Order   `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d70621d7f7645dc1, []int{4}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type UpdateOrderRequest struct {
	// The order resource which replaces the resource on the server.
	Order *Order `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateOrderRequest) Reset()         { *m = UpdateOrderRequest{} }
func (m *UpdateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateOrderRequest) ProtoMessage()    {}
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d70621d7f7645dc1, []int{5}
}

func (m *UpdateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOrderRequest.Unmarshal(m, b)
}
func (m *UpdateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOrderRequest.Marshal(b, m, deterministic)
}
func (m *UpdateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOrderRequest.Merge(m, src)
}
func (m *UpdateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateOrderRequest.Size(m)
}
func (m *UpdateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOrderRequest proto.InternalMessageInfo

func (m *UpdateOrderRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *UpdateOrderRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteOrderRequest struct {
	// The Shopify Store.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// The resource id of the order to be deleted.
	OrderId              int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOrderRequest) Reset()         { *m = DeleteOrderRequest{} }
func (m *DeleteOrderRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteOrderRequest) ProtoMessage()    {}
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d70621d7f7645dc1, []int{6}
}

func (m *DeleteOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOrderRequest.Unmarshal(m, b)
}
func (m *DeleteOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOrderRequest.Marshal(b, m, deterministic)
}
func (m *DeleteOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOrderRequest.Merge(m, src)
}
func (m *DeleteOrderRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteOrderRequest.Size(m)
}
func (m *DeleteOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOrderRequest proto.InternalMessageInfo

func (m *DeleteOrderRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *DeleteOrderRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "Order")
	proto.RegisterType((*ListOrdersRequest)(nil), "ListOrdersRequest")
	proto.RegisterType((*ListOrdersResponse)(nil), "ListOrdersResponse")
	proto.RegisterType((*GetOrderRequest)(nil), "GetOrderRequest")
	proto.RegisterType((*CreateOrderRequest)(nil), "CreateOrderRequest")
	proto.RegisterType((*UpdateOrderRequest)(nil), "UpdateOrderRequest")
	proto.RegisterType((*DeleteOrderRequest)(nil), "DeleteOrderRequest")
}

func init() { proto.RegisterFile("orders/order.proto", fileDescriptor_d70621d7f7645dc1) }

var fileDescriptor_d70621d7f7645dc1 = []byte{
	// 851 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x6f, 0x1b, 0x45,
	0x10, 0x96, 0xe3, 0xda, 0x39, 0x8f, 0x63, 0xbb, 0xac, 0x0a, 0x5d, 0x92, 0x42, 0x5d, 0x23, 0xc0,
	0xbc, 0x38, 0x52, 0x2a, 0x54, 0x21, 0xc4, 0x83, 0x4b, 0xd3, 0xd6, 0xd0, 0xa6, 0xd1, 0x35, 0x15,
	0x12, 0x42, 0x3a, 0xad, 0x6f, 0xc7, 0xce, 0x2a, 0x77, 0xb7, 0xc7, 0xed, 0x9e, 0x08, 0x7d, 0xe7,
	0xcf, 0xe5, 0x7f, 0x40, 0x3b, 0xbb, 0xe7, 0x38, 0x01, 0x14, 0xc4, 0x93, 0x77, 0xbe, 0xf9, 0xe6,
	0xe7, 0xae, 0xbf, 0x03, 0xa6, 0x2b, 0x89, 0x95, 0x39, 0xa4, 0x9f, 0x59, 0x59, 0x69, 0xab, 0xf7,
	0xc7, 0x6b, 0xad, 0xd7, 0x19, 0x1e, 0x92, 0xb5, 0xac, 0x57, 0x87, 0x2b, 0x85, 0x99, 0x4c, 0x72,
	0x61, 0x2e, 0x02, 0xe3, 0xe1, 0x4d, 0x86, 0x55, 0x39, 0x1a, 0x2b, 0xf2, 0x32, 0x10, 0xee, 0x85,
	0xb4, 0x42, 0xca, 0x0a, 0x8d, 0x09, 0xe8, 0x41, 0x40, 0x0b, 0x6d, 0x31, 0x11, 0xd6, 0x56, 0x6a,
	0x59, 0x5b, 0xf4, 0xce, 0xc9, 0x9f, 0x11, 0x74, 0xde, 0x38, 0x3f, 0xbb, 0x0f, 0xbb, 0xe6, 0x5c,
	0x97, 0x89, 0x92, 0xbc, 0x35, 0x6e, 0x4d, 0xdb, 0x71, 0xd7, 0x99, 0x0b, 0xc9, 0x86, 0xb0, 0xa3,
	0x24, 0xdf, 0x21, 0x6c, 0x47, 0x49, 0xf6, 0x08, 0xf6, 0x28, 0x63, 0x52, 0xd4, 0xf9, 0x12, 0x2b,
	0xde, 0x26, 0x4f, 0x9f, 0xb0, 0x13, 0x82, 0xd8, 0x3d, 0xe8, 0x60, 0x2e, 0x54, 0xc6, 0x3b, 0xe3,
	0xd6, 0xb4, 0x17, 0x7b, 0x83, 0x7d, 0x04, 0xdd, 0x10, 0xd2, 0xf5, 0x05, 0xbc, 0xc5, 0x18, 0xdc,
	0x71, 0xbd, 0xf1, 0x5d, 0x22, 0xd3, 0xd9, 0x65, 0xb0, 0xfa, 0x02, 0x0b, 0x1e, 0xf9, 0x0c, 0x64,
	0xb0, 0x6f, 0x00, 0xd2, 0x0a, 0x85, 0x45, 0x99, 0x08, 0xcb, 0x61, 0xdc, 0x9a, 0xf6, 0x8f, 0xf6,
	0x67, 0x7e, 0x2d, 0xb3, 0x66, 0x2d, 0xb3, 0xb3, 0x66, 0x2d, 0x71, 0x2f, 0xb0, 0xe7, 0x96, 0x7d,
	0x07, 0x7b, 0x65, 0xa5, 0x53, 0x34, 0xc6, 0x07, 0xf7, 0x6f, 0x0d, 0xee, 0x6f, 0xf8, 0x73, 0xeb,
	0x2a, 0xd7, 0xa5, 0x6c, 0x2a, 0xef, 0xdd, 0x5e, 0x39, 0xb0, 0xe7, 0x96, 0x3d, 0x81, 0x5e, 0x9a,
	0xe9, 0x50, 0x76, 0x70, 0x6b, 0x64, 0xe4, 0xc9, 0xbe, 0xe5, 0x54, 0x14, 0x29, 0x66, 0x99, 0x8f,
	0x1d, 0xde, 0xde, 0xf2, 0x86, 0x3f, 0xb7, 0xb4, 0x56, 0x91, 0x23, 0xbf, 0x1f, 0xd6, 0x2a, 0x72,
	0x64, 0x9f, 0xc1, 0xc0, 0x53, 0x92, 0x0a, 0x85, 0xd1, 0x05, 0x3f, 0x20, 0x67, 0xa8, 0x13, 0x13,
	0xc6, 0x3e, 0x84, 0xae, 0x28, 0xe9, 0x21, 0x7c, 0x49, 0xf7, 0xd4, 0x11, 0xa5, 0x7b, 0x07, 0x2e,
	0x56, 0x17, 0x56, 0xa4, 0x36, 0xf1, 0x97, 0xfb, 0x75, 0x88, 0xf5, 0xe0, 0x31, 0xdd, 0xf1, 0x63,
	0xb8, 0x6b, 0xce, 0x55, 0x59, 0xaa, 0x62, 0x9d, 0x84, 0x67, 0xc8, 0xe7, 0xd4, 0x77, 0x34, 0x9b,
	0x7b, 0x3b, 0x1e, 0x35, 0x8c, 0x00, 0xb0, 0x09, 0x0c, 0x32, 0x55, 0x60, 0xa2, 0x2c, 0xe6, 0x89,
	0x92, 0x86, 0x3f, 0x1d, 0xb7, 0xdd, 0x93, 0x72, 0xe0, 0xc2, 0x62, 0xbe, 0x90, 0x86, 0x3d, 0x84,
	0x7e, 0x5a, 0x1b, 0xab, 0x73, 0xac, 0x5c, 0x67, 0xcf, 0xa8, 0x33, 0x68, 0xa0, 0x85, 0x64, 0x07,
	0xd0, 0x5b, 0xe9, 0x4c, 0x7a, 0xf7, 0xf1, 0xb8, 0x35, 0xed, 0xc4, 0x91, 0x07, 0x16, 0xd2, 0xed,
	0xc2, 0xa2, 0xb1, 0xfc, 0xf9, 0xb8, 0x35, 0x8d, 0x62, 0x3a, 0xbb, 0x8c, 0x56, 0x5b, 0x91, 0x25,
	0x65, 0xa5, 0x52, 0xe4, 0x2f, 0x7c, 0x46, 0x82, 0x4e, 0x1d, 0xc2, 0x3e, 0x87, 0xa1, 0xa9, 0x97,
	0xdb, 0x9c, 0x97, 0xc4, 0x19, 0x34, 0xa8, 0xa7, 0x3d, 0x82, 0x3d, 0xcf, 0xf9, 0x0d, 0xd5, 0xfa,
	0xdc, 0xf2, 0x85, 0xff, 0x3f, 0x10, 0xf6, 0x13, 0x41, 0xae, 0x37, 0x4f, 0xb1, 0xe2, 0x92, 0xff,
	0x40, 0xfe, 0x88, 0x80, 0x33, 0x71, 0xe9, 0xca, 0x58, 0x71, 0x89, 0x26, 0x51, 0x45, 0x9a, 0xd5,
	0x12, 0x25, 0xff, 0x91, 0xba, 0x1c, 0x10, 0xba, 0x08, 0x20, 0xdb, 0x87, 0x28, 0xad, 0xab, 0x0a,
	0x8b, 0xf4, 0x77, 0xfe, 0x8a, 0x36, 0xbf, 0xb1, 0xd9, 0x57, 0x70, 0x77, 0xa5, 0x0a, 0x51, 0xa4,
	0x4a, 0x64, 0x89, 0xb1, 0xc2, 0xd6, 0x86, 0xbf, 0x26, 0xce, 0x68, 0x83, 0xbf, 0x25, 0x98, 0x3d,
	0x80, 0x5e, 0xaa, 0x8b, 0x95, 0xaa, 0x72, 0x94, 0xfc, 0x84, 0x0a, 0x5d, 0x01, 0xec, 0x0b, 0x18,
	0x6d, 0xcd, 0x9b, 0xd4, 0x46, 0xf2, 0x37, 0x7e, 0xe6, 0xab, 0x81, 0xdf, 0x19, 0xc9, 0x9e, 0xc0,
	0xe8, 0xba, 0x9c, 0x18, 0x7e, 0x3a, 0x6e, 0x4f, 0xfb, 0x47, 0xc3, 0xd9, 0x89, 0xb6, 0x38, 0x6f,
	0xe0, 0x78, 0x58, 0x6c, 0x9b, 0x66, 0xf2, 0x47, 0x0b, 0x3e, 0x78, 0xa5, 0x8c, 0x25, 0xcd, 0x31,
	0x31, 0xfe, 0x5a, 0xbb, 0xab, 0x38, 0x80, 0x5e, 0x29, 0xd6, 0x98, 0x18, 0xf5, 0x1e, 0x49, 0x7d,
	0x3a, 0x71, 0xe4, 0x80, 0xb7, 0xea, 0x3d, 0xb2, 0x4f, 0x00, 0xc8, 0xe9, 0xf5, 0x60, 0x87, 0xc6,
	0x22, 0xfa, 0x19, 0x69, 0xc2, 0x96, 0x6e, 0xb5, 0xaf, 0xe9, 0xd6, 0xb5, 0x07, 0x71, 0xc7, 0x2f,
	0xbd, 0x79, 0x10, 0x93, 0x5f, 0x80, 0x6d, 0xb7, 0x61, 0x4a, 0x5d, 0x18, 0x64, 0x9f, 0x42, 0xd7,
	0x8b, 0x25, 0x6f, 0xd1, 0x34, 0xdd, 0x19, 0x11, 0xe2, 0x80, 0xba, 0xf5, 0x14, 0x78, 0x69, 0x93,
	0xbf, 0xf5, 0x33, 0x70, 0xf0, 0x69, 0xd3, 0xd3, 0xe4, 0x18, 0x46, 0x2f, 0xd0, 0x27, 0x6f, 0x46,
	0xfc, 0x57, 0x79, 0xfd, 0x18, 0x22, 0x2f, 0xa7, 0x1b, 0x91, 0xdd, 0x25, 0x7b, 0x21, 0x27, 0x47,
	0xc0, 0xbe, 0x27, 0x01, 0xbb, 0x96, 0xe9, 0x01, 0x74, 0x88, 0x40, 0xec, 0xab, 0x1e, 0x3d, 0x38,
	0xd1, 0xc0, 0xde, 0x91, 0xf4, 0xfc, 0xf7, 0x18, 0xf6, 0x2d, 0xf4, 0xbd, 0x5c, 0xd1, 0xd7, 0x86,
	0xd6, 0xf8, 0x4f, 0x3a, 0xf3, 0xdc, 0x7d, 0x90, 0x5e, 0x0b, 0x73, 0x11, 0x07, 0x2d, 0x74, 0xe7,
	0xc9, 0x4b, 0x60, 0xcf, 0x30, 0xc3, 0x1b, 0x05, 0xff, 0xc7, 0xb8, 0x4f, 0xa3, 0x9f, 0xc3, 0x9e,
	0x97, 0x5d, 0xaa, 0xf9, 0xf8, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xed, 0x37, 0xde, 0x28,
	0x07, 0x00, 0x00,
}
