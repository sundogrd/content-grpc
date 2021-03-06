// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content/info.proto

package content

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

// request
type EContentState int32

const (
	EContentState_UNKNOWN_STATE EContentState = 0
	EContentState_DRAFT         EContentState = 1
	EContentState_CHECKING      EContentState = 2
	EContentState_PUBLISHED     EContentState = 3
	EContentState_DELETED       EContentState = 255
)

var EContentState_name = map[int32]string{
	0:   "UNKNOWN_STATE",
	1:   "DRAFT",
	2:   "CHECKING",
	3:   "PUBLISHED",
	255: "DELETED",
}

var EContentState_value = map[string]int32{
	"UNKNOWN_STATE": 0,
	"DRAFT":         1,
	"CHECKING":      2,
	"PUBLISHED":     3,
	"DELETED":       255,
}

func (x EContentState) String() string {
	return proto.EnumName(EContentState_name, int32(x))
}

func (EContentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{0}
}

type EContentType int32

const (
	EContentType_UNKNOWN_TYPE EContentType = 0
	EContentType_RICHTEXT     EContentType = 1
)

var EContentType_name = map[int32]string{
	0: "UNKNOWN_TYPE",
	1: "RICHTEXT",
}

var EContentType_value = map[string]int32{
	"UNKNOWN_TYPE": 0,
	"RICHTEXT":     1,
}

func (x EContentType) String() string {
	return proto.EnumName(EContentType_name, int32(x))
}

func (EContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{1}
}

type EContentBodyType int32

const (
	EContentBodyType_UNKNOWN_BODY_TYPE EContentBodyType = 0
	EContentBodyType_TEXT              EContentBodyType = 1
	EContentBodyType_HTML              EContentBodyType = 2
	EContentBodyType_MARKDOWN          EContentBodyType = 3
	EContentBodyType_LATEX             EContentBodyType = 4
)

var EContentBodyType_name = map[int32]string{
	0: "UNKNOWN_BODY_TYPE",
	1: "TEXT",
	2: "HTML",
	3: "MARKDOWN",
	4: "LATEX",
}

var EContentBodyType_value = map[string]int32{
	"UNKNOWN_BODY_TYPE": 0,
	"TEXT":              1,
	"HTML":              2,
	"MARKDOWN":          3,
	"LATEX":             4,
}

func (x EContentBodyType) String() string {
	return proto.EnumName(EContentBodyType_name, int32(x))
}

func (EContentBodyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{2}
}

type Content struct {
	ContentId            int64            `protobuf:"varint,1,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	AppId                string           `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Title                string           `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	AuthorId             int64            `protobuf:"varint,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Category             string           `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Type                 EContentType     `protobuf:"varint,7,opt,name=type,proto3,enum=content.EContentType" json:"type,omitempty"`
	Body                 string           `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	BodyType             EContentBodyType `protobuf:"varint,9,opt,name=body_type,json=bodyType,proto3,enum=content.EContentBodyType" json:"body_type,omitempty"`
	State                EContentState    `protobuf:"varint,10,opt,name=state,proto3,enum=content.EContentState" json:"state,omitempty"`
	Extra                string           `protobuf:"bytes,11,opt,name=extra,proto3" json:"extra,omitempty"`
	CreatedAt            uint32           `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            uint32           `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{0}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetContentId() int64 {
	if m != nil {
		return m.ContentId
	}
	return 0
}

func (m *Content) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Content) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Content) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Content) GetAuthorId() int64 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func (m *Content) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Content) GetType() EContentType {
	if m != nil {
		return m.Type
	}
	return EContentType_UNKNOWN_TYPE
}

func (m *Content) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Content) GetBodyType() EContentBodyType {
	if m != nil {
		return m.BodyType
	}
	return EContentBodyType_UNKNOWN_BODY_TYPE
}

func (m *Content) GetState() EContentState {
	if m != nil {
		return m.State
	}
	return EContentState_UNKNOWN_STATE
}

func (m *Content) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *Content) GetCreatedAt() uint32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Content) GetUpdatedAt() uint32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type ListContentsRequest struct {
	AppId                string        `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Page                 int32         `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32         `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	ContentIds           []int64       `protobuf:"varint,4,rep,packed,name=content_ids,json=contentIds,proto3" json:"content_ids,omitempty"`
	AuthorId             int64         `protobuf:"varint,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                int64         `protobuf:"varint,6,opt,name=title,proto3" json:"title,omitempty"`
	Type                 EContentType  `protobuf:"varint,7,opt,name=type,proto3,enum=content.EContentType" json:"type,omitempty"`
	State                EContentState `protobuf:"varint,8,opt,name=state,proto3,enum=content.EContentState" json:"state,omitempty"`
	StartCreatedAt       uint32        `protobuf:"varint,9,opt,name=start_created_at,json=startCreatedAt,proto3" json:"start_created_at,omitempty"`
	EndCreatedAt         uint32        `protobuf:"varint,10,opt,name=end_created_at,json=endCreatedAt,proto3" json:"end_created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListContentsRequest) Reset()         { *m = ListContentsRequest{} }
func (m *ListContentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListContentsRequest) ProtoMessage()    {}
func (*ListContentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{1}
}

func (m *ListContentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContentsRequest.Unmarshal(m, b)
}
func (m *ListContentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContentsRequest.Marshal(b, m, deterministic)
}
func (m *ListContentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContentsRequest.Merge(m, src)
}
func (m *ListContentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListContentsRequest.Size(m)
}
func (m *ListContentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListContentsRequest proto.InternalMessageInfo

func (m *ListContentsRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ListContentsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListContentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListContentsRequest) GetContentIds() []int64 {
	if m != nil {
		return m.ContentIds
	}
	return nil
}

func (m *ListContentsRequest) GetAuthorId() int64 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func (m *ListContentsRequest) GetTitle() int64 {
	if m != nil {
		return m.Title
	}
	return 0
}

func (m *ListContentsRequest) GetType() EContentType {
	if m != nil {
		return m.Type
	}
	return EContentType_UNKNOWN_TYPE
}

func (m *ListContentsRequest) GetState() EContentState {
	if m != nil {
		return m.State
	}
	return EContentState_UNKNOWN_STATE
}

func (m *ListContentsRequest) GetStartCreatedAt() uint32 {
	if m != nil {
		return m.StartCreatedAt
	}
	return 0
}

func (m *ListContentsRequest) GetEndCreatedAt() uint32 {
	if m != nil {
		return m.EndCreatedAt
	}
	return 0
}

type ListContentsResponse struct {
	AppId                string     `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Contents             []*Content `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents,omitempty"`
	Total                int64      `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Page                 int32      `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32      `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListContentsResponse) Reset()         { *m = ListContentsResponse{} }
func (m *ListContentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListContentsResponse) ProtoMessage()    {}
func (*ListContentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{2}
}

func (m *ListContentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContentsResponse.Unmarshal(m, b)
}
func (m *ListContentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContentsResponse.Marshal(b, m, deterministic)
}
func (m *ListContentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContentsResponse.Merge(m, src)
}
func (m *ListContentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListContentsResponse.Size(m)
}
func (m *ListContentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListContentsResponse proto.InternalMessageInfo

func (m *ListContentsResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ListContentsResponse) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *ListContentsResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListContentsResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListContentsResponse) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type GetContentByIdRequest struct {
	ContentId            int64    `protobuf:"varint,1,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContentByIdRequest) Reset()         { *m = GetContentByIdRequest{} }
func (m *GetContentByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetContentByIdRequest) ProtoMessage()    {}
func (*GetContentByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{3}
}

func (m *GetContentByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContentByIdRequest.Unmarshal(m, b)
}
func (m *GetContentByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContentByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetContentByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContentByIdRequest.Merge(m, src)
}
func (m *GetContentByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetContentByIdRequest.Size(m)
}
func (m *GetContentByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContentByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContentByIdRequest proto.InternalMessageInfo

func (m *GetContentByIdRequest) GetContentId() int64 {
	if m != nil {
		return m.ContentId
	}
	return 0
}

type GetContentByIdResponse struct {
	Content              *Content `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContentByIdResponse) Reset()         { *m = GetContentByIdResponse{} }
func (m *GetContentByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetContentByIdResponse) ProtoMessage()    {}
func (*GetContentByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{4}
}

func (m *GetContentByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContentByIdResponse.Unmarshal(m, b)
}
func (m *GetContentByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContentByIdResponse.Marshal(b, m, deterministic)
}
func (m *GetContentByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContentByIdResponse.Merge(m, src)
}
func (m *GetContentByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetContentByIdResponse.Size(m)
}
func (m *GetContentByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContentByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetContentByIdResponse proto.InternalMessageInfo

func (m *GetContentByIdResponse) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

type CreateContentRequest struct {
	AppId                string           `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Title                string           `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AuthorId             int64            `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Category             string           `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	State                EContentState    `protobuf:"varint,6,opt,name=state,proto3,enum=content.EContentState" json:"state,omitempty"`
	Type                 EContentType     `protobuf:"varint,7,opt,name=type,proto3,enum=content.EContentType" json:"type,omitempty"`
	Body                 string           `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	BodyType             EContentBodyType `protobuf:"varint,9,opt,name=body_type,json=bodyType,proto3,enum=content.EContentBodyType" json:"body_type,omitempty"`
	Extra                string           `protobuf:"bytes,10,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateContentRequest) Reset()         { *m = CreateContentRequest{} }
func (m *CreateContentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateContentRequest) ProtoMessage()    {}
func (*CreateContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{5}
}

func (m *CreateContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateContentRequest.Unmarshal(m, b)
}
func (m *CreateContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateContentRequest.Marshal(b, m, deterministic)
}
func (m *CreateContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateContentRequest.Merge(m, src)
}
func (m *CreateContentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateContentRequest.Size(m)
}
func (m *CreateContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateContentRequest proto.InternalMessageInfo

func (m *CreateContentRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateContentRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateContentRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateContentRequest) GetAuthorId() int64 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func (m *CreateContentRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *CreateContentRequest) GetState() EContentState {
	if m != nil {
		return m.State
	}
	return EContentState_UNKNOWN_STATE
}

func (m *CreateContentRequest) GetType() EContentType {
	if m != nil {
		return m.Type
	}
	return EContentType_UNKNOWN_TYPE
}

func (m *CreateContentRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *CreateContentRequest) GetBodyType() EContentBodyType {
	if m != nil {
		return m.BodyType
	}
	return EContentBodyType_UNKNOWN_BODY_TYPE
}

func (m *CreateContentRequest) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

type CreateContentResponse struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Content              *Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateContentResponse) Reset()         { *m = CreateContentResponse{} }
func (m *CreateContentResponse) String() string { return proto.CompactTextString(m) }
func (*CreateContentResponse) ProtoMessage()    {}
func (*CreateContentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{6}
}

func (m *CreateContentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateContentResponse.Unmarshal(m, b)
}
func (m *CreateContentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateContentResponse.Marshal(b, m, deterministic)
}
func (m *CreateContentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateContentResponse.Merge(m, src)
}
func (m *CreateContentResponse) XXX_Size() int {
	return xxx_messageInfo_CreateContentResponse.Size(m)
}
func (m *CreateContentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateContentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateContentResponse proto.InternalMessageInfo

func (m *CreateContentResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *CreateContentResponse) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

type DeleteContentRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ContentId            int64    `protobuf:"varint,2,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteContentRequest) Reset()         { *m = DeleteContentRequest{} }
func (m *DeleteContentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteContentRequest) ProtoMessage()    {}
func (*DeleteContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{7}
}

func (m *DeleteContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteContentRequest.Unmarshal(m, b)
}
func (m *DeleteContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteContentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteContentRequest.Merge(m, src)
}
func (m *DeleteContentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteContentRequest.Size(m)
}
func (m *DeleteContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteContentRequest proto.InternalMessageInfo

func (m *DeleteContentRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeleteContentRequest) GetContentId() int64 {
	if m != nil {
		return m.ContentId
	}
	return 0
}

type DeleteContentResponse struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ContentId            int64    `protobuf:"varint,2,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteContentResponse) Reset()         { *m = DeleteContentResponse{} }
func (m *DeleteContentResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteContentResponse) ProtoMessage()    {}
func (*DeleteContentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9eb484a5008e3e, []int{8}
}

func (m *DeleteContentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteContentResponse.Unmarshal(m, b)
}
func (m *DeleteContentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteContentResponse.Marshal(b, m, deterministic)
}
func (m *DeleteContentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteContentResponse.Merge(m, src)
}
func (m *DeleteContentResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteContentResponse.Size(m)
}
func (m *DeleteContentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteContentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteContentResponse proto.InternalMessageInfo

func (m *DeleteContentResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeleteContentResponse) GetContentId() int64 {
	if m != nil {
		return m.ContentId
	}
	return 0
}

func init() {
	proto.RegisterEnum("content.EContentState", EContentState_name, EContentState_value)
	proto.RegisterEnum("content.EContentType", EContentType_name, EContentType_value)
	proto.RegisterEnum("content.EContentBodyType", EContentBodyType_name, EContentBodyType_value)
	proto.RegisterType((*Content)(nil), "content.Content")
	proto.RegisterType((*ListContentsRequest)(nil), "content.ListContentsRequest")
	proto.RegisterType((*ListContentsResponse)(nil), "content.ListContentsResponse")
	proto.RegisterType((*GetContentByIdRequest)(nil), "content.GetContentByIdRequest")
	proto.RegisterType((*GetContentByIdResponse)(nil), "content.GetContentByIdResponse")
	proto.RegisterType((*CreateContentRequest)(nil), "content.CreateContentRequest")
	proto.RegisterType((*CreateContentResponse)(nil), "content.CreateContentResponse")
	proto.RegisterType((*DeleteContentRequest)(nil), "content.DeleteContentRequest")
	proto.RegisterType((*DeleteContentResponse)(nil), "content.DeleteContentResponse")
}

func init() { proto.RegisterFile("content/info.proto", fileDescriptor_1e9eb484a5008e3e) }

var fileDescriptor_1e9eb484a5008e3e = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x5d, 0x6f, 0xda, 0x48,
	0x14, 0x5d, 0x7f, 0x81, 0x7d, 0x31, 0xc8, 0x99, 0x85, 0xc8, 0xcb, 0x6e, 0x12, 0x84, 0xf6, 0x81,
	0x45, 0x11, 0x2b, 0x65, 0xa5, 0xbc, 0x03, 0xf6, 0x06, 0x04, 0x21, 0xa9, 0x71, 0x9a, 0xa4, 0x2f,
	0xc8, 0xc1, 0xd3, 0xd4, 0x52, 0x84, 0x5d, 0x7b, 0x52, 0x95, 0xfc, 0x9c, 0xbe, 0xf5, 0xdf, 0xf4,
	0xb5, 0xfd, 0x33, 0xad, 0x3c, 0xfe, 0x88, 0x21, 0xe0, 0xa4, 0x4f, 0x7d, 0xf2, 0xcc, 0xbd, 0x73,
	0x8f, 0xef, 0x9c, 0x73, 0x7c, 0x65, 0x40, 0x73, 0x77, 0x41, 0xf0, 0x82, 0xfc, 0xeb, 0x2c, 0xde,
	0xba, 0x1d, 0xcf, 0x77, 0x89, 0x8b, 0x8a, 0x71, 0xac, 0xf9, 0x99, 0x83, 0x62, 0x3f, 0x5a, 0xa3,
	0x3d, 0x80, 0x38, 0x3c, 0x73, 0x6c, 0x95, 0x69, 0x30, 0x2d, 0xce, 0x90, 0xe2, 0xc8, 0xd0, 0x46,
	0x35, 0x28, 0x58, 0x9e, 0x17, 0xa6, 0xd8, 0x06, 0xd3, 0x92, 0x0c, 0xc1, 0xf2, 0xbc, 0xa1, 0x8d,
	0xaa, 0x20, 0x10, 0x87, 0xdc, 0x61, 0x95, 0x8b, 0xa2, 0x74, 0x83, 0x1a, 0x50, 0xb2, 0x71, 0x30,
	0xf7, 0x1d, 0x8f, 0x38, 0xee, 0x42, 0xe5, 0x69, 0x2e, 0x1b, 0x42, 0x7f, 0x82, 0x64, 0xdd, 0x93,
	0x77, 0xae, 0x1f, 0x22, 0x0a, 0xf4, 0x65, 0x62, 0x14, 0x18, 0xda, 0xa8, 0x0e, 0xe2, 0xdc, 0x22,
	0xf8, 0xd6, 0xf5, 0x97, 0x6a, 0x81, 0xd6, 0xa6, 0x7b, 0xf4, 0x0f, 0xf0, 0x64, 0xe9, 0x61, 0xb5,
	0xd8, 0x60, 0x5a, 0x95, 0xa3, 0x5a, 0x27, 0xee, 0xb0, 0xa3, 0xc7, 0xf7, 0x30, 0x97, 0x1e, 0x36,
	0xe8, 0x11, 0x84, 0x80, 0xbf, 0x71, 0xed, 0xa5, 0x2a, 0x52, 0x08, 0xba, 0x46, 0xc7, 0x20, 0x85,
	0xcf, 0x19, 0xc5, 0x90, 0x28, 0xc6, 0x1f, 0x4f, 0x30, 0x7a, 0xae, 0xbd, 0xa4, 0x38, 0xe2, 0x4d,
	0xbc, 0x42, 0x87, 0x20, 0x04, 0xc4, 0x22, 0x58, 0x05, 0x5a, 0xb3, 0xfb, 0xa4, 0x66, 0x1a, 0x66,
	0x8d, 0xe8, 0x50, 0xc8, 0x0a, 0xfe, 0x48, 0x7c, 0x4b, 0x2d, 0x45, 0xac, 0xd0, 0x0d, 0x65, 0xd8,
	0xc7, 0x16, 0xc1, 0xf6, 0xcc, 0x22, 0xaa, 0xdc, 0x60, 0x5a, 0x65, 0x43, 0x8a, 0x23, 0x5d, 0x2a,
	0xc0, 0xbd, 0x67, 0x27, 0xe9, 0x72, 0x94, 0x8e, 0x23, 0x5d, 0xd2, 0xfc, 0xc6, 0xc2, 0xef, 0x63,
	0x27, 0x20, 0xf1, 0xfb, 0x02, 0x03, 0xbf, 0xbf, 0xc7, 0x01, 0xc9, 0x08, 0xc3, 0x64, 0x85, 0x41,
	0xc0, 0x7b, 0xd6, 0x2d, 0xa6, 0x6a, 0x09, 0x06, 0x5d, 0x87, 0xa4, 0x87, 0xcf, 0x59, 0xe0, 0x3c,
	0x44, 0x82, 0x09, 0x86, 0x18, 0x06, 0xa6, 0xce, 0x03, 0x46, 0x07, 0x50, 0x7a, 0xd4, 0x3f, 0x50,
	0xf9, 0x06, 0xd7, 0xe2, 0x0c, 0x48, 0x0d, 0x10, 0xe4, 0x4b, 0x96, 0xfa, 0xa0, 0x40, 0x13, 0xb1,
	0x0f, 0x7e, 0x42, 0xac, 0x94, 0x60, 0xf1, 0x25, 0x04, 0xb7, 0x40, 0x09, 0x88, 0xe5, 0x93, 0x59,
	0x86, 0x50, 0x89, 0x32, 0x56, 0xa1, 0xf1, 0x7e, 0xca, 0xea, 0xdf, 0x50, 0xc1, 0x0b, 0x3b, 0x7b,
	0x0e, 0xe8, 0x39, 0x19, 0x2f, 0xec, 0xf4, 0x54, 0xf3, 0x13, 0x03, 0xd5, 0x55, 0x72, 0x03, 0xcf,
	0x5d, 0x04, 0x78, 0x1b, 0xbb, 0x87, 0x20, 0xc6, 0xfd, 0x05, 0x2a, 0xdb, 0xe0, 0x5a, 0xa5, 0x23,
	0x25, 0x6d, 0x38, 0xc6, 0x30, 0xd2, 0x13, 0x94, 0x1c, 0x97, 0x58, 0x77, 0x94, 0xf3, 0x90, 0x9c,
	0x70, 0x93, 0x2a, 0xc4, 0x6f, 0x53, 0x48, 0x58, 0x55, 0xa8, 0x79, 0x0c, 0xb5, 0x13, 0x9c, 0xb4,
	0xd8, 0x5b, 0x0e, 0xed, 0xc4, 0x02, 0xf9, 0x9f, 0x6e, 0x53, 0x83, 0xdd, 0xf5, 0xba, 0xf8, 0x76,
	0x6d, 0x48, 0x46, 0x01, 0xad, 0xda, 0x74, 0x8b, 0x74, 0x56, 0x7c, 0x65, 0xa1, 0x1a, 0x11, 0x96,
	0xa4, 0xf2, 0x0d, 0x98, 0x3a, 0x82, 0xcd, 0x99, 0x0c, 0xdc, 0x33, 0x93, 0x81, 0xcf, 0x99, 0x0c,
	0xc2, 0xda, 0x64, 0x48, 0x1d, 0x54, 0x78, 0x89, 0x83, 0x7e, 0xd1, 0x1c, 0x49, 0x27, 0x03, 0x64,
	0x26, 0x43, 0xf3, 0x0d, 0xd4, 0xd6, 0xa8, 0xcd, 0xb7, 0x5f, 0x46, 0x37, 0xf6, 0x39, 0xdd, 0xc6,
	0x50, 0xd5, 0xf0, 0x1d, 0x7e, 0xa9, 0x6c, 0xab, 0x5e, 0x62, 0xd7, 0xbd, 0x74, 0x0a, 0xb5, 0x35,
	0xb4, 0xfc, 0x4e, 0xf3, 0xe1, 0xda, 0x57, 0x50, 0x5e, 0x51, 0x07, 0xed, 0x40, 0xf9, 0x62, 0x32,
	0x9a, 0x9c, 0x5d, 0x4e, 0x66, 0x53, 0xb3, 0x6b, 0xea, 0xca, 0x6f, 0x48, 0x02, 0x41, 0x33, 0xba,
	0xff, 0x9b, 0x0a, 0x83, 0x64, 0x10, 0xfb, 0x03, 0xbd, 0x3f, 0x1a, 0x4e, 0x4e, 0x14, 0x16, 0x95,
	0x41, 0x3a, 0xbf, 0xe8, 0x8d, 0x87, 0xd3, 0x81, 0xae, 0x29, 0x1c, 0x92, 0xa1, 0xa8, 0xe9, 0x63,
	0xdd, 0xd4, 0x35, 0xe5, 0x3b, 0xd3, 0xee, 0x80, 0x9c, 0x95, 0x12, 0x29, 0x20, 0x27, 0xc0, 0xe6,
	0xf5, 0x79, 0x88, 0x2b, 0x83, 0x68, 0x0c, 0xfb, 0x03, 0x53, 0xbf, 0x32, 0x15, 0xa6, 0xfd, 0x1a,
	0x94, 0x75, 0xd9, 0x50, 0x0d, 0x76, 0x92, 0x9a, 0xde, 0x99, 0x76, 0x9d, 0x14, 0x8a, 0xc0, 0x47,
	0x45, 0xe1, 0x6a, 0x60, 0x9e, 0x8e, 0x15, 0x36, 0x04, 0x3b, 0xed, 0x1a, 0x23, 0xed, 0xec, 0x72,
	0xa2, 0x70, 0x61, 0xcb, 0xe3, 0xae, 0xa9, 0x5f, 0x29, 0xfc, 0xd1, 0x17, 0x16, 0x2a, 0xc9, 0x0d,
	0xb1, 0xff, 0xc1, 0x99, 0x63, 0x34, 0x02, 0x39, 0x3b, 0x6b, 0xd0, 0x5f, 0xa9, 0x78, 0x1b, 0xe6,
	0x7b, 0x7d, 0x6f, 0x4b, 0x36, 0xe6, 0xfd, 0x15, 0x54, 0x56, 0x3f, 0x6e, 0xb4, 0x9f, 0x16, 0x6c,
	0x9c, 0x16, 0xf5, 0x83, 0xad, 0xf9, 0x18, 0x72, 0x02, 0xe5, 0x15, 0x37, 0xa2, 0xc7, 0x16, 0x36,
	0x0d, 0x80, 0xfa, 0xfe, 0xb6, 0xf4, 0x23, 0xde, 0x8a, 0x67, 0x32, 0x78, 0x9b, 0x9c, 0x99, 0xc1,
	0xdb, 0x68, 0xb5, 0x9b, 0x02, 0xfd, 0x8b, 0xf9, 0xef, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda,
	0xc0, 0x36, 0x72, 0xdb, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentServiceClient interface {
	ListContents(ctx context.Context, in *ListContentsRequest, opts ...grpc.CallOption) (*ListContentsResponse, error)
	GetContentById(ctx context.Context, in *GetContentByIdRequest, opts ...grpc.CallOption) (*GetContentByIdResponse, error)
	CreateContent(ctx context.Context, in *CreateContentRequest, opts ...grpc.CallOption) (*CreateContentResponse, error)
	DeleteContent(ctx context.Context, in *DeleteContentRequest, opts ...grpc.CallOption) (*DeleteContentResponse, error)
}

type contentServiceClient struct {
	cc *grpc.ClientConn
}

func NewContentServiceClient(cc *grpc.ClientConn) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) ListContents(ctx context.Context, in *ListContentsRequest, opts ...grpc.CallOption) (*ListContentsResponse, error) {
	out := new(ListContentsResponse)
	err := c.cc.Invoke(ctx, "/content.ContentService/ListContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetContentById(ctx context.Context, in *GetContentByIdRequest, opts ...grpc.CallOption) (*GetContentByIdResponse, error) {
	out := new(GetContentByIdResponse)
	err := c.cc.Invoke(ctx, "/content.ContentService/GetContentById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) CreateContent(ctx context.Context, in *CreateContentRequest, opts ...grpc.CallOption) (*CreateContentResponse, error) {
	out := new(CreateContentResponse)
	err := c.cc.Invoke(ctx, "/content.ContentService/CreateContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteContent(ctx context.Context, in *DeleteContentRequest, opts ...grpc.CallOption) (*DeleteContentResponse, error) {
	out := new(DeleteContentResponse)
	err := c.cc.Invoke(ctx, "/content.ContentService/DeleteContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
type ContentServiceServer interface {
	ListContents(context.Context, *ListContentsRequest) (*ListContentsResponse, error)
	GetContentById(context.Context, *GetContentByIdRequest) (*GetContentByIdResponse, error)
	CreateContent(context.Context, *CreateContentRequest) (*CreateContentResponse, error)
	DeleteContent(context.Context, *DeleteContentRequest) (*DeleteContentResponse, error)
}

// UnimplementedContentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (*UnimplementedContentServiceServer) ListContents(ctx context.Context, req *ListContentsRequest) (*ListContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContents not implemented")
}
func (*UnimplementedContentServiceServer) GetContentById(ctx context.Context, req *GetContentByIdRequest) (*GetContentByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContentById not implemented")
}
func (*UnimplementedContentServiceServer) CreateContent(ctx context.Context, req *CreateContentRequest) (*CreateContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContent not implemented")
}
func (*UnimplementedContentServiceServer) DeleteContent(ctx context.Context, req *DeleteContentRequest) (*DeleteContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContent not implemented")
}

func RegisterContentServiceServer(s *grpc.Server, srv ContentServiceServer) {
	s.RegisterService(&_ContentService_serviceDesc, srv)
}

func _ContentService_ListContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).ListContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/ListContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).ListContents(ctx, req.(*ListContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetContentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetContentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/GetContentById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetContentById(ctx, req.(*GetContentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_CreateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).CreateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/CreateContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).CreateContent(ctx, req.(*CreateContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/DeleteContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteContent(ctx, req.(*DeleteContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "content.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContents",
			Handler:    _ContentService_ListContents_Handler,
		},
		{
			MethodName: "GetContentById",
			Handler:    _ContentService_GetContentById_Handler,
		},
		{
			MethodName: "CreateContent",
			Handler:    _ContentService_CreateContent_Handler,
		},
		{
			MethodName: "DeleteContent",
			Handler:    _ContentService_DeleteContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/info.proto",
}
