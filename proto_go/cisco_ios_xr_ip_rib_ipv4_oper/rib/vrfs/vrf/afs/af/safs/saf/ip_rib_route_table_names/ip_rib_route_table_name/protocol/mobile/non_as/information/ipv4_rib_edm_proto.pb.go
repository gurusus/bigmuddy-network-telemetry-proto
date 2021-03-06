// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_proto.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_mobile_non_as_information is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_proto.proto

It has these top-level messages:
	Ipv4RibEdmProto_KEYS
	Ipv4RibEdmProto
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_mobile_non_as_information

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Information of a rib protocol
type Ipv4RibEdmProto_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
}

func (m *Ipv4RibEdmProto_KEYS) Reset()                    { *m = Ipv4RibEdmProto_KEYS{} }
func (m *Ipv4RibEdmProto_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

type Ipv4RibEdmProto struct {
	// Name
	ProtocolNames string `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames" json:"protocol_names,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance" json:"instance,omitempty"`
	// Proto version
	Version uint32 `protobuf:"varint,52,opt,name=version" json:"version,omitempty"`
	// Number of redist clients
	RedistributionClientCount uint32 `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount" json:"redistribution_client_count,omitempty"`
	// Number of proto clients
	ProtocolClientsCount uint32 `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount" json:"protocol_clients_count,omitempty"`
	// Number of routes (including active, backup and deleted), where, number of backup routes = RoutesCounts - ActiveRoutesCount - DeletedRoutesCount
	RoutesCounts uint32 `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts" json:"routes_counts,omitempty"`
	// Number of active routes (not deleted)
	ActiveRoutesCount uint32 `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount" json:"active_routes_count,omitempty"`
	// Number of deleted routes
	DeletedRoutesCount uint32 `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount" json:"deleted_routes_count,omitempty"`
	// Number of paths for all routes
	PathsCount uint32 `protobuf:"varint,58,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// Memory for proto's routes and paths in bytes
	ProtocolRouteMemory uint32 `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory" json:"protocol_route_memory,omitempty"`
	// Number of backup routes
	BackupRoutesCount uint32 `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount" json:"backup_routes_count,omitempty"`
}

func (m *Ipv4RibEdmProto) Reset()                    { *m = Ipv4RibEdmProto{} }
func (m *Ipv4RibEdmProto) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto) ProtoMessage()               {}
func (*Ipv4RibEdmProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.mobile.non_as.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.mobile.non_as.information.ipv4_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x15, 0x16, 0x6d, 0xcb, 0x40, 0x57, 0xe0, 0x5d, 0x58, 0x17, 0x0e, 0xac, 0x16, 0x21,
	0xf5, 0x64, 0xa1, 0xdd, 0xf2, 0x1f, 0x71, 0x59, 0x71, 0x42, 0x70, 0x28, 0x5c, 0x38, 0x59, 0x4e,
	0x3a, 0x11, 0x16, 0x89, 0x1d, 0x79, 0x9c, 0x08, 0x5e, 0x82, 0x03, 0xaf, 0xca, 0x0b, 0xa0, 0x8c,
	0x93, 0xd2, 0x52, 0x71, 0x71, 0xe5, 0xf9, 0x7e, 0x3f, 0xf7, 0x9b, 0x16, 0xa4, 0x6d, 0xba, 0xa5,
	0x0e, 0x36, 0xd7, 0xb8, 0xae, 0x75, 0x13, 0x7c, 0xf4, 0x8a, 0x4f, 0xf1, 0x33, 0x2b, 0x2c, 0x15,
	0x5e, 0x5b, 0x4f, 0xfa, 0x7b, 0xd0, 0xb6, 0x61, 0x8a, 0x71, 0xdf, 0x60, 0x50, 0xc1, 0xe6, 0xaa,
	0x0b, 0x25, 0xf5, 0x87, 0x32, 0x25, 0x29, 0x53, 0x2a, 0xea, 0x3f, 0xc9, 0x94, 0x6a, 0xa0, 0x83,
	0x6f, 0x23, 0xea, 0x68, 0xf2, 0x0a, 0xb5, 0x33, 0x35, 0xd2, 0xff, 0x82, 0xf4, 0x9d, 0x85, 0xaf,
	0x54, 0xed, 0x73, 0x5b, 0xa1, 0x72, 0xde, 0x69, 0x43, 0xca, 0xba, 0xd2, 0x87, 0xda, 0x44, 0xeb,
	0xdd, 0xf9, 0xaf, 0x0c, 0x4e, 0xf7, 0xdb, 0xea, 0xf7, 0xef, 0xbe, 0x7c, 0x12, 0x73, 0x98, 0x76,
	0xa1, 0xe4, 0xa7, 0x64, 0x76, 0x96, 0x2d, 0x6e, 0xac, 0x26, 0x5d, 0x28, 0x3f, 0x9a, 0x1a, 0xc5,
	0x29, 0x4c, 0xcc, 0x90, 0x5c, 0xe3, 0xe4, 0xd0, 0xa4, 0x60, 0x0e, 0x53, 0x1a, 0x93, 0x83, 0xe4,
	0xd0, 0x10, 0x2d, 0xe0, 0xf6, 0xbf, 0x0d, 0xe5, 0x75, 0x46, 0x8e, 0x78, 0xfe, 0xb9, 0x1f, 0xf7,
	0xe4, 0xf9, 0xef, 0x03, 0x10, 0xfb, 0xa5, 0xc4, 0x63, 0x38, 0x1a, 0x37, 0x4a, 0x8b, 0xcb, 0x0b,
	0xd6, 0x67, 0xe3, 0xb4, 0x97, 0x49, 0xdc, 0x87, 0xa9, 0x75, 0x14, 0x8d, 0x2b, 0x50, 0x5e, 0x32,
	0xb0, 0xb9, 0x0b, 0x09, 0x93, 0x0e, 0x03, 0x59, 0xef, 0xe4, 0xf2, 0x2c, 0x5b, 0xcc, 0x56, 0xe3,
	0x55, 0xbc, 0x85, 0x07, 0x01, 0xd7, 0x96, 0x62, 0xb0, 0x79, 0xdb, 0xff, 0x34, 0xba, 0xa8, 0x2c,
	0xba, 0xa8, 0x0b, 0xdf, 0xba, 0x28, 0x9f, 0x32, 0x3d, 0xdf, 0x45, 0xae, 0x98, 0xb8, 0xea, 0x01,
	0xb1, 0x84, 0x7b, 0x9b, 0x72, 0xc9, 0xa4, 0x41, 0x7d, 0xc6, 0xea, 0xc9, 0x98, 0x26, 0x89, 0x92,
	0xf5, 0x08, 0x66, 0xbc, 0xfb, 0xc0, 0x92, 0x7c, 0xce, 0xf0, 0xad, 0x34, 0x64, 0x86, 0x84, 0x82,
	0x63, 0x53, 0x44, 0xdb, 0xa1, 0xde, 0x66, 0xe5, 0x0b, 0x46, 0xef, 0xa4, 0x68, 0xf5, 0x57, 0x10,
	0x4f, 0xe0, 0x64, 0x8d, 0x15, 0x46, 0x5c, 0xef, 0x0a, 0x2f, 0x59, 0x10, 0x43, 0xb6, 0x6d, 0x3c,
	0x84, 0x9b, 0x8d, 0x89, 0x5f, 0x47, 0xf0, 0x15, 0x83, 0xc0, 0xa3, 0x04, 0x5c, 0xc0, 0xdd, 0xcd,
	0x76, 0xe9, 0x4f, 0xac, 0xb1, 0xf6, 0xe1, 0x87, 0x7c, 0xcd, 0xe8, 0xf1, 0x18, 0xf2, 0xa3, 0x1f,
	0x38, 0xea, 0x6b, 0xe7, 0xa6, 0xf8, 0xd6, 0x36, 0xbb, 0x2d, 0xde, 0xa4, 0xda, 0x29, 0xda, 0x2a,
	0x91, 0x1f, 0xf2, 0x23, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x59, 0xde, 0x6f, 0xcc, 0x3e,
	0x03, 0x00, 0x00,
}
