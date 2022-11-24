// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: ClientAbilityInitFinishNotify.proto

package proto

import (
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

// CmdId: 1135
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ClientAbilityInitFinishNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invokes  []*AbilityInvokeEntry `protobuf:"bytes,14,rep,name=invokes,proto3" json:"invokes,omitempty"`
	EntityId uint32                `protobuf:"varint,11,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *ClientAbilityInitFinishNotify) Reset() {
	*x = ClientAbilityInitFinishNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientAbilityInitFinishNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAbilityInitFinishNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAbilityInitFinishNotify) ProtoMessage() {}

func (x *ClientAbilityInitFinishNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ClientAbilityInitFinishNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAbilityInitFinishNotify.ProtoReflect.Descriptor instead.
func (*ClientAbilityInitFinishNotify) Descriptor() ([]byte, []int) {
	return file_ClientAbilityInitFinishNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ClientAbilityInitFinishNotify) GetInvokes() []*AbilityInvokeEntry {
	if x != nil {
		return x.Invokes
	}
	return nil
}

func (x *ClientAbilityInitFinishNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_ClientAbilityInitFinishNotify_proto protoreflect.FileDescriptor

var file_ClientAbilityInitFinishNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x69, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x1d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x33, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientAbilityInitFinishNotify_proto_rawDescOnce sync.Once
	file_ClientAbilityInitFinishNotify_proto_rawDescData = file_ClientAbilityInitFinishNotify_proto_rawDesc
)

func file_ClientAbilityInitFinishNotify_proto_rawDescGZIP() []byte {
	file_ClientAbilityInitFinishNotify_proto_rawDescOnce.Do(func() {
		file_ClientAbilityInitFinishNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientAbilityInitFinishNotify_proto_rawDescData)
	})
	return file_ClientAbilityInitFinishNotify_proto_rawDescData
}

var file_ClientAbilityInitFinishNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClientAbilityInitFinishNotify_proto_goTypes = []interface{}{
	(*ClientAbilityInitFinishNotify)(nil), // 0: proto.ClientAbilityInitFinishNotify
	(*AbilityInvokeEntry)(nil),            // 1: proto.AbilityInvokeEntry
}
var file_ClientAbilityInitFinishNotify_proto_depIdxs = []int32{
	1, // 0: proto.ClientAbilityInitFinishNotify.invokes:type_name -> proto.AbilityInvokeEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ClientAbilityInitFinishNotify_proto_init() }
func file_ClientAbilityInitFinishNotify_proto_init() {
	if File_ClientAbilityInitFinishNotify_proto != nil {
		return
	}
	file_AbilityInvokeEntry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ClientAbilityInitFinishNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAbilityInitFinishNotify); i {
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
			RawDescriptor: file_ClientAbilityInitFinishNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClientAbilityInitFinishNotify_proto_goTypes,
		DependencyIndexes: file_ClientAbilityInitFinishNotify_proto_depIdxs,
		MessageInfos:      file_ClientAbilityInitFinishNotify_proto_msgTypes,
	}.Build()
	File_ClientAbilityInitFinishNotify_proto = out.File
	file_ClientAbilityInitFinishNotify_proto_rawDesc = nil
	file_ClientAbilityInitFinishNotify_proto_goTypes = nil
	file_ClientAbilityInitFinishNotify_proto_depIdxs = nil
}