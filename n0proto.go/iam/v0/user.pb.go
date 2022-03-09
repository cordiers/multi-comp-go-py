// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n0stack/iam/v0/user.proto

package piam

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type User_UserState int32

const (
	User_USER_UNSPECIFIED User_UserState = 0
	// working API
	User_PENDING   User_UserState = 1
	User_AVAILABLE User_UserState = 2
)

var User_UserState_name = map[int32]string{
	0: "USER_UNSPECIFIED",
	1: "PENDING",
	2: "AVAILABLE",
}

var User_UserState_value = map[string]int32{
	"USER_UNSPECIFIED": 0,
	"PENDING":          1,
	"AVAILABLE":        2,
}

func (x User_UserState) String() string {
	return proto.EnumName(User_UserState_name, int32(x))
}

func (User_UserState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{0, 0}
}

type User struct {
	// Name is a unique field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Annotations can store metadata used by the system for control.
	// In particular, implementation-dependent fields that can not be set as protobuf fields are targeted.
	// The control specified by n0stack may delete metadata specified by the user.
	Annotations map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Labels stores user-defined metadata.
	// The n0stack system must not rewrite this value.
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SshPublicKeys        map[string]string `protobuf:"bytes,10,rep,name=ssh_public_keys,json=sshPublicKeys,proto3" json:"ssh_public_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	State                User_UserState    `protobuf:"varint,50,opt,name=state,proto3,enum=n0stack.iam.v0.User_UserState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *User) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *User) GetSshPublicKeys() map[string]string {
	if m != nil {
		return m.SshPublicKeys
	}
	return nil
}

func (m *User) GetState() User_UserState {
	if m != nil {
		return m.State
	}
	return User_USER_UNSPECIFIED
}

type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{1}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{2}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{3}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateUserRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *CreateUserRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type DeleteUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{5}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddSshPublicKeyRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	SshPublicKeyName     string   `protobuf:"bytes,2,opt,name=ssh_public_key_name,json=sshPublicKeyName,proto3" json:"ssh_public_key_name,omitempty"`
	SshPublicKey         string   `protobuf:"bytes,3,opt,name=ssh_public_key,json=sshPublicKey,proto3" json:"ssh_public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSshPublicKeyRequest) Reset()         { *m = AddSshPublicKeyRequest{} }
func (m *AddSshPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*AddSshPublicKeyRequest) ProtoMessage()    {}
func (*AddSshPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{6}
}

func (m *AddSshPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSshPublicKeyRequest.Unmarshal(m, b)
}
func (m *AddSshPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSshPublicKeyRequest.Marshal(b, m, deterministic)
}
func (m *AddSshPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSshPublicKeyRequest.Merge(m, src)
}
func (m *AddSshPublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_AddSshPublicKeyRequest.Size(m)
}
func (m *AddSshPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSshPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSshPublicKeyRequest proto.InternalMessageInfo

func (m *AddSshPublicKeyRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AddSshPublicKeyRequest) GetSshPublicKeyName() string {
	if m != nil {
		return m.SshPublicKeyName
	}
	return ""
}

func (m *AddSshPublicKeyRequest) GetSshPublicKey() string {
	if m != nil {
		return m.SshPublicKey
	}
	return ""
}

type DeleteSshPublicKeyRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	SshPublicKeyName     string   `protobuf:"bytes,2,opt,name=ssh_public_key_name,json=sshPublicKeyName,proto3" json:"ssh_public_key_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSshPublicKeyRequest) Reset()         { *m = DeleteSshPublicKeyRequest{} }
func (m *DeleteSshPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSshPublicKeyRequest) ProtoMessage()    {}
func (*DeleteSshPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55be2fe86ad6ef99, []int{7}
}

func (m *DeleteSshPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSshPublicKeyRequest.Unmarshal(m, b)
}
func (m *DeleteSshPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSshPublicKeyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSshPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSshPublicKeyRequest.Merge(m, src)
}
func (m *DeleteSshPublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSshPublicKeyRequest.Size(m)
}
func (m *DeleteSshPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSshPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSshPublicKeyRequest proto.InternalMessageInfo

func (m *DeleteSshPublicKeyRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *DeleteSshPublicKeyRequest) GetSshPublicKeyName() string {
	if m != nil {
		return m.SshPublicKeyName
	}
	return ""
}

func init() {
	proto.RegisterEnum("n0stack.iam.v0.User_UserState", User_UserState_name, User_UserState_value)
	proto.RegisterType((*User)(nil), "n0stack.iam.v0.User")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v0.User.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v0.User.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v0.User.SshPublicKeysEntry")
	proto.RegisterType((*ListUsersRequest)(nil), "n0stack.iam.v0.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "n0stack.iam.v0.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "n0stack.iam.v0.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "n0stack.iam.v0.CreateUserRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v0.CreateUserRequest.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.iam.v0.CreateUserRequest.LabelsEntry")
	proto.RegisterType((*DeleteUserRequest)(nil), "n0stack.iam.v0.DeleteUserRequest")
	proto.RegisterType((*AddSshPublicKeyRequest)(nil), "n0stack.iam.v0.AddSshPublicKeyRequest")
	proto.RegisterType((*DeleteSshPublicKeyRequest)(nil), "n0stack.iam.v0.DeleteSshPublicKeyRequest")
}

func init() { proto.RegisterFile("n0stack/iam/v0/user.proto", fileDescriptor_55be2fe86ad6ef99) }

var fileDescriptor_55be2fe86ad6ef99 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4f, 0x6f, 0xd3, 0x48,
	0x14, 0x5f, 0xe7, 0x4f, 0xbb, 0x79, 0x69, 0x53, 0x77, 0x9a, 0xad, 0xdc, 0x74, 0x55, 0xa5, 0x56,
	0x77, 0xdb, 0x8d, 0x36, 0x9e, 0x10, 0x38, 0x94, 0xa0, 0x52, 0xd2, 0xd6, 0x94, 0x88, 0x28, 0x54,
	0x09, 0xe5, 0x80, 0x90, 0x22, 0x27, 0x9d, 0x3a, 0xa6, 0x89, 0x6d, 0x32, 0x4e, 0x50, 0x54, 0xf5,
	0xc2, 0x91, 0x23, 0x5c, 0x40, 0xe2, 0xc0, 0x37, 0xe1, 0x43, 0x70, 0xe4, 0xca, 0x07, 0x41, 0x1e,
	0x3b, 0xa9, 0x13, 0xbb, 0x84, 0x0a, 0x21, 0x2e, 0x89, 0xfd, 0xe6, 0xf7, 0x7e, 0xef, 0xcf, 0xfc,
	0xde, 0x8c, 0x61, 0x45, 0xcf, 0x51, 0x4b, 0x69, 0x9e, 0x61, 0x4d, 0xe9, 0xe0, 0x7e, 0x0e, 0xf7,
	0x28, 0xe9, 0x4a, 0x66, 0xd7, 0xb0, 0x0c, 0x94, 0x70, 0x97, 0x24, 0x4d, 0xe9, 0x48, 0xfd, 0x5c,
	0xea, 0x6f, 0xd5, 0x30, 0xd4, 0x36, 0xc1, 0x8a, 0xa9, 0x61, 0x45, 0xd7, 0x0d, 0x4b, 0xb1, 0x34,
	0x43, 0xa7, 0x0e, 0x3a, 0xb5, 0xea, 0xae, 0xb2, 0xb7, 0x46, 0xef, 0x14, 0x93, 0x8e, 0x69, 0x0d,
	0xdc, 0xc5, 0xff, 0xd9, 0x5f, 0x33, 0xab, 0x12, 0x3d, 0x4b, 0x5f, 0x2a, 0xaa, 0x4a, 0xba, 0xd8,
	0x30, 0x99, 0xbb, 0x9f, 0x4a, 0xfc, 0x18, 0x81, 0xc8, 0x31, 0x25, 0x5d, 0x84, 0x20, 0xa2, 0x2b,
	0x1d, 0x22, 0x70, 0x69, 0x6e, 0x2b, 0x56, 0x65, 0xcf, 0xe8, 0x10, 0xe2, 0x1e, 0x0f, 0x21, 0x9c,
	0x0e, 0x6f, 0xc5, 0xf3, 0xff, 0x48, 0xe3, 0xb9, 0x4a, 0xb6, 0xbb, 0x54, 0xbc, 0xc4, 0xc9, 0xba,
	0xd5, 0x1d, 0x54, 0xbd, 0x9e, 0x68, 0x1b, 0x66, 0xda, 0x4a, 0x83, 0xb4, 0xa9, 0x10, 0x61, 0x1c,
	0xe9, 0x40, 0x8e, 0x32, 0x83, 0x38, 0xee, 0x2e, 0x1e, 0x3d, 0x82, 0x05, 0x4a, 0x5b, 0x75, 0xb3,
	0xd7, 0x68, 0x6b, 0xcd, 0xfa, 0x19, 0x19, 0x50, 0x01, 0x18, 0xc5, 0x66, 0x20, 0x45, 0x8d, 0xb6,
	0x8e, 0x18, 0xf4, 0x21, 0x19, 0xb8, 0x4c, 0xf3, 0xd4, 0x6b, 0x43, 0xb7, 0x20, 0x4a, 0x2d, 0xc5,
	0x22, 0x42, 0x3e, 0xcd, 0x6d, 0x25, 0xf2, 0x6b, 0x81, 0x34, 0xf6, 0x4f, 0xcd, 0x46, 0x55, 0x1d,
	0x70, 0xea, 0x2e, 0xf0, 0x93, 0x15, 0x22, 0x1e, 0xc2, 0x67, 0x64, 0xe0, 0x36, 0xcc, 0x7e, 0x44,
	0x49, 0x88, 0xf6, 0x95, 0x76, 0x8f, 0x08, 0x21, 0x66, 0x73, 0x5e, 0x0a, 0xa1, 0x6d, 0x2e, 0x75,
	0x1b, 0xe2, 0x9e, 0xea, 0xae, 0xe5, 0x7a, 0x0f, 0x90, 0xbf, 0xaa, 0xeb, 0x30, 0x88, 0x3b, 0x10,
	0x1b, 0x15, 0x84, 0x92, 0xc0, 0x1f, 0xd7, 0xe4, 0x6a, 0xfd, 0xb8, 0x52, 0x3b, 0x92, 0xf7, 0x4b,
	0xf7, 0x4b, 0xf2, 0x01, 0xff, 0x07, 0x8a, 0xc3, 0xec, 0x91, 0x5c, 0x39, 0x28, 0x55, 0x0e, 0x79,
	0x0e, 0xcd, 0x43, 0xac, 0xf8, 0xa4, 0x58, 0x2a, 0x17, 0xf7, 0xca, 0x32, 0x1f, 0x12, 0x11, 0xf0,
	0x65, 0x8d, 0x5a, 0x36, 0x05, 0xad, 0x92, 0x17, 0x3d, 0x42, 0x2d, 0x71, 0x17, 0x16, 0x3d, 0x36,
	0x6a, 0x1a, 0x3a, 0x25, 0x28, 0x03, 0x51, 0x5b, 0xd2, 0x54, 0xe0, 0xd8, 0x0e, 0x25, 0x83, 0x5a,
	0x5b, 0x75, 0x20, 0xe2, 0x06, 0x24, 0x0e, 0x09, 0xf3, 0x77, 0x29, 0x83, 0x04, 0x28, 0x7e, 0x0a,
	0xc1, 0xe2, 0x7e, 0x97, 0x28, 0x16, 0x99, 0x82, 0x44, 0x8f, 0x83, 0xa4, 0x9a, 0x9f, 0xcc, 0xc0,
	0xc7, 0x35, 0x45, 0xb7, 0xf2, 0x84, 0x6e, 0xb3, 0xd3, 0x09, 0x03, 0x44, 0xfc, 0x1b, 0xd5, 0x23,
	0x6e, 0xc2, 0xe2, 0x01, 0x69, 0x93, 0xa9, 0x0d, 0x14, 0x5f, 0x73, 0xb0, 0x5c, 0x3c, 0x39, 0xf1,
	0x4a, 0x6d, 0x08, 0x5f, 0x85, 0x98, 0xbd, 0x69, 0x75, 0x8f, 0xcf, 0x9f, 0xb6, 0xa1, 0x62, 0x37,
	0x3e, 0x0b, 0x4b, 0xe3, 0x03, 0xea, 0xc0, 0x9c, 0x44, 0x78, 0xef, 0xec, 0x31, 0xf8, 0x06, 0x24,
	0xc6, 0xe1, 0x42, 0x98, 0x21, 0xe7, 0xbc, 0x48, 0x51, 0x85, 0x15, 0x27, 0xeb, 0x5f, 0x9c, 0x4e,
	0xfe, 0x4b, 0x14, 0xe2, 0x6c, 0x36, 0x48, 0xb7, 0xaf, 0x35, 0x09, 0x3a, 0x85, 0xd8, 0x48, 0xd7,
	0xc8, 0x77, 0x4a, 0x4d, 0x8e, 0x41, 0x6a, 0xfd, 0x3b, 0x08, 0x67, 0x28, 0xc4, 0xe4, 0xab, 0xcf,
	0x5f, 0xdf, 0x86, 0x12, 0x68, 0x8e, 0x9d, 0xe5, 0xee, 0xa9, 0x8f, 0x9e, 0xc1, 0xac, 0x2b, 0x7f,
	0xe4, 0x3b, 0x81, 0xc6, 0xe7, 0x22, 0x15, 0x38, 0x46, 0xe2, 0x2a, 0xa3, 0xfd, 0x0b, 0x2d, 0x79,
	0x69, 0xf1, 0xb9, 0x5d, 0xec, 0x05, 0x22, 0x00, 0x97, 0xc2, 0x44, 0xeb, 0x53, 0x45, 0x7b, 0x45,
	0x8c, 0x35, 0x16, 0x43, 0x10, 0x83, 0x62, 0x14, 0xb8, 0x0c, 0x3a, 0x01, 0xb8, 0xd4, 0x96, 0x3f,
	0x8c, 0x4f, 0x77, 0xa9, 0x65, 0xc9, 0xb9, 0xb8, 0xa4, 0xe1, 0xc5, 0x25, 0xc9, 0xf6, 0xc5, 0x35,
	0x2c, 0x26, 0x13, 0x58, 0xcc, 0x7b, 0x0e, 0x16, 0x26, 0x84, 0x89, 0xfe, 0x9d, 0x8c, 0x15, 0xac,
	0xdc, 0x2b, 0xea, 0x7a, 0xc0, 0xc2, 0xed, 0x89, 0x3b, 0xe3, 0xe1, 0x46, 0xa2, 0xba, 0xc0, 0xe3,
	0x12, 0xc2, 0xe7, 0x01, 0x92, 0x62, 0x1d, 0xf8, 0xc0, 0x01, 0xf2, 0x0b, 0x15, 0xfd, 0x17, 0xdc,
	0x8a, 0x1f, 0xcf, 0x50, 0x66, 0x19, 0xee, 0x66, 0x7e, 0x2e, 0xc3, 0xbd, 0x77, 0xdc, 0x9b, 0x62,
	0x03, 0x6d, 0xc3, 0xac, 0x1b, 0x43, 0xcc, 0x8e, 0x1e, 0x91, 0xd8, 0xb2, 0x2c, 0x93, 0x16, 0x30,
	0x56, 0x35, 0xab, 0xd5, 0x6b, 0x48, 0x4d, 0xa3, 0x83, 0x87, 0xdf, 0x27, 0xee, 0x7f, 0x26, 0xc4,
	0x85, 0xf2, 0xbc, 0x62, 0x9a, 0x6d, 0xad, 0xc9, 0x8e, 0x2e, 0xfc, 0x9c, 0x1a, 0x7a, 0xc1, 0x67,
	0x79, 0x7a, 0xe3, 0x6a, 0x0e, 0xac, 0xe7, 0xd8, 0x66, 0x4b, 0xaa, 0xe1, 0x7e, 0xf6, 0xdc, 0x31,
	0x35, 0xa5, 0xd3, 0x98, 0x61, 0xd6, 0x9b, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xbb, 0x11,
	0xeb, 0x14, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddSshPublicKey(ctx context.Context, in *AddSshPublicKeyRequest, opts ...grpc.CallOption) (*User, error)
	DeleteSshPublicKey(ctx context.Context, in *DeleteSshPublicKeyRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v0.UserService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v0.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v0.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v0.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddSshPublicKey(ctx context.Context, in *AddSshPublicKeyRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v0.UserService/AddSshPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteSshPublicKey(ctx context.Context, in *DeleteSshPublicKeyRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/n0stack.iam.v0.UserService/DeleteSshPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	AddSshPublicKey(context.Context, *AddSshPublicKeyRequest) (*User, error)
	DeleteSshPublicKey(context.Context, *DeleteSshPublicKeyRequest) (*User, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedUserServiceServer) AddSshPublicKey(ctx context.Context, req *AddSshPublicKeyRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSshPublicKey not implemented")
}
func (*UnimplementedUserServiceServer) DeleteSshPublicKey(ctx context.Context, req *DeleteSshPublicKeyRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSshPublicKey not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v0.UserService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v0.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v0.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v0.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddSshPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSshPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddSshPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v0.UserService/AddSshPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddSshPublicKey(ctx, req.(*AddSshPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteSshPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSshPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteSshPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.iam.v0.UserService/DeleteSshPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteSshPublicKey(ctx, req.(*DeleteSshPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.iam.v0.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "AddSshPublicKey",
			Handler:    _UserService_AddSshPublicKey_Handler,
		},
		{
			MethodName: "DeleteSshPublicKey",
			Handler:    _UserService_DeleteSshPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "n0stack/iam/v0/user.proto",
}