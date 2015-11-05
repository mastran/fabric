// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package pbft is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Unpack
	PBFT
	Request
	RequestHashes
	Requests
	PBFTArray
*/
package pbft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Unpack_Type int32

const (
	Unpack_UNDEFINED      Unpack_Type = 0
	Unpack_REQUEST        Unpack_Type = 1
	Unpack_PRE_PREPARE    Unpack_Type = 2
	Unpack_PREPARE        Unpack_Type = 3
	Unpack_COMMIT         Unpack_Type = 4
	Unpack_PREPARE_RESULT Unpack_Type = 5
	Unpack_COMMIT_RESULT  Unpack_Type = 6
)

var Unpack_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "REQUEST",
	2: "PRE_PREPARE",
	3: "PREPARE",
	4: "COMMIT",
	5: "PREPARE_RESULT",
	6: "COMMIT_RESULT",
}
var Unpack_Type_value = map[string]int32{
	"UNDEFINED":      0,
	"REQUEST":        1,
	"PRE_PREPARE":    2,
	"PREPARE":        3,
	"COMMIT":         4,
	"PREPARE_RESULT": 5,
	"COMMIT_RESULT":  6,
}

func (x Unpack_Type) String() string {
	return proto.EnumName(Unpack_Type_name, int32(x))
}

type PBFT_Type int32

const (
	PBFT_UNDEFINED      PBFT_Type = 0
	PBFT_REQUEST        PBFT_Type = 1
	PBFT_PRE_PREPARE    PBFT_Type = 2
	PBFT_PREPARE        PBFT_Type = 3
	PBFT_COMMIT         PBFT_Type = 4
	PBFT_PREPARE_RESULT PBFT_Type = 5
	PBFT_COMMIT_RESULT  PBFT_Type = 6
)

var PBFT_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "REQUEST",
	2: "PRE_PREPARE",
	3: "PREPARE",
	4: "COMMIT",
	5: "PREPARE_RESULT",
	6: "COMMIT_RESULT",
}
var PBFT_Type_value = map[string]int32{
	"UNDEFINED":      0,
	"REQUEST":        1,
	"PRE_PREPARE":    2,
	"PREPARE":        3,
	"COMMIT":         4,
	"PREPARE_RESULT": 5,
	"COMMIT_RESULT":  6,
}

func (x PBFT_Type) String() string {
	return proto.EnumName(PBFT_Type_name, int32(x))
}

type Unpack struct {
	Type    Unpack_Type `protobuf:"varint,1,opt,name=type,enum=pbft.Unpack_Type" json:"type,omitempty"`
	Payload []byte      `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Unpack) Reset()         { *m = Unpack{} }
func (m *Unpack) String() string { return proto.CompactTextString(m) }
func (*Unpack) ProtoMessage()    {}

type PBFT struct {
	Type    PBFT_Type `protobuf:"varint,1,opt,name=type,enum=pbft.PBFT_Type" json:"type,omitempty"`
	ID      string    `protobuf:"bytes,2,opt,name=ID" json:"ID,omitempty"`
	Payload []byte    `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *PBFT) Reset()         { *m = PBFT{} }
func (m *PBFT) String() string { return proto.CompactTextString(m) }
func (*PBFT) ProtoMessage()    {}

type Request struct {
	ID      string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type RequestHashes struct {
	Hashes []string `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *RequestHashes) Reset()         { *m = RequestHashes{} }
func (m *RequestHashes) String() string { return proto.CompactTextString(m) }
func (*RequestHashes) ProtoMessage()    {}

type Requests struct {
	Requests []*Request `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
}

func (m *Requests) Reset()         { *m = Requests{} }
func (m *Requests) String() string { return proto.CompactTextString(m) }
func (*Requests) ProtoMessage()    {}

func (m *Requests) GetRequests() []*Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

type PBFTArray struct {
	Pbfts []*PBFT `protobuf:"bytes,1,rep,name=pbfts" json:"pbfts,omitempty"`
}

func (m *PBFTArray) Reset()         { *m = PBFTArray{} }
func (m *PBFTArray) String() string { return proto.CompactTextString(m) }
func (*PBFTArray) ProtoMessage()    {}

func (m *PBFTArray) GetPbfts() []*PBFT {
	if m != nil {
		return m.Pbfts
	}
	return nil
}

func init() {
	proto.RegisterEnum("pbft.Unpack_Type", Unpack_Type_name, Unpack_Type_value)
	proto.RegisterEnum("pbft.PBFT_Type", PBFT_Type_name, PBFT_Type_value)
}
