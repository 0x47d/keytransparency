// Code generated by protoc-gen-go.
// source: proto/core/core.proto
// DO NOT EDIT!

/*
Package google_security_e2ekeys_core is a generated protocol buffer package.

It is generated from these files:
	proto/core/core.proto

It has these top-level messages:
	EntryStorage
	EpochInfo
	DomainInfo
	VerifierInfo
*/
package google_security_e2ekeys_core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_security_e2ekeys_v2 "github.com/google/e2e-key-server/proto/v2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// EntryStorage is what gets serialized to disk / log.
type EntryStorage struct {
	// commitment_timestamp is a sequential update number of the SignedEntryUpdate.
	CommitmentTimestamp uint64 `protobuf:"varint,1,opt,name=commitment_timestamp" json:"commitment_timestamp,omitempty"`
	// entry_update contains a SignedEntryUpdate proto with the actual update
	// contents.
	SignedEntryUpdate *google_security_e2ekeys_v2.SignedEntryUpdate `protobuf:"bytes,2,opt,name=signed_entry_update" json:"signed_entry_update,omitempty"`
	// profile is the serialized protobuf Profile.
	// profile is private and must not be released to verifiers.
	Profile []byte `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	// commitment_key is at least 16 random bytes.
	CommitmentKey []byte `protobuf:"bytes,4,opt,name=commitment_key,proto3" json:"commitment_key,omitempty"`
	// domain allows servers to keep track of multiple trees at once.
	Domain string `protobuf:"bytes,5,opt,name=domain" json:"domain,omitempty"`
}

func (m *EntryStorage) Reset()                    { *m = EntryStorage{} }
func (m *EntryStorage) String() string            { return proto.CompactTextString(m) }
func (*EntryStorage) ProtoMessage()               {}
func (*EntryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EntryStorage) GetSignedEntryUpdate() *google_security_e2ekeys_v2.SignedEntryUpdate {
	if m != nil {
		return m.SignedEntryUpdate
	}
	return nil
}

// EpochInfo is what gets serialized to disk / log.
type EpochInfo struct {
	// signed_epoch_head is the signed epoch head of the created epoch.
	SignedEpochHead *google_security_e2ekeys_v2.SignedEpochHead `protobuf:"bytes,1,opt,name=signed_epoch_head" json:"signed_epoch_head,omitempty"`
	// last_commitment_timestamp is the timestamp of the last update included in
	// created epoch.
	LastCommitmentTimestamp uint64 `protobuf:"varint,2,opt,name=last_commitment_timestamp" json:"last_commitment_timestamp,omitempty"`
}

func (m *EpochInfo) Reset()                    { *m = EpochInfo{} }
func (m *EpochInfo) String() string            { return proto.CompactTextString(m) }
func (*EpochInfo) ProtoMessage()               {}
func (*EpochInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EpochInfo) GetSignedEpochHead() *google_security_e2ekeys_v2.SignedEpochHead {
	if m != nil {
		return m.SignedEpochHead
	}
	return nil
}

// DomainInfo is the information that need to be baked into an application
// in order to verify information for a domain.
type DomainInfo struct {
	PublicKeys []*google_security_e2ekeys_v2.PublicKey `protobuf:"bytes,1,rep,name=public_keys" json:"public_keys,omitempty"`
	// required_sigs is the number of valid signatures to require out of
	// public_keys before considering signed_tree heads legitimate.
	RequiredSigs uint32 `protobuf:"varint,2,opt,name=required_sigs" json:"required_sigs,omitempty"`
	// domain is the doman suffix to use when resolving a user_id to a domain.
	Domain string `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	// api_url is the url prefix to use when querying users on this domain.
	ApiUrl string `protobuf:"bytes,4,opt,name=api_url" json:"api_url,omitempty"`
}

func (m *DomainInfo) Reset()                    { *m = DomainInfo{} }
func (m *DomainInfo) String() string            { return proto.CompactTextString(m) }
func (*DomainInfo) ProtoMessage()               {}
func (*DomainInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DomainInfo) GetPublicKeys() []*google_security_e2ekeys_v2.PublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type VerifierInfo struct {
	PublicKeys []*google_security_e2ekeys_v2.PublicKey `protobuf:"bytes,1,rep,name=public_keys" json:"public_keys,omitempty"`
	// required_sigs is the number of valid signatures to require out of
	// public_keys before considering signed_tree heads legitimate.
	RequiredSigs uint32 `protobuf:"varint,2,opt,name=required_sigs" json:"required_sigs,omitempty"`
	// api_url is the url prefix to use when querying users on this domain.
	ApiUrl string `protobuf:"bytes,4,opt,name=api_url" json:"api_url,omitempty"`
	// domain is the doman suffix that this verifier is responsible for.
	Domain string `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
}

func (m *VerifierInfo) Reset()                    { *m = VerifierInfo{} }
func (m *VerifierInfo) String() string            { return proto.CompactTextString(m) }
func (*VerifierInfo) ProtoMessage()               {}
func (*VerifierInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *VerifierInfo) GetPublicKeys() []*google_security_e2ekeys_v2.PublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*EntryStorage)(nil), "google.security.e2ekeys.core.EntryStorage")
	proto.RegisterType((*EpochInfo)(nil), "google.security.e2ekeys.core.EpochInfo")
	proto.RegisterType((*DomainInfo)(nil), "google.security.e2ekeys.core.DomainInfo")
	proto.RegisterType((*VerifierInfo)(nil), "google.security.e2ekeys.core.VerifierInfo")
}

var fileDescriptor0 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xad, 0x95, 0x6c, 0xff, 0x88, 0xab, 0x95, 0x28, 0x3d, 0x68, 0x40, 0x10, 0xa4,
	0x09, 0xc4, 0x83, 0xe0, 0xd9, 0x8a, 0x7f, 0x2e, 0x42, 0xd1, 0x6b, 0x48, 0x93, 0x69, 0xba, 0x98,
	0x64, 0xe3, 0xee, 0x26, 0x90, 0x9b, 0xf8, 0x8d, 0xfc, 0x86, 0xce, 0x6e, 0x29, 0x94, 0xa2, 0xe2,
	0xc5, 0x4b, 0x60, 0x66, 0xdf, 0xdb, 0xf7, 0x9b, 0xc9, 0x92, 0x61, 0x29, 0xb8, 0xe2, 0x7e, 0xcc,
	0x05, 0x98, 0x8f, 0x67, 0x6a, 0x3a, 0x4a, 0x39, 0x4f, 0x33, 0xf0, 0x24, 0xc4, 0x95, 0x60, 0xaa,
	0xf1, 0x20, 0x80, 0x57, 0x68, 0xa4, 0xa7, 0x35, 0xc7, 0x57, 0x29, 0x53, 0x8b, 0x6a, 0x86, 0x45,
	0xee, 0x2f, 0x85, 0x3e, 0x9e, 0x8f, 0x51, 0x30, 0x96, 0x20, 0x6a, 0x10, 0xfe, 0xf2, 0xd6, 0x3a,
	0xf0, 0x57, 0x3e, 0xd3, 0x70, 0x3f, 0x2d, 0xd2, 0x9b, 0x14, 0x4a, 0x34, 0x53, 0xc5, 0x45, 0x94,
	0x02, 0x1d, 0x91, 0x03, 0xbc, 0x24, 0x67, 0x2a, 0x87, 0x42, 0x85, 0x8a, 0xe5, 0x20, 0x55, 0x94,
	0x97, 0x8e, 0x75, 0x62, 0x9d, 0xb7, 0xe9, 0x03, 0xd9, 0x97, 0x2c, 0x2d, 0x20, 0x09, 0x41, 0x9b,
	0xc2, 0xaa, 0x4c, 0x22, 0x05, 0xce, 0x16, 0x1e, 0x76, 0x83, 0xb1, 0xf7, 0x13, 0x63, 0x1d, 0x78,
	0x53, 0x63, 0x33, 0x51, 0xcf, 0xc6, 0x44, 0x77, 0xc9, 0x0e, 0x32, 0xcc, 0x59, 0x06, 0x4e, 0x0b,
	0xfd, 0x3d, 0x7a, 0x48, 0x06, 0x6b, 0xd1, 0x68, 0x74, 0xda, 0xa6, 0x3f, 0x20, 0x9d, 0x84, 0xe7,
	0x11, 0x2b, 0x9c, 0x6d, 0xac, 0x6d, 0xb7, 0x26, 0xf6, 0xa4, 0xe4, 0xf1, 0xe2, 0xbe, 0x98, 0x73,
	0x7a, 0x4b, 0xf6, 0x56, 0x44, 0xba, 0x17, 0x2e, 0x20, 0x4a, 0x0c, 0x6c, 0x37, 0xb8, 0xf8, 0x03,
	0x8f, 0xf6, 0xdc, 0xa1, 0x85, 0x9e, 0x92, 0xa3, 0x2c, 0x92, 0x2a, 0xfc, 0x76, 0x78, 0x3d, 0x5f,
	0xdb, 0x7d, 0xb7, 0x08, 0xb9, 0x31, 0x20, 0x26, 0xf9, 0x9a, 0x74, 0xcb, 0x6a, 0x96, 0xb1, 0x58,
	0xa3, 0x4a, 0xcc, 0x6c, 0x61, 0xe6, 0xd9, 0x6f, 0x99, 0x4f, 0x46, 0xfe, 0x08, 0x0d, 0x1d, 0x92,
	0xbe, 0x80, 0xb7, 0x8a, 0x09, 0xe4, 0x46, 0x7c, 0x69, 0x12, 0xfa, 0x6b, 0x93, 0xea, 0x8d, 0xd8,
	0x7a, 0x45, 0x51, 0xc9, 0xc2, 0x4a, 0x64, 0x66, 0x15, 0xb6, 0xfb, 0x81, 0xbf, 0xeb, 0x05, 0x04,
	0x9b, 0x33, 0x10, 0xff, 0x05, 0xb1, 0x19, 0xba, 0x49, 0x35, 0xeb, 0x98, 0xa7, 0x73, 0xf9, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xc7, 0x2c, 0x4c, 0x23, 0xaa, 0x02, 0x00, 0x00,
}
