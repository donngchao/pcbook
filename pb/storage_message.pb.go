// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Storage_Driver int32

const (
	Storage_UNKNOWN Storage_Driver = 0
	Storage_HDD     Storage_Driver = 1
	Storage_SSD     Storage_Driver = 2
)

var Storage_Driver_name = map[int32]string{
	0: "UNKNOWN",
	1: "HDD",
	2: "SSD",
}
var Storage_Driver_value = map[string]int32{
	"UNKNOWN": 0,
	"HDD":     1,
	"SSD":     2,
}

func (x Storage_Driver) String() string {
	return proto.EnumName(Storage_Driver_name, int32(x))
}
func (Storage_Driver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_storage_message_c3f038c9ba5881dc, []int{0, 0}
}

type Storage struct {
	Driver               Storage_Driver `protobuf:"varint,1,opt,name=driver,proto3,enum=pcbook.Storage_Driver" json:"driver,omitempty"`
	Memory               *Memory        `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_message_c3f038c9ba5881dc, []int{0}
}
func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (dst *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(dst, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetDriver() Storage_Driver {
	if m != nil {
		return m.Driver
	}
	return Storage_UNKNOWN
}

func (m *Storage) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*Storage)(nil), "pcbook.Storage")
	proto.RegisterEnum("pcbook.Storage_Driver", Storage_Driver_name, Storage_Driver_value)
}

func init() {
	proto.RegisterFile("storage_message.proto", fileDescriptor_storage_message_c3f038c9ba5881dc)
}

var fileDescriptor_storage_message_c3f038c9ba5881dc = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0x8d, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x2b, 0x48, 0x4e, 0xca, 0xcf, 0xcf, 0x96, 0x12, 0xc9, 0x4d, 0xcd, 0xcd, 0x2f,
	0xaa, 0x44, 0x95, 0x55, 0xea, 0x62, 0xe4, 0x62, 0x0f, 0x86, 0xe8, 0x13, 0xd2, 0xe3, 0x62, 0x4b,
	0x29, 0xca, 0x2c, 0x4b, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x12, 0xd3, 0x83, 0x68,
	0xd5, 0x83, 0x2a, 0xd0, 0x73, 0x01, 0xcb, 0x06, 0x41, 0x55, 0x09, 0xa9, 0x71, 0xb1, 0x41, 0xcc,
	0x94, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0x83, 0xa9, 0xf7, 0x05, 0x8b, 0x06, 0x41, 0x65,
	0x95, 0xd4, 0xb9, 0xd8, 0x20, 0x3a, 0x85, 0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3,
	0xfd, 0x04, 0x18, 0x84, 0xd8, 0xb9, 0x98, 0x3d, 0x5c, 0x5c, 0x04, 0x18, 0x41, 0x8c, 0xe0, 0x60,
	0x17, 0x01, 0x26, 0x27, 0x96, 0x28, 0xa6, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0xcb, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb5, 0xac, 0xb1, 0xba, 0xd0, 0x00, 0x00, 0x00,
}
