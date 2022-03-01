// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: proto/timeservice.proto

package gRPC_Streaming

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

// The request message containing the duration_secs
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DurationSecs uint32 `protobuf:"varint,2,opt,name=duration_secs,json=durationSecs,proto3" json:"duration_secs,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timeservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timeservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_timeservice_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetDurationSecs() uint32 {
	if x != nil {
		return x.DurationSecs
	}
	return 0
}

// The response message containing the current_time
type TimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
}

func (x *TimeResponse) Reset() {
	*x = TimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timeservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeResponse) ProtoMessage() {}

func (x *TimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timeservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeResponse.ProtoReflect.Descriptor instead.
func (*TimeResponse) Descriptor() ([]byte, []int) {
	return file_proto_timeservice_proto_rawDescGZIP(), []int{1}
}

func (x *TimeResponse) GetCurrentTime() *timestamp.Timestamp {
	if x != nil {
		return x.CurrentTime
	}
	return nil
}

var File_proto_timeservice_proto protoreflect.FileDescriptor

var file_proto_timeservice_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2e, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63,
	0x73, 0x22, 0x4d, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x32, 0x44, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2d,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x3b, 0x67, 0x52, 0x50, 0x43, 0x5f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_timeservice_proto_rawDescOnce sync.Once
	file_proto_timeservice_proto_rawDescData = file_proto_timeservice_proto_rawDesc
)

func file_proto_timeservice_proto_rawDescGZIP() []byte {
	file_proto_timeservice_proto_rawDescOnce.Do(func() {
		file_proto_timeservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_timeservice_proto_rawDescData)
	})
	return file_proto_timeservice_proto_rawDescData
}

var file_proto_timeservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_timeservice_proto_goTypes = []interface{}{
	(*Request)(nil),             // 0: proto.Request
	(*TimeResponse)(nil),        // 1: proto.TimeResponse
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_timeservice_proto_depIdxs = []int32{
	2, // 0: proto.TimeResponse.current_time:type_name -> google.protobuf.Timestamp
	0, // 1: proto.TimeService.StreamTime:input_type -> proto.Request
	1, // 2: proto.TimeService.StreamTime:output_type -> proto.TimeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_timeservice_proto_init() }
func file_proto_timeservice_proto_init() {
	if File_proto_timeservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_timeservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_timeservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeResponse); i {
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
			RawDescriptor: file_proto_timeservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_timeservice_proto_goTypes,
		DependencyIndexes: file_proto_timeservice_proto_depIdxs,
		MessageInfos:      file_proto_timeservice_proto_msgTypes,
	}.Build()
	File_proto_timeservice_proto = out.File
	file_proto_timeservice_proto_rawDesc = nil
	file_proto_timeservice_proto_goTypes = nil
	file_proto_timeservice_proto_depIdxs = nil
}
