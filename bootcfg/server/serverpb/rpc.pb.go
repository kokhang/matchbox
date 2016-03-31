// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	SelectGroupRequest
	SelectGroupResponse
	SelectProfileRequest
	SelectProfileResponse
	GroupGetRequest
	GroupListRequest
	GroupGetResponse
	GroupListResponse
	ProfilePutRequest
	ProfilePutResponse
	ProfileGetRequest
	ProfileGetResponse
	ProfileListRequest
	ProfileListResponse
*/
package serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import storagepb "github.com/coreos/coreos-baremetal/bootcfg/storage/storagepb"

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
const _ = proto.ProtoPackageIsVersion1

type SelectGroupRequest struct {
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SelectGroupRequest) Reset()                    { *m = SelectGroupRequest{} }
func (m *SelectGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*SelectGroupRequest) ProtoMessage()               {}
func (*SelectGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SelectGroupRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SelectGroupResponse struct {
	Group *storagepb.Group `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *SelectGroupResponse) Reset()                    { *m = SelectGroupResponse{} }
func (m *SelectGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*SelectGroupResponse) ProtoMessage()               {}
func (*SelectGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SelectGroupResponse) GetGroup() *storagepb.Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type SelectProfileRequest struct {
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SelectProfileRequest) Reset()                    { *m = SelectProfileRequest{} }
func (m *SelectProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*SelectProfileRequest) ProtoMessage()               {}
func (*SelectProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SelectProfileRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SelectProfileResponse struct {
	Profile *storagepb.Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
}

func (m *SelectProfileResponse) Reset()                    { *m = SelectProfileResponse{} }
func (m *SelectProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*SelectProfileResponse) ProtoMessage()               {}
func (*SelectProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SelectProfileResponse) GetProfile() *storagepb.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type GroupGetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GroupGetRequest) Reset()                    { *m = GroupGetRequest{} }
func (m *GroupGetRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupGetRequest) ProtoMessage()               {}
func (*GroupGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GroupListRequest struct {
}

func (m *GroupListRequest) Reset()                    { *m = GroupListRequest{} }
func (m *GroupListRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupListRequest) ProtoMessage()               {}
func (*GroupListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GroupGetResponse struct {
	Group *storagepb.Group `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *GroupGetResponse) Reset()                    { *m = GroupGetResponse{} }
func (m *GroupGetResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupGetResponse) ProtoMessage()               {}
func (*GroupGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GroupGetResponse) GetGroup() *storagepb.Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type GroupListResponse struct {
	Groups []*storagepb.Group `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
}

func (m *GroupListResponse) Reset()                    { *m = GroupListResponse{} }
func (m *GroupListResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupListResponse) ProtoMessage()               {}
func (*GroupListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GroupListResponse) GetGroups() []*storagepb.Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type ProfilePutRequest struct {
	Profile *storagepb.Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
}

func (m *ProfilePutRequest) Reset()                    { *m = ProfilePutRequest{} }
func (m *ProfilePutRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfilePutRequest) ProtoMessage()               {}
func (*ProfilePutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ProfilePutRequest) GetProfile() *storagepb.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type ProfilePutResponse struct {
}

func (m *ProfilePutResponse) Reset()                    { *m = ProfilePutResponse{} }
func (m *ProfilePutResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfilePutResponse) ProtoMessage()               {}
func (*ProfilePutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ProfileGetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ProfileGetRequest) Reset()                    { *m = ProfileGetRequest{} }
func (m *ProfileGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileGetRequest) ProtoMessage()               {}
func (*ProfileGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type ProfileGetResponse struct {
	Profile *storagepb.Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
}

func (m *ProfileGetResponse) Reset()                    { *m = ProfileGetResponse{} }
func (m *ProfileGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfileGetResponse) ProtoMessage()               {}
func (*ProfileGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ProfileGetResponse) GetProfile() *storagepb.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type ProfileListRequest struct {
}

func (m *ProfileListRequest) Reset()                    { *m = ProfileListRequest{} }
func (m *ProfileListRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileListRequest) ProtoMessage()               {}
func (*ProfileListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type ProfileListResponse struct {
	Profiles []*storagepb.Profile `protobuf:"bytes,1,rep,name=profiles" json:"profiles,omitempty"`
}

func (m *ProfileListResponse) Reset()                    { *m = ProfileListResponse{} }
func (m *ProfileListResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfileListResponse) ProtoMessage()               {}
func (*ProfileListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ProfileListResponse) GetProfiles() []*storagepb.Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func init() {
	proto.RegisterType((*SelectGroupRequest)(nil), "serverpb.SelectGroupRequest")
	proto.RegisterType((*SelectGroupResponse)(nil), "serverpb.SelectGroupResponse")
	proto.RegisterType((*SelectProfileRequest)(nil), "serverpb.SelectProfileRequest")
	proto.RegisterType((*SelectProfileResponse)(nil), "serverpb.SelectProfileResponse")
	proto.RegisterType((*GroupGetRequest)(nil), "serverpb.GroupGetRequest")
	proto.RegisterType((*GroupListRequest)(nil), "serverpb.GroupListRequest")
	proto.RegisterType((*GroupGetResponse)(nil), "serverpb.GroupGetResponse")
	proto.RegisterType((*GroupListResponse)(nil), "serverpb.GroupListResponse")
	proto.RegisterType((*ProfilePutRequest)(nil), "serverpb.ProfilePutRequest")
	proto.RegisterType((*ProfilePutResponse)(nil), "serverpb.ProfilePutResponse")
	proto.RegisterType((*ProfileGetRequest)(nil), "serverpb.ProfileGetRequest")
	proto.RegisterType((*ProfileGetResponse)(nil), "serverpb.ProfileGetResponse")
	proto.RegisterType((*ProfileListRequest)(nil), "serverpb.ProfileListRequest")
	proto.RegisterType((*ProfileListResponse)(nil), "serverpb.ProfileListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Groups service

type GroupsClient interface {
	// Get a machine Group by id.
	GroupGet(ctx context.Context, in *GroupGetRequest, opts ...grpc.CallOption) (*GroupGetResponse, error)
	// List all machine Groups.
	GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error)
}

type groupsClient struct {
	cc *grpc.ClientConn
}

func NewGroupsClient(cc *grpc.ClientConn) GroupsClient {
	return &groupsClient{cc}
}

func (c *groupsClient) GroupGet(ctx context.Context, in *GroupGetRequest, opts ...grpc.CallOption) (*GroupGetResponse, error) {
	out := new(GroupGetResponse)
	err := grpc.Invoke(ctx, "/serverpb.Groups/GroupGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error) {
	out := new(GroupListResponse)
	err := grpc.Invoke(ctx, "/serverpb.Groups/GroupList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Groups service

type GroupsServer interface {
	// Get a machine Group by id.
	GroupGet(context.Context, *GroupGetRequest) (*GroupGetResponse, error)
	// List all machine Groups.
	GroupList(context.Context, *GroupListRequest) (*GroupListResponse, error)
}

func RegisterGroupsServer(s *grpc.Server, srv GroupsServer) {
	s.RegisterService(&_Groups_serviceDesc, srv)
}

func _Groups_GroupGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GroupGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupsServer).GroupGet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Groups_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupsServer).GroupList(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Groups_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Groups",
	HandlerType: (*GroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GroupGet",
			Handler:    _Groups_GroupGet_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _Groups_GroupList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Profiles service

type ProfilesClient interface {
	// Create a Profile.
	ProfilePut(ctx context.Context, in *ProfilePutRequest, opts ...grpc.CallOption) (*ProfilePutResponse, error)
	// Get a Profile by id.
	ProfileGet(ctx context.Context, in *ProfileGetRequest, opts ...grpc.CallOption) (*ProfileGetResponse, error)
	// List all Profiles.
	ProfileList(ctx context.Context, in *ProfileListRequest, opts ...grpc.CallOption) (*ProfileListResponse, error)
}

type profilesClient struct {
	cc *grpc.ClientConn
}

func NewProfilesClient(cc *grpc.ClientConn) ProfilesClient {
	return &profilesClient{cc}
}

func (c *profilesClient) ProfilePut(ctx context.Context, in *ProfilePutRequest, opts ...grpc.CallOption) (*ProfilePutResponse, error) {
	out := new(ProfilePutResponse)
	err := grpc.Invoke(ctx, "/serverpb.Profiles/ProfilePut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) ProfileGet(ctx context.Context, in *ProfileGetRequest, opts ...grpc.CallOption) (*ProfileGetResponse, error) {
	out := new(ProfileGetResponse)
	err := grpc.Invoke(ctx, "/serverpb.Profiles/ProfileGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) ProfileList(ctx context.Context, in *ProfileListRequest, opts ...grpc.CallOption) (*ProfileListResponse, error) {
	out := new(ProfileListResponse)
	err := grpc.Invoke(ctx, "/serverpb.Profiles/ProfileList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profiles service

type ProfilesServer interface {
	// Create a Profile.
	ProfilePut(context.Context, *ProfilePutRequest) (*ProfilePutResponse, error)
	// Get a Profile by id.
	ProfileGet(context.Context, *ProfileGetRequest) (*ProfileGetResponse, error)
	// List all Profiles.
	ProfileList(context.Context, *ProfileListRequest) (*ProfileListResponse, error)
}

func RegisterProfilesServer(s *grpc.Server, srv ProfilesServer) {
	s.RegisterService(&_Profiles_serviceDesc, srv)
}

func _Profiles_ProfilePut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProfilePutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProfilesServer).ProfilePut(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Profiles_ProfileGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProfileGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProfilesServer).ProfileGet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Profiles_ProfileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProfileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProfilesServer).ProfileList(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Profiles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Profiles",
	HandlerType: (*ProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfilePut",
			Handler:    _Profiles_ProfilePut_Handler,
		},
		{
			MethodName: "ProfileGet",
			Handler:    _Profiles_ProfileGet_Handler,
		},
		{
			MethodName: "ProfileList",
			Handler:    _Profiles_ProfileList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Select service

type SelectClient interface {
	// SelectGroup returns the Group matching the given labels.
	SelectGroup(ctx context.Context, in *SelectGroupRequest, opts ...grpc.CallOption) (*SelectGroupResponse, error)
	// SelectProfile returns the Profile matching the given labels.
	SelectProfile(ctx context.Context, in *SelectProfileRequest, opts ...grpc.CallOption) (*SelectProfileResponse, error)
}

type selectClient struct {
	cc *grpc.ClientConn
}

func NewSelectClient(cc *grpc.ClientConn) SelectClient {
	return &selectClient{cc}
}

func (c *selectClient) SelectGroup(ctx context.Context, in *SelectGroupRequest, opts ...grpc.CallOption) (*SelectGroupResponse, error) {
	out := new(SelectGroupResponse)
	err := grpc.Invoke(ctx, "/serverpb.Select/SelectGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *selectClient) SelectProfile(ctx context.Context, in *SelectProfileRequest, opts ...grpc.CallOption) (*SelectProfileResponse, error) {
	out := new(SelectProfileResponse)
	err := grpc.Invoke(ctx, "/serverpb.Select/SelectProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Select service

type SelectServer interface {
	// SelectGroup returns the Group matching the given labels.
	SelectGroup(context.Context, *SelectGroupRequest) (*SelectGroupResponse, error)
	// SelectProfile returns the Profile matching the given labels.
	SelectProfile(context.Context, *SelectProfileRequest) (*SelectProfileResponse, error)
}

func RegisterSelectServer(s *grpc.Server, srv SelectServer) {
	s.RegisterService(&_Select_serviceDesc, srv)
}

func _Select_SelectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SelectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SelectServer).SelectGroup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Select_SelectProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SelectProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SelectServer).SelectProfile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Select_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Select",
	HandlerType: (*SelectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelectGroup",
			Handler:    _Select_SelectGroup_Handler,
		},
		{
			MethodName: "SelectProfile",
			Handler:    _Select_SelectProfile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xa9, 0x1a, 0x92, 0x89, 0x80, 0x74, 0x1a, 0xa4, 0x62, 0xca, 0x9f, 0x91, 0x50, 0x84,
	0xc0, 0x91, 0xc2, 0x0b, 0x54, 0xaa, 0x04, 0x45, 0xa1, 0x02, 0xe5, 0xa1, 0x0a, 0x27, 0xb0, 0xcd,
	0x34, 0x44, 0xb8, 0x5d, 0xb3, 0xbb, 0xae, 0xd4, 0x63, 0xf0, 0xc0, 0x25, 0x38, 0x1c, 0x67, 0xc0,
	0x59, 0xef, 0x6e, 0x36, 0xb6, 0x5b, 0x29, 0xa8, 0x4f, 0x5e, 0xcf, 0xf7, 0xcd, 0xb7, 0xf3, 0xcd,
	0xd8, 0x03, 0x5d, 0x9e, 0x25, 0x61, 0xc6, 0x99, 0x64, 0xd8, 0x11, 0xc4, 0x2f, 0x88, 0x67, 0xb1,
	0xff, 0x65, 0xbe, 0x90, 0xdf, 0xf3, 0x38, 0x4c, 0xd8, 0xd9, 0x28, 0x61, 0x9c, 0x98, 0xd0, 0x8f,
	0xd7, 0x71, 0xc4, 0xe9, 0x8c, 0x64, 0x94, 0x8e, 0x62, 0xc6, 0x64, 0x72, 0x3a, 0x1f, 0x09, 0xc9,
	0x78, 0x34, 0x27, 0xf3, 0xcc, 0x62, 0x73, 0x2a, 0x55, 0x83, 0x5f, 0x1e, 0xe0, 0x57, 0x4a, 0x29,
	0x91, 0xc7, 0x9c, 0xe5, 0xd9, 0x8c, 0x7e, 0xe6, 0x24, 0x24, 0xbe, 0x87, 0x76, 0x1a, 0xc5, 0x94,
	0x8a, 0x3d, 0xef, 0xe9, 0xd6, 0xb0, 0x37, 0x1e, 0x86, 0xe6, 0xf6, 0xb0, 0xce, 0x0e, 0xa7, 0x8a,
	0x3a, 0x39, 0x97, 0xfc, 0x72, 0xa6, 0xf3, 0xfc, 0x77, 0xd0, 0x73, 0xc2, 0xd8, 0x87, 0xad, 0x1f,
	0x74, 0x59, 0xa8, 0x79, 0xc3, 0xee, 0x6c, 0x79, 0xc4, 0x01, 0x6c, 0x5f, 0x44, 0x69, 0x4e, 0x7b,
	0x2d, 0x15, 0x2b, 0x5f, 0x0e, 0x5a, 0x6f, 0xbd, 0xe0, 0x10, 0x76, 0xd7, 0x2e, 0x11, 0x19, 0x3b,
	0x17, 0x84, 0x2f, 0x60, 0x7b, 0xbe, 0x0c, 0x28, 0x91, 0xde, 0xb8, 0x1f, 0x5a, 0x4f, 0x61, 0x49,
	0x2c, 0xe1, 0xe0, 0xb7, 0x07, 0x83, 0x32, 0xff, 0x84, 0xb3, 0xd3, 0x45, 0x4a, 0xc6, 0xd4, 0x51,
	0xc5, 0xd4, 0xcb, 0xaa, 0xa9, 0x75, 0xfe, 0x4d, 0xdb, 0x9a, 0xc0, 0xfd, 0xca, 0x35, 0xda, 0xd8,
	0x2b, 0xb8, 0x9d, 0x95, 0x21, 0x6d, 0x0d, 0x1d, 0x6b, 0x86, 0x6c, 0x28, 0xc1, 0x33, 0xb8, 0xa7,
	0xec, 0x1e, 0x93, 0x34, 0xc6, 0xee, 0x42, 0x6b, 0xf1, 0x4d, 0x17, 0x51, 0x9c, 0x02, 0x84, 0xbe,
	0xa2, 0x4c, 0x17, 0xc2, 0x70, 0x82, 0x03, 0x1d, 0x53, 0x69, 0x1b, 0x76, 0xf4, 0x10, 0x76, 0x1c,
	0x3d, 0x9d, 0x3c, 0x84, 0xb6, 0x42, 0x4d, 0x37, 0xeb, 0xd9, 0x1a, 0x0f, 0x3e, 0xc0, 0x8e, 0x76,
	0x71, 0x92, 0xdb, 0x9a, 0x37, 0x33, 0x3d, 0x00, 0x74, 0x25, 0xca, 0x12, 0x82, 0xe7, 0x56, 0xf8,
	0x9a, 0x66, 0x1c, 0xd9, 0x54, 0xd7, 0xfa, 0xff, 0x5e, 0xef, 0xb6, 0x74, 0x02, 0xbb, 0x6b, 0x51,
	0x2d, 0x1d, 0x42, 0x47, 0xe7, 0x99, 0xd6, 0x34, 0x69, 0x5b, 0xce, 0xb8, 0xf8, 0x5e, 0xdb, 0xaa,
	0x61, 0x02, 0x3f, 0x42, 0xc7, 0x0c, 0x09, 0x1f, 0xac, 0xbe, 0xce, 0xca, 0xbc, 0x7d, 0xbf, 0x09,
	0xd2, 0x3d, 0xb9, 0x85, 0x9f, 0xa0, 0x6b, 0xa7, 0x85, 0x55, 0xaa, 0x53, 0xbf, 0xff, 0xb0, 0x11,
	0x33, 0x3a, 0xe3, 0xbf, 0x1e, 0x74, 0x74, 0xb5, 0x02, 0x3f, 0x03, 0xac, 0x06, 0x80, 0x4e, 0x66,
	0x6d, 0xb2, 0xfe, 0x7e, 0x33, 0x68, 0xeb, 0x5b, 0x49, 0x2d, 0x6d, 0xd6, 0xa5, 0x1c, 0xa3, 0xfb,
	0xcd, 0xa0, 0x95, 0x9a, 0x42, 0xcf, 0x99, 0x00, 0xd6, 0xe9, 0xae, 0xdd, 0x47, 0x57, 0xa0, 0xd6,
	0xf0, 0x9f, 0x62, 0x10, 0xe5, 0x1f, 0xba, 0x14, 0x76, 0x56, 0x90, 0x2b, 0x5c, 0x5f, 0x7f, 0xae,
	0x70, 0xc3, 0xde, 0x2a, 0xca, 0x9c, 0xc1, 0x9d, 0xb5, 0x3f, 0x1f, 0x1f, 0x5f, 0xbf, 0x79, 0xfc,
	0x27, 0x57, 0xe2, 0x46, 0x33, 0x6e, 0xab, 0xfd, 0xfd, 0xe6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x30, 0x49, 0x0e, 0x1d, 0x22, 0x06, 0x00, 0x00,
}