// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: getambassador.io/v2/Host.proto

package getambassador_io_v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	v11 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HostTLSCertificateSource int32

const (
	HostTLSCertificateSource_Unknown HostTLSCertificateSource = 0
	HostTLSCertificateSource_None    HostTLSCertificateSource = 1
	HostTLSCertificateSource_Other   HostTLSCertificateSource = 2
	HostTLSCertificateSource_ACME    HostTLSCertificateSource = 3
)

var HostTLSCertificateSource_name = map[int32]string{
	0: "Unknown",
	1: "None",
	2: "Other",
	3: "ACME",
}

var HostTLSCertificateSource_value = map[string]int32{
	"Unknown": 0,
	"None":    1,
	"Other":   2,
	"ACME":    3,
}

func (x HostTLSCertificateSource) String() string {
	return proto.EnumName(HostTLSCertificateSource_name, int32(x))
}

func (HostTLSCertificateSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_689b3ac2a9b9e947, []int{0}
}

type HostState int32

const (
	// The default value is the "zero" value, and it would be great if
	// "Pending" could be the default value; but it's Important that the
	// "zero" value be able to be shown as empty/omitted from display,
	// and we really do want `kubectl get hosts` to say "Pending" in the
	// "STATE" column, and not leave the column empty.
	HostState_Initial HostState = 0
	HostState_Pending HostState = 1
	HostState_Ready   HostState = 2
	HostState_Error   HostState = 3
)

var HostState_name = map[int32]string{
	0: "Initial",
	1: "Pending",
	2: "Ready",
	3: "Error",
}

var HostState_value = map[string]int32{
	"Initial": 0,
	"Pending": 1,
	"Ready":   2,
	"Error":   3,
}

func (x HostState) String() string {
	return proto.EnumName(HostState_name, int32(x))
}

func (HostState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_689b3ac2a9b9e947, []int{1}
}

type HostPhase int32

const (
	HostPhase_NA                        HostPhase = 0
	HostPhase_DefaultsFilled            HostPhase = 1
	HostPhase_ACMEUserPrivateKeyCreated HostPhase = 2
	HostPhase_ACMEUserRegistered        HostPhase = 3
	HostPhase_ACMECertificateChallenge  HostPhase = 4
)

var HostPhase_name = map[int32]string{
	0: "NA",
	1: "DefaultsFilled",
	2: "ACMEUserPrivateKeyCreated",
	3: "ACMEUserRegistered",
	4: "ACMECertificateChallenge",
}

var HostPhase_value = map[string]int32{
	"NA":                        0,
	"DefaultsFilled":            1,
	"ACMEUserPrivateKeyCreated": 2,
	"ACMEUserRegistered":        3,
	"ACMECertificateChallenge":  4,
}

func (x HostPhase) String() string {
	return proto.EnumName(HostPhase_name, int32(x))
}

func (HostPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_689b3ac2a9b9e947, []int{2}
}

// The Host resource will usually be a Kubernetes CRD, but it could
// appear in other forms. The HostSpec is the part of the Host resource
// that doesn't change, no matter what form it's in -- when it's a CRD,
// this is the part in the "spec" dictionary.
type HostSpec struct {
	// Common to all Ambassador objects (and optional).
	AmbassadorId []string `protobuf:"bytes,1,rep,name=ambassador_id,json=ambassadorId,proto3" json:"ambassador_id,omitempty"`
	// Common to all Ambassador objects (and optional).
	Generation int32 `protobuf:"varint,2,opt,name=generation,proto3" json:"generation,omitempty"`
	// Hostname by which the Ambassador can be reached.
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Selector by which we can find further configuration. Defaults to hostname=$hostname
	Selector *v1.LabelSelector `protobuf:"bytes,4,opt,name=selector,proto3" json:"selector,omitempty"`
	// Specifies who to talk ACME with to get certs. Defaults to Let's Encrypt; if "none", do
	// not try to do TLS for this Host.
	AcmeProvider *ACMEProviderSpec `protobuf:"bytes,5,opt,name=acmeProvider,proto3" json:"acmeProvider,omitempty"`
	// Name of the Kubernetes secret into which to save generated certificates. Defaults
	// to $hostname
	TlsSecret            *v11.LocalObjectReference `protobuf:"bytes,6,opt,name=tlsSecret,proto3" json:"tlsSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *HostSpec) Reset()         { *m = HostSpec{} }
func (m *HostSpec) String() string { return proto.CompactTextString(m) }
func (*HostSpec) ProtoMessage()    {}
func (*HostSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_689b3ac2a9b9e947, []int{0}
}
func (m *HostSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostSpec.Merge(m, src)
}
func (m *HostSpec) XXX_Size() int {
	return m.Size()
}
func (m *HostSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_HostSpec.DiscardUnknown(m)
}

var xxx_messageInfo_HostSpec proto.InternalMessageInfo

func (m *HostSpec) GetAmbassadorId() []string {
	if m != nil {
		return m.AmbassadorId
	}
	return nil
}

func (m *HostSpec) GetGeneration() int32 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *HostSpec) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostSpec) GetSelector() *v1.LabelSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *HostSpec) GetAcmeProvider() *ACMEProviderSpec {
	if m != nil {
		return m.AcmeProvider
	}
	return nil
}

func (m *HostSpec) GetTlsSecret() *v11.LocalObjectReference {
	if m != nil {
		return m.TlsSecret
	}
	return nil
}

type HostStatus struct {
	TlsCertificateSource HostTLSCertificateSource `protobuf:"varint,1,opt,name=tlsCertificateSource,proto3,enum=getambassador.io.v2.HostTLSCertificateSource" json:"tlsCertificateSource,omitempty"`
	State                HostState                `protobuf:"varint,2,opt,name=state,proto3,enum=getambassador.io.v2.HostState" json:"state,omitempty"`
	// phaseCompleted and phasePending are valid when state==Pending or
	// state==Error.
	PhaseCompleted HostPhase `protobuf:"varint,3,opt,name=phaseCompleted,proto3,enum=getambassador.io.v2.HostPhase" json:"phaseCompleted,omitempty"`
	PhasePending   HostPhase `protobuf:"varint,4,opt,name=phasePending,proto3,enum=getambassador.io.v2.HostPhase" json:"phasePending,omitempty"`
	// errorReason is valid when state==Error.
	ErrorReason          string   `protobuf:"bytes,5,opt,name=errorReason,proto3" json:"errorReason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_689b3ac2a9b9e947, []int{1}
}
func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(m, src)
}
func (m *HostStatus) XXX_Size() int {
	return m.Size()
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetTlsCertificateSource() HostTLSCertificateSource {
	if m != nil {
		return m.TlsCertificateSource
	}
	return HostTLSCertificateSource_Unknown
}

func (m *HostStatus) GetState() HostState {
	if m != nil {
		return m.State
	}
	return HostState_Initial
}

func (m *HostStatus) GetPhaseCompleted() HostPhase {
	if m != nil {
		return m.PhaseCompleted
	}
	return HostPhase_NA
}

func (m *HostStatus) GetPhasePending() HostPhase {
	if m != nil {
		return m.PhasePending
	}
	return HostPhase_NA
}

func (m *HostStatus) GetErrorReason() string {
	if m != nil {
		return m.ErrorReason
	}
	return ""
}

type ACMEProviderSpec struct {
	// Specifies who to talk ACME with to get certs. Defaults to Let's
	// Encrypt; if "none", do not try to do TLS for this Host.
	Authority        string                    `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Email            string                    `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PrivateKeySecret *v11.LocalObjectReference `protobuf:"bytes,3,opt,name=privateKeySecret,proto3" json:"privateKeySecret,omitempty"`
	// This is normally set automatically
	Registration         string   `protobuf:"bytes,4,opt,name=registration,proto3" json:"registration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACMEProviderSpec) Reset()         { *m = ACMEProviderSpec{} }
func (m *ACMEProviderSpec) String() string { return proto.CompactTextString(m) }
func (*ACMEProviderSpec) ProtoMessage()    {}
func (*ACMEProviderSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_689b3ac2a9b9e947, []int{2}
}
func (m *ACMEProviderSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ACMEProviderSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ACMEProviderSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ACMEProviderSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACMEProviderSpec.Merge(m, src)
}
func (m *ACMEProviderSpec) XXX_Size() int {
	return m.Size()
}
func (m *ACMEProviderSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ACMEProviderSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ACMEProviderSpec proto.InternalMessageInfo

func (m *ACMEProviderSpec) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *ACMEProviderSpec) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ACMEProviderSpec) GetPrivateKeySecret() *v11.LocalObjectReference {
	if m != nil {
		return m.PrivateKeySecret
	}
	return nil
}

func (m *ACMEProviderSpec) GetRegistration() string {
	if m != nil {
		return m.Registration
	}
	return ""
}

func init() {
	proto.RegisterEnum("getambassador.io.v2.HostTLSCertificateSource", HostTLSCertificateSource_name, HostTLSCertificateSource_value)
	proto.RegisterEnum("getambassador.io.v2.HostState", HostState_name, HostState_value)
	proto.RegisterEnum("getambassador.io.v2.HostPhase", HostPhase_name, HostPhase_value)
	proto.RegisterType((*HostSpec)(nil), "getambassador.io.v2.HostSpec")
	proto.RegisterType((*HostStatus)(nil), "getambassador.io.v2.HostStatus")
	proto.RegisterType((*ACMEProviderSpec)(nil), "getambassador.io.v2.ACMEProviderSpec")
}

func init() { proto.RegisterFile("getambassador.io/v2/Host.proto", fileDescriptor_689b3ac2a9b9e947) }

var fileDescriptor_689b3ac2a9b9e947 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6b, 0xeb, 0x46,
	0x14, 0x8d, 0xfc, 0x91, 0x5a, 0x37, 0xae, 0x11, 0xd3, 0x50, 0xd4, 0x90, 0x1a, 0xe3, 0x52, 0x30,
	0x81, 0x48, 0xc4, 0xc9, 0xa2, 0xd0, 0x55, 0xe2, 0xc6, 0x34, 0x34, 0x4d, 0xcc, 0x38, 0x59, 0x97,
	0xb1, 0x74, 0x2d, 0x4f, 0x2d, 0x69, 0xc4, 0xcc, 0xd8, 0xc5, 0xd0, 0x1f, 0xd7, 0x65, 0x77, 0x2d,
	0xf4, 0x0f, 0x94, 0xfc, 0x92, 0xc7, 0x8c, 0x1d, 0xdb, 0x49, 0x9c, 0xf7, 0x78, 0x3b, 0xdd, 0x33,
	0xf7, 0x9c, 0xb9, 0xf7, 0x1c, 0x34, 0xd0, 0x4c, 0x50, 0xb3, 0x6c, 0xc4, 0x94, 0x62, 0xb1, 0x90,
	0x01, 0x17, 0xe1, 0xbc, 0x1b, 0xfe, 0x2c, 0x94, 0x0e, 0x0a, 0x29, 0xb4, 0x20, 0x5f, 0xbd, 0x3e,
	0x0f, 0xe6, 0xdd, 0xa3, 0x8b, 0xe9, 0x0f, 0xca, 0xb4, 0xb2, 0x82, 0x67, 0x2c, 0x9a, 0xf0, 0x1c,
	0xe5, 0x22, 0x2c, 0xa6, 0x89, 0x01, 0x54, 0x98, 0xa1, 0x66, 0xe1, 0xfc, 0x2c, 0x4c, 0x30, 0x47,
	0xc9, 0x34, 0xc6, 0x4b, 0xa9, 0xa3, 0xf6, 0x86, 0x15, 0x46, 0x42, 0xe2, 0xae, 0x9e, 0xd3, 0x84,
	0xeb, 0xc9, 0x6c, 0x14, 0x44, 0x22, 0x0b, 0x13, 0x91, 0x88, 0xd0, 0xc2, 0xa3, 0xd9, 0xd8, 0x56,
	0xb6, 0xb0, 0x5f, 0xcb, 0xf6, 0xf6, 0x3f, 0x25, 0xa8, 0x99, 0x61, 0x87, 0x05, 0x46, 0xe4, 0x3b,
	0xf8, 0x72, 0x33, 0xe9, 0x6f, 0x3c, 0xf6, 0x9d, 0x56, 0xb9, 0xe3, 0xd2, 0xfa, 0x06, 0xbc, 0x89,
	0x49, 0x13, 0x60, 0x75, 0x27, 0x17, 0xb9, 0x5f, 0x6a, 0x39, 0x9d, 0x2a, 0xdd, 0x42, 0xc8, 0x11,
	0xd4, 0x26, 0x42, 0xe9, 0x9c, 0x65, 0xe8, 0x97, 0x5b, 0x4e, 0xc7, 0xa5, 0xeb, 0x9a, 0xdc, 0x43,
	0x4d, 0x61, 0x8a, 0x91, 0x16, 0xd2, 0xaf, 0xb4, 0x9c, 0xce, 0x41, 0xf7, 0x3c, 0x58, 0xee, 0x14,
	0x6c, 0x3b, 0x11, 0x14, 0xd3, 0xc4, 0x00, 0x2a, 0x30, 0x4e, 0x04, 0xf3, 0xb3, 0xe0, 0x96, 0x8d,
	0x30, 0x1d, 0xae, 0xa8, 0x74, 0x2d, 0x42, 0x6e, 0xa0, 0xce, 0xa2, 0x0c, 0x07, 0x52, 0xcc, 0x79,
	0x8c, 0xd2, 0xaf, 0x5a, 0xd1, 0xef, 0x83, 0x1d, 0x9e, 0x07, 0x97, 0xbd, 0x5f, 0xaf, 0x9f, 0x1b,
	0xcd, 0xba, 0xf4, 0x05, 0x95, 0xf4, 0xc1, 0xd5, 0xa9, 0x1a, 0x62, 0x24, 0x51, 0xfb, 0xfb, 0x56,
	0xa7, 0xb3, 0x35, 0x5c, 0x60, 0x0c, 0xb7, 0xa3, 0x88, 0x88, 0xa5, 0xf7, 0xa3, 0xdf, 0x31, 0xd2,
	0x14, 0xc7, 0x28, 0x31, 0x8f, 0x90, 0x6e, 0xa8, 0xed, 0xff, 0x4a, 0x00, 0xd6, 0x51, 0xcd, 0xf4,
	0x4c, 0x11, 0x06, 0x87, 0x3a, 0x55, 0x3d, 0x94, 0x9a, 0x8f, 0x79, 0xc4, 0x34, 0x0e, 0xc5, 0x4c,
	0x46, 0xe8, 0x3b, 0x2d, 0xa7, 0xd3, 0xe8, 0x9e, 0xee, 0x9c, 0xd4, 0xd0, 0x1f, 0x6e, 0x87, 0x6f,
	0x48, 0x74, 0xa7, 0x14, 0xb9, 0x80, 0xaa, 0xd2, 0x4c, 0xa3, 0x0d, 0xa3, 0xd1, 0x6d, 0xbe, 0xab,
	0x69, 0x46, 0x42, 0xba, 0x6c, 0x26, 0x7d, 0x68, 0x14, 0x13, 0xa6, 0xb0, 0x27, 0xb2, 0x22, 0x45,
	0x8d, 0xb1, 0x4d, 0xeb, 0x63, 0xf4, 0x81, 0x69, 0xa7, 0xaf, 0x58, 0xe4, 0x0a, 0xea, 0x16, 0x19,
	0x60, 0x1e, 0xf3, 0x3c, 0xb1, 0xb9, 0x7e, 0x5a, 0xe5, 0x05, 0x87, 0xb4, 0xe0, 0x00, 0xa5, 0x14,
	0x92, 0x22, 0x53, 0x22, 0xb7, 0x29, 0xba, 0x74, 0x1b, 0x6a, 0xff, 0xe5, 0x80, 0xf7, 0x3a, 0x40,
	0x72, 0x0c, 0x2e, 0x9b, 0xe9, 0x89, 0x90, 0x5c, 0x2f, 0xac, 0xa1, 0x2e, 0xdd, 0x00, 0xe4, 0x10,
	0xaa, 0x98, 0x31, 0x9e, 0x5a, 0x5b, 0x5c, 0xba, 0x2c, 0xc8, 0x03, 0x78, 0x85, 0xe4, 0x73, 0xa6,
	0xf1, 0x17, 0x5c, 0xac, 0xd2, 0x2e, 0x7f, 0x66, 0xda, 0x6f, 0x14, 0x48, 0x1b, 0xea, 0x12, 0x13,
	0xae, 0xf4, 0xea, 0xb7, 0xa8, 0xd8, 0x2b, 0x5f, 0x60, 0x27, 0x7d, 0xf0, 0xdf, 0x0b, 0x96, 0x1c,
	0xc0, 0x17, 0x8f, 0xf9, 0x34, 0x17, 0x7f, 0xe4, 0xde, 0x1e, 0xa9, 0x41, 0xe5, 0x4e, 0xe4, 0xe8,
	0x39, 0xc4, 0x85, 0xea, 0xbd, 0x9e, 0xa0, 0xf4, 0x4a, 0x06, 0x34, 0xfb, 0x7b, 0xe5, 0x93, 0x1f,
	0xc1, 0x5d, 0x87, 0x69, 0x88, 0x37, 0x39, 0xd7, 0x9c, 0xa5, 0xde, 0x9e, 0x29, 0x56, 0x8e, 0x2e,
	0xb9, 0x14, 0x59, 0xbc, 0xf0, 0x4a, 0xe6, 0xf3, 0xda, 0x78, 0xe9, 0x95, 0x4f, 0xfe, 0x5c, 0x92,
	0x6d, 0x08, 0x64, 0x1f, 0x4a, 0x77, 0x97, 0xde, 0x1e, 0x21, 0xd0, 0xf8, 0x09, 0xc7, 0x6c, 0x96,
	0x6a, 0xd5, 0xe7, 0x69, 0x8a, 0xb1, 0xe7, 0x90, 0x6f, 0xe1, 0x1b, 0x73, 0xdf, 0xa3, 0x42, 0x39,
	0x58, 0x6f, 0xdb, 0x93, 0x68, 0x9e, 0x1a, 0xaf, 0x44, 0xbe, 0x06, 0xf2, 0x7c, 0x4c, 0xed, 0x92,
	0x28, 0x31, 0xf6, 0xca, 0xe4, 0x18, 0x7c, 0x83, 0x6f, 0x6d, 0xd8, 0x9b, 0xb0, 0x34, 0xc5, 0x3c,
	0x41, 0xaf, 0x72, 0x55, 0xff, 0xfb, 0xa9, 0xe9, 0xfc, 0xfb, 0xd4, 0x74, 0xfe, 0x7f, 0x6a, 0x3a,
	0xa3, 0x7d, 0xfb, 0x04, 0x9d, 0x7f, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x3f, 0x48, 0xb6, 0x42,
	0x05, 0x00, 0x00,
}

func (m *HostSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TlsSecret != nil {
		{
			size, err := m.TlsSecret.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.AcmeProvider != nil {
		{
			size, err := m.AcmeProvider.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Selector != nil {
		{
			size, err := m.Selector.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hostname) > 0 {
		i -= len(m.Hostname)
		copy(dAtA[i:], m.Hostname)
		i = encodeVarintHost(dAtA, i, uint64(len(m.Hostname)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Generation != 0 {
		i = encodeVarintHost(dAtA, i, uint64(m.Generation))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AmbassadorId) > 0 {
		for iNdEx := len(m.AmbassadorId) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AmbassadorId[iNdEx])
			copy(dAtA[i:], m.AmbassadorId[iNdEx])
			i = encodeVarintHost(dAtA, i, uint64(len(m.AmbassadorId[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *HostStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ErrorReason) > 0 {
		i -= len(m.ErrorReason)
		copy(dAtA[i:], m.ErrorReason)
		i = encodeVarintHost(dAtA, i, uint64(len(m.ErrorReason)))
		i--
		dAtA[i] = 0x2a
	}
	if m.PhasePending != 0 {
		i = encodeVarintHost(dAtA, i, uint64(m.PhasePending))
		i--
		dAtA[i] = 0x20
	}
	if m.PhaseCompleted != 0 {
		i = encodeVarintHost(dAtA, i, uint64(m.PhaseCompleted))
		i--
		dAtA[i] = 0x18
	}
	if m.State != 0 {
		i = encodeVarintHost(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if m.TlsCertificateSource != 0 {
		i = encodeVarintHost(dAtA, i, uint64(m.TlsCertificateSource))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ACMEProviderSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ACMEProviderSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ACMEProviderSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Registration) > 0 {
		i -= len(m.Registration)
		copy(dAtA[i:], m.Registration)
		i = encodeVarintHost(dAtA, i, uint64(len(m.Registration)))
		i--
		dAtA[i] = 0x22
	}
	if m.PrivateKeySecret != nil {
		{
			size, err := m.PrivateKeySecret.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintHost(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintHost(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHost(dAtA []byte, offset int, v uint64) int {
	offset -= sovHost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AmbassadorId) > 0 {
		for _, s := range m.AmbassadorId {
			l = len(s)
			n += 1 + l + sovHost(uint64(l))
		}
	}
	if m.Generation != 0 {
		n += 1 + sovHost(uint64(m.Generation))
	}
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	if m.AcmeProvider != nil {
		l = m.AcmeProvider.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	if m.TlsSecret != nil {
		l = m.TlsSecret.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HostStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TlsCertificateSource != 0 {
		n += 1 + sovHost(uint64(m.TlsCertificateSource))
	}
	if m.State != 0 {
		n += 1 + sovHost(uint64(m.State))
	}
	if m.PhaseCompleted != 0 {
		n += 1 + sovHost(uint64(m.PhaseCompleted))
	}
	if m.PhasePending != 0 {
		n += 1 + sovHost(uint64(m.PhasePending))
	}
	l = len(m.ErrorReason)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ACMEProviderSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	if m.PrivateKeySecret != nil {
		l = m.PrivateKeySecret.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.Registration)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHost(x uint64) (n int) {
	return sovHost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HostSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmbassadorId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmbassadorId = append(m.AmbassadorId, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Generation", wireType)
			}
			m.Generation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Generation |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &v1.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcmeProvider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AcmeProvider == nil {
				m.AcmeProvider = &ACMEProviderSpec{}
			}
			if err := m.AcmeProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TlsSecret", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TlsSecret == nil {
				m.TlsSecret = &v11.LocalObjectReference{}
			}
			if err := m.TlsSecret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HostStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HostStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TlsCertificateSource", wireType)
			}
			m.TlsCertificateSource = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TlsCertificateSource |= HostTLSCertificateSource(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= HostState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhaseCompleted", wireType)
			}
			m.PhaseCompleted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PhaseCompleted |= HostPhase(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhasePending", wireType)
			}
			m.PhasePending = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PhasePending |= HostPhase(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ACMEProviderSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ACMEProviderSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ACMEProviderSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivateKeySecret", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrivateKeySecret == nil {
				m.PrivateKeySecret = &v11.LocalObjectReference{}
			}
			if err := m.PrivateKeySecret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHost
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHost
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHost
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHost
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHost
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHost
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHost(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHost
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHost = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHost   = fmt.Errorf("proto: integer overflow")
)
