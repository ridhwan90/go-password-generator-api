// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: rpc_list_user_password.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserUuid          string               `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Site              string               `protobuf:"bytes,3,opt,name=site,proto3" json:"site,omitempty"`
	SiteUsername      string               `protobuf:"bytes,4,opt,name=site_username,json=siteUsername,proto3" json:"site_username,omitempty"`
	SiteEmail         string               `protobuf:"bytes,5,opt,name=site_email,json=siteEmail,proto3" json:"site_email,omitempty"`
	GeneratedPassword string               `protobuf:"bytes,6,opt,name=generated_password,json=generatedPassword,proto3" json:"generated_password,omitempty"`
	CreatedAt         *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UserPassword) Reset() {
	*x = UserPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_list_user_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPassword) ProtoMessage() {}

func (x *UserPassword) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_list_user_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPassword.ProtoReflect.Descriptor instead.
func (*UserPassword) Descriptor() ([]byte, []int) {
	return file_rpc_list_user_password_proto_rawDescGZIP(), []int{0}
}

func (x *UserPassword) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserPassword) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *UserPassword) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *UserPassword) GetSiteUsername() string {
	if x != nil {
		return x.SiteUsername
	}
	return ""
}

func (x *UserPassword) GetSiteEmail() string {
	if x != nil {
		return x.SiteEmail
	}
	return ""
}

func (x *UserPassword) GetGeneratedPassword() string {
	if x != nil {
		return x.GeneratedPassword
	}
	return ""
}

func (x *UserPassword) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ListUserPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPassword []*UserPassword `protobuf:"bytes,1,rep,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
}

func (x *ListUserPassword) Reset() {
	*x = ListUserPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_list_user_password_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserPassword) ProtoMessage() {}

func (x *ListUserPassword) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_list_user_password_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserPassword.ProtoReflect.Descriptor instead.
func (*ListUserPassword) Descriptor() ([]byte, []int) {
	return file_rpc_list_user_password_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserPassword) GetUserPassword() []*UserPassword {
	if x != nil {
		return x.UserPassword
	}
	return nil
}

type UsernameParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UsernameParams) Reset() {
	*x = UsernameParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_list_user_password_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernameParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameParams) ProtoMessage() {}

func (x *UsernameParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_list_user_password_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameParams.ProtoReflect.Descriptor instead.
func (*UsernameParams) Descriptor() ([]byte, []int) {
	return file_rpc_list_user_password_proto_rawDescGZIP(), []int{2}
}

func (x *UsernameParams) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_rpc_list_user_password_proto protoreflect.FileDescriptor

var file_rpc_list_user_password_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69,
	0x74, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x12, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x49, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2c,
	0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x64, 0x68, 0x77,
	0x61, 0x6e, 0x39, 0x30, 0x2f, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_list_user_password_proto_rawDescOnce sync.Once
	file_rpc_list_user_password_proto_rawDescData = file_rpc_list_user_password_proto_rawDesc
)

func file_rpc_list_user_password_proto_rawDescGZIP() []byte {
	file_rpc_list_user_password_proto_rawDescOnce.Do(func() {
		file_rpc_list_user_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_list_user_password_proto_rawDescData)
	})
	return file_rpc_list_user_password_proto_rawDescData
}

var file_rpc_list_user_password_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_list_user_password_proto_goTypes = []interface{}{
	(*UserPassword)(nil),        // 0: pb.UserPassword
	(*ListUserPassword)(nil),    // 1: pb.ListUserPassword
	(*UsernameParams)(nil),      // 2: pb.UsernameParams
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_rpc_list_user_password_proto_depIdxs = []int32{
	3, // 0: pb.UserPassword.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: pb.ListUserPassword.user_password:type_name -> pb.UserPassword
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_list_user_password_proto_init() }
func file_rpc_list_user_password_proto_init() {
	if File_rpc_list_user_password_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_list_user_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPassword); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_list_user_password_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserPassword); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_list_user_password_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernameParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_list_user_password_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_list_user_password_proto_goTypes,
		DependencyIndexes: file_rpc_list_user_password_proto_depIdxs,
		MessageInfos:      file_rpc_list_user_password_proto_msgTypes,
	}.Build()
	File_rpc_list_user_password_proto = out.File
	file_rpc_list_user_password_proto_rawDesc = nil
	file_rpc_list_user_password_proto_goTypes = nil
	file_rpc_list_user_password_proto_depIdxs = nil
}
