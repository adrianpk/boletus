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
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
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

func (m *EventRes) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

type IndexEventsRes struct {
	Events               []*IndexEventsRes `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *IndexEventsRes) GetEvents() []*IndexEventsRes {
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
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x54, 0xb2, 0xe0, 0xe2, 0x72, 0x2d, 0x4b,
	0xcd, 0x2b, 0xf1, 0x74, 0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x24, 0xb8, 0xd8, 0x53, 0x21, 0xf2, 0x12, 0x4c,
	0x60, 0x51, 0x18, 0x57, 0xc9, 0x88, 0x8b, 0x03, 0xac, 0x33, 0x28, 0xb5, 0x18, 0x8b, 0x3e, 0x11,
	0x2e, 0x56, 0xb0, 0x42, 0xa8, 0x2e, 0x08, 0x47, 0xc9, 0x86, 0x8b, 0xcf, 0x33, 0x2f, 0x25, 0xb5,
	0x02, 0xac, 0xb1, 0x18, 0xa4, 0x53, 0x8b, 0x8b, 0x0d, 0x2c, 0x55, 0x2c, 0xc1, 0xa8, 0xc0, 0xac,
	0xc1, 0x6d, 0x24, 0xa4, 0x57, 0x66, 0xa8, 0x87, 0xaa, 0x26, 0x08, 0xaa, 0xc2, 0xc8, 0x9e, 0x8b,
	0x23, 0x24, 0x33, 0x39, 0x3b, 0xb5, 0x24, 0xb5, 0x48, 0xc8, 0x98, 0x8b, 0x1b, 0x49, 0x95, 0x10,
	0x1f, 0x48, 0x1b, 0xc2, 0x23, 0x52, 0x58, 0x8c, 0x51, 0x62, 0x48, 0x62, 0x03, 0xfb, 0xdb, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xa6, 0x5e, 0x76, 0x04, 0x01, 0x00, 0x00,
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
