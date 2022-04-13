// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/central/local_scanner.proto

package central

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	storage "github.com/stackrox/stackrox/generated/storage"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LocalScannerCertsIssueError struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalScannerCertsIssueError) Reset()         { *m = LocalScannerCertsIssueError{} }
func (m *LocalScannerCertsIssueError) String() string { return proto.CompactTextString(m) }
func (*LocalScannerCertsIssueError) ProtoMessage()    {}
func (*LocalScannerCertsIssueError) Descriptor() ([]byte, []int) {
	return fileDescriptor_856923c76f63cf0a, []int{0}
}
func (m *LocalScannerCertsIssueError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalScannerCertsIssueError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalScannerCertsIssueError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalScannerCertsIssueError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalScannerCertsIssueError.Merge(m, src)
}
func (m *LocalScannerCertsIssueError) XXX_Size() int {
	return m.Size()
}
func (m *LocalScannerCertsIssueError) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalScannerCertsIssueError.DiscardUnknown(m)
}

var xxx_messageInfo_LocalScannerCertsIssueError proto.InternalMessageInfo

func (m *LocalScannerCertsIssueError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LocalScannerCertsIssueError) MessageClone() proto.Message {
	return m.Clone()
}
func (m *LocalScannerCertsIssueError) Clone() *LocalScannerCertsIssueError {
	if m == nil {
		return nil
	}
	cloned := new(LocalScannerCertsIssueError)
	*cloned = *m

	return cloned
}

type IssueLocalScannerCertsRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueLocalScannerCertsRequest) Reset()         { *m = IssueLocalScannerCertsRequest{} }
func (m *IssueLocalScannerCertsRequest) String() string { return proto.CompactTextString(m) }
func (*IssueLocalScannerCertsRequest) ProtoMessage()    {}
func (*IssueLocalScannerCertsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_856923c76f63cf0a, []int{1}
}
func (m *IssueLocalScannerCertsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IssueLocalScannerCertsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IssueLocalScannerCertsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IssueLocalScannerCertsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueLocalScannerCertsRequest.Merge(m, src)
}
func (m *IssueLocalScannerCertsRequest) XXX_Size() int {
	return m.Size()
}
func (m *IssueLocalScannerCertsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueLocalScannerCertsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueLocalScannerCertsRequest proto.InternalMessageInfo

func (m *IssueLocalScannerCertsRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *IssueLocalScannerCertsRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *IssueLocalScannerCertsRequest) Clone() *IssueLocalScannerCertsRequest {
	if m == nil {
		return nil
	}
	cloned := new(IssueLocalScannerCertsRequest)
	*cloned = *m

	return cloned
}

type IssueLocalScannerCertsResponse struct {
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*IssueLocalScannerCertsResponse_Certificates
	//	*IssueLocalScannerCertsResponse_Error
	Response             isIssueLocalScannerCertsResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *IssueLocalScannerCertsResponse) Reset()         { *m = IssueLocalScannerCertsResponse{} }
func (m *IssueLocalScannerCertsResponse) String() string { return proto.CompactTextString(m) }
func (*IssueLocalScannerCertsResponse) ProtoMessage()    {}
func (*IssueLocalScannerCertsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_856923c76f63cf0a, []int{2}
}
func (m *IssueLocalScannerCertsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IssueLocalScannerCertsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IssueLocalScannerCertsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IssueLocalScannerCertsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueLocalScannerCertsResponse.Merge(m, src)
}
func (m *IssueLocalScannerCertsResponse) XXX_Size() int {
	return m.Size()
}
func (m *IssueLocalScannerCertsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueLocalScannerCertsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueLocalScannerCertsResponse proto.InternalMessageInfo

type isIssueLocalScannerCertsResponse_Response interface {
	isIssueLocalScannerCertsResponse_Response()
	MarshalTo([]byte) (int, error)
	Size() int
	Clone() isIssueLocalScannerCertsResponse_Response
}

type IssueLocalScannerCertsResponse_Certificates struct {
	Certificates *storage.TypedServiceCertificateSet `protobuf:"bytes,2,opt,name=certificates,proto3,oneof" json:"certificates,omitempty"`
}
type IssueLocalScannerCertsResponse_Error struct {
	Error *LocalScannerCertsIssueError `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (*IssueLocalScannerCertsResponse_Certificates) isIssueLocalScannerCertsResponse_Response() {}
func (m *IssueLocalScannerCertsResponse_Certificates) Clone() isIssueLocalScannerCertsResponse_Response {
	if m == nil {
		return nil
	}
	cloned := new(IssueLocalScannerCertsResponse_Certificates)
	*cloned = *m

	cloned.Certificates = m.Certificates.Clone()
	return cloned
}
func (*IssueLocalScannerCertsResponse_Error) isIssueLocalScannerCertsResponse_Response() {}
func (m *IssueLocalScannerCertsResponse_Error) Clone() isIssueLocalScannerCertsResponse_Response {
	if m == nil {
		return nil
	}
	cloned := new(IssueLocalScannerCertsResponse_Error)
	*cloned = *m

	cloned.Error = m.Error.Clone()
	return cloned
}

func (m *IssueLocalScannerCertsResponse) GetResponse() isIssueLocalScannerCertsResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *IssueLocalScannerCertsResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *IssueLocalScannerCertsResponse) GetCertificates() *storage.TypedServiceCertificateSet {
	if x, ok := m.GetResponse().(*IssueLocalScannerCertsResponse_Certificates); ok {
		return x.Certificates
	}
	return nil
}

func (m *IssueLocalScannerCertsResponse) GetError() *LocalScannerCertsIssueError {
	if x, ok := m.GetResponse().(*IssueLocalScannerCertsResponse_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IssueLocalScannerCertsResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IssueLocalScannerCertsResponse_Certificates)(nil),
		(*IssueLocalScannerCertsResponse_Error)(nil),
	}
}

func (m *IssueLocalScannerCertsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *IssueLocalScannerCertsResponse) Clone() *IssueLocalScannerCertsResponse {
	if m == nil {
		return nil
	}
	cloned := new(IssueLocalScannerCertsResponse)
	*cloned = *m

	if m.Response != nil {
		cloned.Response = m.Response.Clone()
	}
	return cloned
}

func init() {
	proto.RegisterType((*LocalScannerCertsIssueError)(nil), "central.LocalScannerCertsIssueError")
	proto.RegisterType((*IssueLocalScannerCertsRequest)(nil), "central.IssueLocalScannerCertsRequest")
	proto.RegisterType((*IssueLocalScannerCertsResponse)(nil), "central.IssueLocalScannerCertsResponse")
}

func init() {
	proto.RegisterFile("internalapi/central/local_scanner.proto", fileDescriptor_856923c76f63cf0a)
}

var fileDescriptor_856923c76f63cf0a = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x3b, 0x8a, 0xd6, 0x8e, 0xae, 0xb2, 0x8a, 0x4a, 0x87, 0x12, 0x05, 0xbb, 0x4a, 0x40,
	0x17, 0x6e, 0xc4, 0x45, 0x8b, 0xd0, 0x80, 0xab, 0xc4, 0x95, 0x9b, 0x30, 0x4e, 0xae, 0x65, 0x20,
	0xce, 0xc4, 0x7b, 0xa7, 0x42, 0xdf, 0xc4, 0x47, 0x72, 0xe9, 0xc2, 0x07, 0x90, 0xf8, 0x22, 0xd2,
	0x64, 0x8a, 0x88, 0xa8, 0xbb, 0xf9, 0x39, 0xdf, 0xe1, 0x9e, 0x73, 0xf9, 0x89, 0x36, 0x0e, 0xd0,
	0xc8, 0x4a, 0xd6, 0x3a, 0x51, 0x60, 0x1c, 0xca, 0x2a, 0xa9, 0xac, 0x92, 0x55, 0x41, 0x4a, 0x1a,
	0x03, 0x18, 0xd7, 0x68, 0x9d, 0x0d, 0xfa, 0xfe, 0xf3, 0x40, 0x90, 0xb3, 0x28, 0xe7, 0x90, 0x10,
	0xe0, 0x93, 0x56, 0x50, 0xe8, 0x12, 0x8c, 0xd3, 0x6e, 0xd9, 0x09, 0xa3, 0x73, 0x7e, 0x78, 0xbd,
	0xe2, 0xf3, 0x0e, 0x9f, 0x02, 0x3a, 0x4a, 0x89, 0x16, 0x70, 0x85, 0x68, 0x31, 0x08, 0x79, 0xff,
	0x01, 0x88, 0xe4, 0x1c, 0x42, 0x36, 0x62, 0xe3, 0x41, 0xb6, 0xbe, 0x46, 0x97, 0x7c, 0xd8, 0xea,
	0x7e, 0xd0, 0x19, 0x3c, 0x2e, 0x80, 0x5c, 0x30, 0xe4, 0x1c, 0xbb, 0x63, 0xa1, 0x4b, 0x4f, 0x0f,
	0xfc, 0x4b, 0x5a, 0x46, 0x6f, 0x8c, 0x8b, 0xdf, 0x0c, 0xa8, 0xb6, 0x86, 0xe0, 0x1f, 0x87, 0x20,
	0xe5, 0x7b, 0x0a, 0xd0, 0xe9, 0x7b, 0xad, 0xa4, 0x03, 0x0a, 0x37, 0x46, 0x6c, 0xbc, 0x7b, 0x7a,
	0x14, 0xfb, 0xc4, 0xf1, 0xcd, 0xb2, 0x86, 0x32, 0xef, 0x62, 0x4f, 0xbf, 0x84, 0x39, 0xb8, 0x59,
	0x2f, 0xfb, 0x86, 0x06, 0x17, 0x7c, 0x0b, 0x56, 0x79, 0xc3, 0xcd, 0xd6, 0xe3, 0x38, 0xf6, 0xf5,
	0xc5, 0x7f, 0x74, 0x33, 0xeb, 0x65, 0x1d, 0x34, 0xe1, 0x7c, 0x07, 0xfd, 0xcc, 0x93, 0xfd, 0x97,
	0x46, 0xb0, 0xd7, 0x46, 0xb0, 0xf7, 0x46, 0xb0, 0xe7, 0x0f, 0xd1, 0xbb, 0x5d, 0xaf, 0xe2, 0x6e,
	0xbb, 0x6d, 0xfc, 0xec, 0x33, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x93, 0x4a, 0x05, 0xc5, 0x01, 0x00,
	0x00,
}

func (m *LocalScannerCertsIssueError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalScannerCertsIssueError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalScannerCertsIssueError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintLocalScanner(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IssueLocalScannerCertsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IssueLocalScannerCertsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IssueLocalScannerCertsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = encodeVarintLocalScanner(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IssueLocalScannerCertsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IssueLocalScannerCertsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IssueLocalScannerCertsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Response != nil {
		{
			size := m.Response.Size()
			i -= size
			if _, err := m.Response.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = encodeVarintLocalScanner(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IssueLocalScannerCertsResponse_Certificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IssueLocalScannerCertsResponse_Certificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Certificates != nil {
		{
			size, err := m.Certificates.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLocalScanner(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *IssueLocalScannerCertsResponse_Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IssueLocalScannerCertsResponse_Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Error != nil {
		{
			size, err := m.Error.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLocalScanner(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func encodeVarintLocalScanner(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocalScanner(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LocalScannerCertsIssueError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IssueLocalScannerCertsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IssueLocalScannerCertsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	if m.Response != nil {
		n += m.Response.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IssueLocalScannerCertsResponse_Certificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Certificates != nil {
		l = m.Certificates.Size()
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	return n
}
func (m *IssueLocalScannerCertsResponse_Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	return n
}

func sovLocalScanner(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocalScanner(x uint64) (n int) {
	return sovLocalScanner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocalScannerCertsIssueError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalScanner
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
			return fmt.Errorf("proto: LocalScannerCertsIssueError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalScannerCertsIssueError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalScanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalScanner
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
func (m *IssueLocalScannerCertsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalScanner
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
			return fmt.Errorf("proto: IssueLocalScannerCertsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueLocalScannerCertsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalScanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalScanner
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
func (m *IssueLocalScannerCertsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalScanner
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
			return fmt.Errorf("proto: IssueLocalScannerCertsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueLocalScannerCertsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &storage.TypedServiceCertificateSet{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &IssueLocalScannerCertsResponse_Certificates{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &LocalScannerCertsIssueError{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &IssueLocalScannerCertsResponse_Error{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalScanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalScanner
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
func skipLocalScanner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocalScanner
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
					return 0, ErrIntOverflowLocalScanner
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
					return 0, ErrIntOverflowLocalScanner
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
				return 0, ErrInvalidLengthLocalScanner
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocalScanner
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocalScanner
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocalScanner        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocalScanner          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocalScanner = fmt.Errorf("proto: unexpected end of group")
)
