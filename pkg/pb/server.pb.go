// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Label struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type LabelSet struct {
	Labels               []*Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelSet) Reset()         { *m = LabelSet{} }
func (m *LabelSet) String() string { return proto.CompactTextString(m) }
func (*LabelSet) ProtoMessage()    {}
func (*LabelSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *LabelSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelSet.Unmarshal(m, b)
}
func (m *LabelSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelSet.Marshal(b, m, deterministic)
}
func (m *LabelSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSet.Merge(m, src)
}
func (m *LabelSet) XXX_Size() int {
	return xxx_messageInfo_LabelSet.Size(m)
}
func (m *LabelSet) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSet.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSet proto.InternalMessageInfo

func (m *LabelSet) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type RegisterGuestAccountRequest struct {
	// server ID is the unique ULID of the Waypoint server requesting a
	// guest account. If this server already has a guest account registered,
	// the same token will be returned.
	ServerId             string   `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterGuestAccountRequest) Reset()         { *m = RegisterGuestAccountRequest{} }
func (m *RegisterGuestAccountRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterGuestAccountRequest) ProtoMessage()    {}
func (*RegisterGuestAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *RegisterGuestAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterGuestAccountRequest.Unmarshal(m, b)
}
func (m *RegisterGuestAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterGuestAccountRequest.Marshal(b, m, deterministic)
}
func (m *RegisterGuestAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterGuestAccountRequest.Merge(m, src)
}
func (m *RegisterGuestAccountRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterGuestAccountRequest.Size(m)
}
func (m *RegisterGuestAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterGuestAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterGuestAccountRequest proto.InternalMessageInfo

func (m *RegisterGuestAccountRequest) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type RegisterGuestAccountResponse struct {
	// API token to use for protected endpoints.
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterGuestAccountResponse) Reset()         { *m = RegisterGuestAccountResponse{} }
func (m *RegisterGuestAccountResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterGuestAccountResponse) ProtoMessage()    {}
func (*RegisterGuestAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *RegisterGuestAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterGuestAccountResponse.Unmarshal(m, b)
}
func (m *RegisterGuestAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterGuestAccountResponse.Marshal(b, m, deterministic)
}
func (m *RegisterGuestAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterGuestAccountResponse.Merge(m, src)
}
func (m *RegisterGuestAccountResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterGuestAccountResponse.Size(m)
}
func (m *RegisterGuestAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterGuestAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterGuestAccountResponse proto.InternalMessageInfo

func (m *RegisterGuestAccountResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RegisterHostnameRequest struct {
	// hostname to register
	//
	// Types that are valid to be assigned to Hostname:
	//	*RegisterHostnameRequest_Generate
	//	*RegisterHostnameRequest_Exact
	Hostname isRegisterHostnameRequest_Hostname `protobuf_oneof:"hostname"`
	// labels to link this hostname to.
	Labels               *LabelSet `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterHostnameRequest) Reset()         { *m = RegisterHostnameRequest{} }
func (m *RegisterHostnameRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterHostnameRequest) ProtoMessage()    {}
func (*RegisterHostnameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *RegisterHostnameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterHostnameRequest.Unmarshal(m, b)
}
func (m *RegisterHostnameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterHostnameRequest.Marshal(b, m, deterministic)
}
func (m *RegisterHostnameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterHostnameRequest.Merge(m, src)
}
func (m *RegisterHostnameRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterHostnameRequest.Size(m)
}
func (m *RegisterHostnameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterHostnameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterHostnameRequest proto.InternalMessageInfo

type isRegisterHostnameRequest_Hostname interface {
	isRegisterHostnameRequest_Hostname()
}

type RegisterHostnameRequest_Generate struct {
	Generate *empty.Empty `protobuf:"bytes,1,opt,name=generate,proto3,oneof"`
}

type RegisterHostnameRequest_Exact struct {
	Exact string `protobuf:"bytes,2,opt,name=exact,proto3,oneof"`
}

func (*RegisterHostnameRequest_Generate) isRegisterHostnameRequest_Hostname() {}

func (*RegisterHostnameRequest_Exact) isRegisterHostnameRequest_Hostname() {}

func (m *RegisterHostnameRequest) GetHostname() isRegisterHostnameRequest_Hostname {
	if m != nil {
		return m.Hostname
	}
	return nil
}

func (m *RegisterHostnameRequest) GetGenerate() *empty.Empty {
	if x, ok := m.GetHostname().(*RegisterHostnameRequest_Generate); ok {
		return x.Generate
	}
	return nil
}

func (m *RegisterHostnameRequest) GetExact() string {
	if x, ok := m.GetHostname().(*RegisterHostnameRequest_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *RegisterHostnameRequest) GetLabels() *LabelSet {
	if m != nil {
		return m.Labels
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegisterHostnameRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegisterHostnameRequest_Generate)(nil),
		(*RegisterHostnameRequest_Exact)(nil),
	}
}

type RegisterHostnameResponse struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Fqdn                 string   `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterHostnameResponse) Reset()         { *m = RegisterHostnameResponse{} }
func (m *RegisterHostnameResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterHostnameResponse) ProtoMessage()    {}
func (*RegisterHostnameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *RegisterHostnameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterHostnameResponse.Unmarshal(m, b)
}
func (m *RegisterHostnameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterHostnameResponse.Marshal(b, m, deterministic)
}
func (m *RegisterHostnameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterHostnameResponse.Merge(m, src)
}
func (m *RegisterHostnameResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterHostnameResponse.Size(m)
}
func (m *RegisterHostnameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterHostnameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterHostnameResponse proto.InternalMessageInfo

func (m *RegisterHostnameResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *RegisterHostnameResponse) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

type ListHostnamesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHostnamesRequest) Reset()         { *m = ListHostnamesRequest{} }
func (m *ListHostnamesRequest) String() string { return proto.CompactTextString(m) }
func (*ListHostnamesRequest) ProtoMessage()    {}
func (*ListHostnamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{6}
}

func (m *ListHostnamesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHostnamesRequest.Unmarshal(m, b)
}
func (m *ListHostnamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHostnamesRequest.Marshal(b, m, deterministic)
}
func (m *ListHostnamesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHostnamesRequest.Merge(m, src)
}
func (m *ListHostnamesRequest) XXX_Size() int {
	return xxx_messageInfo_ListHostnamesRequest.Size(m)
}
func (m *ListHostnamesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHostnamesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListHostnamesRequest proto.InternalMessageInfo

type ListHostnamesResponse struct {
	Hostnames            []*ListHostnamesResponse_Hostname `protobuf:"bytes,1,rep,name=hostnames,proto3" json:"hostnames,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ListHostnamesResponse) Reset()         { *m = ListHostnamesResponse{} }
func (m *ListHostnamesResponse) String() string { return proto.CompactTextString(m) }
func (*ListHostnamesResponse) ProtoMessage()    {}
func (*ListHostnamesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{7}
}

func (m *ListHostnamesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHostnamesResponse.Unmarshal(m, b)
}
func (m *ListHostnamesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHostnamesResponse.Marshal(b, m, deterministic)
}
func (m *ListHostnamesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHostnamesResponse.Merge(m, src)
}
func (m *ListHostnamesResponse) XXX_Size() int {
	return xxx_messageInfo_ListHostnamesResponse.Size(m)
}
func (m *ListHostnamesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHostnamesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListHostnamesResponse proto.InternalMessageInfo

func (m *ListHostnamesResponse) GetHostnames() []*ListHostnamesResponse_Hostname {
	if m != nil {
		return m.Hostnames
	}
	return nil
}

type ListHostnamesResponse_Hostname struct {
	Hostname             string    `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Fqdn                 string    `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Labels               *LabelSet `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListHostnamesResponse_Hostname) Reset()         { *m = ListHostnamesResponse_Hostname{} }
func (m *ListHostnamesResponse_Hostname) String() string { return proto.CompactTextString(m) }
func (*ListHostnamesResponse_Hostname) ProtoMessage()    {}
func (*ListHostnamesResponse_Hostname) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{7, 0}
}

func (m *ListHostnamesResponse_Hostname) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHostnamesResponse_Hostname.Unmarshal(m, b)
}
func (m *ListHostnamesResponse_Hostname) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHostnamesResponse_Hostname.Marshal(b, m, deterministic)
}
func (m *ListHostnamesResponse_Hostname) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHostnamesResponse_Hostname.Merge(m, src)
}
func (m *ListHostnamesResponse_Hostname) XXX_Size() int {
	return xxx_messageInfo_ListHostnamesResponse_Hostname.Size(m)
}
func (m *ListHostnamesResponse_Hostname) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHostnamesResponse_Hostname.DiscardUnknown(m)
}

var xxx_messageInfo_ListHostnamesResponse_Hostname proto.InternalMessageInfo

func (m *ListHostnamesResponse_Hostname) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ListHostnamesResponse_Hostname) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *ListHostnamesResponse_Hostname) GetLabels() *LabelSet {
	if m != nil {
		return m.Labels
	}
	return nil
}

type DeleteHostnameRequest struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteHostnameRequest) Reset()         { *m = DeleteHostnameRequest{} }
func (m *DeleteHostnameRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteHostnameRequest) ProtoMessage()    {}
func (*DeleteHostnameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{8}
}

func (m *DeleteHostnameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteHostnameRequest.Unmarshal(m, b)
}
func (m *DeleteHostnameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteHostnameRequest.Marshal(b, m, deterministic)
}
func (m *DeleteHostnameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteHostnameRequest.Merge(m, src)
}
func (m *DeleteHostnameRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteHostnameRequest.Size(m)
}
func (m *DeleteHostnameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteHostnameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteHostnameRequest proto.InternalMessageInfo

func (m *DeleteHostnameRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*Label)(nil), "hashicorp.waypoint_hzn.Label")
	proto.RegisterType((*LabelSet)(nil), "hashicorp.waypoint_hzn.LabelSet")
	proto.RegisterType((*RegisterGuestAccountRequest)(nil), "hashicorp.waypoint_hzn.RegisterGuestAccountRequest")
	proto.RegisterType((*RegisterGuestAccountResponse)(nil), "hashicorp.waypoint_hzn.RegisterGuestAccountResponse")
	proto.RegisterType((*RegisterHostnameRequest)(nil), "hashicorp.waypoint_hzn.RegisterHostnameRequest")
	proto.RegisterType((*RegisterHostnameResponse)(nil), "hashicorp.waypoint_hzn.RegisterHostnameResponse")
	proto.RegisterType((*ListHostnamesRequest)(nil), "hashicorp.waypoint_hzn.ListHostnamesRequest")
	proto.RegisterType((*ListHostnamesResponse)(nil), "hashicorp.waypoint_hzn.ListHostnamesResponse")
	proto.RegisterType((*ListHostnamesResponse_Hostname)(nil), "hashicorp.waypoint_hzn.ListHostnamesResponse.Hostname")
	proto.RegisterType((*DeleteHostnameRequest)(nil), "hashicorp.waypoint_hzn.DeleteHostnameRequest")
}

func init() {
	proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7)
}

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5d, 0x6f, 0xd2, 0x50,
	0x18, 0xc7, 0x57, 0x0a, 0xb3, 0x7b, 0xf0, 0x85, 0x1c, 0x19, 0x92, 0xa2, 0x09, 0xe9, 0xd5, 0xa2,
	0xae, 0x28, 0x10, 0xe3, 0x4b, 0xa2, 0xb1, 0xd1, 0x88, 0xba, 0xab, 0xce, 0xb8, 0xa4, 0x62, 0x96,
	0x42, 0x1f, 0x68, 0x63, 0x39, 0xa7, 0x6b, 0x0f, 0xe0, 0xb8, 0xf3, 0xd6, 0x4b, 0xbf, 0x98, 0x1f,
	0xc2, 0x2f, 0x61, 0x16, 0x2f, 0x0c, 0x3d, 0x6d, 0x97, 0x21, 0xe0, 0xb8, 0x3b, 0xe7, 0xe1, 0xfc,
	0x9f, 0x97, 0xdf, 0xf3, 0xa7, 0x70, 0x35, 0xc2, 0x70, 0x82, 0xa1, 0x1e, 0x84, 0x8c, 0x33, 0x52,
	0x71, 0xed, 0xc8, 0xf5, 0xfa, 0x2c, 0x0c, 0xf4, 0xa9, 0x7d, 0x1a, 0x30, 0x8f, 0xf2, 0x63, 0x77,
	0x46, 0xd5, 0xda, 0x90, 0xb1, 0xa1, 0x8f, 0x8d, 0xf8, 0x55, 0x6f, 0x3c, 0x68, 0xe0, 0x28, 0xe0,
	0xa7, 0x42, 0xa4, 0xde, 0x74, 0x70, 0x10, 0x35, 0x26, 0xb6, 0xef, 0x39, 0x36, 0x47, 0x11, 0xd4,
	0x1e, 0x42, 0xe1, 0xc0, 0xee, 0xa1, 0x4f, 0x08, 0xe4, 0xa9, 0x3d, 0xc2, 0xaa, 0x54, 0x97, 0xf6,
	0x76, 0xcc, 0xf8, 0x4c, 0xca, 0x50, 0x98, 0xd8, 0xfe, 0x18, 0xab, 0xb9, 0x38, 0x28, 0x2e, 0xda,
	0x7b, 0x50, 0x62, 0xc9, 0x21, 0x72, 0xf2, 0x02, 0xb6, 0xfd, 0xf9, 0x39, 0xaa, 0x4a, 0x75, 0x79,
	0xaf, 0xd8, 0xbc, 0xa3, 0x2f, 0xef, 0x4c, 0x8f, 0x15, 0x86, 0x72, 0x66, 0x14, 0x7e, 0x48, 0x39,
	0x45, 0x32, 0x13, 0x99, 0xf6, 0x14, 0x6a, 0x26, 0x0e, 0xbd, 0x88, 0x63, 0xf8, 0x66, 0x8c, 0x11,
	0x7f, 0xd9, 0xef, 0xb3, 0x31, 0xe5, 0x26, 0x9e, 0xcc, 0x6f, 0xa4, 0x06, 0x3b, 0x62, 0xf0, 0x63,
	0xcf, 0x49, 0x5a, 0x53, 0x44, 0xe0, 0xad, 0xa3, 0xb5, 0xe1, 0xf6, 0x72, 0x6d, 0x14, 0x30, 0x1a,
	0xc5, 0xed, 0x73, 0xf6, 0x05, 0x69, 0x22, 0x14, 0x17, 0xed, 0x8f, 0x04, 0xb7, 0x52, 0x59, 0x87,
	0x45, 0x7c, 0x3e, 0x69, 0x5a, 0xae, 0x0d, 0xca, 0x10, 0x29, 0x86, 0x36, 0x17, 0x20, 0x8a, 0xcd,
	0x8a, 0x2e, 0x90, 0xea, 0x29, 0x52, 0xfd, 0xf5, 0x1c, 0x69, 0x67, 0xcb, 0xcc, 0x5e, 0x92, 0x8f,
	0x50, 0xc0, 0xaf, 0x76, 0x9f, 0x0b, 0x4c, 0xc6, 0xf3, 0x33, 0xe3, 0x59, 0xf8, 0xa4, 0x24, 0x37,
	0x8b, 0xdd, 0xe9, 0xbd, 0x4f, 0xdd, 0x69, 0xd7, 0xd9, 0xff, 0x7c, 0xd7, 0x2a, 0xd8, 0xce, 0xc8,
	0xa3, 0x96, 0x6c, 0x07, 0x9e, 0x95, 0xef, 0xf9, 0x6c, 0x68, 0xc9, 0xee, 0x8c, 0x5a, 0x57, 0x5c,
	0x16, 0x7a, 0x33, 0x46, 0x2d, 0x25, 0xc5, 0xd5, 0xd9, 0x32, 0x45, 0x3a, 0x62, 0x64, 0x70, 0xe5,
	0xb8, 0x97, 0xfa, 0x5a, 0xb8, 0x87, 0xc8, 0x63, 0xbe, 0xdf, 0xa5, 0x5c, 0x29, 0xe3, 0x6b, 0xdc,
	0x00, 0xc5, 0x4d, 0x86, 0x24, 0xf2, 0x6f, 0x43, 0xd2, 0xde, 0x41, 0xf5, 0xdf, 0xe9, 0x13, 0x60,
	0xea, 0xf9, 0xe3, 0x14, 0x76, 0x26, 0x26, 0x90, 0x1f, 0x9c, 0x38, 0x34, 0xb1, 0x42, 0x7c, 0xd6,
	0x2a, 0x50, 0x3e, 0xf0, 0x22, 0x9e, 0xe6, 0x89, 0x12, 0x8c, 0xda, 0x2f, 0x09, 0x76, 0x17, 0x7e,
	0x48, 0x2a, 0x7c, 0x80, 0x9d, 0x34, 0x63, 0x6a, 0x99, 0x47, 0x2b, 0xa7, 0x5a, 0x96, 0x41, 0xcf,
	0x9a, 0x3e, 0x4f, 0xa4, 0x72, 0x50, 0xd2, 0xf0, 0xa6, 0x33, 0x90, 0xc7, 0x9b, 0x42, 0xce, 0xac,
	0xdb, 0x82, 0xdd, 0x57, 0xe8, 0x23, 0xc7, 0x45, 0x17, 0xad, 0x69, 0xa1, 0xf9, 0x53, 0x86, 0xe2,
	0x51, 0xba, 0xe9, 0x19, 0x25, 0xdf, 0x24, 0x28, 0x2f, 0x33, 0x31, 0x69, 0xad, 0xea, 0x63, 0xcd,
	0xdf, 0x45, 0x6d, 0x6f, 0x26, 0x4a, 0x96, 0x32, 0x86, 0xd2, 0xa2, 0x25, 0x48, 0xe3, 0x7f, 0x99,
	0x16, 0x86, 0x56, 0x1f, 0x5c, 0x5e, 0x90, 0x94, 0xf5, 0xe1, 0xda, 0x85, 0x15, 0x93, 0xfb, 0x97,
	0x74, 0x82, 0x28, 0xb8, 0xbf, 0x91, 0x6f, 0xc8, 0x11, 0x5c, 0xbf, 0xb8, 0x2d, 0xb2, 0x32, 0xc1,
	0xd2, 0xad, 0xaa, 0x2b, 0xbe, 0x04, 0x46, 0xde, 0xca, 0x05, 0xbd, 0xde, 0x76, 0x1c, 0x6d, 0xfd,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x09, 0xd0, 0x9c, 0xa8, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WaypointHznClient is the client API for WaypointHzn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WaypointHznClient interface {
	RegisterGuestAccount(ctx context.Context, in *RegisterGuestAccountRequest, opts ...grpc.CallOption) (*RegisterGuestAccountResponse, error)
	RegisterHostname(ctx context.Context, in *RegisterHostnameRequest, opts ...grpc.CallOption) (*RegisterHostnameResponse, error)
	ListHostnames(ctx context.Context, in *ListHostnamesRequest, opts ...grpc.CallOption) (*ListHostnamesResponse, error)
	DeleteHostname(ctx context.Context, in *DeleteHostnameRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type waypointHznClient struct {
	cc grpc.ClientConnInterface
}

func NewWaypointHznClient(cc grpc.ClientConnInterface) WaypointHznClient {
	return &waypointHznClient{cc}
}

func (c *waypointHznClient) RegisterGuestAccount(ctx context.Context, in *RegisterGuestAccountRequest, opts ...grpc.CallOption) (*RegisterGuestAccountResponse, error) {
	out := new(RegisterGuestAccountResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint_hzn.WaypointHzn/RegisterGuestAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waypointHznClient) RegisterHostname(ctx context.Context, in *RegisterHostnameRequest, opts ...grpc.CallOption) (*RegisterHostnameResponse, error) {
	out := new(RegisterHostnameResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint_hzn.WaypointHzn/RegisterHostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waypointHznClient) ListHostnames(ctx context.Context, in *ListHostnamesRequest, opts ...grpc.CallOption) (*ListHostnamesResponse, error) {
	out := new(ListHostnamesResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint_hzn.WaypointHzn/ListHostnames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waypointHznClient) DeleteHostname(ctx context.Context, in *DeleteHostnameRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint_hzn.WaypointHzn/DeleteHostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaypointHznServer is the server API for WaypointHzn service.
type WaypointHznServer interface {
	RegisterGuestAccount(context.Context, *RegisterGuestAccountRequest) (*RegisterGuestAccountResponse, error)
	RegisterHostname(context.Context, *RegisterHostnameRequest) (*RegisterHostnameResponse, error)
	ListHostnames(context.Context, *ListHostnamesRequest) (*ListHostnamesResponse, error)
	DeleteHostname(context.Context, *DeleteHostnameRequest) (*empty.Empty, error)
}

// UnimplementedWaypointHznServer can be embedded to have forward compatible implementations.
type UnimplementedWaypointHznServer struct {
}

func (*UnimplementedWaypointHznServer) RegisterGuestAccount(ctx context.Context, req *RegisterGuestAccountRequest) (*RegisterGuestAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterGuestAccount not implemented")
}
func (*UnimplementedWaypointHznServer) RegisterHostname(ctx context.Context, req *RegisterHostnameRequest) (*RegisterHostnameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterHostname not implemented")
}
func (*UnimplementedWaypointHznServer) ListHostnames(ctx context.Context, req *ListHostnamesRequest) (*ListHostnamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHostnames not implemented")
}
func (*UnimplementedWaypointHznServer) DeleteHostname(ctx context.Context, req *DeleteHostnameRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHostname not implemented")
}

func RegisterWaypointHznServer(s *grpc.Server, srv WaypointHznServer) {
	s.RegisterService(&_WaypointHzn_serviceDesc, srv)
}

func _WaypointHzn_RegisterGuestAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterGuestAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaypointHznServer).RegisterGuestAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint_hzn.WaypointHzn/RegisterGuestAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaypointHznServer).RegisterGuestAccount(ctx, req.(*RegisterGuestAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaypointHzn_RegisterHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterHostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaypointHznServer).RegisterHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint_hzn.WaypointHzn/RegisterHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaypointHznServer).RegisterHostname(ctx, req.(*RegisterHostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaypointHzn_ListHostnames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHostnamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaypointHznServer).ListHostnames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint_hzn.WaypointHzn/ListHostnames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaypointHznServer).ListHostnames(ctx, req.(*ListHostnamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaypointHzn_DeleteHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaypointHznServer).DeleteHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint_hzn.WaypointHzn/DeleteHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaypointHznServer).DeleteHostname(ctx, req.(*DeleteHostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WaypointHzn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.waypoint_hzn.WaypointHzn",
	HandlerType: (*WaypointHznServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterGuestAccount",
			Handler:    _WaypointHzn_RegisterGuestAccount_Handler,
		},
		{
			MethodName: "RegisterHostname",
			Handler:    _WaypointHzn_RegisterHostname_Handler,
		},
		{
			MethodName: "ListHostnames",
			Handler:    _WaypointHzn_ListHostnames_Handler,
		},
		{
			MethodName: "DeleteHostname",
			Handler:    _WaypointHzn_DeleteHostname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
