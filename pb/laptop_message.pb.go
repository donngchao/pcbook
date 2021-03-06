// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamppb "google.golang.org/protobuf/types/known/timestamppb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight               isLaptop_Weight        `protobuf_oneof:"weight"`
	PriceUsd             float64                `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32                 `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_laptop_message_bc47bbb28131d11a, []int{0}
}
func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (dst *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(dst, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Laptop) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Laptop_OneofMarshaler, _Laptop_OneofUnmarshaler, _Laptop_OneofSizer, []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

func _Laptop_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Laptop)
	// weight
	switch x := m.Weight.(type) {
	case *Laptop_WeightKg:
		b.EncodeVarint(10<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.WeightKg))
	case *Laptop_WeightLb:
		b.EncodeVarint(11<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.WeightLb))
	case nil:
	default:
		return fmt.Errorf("Laptop.Weight has unexpected type %T", x)
	}
	return nil
}

func _Laptop_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Laptop)
	switch tag {
	case 10: // weight.weight_kg
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Weight = &Laptop_WeightKg{math.Float64frombits(x)}
		return true, err
	case 11: // weight.weight_lb
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Weight = &Laptop_WeightLb{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Laptop_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Laptop)
	// weight
	switch x := m.Weight.(type) {
	case *Laptop_WeightKg:
		n += 1 // tag and wire
		n += 8
	case *Laptop_WeightLb:
		n += 1 // tag and wire
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Laptop)(nil), "pcbook.Laptop")
}

func init() {
	proto.RegisterFile("laptop_message.proto", fileDescriptor_laptop_message_bc47bbb28131d11a)
}

var fileDescriptor_laptop_message_bc47bbb28131d11a = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x80, 0x49, 0xd3, 0x85, 0xe4, 0x75, 0x2b, 0xc8, 0x2a, 0x60, 0x15, 0x4d, 0x0b, 0x1c, 0x50,
	0x25, 0x50, 0x26, 0xc1, 0x89, 0x23, 0xe3, 0x00, 0xd2, 0x86, 0x34, 0x79, 0xf4, 0x00, 0x97, 0xc8,
	0x89, 0x1f, 0x26, 0x6a, 0x53, 0x5b, 0xb6, 0x23, 0xd4, 0xbf, 0xcc, 0xaf, 0x98, 0x6a, 0x27, 0xd3,
	0x9a, 0x5b, 0xfc, 0x7d, 0x9f, 0x12, 0xbf, 0x3c, 0x58, 0x6c, 0xb9, 0x76, 0x4a, 0x97, 0x2d, 0x5a,
	0xcb, 0x25, 0x16, 0xda, 0x28, 0xa7, 0x48, 0xa2, 0xeb, 0x4a, 0xa9, 0xcd, 0xf2, 0xe5, 0x06, 0xf7,
	0x95, 0xe2, 0x46, 0x1c, 0xfb, 0xe5, 0xa2, 0xc5, 0x56, 0x99, 0xfd, 0x88, 0xbe, 0xd2, 0x46, 0xd5,
	0x68, 0xad, 0x32, 0xe3, 0xdc, 0xd6, 0x06, 0x71, 0x37, 0xa2, 0x2f, 0xac, 0x53, 0x86, 0x4b, 0x1c,
	0xe1, 0x0b, 0xa9, 0x94, 0xdc, 0xe2, 0xa5, 0x3f, 0x55, 0xdd, 0x9f, 0x4b, 0xd7, 0xb4, 0x68, 0x1d,
	0x6f, 0x75, 0x08, 0xde, 0xfe, 0x8f, 0x21, 0xb9, 0xf1, 0xb7, 0x26, 0x73, 0x98, 0x34, 0x82, 0x46,
	0x79, 0xb4, 0xca, 0xd8, 0xa4, 0x11, 0x64, 0x01, 0x27, 0x95, 0xe1, 0x3b, 0x41, 0x27, 0x1e, 0x85,
	0x03, 0x21, 0x30, 0xdd, 0xf1, 0x16, 0x69, 0xec, 0xa1, 0x7f, 0x26, 0xe7, 0x10, 0xd7, 0xba, 0xa3,
	0xd3, 0x3c, 0x5a, 0xcd, 0x3e, 0xce, 0x8a, 0x30, 0x6f, 0xf1, 0xf5, 0x76, 0xcd, 0x0e, 0x9c, 0xe4,
	0x10, 0x1b, 0xde, 0xd2, 0x13, 0xaf, 0xe7, 0x83, 0xfe, 0xe1, 0xa7, 0x66, 0x07, 0x45, 0x2e, 0x60,
	0x2a, 0x75, 0x67, 0x69, 0x92, 0xc7, 0x8f, 0xdf, 0xf0, 0xed, 0x76, 0xcd, 0xbc, 0x20, 0xef, 0x21,
	0xed, 0x07, 0xb4, 0xf4, 0xa9, 0x8f, 0x9e, 0x0d, 0xd1, 0x5d, 0xe0, 0xec, 0x21, 0x20, 0xef, 0x20,
	0x09, 0xff, 0x88, 0xa6, 0xc7, 0x9f, 0xbc, 0xf3, 0x94, 0xf5, 0x96, 0x7c, 0x80, 0x74, 0x58, 0x09,
	0xcd, 0x7c, 0xf9, 0x7c, 0x28, 0xaf, 0x7b, 0xce, 0x1e, 0x0a, 0x72, 0x0e, 0xd9, 0x3f, 0x6c, 0xe4,
	0x5f, 0x57, 0x6e, 0x24, 0x85, 0x3c, 0x5a, 0x45, 0xdf, 0x9f, 0xb0, 0x34, 0xa0, 0x6b, 0xf9, 0x48,
	0x6f, 0x2b, 0x3a, 0x3b, 0xd6, 0x37, 0x15, 0x79, 0x0d, 0x99, 0x36, 0x4d, 0x8d, 0x65, 0x67, 0x05,
	0x3d, 0x3d, 0x68, 0x96, 0x7a, 0xb0, 0xb6, 0x82, 0xbc, 0x81, 0x53, 0x83, 0x5b, 0xe4, 0x16, 0xcb,
	0x3d, 0x72, 0x43, 0xcf, 0xf2, 0x68, 0x75, 0xc6, 0x66, 0x3d, 0xfb, 0x85, 0xdc, 0x90, 0xcf, 0x00,
	0x9d, 0x16, 0xdc, 0xa1, 0x28, 0xb9, 0xa3, 0x73, 0x7f, 0xdb, 0x65, 0x11, 0xb6, 0x5b, 0x0c, 0xdb,
	0x2d, 0x7e, 0x0e, 0xdb, 0x65, 0x59, 0x5f, 0x7f, 0x71, 0x57, 0x29, 0x24, 0xe1, 0x1a, 0x57, 0xd3,
	0xdf, 0x13, 0x5d, 0x55, 0x89, 0xcf, 0x3f, 0xdd, 0x07, 0x00, 0x00, 0xff, 0xff, 0xce, 0x92, 0xee,
	0x98, 0xae, 0x02, 0x00, 0x00,
}
