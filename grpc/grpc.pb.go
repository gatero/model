// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog/catalog.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
	Data                 *CreateRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty" bson:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-" bson:"-"`
	XXX_unrecognized     []byte             `json:"-" bson:"-"`
	XXX_sizecache        int32              `json:"-" bson:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
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
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" bson:"type,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-"`
}

func (m *CreateRequestData) Reset()         { *m = CreateRequestData{} }
func (m *CreateRequestData) String() string { return proto.CompactTextString(m) }
func (*CreateRequestData) ProtoMessage()    {}
func (*CreateRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{1}
}

func (m *CreateRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequestData.Unmarshal(m, b)
}
func (m *CreateRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequestData.Marshal(b, m, deterministic)
}
func (m *CreateRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequestData.Merge(m, src)
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
	Data                 *UpdateRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty" bson:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-" bson:"-"`
	XXX_unrecognized     []byte             `json:"-" bson:"-"`
	XXX_sizecache        int32              `json:"-" bson:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{2}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
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
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" bson:"type,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-"`
}

func (m *UpdateRequestData) Reset()         { *m = UpdateRequestData{} }
func (m *UpdateRequestData) String() string { return proto.CompactTextString(m) }
func (*UpdateRequestData) ProtoMessage()    {}
func (*UpdateRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{3}
}

func (m *UpdateRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequestData.Unmarshal(m, b)
}
func (m *UpdateRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequestData.Marshal(b, m, deterministic)
}
func (m *UpdateRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequestData.Merge(m, src)
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
	Page                 int32             `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" bson:"page,omitempty"`
	Limit                int32             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty" bson:"limit,omitempty"`
	SortBy               string            `protobuf:"bytes,3,opt,name=sortBy,proto3" json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	Filter               map[string]string `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{4}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
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
	Metadata             *Metadata           `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty" bson:"metadata,omitempty"`
	Data                 []*FindResponseData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" bson:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-" bson:"-"`
	XXX_unrecognized     []byte              `json:"-" bson:"-"`
	XXX_sizecache        int32               `json:"-" bson:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{5}
}

func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (m *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(m, src)
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
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" bson:"type,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-"`
}

func (m *FindResponseData) Reset()         { *m = FindResponseData{} }
func (m *FindResponseData) String() string { return proto.CompactTextString(m) }
func (*FindResponseData) ProtoMessage()    {}
func (*FindResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{6}
}

func (m *FindResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponseData.Unmarshal(m, b)
}
func (m *FindResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponseData.Marshal(b, m, deterministic)
}
func (m *FindResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponseData.Merge(m, src)
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
	Paginate             *Paginate `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty" bson:"paginate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{7}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
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
	TotalCount           int32    `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty" bson:"totalCount,omitempty"`
	PageCount            int32    `protobuf:"varint,2,opt,name=pageCount,proto3" json:"pageCount,omitempty" bson:"pageCount,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty" bson:"page,omitempty"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty" bson:"limit,omitempty"`
	SortBy               string   `protobuf:"bytes,5,opt,name=sortBy,proto3" json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Paginate) Reset()         { *m = Paginate{} }
func (m *Paginate) String() string { return proto.CompactTextString(m) }
func (*Paginate) ProtoMessage()    {}
func (*Paginate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{8}
}

func (m *Paginate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paginate.Unmarshal(m, b)
}
func (m *Paginate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paginate.Marshal(b, m, deterministic)
}
func (m *Paginate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paginate.Merge(m, src)
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
type FindByIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *FindByIdRequest) Reset()         { *m = FindByIdRequest{} }
func (m *FindByIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindByIdRequest) ProtoMessage()    {}
func (*FindByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{9}
}

func (m *FindByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIdRequest.Unmarshal(m, b)
}
func (m *FindByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIdRequest.Marshal(b, m, deterministic)
}
func (m *FindByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIdRequest.Merge(m, src)
}
func (m *FindByIdRequest) XXX_Size() int {
	return xxx_messageInfo_FindByIdRequest.Size(m)
}
func (m *FindByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIdRequest proto.InternalMessageInfo

func (m *FindByIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// DELETE BY ID
type DeleteByIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DeleteByIdRequest) Reset()         { *m = DeleteByIdRequest{} }
func (m *DeleteByIdRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteByIdRequest) ProtoMessage()    {}
func (*DeleteByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{10}
}

func (m *DeleteByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteByIdRequest.Unmarshal(m, b)
}
func (m *DeleteByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteByIdRequest.Marshal(b, m, deterministic)
}
func (m *DeleteByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteByIdRequest.Merge(m, src)
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
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" bson:"type,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a439a2eac0f2a9d, []int{11}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
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
	proto.RegisterType((*FindByIdRequest)(nil), "grpc.FindByIdRequest")
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

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error)
	FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*Response, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	DeleteById(ctx context.Context, in *DeleteByIdRequest, opts ...grpc.CallOption) (*Response, error)
}

type catalogClient struct {
	cc *grpc.ClientConn
}

func NewCatalogClient(cc *grpc.ClientConn) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) DeleteById(ctx context.Context, in *DeleteByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
type CatalogServer interface {
	Create(context.Context, *CreateRequest) (*Response, error)
	Update(context.Context, *UpdateRequest) (*Response, error)
	FindById(context.Context, *FindByIdRequest) (*Response, error)
	Find(context.Context, *FindRequest) (*FindResponse, error)
	DeleteById(context.Context, *DeleteByIdRequest) (*Response, error)
}

func RegisterCatalogServer(s *grpc.Server, srv CatalogServer) {
	s.RegisterService(&_Catalog_serviceDesc, srv)
}

func _Catalog_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).FindById(ctx, req.(*FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).DeleteById(ctx, req.(*DeleteByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Catalog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Catalog_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Catalog_Update_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _Catalog_FindById_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _Catalog_Find_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _Catalog_DeleteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog/catalog.proto",
}

func init() { proto.RegisterFile("catalog/catalog.proto", fileDescriptor_0a439a2eac0f2a9d) }

var fileDescriptor_0a439a2eac0f2a9d = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0xee, 0x3a, 0x4e, 0x70, 0x27, 0x90, 0x36, 0x03, 0x2d, 0x51, 0x04, 0x55, 0x31, 0x12, 0x54,
	0x45, 0x4a, 0xa5, 0x54, 0xfc, 0x0a, 0x90, 0x68, 0x4a, 0x10, 0x0f, 0x95, 0x90, 0x25, 0x0e, 0xb0,
	0x6d, 0xb6, 0x91, 0x85, 0x1b, 0x1b, 0x7b, 0x82, 0xe4, 0x2b, 0xf4, 0x2e, 0x88, 0x13, 0x20, 0xc4,
	0xcd, 0x90, 0x77, 0xd7, 0xce, 0x3a, 0x36, 0xd0, 0x3e, 0xe4, 0x29, 0xbb, 0xf3, 0xbb, 0xdf, 0x7c,
	0x5f, 0xc6, 0xb0, 0x75, 0xc6, 0x89, 0x07, 0xe1, 0xf4, 0x40, 0xff, 0x0e, 0xa2, 0x38, 0xa4, 0x10,
	0xed, 0x69, 0x1c, 0x9d, 0xb9, 0xaf, 0xe1, 0xd6, 0x28, 0x16, 0x9c, 0x84, 0x27, 0xbe, 0xce, 0x45,
	0x42, 0xf8, 0x04, 0xec, 0x09, 0x27, 0xde, 0x63, 0xbb, 0x6c, 0xaf, 0x3d, 0xbc, 0x3b, 0xc8, 0xa2,
	0x06, 0xa5, 0x90, 0x63, 0x4e, 0xdc, 0x93, 0x41, 0xee, 0x0f, 0x06, 0xdd, 0x8a, 0x0f, 0x11, 0x6c,
	0x4a, 0x23, 0x21, 0x4b, 0xac, 0x7b, 0xf2, 0x8c, 0x1f, 0x00, 0x38, 0x51, 0xec, 0x9f, 0xce, 0x49,
	0x24, 0x3d, 0x6b, 0xb7, 0xb1, 0xd7, 0x1e, 0x3e, 0xfe, 0x4b, 0xf1, 0xc1, 0xbb, 0x22, 0xf2, 0xfd,
	0x8c, 0xe2, 0xd4, 0x33, 0x52, 0xfb, 0x6f, 0x60, 0x63, 0xc9, 0x8d, 0x9b, 0xd0, 0xf8, 0x22, 0x52,
	0xdd, 0x2e, 0x3b, 0xe2, 0x1d, 0x68, 0x7e, 0xe3, 0xc1, 0x5c, 0xf4, 0x2c, 0x69, 0x53, 0x97, 0x57,
	0xd6, 0x0b, 0x96, 0xe1, 0xfd, 0x1c, 0x4d, 0xfe, 0x87, 0xb7, 0x14, 0x62, 0xe0, 0xfd, 0xc5, 0xa0,
	0x5b, 0xf1, 0xd5, 0xe2, 0xed, 0x80, 0xe5, 0x4f, 0x74, 0x7b, 0xcb, 0x9f, 0x2c, 0xe1, 0x6f, 0x98,
	0xf8, 0x2b, 0x05, 0x57, 0x89, 0xff, 0x37, 0x83, 0xf6, 0xd8, 0x9f, 0x4d, 0x72, 0xf8, 0x08, 0x76,
	0xc4, 0xa7, 0xea, 0xed, 0x4d, 0x4f, 0x9e, 0xb3, 0xec, 0xc0, 0xbf, 0xf0, 0x49, 0x66, 0x37, 0x3d,
	0x75, 0xc1, 0x6d, 0x68, 0x25, 0x61, 0x4c, 0x47, 0x69, 0xaf, 0x21, 0x8b, 0xea, 0x1b, 0x3e, 0x85,
	0xd6, 0xb9, 0x1f, 0x90, 0x88, 0x7b, 0xb6, 0x44, 0x75, 0x5f, 0xa1, 0x32, 0x9a, 0x0c, 0xc6, 0xd2,
	0xaf, 0xb0, 0xe8, 0xe0, 0xfe, 0xcb, 0xec, 0x1d, 0x85, 0xf9, 0x5a, 0x18, 0xce, 0xe1, 0xa6, 0xaa,
	0x9e, 0x44, 0xe1, 0x2c, 0x11, 0xb8, 0x0f, 0xce, 0x85, 0x20, 0x6e, 0xd0, 0xd8, 0x51, 0x6f, 0x38,
	0xd1, 0x56, 0xaf, 0xf0, 0xe3, 0xbe, 0xa6, 0x5b, 0x29, 0x70, 0xdb, 0x7c, 0xab, 0xaa, 0x66, 0xb0,
	0xfd, 0x93, 0xc1, 0xe6, 0xb2, 0xeb, 0x4a, 0x64, 0x8f, 0x6b, 0xc8, 0x7e, 0x54, 0xdf, 0x6a, 0x95,
	0x5c, 0x3f, 0x03, 0xe7, 0x64, 0x81, 0xdb, 0x89, 0xf8, 0xd4, 0x9f, 0x71, 0x12, 0xe5, 0x19, 0x7d,
	0xd2, 0x56, 0xaf, 0xf0, 0xbb, 0x97, 0x0c, 0x9c, 0xdc, 0x8c, 0x3b, 0x00, 0x14, 0x12, 0x0f, 0x46,
	0xe1, 0x7c, 0x46, 0x5a, 0x26, 0x86, 0x05, 0xef, 0xc1, 0x7a, 0x26, 0x1a, 0xe5, 0x56, 0x82, 0x59,
	0x18, 0x0a, 0x79, 0x35, 0xea, 0xe4, 0x65, 0xd7, 0xcb, 0xab, 0x69, 0xca, 0xcb, 0x7d, 0x00, 0x1b,
	0xd9, 0xcc, 0x8e, 0xd2, 0x8f, 0x85, 0x66, 0xd5, 0xb8, 0x59, 0x3e, 0x6e, 0xf7, 0x21, 0x74, 0x8f,
	0x45, 0x20, 0x48, 0xfc, 0x2b, 0xe8, 0x3b, 0x03, 0xa7, 0x50, 0xcc, 0x55, 0x48, 0x7c, 0x5b, 0x43,
	0xe2, 0x8e, 0x9a, 0x59, 0x5e, 0x67, 0x85, 0xe4, 0x0d, 0x2f, 0x2d, 0xb8, 0x31, 0x52, 0x0b, 0x1b,
	0x0f, 0xa0, 0xa5, 0x96, 0x24, 0xde, 0xae, 0x59, 0x99, 0xfd, 0x4e, 0xf9, 0x55, 0xee, 0x5a, 0x96,
	0xa0, 0xb6, 0x4a, 0x9e, 0x50, 0xda, 0x31, 0x35, 0x09, 0x87, 0xe0, 0xe4, 0x53, 0xc6, 0xad, 0x85,
	0x52, 0x8d, 0x81, 0xd6, 0x76, 0xb1, 0xb3, 0x20, 0xec, 0x56, 0xfe, 0xf1, 0x7d, 0xac, 0xaa, 0xdd,
	0x5d, 0xc3, 0xe7, 0x00, 0x0b, 0xa2, 0x50, 0xef, 0xda, 0x0a, 0x75, 0xd5, 0x4e, 0xa7, 0x2d, 0xf9,
	0xc9, 0x3a, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xc9, 0x92, 0x55, 0xcb, 0x06, 0x00, 0x00,
}
