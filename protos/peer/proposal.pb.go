// Code generated by protoc-gen-go.
// source: peer/proposal.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This structure is necessary to sign the proposal which contains the header
// and the payload. Without this structure, we would have to concatenate the
// header and the payload to verify the signature, which could be expensive
// with large payload
//
// When an endorser receives a SignedProposal message, it should verify the
// signature over the proposal bytes. This verification requires the following
// steps:
// 1. Verification of the validity of the certificate that was used to produce
//    the signature.  The certificate will be available once proposalBytes has
//    been unmarshalled to a Proposal message, and Proposal.header has been
//    unmarshalled to a Header message. While this unmarshalling-before-verifying
//    might not be ideal, it is unavoidable because i) the signature needs to also
//    protect the signing certificate; ii) it is desirable that Header is created
//    once by the client and never changed (for the sake of accountability and
//    non-repudiation). Note also that it is actually impossible to conclusively
//    verify the validity of the certificate included in a Proposal, because the
//    proposal needs to first be endorsed and ordered with respect to certificate
//    expiration transactions. Still, it is useful to pre-filter expired
//    certificates at this stage.
// 2. Verification that the certificate is trusted (signed by a trusted CA) and
//    that it is allowed to transact with us (with respect to some ACLs);
// 3. Verification that the signature on proposalBytes is valid;
// 4. Detect replay attacks;
type SignedProposal struct {
	// The bytes of Proposal
	ProposalBytes []byte `protobuf:"bytes,1,opt,name=proposal_bytes,json=proposalBytes,proto3" json:"proposal_bytes,omitempty"`
	// Signaure over proposalBytes; this signature is to be verified against
	// the creator identity contained in the header of the Proposal message
	// marshaled as proposalBytes
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedProposal) Reset()                    { *m = SignedProposal{} }
func (m *SignedProposal) String() string            { return proto.CompactTextString(m) }
func (*SignedProposal) ProtoMessage()               {}
func (*SignedProposal) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

// A Proposal is sent to an endorser for endorsement.  The proposal contains:
// 1. A header which should be unmarshaled to a Header message.  Note that
//    Header is both the header of a Proposal and of a Transaction, in that i)
//    both headers should be unmarshaled to this message; and ii) it is used to
//    compute cryptographic hashes and signatures.  The header has fields common
//    to all proposals/transactions.  In addition it has a type field for
//    additional customization. An example of this is the ChaincodeHeaderExtension
//    message used to extend the Header for type CHAINCODE.
// 2. A payload whose type depends on the header's type field.
// 3. An extension whose type depends on the header's type field.
//
// Let us see an example. For type CHAINCODE (see the Header message),
// we have the following:
// 1. The header is a Header message whose extensions field is a
//    ChaincodeHeaderExtension message.
// 2. The payload is a ChaincodeProposalPayload message.
// 3. The extension is a ChaincodeAction that might be used to ask the
//    endorsers to endorse a specific ChaincodeAction, thus emulating the
//    submitting peer model.
type Proposal struct {
	// The header of the proposal. It is the bytes of the Header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the proposal as defined by the type in the proposal
	// header.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional extensions to the proposal. Its content depends on the Header's
	// type field.  For the type CHAINCODE, it might be the bytes of a
	// ChaincodeAction message.
	Extension []byte `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *Proposal) Reset()                    { *m = Proposal{} }
func (m *Proposal) String() string            { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()               {}
func (*Proposal) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

// ChaincodeHeaderExtension is the Header's extentions message to be used when
// the Header's type is CHAINCODE.  This extensions is used to specify which
// chaincode to invoke and what should appear on the ledger.
type ChaincodeHeaderExtension struct {
	// The PayloadVisibility field controls to what extent the Proposal's payload
	// (recall that for the type CHAINCODE, it is ChaincodeProposalPayload
	// message) field will be visible in the final transaction and in the ledger.
	// Ideally, it would be configurable, supporting at least 3 main “visibility
	// modes”:
	// 1. all bytes of the payload are visible;
	// 2. only a hash of the payload is visible;
	// 3. nothing is visible.
	// Notice that the visibility function may be potentially part of the ESCC.
	// In that case it overrides PayloadVisibility field.  Finally notice that
	// this field impacts the content of ProposalResponsePayload.proposalHash.
	PayloadVisibility []byte `protobuf:"bytes,1,opt,name=payload_visibility,json=payloadVisibility,proto3" json:"payload_visibility,omitempty"`
	// The ID of the chaincode to target.
	ChaincodeID *ChaincodeID `protobuf:"bytes,2,opt,name=chaincodeID" json:"chaincodeID,omitempty"`
}

func (m *ChaincodeHeaderExtension) Reset()                    { *m = ChaincodeHeaderExtension{} }
func (m *ChaincodeHeaderExtension) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeHeaderExtension) ProtoMessage()               {}
func (*ChaincodeHeaderExtension) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *ChaincodeHeaderExtension) GetChaincodeID() *ChaincodeID {
	if m != nil {
		return m.ChaincodeID
	}
	return nil
}

// ChaincodeProposalPayload is the Proposal's payload message to be used when
// the Header's type is CHAINCODE.  It contains the arguments for this
// invocation.
type ChaincodeProposalPayload struct {
	// Input contains the arguments for this invocation. If this invocation
	// deploys a new chaincode, ESCC/VSCC are part of this field.
	Input []byte `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// Transient contains data (e.g. cryptographic material) that might be used
	// to implement some form of application-level confidentiality. The contents
	// of this field are supposed to always be omitted from the transaction and
	// excluded from the ledger.
	Transient []byte `protobuf:"bytes,2,opt,name=transient,proto3" json:"transient,omitempty"`
}

func (m *ChaincodeProposalPayload) Reset()                    { *m = ChaincodeProposalPayload{} }
func (m *ChaincodeProposalPayload) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeProposalPayload) ProtoMessage()               {}
func (*ChaincodeProposalPayload) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

// ChaincodeAction contains the actions the events generated by the execution
// of the chaincode.
type ChaincodeAction struct {
	// This field contains the read set and the write set produced by the
	// chaincode executing this invocation.
	Results []byte `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	// This field contains the events generated by the chaincode executing this
	// invocation.
	Events []byte `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
	// This field contains the result of executing this invocation.
	Response *Response `protobuf:"bytes,3,opt,name=response" json:"response,omitempty"`
}

func (m *ChaincodeAction) Reset()                    { *m = ChaincodeAction{} }
func (m *ChaincodeAction) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeAction) ProtoMessage()               {}
func (*ChaincodeAction) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *ChaincodeAction) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedProposal)(nil), "protos.SignedProposal")
	proto.RegisterType((*Proposal)(nil), "protos.Proposal")
	proto.RegisterType((*ChaincodeHeaderExtension)(nil), "protos.ChaincodeHeaderExtension")
	proto.RegisterType((*ChaincodeProposalPayload)(nil), "protos.ChaincodeProposalPayload")
	proto.RegisterType((*ChaincodeAction)(nil), "protos.ChaincodeAction")
}

func init() { proto.RegisterFile("peer/proposal.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x52, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x65, 0xdf, 0x87, 0xdb, 0xcc, 0x74, 0x6a, 0x36, 0xa4, 0x8c, 0x3d, 0x48, 0x41, 0x50, 0xd4,
	0x16, 0x26, 0xfe, 0x00, 0x37, 0x05, 0x7d, 0x91, 0x51, 0xd1, 0x87, 0xbd, 0x8c, 0xb4, 0xbd, 0xb6,
	0x81, 0x9a, 0xc4, 0x24, 0x1d, 0xf6, 0xcd, 0x9f, 0x2e, 0x6d, 0x93, 0x6c, 0x3e, 0x95, 0x7b, 0xce,
	0xc9, 0xc9, 0x3d, 0xa7, 0x41, 0x23, 0x01, 0x20, 0x43, 0x21, 0xb9, 0xe0, 0x8a, 0x14, 0x81, 0x90,
	0x5c, 0x73, 0xdc, 0x6d, 0x3e, 0x6a, 0x32, 0x6e, 0xc8, 0x24, 0x27, 0x94, 0x25, 0x3c, 0x85, 0x96,
	0x9d, 0x4c, 0xff, 0x1c, 0x59, 0x4b, 0x50, 0x82, 0x33, 0x65, 0x58, 0xff, 0x0d, 0x0d, 0x5f, 0x69,
	0xc6, 0x20, 0x5d, 0x1a, 0x01, 0x3e, 0x47, 0x43, 0x27, 0x8e, 0x2b, 0x0d, 0xca, 0xeb, 0x9c, 0x75,
	0x2e, 0x0e, 0xa2, 0x43, 0x8b, 0xce, 0x6b, 0x10, 0x4f, 0xd1, 0xbe, 0xa2, 0x19, 0x23, 0xba, 0x94,
	0xe0, 0xfd, 0x6b, 0x14, 0x5b, 0xc0, 0x5f, 0xa1, 0xbe, 0x33, 0x3c, 0x45, 0xdd, 0x1c, 0x48, 0x0a,
	0xd2, 0x18, 0x99, 0x09, 0x7b, 0xa8, 0x27, 0x48, 0x55, 0x70, 0x92, 0x9a, 0xf3, 0x76, 0xac, 0xbd,
	0xe1, 0x5b, 0x03, 0x53, 0x94, 0x33, 0xef, 0x7f, 0xeb, 0xed, 0x00, 0xff, 0xa7, 0x83, 0xbc, 0x85,
	0x0d, 0xf9, 0xd4, 0x78, 0x3d, 0x5a, 0x12, 0xdf, 0x20, 0x6c, 0x5c, 0xd6, 0x1b, 0xaa, 0x68, 0x4c,
	0x0b, 0xaa, 0x2b, 0x73, 0xf1, 0x89, 0x61, 0xde, 0x1d, 0x81, 0xef, 0xd0, 0xc0, 0xf5, 0xf5, 0xfc,
	0xd0, 0xec, 0x31, 0x98, 0x8d, 0xda, 0x6e, 0x54, 0xb0, 0xd8, 0x52, 0xd1, 0xae, 0xce, 0x7f, 0xd9,
	0xd9, 0xc0, 0xe6, 0x5c, 0x9a, 0xe5, 0xc7, 0x68, 0x8f, 0x32, 0x51, 0x6a, 0x73, 0x69, 0x3b, 0xd4,
	0x91, 0xb4, 0x24, 0x4c, 0x51, 0x60, 0xda, 0xd6, 0xe5, 0x00, 0xff, 0x0b, 0x1d, 0x39, 0xbf, 0xfb,
	0x44, 0xd7, 0x41, 0x3c, 0xd4, 0x93, 0xa0, 0xca, 0x42, 0xdb, 0xfe, 0xed, 0x58, 0xf7, 0x09, 0x1b,
	0x60, 0x5a, 0x19, 0x1f, 0x33, 0xe1, 0x6b, 0xd4, 0xb7, 0x3f, 0xb7, 0x29, 0x6d, 0x30, 0x3b, 0xb6,
	0x41, 0x22, 0x83, 0x47, 0x4e, 0x31, 0xbf, 0x5a, 0x5d, 0x66, 0x54, 0xe7, 0x65, 0x1c, 0x24, 0xfc,
	0x33, 0xcc, 0x2b, 0x01, 0xb2, 0x80, 0x34, 0x03, 0x19, 0x7e, 0x90, 0x58, 0xd2, 0x24, 0x6c, 0x8f,
	0x86, 0xf5, 0xeb, 0x89, 0xdb, 0x17, 0x76, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x0b, 0x8f,
	0x64, 0x7f, 0x02, 0x00, 0x00,
}
