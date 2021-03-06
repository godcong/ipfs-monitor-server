// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ManagerType ...
type ManagerType int32

// ManagerType_BootstrapAdd ...
const (
	ManagerType_BootstrapAdd    ManagerType = 0
	ManagerType_BootstrapRemove ManagerType = 1
	ManagerType_PinAdd          ManagerType = 2
	ManagerType_PinRemove       ManagerType = 3
)

// ManagerType_name ...
var ManagerType_name = map[int32]string{
	0: "BootstrapAdd",
	1: "BootstrapRemove",
	2: "PinAdd",
	3: "PinRemove",
}

// ManagerType_value ...
var ManagerType_value = map[string]int32{
	"BootstrapAdd":    0,
	"BootstrapRemove": 1,
	"PinAdd":          2,
	"PinRemove":       3,
}

// String ...
func (x ManagerType) String() string {
	return proto.EnumName(ManagerType_name, int32(x))
}

// EnumDescriptor ...
func (ManagerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{0}
}

// StartMode ...
type StartMode int32

// StartMode_Cluster ...
const (
	StartMode_Cluster StartMode = 0
	StartMode_Simple  StartMode = 1
)

// StartMode_name ...
var StartMode_name = map[int32]string{
	0: "Cluster",
	1: "Simple",
}

// StartMode_value ...
var StartMode_value = map[string]int32{
	"Cluster": 0,
	"Simple":  1,
}

// String ...
func (x StartMode) String() string {
	return proto.EnumName(StartMode_name, int32(x))
}

// EnumDescriptor ...
func (StartMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{1}
}

// MonitorType ...
type MonitorType int32

// MonitorType_Init ...
const (
	MonitorType_Init  MonitorType = 0
	MonitorType_Reset MonitorType = 1
	MonitorType_Info  MonitorType = 2
)

// MonitorType_name ...
var MonitorType_name = map[int32]string{
	0: "Init",
	1: "Reset",
	2: "Info",
}

// MonitorType_value ...
var MonitorType_value = map[string]int32{
	"Init":  0,
	"Reset": 1,
	"Info":  2,
}

// String ...
func (x MonitorType) String() string {
	return proto.EnumName(MonitorType_name, int32(x))
}

// EnumDescriptor ...
func (MonitorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{2}
}

// MonitorRequest ...
type MonitorRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *MonitorRequest) Reset() { *m = MonitorRequest{} }

// String ...
func (m *MonitorRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorRequest) ProtoMessage() {}

// Descriptor ...
func (*MonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{0}
}

// XXX_Unmarshal ...
func (m *MonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorRequest.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorRequest) XXX_Size() int {
	return xxx_messageInfo_MonitorRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorRequest proto.InternalMessageInfo

// MonitorManagerRequest ...
type MonitorManagerRequest struct {
	Type                 ManagerType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.ManagerType" json:"type,omitempty"`
	Data                 []string    `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

// Reset ...
func (m *MonitorManagerRequest) Reset() { *m = MonitorManagerRequest{} }

// String ...
func (m *MonitorManagerRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorManagerRequest) ProtoMessage() {}

// Descriptor ...
func (*MonitorManagerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{1}
}

// XXX_Unmarshal ...
func (m *MonitorManagerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorManagerRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorManagerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorManagerRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorManagerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorManagerRequest.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorManagerRequest) XXX_Size() int {
	return xxx_messageInfo_MonitorManagerRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorManagerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorManagerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorManagerRequest proto.InternalMessageInfo

// GetType ...
func (m *MonitorManagerRequest) GetType() ManagerType {
	if m != nil {
		return m.Type
	}
	return ManagerType_BootstrapAdd
}

// GetData ...
func (m *MonitorManagerRequest) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

// MonitorInitRequest ...
type MonitorInitRequest struct {
	StartMode            StartMode `protobuf:"varint,1,opt,name=start_mode,json=startMode,proto3,enum=proto.StartMode" json:"start_mode,omitempty"`
	Host                 string    `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Bootstrap            string    `protobuf:"bytes,3,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	Secret               string    `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	Workspace            string    `protobuf:"bytes,5,opt,name=workspace,proto3" json:"workspace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

// Reset ...
func (m *MonitorInitRequest) Reset() { *m = MonitorInitRequest{} }

// String ...
func (m *MonitorInitRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorInitRequest) ProtoMessage() {}

// Descriptor ...
func (*MonitorInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{2}
}

// XXX_Unmarshal ...
func (m *MonitorInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorInitRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorInitRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorInitRequest.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorInitRequest) XXX_Size() int {
	return xxx_messageInfo_MonitorInitRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorInitRequest proto.InternalMessageInfo

// GetStartMode ...
func (m *MonitorInitRequest) GetStartMode() StartMode {
	if m != nil {
		return m.StartMode
	}
	return StartMode_Cluster
}

// GetHost ...
func (m *MonitorInitRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

// GetBootstrap ...
func (m *MonitorInitRequest) GetBootstrap() string {
	if m != nil {
		return m.Bootstrap
	}
	return ""
}

// GetSecret ...
func (m *MonitorInitRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

// GetWorkspace ...
func (m *MonitorInitRequest) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

// MonitorBootstrapReply ...
type MonitorBootstrapReply struct {
	Bootstraps           []string `protobuf:"bytes,1,rep,name=bootstraps,proto3" json:"bootstraps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *MonitorBootstrapReply) Reset() { *m = MonitorBootstrapReply{} }

// String ...
func (m *MonitorBootstrapReply) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorBootstrapReply) ProtoMessage() {}

// Descriptor ...
func (*MonitorBootstrapReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{3}
}

// XXX_Unmarshal ...
func (m *MonitorBootstrapReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorBootstrapReply.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorBootstrapReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorBootstrapReply.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorBootstrapReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorBootstrapReply.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorBootstrapReply) XXX_Size() int {
	return xxx_messageInfo_MonitorBootstrapReply.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorBootstrapReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorBootstrapReply.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorBootstrapReply proto.InternalMessageInfo

// GetBootstraps ...
func (m *MonitorBootstrapReply) GetBootstraps() []string {
	if m != nil {
		return m.Bootstraps
	}
	return nil
}

// MonitorAddressReply ...
type MonitorAddressReply struct {
	Addresses            []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *MonitorAddressReply) Reset() { *m = MonitorAddressReply{} }

// String ...
func (m *MonitorAddressReply) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorAddressReply) ProtoMessage() {}

// Descriptor ...
func (*MonitorAddressReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{4}
}

// XXX_Unmarshal ...
func (m *MonitorAddressReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorAddressReply.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorAddressReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorAddressReply.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorAddressReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorAddressReply.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorAddressReply) XXX_Size() int {
	return xxx_messageInfo_MonitorAddressReply.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorAddressReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorAddressReply.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorAddressReply proto.InternalMessageInfo

// GetAddresses ...
func (m *MonitorAddressReply) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// MonitorPinReply ...
type MonitorPinReply struct {
	Pins                 []string `protobuf:"bytes,1,rep,name=pins,proto3" json:"pins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *MonitorPinReply) Reset() { *m = MonitorPinReply{} }

// String ...
func (m *MonitorPinReply) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorPinReply) ProtoMessage() {}

// Descriptor ...
func (*MonitorPinReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{5}
}

// XXX_Unmarshal ...
func (m *MonitorPinReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorPinReply.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorPinReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorPinReply.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorPinReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorPinReply.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorPinReply) XXX_Size() int {
	return xxx_messageInfo_MonitorPinReply.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorPinReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorPinReply.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorPinReply proto.InternalMessageInfo

// GetPins ...
func (m *MonitorPinReply) GetPins() []string {
	if m != nil {
		return m.Pins
	}
	return nil
}

// MonitorProcRequest ...
type MonitorProcRequest struct {
	Type                 MonitorType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.MonitorType" json:"type,omitempty"`
	BootStrap            string      `protobuf:"bytes,2,opt,name=boot_strap,json=bootStrap,proto3" json:"boot_strap,omitempty"`
	Secret               string      `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	Workspace            string      `protobuf:"bytes,4,opt,name=workspace,proto3" json:"workspace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

// Reset ...
func (m *MonitorProcRequest) Reset() { *m = MonitorProcRequest{} }

// String ...
func (m *MonitorProcRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorProcRequest) ProtoMessage() {}

// Descriptor ...
func (*MonitorProcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{6}
}

// XXX_Unmarshal ...
func (m *MonitorProcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorProcRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorProcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorProcRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorProcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorProcRequest.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorProcRequest) XXX_Size() int {
	return xxx_messageInfo_MonitorProcRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorProcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorProcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorProcRequest proto.InternalMessageInfo

// GetType ...
func (m *MonitorProcRequest) GetType() MonitorType {
	if m != nil {
		return m.Type
	}
	return MonitorType_Init
}

// GetBootStrap ...
func (m *MonitorProcRequest) GetBootStrap() string {
	if m != nil {
		return m.BootStrap
	}
	return ""
}

// GetSecret ...
func (m *MonitorProcRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

// GetWorkspace ...
func (m *MonitorProcRequest) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

// MonitorCensorRequest ...
type MonitorCensorRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *MonitorCensorRequest) Reset() { *m = MonitorCensorRequest{} }

// String ...
func (m *MonitorCensorRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorCensorRequest) ProtoMessage() {}

// Descriptor ...
func (*MonitorCensorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{7}
}

// XXX_Unmarshal ...
func (m *MonitorCensorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorCensorRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorCensorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorCensorRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorCensorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorCensorRequest.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorCensorRequest) XXX_Size() int {
	return xxx_messageInfo_MonitorCensorRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorCensorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorCensorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorCensorRequest proto.InternalMessageInfo

// GetID ...
func (m *MonitorCensorRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

// GetDetail ...
func (m *MonitorCensorRequest) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

// MonitorReply ...
type MonitorReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Detail               string   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *MonitorReply) Reset() { *m = MonitorReply{} }

// String ...
func (m *MonitorReply) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*MonitorReply) ProtoMessage() {}

// Descriptor ...
func (*MonitorReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{8}
}

// XXX_Unmarshal ...
func (m *MonitorReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorReply.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *MonitorReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorReply.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *MonitorReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorReply.Merge(m, src)
}

// XXX_Size ...
func (m *MonitorReply) XXX_Size() int {
	return xxx_messageInfo_MonitorReply.Size(m)
}

// XXX_DiscardUnknown ...
func (m *MonitorReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorReply.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorReply proto.InternalMessageInfo

// GetCode ...
func (m *MonitorReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

// GetMessage ...
func (m *MonitorReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// GetDetail ...
func (m *MonitorReply) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.ManagerType", ManagerType_name, ManagerType_value)
	proto.RegisterEnum("proto.StartMode", StartMode_name, StartMode_value)
	proto.RegisterEnum("proto.MonitorType", MonitorType_name, MonitorType_value)
	proto.RegisterType((*MonitorRequest)(nil), "proto.MonitorRequest")
	proto.RegisterType((*MonitorManagerRequest)(nil), "proto.MonitorManagerRequest")
	proto.RegisterType((*MonitorInitRequest)(nil), "proto.MonitorInitRequest")
	proto.RegisterType((*MonitorBootstrapReply)(nil), "proto.MonitorBootstrapReply")
	proto.RegisterType((*MonitorAddressReply)(nil), "proto.MonitorAddressReply")
	proto.RegisterType((*MonitorPinReply)(nil), "proto.MonitorPinReply")
	proto.RegisterType((*MonitorProcRequest)(nil), "proto.MonitorProcRequest")
	proto.RegisterType((*MonitorCensorRequest)(nil), "proto.MonitorCensorRequest")
	proto.RegisterType((*MonitorReply)(nil), "proto.MonitorReply")
}

func init() { proto.RegisterFile("monitor.proto", fileDescriptor_44174b7b2a306b71) }

var fileDescriptor_44174b7b2a306b71 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xcd, 0x57, 0x37, 0x72, 0xb7, 0x75, 0xd1, 0x2d, 0x9b, 0xc2, 0x34, 0xa0, 0x8a, 0x00, 0x4d,
	0x15, 0x2a, 0xd2, 0xf6, 0x00, 0x2f, 0x80, 0xd6, 0x8d, 0x87, 0x3e, 0x54, 0x54, 0xe9, 0xde, 0xa7,
	0x2c, 0x31, 0x25, 0xa2, 0x89, 0x43, 0xec, 0x15, 0xf5, 0x67, 0xf0, 0x37, 0x90, 0xf8, 0x8f, 0xc8,
	0x8e, 0x93, 0x3a, 0x91, 0x2a, 0x9e, 0x62, 0x9f, 0xeb, 0x73, 0x1c, 0x9f, 0x7b, 0x6c, 0x38, 0xca,
	0x68, 0x9e, 0x72, 0x5a, 0x8e, 0x8b, 0x92, 0x72, 0x8a, 0x3d, 0xf9, 0x09, 0x3c, 0xe8, 0xcf, 0x2a,
	0x3c, 0x24, 0x3f, 0x1f, 0x09, 0xe3, 0xc1, 0x02, 0x4e, 0x14, 0x32, 0x8b, 0xf2, 0x68, 0x49, 0xea,
	0x02, 0xbe, 0x01, 0x87, 0x6f, 0x0a, 0xe2, 0x9b, 0x43, 0xf3, 0xa2, 0x7f, 0x89, 0x95, 0xce, 0x58,
	0x2d, 0xba, 0xdb, 0x14, 0x24, 0x94, 0x75, 0x44, 0x70, 0x92, 0x88, 0x47, 0xbe, 0x35, 0xb4, 0x2f,
	0xdc, 0x50, 0x8e, 0x83, 0xbf, 0x26, 0xa0, 0x52, 0x9d, 0xe6, 0x29, 0xaf, 0x25, 0xdf, 0x01, 0x30,
	0x1e, 0x95, 0xfc, 0x3e, 0xa3, 0x49, 0x2d, 0xec, 0x29, 0xe1, 0x85, 0x28, 0xcc, 0x68, 0x42, 0x42,
	0x97, 0xd5, 0x43, 0xa1, 0xfd, 0x9d, 0x32, 0xee, 0x5b, 0x43, 0x53, 0x68, 0x8b, 0x31, 0x9e, 0x83,
	0xfb, 0x40, 0x29, 0x67, 0xbc, 0x8c, 0x0a, 0xdf, 0x96, 0x85, 0x2d, 0x80, 0xa7, 0xb0, 0xc7, 0x48,
	0x5c, 0x12, 0xee, 0x3b, 0xb2, 0xa4, 0x66, 0x82, 0xf5, 0x8b, 0x96, 0x3f, 0x58, 0x11, 0xc5, 0xc4,
	0xef, 0x55, 0xac, 0x06, 0x08, 0xde, 0x37, 0x26, 0x4c, 0x6a, 0xa5, 0x90, 0x14, 0xab, 0x0d, 0xbe,
	0x00, 0x68, 0xb4, 0x99, 0x6f, 0xca, 0x23, 0x6a, 0x48, 0x70, 0x05, 0x03, 0x45, 0xbc, 0x4e, 0x92,
	0x92, 0x30, 0x56, 0xd1, 0xce, 0xc1, 0x8d, 0xaa, 0x39, 0xa9, 0x59, 0x5b, 0x20, 0x78, 0x0d, 0xc7,
	0x8a, 0x34, 0x4f, 0xf3, 0x8a, 0x80, 0xe0, 0x14, 0x69, 0x5e, 0xaf, 0x95, 0xe3, 0xe0, 0xf7, 0xd6,
	0xc4, 0x79, 0x49, 0xe3, 0xff, 0xf4, 0xa5, 0x5a, 0xa8, 0xf5, 0xe5, 0x79, 0xf5, 0xeb, 0xf7, 0x95,
	0x51, 0xd6, 0xd6, 0xa8, 0x45, 0xc7, 0x28, 0x7b, 0xb7, 0x51, 0x4e, 0xd7, 0xa8, 0x4f, 0xf0, 0x54,
	0xed, 0x74, 0x43, 0x72, 0xd6, 0xa4, 0x08, 0xfb, 0x60, 0x4d, 0x6f, 0xe5, 0x2f, 0xb9, 0xa1, 0x35,
	0xbd, 0x15, 0xea, 0x09, 0xe1, 0x51, 0xba, 0x52, 0x1b, 0xab, 0x59, 0x70, 0x07, 0x87, 0x4d, 0xfe,
	0xd4, 0xb9, 0xe3, 0x3a, 0x0b, 0xbd, 0x50, 0x8e, 0xd1, 0x87, 0xfd, 0x8c, 0x30, 0x16, 0x2d, 0x89,
	0x22, 0xd7, 0x53, 0x4d, 0xd5, 0xd6, 0x55, 0x47, 0x5f, 0xe1, 0x40, 0xcb, 0x25, 0x7a, 0x70, 0xd8,
	0xb4, 0xf1, 0x3a, 0x49, 0x3c, 0x03, 0x07, 0x70, 0xac, 0x35, 0x36, 0xa3, 0x6b, 0xe2, 0x99, 0x08,
	0xb0, 0x37, 0x4f, 0x73, 0xb1, 0xc0, 0xc2, 0x23, 0x70, 0x65, 0x2f, 0x64, 0xc9, 0x1e, 0xbd, 0x02,
	0xb7, 0xc9, 0x23, 0x1e, 0xc0, 0xfe, 0xcd, 0xea, 0x91, 0x71, 0x52, 0x7a, 0x86, 0x20, 0x2d, 0xd2,
	0xac, 0x58, 0x11, 0xcf, 0x1c, 0xbd, 0x85, 0x03, 0xcd, 0x76, 0x7c, 0x02, 0x8e, 0x08, 0xbb, 0x67,
	0xa0, 0x0b, 0xbd, 0x90, 0x30, 0xc2, 0x3d, 0xb3, 0x02, 0xbf, 0x51, 0xcf, 0xba, 0xfc, 0x63, 0x43,
	0x5f, 0xe9, 0x28, 0x16, 0x7e, 0x6e, 0x04, 0x04, 0x11, 0x9f, 0xb5, 0x7b, 0xa9, 0xdd, 0x9c, 0xb3,
	0x41, 0xbb, 0x24, 0xcd, 0x0b, 0x0c, 0x4d, 0x40, 0x24, 0xa4, 0x2b, 0xa0, 0xa5, 0x66, 0x97, 0xc0,
	0x97, 0xe6, 0x3d, 0x50, 0xf9, 0xc5, 0x93, 0xee, 0xc2, 0x8a, 0x7f, 0xd6, 0x86, 0xf5, 0xb4, 0x07,
	0x06, 0x4e, 0xc1, 0xeb, 0xde, 0x9f, 0x5d, 0x42, 0xe7, 0x6d, 0xb8, 0x7d, 0xdf, 0x02, 0x03, 0x3f,
	0x02, 0x6c, 0x2f, 0xc7, 0x2e, 0x91, 0xd3, 0xce, 0x41, 0xd5, 0x35, 0x6a, 0x1d, 0x48, 0x25, 0x02,
	0x3b, 0x1b, 0xb6, 0x5f, 0xb9, 0x1d, 0xbe, 0x4c, 0x3e, 0xc0, 0xcb, 0x98, 0x66, 0xe3, 0xf5, 0x2a,
	0x5a, 0x97, 0xe3, 0x25, 0x4d, 0x62, 0x9a, 0x2f, 0xc7, 0xad, 0x17, 0x75, 0x32, 0x68, 0x37, 0x73,
	0x2e, 0xc0, 0xb9, 0xf9, 0xb0, 0x27, 0xab, 0x57, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xec,
	0xaa, 0xfe, 0x80, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterMonitorClient is the client API for ClusterMonitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterMonitorClient interface {
	MonitorInit(ctx context.Context, in *MonitorInitRequest, opts ...grpc.CallOption) (*MonitorReply, error)
	MonitorProc(ctx context.Context, in *MonitorProcRequest, opts ...grpc.CallOption) (*MonitorReply, error)
	MonitorAddress(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorAddressReply, error)
	MonitorBootstrap(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorBootstrapReply, error)
	MonitorPin(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorPinReply, error)
	MonitorManager(ctx context.Context, in *MonitorManagerRequest, opts ...grpc.CallOption) (*MonitorReply, error)
}

type clusterMonitorClient struct {
	cc *grpc.ClientConn
}

// NewClusterMonitorClient ...
func NewClusterMonitorClient(cc *grpc.ClientConn) ClusterMonitorClient {
	return &clusterMonitorClient{cc}
}

// MonitorInit ...
func (c *clusterMonitorClient) MonitorInit(ctx context.Context, in *MonitorInitRequest, opts ...grpc.CallOption) (*MonitorReply, error) {
	out := new(MonitorReply)
	err := c.cc.Invoke(ctx, "/proto.ClusterMonitor/MonitorInit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorProc ...
func (c *clusterMonitorClient) MonitorProc(ctx context.Context, in *MonitorProcRequest, opts ...grpc.CallOption) (*MonitorReply, error) {
	out := new(MonitorReply)
	err := c.cc.Invoke(ctx, "/proto.ClusterMonitor/MonitorProc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorAddress ...
func (c *clusterMonitorClient) MonitorAddress(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorAddressReply, error) {
	out := new(MonitorAddressReply)
	err := c.cc.Invoke(ctx, "/proto.ClusterMonitor/MonitorAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorBootstrap ...
func (c *clusterMonitorClient) MonitorBootstrap(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorBootstrapReply, error) {
	out := new(MonitorBootstrapReply)
	err := c.cc.Invoke(ctx, "/proto.ClusterMonitor/MonitorBootstrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorPin ...
func (c *clusterMonitorClient) MonitorPin(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (*MonitorPinReply, error) {
	out := new(MonitorPinReply)
	err := c.cc.Invoke(ctx, "/proto.ClusterMonitor/MonitorPin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorManager ...
func (c *clusterMonitorClient) MonitorManager(ctx context.Context, in *MonitorManagerRequest, opts ...grpc.CallOption) (*MonitorReply, error) {
	out := new(MonitorReply)
	err := c.cc.Invoke(ctx, "/proto.ClusterMonitor/MonitorManager", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterMonitorServer is the server API for ClusterMonitor service.
type ClusterMonitorServer interface {
	MonitorInit(context.Context, *MonitorInitRequest) (*MonitorReply, error)
	MonitorProc(context.Context, *MonitorProcRequest) (*MonitorReply, error)
	MonitorAddress(context.Context, *MonitorRequest) (*MonitorAddressReply, error)
	MonitorBootstrap(context.Context, *MonitorRequest) (*MonitorBootstrapReply, error)
	MonitorPin(context.Context, *MonitorRequest) (*MonitorPinReply, error)
	MonitorManager(context.Context, *MonitorManagerRequest) (*MonitorReply, error)
}

// RegisterClusterMonitorServer ...
func RegisterClusterMonitorServer(s *grpc.Server, srv ClusterMonitorServer) {
	s.RegisterService(&_ClusterMonitor_serviceDesc, srv)
}

func _ClusterMonitor_MonitorInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterMonitorServer).MonitorInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ClusterMonitor/MonitorInit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterMonitorServer).MonitorInit(ctx, req.(*MonitorInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterMonitor_MonitorProc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorProcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterMonitorServer).MonitorProc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ClusterMonitor/MonitorProc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterMonitorServer).MonitorProc(ctx, req.(*MonitorProcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterMonitor_MonitorAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterMonitorServer).MonitorAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ClusterMonitor/MonitorAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterMonitorServer).MonitorAddress(ctx, req.(*MonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterMonitor_MonitorBootstrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterMonitorServer).MonitorBootstrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ClusterMonitor/MonitorBootstrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterMonitorServer).MonitorBootstrap(ctx, req.(*MonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterMonitor_MonitorPin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterMonitorServer).MonitorPin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ClusterMonitor/MonitorPin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterMonitorServer).MonitorPin(ctx, req.(*MonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterMonitor_MonitorManager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorManagerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterMonitorServer).MonitorManager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ClusterMonitor/MonitorManager",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterMonitorServer).MonitorManager(ctx, req.(*MonitorManagerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterMonitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ClusterMonitor",
	HandlerType: (*ClusterMonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MonitorInit",
			Handler:    _ClusterMonitor_MonitorInit_Handler,
		},
		{
			MethodName: "MonitorProc",
			Handler:    _ClusterMonitor_MonitorProc_Handler,
		},
		{
			MethodName: "MonitorAddress",
			Handler:    _ClusterMonitor_MonitorAddress_Handler,
		},
		{
			MethodName: "MonitorBootstrap",
			Handler:    _ClusterMonitor_MonitorBootstrap_Handler,
		},
		{
			MethodName: "MonitorPin",
			Handler:    _ClusterMonitor_MonitorPin_Handler,
		},
		{
			MethodName: "MonitorManager",
			Handler:    _ClusterMonitor_MonitorManager_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor.proto",
}
