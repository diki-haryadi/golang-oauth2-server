// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: shared/error/v1/custom_error.proto

package error_v1

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

type CustomError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code       int64             `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg        string            `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp  string            `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	StackTrace *string           `protobuf:"bytes,5,opt,name=stack_trace,json=stackTrace,proto3,oneof" json:"stack_trace,omitempty"`
	Details    map[string]string `protobuf:"bytes,6,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CustomError) Reset() {
	*x = CustomError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_error_v1_custom_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomError) ProtoMessage() {}

func (x *CustomError) ProtoReflect() protoreflect.Message {
	mi := &file_shared_error_v1_custom_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomError.ProtoReflect.Descriptor instead.
func (*CustomError) Descriptor() ([]byte, []int) {
	return file_shared_error_v1_custom_error_proto_rawDescGZIP(), []int{0}
}

func (x *CustomError) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CustomError) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CustomError) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CustomError) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CustomError) GetStackTrace() string {
	if x != nil && x.StackTrace != nil {
		return *x.StackTrace
	}
	return ""
}

func (x *CustomError) GetDetails() map[string]string {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_shared_error_v1_custom_error_proto protoreflect.FileDescriptor

var file_shared_error_v1_custom_error_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x9e, 0x02, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x24, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_error_v1_custom_error_proto_rawDescOnce sync.Once
	file_shared_error_v1_custom_error_proto_rawDescData = file_shared_error_v1_custom_error_proto_rawDesc
)

func file_shared_error_v1_custom_error_proto_rawDescGZIP() []byte {
	file_shared_error_v1_custom_error_proto_rawDescOnce.Do(func() {
		file_shared_error_v1_custom_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_error_v1_custom_error_proto_rawDescData)
	})
	return file_shared_error_v1_custom_error_proto_rawDescData
}

var file_shared_error_v1_custom_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shared_error_v1_custom_error_proto_goTypes = []interface{}{
	(*CustomError)(nil), // 0: shared.error.v1.CustomError
	nil,                 // 1: shared.error.v1.CustomError.DetailsEntry
}
var file_shared_error_v1_custom_error_proto_depIdxs = []int32{
	1, // 0: shared.error.v1.CustomError.details:type_name -> shared.error.v1.CustomError.DetailsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shared_error_v1_custom_error_proto_init() }
func file_shared_error_v1_custom_error_proto_init() {
	if File_shared_error_v1_custom_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_error_v1_custom_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomError); i {
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
	file_shared_error_v1_custom_error_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_error_v1_custom_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_error_v1_custom_error_proto_goTypes,
		DependencyIndexes: file_shared_error_v1_custom_error_proto_depIdxs,
		MessageInfos:      file_shared_error_v1_custom_error_proto_msgTypes,
	}.Build()
	File_shared_error_v1_custom_error_proto = out.File
	file_shared_error_v1_custom_error_proto_rawDesc = nil
	file_shared_error_v1_custom_error_proto_goTypes = nil
	file_shared_error_v1_custom_error_proto_depIdxs = nil
}
