// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sf/metering/v1/metering.proto

package pbmetering

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ApiKeyId  string `protobuf:"bytes,2,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// Defines the service that emitted the event (sf.firehose.v1/Blocks ...)
	Service string `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	// Defines the blockchain (eth-mainnet, sol-mainnet ...)
	Network   string                 `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
	Metrics   []*Metric              `protobuf:"bytes,20,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,30,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_metering_v1_metering_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_sf_metering_v1_metering_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_sf_metering_v1_metering_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Event) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

func (x *Event) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *Event) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Event) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Event) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Event) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_metering_v1_metering_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_sf_metering_v1_metering_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_sf_metering_v1_metering_proto_rawDescGZIP(), []int{1}
}

func (x *Metric) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Metric) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_sf_metering_v1_metering_proto protoreflect.FileDescriptor

var file_sf_metering_v1_metering_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x73, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x14, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x30, 0x0a,
	0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x43, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x04, 0x45,
	0x6d, 0x69, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74,
	0x2f, 0x64, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66,
	0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_metering_v1_metering_proto_rawDescOnce sync.Once
	file_sf_metering_v1_metering_proto_rawDescData = file_sf_metering_v1_metering_proto_rawDesc
)

func file_sf_metering_v1_metering_proto_rawDescGZIP() []byte {
	file_sf_metering_v1_metering_proto_rawDescOnce.Do(func() {
		file_sf_metering_v1_metering_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_metering_v1_metering_proto_rawDescData)
	})
	return file_sf_metering_v1_metering_proto_rawDescData
}

var file_sf_metering_v1_metering_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sf_metering_v1_metering_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: sf.metering.v1.Event
	(*Metric)(nil),                // 1: sf.metering.v1.Metric
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_sf_metering_v1_metering_proto_depIdxs = []int32{
	1, // 0: sf.metering.v1.Event.metrics:type_name -> sf.metering.v1.Metric
	2, // 1: sf.metering.v1.Event.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: sf.metering.v1.Metering.Emit:input_type -> sf.metering.v1.Event
	3, // 3: sf.metering.v1.Metering.Emit:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sf_metering_v1_metering_proto_init() }
func file_sf_metering_v1_metering_proto_init() {
	if File_sf_metering_v1_metering_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_metering_v1_metering_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_sf_metering_v1_metering_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
			RawDescriptor: file_sf_metering_v1_metering_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sf_metering_v1_metering_proto_goTypes,
		DependencyIndexes: file_sf_metering_v1_metering_proto_depIdxs,
		MessageInfos:      file_sf_metering_v1_metering_proto_msgTypes,
	}.Build()
	File_sf_metering_v1_metering_proto = out.File
	file_sf_metering_v1_metering_proto_rawDesc = nil
	file_sf_metering_v1_metering_proto_goTypes = nil
	file_sf_metering_v1_metering_proto_depIdxs = nil
}