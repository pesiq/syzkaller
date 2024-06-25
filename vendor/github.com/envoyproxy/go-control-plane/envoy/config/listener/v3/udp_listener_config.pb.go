// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/config/listener/v3/udp_listener_config.proto

package listenerv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

// [#next-free-field: 9]
type UdpListenerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UDP socket configuration for the listener. The default for
	// :ref:`prefer_gro <envoy_v3_api_field_config.core.v3.UdpSocketConfig.prefer_gro>` is false for
	// listener sockets. If receiving a large amount of datagrams from a small number of sources, it
	// may be worthwhile to enable this option after performance testing.
	DownstreamSocketConfig *v3.UdpSocketConfig `protobuf:"bytes,5,opt,name=downstream_socket_config,json=downstreamSocketConfig,proto3" json:"downstream_socket_config,omitempty"`
	// Configuration for QUIC protocol. If empty, QUIC will not be enabled on this listener. Set
	// to the default object to enable QUIC without modifying any additional options.
	QuicOptions *QuicProtocolOptions `protobuf:"bytes,7,opt,name=quic_options,json=quicOptions,proto3" json:"quic_options,omitempty"`
	// Configuration for the UDP packet writer. If empty, HTTP/3 will use GSO if available
	// (:ref:`UdpDefaultWriterFactory <envoy_v3_api_msg_extensions.udp_packet_writer.v3.UdpGsoBatchWriterFactory>`)
	// or the default kernel sendmsg if not,
	// (:ref:`UdpDefaultWriterFactory <envoy_v3_api_msg_extensions.udp_packet_writer.v3.UdpDefaultWriterFactory>`)
	// and raw UDP will use kernel sendmsg.
	// [#extension-category: envoy.udp_packet_writer]
	UdpPacketPacketWriterConfig *v3.TypedExtensionConfig `protobuf:"bytes,8,opt,name=udp_packet_packet_writer_config,json=udpPacketPacketWriterConfig,proto3" json:"udp_packet_packet_writer_config,omitempty"`
}

func (x *UdpListenerConfig) Reset() {
	*x = UdpListenerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpListenerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpListenerConfig) ProtoMessage() {}

func (x *UdpListenerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpListenerConfig.ProtoReflect.Descriptor instead.
func (*UdpListenerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP(), []int{0}
}

func (x *UdpListenerConfig) GetDownstreamSocketConfig() *v3.UdpSocketConfig {
	if x != nil {
		return x.DownstreamSocketConfig
	}
	return nil
}

func (x *UdpListenerConfig) GetQuicOptions() *QuicProtocolOptions {
	if x != nil {
		return x.QuicOptions
	}
	return nil
}

func (x *UdpListenerConfig) GetUdpPacketPacketWriterConfig() *v3.TypedExtensionConfig {
	if x != nil {
		return x.UdpPacketPacketWriterConfig
	}
	return nil
}

type ActiveRawUdpListenerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActiveRawUdpListenerConfig) Reset() {
	*x = ActiveRawUdpListenerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRawUdpListenerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRawUdpListenerConfig) ProtoMessage() {}

func (x *ActiveRawUdpListenerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRawUdpListenerConfig.ProtoReflect.Descriptor instead.
func (*ActiveRawUdpListenerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP(), []int{1}
}

var File_envoy_config_listener_v3_udp_listener_config_proto protoreflect.FileDescriptor

var file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x24,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x71, 0x75, 0x69,
	0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8e, 0x03, 0x0a, 0x11, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5f, 0x0a, 0x18, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x55, 0x64, 0x70, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x16, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0c, 0x71, 0x75, 0x69, 0x63, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x75, 0x69, 0x63, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x71, 0x75,
	0x69, 0x63, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x70, 0x0a, 0x1f, 0x75, 0x64, 0x70,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x1b,
	0x75, 0x64, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x2e, 0x9a, 0xc5, 0x88,
	0x1e, 0x29, 0x0a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x04, 0x08, 0x01, 0x10,
	0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08,
	0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x55, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x77, 0x55, 0x64,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a,
	0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x77, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x96, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x16, 0x55, 0x64, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x76,
	0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescOnce sync.Once
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData = file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc
)

func file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP() []byte {
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData)
	})
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData
}

var file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_listener_v3_udp_listener_config_proto_goTypes = []interface{}{
	(*UdpListenerConfig)(nil),          // 0: envoy.config.listener.v3.UdpListenerConfig
	(*ActiveRawUdpListenerConfig)(nil), // 1: envoy.config.listener.v3.ActiveRawUdpListenerConfig
	(*v3.UdpSocketConfig)(nil),         // 2: envoy.config.core.v3.UdpSocketConfig
	(*QuicProtocolOptions)(nil),        // 3: envoy.config.listener.v3.QuicProtocolOptions
	(*v3.TypedExtensionConfig)(nil),    // 4: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs = []int32{
	2, // 0: envoy.config.listener.v3.UdpListenerConfig.downstream_socket_config:type_name -> envoy.config.core.v3.UdpSocketConfig
	3, // 1: envoy.config.listener.v3.UdpListenerConfig.quic_options:type_name -> envoy.config.listener.v3.QuicProtocolOptions
	4, // 2: envoy.config.listener.v3.UdpListenerConfig.udp_packet_packet_writer_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_config_listener_v3_udp_listener_config_proto_init() }
func file_envoy_config_listener_v3_udp_listener_config_proto_init() {
	if File_envoy_config_listener_v3_udp_listener_config_proto != nil {
		return
	}
	file_envoy_config_listener_v3_quic_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpListenerConfig); i {
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
		file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRawUdpListenerConfig); i {
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
			RawDescriptor: file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_listener_v3_udp_listener_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes,
	}.Build()
	File_envoy_config_listener_v3_udp_listener_config_proto = out.File
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc = nil
	file_envoy_config_listener_v3_udp_listener_config_proto_goTypes = nil
	file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs = nil
}
