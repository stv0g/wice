// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: interface.proto

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

type Interface_Type int32

const (
	Interface_UNKNOWN        Interface_Type = 0
	Interface_LINUX_KERNEL   Interface_Type = 1
	Interface_OPENBSD_KERNEL Interface_Type = 2
	Interface_WINDOWS_KERNEL Interface_Type = 3
	Interface_USERSPACE      Interface_Type = 4
)

// Enum value maps for Interface_Type.
var (
	Interface_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "LINUX_KERNEL",
		2: "OPENBSD_KERNEL",
		3: "WINDOWS_KERNEL",
		4: "USERSPACE",
	}
	Interface_Type_value = map[string]int32{
		"UNKNOWN":        0,
		"LINUX_KERNEL":   1,
		"OPENBSD_KERNEL": 2,
		"WINDOWS_KERNEL": 3,
		"USERSPACE":      4,
	}
)

func (x Interface_Type) Enum() *Interface_Type {
	p := new(Interface_Type)
	*p = x
	return p
}

func (x Interface_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Interface_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_interface_proto_enumTypes[0].Descriptor()
}

func (Interface_Type) Type() protoreflect.EnumType {
	return &file_interface_proto_enumTypes[0]
}

func (x Interface_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interface_Type.Descriptor instead.
func (Interface_Type) EnumDescriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{0, 0}
}

// A Wireguard interface
// See: https://pkg.go.dev/golang.zx2c4.com/wireguard/wgctrl/wgtypes#Device
type Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type         Interface_Type `protobuf:"varint,2,opt,name=type,proto3,enum=wice.Interface_Type" json:"type,omitempty"`
	PublicKey    []byte         `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PrivateKey   []byte         `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	ListenPort   uint32         `protobuf:"varint,5,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	FirewallMark uint32         `protobuf:"varint,6,opt,name=firewall_mark,json=firewallMark,proto3" json:"firewall_mark,omitempty"`
	Peers        []*Peer        `protobuf:"bytes,7,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *Interface) Reset() {
	*x = Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface) ProtoMessage() {}

func (x *Interface) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface.ProtoReflect.Descriptor instead.
func (*Interface) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{0}
}

func (x *Interface) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Interface) GetType() Interface_Type {
	if x != nil {
		return x.Type
	}
	return Interface_UNKNOWN
}

func (x *Interface) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Interface) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *Interface) GetListenPort() uint32 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *Interface) GetFirewallMark() uint32 {
	if x != nil {
		return x.FirewallMark
	}
	return 0
}

func (x *Interface) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_interface_proto protoreflect.FileDescriptor

var file_interface_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x77, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x5c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x4c, 0x49, 0x4e, 0x55, 0x58, 0x5f, 0x4b, 0x45, 0x52, 0x4e, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x4f, 0x50, 0x45, 0x4e, 0x42, 0x53, 0x44, 0x5f, 0x4b, 0x45, 0x52, 0x4e, 0x45, 0x4c,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x5f, 0x4b, 0x45,
	0x52, 0x4e, 0x45, 0x4c, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x53, 0x45, 0x52, 0x53, 0x50,
	0x41, 0x43, 0x45, 0x10, 0x04, 0x42, 0x16, 0x5a, 0x14, 0x72, 0x69, 0x61, 0x73, 0x63, 0x2e, 0x65,
	0x75, 0x2f, 0x77, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interface_proto_rawDescOnce sync.Once
	file_interface_proto_rawDescData = file_interface_proto_rawDesc
)

func file_interface_proto_rawDescGZIP() []byte {
	file_interface_proto_rawDescOnce.Do(func() {
		file_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_interface_proto_rawDescData)
	})
	return file_interface_proto_rawDescData
}

var file_interface_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_interface_proto_goTypes = []interface{}{
	(Interface_Type)(0), // 0: wice.Interface.Type
	(*Interface)(nil),   // 1: wice.Interface
	(*Peer)(nil),        // 2: wice.Peer
}
var file_interface_proto_depIdxs = []int32{
	0, // 0: wice.Interface.type:type_name -> wice.Interface.Type
	2, // 1: wice.Interface.peers:type_name -> wice.Peer
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_interface_proto_init() }
func file_interface_proto_init() {
	if File_interface_proto != nil {
		return
	}
	file_peer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface); i {
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
			RawDescriptor: file_interface_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interface_proto_goTypes,
		DependencyIndexes: file_interface_proto_depIdxs,
		EnumInfos:         file_interface_proto_enumTypes,
		MessageInfos:      file_interface_proto_msgTypes,
	}.Build()
	File_interface_proto = out.File
	file_interface_proto_rawDesc = nil
	file_interface_proto_goTypes = nil
	file_interface_proto_depIdxs = nil
}
