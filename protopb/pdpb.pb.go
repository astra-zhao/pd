// Code generated by protoc-gen-go.
// source: pdpb.proto
// DO NOT EDIT!

package protopb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommandType int32

const (
	CommandType_Invalid        CommandType = 0
	CommandType_Tso            CommandType = 1
	CommandType_Bootstrap      CommandType = 2
	CommandType_IsBootstrapped CommandType = 3
	CommandType_AllocId        CommandType = 4
	CommandType_GetMeta        CommandType = 5
	CommandType_PutMeta        CommandType = 6
	CommandType_DeleteMeta     CommandType = 7
	CommandType_AskChangePeer  CommandType = 8
	CommandType_AskSplit       CommandType = 9
)

var CommandType_name = map[int32]string{
	0: "Invalid",
	1: "Tso",
	2: "Bootstrap",
	3: "IsBootstrapped",
	4: "AllocId",
	5: "GetMeta",
	6: "PutMeta",
	7: "DeleteMeta",
	8: "AskChangePeer",
	9: "AskSplit",
}
var CommandType_value = map[string]int32{
	"Invalid":        0,
	"Tso":            1,
	"Bootstrap":      2,
	"IsBootstrapped": 3,
	"AllocId":        4,
	"GetMeta":        5,
	"PutMeta":        6,
	"DeleteMeta":     7,
	"AskChangePeer":  8,
	"AskSplit":       9,
}

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}
func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (x *CommandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandType_value, data, "CommandType")
	if err != nil {
		return err
	}
	*x = CommandType(value)
	return nil
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type MetaType int32

const (
	MetaType_InvalidMeta MetaType = 0
	MetaType_NodeType    MetaType = 1
	MetaType_StoreType   MetaType = 2
	MetaType_RegionType  MetaType = 3
	MetaType_PeerType    MetaType = 4
)

var MetaType_name = map[int32]string{
	0: "InvalidMeta",
	1: "NodeType",
	2: "StoreType",
	3: "RegionType",
	4: "PeerType",
}
var MetaType_value = map[string]int32{
	"InvalidMeta": 0,
	"NodeType":    1,
	"StoreType":   2,
	"RegionType":  3,
	"PeerType":    4,
}

func (x MetaType) Enum() *MetaType {
	p := new(MetaType)
	*p = x
	return p
}
func (x MetaType) String() string {
	return proto.EnumName(MetaType_name, int32(x))
}
func (x *MetaType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetaType_value, data, "MetaType")
	if err != nil {
		return err
	}
	*x = MetaType(value)
	return nil
}
func (MetaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type Leader struct {
	Addr             *string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Pid              *int64  `protobuf:"varint,2,opt,name=pid" json:"pid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Leader) Reset()                    { *m = Leader{} }
func (m *Leader) String() string            { return proto.CompactTextString(m) }
func (*Leader) ProtoMessage()               {}
func (*Leader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Leader) GetAddr() string {
	if m != nil && m.Addr != nil {
		return *m.Addr
	}
	return ""
}

func (m *Leader) GetPid() int64 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

type TsoRequest struct {
	Number           *uint32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TsoRequest) Reset()                    { *m = TsoRequest{} }
func (m *TsoRequest) String() string            { return proto.CompactTextString(m) }
func (*TsoRequest) ProtoMessage()               {}
func (*TsoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TsoRequest) GetNumber() uint32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

type Timestamp struct {
	Physical         *int64 `protobuf:"varint,1,opt,name=physical" json:"physical,omitempty"`
	Logical          *int64 `protobuf:"varint,2,opt,name=logical" json:"logical,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Timestamp) GetPhysical() int64 {
	if m != nil && m.Physical != nil {
		return *m.Physical
	}
	return 0
}

func (m *Timestamp) GetLogical() int64 {
	if m != nil && m.Logical != nil {
		return *m.Logical
	}
	return 0
}

type TsoResponse struct {
	Timestamps       []*Timestamp `protobuf:"bytes,1,rep,name=timestamps" json:"timestamps,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TsoResponse) Reset()                    { *m = TsoResponse{} }
func (m *TsoResponse) String() string            { return proto.CompactTextString(m) }
func (*TsoResponse) ProtoMessage()               {}
func (*TsoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *TsoResponse) GetTimestamps() []*Timestamp {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

type BootstrapRequest struct {
	Node             *Node    `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Stores           []*Store `protobuf:"bytes,2,rep,name=stores" json:"stores,omitempty"`
	Region           *Region  `protobuf:"bytes,3,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BootstrapRequest) Reset()                    { *m = BootstrapRequest{} }
func (m *BootstrapRequest) String() string            { return proto.CompactTextString(m) }
func (*BootstrapRequest) ProtoMessage()               {}
func (*BootstrapRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *BootstrapRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *BootstrapRequest) GetStores() []*Store {
	if m != nil {
		return m.Stores
	}
	return nil
}

func (m *BootstrapRequest) GetRegion() *Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type BootstrapResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BootstrapResponse) Reset()                    { *m = BootstrapResponse{} }
func (m *BootstrapResponse) String() string            { return proto.CompactTextString(m) }
func (*BootstrapResponse) ProtoMessage()               {}
func (*BootstrapResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type IsBootstrappedRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *IsBootstrappedRequest) Reset()                    { *m = IsBootstrappedRequest{} }
func (m *IsBootstrappedRequest) String() string            { return proto.CompactTextString(m) }
func (*IsBootstrappedRequest) ProtoMessage()               {}
func (*IsBootstrappedRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type IsBootstrappedResponse struct {
	Bootstrapped     *bool  `protobuf:"varint,1,opt,name=bootstrapped" json:"bootstrapped,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IsBootstrappedResponse) Reset()                    { *m = IsBootstrappedResponse{} }
func (m *IsBootstrappedResponse) String() string            { return proto.CompactTextString(m) }
func (*IsBootstrappedResponse) ProtoMessage()               {}
func (*IsBootstrappedResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *IsBootstrappedResponse) GetBootstrapped() bool {
	if m != nil && m.Bootstrapped != nil {
		return *m.Bootstrapped
	}
	return false
}

type AllocIdRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AllocIdRequest) Reset()                    { *m = AllocIdRequest{} }
func (m *AllocIdRequest) String() string            { return proto.CompactTextString(m) }
func (*AllocIdRequest) ProtoMessage()               {}
func (*AllocIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type AllocIdResponse struct {
	Id               *uint64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AllocIdResponse) Reset()                    { *m = AllocIdResponse{} }
func (m *AllocIdResponse) String() string            { return proto.CompactTextString(m) }
func (*AllocIdResponse) ProtoMessage()               {}
func (*AllocIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *AllocIdResponse) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type GetMetaRequest struct {
	MetaType         *MetaType `protobuf:"varint,1,opt,name=meta_type,enum=pdpb.MetaType" json:"meta_type,omitempty"`
	NodeId           *uint64   `protobuf:"varint,2,opt,name=node_id" json:"node_id,omitempty"`
	StoreId          *uint64   `protobuf:"varint,3,opt,name=store_id" json:"store_id,omitempty"`
	RegionKey        []byte    `protobuf:"bytes,4,opt,name=region_key" json:"region_key,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *GetMetaRequest) Reset()                    { *m = GetMetaRequest{} }
func (m *GetMetaRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMetaRequest) ProtoMessage()               {}
func (*GetMetaRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *GetMetaRequest) GetMetaType() MetaType {
	if m != nil && m.MetaType != nil {
		return *m.MetaType
	}
	return MetaType_InvalidMeta
}

func (m *GetMetaRequest) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *GetMetaRequest) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

func (m *GetMetaRequest) GetRegionKey() []byte {
	if m != nil {
		return m.RegionKey
	}
	return nil
}

type GetMetaResponse struct {
	MetaType         *MetaType `protobuf:"varint,1,opt,name=meta_type,enum=pdpb.MetaType" json:"meta_type,omitempty"`
	Node             *Node     `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	Store            *Store    `protobuf:"bytes,3,opt,name=store" json:"store,omitempty"`
	Region           *Region   `protobuf:"bytes,4,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *GetMetaResponse) Reset()                    { *m = GetMetaResponse{} }
func (m *GetMetaResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMetaResponse) ProtoMessage()               {}
func (*GetMetaResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *GetMetaResponse) GetMetaType() MetaType {
	if m != nil && m.MetaType != nil {
		return *m.MetaType
	}
	return MetaType_InvalidMeta
}

func (m *GetMetaResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *GetMetaResponse) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *GetMetaResponse) GetRegion() *Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type PutMetaRequest struct {
	MetaType         *MetaType `protobuf:"varint,1,opt,name=meta_type,enum=pdpb.MetaType" json:"meta_type,omitempty"`
	Node             *Node     `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	Store            *Store    `protobuf:"bytes,3,opt,name=store" json:"store,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *PutMetaRequest) Reset()                    { *m = PutMetaRequest{} }
func (m *PutMetaRequest) String() string            { return proto.CompactTextString(m) }
func (*PutMetaRequest) ProtoMessage()               {}
func (*PutMetaRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *PutMetaRequest) GetMetaType() MetaType {
	if m != nil && m.MetaType != nil {
		return *m.MetaType
	}
	return MetaType_InvalidMeta
}

func (m *PutMetaRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *PutMetaRequest) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

type PutMetaResponse struct {
	MetaType         *MetaType `protobuf:"varint,1,opt,name=meta_type,enum=pdpb.MetaType" json:"meta_type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *PutMetaResponse) Reset()                    { *m = PutMetaResponse{} }
func (m *PutMetaResponse) String() string            { return proto.CompactTextString(m) }
func (*PutMetaResponse) ProtoMessage()               {}
func (*PutMetaResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *PutMetaResponse) GetMetaType() MetaType {
	if m != nil && m.MetaType != nil {
		return *m.MetaType
	}
	return MetaType_InvalidMeta
}

type DeleteMetaRequest struct {
	MetaType         *MetaType `protobuf:"varint,1,opt,name=meta_type,enum=pdpb.MetaType" json:"meta_type,omitempty"`
	NodeId           *uint64   `protobuf:"varint,2,opt,name=node_id" json:"node_id,omitempty"`
	StoreId          *uint64   `protobuf:"varint,3,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *DeleteMetaRequest) Reset()                    { *m = DeleteMetaRequest{} }
func (m *DeleteMetaRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteMetaRequest) ProtoMessage()               {}
func (*DeleteMetaRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *DeleteMetaRequest) GetMetaType() MetaType {
	if m != nil && m.MetaType != nil {
		return *m.MetaType
	}
	return MetaType_InvalidMeta
}

func (m *DeleteMetaRequest) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *DeleteMetaRequest) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

type DeleteMetaResponse struct {
	MetaType         *MetaType `protobuf:"varint,1,opt,name=meta_type,enum=pdpb.MetaType" json:"meta_type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *DeleteMetaResponse) Reset()                    { *m = DeleteMetaResponse{} }
func (m *DeleteMetaResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteMetaResponse) ProtoMessage()               {}
func (*DeleteMetaResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *DeleteMetaResponse) GetMetaType() MetaType {
	if m != nil && m.MetaType != nil {
		return *m.MetaType
	}
	return MetaType_InvalidMeta
}

type AskChangePeerRequest struct {
	Region           *Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AskChangePeerRequest) Reset()                    { *m = AskChangePeerRequest{} }
func (m *AskChangePeerRequest) String() string            { return proto.CompactTextString(m) }
func (*AskChangePeerRequest) ProtoMessage()               {}
func (*AskChangePeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *AskChangePeerRequest) GetRegion() *Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type AskChangePeerResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AskChangePeerResponse) Reset()                    { *m = AskChangePeerResponse{} }
func (m *AskChangePeerResponse) String() string            { return proto.CompactTextString(m) }
func (*AskChangePeerResponse) ProtoMessage()               {}
func (*AskChangePeerResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

type AskSplitRequest struct {
	Region           *Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	SplitKey         []byte  `protobuf:"bytes,2,opt,name=split_key" json:"split_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AskSplitRequest) Reset()                    { *m = AskSplitRequest{} }
func (m *AskSplitRequest) String() string            { return proto.CompactTextString(m) }
func (*AskSplitRequest) ProtoMessage()               {}
func (*AskSplitRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

func (m *AskSplitRequest) GetRegion() *Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *AskSplitRequest) GetSplitKey() []byte {
	if m != nil {
		return m.SplitKey
	}
	return nil
}

type AskSplitResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AskSplitResponse) Reset()                    { *m = AskSplitResponse{} }
func (m *AskSplitResponse) String() string            { return proto.CompactTextString(m) }
func (*AskSplitResponse) ProtoMessage()               {}
func (*AskSplitResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

type RequestHeader struct {
	// 16 bytes, to distinguish request.
	Uuid             []byte  `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	ClusterId        *uint64 `protobuf:"varint,2,opt,name=cluster_id" json:"cluster_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{20} }

func (m *RequestHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *RequestHeader) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

type ResponseHeader struct {
	// 16 bytes, to distinguish request.
	Uuid             []byte  `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	ClusterId        *uint64 `protobuf:"varint,2,opt,name=cluster_id" json:"cluster_id,omitempty"`
	Error            *Error  `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{21} }

func (m *ResponseHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *ResponseHeader) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

func (m *ResponseHeader) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Request struct {
	Header           *RequestHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CmdType          *CommandType           `protobuf:"varint,2,opt,name=cmd_type,enum=pdpb.CommandType" json:"cmd_type,omitempty"`
	Tso              *TsoRequest            `protobuf:"bytes,3,opt,name=tso" json:"tso,omitempty"`
	Bootstrap        *BootstrapRequest      `protobuf:"bytes,4,opt,name=bootstrap" json:"bootstrap,omitempty"`
	IsBootstrapped   *IsBootstrappedRequest `protobuf:"bytes,5,opt,name=is_bootstrapped" json:"is_bootstrapped,omitempty"`
	AllocId          *AllocIdRequest        `protobuf:"bytes,6,opt,name=alloc_id" json:"alloc_id,omitempty"`
	GetMeta          *GetMetaRequest        `protobuf:"bytes,7,opt,name=get_meta" json:"get_meta,omitempty"`
	PutMeta          *PutMetaRequest        `protobuf:"bytes,8,opt,name=put_meta" json:"put_meta,omitempty"`
	DeleteMeta       *DeleteMetaRequest     `protobuf:"bytes,9,opt,name=delete_meta" json:"delete_meta,omitempty"`
	AskChangePeer    *AskChangePeerRequest  `protobuf:"bytes,10,opt,name=ask_change_peer" json:"ask_change_peer,omitempty"`
	AskSplit         *AskSplitRequest       `protobuf:"bytes,11,opt,name=ask_split" json:"ask_split,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{22} }

func (m *Request) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetCmdType() CommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CommandType_Invalid
}

func (m *Request) GetTso() *TsoRequest {
	if m != nil {
		return m.Tso
	}
	return nil
}

func (m *Request) GetBootstrap() *BootstrapRequest {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *Request) GetIsBootstrapped() *IsBootstrappedRequest {
	if m != nil {
		return m.IsBootstrapped
	}
	return nil
}

func (m *Request) GetAllocId() *AllocIdRequest {
	if m != nil {
		return m.AllocId
	}
	return nil
}

func (m *Request) GetGetMeta() *GetMetaRequest {
	if m != nil {
		return m.GetMeta
	}
	return nil
}

func (m *Request) GetPutMeta() *PutMetaRequest {
	if m != nil {
		return m.PutMeta
	}
	return nil
}

func (m *Request) GetDeleteMeta() *DeleteMetaRequest {
	if m != nil {
		return m.DeleteMeta
	}
	return nil
}

func (m *Request) GetAskChangePeer() *AskChangePeerRequest {
	if m != nil {
		return m.AskChangePeer
	}
	return nil
}

func (m *Request) GetAskSplit() *AskSplitRequest {
	if m != nil {
		return m.AskSplit
	}
	return nil
}

type Response struct {
	Header           *ResponseHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CmdType          *CommandType            `protobuf:"varint,2,opt,name=cmd_type,enum=pdpb.CommandType" json:"cmd_type,omitempty"`
	Tso              *TsoResponse            `protobuf:"bytes,3,opt,name=tso" json:"tso,omitempty"`
	Bootstrap        *BootstrapResponse      `protobuf:"bytes,4,opt,name=bootstrap" json:"bootstrap,omitempty"`
	IsBootstrapped   *IsBootstrappedResponse `protobuf:"bytes,5,opt,name=is_bootstrapped" json:"is_bootstrapped,omitempty"`
	AllocId          *AllocIdResponse        `protobuf:"bytes,6,opt,name=alloc_id" json:"alloc_id,omitempty"`
	GetMeta          *GetMetaResponse        `protobuf:"bytes,7,opt,name=get_meta" json:"get_meta,omitempty"`
	PutMeta          *PutMetaResponse        `protobuf:"bytes,8,opt,name=put_meta" json:"put_meta,omitempty"`
	DeleteMeta       *DeleteMetaResponse     `protobuf:"bytes,9,opt,name=delete_meta" json:"delete_meta,omitempty"`
	AskChangePeer    *AskChangePeerResponse  `protobuf:"bytes,10,opt,name=ask_change_peer" json:"ask_change_peer,omitempty"`
	AskSplit         *AskSplitResponse       `protobuf:"bytes,11,opt,name=ask_split" json:"ask_split,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{23} }

func (m *Response) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetCmdType() CommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CommandType_Invalid
}

func (m *Response) GetTso() *TsoResponse {
	if m != nil {
		return m.Tso
	}
	return nil
}

func (m *Response) GetBootstrap() *BootstrapResponse {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *Response) GetIsBootstrapped() *IsBootstrappedResponse {
	if m != nil {
		return m.IsBootstrapped
	}
	return nil
}

func (m *Response) GetAllocId() *AllocIdResponse {
	if m != nil {
		return m.AllocId
	}
	return nil
}

func (m *Response) GetGetMeta() *GetMetaResponse {
	if m != nil {
		return m.GetMeta
	}
	return nil
}

func (m *Response) GetPutMeta() *PutMetaResponse {
	if m != nil {
		return m.PutMeta
	}
	return nil
}

func (m *Response) GetDeleteMeta() *DeleteMetaResponse {
	if m != nil {
		return m.DeleteMeta
	}
	return nil
}

func (m *Response) GetAskChangePeer() *AskChangePeerResponse {
	if m != nil {
		return m.AskChangePeer
	}
	return nil
}

func (m *Response) GetAskSplit() *AskSplitResponse {
	if m != nil {
		return m.AskSplit
	}
	return nil
}

type BootstrappedError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BootstrappedError) Reset()                    { *m = BootstrappedError{} }
func (m *BootstrappedError) String() string            { return proto.CompactTextString(m) }
func (*BootstrappedError) ProtoMessage()               {}
func (*BootstrappedError) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{24} }

type Error struct {
	Message          *string            `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Bootstrapped     *BootstrappedError `protobuf:"bytes,2,opt,name=bootstrapped" json:"bootstrapped,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{25} }

func (m *Error) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Error) GetBootstrapped() *BootstrappedError {
	if m != nil {
		return m.Bootstrapped
	}
	return nil
}

func init() {
	proto.RegisterType((*Leader)(nil), "pdpb.Leader")
	proto.RegisterType((*TsoRequest)(nil), "pdpb.TsoRequest")
	proto.RegisterType((*Timestamp)(nil), "pdpb.Timestamp")
	proto.RegisterType((*TsoResponse)(nil), "pdpb.TsoResponse")
	proto.RegisterType((*BootstrapRequest)(nil), "pdpb.BootstrapRequest")
	proto.RegisterType((*BootstrapResponse)(nil), "pdpb.BootstrapResponse")
	proto.RegisterType((*IsBootstrappedRequest)(nil), "pdpb.IsBootstrappedRequest")
	proto.RegisterType((*IsBootstrappedResponse)(nil), "pdpb.IsBootstrappedResponse")
	proto.RegisterType((*AllocIdRequest)(nil), "pdpb.AllocIdRequest")
	proto.RegisterType((*AllocIdResponse)(nil), "pdpb.AllocIdResponse")
	proto.RegisterType((*GetMetaRequest)(nil), "pdpb.GetMetaRequest")
	proto.RegisterType((*GetMetaResponse)(nil), "pdpb.GetMetaResponse")
	proto.RegisterType((*PutMetaRequest)(nil), "pdpb.PutMetaRequest")
	proto.RegisterType((*PutMetaResponse)(nil), "pdpb.PutMetaResponse")
	proto.RegisterType((*DeleteMetaRequest)(nil), "pdpb.DeleteMetaRequest")
	proto.RegisterType((*DeleteMetaResponse)(nil), "pdpb.DeleteMetaResponse")
	proto.RegisterType((*AskChangePeerRequest)(nil), "pdpb.AskChangePeerRequest")
	proto.RegisterType((*AskChangePeerResponse)(nil), "pdpb.AskChangePeerResponse")
	proto.RegisterType((*AskSplitRequest)(nil), "pdpb.AskSplitRequest")
	proto.RegisterType((*AskSplitResponse)(nil), "pdpb.AskSplitResponse")
	proto.RegisterType((*RequestHeader)(nil), "pdpb.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "pdpb.ResponseHeader")
	proto.RegisterType((*Request)(nil), "pdpb.Request")
	proto.RegisterType((*Response)(nil), "pdpb.Response")
	proto.RegisterType((*BootstrappedError)(nil), "pdpb.BootstrappedError")
	proto.RegisterType((*Error)(nil), "pdpb.Error")
	proto.RegisterEnum("pdpb.CommandType", CommandType_name, CommandType_value)
	proto.RegisterEnum("pdpb.MetaType", MetaType_name, MetaType_value)
}

var fileDescriptor1 = []byte{
	// 937 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xde, 0xd8, 0x4e, 0x62, 0x1f, 0x27, 0x8e, 0xe3, 0x4d, 0x5a, 0xab, 0xdb, 0xae, 0xc0, 0x45,
	0xd0, 0xad, 0xd8, 0x48, 0x94, 0x05, 0xae, 0xf7, 0x07, 0x95, 0x4a, 0x50, 0x55, 0x6d, 0x85, 0x04,
	0x37, 0x91, 0x13, 0x8f, 0x52, 0x13, 0x3b, 0x36, 0x1e, 0x07, 0xa9, 0x2f, 0xc1, 0x23, 0xf0, 0x5c,
	0xdc, 0xf3, 0x22, 0xcc, 0x9c, 0x19, 0x3b, 0x8e, 0xe3, 0x4a, 0x01, 0xed, 0x9d, 0xe7, 0xcc, 0x77,
	0x7e, 0xe6, 0xfb, 0xce, 0x9c, 0x31, 0x40, 0x1a, 0xa4, 0xb3, 0x49, 0x9a, 0x25, 0x79, 0xe2, 0x68,
	0xfc, 0xfb, 0xa8, 0x17, 0x93, 0xdc, 0x2f, 0x6c, 0xde, 0x29, 0x74, 0x7e, 0x24, 0x7e, 0x40, 0x32,
	0xa7, 0x07, 0x9a, 0x1f, 0x04, 0x99, 0xdb, 0xfa, 0xa4, 0x75, 0x66, 0x38, 0x26, 0xa8, 0x69, 0x18,
	0xb8, 0x0a, 0x5b, 0xa8, 0xde, 0x31, 0xc0, 0x3d, 0x4d, 0x6e, 0xc9, 0xef, 0x6b, 0x42, 0x73, 0xc7,
	0x82, 0xce, 0x6a, 0x1d, 0xcf, 0x88, 0x80, 0xf6, 0xbd, 0x09, 0x18, 0xf7, 0x61, 0xcc, 0x76, 0xfc,
	0x38, 0x75, 0x6c, 0xd0, 0xd3, 0x87, 0x47, 0x1a, 0xce, 0xfd, 0x08, 0xb7, 0x55, 0x67, 0x00, 0xdd,
	0x28, 0x59, 0xa0, 0x41, 0x44, 0xbb, 0x00, 0x13, 0xa3, 0xd1, 0x34, 0x59, 0x51, 0xe2, 0x9c, 0x02,
	0xe4, 0x85, 0x3b, 0x65, 0x3e, 0xea, 0x99, 0x79, 0x31, 0x98, 0x60, 0xd9, 0x65, 0x58, 0x2f, 0x06,
	0xfb, 0x5d, 0x92, 0xe4, 0x34, 0xcf, 0xfc, 0xb4, 0xa8, 0xe3, 0x08, 0xb4, 0x55, 0x12, 0x10, 0x4c,
	0x63, 0x5e, 0xf4, 0x26, 0xf2, 0x5c, 0xd7, 0xcc, 0xe6, 0x9c, 0x40, 0x87, 0xe6, 0x49, 0x46, 0x28,
	0xcb, 0xc9, 0x03, 0xf6, 0x8b, 0xdd, 0x3b, 0x6e, 0x75, 0x5e, 0x42, 0x27, 0x23, 0x8b, 0x30, 0x59,
	0xb9, 0x2a, 0x3a, 0x5b, 0xc5, 0xf6, 0x2d, 0x5a, 0xbd, 0xe7, 0x30, 0xac, 0xa4, 0x13, 0x85, 0x7a,
	0x87, 0x30, 0xbe, 0xa2, 0xa5, 0x39, 0x25, 0x81, 0x2c, 0x84, 0x11, 0x70, 0x50, 0xdf, 0x90, 0x67,
	0x1b, 0x41, 0x6f, 0x56, 0xb1, 0x63, 0xa9, 0xba, 0x67, 0x83, 0xf5, 0x36, 0x8a, 0x92, 0xf9, 0x55,
	0x19, 0xe1, 0x04, 0x06, 0xa5, 0x45, 0xba, 0x02, 0x28, 0x92, 0x7f, 0xcd, 0xfb, 0x0d, 0xac, 0x4b,
	0x92, 0xff, 0xc4, 0x4a, 0x2c, 0xce, 0xfe, 0x29, 0x18, 0xbc, 0xe2, 0x69, 0xfe, 0x98, 0x0a, 0x02,
	0x2c, 0x76, 0x06, 0xe4, 0x8c, 0xa3, 0xee, 0x99, 0x95, 0xf3, 0xce, 0xe9, 0x99, 0x16, 0x51, 0xb8,
	0x34, 0xc8, 0x09, 0xb7, 0xa8, 0x68, 0x61, 0x49, 0x04, 0x0d, 0xd3, 0x25, 0x79, 0x74, 0x35, 0x66,
	0xeb, 0x79, 0x7f, 0xb6, 0x60, 0x50, 0x26, 0x93, 0xb5, 0xec, 0x91, 0xad, 0x10, 0x43, 0x69, 0x10,
	0xe3, 0x18, 0xda, 0x98, 0x58, 0x92, 0xfd, 0xa4, 0x16, 0x5a, 0xa3, 0x16, 0x31, 0x58, 0x37, 0xeb,
	0xff, 0x7a, 0xf8, 0xff, 0x5d, 0x8e, 0xf7, 0x06, 0x06, 0x65, 0xba, 0xbd, 0x8f, 0xef, 0xfd, 0x02,
	0xc3, 0x0f, 0x24, 0x22, 0x39, 0xf9, 0xe8, 0x22, 0x79, 0xdf, 0x81, 0x53, 0x0d, 0xbd, 0x7f, 0x4d,
	0xdf, 0xc2, 0xe8, 0x2d, 0x5d, 0xbe, 0x7f, 0xf0, 0x57, 0x0b, 0x72, 0x43, 0x48, 0x56, 0x94, 0xb5,
	0x21, 0xbc, 0xd5, 0x48, 0x38, 0xeb, 0xf3, 0x9a, 0x9f, 0xbc, 0x00, 0x1f, 0x58, 0x97, 0xd2, 0xe5,
	0x5d, 0x1a, 0x85, 0xf9, 0x9e, 0xb1, 0x9c, 0x21, 0x18, 0x94, 0xe3, 0xb1, 0xc1, 0x14, 0x6c, 0x30,
	0x07, 0xec, 0x4d, 0x14, 0x19, 0xf9, 0x2b, 0xe8, 0xcb, 0x88, 0x3f, 0x94, 0xc3, 0x68, 0xbd, 0x0e,
	0xc5, 0x85, 0xe9, 0xf1, 0x3e, 0x9d, 0x47, 0x6b, 0x9a, 0x93, 0xac, 0x24, 0xca, 0xbb, 0x06, 0xab,
	0x70, 0xdf, 0xd7, 0x87, 0x75, 0x45, 0x9b, 0x64, 0x59, 0x92, 0x49, 0xe5, 0x4d, 0x41, 0xd8, 0xf7,
	0xdc, 0xe4, 0xfd, 0xad, 0x42, 0xb7, 0x38, 0x15, 0x1b, 0x8a, 0x0f, 0x18, 0x53, 0x9e, 0xea, 0xb9,
	0x00, 0x6e, 0x97, 0x78, 0x0a, 0xfa, 0x3c, 0x0e, 0x84, 0x00, 0x0a, 0x0a, 0x30, 0x14, 0xb0, 0xf7,
	0x49, 0x1c, 0xfb, 0xab, 0x00, 0xf5, 0x3d, 0x01, 0x35, 0xa7, 0x89, 0xcc, 0x67, 0xcb, 0xa9, 0xb6,
	0x19, 0xa5, 0xaf, 0xc0, 0x28, 0xe7, 0x83, 0x6c, 0xff, 0x03, 0x01, 0xda, 0x99, 0x76, 0xac, 0x2f,
	0x43, 0x3a, 0xdd, 0x9a, 0x26, 0x6d, 0x74, 0x78, 0x21, 0x1c, 0x1a, 0x47, 0x93, 0xf3, 0x39, 0xe8,
	0x3e, 0x1f, 0x2c, 0x9c, 0x83, 0x0e, 0xc2, 0x47, 0x02, 0xbe, 0x3d, 0x80, 0x38, 0x6e, 0x41, 0xf2,
	0x29, 0x17, 0xcf, 0xed, 0x56, 0x71, 0xb5, 0xb9, 0xc3, 0x70, 0xe9, 0x5a, 0xe2, 0xf4, 0x2a, 0xae,
	0x76, 0x45, 0xbf, 0x04, 0x33, 0xc0, 0xa6, 0x15, 0x50, 0x03, 0xa1, 0x87, 0x02, 0xba, 0x7b, 0x51,
	0xbe, 0x86, 0x81, 0x4f, 0x97, 0xd3, 0x39, 0xb6, 0xdc, 0x34, 0x65, 0x3d, 0xe7, 0x02, 0x7a, 0x1c,
	0xc9, 0x62, 0x9b, 0xda, 0xf8, 0x0c, 0x0c, 0xee, 0x84, 0xed, 0xe5, 0x9a, 0x08, 0x1f, 0x97, 0xf0,
	0x6a, 0x93, 0x7a, 0xff, 0xa8, 0xa0, 0x97, 0x17, 0xe7, 0xb3, 0x9a, 0xb6, 0xa3, 0x42, 0xdb, 0xad,
	0x5e, 0xda, 0x4b, 0xdc, 0x97, 0x55, 0x71, 0x87, 0x15, 0x71, 0x65, 0xaa, 0xf3, 0x5d, 0x75, 0x0f,
	0x77, 0xd4, 0x95, 0xd8, 0x6f, 0x9e, 0x92, 0xf7, 0xb8, 0x59, 0x5e, 0xe9, 0xf6, 0xc5, 0x8e, 0xbe,
	0xe3, 0x9a, 0xbe, 0x1b, 0x60, 0x4d, 0xe0, 0x71, 0x4d, 0xe0, 0x0d, 0xb0, 0xa6, 0xf0, 0xb8, 0xa6,
	0xb0, 0x04, 0xbe, 0x6e, 0x92, 0xd8, 0xdd, 0x95, 0x58, 0xc2, 0xdf, 0x3c, 0xa5, 0xf1, 0x8b, 0x46,
	0x8d, 0xa5, 0xd7, 0xab, 0x5d, 0x91, 0x0f, 0xea, 0x22, 0xcb, 0x19, 0x52, 0x7d, 0xb3, 0x19, 0x45,
	0xe2, 0x56, 0x5f, 0x42, 0x1b, 0x3f, 0xf8, 0xa0, 0x65, 0xff, 0x12, 0xd4, 0x5f, 0x10, 0xf9, 0x83,
	0xf3, 0xba, 0xf6, 0x34, 0x2b, 0x8d, 0xfa, 0x14, 0x81, 0xce, 0xff, 0x6a, 0x81, 0x59, 0xd5, 0xde,
	0x84, 0xee, 0xd5, 0xea, 0x0f, 0x3f, 0x0a, 0x03, 0xfb, 0x99, 0xd3, 0x05, 0x95, 0xe9, 0x6e, 0xb7,
	0x9c, 0x3e, 0x18, 0xa5, 0xab, 0xad, 0xb0, 0x19, 0x64, 0x6d, 0xeb, 0x66, 0xab, 0xdc, 0x51, 0x6a,
	0x63, 0x6b, 0x7c, 0x21, 0xf9, 0xb7, 0xdb, 0x7c, 0x21, 0x39, 0xb6, 0x3b, 0xec, 0x27, 0x0b, 0x36,
	0x24, 0xda, 0x5d, 0x36, 0x48, 0xfb, 0x5b, 0x0c, 0xd9, 0x3a, 0x9b, 0x77, 0x7a, 0x41, 0x82, 0x6d,
	0x9c, 0xff, 0x0c, 0x7a, 0xe5, 0x55, 0x31, 0x65, 0x71, 0xe8, 0xfd, 0x8c, 0x43, 0xf9, 0xd3, 0xc7,
	0x37, 0x45, 0x95, 0xf8, 0xd6, 0xe1, 0x52, 0xe1, 0xa9, 0xc4, 0xb4, 0xc6, 0xb5, 0xca, 0xc1, 0x3c,
	0x03, 0xae, 0xb4, 0x77, 0xc6, 0xaf, 0x5d, 0xfc, 0x53, 0x4c, 0x67, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0x20, 0x26, 0x15, 0x4c, 0x0a, 0x00, 0x00,
}
