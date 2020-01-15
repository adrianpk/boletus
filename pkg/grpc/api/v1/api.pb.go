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
	EventID              string   `protobuf:"bytes,2,opt,name=eventID,proto3" json:"eventID,omitempty"`
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

func (m *EventIDReq) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

type EventRes struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Place                string   `protobuf:"bytes,5,opt,name=place,proto3" json:"place,omitempty"`
	ScheduledAt          string   `protobuf:"bytes,6,opt,name=scheduledAt,proto3" json:"scheduledAt,omitempty"`
	Timezone             string   `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty"`
	IsNew                bool     `protobuf:"varint,8,opt,name=isNew,proto3" json:"isNew,omitempty"`
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
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0xd2, 0xa6, 0xe9, 0x44, 0x8a, 0x0c, 0x1e, 0x86, 0x9e, 0x42, 0xf0, 0x90, 0x53,
	0xa0, 0xed, 0xc5, 0x9b, 0x08, 0x0a, 0xf6, 0xe2, 0x21, 0xf8, 0x07, 0x62, 0x32, 0xe8, 0x62, 0x9a,
	0xc4, 0xec, 0x36, 0x8a, 0xbf, 0xd2, 0x9f, 0x24, 0x3b, 0x69, 0x4b, 0x84, 0xde, 0xe6, 0xbd, 0xc7,
	0x37, 0xcb, 0xbc, 0x85, 0x79, 0xde, 0xaa, 0xb4, 0xed, 0x1a, 0xd3, 0xa0, 0xdb, 0xaf, 0xe2, 0x5b,
	0x80, 0xc7, 0x9e, 0x6b, 0xb3, 0x7d, 0xc8, 0xf8, 0x13, 0xaf, 0xc0, 0xcb, 0x5b, 0x45, 0x4e, 0xe4,
	0x24, 0xf3, 0xcc, 0x8e, 0x48, 0x30, 0xe3, 0x21, 0x27, 0x57, 0xdc, 0xa3, 0x8c, 0x7f, 0x1d, 0x08,
	0x04, 0xcd, 0x58, 0x9f, 0x01, 0x11, 0x26, 0xba, 0xda, 0xbf, 0x1d, 0x28, 0x99, 0xad, 0x57, 0xe7,
	0x3b, 0x26, 0x6f, 0xf0, 0xec, 0x8c, 0x11, 0x84, 0x25, 0xeb, 0xa2, 0x53, 0xad, 0x51, 0x4d, 0x4d,
	0x13, 0x89, 0xc6, 0x16, 0x5e, 0xc3, 0xb4, 0xad, 0xf2, 0x82, 0x69, 0x2a, 0xd9, 0x20, 0x2c, 0xa7,
	0x8b, 0x77, 0x2e, 0xf7, 0x15, 0x97, 0xf7, 0x86, 0xfc, 0x81, 0x1b, 0x59, 0xb8, 0x84, 0xc0, 0xa8,
	0x1d, 0xff, 0x34, 0x35, 0xd3, 0x4c, 0xe2, 0x93, 0xb6, 0x3b, 0x95, 0x7e, 0xe6, 0x2f, 0x0a, 0x22,
	0x27, 0x09, 0xb2, 0x41, 0xc4, 0x4f, 0xb0, 0xd8, 0xd6, 0x25, 0x7f, 0xcb, 0x59, 0xfa, 0xfc, 0x5d,
	0x37, 0xe0, 0x4b, 0x03, 0x9a, 0xdc, 0xc8, 0x4b, 0xc2, 0xf5, 0x65, 0xda, 0xaf, 0xd2, 0x63, 0x0f,
	0xd9, 0x21, 0x5b, 0xdf, 0x41, 0xf0, 0xa2, 0x8a, 0x0f, 0x36, 0xdc, 0xe1, 0x06, 0xc2, 0xd1, 0x56,
	0x5c, 0x9c, 0x00, 0xe9, 0x7c, 0x89, 0x56, 0xff, 0x7f, 0x36, 0xbe, 0x78, 0xf5, 0xe5, 0x8b, 0x36,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x34, 0x56, 0xf4, 0x11, 0xaf, 0x01, 0x00, 0x00,
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
