// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/grpc.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// CREATE
type CreateRequest struct {
	Data                 *CreateRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetData() *CreateRequestData {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateRequestData struct {
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateRequestData) Reset()         { *m = CreateRequestData{} }
func (m *CreateRequestData) String() string { return proto.CompactTextString(m) }
func (*CreateRequestData) ProtoMessage()    {}
func (*CreateRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{1}
}
func (m *CreateRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequestData.Unmarshal(m, b)
}
func (m *CreateRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequestData.Marshal(b, m, deterministic)
}
func (dst *CreateRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequestData.Merge(dst, src)
}
func (m *CreateRequestData) XXX_Size() int {
	return xxx_messageInfo_CreateRequestData.Size(m)
}
func (m *CreateRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequestData proto.InternalMessageInfo

func (m *CreateRequestData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateRequestData) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// UPDATE
type UpdateRequest struct {
	Data                 *UpdateRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{2}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetData() *UpdateRequestData {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateRequestData struct {
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateRequestData) Reset()         { *m = UpdateRequestData{} }
func (m *UpdateRequestData) String() string { return proto.CompactTextString(m) }
func (*UpdateRequestData) ProtoMessage()    {}
func (*UpdateRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{3}
}
func (m *UpdateRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequestData.Unmarshal(m, b)
}
func (m *UpdateRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequestData.Marshal(b, m, deterministic)
}
func (dst *UpdateRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequestData.Merge(dst, src)
}
func (m *UpdateRequestData) XXX_Size() int {
	return xxx_messageInfo_UpdateRequestData.Size(m)
}
func (m *UpdateRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequestData proto.InternalMessageInfo

func (m *UpdateRequestData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateRequestData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequestData) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// FIND
type FindRequest struct {
	Page                 int32             `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	SortBy               string            `protobuf:"bytes,3,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	Filter               map[string]string `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{4}
}
func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (dst *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(dst, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *FindRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FindRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *FindRequest) GetFilter() map[string]string {
	if m != nil {
		return m.Filter
	}
	return nil
}

type FindResponse struct {
	Metadata             *Metadata           `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Data                 []*FindResponseData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{5}
}
func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (dst *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(dst, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FindResponse) GetData() []*FindResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FindResponseData struct {
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FindResponseData) Reset()         { *m = FindResponseData{} }
func (m *FindResponseData) String() string { return proto.CompactTextString(m) }
func (*FindResponseData) ProtoMessage()    {}
func (*FindResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{6}
}
func (m *FindResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponseData.Unmarshal(m, b)
}
func (m *FindResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponseData.Marshal(b, m, deterministic)
}
func (dst *FindResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponseData.Merge(dst, src)
}
func (m *FindResponseData) XXX_Size() int {
	return xxx_messageInfo_FindResponseData.Size(m)
}
func (m *FindResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponseData proto.InternalMessageInfo

func (m *FindResponseData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FindResponseData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FindResponseData) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Metadata struct {
	Paginate             *Paginate `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{7}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetPaginate() *Paginate {
	if m != nil {
		return m.Paginate
	}
	return nil
}

type Paginate struct {
	TotalCount           int32    `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	PageCount            int32    `protobuf:"varint,2,opt,name=pageCount,proto3" json:"pageCount,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	SortBy               string   `protobuf:"bytes,5,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paginate) Reset()         { *m = Paginate{} }
func (m *Paginate) String() string { return proto.CompactTextString(m) }
func (*Paginate) ProtoMessage()    {}
func (*Paginate) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{8}
}
func (m *Paginate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paginate.Unmarshal(m, b)
}
func (m *Paginate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paginate.Marshal(b, m, deterministic)
}
func (dst *Paginate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paginate.Merge(dst, src)
}
func (m *Paginate) XXX_Size() int {
	return xxx_messageInfo_Paginate.Size(m)
}
func (m *Paginate) XXX_DiscardUnknown() {
	xxx_messageInfo_Paginate.DiscardUnknown(m)
}

var xxx_messageInfo_Paginate proto.InternalMessageInfo

func (m *Paginate) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Paginate) GetPageCount() int32 {
	if m != nil {
		return m.PageCount
	}
	return 0
}

func (m *Paginate) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Paginate) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Paginate) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

// FIND ONE
type FindOneRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindOneRequest) Reset()         { *m = FindOneRequest{} }
func (m *FindOneRequest) String() string { return proto.CompactTextString(m) }
func (*FindOneRequest) ProtoMessage()    {}
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{9}
}
func (m *FindOneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneRequest.Unmarshal(m, b)
}
func (m *FindOneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneRequest.Marshal(b, m, deterministic)
}
func (dst *FindOneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneRequest.Merge(dst, src)
}
func (m *FindOneRequest) XXX_Size() int {
	return xxx_messageInfo_FindOneRequest.Size(m)
}
func (m *FindOneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneRequest proto.InternalMessageInfo

func (m *FindOneRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// DELETE BY ID
type DeleteByIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteByIdRequest) Reset()         { *m = DeleteByIdRequest{} }
func (m *DeleteByIdRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteByIdRequest) ProtoMessage()    {}
func (*DeleteByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{10}
}
func (m *DeleteByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteByIdRequest.Unmarshal(m, b)
}
func (m *DeleteByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteByIdRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteByIdRequest.Merge(dst, src)
}
func (m *DeleteByIdRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteByIdRequest.Size(m)
}
func (m *DeleteByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteByIdRequest proto.InternalMessageInfo

func (m *DeleteByIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// RESPONSE
type Response struct {
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ff991094aadd815f, []int{11}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "grpc.CreateRequest")
	proto.RegisterType((*CreateRequestData)(nil), "grpc.CreateRequestData")
	proto.RegisterMapType((map[string]string)(nil), "grpc.CreateRequestData.AttributesEntry")
	proto.RegisterType((*UpdateRequest)(nil), "grpc.UpdateRequest")
	proto.RegisterType((*UpdateRequestData)(nil), "grpc.UpdateRequestData")
	proto.RegisterMapType((map[string]string)(nil), "grpc.UpdateRequestData.AttributesEntry")
	proto.RegisterType((*FindRequest)(nil), "grpc.FindRequest")
	proto.RegisterMapType((map[string]string)(nil), "grpc.FindRequest.FilterEntry")
	proto.RegisterType((*FindResponse)(nil), "grpc.FindResponse")
	proto.RegisterType((*FindResponseData)(nil), "grpc.FindResponseData")
	proto.RegisterMapType((map[string]string)(nil), "grpc.FindResponseData.AttributesEntry")
	proto.RegisterType((*Metadata)(nil), "grpc.Metadata")
	proto.RegisterType((*Paginate)(nil), "grpc.Paginate")
	proto.RegisterType((*FindOneRequest)(nil), "grpc.FindOneRequest")
	proto.RegisterType((*DeleteByIdRequest)(nil), "grpc.DeleteByIdRequest")
	proto.RegisterType((*Response)(nil), "grpc.Response")
	proto.RegisterMapType((map[string]string)(nil), "grpc.Response.AttributesEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModelClient is the client API for Model service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error)
	FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*Response, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	DeleteById(ctx context.Context, in *DeleteByIdRequest, opts ...grpc.CallOption) (*Response, error)
}

type modelClient struct {
	cc *grpc.ClientConn
}

func NewModelClient(cc *grpc.ClientConn) ModelClient {
	return &modelClient{cc}
}

func (c *modelClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Model/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Model/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelClient) FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Model/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/grpc.Model/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelClient) DeleteById(ctx context.Context, in *DeleteByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Model/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelServer is the server API for Model service.
type ModelServer interface {
	Create(context.Context, *CreateRequest) (*Response, error)
	Update(context.Context, *UpdateRequest) (*Response, error)
	FindOne(context.Context, *FindOneRequest) (*Response, error)
	Find(context.Context, *FindRequest) (*FindResponse, error)
	DeleteById(context.Context, *DeleteByIdRequest) (*Response, error)
}

func RegisterModelServer(s *grpc.Server, srv ModelServer) {
	s.RegisterService(&_Model_serviceDesc, srv)
}

func _Model_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Model/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Model_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Model/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Model_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Model/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServer).FindOne(ctx, req.(*FindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Model_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Model/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Model_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Model/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServer).DeleteById(ctx, req.(*DeleteByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Model_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Model",
	HandlerType: (*ModelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Model_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Model_Update_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _Model_FindOne_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _Model_Find_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _Model_DeleteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/grpc.proto",
}

func init() { proto.RegisterFile("grpc/grpc.proto", fileDescriptor_grpc_ff991094aadd815f) }

var fileDescriptor_grpc_ff991094aadd815f = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xed, 0x6a, 0xd3, 0x50,
	0x18, 0xde, 0x49, 0xd3, 0x9a, 0xbd, 0xd5, 0x6e, 0x7d, 0x1d, 0x33, 0x14, 0x1d, 0x25, 0x82, 0x8e,
	0x09, 0x1d, 0x56, 0xfc, 0x44, 0x05, 0xb7, 0x59, 0xf1, 0x47, 0x51, 0x02, 0x5e, 0xc0, 0x99, 0x39,
	0x2b, 0xc1, 0x2c, 0x89, 0xc9, 0xa9, 0x90, 0x1b, 0xf0, 0x87, 0xf7, 0x22, 0x5e, 0x81, 0x88, 0x77,
	0x26, 0xe7, 0x23, 0x5f, 0x4d, 0x86, 0xdb, 0x8f, 0xfe, 0x29, 0xe7, 0xbc, 0x9f, 0xe7, 0x79, 0x9f,
	0x27, 0x6f, 0x61, 0x6b, 0x91, 0xc4, 0x9f, 0x0f, 0xc5, 0xcf, 0x24, 0x4e, 0x22, 0x1e, 0xa1, 0x29,
	0xce, 0xce, 0x4b, 0xb8, 0x71, 0x9c, 0x30, 0xca, 0x99, 0xcb, 0xbe, 0x2e, 0x59, 0xca, 0xf1, 0x01,
	0x98, 0x1e, 0xe5, 0xd4, 0x26, 0x63, 0xb2, 0xdf, 0x9f, 0xde, 0x9a, 0xc8, 0x8c, 0x5a, 0xc8, 0x09,
	0xe5, 0xd4, 0x95, 0x41, 0xce, 0x2f, 0x02, 0xc3, 0x86, 0x0f, 0x11, 0x4c, 0x9e, 0xc5, 0x4c, 0x96,
	0xd8, 0x74, 0xe5, 0x19, 0xdf, 0x01, 0x50, 0xce, 0x13, 0xff, 0x74, 0xc9, 0x59, 0x6a, 0x1b, 0xe3,
	0xce, 0x7e, 0x7f, 0x7a, 0xff, 0x82, 0xe2, 0x93, 0x37, 0x45, 0xe4, 0xdb, 0x90, 0x27, 0x99, 0x5b,
	0x49, 0x1d, 0xbd, 0x82, 0xad, 0x15, 0x37, 0x6e, 0x43, 0xe7, 0x0b, 0xcb, 0x74, 0x3b, 0x71, 0xc4,
	0x1d, 0xe8, 0x7e, 0xa3, 0xc1, 0x92, 0xd9, 0x86, 0xb4, 0xa9, 0xcb, 0x0b, 0xe3, 0x19, 0x11, 0x78,
	0x3f, 0xc5, 0xde, 0xff, 0xf0, 0xd6, 0x42, 0x2a, 0x78, 0xff, 0x10, 0x18, 0x36, 0x7c, 0xad, 0x78,
	0x07, 0x60, 0xf8, 0x9e, 0x6e, 0x6f, 0xf8, 0xde, 0x0a, 0xfe, 0x4e, 0x15, 0x7f, 0xa3, 0xe0, 0x3a,
	0xf1, 0xff, 0x25, 0xd0, 0x9f, 0xf9, 0xa1, 0x97, 0xc3, 0x47, 0x30, 0x63, 0xba, 0x50, 0x6f, 0xef,
	0xba, 0xf2, 0x2c, 0xb2, 0x03, 0xff, 0xdc, 0xe7, 0x32, 0xbb, 0xeb, 0xaa, 0x0b, 0xee, 0x42, 0x2f,
	0x8d, 0x12, 0x7e, 0x94, 0xd9, 0x1d, 0x59, 0x54, 0xdf, 0xf0, 0x31, 0xf4, 0xce, 0xfc, 0x80, 0xb3,
	0xc4, 0x36, 0x25, 0xaa, 0x3b, 0x0a, 0x55, 0xa5, 0xc9, 0x64, 0x26, 0xfd, 0x0a, 0x8b, 0x0e, 0x1e,
	0x3d, 0x17, 0xef, 0x28, 0xcc, 0x57, 0xc2, 0x70, 0x06, 0xd7, 0x55, 0xf5, 0x34, 0x8e, 0xc2, 0x94,
	0xe1, 0x01, 0x58, 0xe7, 0x8c, 0xd3, 0x0a, 0x8d, 0x03, 0xf5, 0x86, 0xb9, 0xb6, 0xba, 0x85, 0x1f,
	0x0f, 0x34, 0xdd, 0x4a, 0x81, 0xbb, 0xd5, 0xb7, 0xaa, 0x6a, 0x15, 0xb6, 0x7f, 0x13, 0xd8, 0x5e,
	0x75, 0x5d, 0x8a, 0xec, 0x59, 0x0b, 0xd9, 0xf7, 0xda, 0x5b, 0xad, 0x93, 0xeb, 0x27, 0x60, 0xcd,
	0x4b, 0xdc, 0x56, 0x4c, 0x17, 0x7e, 0x48, 0x39, 0xab, 0xcf, 0xe8, 0xa3, 0xb6, 0xba, 0x85, 0xdf,
	0xf9, 0x41, 0xc0, 0xca, 0xcd, 0xb8, 0x07, 0xc0, 0x23, 0x4e, 0x83, 0xe3, 0x68, 0x19, 0x72, 0x2d,
	0x93, 0x8a, 0x05, 0x6f, 0xc3, 0xa6, 0x10, 0x8d, 0x72, 0x2b, 0xc1, 0x94, 0x86, 0x42, 0x5e, 0x9d,
	0x36, 0x79, 0x99, 0xed, 0xf2, 0xea, 0x56, 0xe5, 0xe5, 0x8c, 0x61, 0x20, 0x66, 0xf6, 0x21, 0x2c,
	0xbe, 0x58, 0x35, 0x6d, 0x92, 0x4f, 0xdb, 0xb9, 0x0b, 0xc3, 0x13, 0x16, 0x30, 0xce, 0x8e, 0xb2,
	0xf7, 0xde, 0x45, 0x41, 0x3f, 0x09, 0x58, 0x85, 0x60, 0x2e, 0xc3, 0xe1, 0xeb, 0x16, 0x0e, 0xf7,
	0xd4, 0xc8, 0xf2, 0x3a, 0x6b, 0xe4, 0x6e, 0xfa, 0xdd, 0x80, 0xee, 0x3c, 0xf2, 0x58, 0x80, 0x87,
	0xd0, 0x53, 0x1b, 0x12, 0x6f, 0xb6, 0xec, 0xcb, 0xd1, 0xa0, 0xfe, 0x26, 0x67, 0x43, 0x24, 0xa8,
	0x95, 0x92, 0x27, 0xd4, 0x16, 0x4c, 0x4b, 0xc2, 0x43, 0xb8, 0xa6, 0x47, 0x8c, 0x3b, 0xa5, 0x4a,
	0xcb, 0x89, 0xb7, 0xf6, 0x30, 0x45, 0x0c, 0x0e, 0x1b, 0x1f, 0xfb, 0x08, 0x9b, 0x42, 0x77, 0x36,
	0xf0, 0x29, 0x40, 0x49, 0x12, 0xea, 0x35, 0xdb, 0xa0, 0xad, 0xd9, 0xe9, 0xb4, 0x27, 0xff, 0xad,
	0x1e, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xd7, 0x25, 0xed, 0xc0, 0x06, 0x00, 0x00,
}
