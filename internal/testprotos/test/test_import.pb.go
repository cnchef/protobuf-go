// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test_import.proto

package test

import (
	bytes "bytes"
	gzip "compress/gzip"
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ImportEnum int32

const (
	ImportEnum_IMPORT_ZERO ImportEnum = 0
)

func (e ImportEnum) Type() protoreflect.EnumType {
	return xxx_ProtoFile_test_import_enumTypes[0]
}
func (e ImportEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var ImportEnum_name = map[int32]string{
	0: "IMPORT_ZERO",
}

var ImportEnum_value = map[string]int32{
	"IMPORT_ZERO": 0,
}

func (x ImportEnum) Enum() *ImportEnum {
	p := new(ImportEnum)
	*p = x
	return p
}

func (x ImportEnum) String() string {
	return proto.EnumName(ImportEnum_name, int32(x))
}

func (x *ImportEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImportEnum_value, data, "ImportEnum")
	if err != nil {
		return err
	}
	*x = ImportEnum(value)
	return nil
}

func (ImportEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89be98e26346f54e_gzipped, []int{0}
}

type ImportMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportMessage) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_test_import_messageTypes[0].MessageOf(m)
}
func (m *ImportMessage) Reset()         { *m = ImportMessage{} }
func (m *ImportMessage) String() string { return proto.CompactTextString(m) }
func (*ImportMessage) ProtoMessage()    {}
func (*ImportMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_89be98e26346f54e_gzipped, []int{0}
}

func (m *ImportMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportMessage.Unmarshal(m, b)
}
func (m *ImportMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportMessage.Marshal(b, m, deterministic)
}
func (m *ImportMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportMessage.Merge(m, src)
}
func (m *ImportMessage) XXX_Size() int {
	return xxx_messageInfo_ImportMessage.Size(m)
}
func (m *ImportMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ImportMessage proto.InternalMessageInfo

func init() {
	proto.RegisterFile("test_import.proto", fileDescriptor_89be98e26346f54e_gzipped)
	proto.RegisterEnum("goproto.proto.test.ImportEnum", ImportEnum_name, ImportEnum_value)
	proto.RegisterType((*ImportMessage)(nil), "goproto.proto.test.ImportMessage")
}

var fileDescriptor_89be98e26346f54e = []byte{
	// 145 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x1d, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74,
}

var fileDescriptor_89be98e26346f54e_gzipped = func() []byte {
	bb := new(bytes.Buffer)
	zw, _ := gzip.NewWriterLevel(bb, gzip.NoCompression)
	zw.Write(fileDescriptor_89be98e26346f54e)
	zw.Close()
	return bb.Bytes()
}()

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_test_import protoreflect.FileDescriptor

var xxx_ProtoFile_test_import_enumTypes [1]protoreflect.EnumType
var xxx_ProtoFile_test_import_messageTypes [1]protoimpl.MessageType
var xxx_ProtoFile_test_import_goTypes = []interface{}{
	(ImportEnum)(0),       // 0: goproto.proto.test.ImportEnum
	(*ImportMessage)(nil), // 1: goproto.proto.test.ImportMessage
}
var xxx_ProtoFile_test_import_depIdxs = []int32{}

func init() {
	var messageTypes [1]protoreflect.MessageType
	ProtoFile_test_import = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_89be98e26346f54e,
		GoTypes:            xxx_ProtoFile_test_import_goTypes,
		DependencyIndexes:  xxx_ProtoFile_test_import_depIdxs,
		EnumOutputTypes:    xxx_ProtoFile_test_import_enumTypes[:],
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_test_import_goTypes[1:][:1]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_test_import_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_test_import_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_test_import_goTypes = nil
	xxx_ProtoFile_test_import_depIdxs = nil
}
