// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: orders/address.proto

package orders

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Shopify Store to which address record belongs to.
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	// The full name of the person associated with the payment method.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The phone number at the shipping address.
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	// The company of the person associated with the shipping address.
	Company string `protobuf:"bytes,4,opt,name=company,proto3" json:"company,omitempty"`
	// The first name of the person associated with the shipping address.
	FirstName string `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// The last name of the person associated with the shipping address.
	LastName string `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// The street address of the shipping address.
	Address1 string `protobuf:"bytes,7,opt,name=address1,proto3" json:"address1,omitempty"`
	// An optional additional field for the street address of the
	// shipping address.
	Address2 string `protobuf:"bytes,8,opt,name=address2,proto3" json:"address2,omitempty"`
	// The city, town, or village of the shipping address.
	City string `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	// The two-letter abbreviation of the region of the shipping address.
	ProvinceCode string `protobuf:"bytes,10,opt,name=province_code,json=provinceCode,proto3" json:"province_code,omitempty"`
	// The name of the region (province, state, prefecture, …) of the
	// shipping address.
	Province string `protobuf:"bytes,11,opt,name=province,proto3" json:"province,omitempty"`
	// The postal code (zip, postcode, Eircode, …) of the shipping address.
	Zip string `protobuf:"bytes,12,opt,name=zip,proto3" json:"zip,omitempty"`
	// The two-letter code (ISO 3166-1 format) for the country of the
	// shipping address.
	CountryCode string `protobuf:"bytes,13,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// The name of the country of the shipping address.
	Country string `protobuf:"bytes,14,opt,name=country,proto3" json:"country,omitempty"`
	// The latitude of the shipping address.
	Latitude float64 `protobuf:"fixed64,15,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The longitude of the shipping address.
	Longitude float64 `protobuf:"fixed64,16,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_orders_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_orders_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *Address) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Address) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Address) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Address) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Address) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Address) GetAddress1() string {
	if x != nil {
		return x.Address1
	}
	return ""
}

func (x *Address) GetAddress2() string {
	if x != nil {
		return x.Address2
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetProvinceCode() string {
	if x != nil {
		return x.ProvinceCode
	}
	return ""
}

func (x *Address) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Address) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *Address) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Address) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_orders_address_proto protoreflect.FileDescriptor

var file_orders_address_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x03, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x68, 0x69, 0x66, 0x79, 0x69, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_address_proto_rawDescOnce sync.Once
	file_orders_address_proto_rawDescData = file_orders_address_proto_rawDesc
)

func file_orders_address_proto_rawDescGZIP() []byte {
	file_orders_address_proto_rawDescOnce.Do(func() {
		file_orders_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_address_proto_rawDescData)
	})
	return file_orders_address_proto_rawDescData
}

var file_orders_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_orders_address_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: Address
}
var file_orders_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orders_address_proto_init() }
func file_orders_address_proto_init() {
	if File_orders_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_orders_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orders_address_proto_goTypes,
		DependencyIndexes: file_orders_address_proto_depIdxs,
		MessageInfos:      file_orders_address_proto_msgTypes,
	}.Build()
	File_orders_address_proto = out.File
	file_orders_address_proto_rawDesc = nil
	file_orders_address_proto_goTypes = nil
	file_orders_address_proto_depIdxs = nil
}
