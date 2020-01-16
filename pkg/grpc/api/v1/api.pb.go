// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package v1

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

type EventIDReq struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventIDReq) Reset()         { *m = EventIDReq{} }
func (m *EventIDReq) String() string { return proto.CompactTextString(m) }
func (*EventIDReq) ProtoMessage()    {}
func (*EventIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *EventIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventIDReq.Unmarshal(m, b)
}
func (m *EventIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventIDReq.Marshal(b, m, deterministic)
}
func (m *EventIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventIDReq.Merge(m, src)
}
func (m *EventIDReq) XXX_Size() int {
	return xxx_messageInfo_EventIDReq.Size(m)
}
func (m *EventIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EventIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_EventIDReq proto.InternalMessageInfo

func (m *EventIDReq) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *EventIDReq) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type EventRes struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Place                string   `protobuf:"bytes,6,opt,name=place,proto3" json:"place,omitempty"`
	ScheduledAt          string   `protobuf:"bytes,7,opt,name=scheduledAt,proto3" json:"scheduledAt,omitempty"`
	Timezone             string   `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	IsNew                bool     `protobuf:"varint,9,opt,name=isNew,proto3" json:"isNew,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventRes) Reset()         { *m = EventRes{} }
func (m *EventRes) String() string { return proto.CompactTextString(m) }
func (*EventRes) ProtoMessage()    {}
func (*EventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *EventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRes.Unmarshal(m, b)
}
func (m *EventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRes.Marshal(b, m, deterministic)
}
func (m *EventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRes.Merge(m, src)
}
func (m *EventRes) XXX_Size() int {
	return xxx_messageInfo_EventRes.Size(m)
}
func (m *EventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRes.DiscardUnknown(m)
}

var xxx_messageInfo_EventRes proto.InternalMessageInfo

func (m *EventRes) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *EventRes) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *EventRes) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EventRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventRes) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventRes) GetPlace() string {
	if m != nil {
		return m.Place
	}
	return ""
}

func (m *EventRes) GetScheduledAt() string {
	if m != nil {
		return m.ScheduledAt
	}
	return ""
}

func (m *EventRes) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *EventRes) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

type IndexEventsRes struct {
	Api                  string      `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Events               []*EventRes `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *IndexEventsRes) Reset()         { *m = IndexEventsRes{} }
func (m *IndexEventsRes) String() string { return proto.CompactTextString(m) }
func (*IndexEventsRes) ProtoMessage()    {}
func (*IndexEventsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *IndexEventsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexEventsRes.Unmarshal(m, b)
}
func (m *IndexEventsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexEventsRes.Marshal(b, m, deterministic)
}
func (m *IndexEventsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexEventsRes.Merge(m, src)
}
func (m *IndexEventsRes) XXX_Size() int {
	return xxx_messageInfo_IndexEventsRes.Size(m)
}
func (m *IndexEventsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexEventsRes.DiscardUnknown(m)
}

var xxx_messageInfo_IndexEventsRes proto.InternalMessageInfo

func (m *IndexEventsRes) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *IndexEventsRes) GetEvents() []*EventRes {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*EventIDReq)(nil), "v1.EventIDReq")
	proto.RegisterType((*EventRes)(nil), "v1.EventRes")
	proto.RegisterType((*IndexEventsRes)(nil), "v1.IndexEventsRes")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4a, 0xf4, 0x40,
	0x10, 0xc5, 0xbf, 0x24, 0x33, 0xf9, 0x92, 0x8a, 0x0c, 0xd2, 0xb8, 0x28, 0x66, 0x15, 0x82, 0x8b,
	0xac, 0x02, 0x93, 0x39, 0x80, 0x08, 0x0a, 0xce, 0xc6, 0x45, 0xf0, 0x02, 0x31, 0x29, 0xb4, 0x31,
	0x93, 0xb4, 0xe9, 0x9e, 0xf8, 0xe7, 0xbc, 0x1e, 0x44, 0xba, 0x5a, 0x87, 0x08, 0x2e, 0xdc, 0xbd,
	0xf7, 0x7b, 0xbc, 0x6a, 0xaa, 0x0b, 0xe2, 0x5a, 0xc9, 0x42, 0x8d, 0x83, 0x19, 0x84, 0x3f, 0x6d,
	0xb2, 0x12, 0xe0, 0x7a, 0xa2, 0xde, 0xec, 0xae, 0x2a, 0x7a, 0x16, 0xa7, 0x10, 0xd4, 0x4a, 0xa2,
	0x97, 0x7a, 0x79, 0x5c, 0x59, 0x29, 0x04, 0x2c, 0x74, 0x77, 0x78, 0x40, 0x9f, 0x11, 0xeb, 0xec,
	0xc3, 0x83, 0x88, 0x4b, 0x15, 0xe9, 0xbf, 0x55, 0x2c, 0x33, 0x6f, 0x8a, 0x30, 0x70, 0xcc, 0x6a,
	0xcb, 0xfa, 0x7a, 0x4f, 0xb8, 0x70, 0xcc, 0x6a, 0x91, 0x42, 0xd2, 0x92, 0x6e, 0x46, 0xa9, 0x8c,
	0x1c, 0x7a, 0x5c, 0x72, 0x34, 0x47, 0xe2, 0x0c, 0x96, 0xaa, 0xab, 0x1b, 0xc2, 0x90, 0x33, 0x67,
	0x6c, 0x4f, 0x37, 0x8f, 0xd4, 0x1e, 0x3a, 0x6a, 0x2f, 0x0d, 0xfe, 0x77, 0xbd, 0x19, 0x12, 0x6b,
	0x88, 0x8c, 0xdc, 0xd3, 0xfb, 0xd0, 0x13, 0x46, 0x1c, 0x1f, 0xbd, 0x9d, 0x29, 0xf5, 0x2d, 0xbd,
	0x60, 0x9c, 0x7a, 0x79, 0x54, 0x39, 0x93, 0xdd, 0xc0, 0x6a, 0xd7, 0xb7, 0xf4, 0xca, 0xab, 0xea,
	0xdf, 0x77, 0x3d, 0x87, 0x90, 0x38, 0x46, 0x3f, 0x0d, 0xf2, 0xa4, 0x3c, 0x29, 0xa6, 0x4d, 0xf1,
	0xfd, 0x37, 0xd5, 0x57, 0x56, 0x5e, 0x40, 0x74, 0x27, 0x9b, 0x27, 0x32, 0x34, 0x8a, 0x2d, 0x24,
	0xb3, 0xa9, 0x62, 0x75, 0x2c, 0xf0, 0x05, 0xd6, 0xc2, 0xfa, 0x9f, 0xcf, 0x66, 0xff, 0xee, 0x43,
	0x3e, 0xd8, 0xf6, 0x33, 0x00, 0x00, 0xff, 0xff, 0x69, 0xa2, 0x08, 0x94, 0xbd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TicketerClient is the client API for Ticketer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TicketerClient interface {
	IndexEvents(ctx context.Context, in *EventIDReq, opts ...grpc.CallOption) (*IndexEventsRes, error)
}

type ticketerClient struct {
	cc *grpc.ClientConn
}

func NewTicketerClient(cc *grpc.ClientConn) TicketerClient {
	return &ticketerClient{cc}
}

func (c *ticketerClient) IndexEvents(ctx context.Context, in *EventIDReq, opts ...grpc.CallOption) (*IndexEventsRes, error) {
	out := new(IndexEventsRes)
	err := c.cc.Invoke(ctx, "/v1.Ticketer/IndexEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketerServer is the server API for Ticketer service.
type TicketerServer interface {
	IndexEvents(context.Context, *EventIDReq) (*IndexEventsRes, error)
}

// UnimplementedTicketerServer can be embedded to have forward compatible implementations.
type UnimplementedTicketerServer struct {
}

func (*UnimplementedTicketerServer) IndexEvents(ctx context.Context, req *EventIDReq) (*IndexEventsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndexEvents not implemented")
}

func RegisterTicketerServer(s *grpc.Server, srv TicketerServer) {
	s.RegisterService(&_Ticketer_serviceDesc, srv)
}

func _Ticketer_IndexEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketerServer).IndexEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Ticketer/IndexEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketerServer).IndexEvents(ctx, req.(*EventIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ticketer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Ticketer",
	HandlerType: (*TicketerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IndexEvents",
			Handler:    _Ticketer_IndexEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
