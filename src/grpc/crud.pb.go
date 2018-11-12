// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/crud/crud.proto

package api

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
	return fileDescriptor_8f6983d17f831b9d, []int{0}
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
	return fileDescriptor_8f6983d17f831b9d, []int{1}
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
	return fileDescriptor_8f6983d17f831b9d, []int{2}
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
	return fileDescriptor_8f6983d17f831b9d, []int{3}
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
	return fileDescriptor_8f6983d17f831b9d, []int{4}
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
	return fileDescriptor_8f6983d17f831b9d, []int{5}
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
	return fileDescriptor_8f6983d17f831b9d, []int{6}
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
	return fileDescriptor_8f6983d17f831b9d, []int{7}
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
	return fileDescriptor_8f6983d17f831b9d, []int{8}
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
type FindByIdRequestData struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" bson:"type,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *FindByIdRequestData) Reset()         { *m = FindByIdRequestData{} }
func (m *FindByIdRequestData) String() string { return proto.CompactTextString(m) }
func (*FindByIdRequestData) ProtoMessage()    {}
func (*FindByIdRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6983d17f831b9d, []int{9}
}

func (m *FindByIdRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIdRequestData.Unmarshal(m, b)
}
func (m *FindByIdRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIdRequestData.Marshal(b, m, deterministic)
}
func (m *FindByIdRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIdRequestData.Merge(m, src)
}
func (m *FindByIdRequestData) XXX_Size() int {
	return xxx_messageInfo_FindByIdRequestData.Size(m)
}
func (m *FindByIdRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIdRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIdRequestData proto.InternalMessageInfo

func (m *FindByIdRequestData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FindByIdRequestData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FindByIdRequest struct {
	Data                 *FindByIdRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty" bson:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" bson:"-"`
	XXX_unrecognized     []byte               `json:"-" bson:"-"`
	XXX_sizecache        int32                `json:"-" bson:"-"`
}

func (m *FindByIdRequest) Reset()         { *m = FindByIdRequest{} }
func (m *FindByIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindByIdRequest) ProtoMessage()    {}
func (*FindByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6983d17f831b9d, []int{10}
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

func (m *FindByIdRequest) GetData() *FindByIdRequestData {
	if m != nil {
		return m.Data
	}
	return nil
}

// DELETE BY ID
type DeleteByIdRequestData struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" bson:"type,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" bson:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DeleteByIdRequestData) Reset()         { *m = DeleteByIdRequestData{} }
func (m *DeleteByIdRequestData) String() string { return proto.CompactTextString(m) }
func (*DeleteByIdRequestData) ProtoMessage()    {}
func (*DeleteByIdRequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6983d17f831b9d, []int{11}
}

func (m *DeleteByIdRequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteByIdRequestData.Unmarshal(m, b)
}
func (m *DeleteByIdRequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteByIdRequestData.Marshal(b, m, deterministic)
}
func (m *DeleteByIdRequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteByIdRequestData.Merge(m, src)
}
func (m *DeleteByIdRequestData) XXX_Size() int {
	return xxx_messageInfo_DeleteByIdRequestData.Size(m)
}
func (m *DeleteByIdRequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteByIdRequestData.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteByIdRequestData proto.InternalMessageInfo

func (m *DeleteByIdRequestData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DeleteByIdRequestData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteByIdRequest struct {
	Data                 *DeleteByIdRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty" bson:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-" bson:"-"`
	XXX_unrecognized     []byte                 `json:"-" bson:"-"`
	XXX_sizecache        int32                  `json:"-" bson:"-"`
}

func (m *DeleteByIdRequest) Reset()         { *m = DeleteByIdRequest{} }
func (m *DeleteByIdRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteByIdRequest) ProtoMessage()    {}
func (*DeleteByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f6983d17f831b9d, []int{12}
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

func (m *DeleteByIdRequest) GetData() *DeleteByIdRequestData {
	if m != nil {
		return m.Data
	}
	return nil
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
	return fileDescriptor_8f6983d17f831b9d, []int{13}
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
	proto.RegisterType((*CreateRequest)(nil), "api.CreateRequest")
	proto.RegisterType((*CreateRequestData)(nil), "api.CreateRequestData")
	proto.RegisterMapType((map[string]string)(nil), "api.CreateRequestData.AttributesEntry")
	proto.RegisterType((*UpdateRequest)(nil), "api.UpdateRequest")
	proto.RegisterType((*UpdateRequestData)(nil), "api.UpdateRequestData")
	proto.RegisterMapType((map[string]string)(nil), "api.UpdateRequestData.AttributesEntry")
	proto.RegisterType((*FindRequest)(nil), "api.FindRequest")
	proto.RegisterMapType((map[string]string)(nil), "api.FindRequest.FilterEntry")
	proto.RegisterType((*FindResponse)(nil), "api.FindResponse")
	proto.RegisterType((*FindResponseData)(nil), "api.FindResponseData")
	proto.RegisterMapType((map[string]string)(nil), "api.FindResponseData.AttributesEntry")
	proto.RegisterType((*Metadata)(nil), "api.Metadata")
	proto.RegisterType((*Paginate)(nil), "api.Paginate")
	proto.RegisterType((*FindByIdRequestData)(nil), "api.FindByIdRequestData")
	proto.RegisterType((*FindByIdRequest)(nil), "api.FindByIdRequest")
	proto.RegisterType((*DeleteByIdRequestData)(nil), "api.DeleteByIdRequestData")
	proto.RegisterType((*DeleteByIdRequest)(nil), "api.DeleteByIdRequest")
	proto.RegisterType((*Response)(nil), "api.Response")
	proto.RegisterMapType((map[string]string)(nil), "api.Response.AttributesEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CrudClient is the client API for Crud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrudClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error)
	FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*Response, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	DeleteById(ctx context.Context, in *DeleteByIdRequest, opts ...grpc.CallOption) (*Response, error)
}

type crudClient struct {
	cc *grpc.ClientConn
}

func NewCrudClient(cc *grpc.ClientConn) CrudClient {
	return &crudClient{cc}
}

func (c *crudClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Crud/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Crud/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Crud/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/api.Crud/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) DeleteById(ctx context.Context, in *DeleteByIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Crud/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudServer is the server API for Crud service.
type CrudServer interface {
	Create(context.Context, *CreateRequest) (*Response, error)
	Update(context.Context, *UpdateRequest) (*Response, error)
	FindById(context.Context, *FindByIdRequest) (*Response, error)
	Find(context.Context, *FindRequest) (*FindResponse, error)
	DeleteById(context.Context, *DeleteByIdRequest) (*Response, error)
}

func RegisterCrudServer(s *grpc.Server, srv CrudServer) {
	s.RegisterService(&_Crud_serviceDesc, srv)
}

func _Crud_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Crud/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Crud/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Crud/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).FindById(ctx, req.(*FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Crud/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Crud/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).DeleteById(ctx, req.(*DeleteByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crud_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Crud",
	HandlerType: (*CrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Crud_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Crud_Update_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _Crud_FindById_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _Crud_Find_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _Crud_DeleteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/crud/crud.proto",
}

func init() { proto.RegisterFile("src/crud/crud.proto", fileDescriptor_8f6983d17f831b9d) }

var fileDescriptor_8f6983d17f831b9d = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xff, 0x6a, 0x13, 0x41,
	0x10, 0xee, 0xe6, 0x2e, 0xe1, 0x3a, 0x35, 0x36, 0x99, 0xb6, 0xe1, 0x08, 0x55, 0xca, 0x81, 0xd2,
	0x8a, 0x9e, 0x18, 0x2d, 0x58, 0x4b, 0x11, 0x9b, 0x36, 0xe0, 0x1f, 0x05, 0x39, 0xf0, 0x01, 0xb6,
	0xbd, 0xb5, 0x2c, 0xa6, 0xb9, 0xf3, 0x6e, 0x4f, 0xc8, 0x2b, 0xf8, 0x2c, 0xe2, 0x13, 0x88, 0xf8,
	0x62, 0x82, 0xdc, 0xee, 0xe6, 0x72, 0x3f, 0x16, 0x6c, 0x85, 0xfc, 0x13, 0x76, 0x67, 0x66, 0xe7,
	0x9b, 0xef, 0xbe, 0x6f, 0x37, 0xb0, 0x95, 0x26, 0x57, 0xcf, 0xaf, 0x92, 0x2c, 0x94, 0x3f, 0x7e,
	0x9c, 0x44, 0x22, 0x42, 0x8b, 0xc6, 0xdc, 0x3b, 0x86, 0xee, 0x38, 0x61, 0x54, 0xb0, 0x80, 0x7d,
	0xc9, 0x58, 0x2a, 0xf0, 0x09, 0xd8, 0x21, 0x15, 0xd4, 0x25, 0x7b, 0x64, 0x7f, 0x63, 0x34, 0xf0,
	0x69, 0xcc, 0xfd, 0x4a, 0xc5, 0x19, 0x15, 0x34, 0x90, 0x35, 0xde, 0x0f, 0x02, 0xfd, 0x46, 0x0e,
	0x11, 0x6c, 0x31, 0x8f, 0x99, 0xec, 0xb0, 0x1e, 0xc8, 0x35, 0x4e, 0x00, 0xa8, 0x10, 0x09, 0xbf,
	0xcc, 0x04, 0x4b, 0xdd, 0xd6, 0x9e, 0xb5, 0xbf, 0x31, 0x7a, 0x6c, 0xee, 0xed, 0xbf, 0x2b, 0x0a,
	0xcf, 0x67, 0x22, 0x99, 0x07, 0xa5, 0x93, 0xc3, 0x13, 0xd8, 0xac, 0xa5, 0xb1, 0x07, 0xd6, 0x67,
	0x36, 0xd7, 0x68, 0xf9, 0x12, 0xb7, 0xa1, 0xfd, 0x95, 0x4e, 0x33, 0xe6, 0xb6, 0x64, 0x4c, 0x6d,
	0xde, 0xb4, 0x5e, 0x93, 0x9c, 0xed, 0xc7, 0x38, 0xfc, 0x07, 0xdb, 0x4a, 0x45, 0x89, 0xed, 0x2f,
	0x02, 0xfd, 0x46, 0xce, 0xc8, 0xf6, 0x3e, 0xb4, 0x78, 0xa8, 0xd1, 0x5b, 0x3c, 0xac, 0xb1, 0xb7,
	0x4a, 0xec, 0x1b, 0xfd, 0x56, 0xc9, 0xfe, 0x37, 0x81, 0x8d, 0x09, 0x9f, 0x85, 0x0b, 0xf2, 0x08,
	0x76, 0x4c, 0xaf, 0xd5, 0xe8, 0xed, 0x40, 0xae, 0xf3, 0xd3, 0x53, 0x7e, 0xc3, 0x85, 0x3c, 0xdd,
	0x0e, 0xd4, 0x06, 0x07, 0xd0, 0x49, 0xa3, 0x44, 0x9c, 0xce, 0x5d, 0x4b, 0x36, 0xd5, 0x3b, 0x7c,
	0x05, 0x9d, 0x4f, 0x7c, 0x2a, 0x58, 0xe2, 0xda, 0x92, 0xd4, 0xae, 0x24, 0x55, 0xc2, 0xf0, 0x27,
	0x32, 0xad, 0xa8, 0xe8, 0xda, 0xe1, 0x51, 0x3e, 0x46, 0x11, 0xbe, 0x13, 0x85, 0x10, 0xee, 0xa9,
	0xee, 0x69, 0x1c, 0xcd, 0x52, 0x86, 0x07, 0xe0, 0xdc, 0x30, 0x41, 0x4b, 0x1a, 0x76, 0xe5, 0x08,
	0x17, 0x3a, 0x18, 0x14, 0x69, 0x3c, 0xd0, 0x52, 0x2b, 0xf3, 0xed, 0x94, 0x26, 0x55, 0xbd, 0x4a,
	0x4a, 0xff, 0x24, 0xd0, 0xab, 0xa7, 0x6e, 0x25, 0xf4, 0xb9, 0x41, 0xe8, 0x47, 0x46, 0xa4, 0x55,
	0xea, 0x7c, 0x08, 0xce, 0xc5, 0x92, 0xb5, 0x13, 0xd3, 0x6b, 0x3e, 0xa3, 0x82, 0x55, 0x3e, 0xd0,
	0x07, 0x1d, 0x0c, 0x8a, 0xb4, 0xf7, 0x8d, 0x80, 0xb3, 0x08, 0xe3, 0x43, 0x00, 0x11, 0x09, 0x3a,
	0x1d, 0x47, 0xd9, 0x4c, 0x68, 0x87, 0x94, 0x22, 0xb8, 0x0b, 0xeb, 0xb9, 0x5f, 0x54, 0x5a, 0x79,
	0x65, 0x19, 0x28, 0x9c, 0x65, 0x99, 0x9c, 0x65, 0x9b, 0x9d, 0xd5, 0x2e, 0x3b, 0xcb, 0x3b, 0x82,
	0xad, 0xfc, 0x93, 0x9d, 0xce, 0xdf, 0x87, 0x77, 0xbc, 0x6d, 0xde, 0x5b, 0xd8, 0xac, 0x1d, 0xc5,
	0xa7, 0x95, 0x6b, 0xee, 0x16, 0x8a, 0xd4, 0xda, 0x6b, 0xf9, 0x8f, 0x61, 0xe7, 0x8c, 0x4d, 0x99,
	0x60, 0xff, 0x83, 0x3e, 0x86, 0x7e, 0xe3, 0x30, 0xfa, 0x15, 0xfc, 0xa1, 0xc4, 0x37, 0x42, 0xe8,
	0x09, 0xbe, 0x13, 0x70, 0x0a, 0x8f, 0xdf, 0xc6, 0x78, 0x27, 0x06, 0xe3, 0x3d, 0x90, 0x30, 0x8b,
	0x36, 0x2b, 0x34, 0xdc, 0xe8, 0x0f, 0x01, 0x7b, 0x9c, 0x64, 0x21, 0x3e, 0x83, 0x8e, 0x7a, 0xcf,
	0x11, 0x9b, 0x8f, 0xfb, 0xb0, 0x5b, 0x19, 0xc8, 0x5b, 0xcb, 0xcb, 0xd5, 0x03, 0xa8, 0xcb, 0x2b,
	0xaf, 0x61, 0xb3, 0xfc, 0x05, 0x38, 0x0b, 0xd1, 0x70, 0xdb, 0xa4, 0xa1, 0x09, 0xc1, 0xce, 0x6b,
	0xb0, 0x57, 0x7f, 0x98, 0x86, 0xfd, 0xc6, 0xb5, 0xf4, 0xd6, 0xf0, 0x10, 0x60, 0x29, 0x0b, 0x0e,
	0xcc, 0x3a, 0x35, 0x50, 0x2e, 0x3b, 0xf2, 0x0f, 0xf5, 0xe5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0x90, 0x33, 0x1e, 0x67, 0x07, 0x00, 0x00,
}
