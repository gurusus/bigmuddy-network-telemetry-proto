// Code generated by protoc-gen-go.
// source: bgp_rpki_routes_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_rpki_routes_rpki_route is a generated protocol buffer package.

It is generated from these files:
	bgp_rpki_routes_bag.proto

It has these top-level messages:
	BgpRpkiRoutesBag_KEYS
	BgpRpkiRoutesBag
	BgpEdmRpkiRouteType_
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_rpki_routes_rpki_route

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

type BgpRpkiRoutesBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName       string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	Address      string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Minimum      uint32 `protobuf:"varint,4,opt,name=minimum" json:"minimum,omitempty"`
	Maximum      uint32 `protobuf:"varint,5,opt,name=maximum" json:"maximum,omitempty"`
}

func (m *BgpRpkiRoutesBag_KEYS) Reset()                    { *m = BgpRpkiRoutesBag_KEYS{} }
func (m *BgpRpkiRoutesBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpRpkiRoutesBag_KEYS) ProtoMessage()               {}
func (*BgpRpkiRoutesBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpRpkiRoutesBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpRpkiRoutesBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpRpkiRoutesBag_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BgpRpkiRoutesBag_KEYS) GetMinimum() uint32 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

func (m *BgpRpkiRoutesBag_KEYS) GetMaximum() uint32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

type BgpRpkiRoutesBag struct {
	// Array or RPKI routes
	RpkiRoutes []*BgpEdmRpkiRouteType_ `protobuf:"bytes,50,rep,name=rpki_routes,json=rpkiRoutes" json:"rpki_routes,omitempty"`
}

func (m *BgpRpkiRoutesBag) Reset()                    { *m = BgpRpkiRoutesBag{} }
func (m *BgpRpkiRoutesBag) String() string            { return proto.CompactTextString(m) }
func (*BgpRpkiRoutesBag) ProtoMessage()               {}
func (*BgpRpkiRoutesBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpRpkiRoutesBag) GetRpkiRoutes() []*BgpEdmRpkiRouteType_ {
	if m != nil {
		return m.RpkiRoutes
	}
	return nil
}

type BgpEdmRpkiRouteType_ struct {
	// Address Family
	AfName string `protobuf:"bytes,1,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	// Address Prefix
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	// Minimum Prefix Length
	MinPrefixLen uint32 `protobuf:"varint,3,opt,name=min_prefix_len,json=minPrefixLen" json:"min_prefix_len,omitempty"`
	// Maximum Prefix Length
	MaxPrefixLen uint32 `protobuf:"varint,4,opt,name=max_prefix_len,json=maxPrefixLen" json:"max_prefix_len,omitempty"`
	// AS number
	As uint32 `protobuf:"varint,5,opt,name=as" json:"as,omitempty"`
	// Refcount
	Refcount uint32 `protobuf:"varint,6,opt,name=refcount" json:"refcount,omitempty"`
	// Source Server
	Server string `protobuf:"bytes,7,opt,name=server" json:"server,omitempty"`
	// ROA is stale
	Stale bool `protobuf:"varint,8,opt,name=stale" json:"stale,omitempty"`
}

func (m *BgpEdmRpkiRouteType_) Reset()                    { *m = BgpEdmRpkiRouteType_{} }
func (m *BgpEdmRpkiRouteType_) String() string            { return proto.CompactTextString(m) }
func (*BgpEdmRpkiRouteType_) ProtoMessage()               {}
func (*BgpEdmRpkiRouteType_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BgpEdmRpkiRouteType_) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpEdmRpkiRouteType_) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BgpEdmRpkiRouteType_) GetMinPrefixLen() uint32 {
	if m != nil {
		return m.MinPrefixLen
	}
	return 0
}

func (m *BgpEdmRpkiRouteType_) GetMaxPrefixLen() uint32 {
	if m != nil {
		return m.MaxPrefixLen
	}
	return 0
}

func (m *BgpEdmRpkiRouteType_) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

func (m *BgpEdmRpkiRouteType_) GetRefcount() uint32 {
	if m != nil {
		return m.Refcount
	}
	return 0
}

func (m *BgpEdmRpkiRouteType_) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *BgpEdmRpkiRouteType_) GetStale() bool {
	if m != nil {
		return m.Stale
	}
	return false
}

func init() {
	proto.RegisterType((*BgpRpkiRoutesBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.rpki_routes.rpki_route.bgp_rpki_routes_bag_KEYS")
	proto.RegisterType((*BgpRpkiRoutesBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.rpki_routes.rpki_route.bgp_rpki_routes_bag")
	proto.RegisterType((*BgpEdmRpkiRouteType_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.rpki_routes.rpki_route.bgp_edm_rpki_route_type_")
}

func init() { proto.RegisterFile("bgp_rpki_routes_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcd, 0x4e, 0xbb, 0x40,
	0x14, 0xc5, 0x33, 0xb4, 0xa5, 0xfd, 0x4f, 0x3f, 0x16, 0xfc, 0x8d, 0x8e, 0xae, 0x48, 0x75, 0xc1,
	0x8a, 0x45, 0xf5, 0x15, 0x5c, 0x69, 0x8c, 0xc1, 0x55, 0x57, 0x93, 0x01, 0x2e, 0xcd, 0xc4, 0xce,
	0x47, 0x66, 0x68, 0x43, 0xdf, 0xc3, 0x67, 0x70, 0xe5, 0xdb, 0xf9, 0x02, 0x86, 0x81, 0x22, 0x26,
	0x75, 0xe9, 0x8a, 0x7b, 0xce, 0xef, 0x70, 0xb9, 0x39, 0x01, 0x5f, 0xa6, 0x1b, 0x4d, 0x8d, 0x7e,
	0xe5, 0xd4, 0xa8, 0x5d, 0x09, 0x96, 0xa6, 0x6c, 0x13, 0x6b, 0xa3, 0x4a, 0x15, 0xac, 0x33, 0x6e,
	0x33, 0x45, 0xb9, 0xb2, 0xb4, 0x32, 0x94, 0xeb, 0xfd, 0x1d, 0xad, 0xc3, 0x4a, 0x83, 0x89, 0xd3,
	0x8d, 0x8e, 0xb9, 0xb4, 0x25, 0x93, 0x19, 0xd8, 0x6e, 0xea, 0x06, 0x5a, 0x3f, 0xf2, 0xf4, 0x10,
	0xf7, 0x36, 0xf7, 0xe6, 0xe5, 0x3b, 0xc2, 0xe4, 0xc4, 0x87, 0xe9, 0xc3, 0xfd, 0xfa, 0x25, 0xb8,
	0xc6, 0xf3, 0x6e, 0x8f, 0x64, 0x02, 0x08, 0x0a, 0x51, 0xf4, 0x2f, 0x99, 0x1d, 0xcd, 0x27, 0x26,
	0x20, 0xb8, 0xc0, 0x63, 0x56, 0x34, 0xd8, 0x73, 0xd8, 0x67, 0x85, 0x03, 0x04, 0x8f, 0x59, 0x9e,
	0x1b, 0xb0, 0x96, 0x0c, 0x1c, 0x38, 0xca, 0x9a, 0x08, 0x2e, 0xb9, 0xd8, 0x09, 0x32, 0x0c, 0x51,
	0x34, 0x4f, 0x8e, 0xd2, 0x11, 0x56, 0x39, 0x32, 0x6a, 0x49, 0x23, 0x97, 0x1f, 0x08, 0xff, 0x3f,
	0x71, 0x68, 0xf0, 0x86, 0xf0, 0xb4, 0xe7, 0x91, 0x55, 0x38, 0x88, 0xa6, 0x2b, 0x1b, 0xff, 0x59,
	0x65, 0xf5, 0xeb, 0x14, 0x72, 0xd1, 0xbb, 0x84, 0x96, 0x07, 0x0d, 0x34, 0xc1, 0xb5, 0x93, 0xb8,
	0xfc, 0xf2, 0xb3, 0xed, 0xf5, 0x54, 0xb0, 0x5f, 0x19, 0xfa, 0xad, 0x32, 0xef, 0x67, 0x65, 0x37,
	0x78, 0x21, 0xb8, 0xa4, 0xda, 0x40, 0xc1, 0x2b, 0xba, 0x05, 0xe9, 0x3a, 0x9d, 0x27, 0x33, 0xc1,
	0xe5, 0xb3, 0x33, 0x1f, 0x41, 0xba, 0x14, 0xab, 0xfa, 0xa9, 0x61, 0x9b, 0x62, 0xd5, 0x77, 0x6a,
	0x81, 0x3d, 0x66, 0xdb, 0x7e, 0x3d, 0x66, 0x83, 0x2b, 0x3c, 0x31, 0x50, 0x64, 0x6a, 0x27, 0x4b,
	0xe2, 0x3b, 0xb7, 0xd3, 0xc1, 0x39, 0xf6, 0x2d, 0x98, 0x3d, 0x18, 0x32, 0x6e, 0x2e, 0x6d, 0x54,
	0x70, 0x86, 0x47, 0xb6, 0x64, 0x5b, 0x20, 0x93, 0x10, 0x45, 0x93, 0xa4, 0x11, 0xa9, 0xef, 0xfe,
	0xd7, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x0f, 0xdc, 0x34, 0xcc, 0x02, 0x00, 0x00,
}