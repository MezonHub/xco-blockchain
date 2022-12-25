// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/token/v1beta1/token.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_ixofoundation_ixo_blockchain_x_iid_types "github.com/xcohub/xco-blockchain/x/iid/types"
	io "io"
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

type ContractType int32

const (
	ContractType_CW20    ContractType = 0
	ContractType_CW721   ContractType = 1
	ContractType_XCO1155 ContractType = 2
)

var ContractType_name = map[int32]string{
	0: "CW20",
	1: "CW721",
	2: "XCO1155",
}

var ContractType_value = map[string]int32{
	"CW20":    0,
	"CW721":   1,
	"XCO1155": 2,
}

func (x ContractType) String() string {
	return proto.EnumName(ContractType_name, int32(x))
}

func (ContractType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5772c07022bd1ddf, []int{0}
}

type Contract struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address" yaml:"address"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_5772c07022bd1ddf, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Contract) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Params struct {
	Cw20ContractCode    string `protobuf:"bytes,1,opt,name=cw20ContractCode,proto3" json:"cw20_contract_code" yaml:"cw20_contract_code"`
	Cw721ContractCode   string `protobuf:"bytes,2,opt,name=cw721ContractCode,proto3" json:"cw721_contract_code" yaml:"cw721_contract_code"`
	Xco1155ContractCode string `protobuf:"bytes,3,opt,name=ixo1155ContractCode,proto3" json:"ixo1155_contract_code" yaml:"ixo1155_contract_code"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5772c07022bd1ddf, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type TokenMinter struct {
	MinterDid       github_com_ixofoundation_ixo_blockchain_x_iid_types.DIDFragment `protobuf:"bytes,1,opt,name=minterDid,proto3,casttype=github.com/xcohub/xco-blockchain/x/iid/types.DIDFragment" json:"minter_did" yaml:"minter_did"`
	MinterAddress   string                                                          `protobuf:"bytes,2,opt,name=minterAddress,proto3" json:"minter_address" yaml:"minter_address"`
	ContractAddress string                                                          `protobuf:"bytes,3,opt,name=contractAddress,proto3" json:"contract_address" yaml:"contract_address"`
	ContractType    ContractType                                                    `protobuf:"varint,4,opt,name=contractType,proto3,enum=ixo.token.v1beta1.ContractType" json:"contract_type" yaml:"contract_type"`
	Name            string                                                          `protobuf:"bytes,5,opt,name=name,proto3" json:"name" yaml:"name"`
	Description     string                                                          `protobuf:"bytes,6,opt,name=description,proto3" json:"description" yaml:"description"`
}

func (m *TokenMinter) Reset()         { *m = TokenMinter{} }
func (m *TokenMinter) String() string { return proto.CompactTextString(m) }
func (*TokenMinter) ProtoMessage()    {}
func (*TokenMinter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5772c07022bd1ddf, []int{2}
}
func (m *TokenMinter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenMinter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenMinter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenMinter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenMinter.Merge(m, src)
}
func (m *TokenMinter) XXX_Size() int {
	return m.Size()
}
func (m *TokenMinter) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenMinter.DiscardUnknown(m)
}

var xxx_messageInfo_TokenMinter proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ixo.token.v1beta1.ContractType", ContractType_name, ContractType_value)
	proto.RegisterType((*Contract)(nil), "ixo.token.v1beta1.Contract")
	proto.RegisterType((*Params)(nil), "ixo.token.v1beta1.Params")
	proto.RegisterType((*TokenMinter)(nil), "ixo.token.v1beta1.TokenMinter")
}

func init() { proto.RegisterFile("ixo/token/v1beta1/token.proto", fileDescriptor_5772c07022bd1ddf) }

var fileDescriptor_5772c07022bd1ddf = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x3f, 0x4f, 0xdb, 0x4e,
	0x18, 0x8e, 0xfd, 0x0b, 0x7f, 0x72, 0x01, 0x7e, 0xe1, 0x00, 0x91, 0xa2, 0xe2, 0x43, 0x57, 0x55,
	0xa2, 0x45, 0xb5, 0x71, 0x50, 0x84, 0xca, 0x52, 0x95, 0xa0, 0x56, 0x0c, 0x15, 0xad, 0x85, 0x44,
	0xdb, 0x25, 0x72, 0x7c, 0x26, 0x9c, 0xc0, 0xbe, 0xc8, 0x31, 0x25, 0xec, 0x1d, 0x3a, 0xf6, 0x23,
	0xf4, 0x13, 0xf4, 0x3b, 0x74, 0xeb, 0xc8, 0xd8, 0xe9, 0x54, 0xc1, 0xe6, 0xd1, 0x63, 0xa7, 0xca,
	0x77, 0x76, 0x12, 0xc7, 0x0c, 0xdd, 0xfc, 0x3e, 0xcf, 0x7b, 0xcf, 0x73, 0x79, 0xdf, 0x27, 0x07,
	0xd6, 0xe9, 0x80, 0x19, 0x21, 0x3b, 0x77, 0x7d, 0xe3, 0x93, 0xd9, 0x71, 0x43, 0xdb, 0x94, 0x95,
	0xde, 0x0b, 0x58, 0xc8, 0xe0, 0x22, 0x1d, 0x30, 0x5d, 0x02, 0x29, 0xbd, 0xb6, 0xdc, 0x65, 0x5d,
	0x26, 0x58, 0x23, 0xf9, 0x92, 0x8d, 0xf8, 0x0c, 0xcc, 0xb6, 0x98, 0x1f, 0x06, 0xb6, 0x13, 0xc2,
	0x47, 0x40, 0xa5, 0xa4, 0xae, 0x6c, 0x28, 0x9b, 0x95, 0xfd, 0xa5, 0x88, 0x23, 0x95, 0x92, 0x98,
	0xa3, 0xca, 0xb5, 0xed, 0x5d, 0xec, 0x61, 0x4a, 0xb0, 0xa5, 0x52, 0x02, 0x77, 0xc1, 0x8c, 0x4d,
	0x48, 0xe0, 0xf6, 0xfb, 0x75, 0x55, 0x74, 0xae, 0x47, 0x1c, 0x65, 0x50, 0xcc, 0xd1, 0x82, 0x6c,
	0x4f, 0x01, 0x6c, 0x65, 0x14, 0xfe, 0xa1, 0x82, 0xe9, 0xb7, 0x76, 0x60, 0x7b, 0x7d, 0xd8, 0x06,
	0x35, 0xe7, 0xaa, 0xb1, 0x9d, 0x19, 0xb7, 0x18, 0x71, 0x53, 0xdb, 0x9d, 0x88, 0x23, 0x98, 0x70,
	0x6d, 0x27, 0x25, 0xdb, 0x0e, 0x23, 0x6e, 0xcc, 0xd1, 0x03, 0xa9, 0x5b, 0xe4, 0xb0, 0x55, 0x10,
	0x83, 0x0e, 0x58, 0x74, 0xae, 0x76, 0x1b, 0x66, 0xce, 0x41, 0x5e, 0xb7, 0x19, 0x71, 0xb4, 0x24,
	0xc8, 0x82, 0xc5, 0x5a, 0x66, 0x51, 0x20, 0xb1, 0x55, 0xd4, 0x83, 0xe7, 0x60, 0x89, 0x0e, 0x98,
	0x69, 0x36, 0x9b, 0x39, 0x9b, 0xff, 0x84, 0xcd, 0xf3, 0x88, 0xa3, 0x95, 0x94, 0x2e, 0x18, 0x3d,
	0x4c, 0x47, 0x7a, 0x1f, 0x8d, 0xad, 0xfb, 0x54, 0xf7, 0xca, 0x5f, 0xbe, 0xa1, 0x12, 0xfe, 0x5e,
	0x06, 0xd5, 0xe3, 0x64, 0xab, 0x6f, 0xa8, 0x1f, 0xba, 0x01, 0xfc, 0xac, 0x80, 0x8a, 0x27, 0x3e,
	0x0f, 0x86, 0x9b, 0x3b, 0x8d, 0x38, 0x02, 0x12, 0x6c, 0x13, 0xb1, 0xc1, 0x45, 0x69, 0x37, 0xc2,
	0xf0, 0x1f, 0x8e, 0x5e, 0x74, 0x69, 0x78, 0x76, 0xd9, 0xd1, 0x1d, 0xe6, 0x19, 0x74, 0xc0, 0x4e,
	0xd9, 0xa5, 0x4f, 0xec, 0x90, 0x32, 0x3f, 0xa9, 0x9e, 0x75, 0x2e, 0x98, 0x73, 0xee, 0x9c, 0xd9,
	0xd4, 0x37, 0x06, 0x06, 0xa5, 0xc4, 0x08, 0xaf, 0x7b, 0x6e, 0x5f, 0x3f, 0x38, 0x3c, 0x78, 0x15,
	0xd8, 0x5d, 0xcf, 0xf5, 0x43, 0x6b, 0x64, 0x0c, 0xdf, 0x81, 0x79, 0x59, 0xbc, 0xcc, 0x25, 0x63,
	0x2b, 0xe2, 0x68, 0x21, 0x75, 0x1d, 0x05, 0x64, 0x25, 0x77, 0x9b, 0x61, 0x4e, 0xf2, 0x0a, 0xf0,
	0x03, 0xf8, 0x3f, 0x1b, 0x4b, 0x26, 0x2a, 0x07, 0x6b, 0x44, 0x1c, 0xd5, 0x86, 0x13, 0x1b, 0xc9,
	0xae, 0xa6, 0xcb, 0x9b, 0x60, 0xb0, 0x35, 0xa9, 0x03, 0x7d, 0x30, 0x97, 0x41, 0xc7, 0xd7, 0x3d,
	0xb7, 0x5e, 0xde, 0x50, 0x36, 0x17, 0x1a, 0x48, 0x2f, 0xfc, 0x65, 0xf4, 0xd6, 0x58, 0xdb, 0xfe,
	0x93, 0x88, 0xa3, 0xf9, 0xa1, 0x7c, 0x32, 0x8f, 0x98, 0xa3, 0xe5, 0x09, 0xd7, 0x04, 0xc6, 0x56,
	0x4e, 0x1f, 0x6e, 0x81, 0xb2, 0x6f, 0x7b, 0x6e, 0x7d, 0x4a, 0xdc, 0x7f, 0x35, 0xe2, 0x48, 0xd4,
	0x31, 0x47, 0x55, 0x79, 0x3a, 0xa9, 0xb0, 0x25, 0x40, 0xf8, 0x1a, 0x54, 0x89, 0xdb, 0x77, 0x02,
	0xda, 0x4b, 0x76, 0x51, 0x9f, 0x16, 0x67, 0x1e, 0x47, 0x1c, 0x8d, 0xc3, 0x31, 0x47, 0x50, 0x1e,
	0x1d, 0x03, 0xb1, 0x35, 0xde, 0x22, 0x03, 0xf3, 0x74, 0x1b, 0xcc, 0x8d, 0xff, 0x08, 0x38, 0x0b,
	0xca, 0xad, 0x93, 0xc6, 0x76, 0xad, 0x04, 0x2b, 0x60, 0xaa, 0x75, 0xb2, 0xdb, 0x30, 0x6b, 0x0a,
	0xac, 0x82, 0x99, 0xc3, 0xf7, 0x47, 0x49, 0xe4, 0x6a, 0xea, 0xfe, 0xd1, 0xcf, 0x5b, 0x4d, 0xb9,
	0xb9, 0xd5, 0x94, 0xdf, 0xb7, 0x9a, 0xf2, 0xf5, 0x4e, 0x2b, 0xdd, 0xdc, 0x69, 0xa5, 0x5f, 0x77,
	0x5a, 0xe9, 0x63, 0xf3, 0xdf, 0x13, 0x23, 0x1f, 0x26, 0x91, 0x99, 0xce, 0xb4, 0x78, 0x68, 0x76,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x0c, 0x50, 0x17, 0xb2, 0x04, 0x00, 0x00,
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Xco1155ContractCode) > 0 {
		i -= len(m.Xco1155ContractCode)
		copy(dAtA[i:], m.Xco1155ContractCode)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Xco1155ContractCode)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cw721ContractCode) > 0 {
		i -= len(m.Cw721ContractCode)
		copy(dAtA[i:], m.Cw721ContractCode)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Cw721ContractCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cw20ContractCode) > 0 {
		i -= len(m.Cw20ContractCode)
		copy(dAtA[i:], m.Cw20ContractCode)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Cw20ContractCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenMinter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenMinter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenMinter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ContractType != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.ContractType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintToken(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MinterAddress) > 0 {
		i -= len(m.MinterAddress)
		copy(dAtA[i:], m.MinterAddress)
		i = encodeVarintToken(dAtA, i, uint64(len(m.MinterAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MinterDid) > 0 {
		i -= len(m.MinterDid)
		copy(dAtA[i:], m.MinterDid)
		i = encodeVarintToken(dAtA, i, uint64(len(m.MinterDid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cw20ContractCode)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Cw721ContractCode)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Xco1155ContractCode)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func (m *TokenMinter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MinterDid)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.MinterAddress)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.ContractType != 0 {
		n += 1 + sovToken(uint64(m.ContractType))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cw20ContractCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cw20ContractCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cw721ContractCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cw721ContractCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Xco1155ContractCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Xco1155ContractCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenMinter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: TokenMinter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenMinter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinterDid = github_com_ixofoundation_ixo_blockchain_x_iid_types.DIDFragment(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractType", wireType)
			}
			m.ContractType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractType |= ContractType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
