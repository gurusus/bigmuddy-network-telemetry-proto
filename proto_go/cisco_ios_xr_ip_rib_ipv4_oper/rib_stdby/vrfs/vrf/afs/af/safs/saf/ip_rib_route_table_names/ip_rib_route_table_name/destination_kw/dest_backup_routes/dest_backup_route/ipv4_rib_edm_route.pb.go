// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_route.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_destination_kw_dest_backup_routes_dest_backup_route is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_route.proto

It has these top-level messages:
	Ipv4RibEdmRoute_KEYS
	Ipv4RibEdmRoute
	Ipv4RibEdmPath
	Ipv4RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_destination_kw_dest_backup_routes_dest_backup_route

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

// Information of a rib route head and rib proto route
type Ipv4RibEdmRoute_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *Ipv4RibEdmRoute_KEYS) Reset()                    { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv4RibEdmRoute struct {
	// Route prefix
	Prefix string `protobuf:"bytes,50,opt,name=prefix" json:"prefix,omitempty"`
	// Length of prefix
	PrefixLength uint32 `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	// Route version
	RouteVersion uint32 `protobuf:"varint,52,opt,name=route_version,json=routeVersion" json:"route_version,omitempty"`
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,53,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	//  Name of Protocol
	ProtocolName string `protobuf:"bytes,54,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Instance name
	Instance string `protobuf:"bytes,55,opt,name=instance" json:"instance,omitempty"`
	// Client adding the route to RIB
	ClientId uint32 `protobuf:"varint,56,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Route type
	RouteType uint32 `protobuf:"varint,57,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// Route priority
	Priority uint32 `protobuf:"varint,58,opt,name=priority" json:"priority,omitempty"`
	// SVD Type of route
	SvdType uint32 `protobuf:"varint,59,opt,name=svd_type,json=svdType" json:"svd_type,omitempty"`
	// Route flags
	Flags uint32 `protobuf:"varint,60,opt,name=flags" json:"flags,omitempty"`
	// Extended Route flags
	ExtendedFlags uint64 `protobuf:"varint,61,opt,name=extended_flags,json=extendedFlags" json:"extended_flags,omitempty"`
	// Opaque proto specific info
	Tag uint32 `protobuf:"varint,62,opt,name=tag" json:"tag,omitempty"`
	// Distance of the route
	Distance uint32 `protobuf:"varint,63,opt,name=distance" json:"distance,omitempty"`
	// Diversion distance of the route
	DiversionDistance uint32 `protobuf:"varint,64,opt,name=diversion_distance,json=diversionDistance" json:"diversion_distance,omitempty"`
	// Route metric
	Metric uint32 `protobuf:"varint,65,opt,name=metric" json:"metric,omitempty"`
	// Number of paths
	PathsCount uint32 `protobuf:"varint,66,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// BGP Attribute ID
	AttributeIdentity uint32 `protobuf:"varint,67,opt,name=attribute_identity,json=attributeIdentity" json:"attribute_identity,omitempty"`
	// BGP Traffic Index
	TrafficIndex uint32 `protobuf:"varint,68,opt,name=traffic_index,json=trafficIndex" json:"traffic_index,omitempty"`
	// Route ip precedence
	RoutePrecedence uint32 `protobuf:"varint,69,opt,name=route_precedence,json=routePrecedence" json:"route_precedence,omitempty"`
	// Route qos group
	QosGroup uint32 `protobuf:"varint,70,opt,name=qos_group,json=qosGroup" json:"qos_group,omitempty"`
	// Flow tag
	FlowTag uint32 `protobuf:"varint,71,opt,name=flow_tag,json=flowTag" json:"flow_tag,omitempty"`
	// Forward Class
	FwdClass uint32 `protobuf:"varint,72,opt,name=fwd_class,json=fwdClass" json:"fwd_class,omitempty"`
	// Number of pic paths in this route
	PicCount uint32 `protobuf:"varint,73,opt,name=pic_count,json=picCount" json:"pic_count,omitempty"`
	// Is the route active or backup
	Active bool `protobuf:"varint,74,opt,name=active" json:"active,omitempty"`
	// Route has a diversion path
	Diversion bool `protobuf:"varint,75,opt,name=diversion" json:"diversion,omitempty"`
	// Diversion route protocol name
	DiversionProtoName string `protobuf:"bytes,76,opt,name=diversion_proto_name,json=diversionProtoName" json:"diversion_proto_name,omitempty"`
	// Age of route (seconds)
	RouteAge uint32 `protobuf:"varint,77,opt,name=route_age,json=routeAge" json:"route_age,omitempty"`
	// Local label of the route
	RouteLabel uint32 `protobuf:"varint,78,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Route Version
	Version uint32 `protobuf:"varint,79,opt,name=version" json:"version,omitempty"`
	// Table Version
	TblVersion uint64 `protobuf:"varint,80,opt,name=tbl_version,json=tblVersion" json:"tbl_version,omitempty"`
	// Route modification time(nanoseconds)
	RouteModifyTime uint64 `protobuf:"varint,81,opt,name=route_modify_time,json=routeModifyTime" json:"route_modify_time,omitempty"`
	// Path(s) of the route
	RoutePath *Ipv4RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath" json:"route_path,omitempty"`
}

func (m *Ipv4RibEdmRoute) Reset()                    { *m = Ipv4RibEdmRoute{} }
func (m *Ipv4RibEdmRoute) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute) ProtoMessage()               {}
func (*Ipv4RibEdmRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePath() *Ipv4RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

// Information of a rib path
type Ipv4RibEdmPath struct {
	// Next path
	Ipv4RibEdmPath []*Ipv4RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv4_rib_edm_path,json=ipv4RibEdmPath" json:"ipv4_rib_edm_path,omitempty"`
}

func (m *Ipv4RibEdmPath) Reset()                    { *m = Ipv4RibEdmPath{} }
func (m *Ipv4RibEdmPath) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPath) ProtoMessage()               {}
func (*Ipv4RibEdmPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv4RibEdmPath) GetIpv4RibEdmPath() []*Ipv4RibEdmPathItem {
	if m != nil {
		return m.Ipv4RibEdmPath
	}
	return nil
}

type Ipv4RibEdmPathItem struct {
	// Nexthop
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Infosource
	InformationSource string `protobuf:"bytes,2,opt,name=information_source,json=informationSource" json:"information_source,omitempty"`
	// V6 nexthop
	V6Nexthop string `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop" json:"v6_nexthop,omitempty"`
	// Interface Name
	InterfaceName string `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Metrics
	Metric uint32 `protobuf:"varint,5,opt,name=metric" json:"metric,omitempty"`
	// Load Metrics
	LoadMetric uint32 `protobuf:"varint,6,opt,name=load_metric,json=loadMetric" json:"load_metric,omitempty"`
	// Flags extended to 64 bits
	Flags64 uint64 `protobuf:"varint,7,opt,name=flags64" json:"flags64,omitempty"`
	// Flags
	Flags uint32 `protobuf:"varint,8,opt,name=flags" json:"flags,omitempty"`
	// Private Flags
	PrivateFlags uint32 `protobuf:"varint,9,opt,name=private_flags,json=privateFlags" json:"private_flags,omitempty"`
	// Looping path
	Looped bool `protobuf:"varint,10,opt,name=looped" json:"looped,omitempty"`
	// Nexthop tableid
	NextHopTableId uint32 `protobuf:"varint,11,opt,name=next_hop_table_id,json=nextHopTableId" json:"next_hop_table_id,omitempty"`
	// VRF Name of the nh table
	NextHopVrfName string `protobuf:"bytes,12,opt,name=next_hop_vrf_name,json=nextHopVrfName" json:"next_hop_vrf_name,omitempty"`
	// NH table name
	NextHopTableName string `protobuf:"bytes,13,opt,name=next_hop_table_name,json=nextHopTableName" json:"next_hop_table_name,omitempty"`
	// NH afi
	NextHopAfi uint32 `protobuf:"varint,14,opt,name=next_hop_afi,json=nextHopAfi" json:"next_hop_afi,omitempty"`
	// NH safi
	NextHopSafi uint32 `protobuf:"varint,15,opt,name=next_hop_safi,json=nextHopSafi" json:"next_hop_safi,omitempty"`
	// Label associated with this path
	RouteLabel uint32 `protobuf:"varint,16,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Tunnel ID associated with this path
	TunnelId uint32 `protobuf:"varint,17,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// Path id of this path
	Pathid uint32 `protobuf:"varint,18,opt,name=pathid" json:"pathid,omitempty"`
	// Path id of this path's backup
	BackupPathid uint32 `protobuf:"varint,19,opt,name=backup_pathid,json=backupPathid" json:"backup_pathid,omitempty"`
	// Refcnt of backup
	RefCntOfBackup uint32 `protobuf:"varint,20,opt,name=ref_cnt_of_backup,json=refCntOfBackup" json:"ref_cnt_of_backup,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,21,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities" json:"number_of_extended_communities,omitempty"`
	// MVPN attribute present
	MvpnPresent bool `protobuf:"varint,22,opt,name=mvpn_present,json=mvpnPresent" json:"mvpn_present,omitempty"`
	// Path RT present
	PathrtPresent           bool `protobuf:"varint,23,opt,name=pathrt_present,json=pathrtPresent" json:"pathrt_present,omitempty"`
	VrfimportrtPresent      bool `protobuf:"varint,24,opt,name=vrfimportrt_present,json=vrfimportrtPresent" json:"vrfimportrt_present,omitempty"`
	SourceasrtPresent       bool `protobuf:"varint,25,opt,name=sourceasrt_present,json=sourceasrtPresent" json:"sourceasrt_present,omitempty"`
	SourcerdPresent         bool `protobuf:"varint,26,opt,name=sourcerd_present,json=sourcerdPresent" json:"sourcerd_present,omitempty"`
	SegmentedNexthopPresent bool `protobuf:"varint,27,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent" json:"segmented_nexthop_present,omitempty"`
	// NHID associated with this path
	NextHopId uint32 `protobuf:"varint,28,opt,name=next_hop_id,json=nextHopId" json:"next_hop_id,omitempty"`
	// NHID references
	NextHopIdRefcount uint32 `protobuf:"varint,29,opt,name=next_hop_id_refcount,json=nextHopIdRefcount" json:"next_hop_id_refcount,omitempty"`
	// OSPF area associated with the path
	OspfAreaId string `protobuf:"bytes,30,opt,name=ospf_area_id,json=ospfAreaId" json:"ospf_area_id,omitempty"`
	// Remote backup node address
	RemoteBackupAddr [][]byte `protobuf:"bytes,31,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
	// Path has a label stack
	HasLabelstk bool `protobuf:"varint,32,opt,name=has_labelstk,json=hasLabelstk" json:"has_labelstk,omitempty"`
	// Number of labels in stack
	NumLabels uint32 `protobuf:"varint,33,opt,name=num_labels,json=numLabels" json:"num_labels,omitempty"`
	// Labels for this path
	Labelstk []uint32 `protobuf:"varint,34,rep,packed,name=labelstk" json:"labelstk,omitempty"`
	// binding Label for this path
	BindingLabel uint32 `protobuf:"varint,35,opt,name=binding_label,json=bindingLabel" json:"binding_label,omitempty"`
	// Fib nhid encap id
	NhidFeid uint64 `protobuf:"varint,36,opt,name=nhid_feid,json=nhidFeid" json:"nhid_feid,omitempty"`
	// Fib mpls encap id
	MplsFeid uint64 `protobuf:"varint,37,opt,name=mpls_feid,json=mplsFeid" json:"mpls_feid,omitempty"`
}

func (m *Ipv4RibEdmPathItem) Reset()                    { *m = Ipv4RibEdmPathItem{} }
func (m *Ipv4RibEdmPathItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPathItem) ProtoMessage()               {}
func (*Ipv4RibEdmPathItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv4RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetRemoteBackupAddr() [][]byte {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

// Information of local label for route head
type RibEdmLocalLabel struct {
	// Protocol Name
	ProtocolName string `protobuf:"bytes,1,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Client ID
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Stale
	Stale uint32 `protobuf:"varint,3,opt,name=stale" json:"stale,omitempty"`
	// Mirrored
	Mirrored uint32 `protobuf:"varint,4,opt,name=mirrored" json:"mirrored,omitempty"`
	// Merge disable
	MergeDisable uint32 `protobuf:"varint,5,opt,name=merge_disable,json=mergeDisable" json:"merge_disable,omitempty"`
	// Redist only
	RedistOnly uint32 `protobuf:"varint,6,opt,name=redist_only,json=redistOnly" json:"redist_only,omitempty"`
	// Local label
	Label uint32 `protobuf:"varint,7,opt,name=label" json:"label,omitempty"`
	// Distance
	Distance uint32 `protobuf:"varint,8,opt,name=distance" json:"distance,omitempty"`
}

func (m *RibEdmLocalLabel) Reset()                    { *m = RibEdmLocalLabel{} }
func (m *RibEdmLocalLabel) String() string            { return proto.CompactTextString(m) }
func (*RibEdmLocalLabel) ProtoMessage()               {}
func (*RibEdmLocalLabel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RibEdmLocalLabel) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *RibEdmLocalLabel) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RibEdmLocalLabel) GetStale() uint32 {
	if m != nil {
		return m.Stale
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMirrored() uint32 {
	if m != nil {
		return m.Mirrored
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMergeDisable() uint32 {
	if m != nil {
		return m.MergeDisable
	}
	return 0
}

func (m *RibEdmLocalLabel) GetRedistOnly() uint32 {
	if m != nil {
		return m.RedistOnly
	}
	return 0
}

func (m *RibEdmLocalLabel) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *RibEdmLocalLabel) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_backup_routes.dest_backup_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_backup_routes.dest_backup_route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_backup_routes.dest_backup_route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_backup_routes.dest_backup_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_backup_routes.dest_backup_route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x52, 0x1b, 0xbf,
	0x15, 0x1f, 0x87, 0x3f, 0x60, 0x0b, 0x4c, 0xf0, 0x42, 0x41, 0x84, 0x7c, 0x38, 0xa4, 0x99, 0x31,
	0x9d, 0xe2, 0x74, 0x92, 0x94, 0xb6, 0xe9, 0x27, 0x21, 0x24, 0x71, 0x43, 0x02, 0x75, 0x32, 0x99,
	0xe9, 0x95, 0x46, 0x5e, 0x69, 0x6d, 0x4d, 0x76, 0x57, 0x1b, 0x49, 0x36, 0x70, 0xd9, 0x87, 0xe9,
	0x3b, 0xf4, 0xa6, 0x0f, 0xd1, 0xbb, 0xbe, 0x42, 0x6f, 0x7b, 0xd3, 0xcb, 0x76, 0x74, 0x8e, 0x76,
	0x6d, 0x70, 0xfa, 0x00, 0xb9, 0x01, 0xce, 0xef, 0xf7, 0x93, 0x56, 0xe7, 0x43, 0x47, 0x07, 0x42,
	0x55, 0x31, 0x79, 0xce, 0x8c, 0x1a, 0x30, 0x29, 0x32, 0x66, 0xf4, 0xd8, 0xc9, 0x6e, 0x61, 0xb4,
	0xd3, 0xd1, 0x5f, 0x6b, 0xb1, 0xb2, 0xb1, 0x66, 0x4a, 0x5b, 0x76, 0x69, 0x98, 0x2a, 0x40, 0x05,
	0x72, 0x5d, 0x48, 0xd3, 0xf5, 0x96, 0x75, 0x62, 0x70, 0xd5, 0x9d, 0x98, 0xc4, 0xfa, 0x1f, 0x5d,
	0x9e, 0xd8, 0x2e, 0x4f, 0xba, 0xd6, 0xff, 0xb6, 0x3c, 0xe9, 0x86, 0x35, 0xb0, 0x2b, 0x73, 0x7c,
	0x90, 0x4a, 0x96, 0xf3, 0x4c, 0xda, 0xff, 0x47, 0x74, 0x85, 0xb4, 0x4e, 0xe5, 0xdc, 0x29, 0x9d,
	0xb3, 0x2f, 0x17, 0x60, 0xb2, 0x01, 0x8f, 0xbf, 0x8c, 0x0b, 0xd4, 0xda, 0x79, 0x68, 0xef, 0x1f,
	0x35, 0xb2, 0x3d, 0xef, 0x04, 0x7b, 0x77, 0xf2, 0xe7, 0x8f, 0xd1, 0x0e, 0xa9, 0x4f, 0x4c, 0x02,
	0x7b, 0xd3, 0x5a, 0xbb, 0xd6, 0x69, 0xf4, 0x97, 0x27, 0x26, 0xf9, 0xc0, 0x33, 0x19, 0x6d, 0x93,
	0x65, 0x1e, 0x98, 0x5b, 0xc0, 0x2c, 0x71, 0x24, 0x76, 0x48, 0xdd, 0x96, 0xcc, 0x02, 0xae, 0xb1,
	0x81, 0xea, 0x90, 0xf5, 0x9b, 0x47, 0xa6, 0x3f, 0x80, 0x64, 0x0d, 0xf0, 0x4f, 0x1e, 0x06, 0x25,
	0x25, 0xcb, 0x5c, 0x08, 0x23, 0xad, 0xa5, 0x8b, 0xb8, 0x47, 0x30, 0xa3, 0x47, 0xa4, 0x59, 0x18,
	0x99, 0xa8, 0x4b, 0x96, 0xca, 0x7c, 0xe8, 0x46, 0x74, 0xa9, 0x5d, 0xeb, 0x34, 0xfb, 0xab, 0x08,
	0x9e, 0x02, 0xb6, 0xf7, 0x9f, 0x06, 0x89, 0xe6, 0x7d, 0x8a, 0xb6, 0xc8, 0x12, 0xca, 0xe8, 0x53,
	0x3c, 0x32, 0x5a, 0xf3, 0x7b, 0x3e, 0x9b, 0xdf, 0xd3, 0x8b, 0xf0, 0xf0, 0x13, 0x69, 0xac, 0xd2,
	0x39, 0x7d, 0x8e, 0x22, 0x00, 0x3f, 0x23, 0x16, 0x3d, 0x20, 0x2b, 0x90, 0xfd, 0x58, 0xa7, 0x4c,
	0x09, 0xfa, 0x73, 0x90, 0x90, 0x12, 0xea, 0x09, 0xfc, 0x54, 0x10, 0x80, 0xff, 0x87, 0x70, 0x92,
	0xd5, 0x12, 0x04, 0xef, 0xef, 0x90, 0xba, 0xca, 0xad, 0xe3, 0x79, 0x2c, 0xe9, 0x2f, 0x80, 0xaf,
	0xec, 0x68, 0x97, 0x34, 0xe2, 0x54, 0xc9, 0xdc, 0xf9, 0xfd, 0x7f, 0x09, 0xfb, 0xd7, 0x11, 0xe8,
	0x89, 0xe8, 0x1e, 0x21, 0x21, 0xc0, 0x57, 0x85, 0xa4, 0xbf, 0x02, 0xb6, 0x81, 0xa1, 0xbd, 0x2a,
	0x60, 0xdf, 0xc2, 0x28, 0x6d, 0x94, 0xbb, 0xa2, 0x2f, 0x70, 0x69, 0x69, 0x43, 0xda, 0x26, 0x02,
	0x17, 0xfe, 0x1a, 0xb8, 0x65, 0x3b, 0x11, 0xb0, 0x6c, 0x93, 0x2c, 0x26, 0x29, 0x1f, 0x5a, 0xfa,
	0x1b, 0xc0, 0xd1, 0x88, 0x1e, 0x93, 0x35, 0x79, 0xe9, 0x64, 0x2e, 0xa4, 0x60, 0x48, 0xff, 0xb6,
	0x5d, 0xeb, 0xfc, 0xd0, 0x6f, 0x96, 0xe8, 0x6b, 0x90, 0xad, 0x93, 0x05, 0xc7, 0x87, 0xf4, 0x77,
	0xb0, 0xd4, 0xff, 0xe9, 0x4f, 0x21, 0x54, 0xf0, 0xee, 0xf7, 0x78, 0x8a, 0xd2, 0x8e, 0x0e, 0x48,
	0x24, 0x54, 0x08, 0x30, 0xab, 0x54, 0x7f, 0x00, 0x55, 0xab, 0x62, 0x5e, 0x95, 0xf2, 0x2d, 0xb2,
	0x94, 0x49, 0x67, 0x54, 0x4c, 0x8f, 0x40, 0x12, 0x2c, 0x48, 0x03, 0x77, 0x23, 0xcb, 0x62, 0x3d,
	0xce, 0x1d, 0x7d, 0x19, 0xd2, 0xe0, 0xa1, 0x63, 0x8f, 0xf8, 0xef, 0x70, 0xe7, 0x8c, 0x1a, 0xf8,
	0x60, 0x29, 0x21, 0x73, 0xe7, 0x63, 0x72, 0x8c, 0xdf, 0xa9, 0x98, 0x5e, 0x20, 0x7c, 0xd6, 0x9c,
	0xe1, 0x49, 0xa2, 0x62, 0xa6, 0x72, 0x21, 0x2f, 0xe9, 0x2b, 0xcc, 0x7d, 0x00, 0x7b, 0x1e, 0x8b,
	0xf6, 0xcb, 0xea, 0x2e, 0x8c, 0x8c, 0xa5, 0x90, 0xfe, 0xe4, 0x27, 0xa0, 0xbb, 0x0d, 0xf8, 0x79,
	0x05, 0xfb, 0x24, 0x7e, 0xd5, 0x96, 0x0d, 0x8d, 0x1e, 0x17, 0xf4, 0x35, 0xc6, 0xe0, 0xab, 0xb6,
	0x6f, 0xbc, 0xed, 0x33, 0x91, 0xa4, 0xfa, 0x82, 0xf9, 0xb0, 0xbd, 0xc1, 0x4c, 0x78, 0xfb, 0x13,
	0x1f, 0xfa, 0x75, 0xc9, 0x85, 0x60, 0x71, 0xca, 0xad, 0xa5, 0x6f, 0x71, 0x5d, 0x72, 0x21, 0x8e,
	0xbd, 0xed, 0xc9, 0x42, 0xc5, 0xc1, 0xe5, 0x5e, 0x48, 0xaf, 0x8a, 0xd1, 0xe1, 0x2d, 0xb2, 0xc4,
	0x63, 0xa7, 0x26, 0x92, 0xfe, 0xb1, 0x5d, 0xeb, 0xd4, 0xfb, 0xc1, 0x8a, 0xee, 0x92, 0x46, 0x15,
	0x56, 0xfa, 0x0e, 0xa8, 0x29, 0x10, 0xfd, 0x8c, 0x6c, 0x4e, 0xd3, 0x01, 0x25, 0x8a, 0x45, 0x7b,
	0x0a, 0x45, 0x39, 0x4d, 0xd5, 0xb9, 0xa7, 0xa0, 0x74, 0x77, 0x09, 0xd6, 0x1b, 0xe3, 0x43, 0x49,
	0xdf, 0xe3, 0x21, 0x00, 0x38, 0x1a, 0x4a, 0x9f, 0x16, 0x24, 0x53, 0x3e, 0x90, 0x29, 0xfd, 0x80,
	0x69, 0x01, 0xe8, 0xd4, 0x23, 0xfe, 0xda, 0x97, 0x67, 0x39, 0x43, 0xcf, 0x27, 0xd3, 0x8b, 0xe5,
	0x06, 0x69, 0x75, 0xf7, 0xce, 0xa1, 0xd4, 0x88, 0x1b, 0xa4, 0xe5, 0xcd, 0xfb, 0x09, 0x69, 0xe1,
	0xde, 0x99, 0x16, 0x2a, 0xb9, 0x62, 0x4e, 0x65, 0x92, 0xfe, 0x09, 0x64, 0x18, 0xfe, 0xf7, 0x80,
	0x7f, 0x52, 0x99, 0x8c, 0xfe, 0x59, 0x2b, 0xef, 0x89, 0x2f, 0x09, 0xda, 0x6f, 0xd7, 0x3a, 0x2b,
	0x4f, 0xff, 0x56, 0xeb, 0x7e, 0x17, 0x0d, 0xbb, 0x7b, 0xad, 0xb1, 0x79, 0x07, 0xc2, 0x15, 0x3f,
	0xe7, 0x6e, 0xb4, 0xf7, 0xdf, 0x1a, 0x69, 0xcd, 0x09, 0xa2, 0x7f, 0x7d, 0x0b, 0xa5, 0xb5, 0xf6,
	0x42, 0x67, 0xe5, 0xe9, 0xdf, 0xbf, 0x5b, 0xbf, 0x99, 0x72, 0x32, 0xeb, 0xaf, 0x79, 0xbc, 0xaf,
	0x06, 0x27, 0x22, 0x83, 0x08, 0xfc, 0x9b, 0x90, 0xad, 0x6f, 0x4b, 0x67, 0x5f, 0x95, 0xda, 0xf5,
	0x57, 0xe5, 0x80, 0x44, 0x2a, 0x4f, 0xb4, 0xc9, 0xf0, 0x48, 0x56, 0x8f, 0x4d, 0x5c, 0x3e, 0x6c,
	0xad, 0x19, 0xe6, 0x23, 0x10, 0xbe, 0xcf, 0x4e, 0x0e, 0x59, 0x2e, 0x2f, 0xdd, 0x48, 0x17, 0xe1,
	0x95, 0x6b, 0x4c, 0x0e, 0x3f, 0x20, 0xe0, 0x5b, 0xa3, 0xca, 0x9d, 0x34, 0x09, 0x8f, 0xaf, 0xbd,
	0x72, 0xcd, 0x0a, 0x85, 0xbb, 0x32, 0xed, 0x5e, 0x8b, 0x37, 0xbb, 0x57, 0xaa, 0xb9, 0x60, 0x81,
	0xc4, 0x07, 0x8e, 0x78, 0xe8, 0x3d, 0x0a, 0x28, 0x59, 0x86, 0x8e, 0x7b, 0xf8, 0x9c, 0x2e, 0x43,
	0x85, 0x97, 0xe6, 0xb4, 0x55, 0xd7, 0x67, 0x5b, 0x35, 0x3c, 0x3a, 0x6a, 0xc2, 0x9d, 0x0c, 0x9d,
	0xba, 0x51, 0xbe, 0x6f, 0x00, 0x62, 0xa3, 0xde, 0x22, 0x4b, 0xa9, 0xd6, 0x85, 0x14, 0x94, 0x60,
	0x87, 0x40, 0x2b, 0xda, 0x27, 0x2d, 0xef, 0x28, 0x1b, 0xe9, 0x22, 0x64, 0x50, 0x09, 0xba, 0x02,
	0x1b, 0xac, 0x79, 0xe2, 0xad, 0x2e, 0xe0, 0xdd, 0xee, 0x5d, 0x97, 0x56, 0x73, 0xc3, 0x2a, 0x3e,
	0xf0, 0x41, 0xfa, 0x39, 0x8c, 0x0f, 0x07, 0x64, 0xe3, 0xc6, 0xae, 0x20, 0x6e, 0x82, 0x78, 0x7d,
	0x76, 0x5f, 0x90, 0xb7, 0xc9, 0x6a, 0x25, 0xe7, 0x89, 0xa2, 0x6b, 0x18, 0x93, 0xa0, 0x3b, 0x4a,
	0x54, 0xb4, 0x47, 0x9a, 0x95, 0xc2, 0x7a, 0xc9, 0x6d, 0x90, 0xac, 0x04, 0xc9, 0x47, 0x9e, 0xa8,
	0x9b, 0xfd, 0x67, 0x7d, 0xae, 0xff, 0xec, 0x92, 0x86, 0x1b, 0xe7, 0xb9, 0x84, 0xc7, 0xbb, 0x85,
	0xdd, 0x0b, 0x81, 0x9e, 0x80, 0xe9, 0x81, 0xbb, 0x91, 0x12, 0x34, 0xc2, 0x74, 0xa1, 0xe5, 0xa3,
	0x1b, 0xea, 0x34, 0xd0, 0x1b, 0x18, 0x5d, 0x04, 0xcf, 0x51, 0xb4, 0x4f, 0x5a, 0x46, 0x26, 0x2c,
	0xce, 0x1d, 0xd3, 0x49, 0xa8, 0x6b, 0xba, 0x89, 0x51, 0x34, 0x32, 0x39, 0xce, 0xdd, 0x59, 0xf2,
	0x12, 0xd0, 0xe8, 0x98, 0xdc, 0xcf, 0xc7, 0xd9, 0x40, 0x1a, 0xaf, 0xac, 0x9e, 0xd8, 0x58, 0x67,
	0xd9, 0x38, 0x57, 0x4e, 0x49, 0x4b, 0x7f, 0x04, 0xeb, 0x76, 0x51, 0x75, 0x96, 0x9c, 0x04, 0xcd,
	0xf1, 0x54, 0x12, 0x3d, 0x24, 0xab, 0xd9, 0xa4, 0xf0, 0x4d, 0x5b, 0x5a, 0x99, 0x3b, 0xba, 0x05,
	0x39, 0x5d, 0xf1, 0xd8, 0x39, 0x42, 0xbe, 0x4a, 0xfd, 0x81, 0x8d, 0xab, 0x44, 0xdb, 0x20, 0x6a,
	0x22, 0x5a, 0xca, 0x9e, 0x90, 0x8d, 0x89, 0x49, 0x54, 0x56, 0x68, 0xe3, 0x66, 0xb4, 0x14, 0xb4,
	0xd1, 0x0c, 0x55, 0x2e, 0x38, 0x20, 0x11, 0xde, 0x1f, 0x6e, 0x67, 0xf4, 0x3b, 0xa0, 0x6f, 0x4d,
	0x99, 0x52, 0xbe, 0x4f, 0xd6, 0x11, 0x34, 0xa2, 0x12, 0xdf, 0x01, 0xf1, 0xed, 0x12, 0x2f, 0xa5,
	0x2f, 0xc8, 0x8e, 0x95, 0xc3, 0x4c, 0xe6, 0x4e, 0x8a, 0xf2, 0xf6, 0x55, 0x6b, 0x76, 0x61, 0xcd,
	0x76, 0x25, 0x08, 0x97, 0xb1, 0x5c, 0x7b, 0x9f, 0xac, 0x54, 0xf5, 0xa1, 0x04, 0xbd, 0x8b, 0xb3,
	0x51, 0xa8, 0x8e, 0x9e, 0x88, 0x9e, 0x90, 0xcd, 0x19, 0x9e, 0x19, 0x99, 0xe0, 0x43, 0x7a, 0x0f,
	0x67, 0x82, 0x4a, 0xd8, 0x0f, 0x84, 0x2f, 0x49, 0x6d, 0x8b, 0x84, 0x71, 0x23, 0xb9, 0xdf, 0xf1,
	0x3e, 0x94, 0x2e, 0xf1, 0xd8, 0x91, 0x91, 0xbc, 0x27, 0xa2, 0x9f, 0x92, 0xc8, 0xc8, 0x4c, 0x3b,
	0x59, 0xf6, 0x31, 0xdf, 0x6d, 0xe8, 0x83, 0xf6, 0x42, 0x67, 0xb5, 0xbf, 0x8e, 0x0c, 0xa6, 0xfc,
	0x48, 0x08, 0xe3, 0x33, 0x36, 0xe2, 0x16, 0x4b, 0xd3, 0xba, 0x2f, 0xb4, 0x8d, 0x19, 0x1b, 0x71,
	0x7b, 0x1a, 0x20, 0xdf, 0x76, 0xf2, 0x71, 0x16, 0x24, 0xf4, 0x61, 0x70, 0x61, 0x9c, 0xa1, 0xc0,
	0x0f, 0x56, 0xd5, 0xea, 0xbd, 0xf6, 0x82, 0x2f, 0xde, 0xd2, 0x86, 0x22, 0x55, 0xb9, 0x50, 0xf9,
	0x30, 0x14, 0xff, 0xa3, 0x50, 0xa4, 0x08, 0x56, 0xe5, 0x9f, 0x8f, 0x94, 0x60, 0x89, 0x54, 0x82,
	0xfe, 0x18, 0x3a, 0x4b, 0xdd, 0x03, 0xaf, 0xa5, 0x12, 0x9e, 0xcc, 0x8a, 0xd4, 0x22, 0xf9, 0x18,
	0x49, 0x0f, 0x78, 0x72, 0xef, 0x2f, 0xb7, 0xc8, 0x46, 0xd9, 0x6f, 0x53, 0x1d, 0xf3, 0x14, 0xbf,
	0x32, 0x3f, 0xee, 0xd6, 0xbe, 0x31, 0xee, 0x5e, 0x1b, 0x69, 0x6f, 0xdd, 0x18, 0x69, 0x37, 0xc9,
	0xa2, 0x75, 0x3c, 0xc5, 0xff, 0x25, 0x9a, 0x7d, 0x34, 0xbc, 0xab, 0x99, 0x32, 0x46, 0x1b, 0x29,
	0xa0, 0xb7, 0x36, 0xfb, 0x95, 0xed, 0xbf, 0x99, 0x49, 0x33, 0x94, 0x7e, 0x7e, 0xf4, 0x0d, 0x24,
	0x74, 0xd7, 0x55, 0x00, 0x5f, 0x21, 0x06, 0xad, 0x40, 0xfa, 0x01, 0x93, 0xe9, 0x3c, 0xbd, 0x2a,
	0x7b, 0x2c, 0x42, 0x67, 0x79, 0x7a, 0xe5, 0xbf, 0x8b, 0x81, 0x5a, 0xc6, 0xef, 0xa2, 0x3f, 0xb3,
	0xb3, 0x6b, 0xfd, 0xfa, 0xec, 0x3a, 0x58, 0x02, 0xa7, 0x9e, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0x1e, 0x81, 0x6f, 0x13, 0x0e, 0x00, 0x00,
}
