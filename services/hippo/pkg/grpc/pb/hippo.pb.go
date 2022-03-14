// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hippo.proto

package hippo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ValueType int32

const (
	ValueType_UNKNOWN    ValueType = 0
	ValueType_FLOAT16    ValueType = 1
	ValueType_FLOAT32    ValueType = 2
	ValueType_FLOAT64    ValueType = 3
	ValueType_UINT8      ValueType = 4
	ValueType_UINT16     ValueType = 5
	ValueType_UINT32     ValueType = 6
	ValueType_UINT64     ValueType = 7
	ValueType_INT8       ValueType = 8
	ValueType_INT16      ValueType = 9
	ValueType_INT32      ValueType = 10
	ValueType_INT64      ValueType = 11
	ValueType_COMPLEX64  ValueType = 12
	ValueType_COMPLEX128 ValueType = 13
)

var ValueType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "FLOAT16",
	2:  "FLOAT32",
	3:  "FLOAT64",
	4:  "UINT8",
	5:  "UINT16",
	6:  "UINT32",
	7:  "UINT64",
	8:  "INT8",
	9:  "INT16",
	10: "INT32",
	11: "INT64",
	12: "COMPLEX64",
	13: "COMPLEX128",
}

var ValueType_value = map[string]int32{
	"UNKNOWN":    0,
	"FLOAT16":    1,
	"FLOAT32":    2,
	"FLOAT64":    3,
	"UINT8":      4,
	"UINT16":     5,
	"UINT32":     6,
	"UINT64":     7,
	"INT8":       8,
	"INT16":      9,
	"INT32":      10,
	"INT64":      11,
	"COMPLEX64":  12,
	"COMPLEX128": 13,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db052698e7b66ea3, []int{0}
}

type Shape struct {
	Value                []int64  `protobuf:"varint,1,rep,packed,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shape) Reset()         { *m = Shape{} }
func (m *Shape) String() string { return proto.CompactTextString(m) }
func (*Shape) ProtoMessage()    {}
func (*Shape) Descriptor() ([]byte, []int) {
	return fileDescriptor_db052698e7b66ea3, []int{0}
}

func (m *Shape) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shape.Unmarshal(m, b)
}
func (m *Shape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shape.Marshal(b, m, deterministic)
}
func (m *Shape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shape.Merge(m, src)
}
func (m *Shape) XXX_Size() int {
	return xxx_messageInfo_Shape.Size(m)
}
func (m *Shape) XXX_DiscardUnknown() {
	xxx_messageInfo_Shape.DiscardUnknown(m)
}

var xxx_messageInfo_Shape proto.InternalMessageInfo

func (m *Shape) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Tensor struct {
	Type                 ValueType `protobuf:"varint,1,opt,name=Type,proto3,enum=hippo.ValueType" json:"Type,omitempty"`
	Shape                *Shape    `protobuf:"bytes,2,opt,name=Shape,proto3" json:"Shape,omitempty"`
	Data                 []byte    `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Tensor) Reset()         { *m = Tensor{} }
func (m *Tensor) String() string { return proto.CompactTextString(m) }
func (*Tensor) ProtoMessage()    {}
func (*Tensor) Descriptor() ([]byte, []int) {
	return fileDescriptor_db052698e7b66ea3, []int{1}
}

func (m *Tensor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tensor.Unmarshal(m, b)
}
func (m *Tensor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tensor.Marshal(b, m, deterministic)
}
func (m *Tensor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tensor.Merge(m, src)
}
func (m *Tensor) XXX_Size() int {
	return xxx_messageInfo_Tensor.Size(m)
}
func (m *Tensor) XXX_DiscardUnknown() {
	xxx_messageInfo_Tensor.DiscardUnknown(m)
}

var xxx_messageInfo_Tensor proto.InternalMessageInfo

func (m *Tensor) GetType() ValueType {
	if m != nil {
		return m.Type
	}
	return ValueType_UNKNOWN
}

func (m *Tensor) GetShape() *Shape {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *Tensor) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RunRequest struct {
	ModelID              string    `protobuf:"bytes,1,opt,name=ModelID,proto3" json:"ModelID,omitempty"`
	Tensors              []*Tensor `protobuf:"bytes,2,rep,name=Tensors,proto3" json:"Tensors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db052698e7b66ea3, []int{2}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

func (m *RunRequest) GetModelID() string {
	if m != nil {
		return m.ModelID
	}
	return ""
}

func (m *RunRequest) GetTensors() []*Tensor {
	if m != nil {
		return m.Tensors
	}
	return nil
}

type RunReply struct {
	Tensors              []*Tensor `protobuf:"bytes,1,rep,name=Tensors,proto3" json:"Tensors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RunReply) Reset()         { *m = RunReply{} }
func (m *RunReply) String() string { return proto.CompactTextString(m) }
func (*RunReply) ProtoMessage()    {}
func (*RunReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_db052698e7b66ea3, []int{3}
}

func (m *RunReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunReply.Unmarshal(m, b)
}
func (m *RunReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunReply.Marshal(b, m, deterministic)
}
func (m *RunReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunReply.Merge(m, src)
}
func (m *RunReply) XXX_Size() int {
	return xxx_messageInfo_RunReply.Size(m)
}
func (m *RunReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RunReply.DiscardUnknown(m)
}

var xxx_messageInfo_RunReply proto.InternalMessageInfo

func (m *RunReply) GetTensors() []*Tensor {
	if m != nil {
		return m.Tensors
	}
	return nil
}

func init() {
	proto.RegisterEnum("hippo.ValueType", ValueType_name, ValueType_value)
	proto.RegisterType((*Shape)(nil), "hippo.Shape")
	proto.RegisterType((*Tensor)(nil), "hippo.Tensor")
	proto.RegisterType((*RunRequest)(nil), "hippo.RunRequest")
	proto.RegisterType((*RunReply)(nil), "hippo.RunReply")
}

func init() { proto.RegisterFile("hippo.proto", fileDescriptor_db052698e7b66ea3) }

var fileDescriptor_db052698e7b66ea3 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0xf3, 0xaf, 0x99, 0xb4, 0x75, 0x1d, 0x3c, 0x04, 0x41, 0x08, 0x41, 0x30, 0x7a,
	0x28, 0x74, 0x13, 0x42, 0xaf, 0x62, 0x15, 0x8b, 0x6d, 0x22, 0x6b, 0xaa, 0x5e, 0x23, 0x46, 0x2a,
	0x84, 0x66, 0x6d, 0x9b, 0x43, 0xbf, 0x99, 0x1f, 0x4f, 0xb2, 0xd9, 0xd4, 0x9e, 0xbc, 0xbd, 0xdf,
	0x9b, 0x37, 0x33, 0x99, 0xb0, 0x60, 0x2f, 0xbf, 0x38, 0x2f, 0x87, 0x7c, 0x5d, 0x6e, 0x4b, 0xd4,
	0x05, 0x78, 0xe7, 0xa0, 0x3f, 0x2f, 0x33, 0x9e, 0xe3, 0x29, 0xe8, 0x2f, 0x59, 0x51, 0xe5, 0x8e,
	0xe2, 0xaa, 0xbe, 0xca, 0x1a, 0xf0, 0x3e, 0xc1, 0x48, 0xf3, 0xd5, 0xa6, 0x5c, 0xe3, 0x05, 0x68,
	0xe9, 0x8e, 0xd7, 0x65, 0xc5, 0x1f, 0x50, 0x32, 0x6c, 0x66, 0x89, 0x54, 0xed, 0x33, 0x51, 0x45,
	0x4f, 0x8e, 0x73, 0x3a, 0xae, 0xe2, 0xdb, 0xb4, 0x27, 0x63, 0xc2, 0x63, 0x72, 0x13, 0x82, 0x36,
	0xc9, 0xb6, 0x99, 0xa3, 0xba, 0x8a, 0xdf, 0x63, 0x42, 0x7b, 0x09, 0x00, 0xab, 0x56, 0x2c, 0xff,
	0xae, 0xf2, 0xcd, 0x16, 0x1d, 0x30, 0xe7, 0xe5, 0x47, 0x5e, 0x4c, 0x27, 0x62, 0x9d, 0xc5, 0x5a,
	0xc4, 0x4b, 0x30, 0x9b, 0xef, 0xd9, 0x38, 0x1d, 0x57, 0xf5, 0x6d, 0xda, 0x97, 0x1b, 0x1a, 0x97,
	0xb5, 0x55, 0x2f, 0x80, 0xae, 0x18, 0xc8, 0x8b, 0xdd, 0x61, 0x93, 0xf2, 0x5f, 0xd3, 0xf5, 0x8f,
	0x02, 0xd6, 0xfe, 0x22, 0xb4, 0xc1, 0x5c, 0xc4, 0x8f, 0x71, 0xf2, 0x1a, 0x93, 0xa3, 0x1a, 0xee,
	0x67, 0xc9, 0x4d, 0x3a, 0x8a, 0x88, 0xb2, 0x87, 0x80, 0x92, 0xce, 0x1e, 0xa2, 0x90, 0xa8, 0x68,
	0x81, 0xbe, 0x98, 0xc6, 0xe9, 0x98, 0x68, 0x08, 0x60, 0xd4, 0x72, 0x14, 0x11, 0xbd, 0xd5, 0x01,
	0x25, 0x46, 0xab, 0xa3, 0x90, 0x98, 0xd8, 0x05, 0x4d, 0xa4, 0xbb, 0x75, 0x63, 0x13, 0xb6, 0xa4,
	0x0c, 0x28, 0x01, 0x29, 0xa3, 0x90, 0xd8, 0xd8, 0x07, 0xeb, 0x36, 0x99, 0x3f, 0xcd, 0xee, 0xde,
	0xa2, 0x90, 0xf4, 0x70, 0x00, 0x20, 0x71, 0x44, 0xc7, 0xa4, 0x4f, 0x29, 0xe8, 0x0f, 0xf5, 0x4d,
	0x78, 0x05, 0x2a, 0xab, 0x56, 0x78, 0x22, 0x4f, 0xfc, 0xfb, 0xab, 0x67, 0xc7, 0x87, 0x16, 0x2f,
	0x76, 0xef, 0x86, 0x78, 0x09, 0xc1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x58, 0x35, 0x23,
	0x18, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HippoClient is the client API for Hippo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HippoClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunReply, error)
}

type hippoClient struct {
	cc *grpc.ClientConn
}

func NewHippoClient(cc *grpc.ClientConn) HippoClient {
	return &hippoClient{cc}
}

func (c *hippoClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunReply, error) {
	out := new(RunReply)
	err := c.cc.Invoke(ctx, "/hippo.Hippo/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HippoServer is the server API for Hippo service.
type HippoServer interface {
	Run(context.Context, *RunRequest) (*RunReply, error)
}

// UnimplementedHippoServer can be embedded to have forward compatible implementations.
type UnimplementedHippoServer struct {
}

func (*UnimplementedHippoServer) Run(ctx context.Context, req *RunRequest) (*RunReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterHippoServer(s *grpc.Server, srv HippoServer) {
	s.RegisterService(&_Hippo_serviceDesc, srv)
}

func _Hippo_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HippoServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hippo.Hippo/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HippoServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hippo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hippo.Hippo",
	HandlerType: (*HippoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Hippo_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hippo.proto",
}
