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

type TicketSummaryListRes struct {
	Api                  string              `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	List                 []*TicketSummaryRes `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TicketSummaryListRes) Reset()         { *m = TicketSummaryListRes{} }
func (m *TicketSummaryListRes) String() string { return proto.CompactTextString(m) }
func (*TicketSummaryListRes) ProtoMessage()    {}
func (*TicketSummaryListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *TicketSummaryListRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TicketSummaryListRes.Unmarshal(m, b)
}
func (m *TicketSummaryListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TicketSummaryListRes.Marshal(b, m, deterministic)
}
func (m *TicketSummaryListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TicketSummaryListRes.Merge(m, src)
}
func (m *TicketSummaryListRes) XXX_Size() int {
	return xxx_messageInfo_TicketSummaryListRes.Size(m)
}
func (m *TicketSummaryListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_TicketSummaryListRes.DiscardUnknown(m)
}

var xxx_messageInfo_TicketSummaryListRes proto.InternalMessageInfo

func (m *TicketSummaryListRes) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *TicketSummaryListRes) GetList() []*TicketSummaryRes {
	if m != nil {
		return m.List
	}
	return nil
}

type TicketSummaryRes struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EventSlug            string   `protobuf:"bytes,3,opt,name=eventSlug,proto3" json:"eventSlug,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Qty                  int32    `protobuf:"varint,5,opt,name=qty,proto3" json:"qty,omitempty"`
	Price                float32  `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	Currency             string   `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TicketSummaryRes) Reset()         { *m = TicketSummaryRes{} }
func (m *TicketSummaryRes) String() string { return proto.CompactTextString(m) }
func (*TicketSummaryRes) ProtoMessage()    {}
func (*TicketSummaryRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *TicketSummaryRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TicketSummaryRes.Unmarshal(m, b)
}
func (m *TicketSummaryRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TicketSummaryRes.Marshal(b, m, deterministic)
}
func (m *TicketSummaryRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TicketSummaryRes.Merge(m, src)
}
func (m *TicketSummaryRes) XXX_Size() int {
	return xxx_messageInfo_TicketSummaryRes.Size(m)
}
func (m *TicketSummaryRes) XXX_DiscardUnknown() {
	xxx_messageInfo_TicketSummaryRes.DiscardUnknown(m)
}

var xxx_messageInfo_TicketSummaryRes proto.InternalMessageInfo

func (m *TicketSummaryRes) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *TicketSummaryRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TicketSummaryRes) GetEventSlug() string {
	if m != nil {
		return m.EventSlug
	}
	return ""
}

func (m *TicketSummaryRes) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TicketSummaryRes) GetQty() int32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *TicketSummaryRes) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TicketSummaryRes) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type PreBookReq struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	EventSlug            string   `protobuf:"bytes,2,opt,name=eventSlug,proto3" json:"eventSlug,omitempty"`
	TicketType           string   `protobuf:"bytes,3,opt,name=ticketType,proto3" json:"ticketType,omitempty"`
	Qty                  int32    `protobuf:"varint,4,opt,name=qty,proto3" json:"qty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreBookReq) Reset()         { *m = PreBookReq{} }
func (m *PreBookReq) String() string { return proto.CompactTextString(m) }
func (*PreBookReq) ProtoMessage()    {}
func (*PreBookReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *PreBookReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreBookReq.Unmarshal(m, b)
}
func (m *PreBookReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreBookReq.Marshal(b, m, deterministic)
}
func (m *PreBookReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreBookReq.Merge(m, src)
}
func (m *PreBookReq) XXX_Size() int {
	return xxx_messageInfo_PreBookReq.Size(m)
}
func (m *PreBookReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PreBookReq.DiscardUnknown(m)
}

var xxx_messageInfo_PreBookReq proto.InternalMessageInfo

func (m *PreBookReq) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *PreBookReq) GetEventSlug() string {
	if m != nil {
		return m.EventSlug
	}
	return ""
}

func (m *PreBookReq) GetTicketType() string {
	if m != nil {
		return m.TicketType
	}
	return ""
}

func (m *PreBookReq) GetQty() int32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

type PreBookRes struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	EventSlug            string   `protobuf:"bytes,2,opt,name=eventSlug,proto3" json:"eventSlug,omitempty"`
	TicketType           string   `protobuf:"bytes,3,opt,name=ticketType,proto3" json:"ticketType,omitempty"`
	Qty                  int32    `protobuf:"varint,4,opt,name=qty,proto3" json:"qty,omitempty"`
	Price                float32  `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Total                float32  `protobuf:"fixed32,6,opt,name=total,proto3" json:"total,omitempty"`
	Currency             string   `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreBookRes) Reset()         { *m = PreBookRes{} }
func (m *PreBookRes) String() string { return proto.CompactTextString(m) }
func (*PreBookRes) ProtoMessage()    {}
func (*PreBookRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *PreBookRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreBookRes.Unmarshal(m, b)
}
func (m *PreBookRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreBookRes.Marshal(b, m, deterministic)
}
func (m *PreBookRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreBookRes.Merge(m, src)
}
func (m *PreBookRes) XXX_Size() int {
	return xxx_messageInfo_PreBookRes.Size(m)
}
func (m *PreBookRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PreBookRes.DiscardUnknown(m)
}

var xxx_messageInfo_PreBookRes proto.InternalMessageInfo

func (m *PreBookRes) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *PreBookRes) GetEventSlug() string {
	if m != nil {
		return m.EventSlug
	}
	return ""
}

func (m *PreBookRes) GetTicketType() string {
	if m != nil {
		return m.TicketType
	}
	return ""
}

func (m *PreBookRes) GetQty() int32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *PreBookRes) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *PreBookRes) GetTotal() float32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PreBookRes) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PreBookRes) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*EventIDReq)(nil), "v1.EventIDReq")
	proto.RegisterType((*EventRes)(nil), "v1.EventRes")
	proto.RegisterType((*IndexEventsRes)(nil), "v1.IndexEventsRes")
	proto.RegisterType((*TicketSummaryListRes)(nil), "v1.TicketSummaryListRes")
	proto.RegisterType((*TicketSummaryRes)(nil), "v1.TicketSummaryRes")
	proto.RegisterType((*PreBookReq)(nil), "v1.PreBookReq")
	proto.RegisterType((*PreBookRes)(nil), "v1.PreBookRes")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x6d, 0xd2, 0xa4, 0xa6, 0xb7, 0x52, 0xca, 0x50, 0x64, 0x28, 0x22, 0x65, 0xf0, 0xa1, 0x4f,
	0x85, 0xed, 0xfe, 0x80, 0x8a, 0x82, 0x0b, 0x22, 0x92, 0xdd, 0x1f, 0x88, 0xe9, 0x45, 0x87, 0x4d,
	0x93, 0x6c, 0x66, 0x52, 0xad, 0x1f, 0xe5, 0x9b, 0xbf, 0xe1, 0x57, 0xf8, 0x21, 0x32, 0x77, 0xa6,
	0xe9, 0x64, 0xb7, 0x8a, 0x2f, 0xbe, 0xdd, 0x73, 0x6e, 0xee, 0xcc, 0x9d, 0x73, 0x4e, 0x0b, 0xe3,
	0xac, 0x96, 0xeb, 0xba, 0xa9, 0x74, 0xc5, 0xc2, 0xfd, 0x85, 0xd8, 0x00, 0xbc, 0xd9, 0x63, 0xa9,
	0xaf, 0x5e, 0xa7, 0x78, 0xc7, 0x66, 0x30, 0xcc, 0x6a, 0xc9, 0x83, 0x65, 0xb0, 0x1a, 0xa7, 0xa6,
	0x64, 0x0c, 0x22, 0x55, 0xb4, 0x9f, 0x78, 0x48, 0x14, 0xd5, 0xe2, 0x57, 0x00, 0x09, 0x0d, 0xa5,
	0xa8, 0xfe, 0x6d, 0xc4, 0x70, 0xfa, 0x50, 0x23, 0x1f, 0x5a, 0xce, 0xd4, 0x86, 0x2b, 0xb3, 0x1d,
	0xf2, 0xc8, 0x72, 0xa6, 0x66, 0x4b, 0x98, 0x6c, 0x51, 0xe5, 0x8d, 0xac, 0xb5, 0xac, 0x4a, 0x1e,
	0x53, 0xcb, 0xa7, 0xd8, 0x1c, 0xe2, 0xba, 0xc8, 0x72, 0xe4, 0x23, 0xea, 0x59, 0x60, 0xe6, 0x54,
	0xfe, 0x19, 0xb7, 0x6d, 0x81, 0xdb, 0x97, 0x9a, 0x3f, 0xb2, 0x73, 0x1e, 0xc5, 0x16, 0x90, 0x68,
	0xb9, 0xc3, 0x6f, 0x55, 0x89, 0x3c, 0xa1, 0x76, 0x87, 0xcd, 0x99, 0x52, 0xbd, 0xc7, 0x2f, 0x7c,
	0xbc, 0x0c, 0x56, 0x49, 0x6a, 0x81, 0x78, 0x0b, 0xd3, 0xab, 0x72, 0x8b, 0x5f, 0xe9, 0xa9, 0xea,
	0xfc, 0x5b, 0x9f, 0xc3, 0x08, 0xa9, 0xcd, 0xc3, 0xe5, 0x70, 0x35, 0xd9, 0x3c, 0x5e, 0xef, 0x2f,
	0xd6, 0x47, 0x6d, 0x52, 0xd7, 0x13, 0x29, 0xcc, 0x6f, 0x64, 0x7e, 0x8b, 0xfa, 0xba, 0xdd, 0xed,
	0xb2, 0xe6, 0xf0, 0x4e, 0xaa, 0x3f, 0x68, 0xb7, 0x82, 0xa8, 0x90, 0x4a, 0xbb, 0xd3, 0xe6, 0xe6,
	0xb4, 0xde, 0xa4, 0x39, 0x95, 0xbe, 0x10, 0xdf, 0x03, 0x98, 0xdd, 0x6f, 0x9d, 0x37, 0x83, 0x44,
	0x0e, 0x3d, 0x91, 0x9f, 0xc2, 0x98, 0x16, 0xbb, 0x36, 0x2e, 0x59, 0x47, 0x4e, 0x44, 0x67, 0x55,
	0xe4, 0x59, 0x35, 0x83, 0xe1, 0x9d, 0x3e, 0x90, 0x1d, 0x71, 0x6a, 0x4a, 0xb2, 0xa1, 0x91, 0xce,
	0x86, 0x30, 0xb5, 0xc0, 0x88, 0x9c, 0xb7, 0x4d, 0x83, 0x65, 0x7e, 0x70, 0x1e, 0x74, 0x58, 0x94,
	0x00, 0x1f, 0x1a, 0x7c, 0x55, 0x55, 0xb7, 0xe7, 0x93, 0xd6, 0xdb, 0x2a, 0xbc, 0xbf, 0xd5, 0x33,
	0x00, 0x4d, 0xaf, 0xbd, 0x39, 0xc5, 0xc8, 0x63, 0x8e, 0x1b, 0x46, 0xdd, 0x86, 0xe2, 0x67, 0xe0,
	0x5d, 0xa8, 0xfe, 0xff, 0x85, 0x27, 0x49, 0x62, 0x5f, 0x92, 0x39, 0xc4, 0xba, 0xd2, 0x59, 0x71,
	0x14, 0x8a, 0xc0, 0xdf, 0x84, 0x62, 0x4f, 0x60, 0xa4, 0x74, 0xa6, 0x5b, 0xe5, 0x72, 0xea, 0xd0,
	0xe6, 0x47, 0x00, 0x89, 0x75, 0x1c, 0x1b, 0x76, 0x09, 0x13, 0x2f, 0x9c, 0x6c, 0xda, 0xe5, 0x8e,
	0x7e, 0xc8, 0x0b, 0x66, 0x70, 0x3f, 0xbd, 0x62, 0xc0, 0x5e, 0x00, 0x23, 0xd8, 0xcb, 0xcd, 0x83,
	0x59, 0xfe, 0x20, 0x75, 0x2e, 0xaf, 0x62, 0xc0, 0x36, 0x30, 0x75, 0x9a, 0xda, 0x0f, 0xdc, 0xcd,
	0x27, 0x63, 0x17, 0x7d, 0xac, 0xc4, 0xe0, 0xe3, 0x88, 0xfe, 0x6d, 0x2e, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xfb, 0xa7, 0x5b, 0x8a, 0x7a, 0x04, 0x00, 0x00,
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
	EventTicketSummary(ctx context.Context, in *EventIDReq, opts ...grpc.CallOption) (*TicketSummaryListRes, error)
	PreBookTickets(ctx context.Context, in *PreBookReq, opts ...grpc.CallOption) (*PreBookRes, error)
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

func (c *ticketerClient) EventTicketSummary(ctx context.Context, in *EventIDReq, opts ...grpc.CallOption) (*TicketSummaryListRes, error) {
	out := new(TicketSummaryListRes)
	err := c.cc.Invoke(ctx, "/v1.Ticketer/EventTicketSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketerClient) PreBookTickets(ctx context.Context, in *PreBookReq, opts ...grpc.CallOption) (*PreBookRes, error) {
	out := new(PreBookRes)
	err := c.cc.Invoke(ctx, "/v1.Ticketer/PreBookTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketerServer is the server API for Ticketer service.
type TicketerServer interface {
	IndexEvents(context.Context, *EventIDReq) (*IndexEventsRes, error)
	EventTicketSummary(context.Context, *EventIDReq) (*TicketSummaryListRes, error)
	PreBookTickets(context.Context, *PreBookReq) (*PreBookRes, error)
}

// UnimplementedTicketerServer can be embedded to have forward compatible implementations.
type UnimplementedTicketerServer struct {
}

func (*UnimplementedTicketerServer) IndexEvents(ctx context.Context, req *EventIDReq) (*IndexEventsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndexEvents not implemented")
}
func (*UnimplementedTicketerServer) EventTicketSummary(ctx context.Context, req *EventIDReq) (*TicketSummaryListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventTicketSummary not implemented")
}
func (*UnimplementedTicketerServer) PreBookTickets(ctx context.Context, req *PreBookReq) (*PreBookRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreBookTickets not implemented")
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

func _Ticketer_EventTicketSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketerServer).EventTicketSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Ticketer/EventTicketSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketerServer).EventTicketSummary(ctx, req.(*EventIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticketer_PreBookTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketerServer).PreBookTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Ticketer/PreBookTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketerServer).PreBookTickets(ctx, req.(*PreBookReq))
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
		{
			MethodName: "EventTicketSummary",
			Handler:    _Ticketer_EventTicketSummary_Handler,
		},
		{
			MethodName: "PreBookTickets",
			Handler:    _Ticketer_PreBookTickets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
