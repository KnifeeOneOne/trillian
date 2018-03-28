// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trillian.proto

package trillian

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keyspb "github.com/google/trillian/crypto/keyspb"
import sigpb "github.com/google/trillian/crypto/sigpb"
import google_protobuf2 "github.com/golang/protobuf/ptypes/any"
import google_protobuf3 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// LogRootFormat specifies the fields that are covered by the
// SignedLogRoot signature, as well as their ordering and formats.
type LogRootFormat int32

const (
	LogRootFormat_LOG_ROOT_FORMAT_UNKNOWN LogRootFormat = 0
	LogRootFormat_LOG_ROOT_FORMAT_V1      LogRootFormat = 1
)

var LogRootFormat_name = map[int32]string{
	0: "LOG_ROOT_FORMAT_UNKNOWN",
	1: "LOG_ROOT_FORMAT_V1",
}
var LogRootFormat_value = map[string]int32{
	"LOG_ROOT_FORMAT_UNKNOWN": 0,
	"LOG_ROOT_FORMAT_V1":      1,
}

func (x LogRootFormat) String() string {
	return proto.EnumName(LogRootFormat_name, int32(x))
}
func (LogRootFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// MapRootFormat specifies the fields that are covered by the
// SignedMapRoot signature, as well as their ordering and formats.
type MapRootFormat int32

const (
	MapRootFormat_MAP_ROOT_FORMAT_UNKNOWN MapRootFormat = 0
	MapRootFormat_MAP_ROOT_FORMAT_V1      MapRootFormat = 1
)

var MapRootFormat_name = map[int32]string{
	0: "MAP_ROOT_FORMAT_UNKNOWN",
	1: "MAP_ROOT_FORMAT_V1",
}
var MapRootFormat_value = map[string]int32{
	"MAP_ROOT_FORMAT_UNKNOWN": 0,
	"MAP_ROOT_FORMAT_V1":      1,
}

func (x MapRootFormat) String() string {
	return proto.EnumName(MapRootFormat_name, int32(x))
}
func (MapRootFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

// Defines the way empty / node / leaf hashes are constructed incorporating
// preimage protection, which can be application specific.
type HashStrategy int32

const (
	// Hash strategy cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	HashStrategy_UNKNOWN_HASH_STRATEGY HashStrategy = 0
	// Certificate Transparency strategy: leaf hash prefix = 0x00, node prefix =
	// 0x01, empty hash is digest([]byte{}), as defined in the specification.
	HashStrategy_RFC6962_SHA256 HashStrategy = 1
	// Sparse Merkle Tree strategy:  leaf hash prefix = 0x00, node prefix = 0x01,
	// empty branch is recursively computed from empty leaf nodes.
	// NOT secure in a multi tree environment. For testing only.
	HashStrategy_TEST_MAP_HASHER HashStrategy = 2
	// Append-only log strategy where leaf nodes are defined as the ObjectHash.
	// All other properties are equal to RFC6962_SHA256.
	HashStrategy_OBJECT_RFC6962_SHA256 HashStrategy = 3
	// The CONIKS sparse tree hasher with SHA512_256 as the hash algorithm.
	HashStrategy_CONIKS_SHA512_256 HashStrategy = 4
)

var HashStrategy_name = map[int32]string{
	0: "UNKNOWN_HASH_STRATEGY",
	1: "RFC6962_SHA256",
	2: "TEST_MAP_HASHER",
	3: "OBJECT_RFC6962_SHA256",
	4: "CONIKS_SHA512_256",
}
var HashStrategy_value = map[string]int32{
	"UNKNOWN_HASH_STRATEGY": 0,
	"RFC6962_SHA256":        1,
	"TEST_MAP_HASHER":       2,
	"OBJECT_RFC6962_SHA256": 3,
	"CONIKS_SHA512_256":     4,
}

func (x HashStrategy) String() string {
	return proto.EnumName(HashStrategy_name, int32(x))
}
func (HashStrategy) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

// State of the tree.
type TreeState int32

const (
	// Tree state cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	TreeState_UNKNOWN_TREE_STATE TreeState = 0
	// Active trees are able to respond to both read and write requests.
	TreeState_ACTIVE TreeState = 1
	// Frozen trees are only able to respond to read requests, writing to a frozen
	// tree is forbidden. Trees should not be frozen when there are entries
	// in the queue that have not yet been integrated. See the DRAINING
	// state for this case.
	TreeState_FROZEN TreeState = 2
	// Deprecated in favor of Tree.deleted.
	TreeState_DEPRECATED_SOFT_DELETED TreeState = 3
	// Deprecated in favor of Tree.deleted.
	TreeState_DEPRECATED_HARD_DELETED TreeState = 4
	// A tree that is draining will continue to integrate queued entries.
	// No new entries should be accepted.
	TreeState_DRAINING TreeState = 5
)

var TreeState_name = map[int32]string{
	0: "UNKNOWN_TREE_STATE",
	1: "ACTIVE",
	2: "FROZEN",
	3: "DEPRECATED_SOFT_DELETED",
	4: "DEPRECATED_HARD_DELETED",
	5: "DRAINING",
}
var TreeState_value = map[string]int32{
	"UNKNOWN_TREE_STATE":      0,
	"ACTIVE":                  1,
	"FROZEN":                  2,
	"DEPRECATED_SOFT_DELETED": 3,
	"DEPRECATED_HARD_DELETED": 4,
	"DRAINING":                5,
}

func (x TreeState) String() string {
	return proto.EnumName(TreeState_name, int32(x))
}
func (TreeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

// Type of the tree.
type TreeType int32

const (
	// Tree type cannot be determined. Included to enable detection of mismatched
	// proto versions being used. Represents an invalid value.
	TreeType_UNKNOWN_TREE_TYPE TreeType = 0
	// Tree represents a verifiable log.
	TreeType_LOG TreeType = 1
	// Tree represents a verifiable map.
	TreeType_MAP TreeType = 2
	// Tree represents a verifiable pre-ordered log, i.e., a log whose entries are
	// placed according to sequence numbers assigned outside of Trillian.
	TreeType_PREORDERED_LOG TreeType = 3
)

var TreeType_name = map[int32]string{
	0: "UNKNOWN_TREE_TYPE",
	1: "LOG",
	2: "MAP",
	3: "PREORDERED_LOG",
}
var TreeType_value = map[string]int32{
	"UNKNOWN_TREE_TYPE": 0,
	"LOG":               1,
	"MAP":               2,
	"PREORDERED_LOG":    3,
}

func (x TreeType) String() string {
	return proto.EnumName(TreeType_name, int32(x))
}
func (TreeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

// Represents a tree, which may be either a verifiable log or map.
// Readonly attributes are assigned at tree creation, after which they may not
// be modified.
//
// Note: Many APIs within the rest of the code require these objects to
// be provided. For safety they should be obtained via Admin API calls and
// not created dynamically.
type Tree struct {
	// ID of the tree.
	// Readonly.
	TreeId int64 `protobuf:"varint,1,opt,name=tree_id,json=treeId" json:"tree_id,omitempty"`
	// State of the tree.
	// Trees are ACTIVE after creation. At any point the tree may transition
	// between ACTIVE, DRAINING and FROZEN states.
	TreeState TreeState `protobuf:"varint,2,opt,name=tree_state,json=treeState,enum=trillian.TreeState" json:"tree_state,omitempty"`
	// Type of the tree.
	// Readonly after Tree creation. Exception: Can be switched from
	// PREORDERED_LOG to LOG if the Tree is and remains in the FROZEN state.
	TreeType TreeType `protobuf:"varint,3,opt,name=tree_type,json=treeType,enum=trillian.TreeType" json:"tree_type,omitempty"`
	// Hash strategy to be used by the tree.
	// Readonly.
	HashStrategy HashStrategy `protobuf:"varint,4,opt,name=hash_strategy,json=hashStrategy,enum=trillian.HashStrategy" json:"hash_strategy,omitempty"`
	// Hash algorithm to be used by the tree.
	// Readonly.
	HashAlgorithm sigpb.DigitallySigned_HashAlgorithm `protobuf:"varint,5,opt,name=hash_algorithm,json=hashAlgorithm,enum=sigpb.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	// Signature algorithm to be used by the tree.
	// Readonly.
	SignatureAlgorithm sigpb.DigitallySigned_SignatureAlgorithm `protobuf:"varint,6,opt,name=signature_algorithm,json=signatureAlgorithm,enum=sigpb.DigitallySigned_SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	// Display name of the tree.
	// Optional.
	DisplayName string `protobuf:"bytes,8,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Description of the tree,
	// Optional.
	Description string `protobuf:"bytes,9,opt,name=description" json:"description,omitempty"`
	// Identifies the private key used for signing tree heads and entry
	// timestamps.
	// This can be any type of message to accommodate different key management
	// systems, e.g. PEM files, HSMs, etc.
	// Private keys are write-only: they're never returned by RPCs.
	// The private_key message can be changed after a tree is created, but the
	// underlying key must remain the same - this is to enable migrating a key
	// from one provider to another.
	PrivateKey *google_protobuf2.Any `protobuf:"bytes,12,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	// Storage-specific settings.
	// Varies according to the storage implementation backing Trillian.
	StorageSettings *google_protobuf2.Any `protobuf:"bytes,13,opt,name=storage_settings,json=storageSettings" json:"storage_settings,omitempty"`
	// The public key used for verifying tree heads and entry timestamps.
	// Readonly.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,14,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	// Interval after which a new signed root is produced even if there have been
	// no submission.  If zero, this behavior is disabled.
	MaxRootDuration *google_protobuf3.Duration `protobuf:"bytes,15,opt,name=max_root_duration,json=maxRootDuration" json:"max_root_duration,omitempty"`
	// Time of tree creation.
	// Readonly.
	CreateTime *google_protobuf1.Timestamp `protobuf:"bytes,16,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// Time of last tree update.
	// Readonly (automatically assigned on updates).
	UpdateTime *google_protobuf1.Timestamp `protobuf:"bytes,17,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	// If true, the tree has been deleted.
	// Deleted trees may be undeleted during a certain time window, after which
	// they're permanently deleted (and unrecoverable).
	// Readonly.
	Deleted bool `protobuf:"varint,19,opt,name=deleted" json:"deleted,omitempty"`
	// Time of tree deletion, if any.
	// Readonly.
	DeleteTime *google_protobuf1.Timestamp `protobuf:"bytes,20,opt,name=delete_time,json=deleteTime" json:"delete_time,omitempty"`
}

func (m *Tree) Reset()                    { *m = Tree{} }
func (m *Tree) String() string            { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()               {}
func (*Tree) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Tree) GetTreeId() int64 {
	if m != nil {
		return m.TreeId
	}
	return 0
}

func (m *Tree) GetTreeState() TreeState {
	if m != nil {
		return m.TreeState
	}
	return TreeState_UNKNOWN_TREE_STATE
}

func (m *Tree) GetTreeType() TreeType {
	if m != nil {
		return m.TreeType
	}
	return TreeType_UNKNOWN_TREE_TYPE
}

func (m *Tree) GetHashStrategy() HashStrategy {
	if m != nil {
		return m.HashStrategy
	}
	return HashStrategy_UNKNOWN_HASH_STRATEGY
}

func (m *Tree) GetHashAlgorithm() sigpb.DigitallySigned_HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return sigpb.DigitallySigned_NONE
}

func (m *Tree) GetSignatureAlgorithm() sigpb.DigitallySigned_SignatureAlgorithm {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return sigpb.DigitallySigned_ANONYMOUS
}

func (m *Tree) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Tree) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tree) GetPrivateKey() *google_protobuf2.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *Tree) GetStorageSettings() *google_protobuf2.Any {
	if m != nil {
		return m.StorageSettings
	}
	return nil
}

func (m *Tree) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Tree) GetMaxRootDuration() *google_protobuf3.Duration {
	if m != nil {
		return m.MaxRootDuration
	}
	return nil
}

func (m *Tree) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Tree) GetUpdateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Tree) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Tree) GetDeleteTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.DeleteTime
	}
	return nil
}

type SignedEntryTimestamp struct {
	TimestampNanos int64                  `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	LogId          int64                  `protobuf:"varint,2,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	Signature      *sigpb.DigitallySigned `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (m *SignedEntryTimestamp) Reset()                    { *m = SignedEntryTimestamp{} }
func (m *SignedEntryTimestamp) String() string            { return proto.CompactTextString(m) }
func (*SignedEntryTimestamp) ProtoMessage()               {}
func (*SignedEntryTimestamp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SignedEntryTimestamp) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedEntryTimestamp) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *SignedEntryTimestamp) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

// SignedLogRoot represents a commitment by a Log to a particular tree.
type SignedLogRoot struct {
	// Deprecated: TimestampNanos moved to LogRoot.
	TimestampNanos int64 `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	// Deprecated: RootHash moved to LogRoot.
	RootHash []byte `protobuf:"bytes,2,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	// Deprecated: TreeSize moved to LogRoot.
	TreeSize int64 `protobuf:"varint,3,opt,name=tree_size,json=treeSize" json:"tree_size,omitempty"`
	// Deprecated TreeRevision moved to LogRoot.
	TreeRevision int64 `protobuf:"varint,6,opt,name=tree_revision,json=treeRevision" json:"tree_revision,omitempty"`
	// key_hint is a hint to identify the public key for signature verification.
	// key_hint is not authenticated and may be incorrect or missing, in which
	// case all known public keys may be used to verify the signature.
	// When directly communicating with a Trillian gRPC server, the key_hint will
	// typically contain the LogID encoded as a big-endian 64-bit integer;
	// however, in other contexts the key_hint is likely to have different
	// contents (e.g. it could be a GUID, a URL + TreeID, or it could be
	// derived from the public key itself).
	KeyHint []byte `protobuf:"bytes,7,opt,name=key_hint,json=keyHint,proto3" json:"key_hint,omitempty"`
	// log_root holds the TLS-serialization of the following
	// structure (described in RFC5246 notation):
	// enum { v1(1), (65535)} Version;
	// struct {
	//   uint64 tree_size;
	//   opaque root_hash<0..128>;
	//   uint64 timestamp_nanos;
	//   uint64 revision;
	//   opaque metadata<0..65535>;
	// } LogRootV1;
	// struct {
	//   Version version;
	//   select(version) {
	//     case v1: LogRootV1;
	//   }
	// } LogRoot;
	LogRoot []byte `protobuf:"bytes,8,opt,name=log_root,json=logRoot,proto3" json:"log_root,omitempty"`
	// LogRootSignature is the signature over LogRoot.
	LogRootSignature []byte `protobuf:"bytes,9,opt,name=log_root_signature,json=logRootSignature,proto3" json:"log_root_signature,omitempty"`
}

func (m *SignedLogRoot) Reset()                    { *m = SignedLogRoot{} }
func (m *SignedLogRoot) String() string            { return proto.CompactTextString(m) }
func (*SignedLogRoot) ProtoMessage()               {}
func (*SignedLogRoot) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *SignedLogRoot) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedLogRoot) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *SignedLogRoot) GetTreeSize() int64 {
	if m != nil {
		return m.TreeSize
	}
	return 0
}

func (m *SignedLogRoot) GetTreeRevision() int64 {
	if m != nil {
		return m.TreeRevision
	}
	return 0
}

func (m *SignedLogRoot) GetKeyHint() []byte {
	if m != nil {
		return m.KeyHint
	}
	return nil
}

func (m *SignedLogRoot) GetLogRoot() []byte {
	if m != nil {
		return m.LogRoot
	}
	return nil
}

func (m *SignedLogRoot) GetLogRootSignature() []byte {
	if m != nil {
		return m.LogRootSignature
	}
	return nil
}

// SignedMapRoot represents a commitment by a Map to a particular tree.
type SignedMapRoot struct {
	// map_root holds the TLS-serialization of the following
	// structure (described in RFC5246 notation):
	// enum { v1(1), (65535)} Version;
	// struct {
	//   opaque root_hash<0..128>;
	//   uint64 timestamp_nanos;
	//   uint64 revision;
	//   opaque metadata<0..65535>;
	// } MapRootV1;
	// struct {
	//   Version version;
	//   select(version) {
	//     case v1: MapRootV1;
	//   }
	// } MapRoot;
	MapRoot []byte `protobuf:"bytes,9,opt,name=map_root,json=mapRoot,proto3" json:"map_root,omitempty"`
	// Signature is the signature over MapRoot.
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedMapRoot) Reset()                    { *m = SignedMapRoot{} }
func (m *SignedMapRoot) String() string            { return proto.CompactTextString(m) }
func (*SignedMapRoot) ProtoMessage()               {}
func (*SignedMapRoot) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *SignedMapRoot) GetMapRoot() []byte {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

func (m *SignedMapRoot) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Tree)(nil), "trillian.Tree")
	proto.RegisterType((*SignedEntryTimestamp)(nil), "trillian.SignedEntryTimestamp")
	proto.RegisterType((*SignedLogRoot)(nil), "trillian.SignedLogRoot")
	proto.RegisterType((*SignedMapRoot)(nil), "trillian.SignedMapRoot")
	proto.RegisterEnum("trillian.LogRootFormat", LogRootFormat_name, LogRootFormat_value)
	proto.RegisterEnum("trillian.MapRootFormat", MapRootFormat_name, MapRootFormat_value)
	proto.RegisterEnum("trillian.HashStrategy", HashStrategy_name, HashStrategy_value)
	proto.RegisterEnum("trillian.TreeState", TreeState_name, TreeState_value)
	proto.RegisterEnum("trillian.TreeType", TreeType_name, TreeType_value)
}

func init() { proto.RegisterFile("trillian.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1097 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdb, 0x6f, 0xdb, 0x36,
	0x17, 0xaf, 0x6c, 0xc5, 0x96, 0xe9, 0x4b, 0x18, 0xa6, 0x17, 0xd9, 0xfd, 0xf0, 0xd5, 0xcb, 0x06,
	0x2c, 0x2b, 0x06, 0x67, 0xf5, 0xd6, 0x02, 0x43, 0x1f, 0x06, 0x35, 0x56, 0x62, 0x3b, 0x89, 0x6d,
	0x50, 0x5a, 0x87, 0xe6, 0x85, 0x90, 0x6d, 0x4e, 0x16, 0xa2, 0x1b, 0x24, 0xba, 0xa8, 0xfa, 0x3c,
	0x60, 0x18, 0xb6, 0x3f, 0x7a, 0x20, 0x45, 0xd9, 0x89, 0xb3, 0xae, 0x7b, 0x49, 0x78, 0xce, 0xef,
	0xc2, 0x23, 0xf2, 0x1c, 0xc9, 0xa0, 0xc5, 0x12, 0xcf, 0xf7, 0x3d, 0x27, 0xec, 0xc5, 0x49, 0xc4,
	0x22, 0xa4, 0x15, 0x71, 0xa7, 0xb3, 0x48, 0xb2, 0x98, 0x45, 0x27, 0x37, 0x34, 0x4b, 0xe3, 0xb9,
	0xfc, 0x97, 0xb3, 0x3a, 0xba, 0xc4, 0x52, 0xcf, 0x8d, 0xe7, 0xf9, 0x5f, 0x89, 0xb4, 0xdd, 0x28,
	0x72, 0x7d, 0x7a, 0x22, 0xa2, 0xf9, 0xfa, 0xd7, 0x13, 0x27, 0xcc, 0x24, 0xf4, 0xff, 0x5d, 0x68,
	0xb9, 0x4e, 0x1c, 0xe6, 0x45, 0x72, 0xeb, 0xce, 0xb3, 0x5d, 0x9c, 0x79, 0x01, 0x4d, 0x99, 0x13,
	0xc4, 0x39, 0xe1, 0xe8, 0x8f, 0x2a, 0x50, 0xed, 0x84, 0x52, 0xf4, 0x04, 0x54, 0x59, 0x42, 0x29,
	0xf1, 0x96, 0xba, 0xd2, 0x55, 0x8e, 0xcb, 0xb8, 0xc2, 0xc3, 0xd1, 0x12, 0xf5, 0x01, 0x10, 0x40,
	0xca, 0x1c, 0x46, 0xf5, 0x52, 0x57, 0x39, 0x6e, 0xf5, 0x0f, 0x7b, 0x9b, 0x47, 0xe4, 0x62, 0x8b,
	0x43, 0xb8, 0xc6, 0x8a, 0x25, 0x3a, 0x01, 0x22, 0x20, 0x2c, 0x8b, 0xa9, 0x5e, 0x16, 0x12, 0x74,
	0x57, 0x62, 0x67, 0x31, 0xc5, 0x1a, 0x93, 0x2b, 0xf4, 0x1a, 0x34, 0x57, 0x4e, 0xba, 0x22, 0x29,
	0x4b, 0x1c, 0x46, 0xdd, 0x4c, 0x57, 0x85, 0xe8, 0xf1, 0x56, 0x34, 0x74, 0xd2, 0x95, 0x25, 0x51,
	0xdc, 0x58, 0xdd, 0x8a, 0xd0, 0x05, 0x68, 0x09, 0xb1, 0xe3, 0xbb, 0x51, 0xe2, 0xb1, 0x55, 0xa0,
	0xef, 0x09, 0xf5, 0x57, 0xbd, 0xfc, 0x14, 0x07, 0x9e, 0xeb, 0x31, 0xc7, 0xf7, 0x33, 0xcb, 0x73,
	0x43, 0xba, 0x14, 0x56, 0x46, 0xc1, 0xc5, 0x62, 0xe3, 0x4d, 0x88, 0xae, 0xc1, 0x61, 0xea, 0xb9,
	0xa1, 0xc3, 0xd6, 0x09, 0xbd, 0xe5, 0x58, 0x11, 0x8e, 0xdf, 0x7c, 0xc2, 0xd1, 0x2a, 0x14, 0x5b,
	0x5b, 0x94, 0xde, 0xcb, 0xa1, 0x2f, 0x40, 0x63, 0xe9, 0xa5, 0xb1, 0xef, 0x64, 0x24, 0x74, 0x02,
	0xaa, 0x6b, 0x5d, 0xe5, 0xb8, 0x86, 0xeb, 0x32, 0x37, 0x71, 0x02, 0x8a, 0xba, 0xa0, 0xbe, 0xa4,
	0xe9, 0x22, 0xf1, 0x62, 0x7e, 0x8b, 0x7a, 0x4d, 0x32, 0xb6, 0x29, 0xf4, 0x12, 0xd4, 0xe3, 0xc4,
	0x7b, 0xef, 0x30, 0x4a, 0x6e, 0x68, 0xa6, 0x37, 0xba, 0xca, 0x71, 0xbd, 0xff, 0xb0, 0x97, 0x5f,
	0x74, 0xaf, 0xb8, 0xe8, 0x9e, 0x11, 0x66, 0x18, 0x48, 0xe2, 0x05, 0xcd, 0xd0, 0x4f, 0x00, 0xa6,
	0x2c, 0x4a, 0x1c, 0x97, 0x92, 0x94, 0x32, 0xe6, 0x85, 0x6e, 0xaa, 0x37, 0xff, 0x45, 0xbb, 0x2f,
	0xd9, 0x96, 0x24, 0xa3, 0xef, 0x00, 0x88, 0xd7, 0x73, 0xdf, 0x5b, 0x88, 0x6d, 0x5b, 0x42, 0x7a,
	0xd0, 0x93, 0x2d, 0x3c, 0x13, 0xc8, 0x05, 0xcd, 0x70, 0x2d, 0x2e, 0x96, 0xc8, 0x04, 0x07, 0x81,
	0xf3, 0x81, 0x24, 0x51, 0xc4, 0x48, 0xd1, 0x97, 0xfa, 0xbe, 0x10, 0xb6, 0xef, 0xed, 0x39, 0x90,
	0x04, 0xbc, 0x1f, 0x38, 0x1f, 0x70, 0x14, 0xb1, 0x22, 0x81, 0x5e, 0x83, 0xfa, 0x22, 0xa1, 0xfc,
	0x79, 0x79, 0xf3, 0xea, 0x50, 0x18, 0x74, 0xee, 0x19, 0xd8, 0x45, 0x67, 0x63, 0x90, 0xd3, 0x79,
	0x82, 0x8b, 0xd7, 0xf1, 0x72, 0x23, 0x3e, 0xf8, 0xbc, 0x38, 0xa7, 0x0b, 0xb1, 0x0e, 0xaa, 0x4b,
	0xea, 0x53, 0x46, 0x97, 0xfa, 0x61, 0x57, 0x39, 0xd6, 0x70, 0x11, 0x72, 0xdb, 0x7c, 0x99, 0xdb,
	0x3e, 0xfc, 0xbc, 0x6d, 0x4e, 0xe7, 0x89, 0xb1, 0xaa, 0x21, 0x78, 0x38, 0x56, 0xb5, 0x2a, 0xd4,
	0xc6, 0xaa, 0x06, 0x60, 0x7d, 0xac, 0x6a, 0x75, 0xd8, 0x38, 0xfa, 0x4b, 0x01, 0x0f, 0xf3, 0x86,
	0x32, 0x43, 0x96, 0x64, 0x1b, 0x31, 0xfa, 0x1a, 0xec, 0x6f, 0xe6, 0x96, 0x84, 0x4e, 0x18, 0xa5,
	0x72, 0x46, 0x5b, 0x9b, 0xf4, 0x84, 0x67, 0xd1, 0x23, 0x50, 0xf1, 0x23, 0x97, 0xcf, 0x70, 0x49,
	0xe0, 0x7b, 0x7e, 0xe4, 0x8e, 0x96, 0xe8, 0x07, 0x50, 0xdb, 0x74, 0xa3, 0x18, 0xc7, 0x7a, 0xff,
	0xf1, 0x3f, 0x77, 0x32, 0xde, 0x12, 0x8f, 0x7e, 0x2f, 0x81, 0x66, 0x9e, 0xbd, 0x8c, 0x5c, 0x7e,
	0x23, 0xff, 0xbd, 0x8e, 0xa7, 0xa0, 0x26, 0x6e, 0x9d, 0x8f, 0x96, 0x28, 0xa5, 0x81, 0x35, 0x9e,
	0xe0, 0x93, 0xc7, 0xc1, 0xfc, 0x85, 0xe2, 0x7d, 0xcc, 0xab, 0x29, 0xe7, 0x2f, 0x02, 0xcb, 0xfb,
	0x48, 0xd1, 0x97, 0xa0, 0x29, 0xc0, 0x84, 0xbe, 0xf7, 0x52, 0xde, 0x2f, 0x15, 0x41, 0x68, 0xf0,
	0x24, 0x96, 0x39, 0xd4, 0x06, 0xda, 0x0d, 0xcd, 0xc8, 0xca, 0x0b, 0x99, 0x5e, 0x15, 0xee, 0xd5,
	0x1b, 0x9a, 0x0d, 0xbd, 0x90, 0x71, 0x88, 0x9f, 0x00, 0xdf, 0x4c, 0x8c, 0x57, 0x03, 0x57, 0x7d,
	0x59, 0xfd, 0xb7, 0x00, 0x15, 0x10, 0xd9, 0x1e, 0x47, 0x4d, 0x90, 0xa0, 0x24, 0x6d, 0x06, 0x79,
	0xac, 0x6a, 0x2a, 0xdc, 0x1b, 0xab, 0xda, 0x1e, 0xac, 0x1c, 0x25, 0xc5, 0x41, 0x5c, 0x39, 0xb1,
	0xb0, 0x6a, 0x03, 0x2d, 0x70, 0xe2, 0x7c, 0x97, 0xdc, 0xa0, 0x1a, 0x48, 0xe8, 0x7f, 0xb7, 0xcf,
	0x5a, 0x15, 0xd8, 0x36, 0x31, 0x56, 0x35, 0x05, 0x96, 0xc6, 0xaa, 0x56, 0x82, 0xe5, 0xb1, 0xaa,
	0x95, 0xa1, 0x9a, 0xef, 0x30, 0x56, 0xb5, 0x0a, 0xac, 0x6e, 0x5a, 0x42, 0x83, 0xb5, 0xe7, 0x03,
	0xd0, 0x94, 0xc7, 0x7e, 0x16, 0x25, 0x81, 0xc3, 0xd0, 0x53, 0xf0, 0xe4, 0x72, 0x7a, 0x4e, 0xf0,
	0x74, 0x6a, 0x93, 0xb3, 0x29, 0xbe, 0x32, 0x6c, 0xf2, 0xf3, 0xe4, 0x62, 0x32, 0xfd, 0x65, 0x02,
	0x1f, 0xa0, 0xc7, 0x00, 0xed, 0x82, 0x6f, 0x5f, 0x40, 0x85, 0xbb, 0xc8, 0x9a, 0xb7, 0x2e, 0x57,
	0xc6, 0xec, 0xd3, 0x2e, 0xbb, 0xa0, 0x70, 0xf9, 0x4d, 0x01, 0x8d, 0xdb, 0xef, 0x5f, 0xd4, 0x06,
	0x8f, 0xa4, 0x8a, 0x0c, 0x0d, 0x6b, 0x48, 0x2c, 0x1b, 0x1b, 0xb6, 0x79, 0xfe, 0x0e, 0x3e, 0x40,
	0x08, 0xb4, 0xf0, 0xd9, 0xe9, 0xab, 0x1f, 0x5f, 0xf5, 0x89, 0x35, 0x34, 0xfa, 0x2f, 0x5f, 0x41,
	0x05, 0x1d, 0x82, 0x7d, 0xdb, 0xb4, 0x6c, 0xc2, 0xcd, 0x39, 0xdf, 0xc4, 0xb0, 0xc4, 0x3d, 0xa6,
	0x6f, 0xc6, 0xe6, 0xa9, 0x4d, 0x76, 0xf8, 0x65, 0xf4, 0x08, 0x1c, 0x9c, 0x4e, 0x27, 0xa3, 0x0b,
	0x8b, 0xa7, 0x5e, 0xbe, 0xe8, 0x13, 0x9e, 0x56, 0x9f, 0xff, 0xa9, 0x80, 0xda, 0xe6, 0x73, 0xc3,
	0x8b, 0x2d, 0x6a, 0xb0, 0xb1, 0x69, 0x12, 0xcb, 0x36, 0x6c, 0x13, 0x3e, 0x40, 0x00, 0x54, 0x8c,
	0x53, 0x7b, 0xf4, 0xd6, 0x84, 0x0a, 0x5f, 0x9f, 0xe1, 0xe9, 0xb5, 0x39, 0x81, 0x25, 0xf4, 0x0c,
	0x3c, 0x19, 0x98, 0x33, 0x6c, 0x9e, 0x1a, 0xb6, 0x39, 0x20, 0xd6, 0xf4, 0xcc, 0x26, 0x03, 0xf3,
	0xd2, 0xb4, 0xcd, 0x01, 0x2c, 0x77, 0x4a, 0x9a, 0xb2, 0x43, 0x18, 0x1a, 0x78, 0xb0, 0x21, 0xa8,
	0x82, 0xd0, 0x00, 0xda, 0x00, 0x1b, 0xa3, 0xc9, 0x68, 0x72, 0x0e, 0xf7, 0x9e, 0x9f, 0x03, 0xad,
	0xf8, 0x90, 0xf1, 0x82, 0xef, 0xd4, 0x62, 0xbf, 0x9b, 0xf1, 0x52, 0xaa, 0xa0, 0x7c, 0x39, 0x3d,
	0x87, 0x0a, 0x5f, 0x5c, 0x19, 0x33, 0x58, 0xe2, 0xa7, 0x33, 0xc3, 0xe6, 0x14, 0x0f, 0x4c, 0x6c,
	0x0e, 0x08, 0x07, 0xcb, 0x6f, 0x86, 0xa0, 0xbd, 0x88, 0x82, 0xe2, 0xdd, 0x71, 0xf7, 0xb7, 0xc3,
	0x9b, 0xa6, 0x2d, 0xe3, 0x19, 0x0f, 0x67, 0xca, 0x75, 0xc7, 0xf5, 0xd8, 0x6a, 0x3d, 0xef, 0x2d,
	0xa2, 0xe0, 0x44, 0x7e, 0xdc, 0x0b, 0xc9, 0xbc, 0x22, 0x34, 0xdf, 0xff, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x19, 0x4e, 0xf0, 0x51, 0x81, 0x08, 0x00, 0x00,
}
