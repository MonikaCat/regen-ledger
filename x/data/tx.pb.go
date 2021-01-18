// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/data/v1alpha2/tx.proto

package data

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// MsgAnchorDataRequest is the Msg/AnchorData request type.
type MsgAnchorDataRequest struct {
	// sender is the address of the sender of the transaction.
	// The sender in StoreData is not attesting to the veracity of the underlying
	// data. They can simply be a intermediary providing services.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// id is the hash-based identifier for the anchored content.
	Id *ID `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgAnchorDataRequest) Reset()         { *m = MsgAnchorDataRequest{} }
func (m *MsgAnchorDataRequest) String() string { return proto.CompactTextString(m) }
func (*MsgAnchorDataRequest) ProtoMessage()    {}
func (*MsgAnchorDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff31907a513a4b24, []int{0}
}
func (m *MsgAnchorDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAnchorDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnchorDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAnchorDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnchorDataRequest.Merge(m, src)
}
func (m *MsgAnchorDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgAnchorDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnchorDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnchorDataRequest proto.InternalMessageInfo

func (m *MsgAnchorDataRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgAnchorDataRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

// MsgAnchorDataRequest is the Msg/AnchorData response type.
type MsgAnchorDataResponse struct {
	// timestamp is the timestamp of the block at which the data was anchored.
	Timestamp *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *MsgAnchorDataResponse) Reset()         { *m = MsgAnchorDataResponse{} }
func (m *MsgAnchorDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAnchorDataResponse) ProtoMessage()    {}
func (*MsgAnchorDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff31907a513a4b24, []int{1}
}
func (m *MsgAnchorDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAnchorDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnchorDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAnchorDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnchorDataResponse.Merge(m, src)
}
func (m *MsgAnchorDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAnchorDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnchorDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnchorDataResponse proto.InternalMessageInfo

func (m *MsgAnchorDataResponse) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// MsgSignDataRequest is the Msg/SignData request type.
type MsgSignDataRequest struct {
	// signers are the addresses of the accounts signing the data.
	// By making a SignData request, the signers are attesting to the veracity
	// of the data referenced by the cid. The precise meaning of this may vary
	// depending on the underlying data.
	Signers []string `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty"`
	// id is the hash-based identifier for the anchored content.
	Id *ID `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgSignDataRequest) Reset()         { *m = MsgSignDataRequest{} }
func (m *MsgSignDataRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSignDataRequest) ProtoMessage()    {}
func (*MsgSignDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff31907a513a4b24, []int{2}
}
func (m *MsgSignDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignDataRequest.Merge(m, src)
}
func (m *MsgSignDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignDataRequest proto.InternalMessageInfo

// MsgSignDataResponse is the Msg/SignData response type.
type MsgSignDataResponse struct {
}

func (m *MsgSignDataResponse) Reset()         { *m = MsgSignDataResponse{} }
func (m *MsgSignDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSignDataResponse) ProtoMessage()    {}
func (*MsgSignDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff31907a513a4b24, []int{3}
}
func (m *MsgSignDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignDataResponse.Merge(m, src)
}
func (m *MsgSignDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignDataResponse proto.InternalMessageInfo

// MsgStoreRawDataRequest is the Msg/StoreRawData request type.
type MsgStoreRawDataRequest struct {
	// sender is the address of the sender of the transaction.
	// The sender in StoreData is not attesting to the veracity of the underlying
	// data. They can simply be a intermediary providing services.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// id is the hash-based identifier for the anchored content.
	// The id's type must equal ID_TYPE_RAW_UNSPECIFIED.
	Id *ID `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// content is the content of the raw data corresponding to the provided ID.
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *MsgStoreRawDataRequest) Reset()         { *m = MsgStoreRawDataRequest{} }
func (m *MsgStoreRawDataRequest) String() string { return proto.CompactTextString(m) }
func (*MsgStoreRawDataRequest) ProtoMessage()    {}
func (*MsgStoreRawDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff31907a513a4b24, []int{4}
}
func (m *MsgStoreRawDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreRawDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreRawDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreRawDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreRawDataRequest.Merge(m, src)
}
func (m *MsgStoreRawDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreRawDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreRawDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreRawDataRequest proto.InternalMessageInfo

func (m *MsgStoreRawDataRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgStoreRawDataRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MsgStoreRawDataRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// MsgStoreRawDataRequest is the Msg/StoreRawData response type.
type MsgStoreRawDataResponse struct {
}

func (m *MsgStoreRawDataResponse) Reset()         { *m = MsgStoreRawDataResponse{} }
func (m *MsgStoreRawDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStoreRawDataResponse) ProtoMessage()    {}
func (*MsgStoreRawDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff31907a513a4b24, []int{5}
}
func (m *MsgStoreRawDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreRawDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreRawDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreRawDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreRawDataResponse.Merge(m, src)
}
func (m *MsgStoreRawDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreRawDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreRawDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreRawDataResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAnchorDataRequest)(nil), "regen.data.v1alpha2.MsgAnchorDataRequest")
	proto.RegisterType((*MsgAnchorDataResponse)(nil), "regen.data.v1alpha2.MsgAnchorDataResponse")
	proto.RegisterType((*MsgSignDataRequest)(nil), "regen.data.v1alpha2.MsgSignDataRequest")
	proto.RegisterType((*MsgSignDataResponse)(nil), "regen.data.v1alpha2.MsgSignDataResponse")
	proto.RegisterType((*MsgStoreRawDataRequest)(nil), "regen.data.v1alpha2.MsgStoreRawDataRequest")
	proto.RegisterType((*MsgStoreRawDataResponse)(nil), "regen.data.v1alpha2.MsgStoreRawDataResponse")
}

func init() { proto.RegisterFile("regen/data/v1alpha2/tx.proto", fileDescriptor_ff31907a513a4b24) }

var fileDescriptor_ff31907a513a4b24 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0x26, 0xa8, 0x90, 0x69, 0x4f, 0xdb, 0x3f, 0x63, 0x21, 0x37, 0xf2, 0xa5, 0x06, 0xca,
	0xae, 0x08, 0x17, 0xc4, 0x0d, 0x54, 0x21, 0x71, 0xc8, 0x01, 0x83, 0x84, 0x84, 0xd4, 0xc3, 0xc6,
	0x1e, 0x36, 0x16, 0xc9, 0xae, 0xd9, 0xdd, 0xd0, 0xf2, 0x06, 0x1c, 0x79, 0x04, 0x1e, 0x82, 0x87,
	0xe0, 0xd8, 0x23, 0x47, 0x94, 0xbc, 0x08, 0xf2, 0x1f, 0x2d, 0xad, 0x23, 0x72, 0xe8, 0xcd, 0x9f,
	0xe7, 0x9b, 0x6f, 0xbe, 0xf9, 0xc6, 0x86, 0x7b, 0x06, 0x25, 0x2a, 0x9e, 0x0a, 0x27, 0xf8, 0xe7,
	0xc7, 0x62, 0x9a, 0x4f, 0xc4, 0x90, 0xbb, 0x33, 0x96, 0x1b, 0xed, 0x34, 0xdd, 0x2e, 0xab, 0xac,
	0xa8, 0xb2, 0xa6, 0xea, 0xef, 0x48, 0x2d, 0x75, 0x59, 0xe7, 0xc5, 0x53, 0x45, 0xf5, 0x0f, 0xa4,
	0xd6, 0x72, 0x8a, 0xbc, 0x44, 0xe3, 0xf9, 0x07, 0xee, 0xb2, 0x19, 0x5a, 0x27, 0x66, 0x79, 0x43,
	0x68, 0x9d, 0xf4, 0x25, 0x47, 0x5b, 0x11, 0xc2, 0x77, 0xb0, 0x33, 0xb2, 0xf2, 0xb9, 0x4a, 0x26,
	0xda, 0x1c, 0x0b, 0x27, 0x62, 0xfc, 0x34, 0x47, 0xeb, 0xe8, 0x1e, 0x6c, 0x58, 0x54, 0x29, 0x1a,
	0x8f, 0x0c, 0x48, 0xd4, 0x8f, 0x6b, 0x44, 0x0f, 0xa1, 0x9b, 0xa5, 0x5e, 0x77, 0x40, 0xa2, 0xcd,
	0xe1, 0x3e, 0x6b, 0x71, 0xca, 0x5e, 0x1d, 0xc7, 0xdd, 0x2c, 0x0d, 0x5f, 0xc3, 0xee, 0x15, 0x61,
	0x9b, 0x6b, 0x65, 0x91, 0x3e, 0x85, 0xfe, 0x5f, 0x97, 0xa5, 0xf8, 0xe6, 0xd0, 0x67, 0xd5, 0x1e,
	0xac, 0xd9, 0x83, 0xbd, 0x6d, 0x18, 0xf1, 0x05, 0x39, 0x3c, 0x01, 0x3a, 0xb2, 0xf2, 0x4d, 0x26,
	0xd5, 0x65, 0xa7, 0x1e, 0xdc, 0xb6, 0x99, 0x54, 0x68, 0xac, 0x47, 0x06, 0xbd, 0xa8, 0x1f, 0x37,
	0x70, 0x6d, 0xaf, 0xcf, 0x6e, 0x7d, 0xfd, 0x7e, 0xd0, 0x09, 0x77, 0x61, 0xfb, 0x1f, 0xf9, 0xca,
	0x6f, 0x68, 0x61, 0xaf, 0x78, 0xed, 0xb4, 0xc1, 0x58, 0x9c, 0xde, 0x64, 0x46, 0x85, 0xf5, 0x44,
	0x2b, 0x87, 0xca, 0x79, 0xbd, 0x01, 0x89, 0xb6, 0xe2, 0x06, 0x86, 0x77, 0x61, 0xff, 0xda, 0xd0,
	0xca, 0xcf, 0xf0, 0x47, 0x17, 0x7a, 0x23, 0x2b, 0x69, 0x02, 0x70, 0x91, 0x2e, 0xbd, 0xdf, 0x3a,
	0xa7, 0xed, 0xb4, 0xfe, 0x83, 0x75, 0xa8, 0xf5, 0xb1, 0x4e, 0xe0, 0x4e, 0x13, 0x08, 0x3d, 0x5c,
	0xd5, 0x77, 0xe5, 0x22, 0x7e, 0xf4, 0x7f, 0x62, 0x2d, 0x9f, 0xc1, 0xd6, 0xe5, 0x1d, 0xe9, 0xc3,
	0x95, 0x9d, 0xd7, 0xe3, 0xf7, 0x8f, 0xd6, 0x23, 0x57, 0xa3, 0x5e, 0xbc, 0xfc, 0xb9, 0x08, 0xc8,
	0xf9, 0x22, 0x20, 0xbf, 0x17, 0x01, 0xf9, 0xb6, 0x0c, 0x3a, 0xe7, 0xcb, 0xa0, 0xf3, 0x6b, 0x19,
	0x74, 0xde, 0x1f, 0xc9, 0xcc, 0x4d, 0xe6, 0x63, 0x96, 0xe8, 0x19, 0x2f, 0x15, 0x1f, 0x29, 0x74,
	0xa7, 0xda, 0x7c, 0xac, 0xd1, 0x14, 0x53, 0x89, 0x86, 0x9f, 0x95, 0x7f, 0xd1, 0x78, 0xa3, 0xfc,
	0x46, 0x9f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x90, 0x8d, 0xc5, 0x19, 0xc4, 0x03, 0x00, 0x00,
}

func (m *MsgAnchorDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnchorDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnchorDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		{
			size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAnchorDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnchorDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnchorDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSignDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		{
			size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgSignDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgStoreRawDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreRawDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreRawDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != nil {
		{
			size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStoreRawDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreRawDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreRawDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAnchorDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAnchorDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSignDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSignDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgStoreRawDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgStoreRawDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAnchorDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAnchorDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnchorDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &ID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAnchorDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAnchorDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnchorDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSignDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSignDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &ID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSignDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSignDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStoreRawDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStoreRawDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreRawDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &ID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStoreRawDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStoreRawDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreRawDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
