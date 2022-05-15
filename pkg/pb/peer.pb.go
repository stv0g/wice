// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: peer.proto

package pb

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

// Type of the Wireguard interface impelementation
type PeerDescription_InterfaceType int32

const (
	PeerDescription_UNKNOWN        PeerDescription_InterfaceType = 0
	PeerDescription_LINUX_KERNEL   PeerDescription_InterfaceType = 1
	PeerDescription_OPENBSD_KERNEL PeerDescription_InterfaceType = 2
	PeerDescription_WINDOWS_KERNEL PeerDescription_InterfaceType = 3
	PeerDescription_USERSPACE      PeerDescription_InterfaceType = 4
)

// Enum value maps for PeerDescription_InterfaceType.
var (
	PeerDescription_InterfaceType_name = map[int32]string{
		0: "UNKNOWN",
		1: "LINUX_KERNEL",
		2: "OPENBSD_KERNEL",
		3: "WINDOWS_KERNEL",
		4: "USERSPACE",
	}
	PeerDescription_InterfaceType_value = map[string]int32{
		"UNKNOWN":        0,
		"LINUX_KERNEL":   1,
		"OPENBSD_KERNEL": 2,
		"WINDOWS_KERNEL": 3,
		"USERSPACE":      4,
	}
)

func (x PeerDescription_InterfaceType) Enum() *PeerDescription_InterfaceType {
	p := new(PeerDescription_InterfaceType)
	*p = x
	return p
}

func (x PeerDescription_InterfaceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerDescription_InterfaceType) Descriptor() protoreflect.EnumDescriptor {
	return file_peer_proto_enumTypes[0].Descriptor()
}

func (PeerDescription_InterfaceType) Type() protoreflect.EnumType {
	return &file_peer_proto_enumTypes[0]
}

func (x PeerDescription_InterfaceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerDescription_InterfaceType.Descriptor instead.
func (PeerDescription_InterfaceType) EnumDescriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{0, 0}
}

type PeerDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hostname of the node
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// List of allowed IPs
	AllowedIps []string `protobuf:"bytes,3,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
	// Wireguard endpoint address
	Endpoint string `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Version of ɯice agent
	WiceVersion   string                        `protobuf:"bytes,5,opt,name=wice_version,json=wiceVersion,proto3" json:"wice_version,omitempty"`
	InterfaceType PeerDescription_InterfaceType `protobuf:"varint,6,opt,name=interface_type,json=interfaceType,proto3,enum=wice.PeerDescription_InterfaceType" json:"interface_type,omitempty"`
}

func (x *PeerDescription) Reset() {
	*x = PeerDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerDescription) ProtoMessage() {}

func (x *PeerDescription) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerDescription.ProtoReflect.Descriptor instead.
func (*PeerDescription) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{0}
}

func (x *PeerDescription) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *PeerDescription) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

func (x *PeerDescription) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *PeerDescription) GetWiceVersion() string {
	if x != nil {
		return x.WiceVersion
	}
	return ""
}

func (x *PeerDescription) GetInterfaceType() PeerDescription_InterfaceType {
	if x != nil {
		return x.InterfaceType
	}
	return PeerDescription_UNKNOWN
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey                   []byte          `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Endpoint                    string          `protobuf:"bytes,2,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	PersistentKeepaliveInterval uint32          `protobuf:"varint,3,opt,name=persistent_keepalive_interval,json=persistentKeepaliveInterval,proto3" json:"persistent_keepalive_interval,omitempty"`
	LastHandshake               *Timestamp      `protobuf:"bytes,4,opt,name=last_handshake,json=lastHandshake,proto3" json:"last_handshake,omitempty"`
	TransmitBytes               int64           `protobuf:"varint,5,opt,name=transmit_bytes,json=transmitBytes,proto3" json:"transmit_bytes,omitempty"`
	ReceiveBytes                int64           `protobuf:"varint,6,opt,name=receive_bytes,json=receiveBytes,proto3" json:"receive_bytes,omitempty"`
	AllowedIps                  []string        `protobuf:"bytes,7,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
	ProtocolVersion             uint32          `protobuf:"varint,8,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	PresharedKey                []byte          `protobuf:"bytes,9,opt,name=preshared_key,json=presharedKey,proto3" json:"preshared_key,omitempty"`
	State                       ConnectionState `protobuf:"varint,10,opt,name=state,proto3,enum=wice.ConnectionState" json:"state,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{1}
}

func (x *Peer) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Peer) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Peer) GetPersistentKeepaliveInterval() uint32 {
	if x != nil {
		return x.PersistentKeepaliveInterval
	}
	return 0
}

func (x *Peer) GetLastHandshake() *Timestamp {
	if x != nil {
		return x.LastHandshake
	}
	return nil
}

func (x *Peer) GetTransmitBytes() int64 {
	if x != nil {
		return x.TransmitBytes
	}
	return 0
}

func (x *Peer) GetReceiveBytes() int64 {
	if x != nil {
		return x.ReceiveBytes
	}
	return 0
}

func (x *Peer) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

func (x *Peer) GetProtocolVersion() uint32 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *Peer) GetPresharedKey() []byte {
	if x != nil {
		return x.PresharedKey
	}
	return nil
}

func (x *Peer) GetState() ConnectionState {
	if x != nil {
		return x.State
	}
	return ConnectionState_NEW
}

var File_peer_proto protoreflect.FileDescriptor

var file_peer_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x77, 0x69,
	0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc0, 0x02, 0x0a, 0x0f, 0x50, 0x65, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x4a, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x65, 0x0a, 0x0d,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x49,
	0x4e, 0x55, 0x58, 0x5f, 0x4b, 0x45, 0x52, 0x4e, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x4f, 0x50, 0x45, 0x4e, 0x42, 0x53, 0x44, 0x5f, 0x4b, 0x45, 0x52, 0x4e, 0x45, 0x4c, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x5f, 0x4b, 0x45, 0x52, 0x4e,
	0x45, 0x4c, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x53, 0x45, 0x52, 0x53, 0x50, 0x41, 0x43,
	0x45, 0x10, 0x04, 0x22, 0xa7, 0x03, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x1d, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68,
	0x61, 0x6b, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70, 0x73,
	0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x72, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79,
	0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x16, 0x5a,
	0x14, 0x72, 0x69, 0x61, 0x73, 0x63, 0x2e, 0x65, 0x75, 0x2f, 0x77, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_proto_rawDescOnce sync.Once
	file_peer_proto_rawDescData = file_peer_proto_rawDesc
)

func file_peer_proto_rawDescGZIP() []byte {
	file_peer_proto_rawDescOnce.Do(func() {
		file_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_proto_rawDescData)
	})
	return file_peer_proto_rawDescData
}

var file_peer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_peer_proto_goTypes = []interface{}{
	(PeerDescription_InterfaceType)(0), // 0: wice.PeerDescription.InterfaceType
	(*PeerDescription)(nil),            // 1: wice.PeerDescription
	(*Peer)(nil),                       // 2: wice.Peer
	(*Timestamp)(nil),                  // 3: wice.Timestamp
	(ConnectionState)(0),               // 4: wice.ConnectionState
}
var file_peer_proto_depIdxs = []int32{
	0, // 0: wice.PeerDescription.interface_type:type_name -> wice.PeerDescription.InterfaceType
	3, // 1: wice.Peer.last_handshake:type_name -> wice.Timestamp
	4, // 2: wice.Peer.state:type_name -> wice.ConnectionState
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_peer_proto_init() }
func file_peer_proto_init() {
	if File_peer_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerDescription); i {
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
		file_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
			RawDescriptor: file_peer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_proto_goTypes,
		DependencyIndexes: file_peer_proto_depIdxs,
		EnumInfos:         file_peer_proto_enumTypes,
		MessageInfos:      file_peer_proto_msgTypes,
	}.Build()
	File_peer_proto = out.File
	file_peer_proto_rawDesc = nil
	file_peer_proto_goTypes = nil
	file_peer_proto_depIdxs = nil
}
