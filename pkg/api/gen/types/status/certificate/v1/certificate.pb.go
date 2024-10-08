// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: types/status/certificate/v1/certificate.proto

package certificatev1

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

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCnameSet bool   `protobuf:"varint,1,opt,name=is_cname_set,json=isCnameSet,proto3" json:"is_cname_set,omitempty"`
	IsIssued   bool   `protobuf:"varint,2,opt,name=is_issued,json=isIssued,proto3" json:"is_issued,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_status_certificate_v1_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_types_status_certificate_v1_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_types_status_certificate_v1_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetIsCnameSet() bool {
	if x != nil {
		return x.IsCnameSet
	}
	return false
}

func (x *Status) GetIsIssued() bool {
	if x != nil {
		return x.IsIssued
	}
	return false
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_types_status_certificate_v1_certificate_proto protoreflect.FileDescriptor

var file_types_status_certificate_v1_certificate_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x61, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xc2, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x10, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x69, 0x6f, 0x2f, 0x61, 0x6b, 0x75, 0x69, 0x74,
	0x79, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x41, 0x54, 0x53, 0x43, 0xaa, 0x02, 0x22, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x22, 0x41,
	0x6b, 0x75, 0x69, 0x74, 0x79, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x2e, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x5c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x26, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x3a, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_types_status_certificate_v1_certificate_proto_rawDescOnce sync.Once
	file_types_status_certificate_v1_certificate_proto_rawDescData = file_types_status_certificate_v1_certificate_proto_rawDesc
)

func file_types_status_certificate_v1_certificate_proto_rawDescGZIP() []byte {
	file_types_status_certificate_v1_certificate_proto_rawDescOnce.Do(func() {
		file_types_status_certificate_v1_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_status_certificate_v1_certificate_proto_rawDescData)
	})
	return file_types_status_certificate_v1_certificate_proto_rawDescData
}

var file_types_status_certificate_v1_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_status_certificate_v1_certificate_proto_goTypes = []interface{}{
	(*Status)(nil), // 0: akuity.types.status.certificate.v1.Status
}
var file_types_status_certificate_v1_certificate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_status_certificate_v1_certificate_proto_init() }
func file_types_status_certificate_v1_certificate_proto_init() {
	if File_types_status_certificate_v1_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_status_certificate_v1_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_types_status_certificate_v1_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_status_certificate_v1_certificate_proto_goTypes,
		DependencyIndexes: file_types_status_certificate_v1_certificate_proto_depIdxs,
		MessageInfos:      file_types_status_certificate_v1_certificate_proto_msgTypes,
	}.Build()
	File_types_status_certificate_v1_certificate_proto = out.File
	file_types_status_certificate_v1_certificate_proto_rawDesc = nil
	file_types_status_certificate_v1_certificate_proto_goTypes = nil
	file_types_status_certificate_v1_certificate_proto_depIdxs = nil
}
