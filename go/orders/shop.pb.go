// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: orders/shop.proto

package orders

import (
	proto "github.com/golang/protobuf/proto"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The shop ID.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The domain for Shopify Store
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// The name of Shopify Store
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_orders_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_orders_shop_proto_rawDescGZIP(), []int{0}
}

func (x *Shop) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Shop) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ListShopsRequest
//
// All filter fields applied with AND operator.
// Zero value in filter (by default) means no filtering.
//
type ListShopsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListShopsRequest) Reset() {
	*x = ListShopsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShopsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShopsRequest) ProtoMessage() {}

func (x *ListShopsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShopsRequest.ProtoReflect.Descriptor instead.
func (*ListShopsRequest) Descriptor() ([]byte, []int) {
	return file_orders_shop_proto_rawDescGZIP(), []int{1}
}

func (x *ListShopsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListShopsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListShopsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// There will be a maximum number of items returned based on the page_size
	// field in the request.
	Shops []*Shop `protobuf:"bytes,1,rep,name=shops,proto3" json:"shops,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListShopsResponse) Reset() {
	*x = ListShopsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_shop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShopsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShopsResponse) ProtoMessage() {}

func (x *ListShopsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_shop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShopsResponse.ProtoReflect.Descriptor instead.
func (*ListShopsResponse) Descriptor() ([]byte, []int) {
	return file_orders_shop_proto_rawDescGZIP(), []int{2}
}

func (x *ListShopsResponse) GetShops() []*Shop {
	if x != nil {
		return x.Shops
	}
	return nil
}

func (x *ListShopsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field will contain id of shop.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *GetShopRequest) Reset() {
	*x = GetShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_shop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopRequest) ProtoMessage() {}

func (x *GetShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_shop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopRequest.ProtoReflect.Descriptor instead.
func (*GetShopRequest) Descriptor() ([]byte, []int) {
	return file_orders_shop_proto_rawDescGZIP(), []int{3}
}

func (x *GetShopRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type GetShopByDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field will contain domain of shop.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *GetShopByDomainRequest) Reset() {
	*x = GetShopByDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_shop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShopByDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopByDomainRequest) ProtoMessage() {}

func (x *GetShopByDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_shop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopByDomainRequest.ProtoReflect.Descriptor instead.
func (*GetShopByDomainRequest) Descriptor() ([]byte, []int) {
	return file_orders_shop_proto_rawDescGZIP(), []int{4}
}

func (x *GetShopByDomainRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type CreateShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The shop resource to create.
	Shop *Shop `protobuf:"bytes,3,opt,name=shop,proto3" json:"shop,omitempty"`
}

func (x *CreateShopRequest) Reset() {
	*x = CreateShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_shop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShopRequest) ProtoMessage() {}

func (x *CreateShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_shop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShopRequest.ProtoReflect.Descriptor instead.
func (*CreateShopRequest) Descriptor() ([]byte, []int) {
	return file_orders_shop_proto_rawDescGZIP(), []int{5}
}

func (x *CreateShopRequest) GetShop() *Shop {
	if x != nil {
		return x.Shop
	}
	return nil
}

type UpdateShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The shop resource which replaces the resource on the server.
	Shop *Shop `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateShopRequest) Reset() {
	*x = UpdateShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_shop_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShopRequest) ProtoMessage() {}

func (x *UpdateShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_shop_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShopRequest.ProtoReflect.Descriptor instead.
func (*UpdateShopRequest) Descriptor() ([]byte, []int) {
	return file_orders_shop_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateShopRequest) GetShop() *Shop {
	if x != nil {
		return x.Shop
	}
	return nil
}

func (x *UpdateShopRequest) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource id of the shop to be deleted.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *DeleteShopRequest) Reset() {
	*x = DeleteShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_shop_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShopRequest) ProtoMessage() {}

func (x *DeleteShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_shop_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShopRequest.ProtoReflect.Descriptor instead.
func (*DeleteShopRequest) Descriptor() ([]byte, []int) {
	return file_orders_shop_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteShopRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

var File_orders_shop_proto protoreflect.FileDescriptor

var file_orders_shop_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x58, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x30,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x22, 0x6b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x2c, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x66,
	0x79, 0x69, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_shop_proto_rawDescOnce sync.Once
	file_orders_shop_proto_rawDescData = file_orders_shop_proto_rawDesc
)

func file_orders_shop_proto_rawDescGZIP() []byte {
	file_orders_shop_proto_rawDescOnce.Do(func() {
		file_orders_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_shop_proto_rawDescData)
	})
	return file_orders_shop_proto_rawDescData
}

var file_orders_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_orders_shop_proto_goTypes = []interface{}{
	(*Shop)(nil),                   // 0: Shop
	(*ListShopsRequest)(nil),       // 1: ListShopsRequest
	(*ListShopsResponse)(nil),      // 2: ListShopsResponse
	(*GetShopRequest)(nil),         // 3: GetShopRequest
	(*GetShopByDomainRequest)(nil), // 4: GetShopByDomainRequest
	(*CreateShopRequest)(nil),      // 5: CreateShopRequest
	(*UpdateShopRequest)(nil),      // 6: UpdateShopRequest
	(*DeleteShopRequest)(nil),      // 7: DeleteShopRequest
	(*field_mask.FieldMask)(nil),   // 8: google.protobuf.FieldMask
}
var file_orders_shop_proto_depIdxs = []int32{
	0, // 0: ListShopsResponse.shops:type_name -> Shop
	0, // 1: CreateShopRequest.shop:type_name -> Shop
	0, // 2: UpdateShopRequest.shop:type_name -> Shop
	8, // 3: UpdateShopRequest.update_mask:type_name -> google.protobuf.FieldMask
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_orders_shop_proto_init() }
func file_orders_shop_proto_init() {
	if File_orders_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
		file_orders_shop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShopsRequest); i {
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
		file_orders_shop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShopsResponse); i {
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
		file_orders_shop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShopRequest); i {
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
		file_orders_shop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShopByDomainRequest); i {
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
		file_orders_shop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShopRequest); i {
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
		file_orders_shop_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateShopRequest); i {
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
		file_orders_shop_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShopRequest); i {
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
			RawDescriptor: file_orders_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orders_shop_proto_goTypes,
		DependencyIndexes: file_orders_shop_proto_depIdxs,
		MessageInfos:      file_orders_shop_proto_msgTypes,
	}.Build()
	File_orders_shop_proto = out.File
	file_orders_shop_proto_rawDesc = nil
	file_orders_shop_proto_goTypes = nil
	file_orders_shop_proto_depIdxs = nil
}
