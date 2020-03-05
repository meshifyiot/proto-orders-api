// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders/addresses.proto

package orders

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Address struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Company              string   `protobuf:"bytes,4,opt,name=company,proto3" json:"company,omitempty"`
	FirstName            string   `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address1             string   `protobuf:"bytes,7,opt,name=address1,proto3" json:"address1,omitempty"`
	Address2             string   `protobuf:"bytes,8,opt,name=address2,proto3" json:"address2,omitempty"`
	City                 string   `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	ProvinceCode         string   `protobuf:"bytes,10,opt,name=province_code,json=provinceCode,proto3" json:"province_code,omitempty"`
	Province             string   `protobuf:"bytes,11,opt,name=province,proto3" json:"province,omitempty"`
	Zip                  string   `protobuf:"bytes,12,opt,name=zip,proto3" json:"zip,omitempty"`
	CountryCode          string   `protobuf:"bytes,13,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Country              string   `protobuf:"bytes,14,opt,name=country,proto3" json:"country,omitempty"`
	Latitude             float64  `protobuf:"fixed64,15,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,16,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_befdb0e26035decf, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Address) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Address) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Address) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Address) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Address) GetAddress1() string {
	if m != nil {
		return m.Address1
	}
	return ""
}

func (m *Address) GetAddress2() string {
	if m != nil {
		return m.Address2
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetProvinceCode() string {
	if m != nil {
		return m.ProvinceCode
	}
	return ""
}

func (m *Address) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *Address) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Address) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*Address)(nil), "Address")
}

func init() { proto.RegisterFile("orders/addresses.proto", fileDescriptor_befdb0e26035decf) }

var fileDescriptor_befdb0e26035decf = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0x59, 0xb7, 0xdd, 0x8f, 0xe9, 0x56, 0x4b, 0x10, 0x19, 0xfc, 0x80, 0xaa, 0x97, 0x9e,
	0x14, 0xeb, 0x2f, 0x50, 0xef, 0x1e, 0x7a, 0xf4, 0x52, 0xd6, 0x4d, 0xd4, 0x85, 0x6d, 0x66, 0x49,
	0x52, 0x61, 0xfd, 0x33, 0xfe, 0x55, 0xc9, 0xa4, 0x51, 0x7b, 0x9b, 0xf7, 0x79, 0xe0, 0xcd, 0x90,
	0x81, 0x13, 0x32, 0x52, 0x19, 0x7b, 0x5b, 0x4b, 0x69, 0x94, 0xb5, 0xca, 0xde, 0xf4, 0x86, 0x1c,
	0x5d, 0x7d, 0xa7, 0x90, 0x3f, 0x04, 0x26, 0x04, 0x8c, 0x74, 0xbd, 0x51, 0x78, 0x30, 0x4f, 0x16,
	0xe5, 0x8a, 0x67, 0x71, 0x0c, 0xe3, 0xfe, 0x83, 0xb4, 0xc2, 0x94, 0x61, 0x08, 0x02, 0x21, 0x6f,
	0x68, 0xd3, 0xd7, 0x7a, 0xc0, 0x11, 0xf3, 0x18, 0xc5, 0x05, 0xc0, 0x5b, 0x6b, 0xac, 0x5b, 0x73,
	0xd3, 0x98, 0x65, 0xc9, 0xe4, 0xd9, 0xd7, 0x9d, 0x41, 0xd9, 0xd5, 0xd1, 0x66, 0x6c, 0x0b, 0x0f,
	0x58, 0x9e, 0x42, 0xb1, 0x5b, 0xef, 0x0e, 0xf3, 0xe0, 0x62, 0xfe, 0xe7, 0x96, 0x58, 0xec, 0xb9,
	0xa5, 0xdf, 0xbb, 0x69, 0xdd, 0x80, 0x65, 0xd8, 0xdb, 0xcf, 0xe2, 0x1a, 0xa6, 0xbd, 0xa1, 0xcf,
	0x56, 0x37, 0x6a, 0xdd, 0x90, 0x54, 0x08, 0x2c, 0xab, 0x08, 0x9f, 0x48, 0xf2, 0x83, 0x31, 0xe3,
	0x24, 0x94, 0xc6, 0x2c, 0x66, 0x90, 0x7e, 0xb5, 0x3d, 0x56, 0x8c, 0xfd, 0x28, 0x2e, 0xa1, 0x6a,
	0x68, 0xab, 0x9d, 0x19, 0x42, 0xe3, 0x94, 0xd5, 0x64, 0xc7, 0xb8, 0x90, 0xff, 0x85, 0x23, 0x1e,
	0xc6, 0x7f, 0xe1, 0xe8, 0x9f, 0xea, 0x6a, 0xd7, 0xba, 0xad, 0x54, 0x78, 0x34, 0x4f, 0x16, 0xc9,
	0xea, 0x37, 0x8b, 0x73, 0x28, 0x3b, 0xd2, 0xef, 0x41, 0xce, 0x58, 0xfe, 0x81, 0xc7, 0xe2, 0x25,
	0x0b, 0xb7, 0x7b, 0xcd, 0xf8, 0x64, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x07, 0x8d,
	0x26, 0xcc, 0x01, 0x00, 0x00,
}