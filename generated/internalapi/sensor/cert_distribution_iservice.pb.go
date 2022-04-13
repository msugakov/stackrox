// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/sensor/cert_distribution_iservice.proto

package sensor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	storage "github.com/stackrox/stackrox/generated/storage"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FetchCertificateRequest struct {
	ServiceType          storage.ServiceType `protobuf:"varint,1,opt,name=service_type,json=serviceType,proto3,enum=storage.ServiceType" json:"service_type,omitempty"`
	ServiceAccountToken  string              `protobuf:"bytes,2,opt,name=service_account_token,json=serviceAccountToken,proto3" json:"service_account_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FetchCertificateRequest) Reset()         { *m = FetchCertificateRequest{} }
func (m *FetchCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*FetchCertificateRequest) ProtoMessage()    {}
func (*FetchCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc1f5deeeb84e10, []int{0}
}
func (m *FetchCertificateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchCertificateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchCertificateRequest.Merge(m, src)
}
func (m *FetchCertificateRequest) XXX_Size() int {
	return m.Size()
}
func (m *FetchCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchCertificateRequest proto.InternalMessageInfo

func (m *FetchCertificateRequest) GetServiceType() storage.ServiceType {
	if m != nil {
		return m.ServiceType
	}
	return storage.ServiceType_UNKNOWN_SERVICE
}

func (m *FetchCertificateRequest) GetServiceAccountToken() string {
	if m != nil {
		return m.ServiceAccountToken
	}
	return ""
}

func (m *FetchCertificateRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *FetchCertificateRequest) Clone() *FetchCertificateRequest {
	if m == nil {
		return nil
	}
	cloned := new(FetchCertificateRequest)
	*cloned = *m

	return cloned
}

type FetchCertificateResponse struct {
	PemCert              string   `protobuf:"bytes,1,opt,name=pem_cert,json=pemCert,proto3" json:"pem_cert,omitempty"`
	PemKey               string   `protobuf:"bytes,2,opt,name=pem_key,json=pemKey,proto3" json:"pem_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchCertificateResponse) Reset()         { *m = FetchCertificateResponse{} }
func (m *FetchCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*FetchCertificateResponse) ProtoMessage()    {}
func (*FetchCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc1f5deeeb84e10, []int{1}
}
func (m *FetchCertificateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchCertificateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchCertificateResponse.Merge(m, src)
}
func (m *FetchCertificateResponse) XXX_Size() int {
	return m.Size()
}
func (m *FetchCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchCertificateResponse proto.InternalMessageInfo

func (m *FetchCertificateResponse) GetPemCert() string {
	if m != nil {
		return m.PemCert
	}
	return ""
}

func (m *FetchCertificateResponse) GetPemKey() string {
	if m != nil {
		return m.PemKey
	}
	return ""
}

func (m *FetchCertificateResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *FetchCertificateResponse) Clone() *FetchCertificateResponse {
	if m == nil {
		return nil
	}
	cloned := new(FetchCertificateResponse)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*FetchCertificateRequest)(nil), "sensor.FetchCertificateRequest")
	proto.RegisterType((*FetchCertificateResponse)(nil), "sensor.FetchCertificateResponse")
}

func init() {
	proto.RegisterFile("internalapi/sensor/cert_distribution_iservice.proto", fileDescriptor_bfc1f5deeeb84e10)
}

var fileDescriptor_bfc1f5deeeb84e10 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x3f, 0x7f, 0x43, 0xa1, 0x06, 0x21, 0x64, 0x40, 0x0d, 0x1d, 0x42, 0xd4, 0xa9, 0x53,
	0x22, 0xa5, 0x03, 0x33, 0x3f, 0x62, 0x41, 0x62, 0x08, 0x65, 0x61, 0xb1, 0x52, 0xf7, 0x00, 0x56,
	0x89, 0x6d, 0xec, 0x13, 0xa4, 0xdc, 0x00, 0xd7, 0xc0, 0x25, 0x31, 0x72, 0x09, 0x28, 0xdc, 0x08,
	0x4a, 0xed, 0x0a, 0x04, 0xea, 0x96, 0xe8, 0x79, 0xce, 0xf1, 0xeb, 0xd7, 0x74, 0x22, 0x15, 0x82,
	0x55, 0xe5, 0x63, 0x69, 0x64, 0xe6, 0x40, 0x39, 0x6d, 0x33, 0x01, 0x16, 0xf9, 0x5c, 0x3a, 0xb4,
	0x72, 0x56, 0xa3, 0xd4, 0x8a, 0x4b, 0x07, 0xf6, 0x59, 0x0a, 0x48, 0x8d, 0xd5, 0xa8, 0x59, 0xcf,
	0x8b, 0xc3, 0xd8, 0xa1, 0xb6, 0xe5, 0x3d, 0x64, 0x01, 0x73, 0x39, 0x07, 0x85, 0x12, 0x1b, 0xef,
	0x8d, 0x5e, 0x08, 0x1d, 0x5c, 0x00, 0x8a, 0x87, 0x33, 0xb0, 0x28, 0xef, 0xa4, 0x28, 0x11, 0x0a,
	0x78, 0xaa, 0xc1, 0x21, 0x3b, 0xa6, 0xdb, 0xab, 0x29, 0x6c, 0x0c, 0x44, 0x24, 0x21, 0xe3, 0x9d,
	0x7c, 0x3f, 0x0d, 0x2b, 0xd3, 0x6b, 0x0f, 0xa7, 0x8d, 0x81, 0x62, 0xcb, 0x7d, 0xff, 0xb0, 0x9c,
	0x1e, 0xac, 0x06, 0x4b, 0x21, 0x74, 0xad, 0x90, 0xa3, 0x5e, 0x80, 0x8a, 0xfe, 0x27, 0x64, 0xdc,
	0x2f, 0xf6, 0x02, 0x3c, 0xf1, 0x6c, 0xda, 0xa1, 0xd1, 0x15, 0x8d, 0xfe, 0xe6, 0x70, 0x46, 0x2b,
	0x07, 0xec, 0x90, 0x6e, 0x1a, 0xa8, 0x78, 0x77, 0xe9, 0x65, 0x88, 0x7e, 0xb1, 0x61, 0xa0, 0xea,
	0x4c, 0x36, 0xa0, 0xdd, 0x27, 0x5f, 0x40, 0x13, 0x96, 0xf7, 0x0c, 0x54, 0x97, 0xd0, 0xe4, 0x86,
	0x0e, 0x3a, 0xe1, 0xfc, 0x47, 0x47, 0x21, 0x2f, 0xbb, 0xa1, 0xbb, 0xbf, 0x8f, 0x62, 0x47, 0xa9,
	0x2f, 0x2c, 0x5d, 0x53, 0xc6, 0x30, 0x59, 0x2f, 0xf8, 0x94, 0xa7, 0xd1, 0x5b, 0x1b, 0x93, 0xf7,
	0x36, 0x26, 0x1f, 0x6d, 0x4c, 0x5e, 0x3f, 0xe3, 0x7f, 0xb7, 0xe1, 0x11, 0x66, 0xbd, 0x65, 0xd7,
	0x93, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xc7, 0x85, 0x08, 0xca, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CertDistributionServiceClient is the client API for CertDistributionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type CertDistributionServiceClient interface {
	FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error)
}

type certDistributionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertDistributionServiceClient(cc grpc.ClientConnInterface) CertDistributionServiceClient {
	return &certDistributionServiceClient{cc}
}

func (c *certDistributionServiceClient) FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error) {
	out := new(FetchCertificateResponse)
	err := c.cc.Invoke(ctx, "/sensor.CertDistributionService/FetchCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertDistributionServiceServer is the server API for CertDistributionService service.
type CertDistributionServiceServer interface {
	FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error)
}

// UnimplementedCertDistributionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCertDistributionServiceServer struct {
}

func (*UnimplementedCertDistributionServiceServer) FetchCertificate(ctx context.Context, req *FetchCertificateRequest) (*FetchCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCertificate not implemented")
}

func RegisterCertDistributionServiceServer(s *grpc.Server, srv CertDistributionServiceServer) {
	s.RegisterService(&_CertDistributionService_serviceDesc, srv)
}

func _CertDistributionService_FetchCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertDistributionServiceServer).FetchCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.CertDistributionService/FetchCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertDistributionServiceServer).FetchCertificate(ctx, req.(*FetchCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertDistributionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.CertDistributionService",
	HandlerType: (*CertDistributionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCertificate",
			Handler:    _CertDistributionService_FetchCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/sensor/cert_distribution_iservice.proto",
}

func (m *FetchCertificateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchCertificateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchCertificateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ServiceAccountToken) > 0 {
		i -= len(m.ServiceAccountToken)
		copy(dAtA[i:], m.ServiceAccountToken)
		i = encodeVarintCertDistributionIservice(dAtA, i, uint64(len(m.ServiceAccountToken)))
		i--
		dAtA[i] = 0x12
	}
	if m.ServiceType != 0 {
		i = encodeVarintCertDistributionIservice(dAtA, i, uint64(m.ServiceType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FetchCertificateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchCertificateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchCertificateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PemKey) > 0 {
		i -= len(m.PemKey)
		copy(dAtA[i:], m.PemKey)
		i = encodeVarintCertDistributionIservice(dAtA, i, uint64(len(m.PemKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PemCert) > 0 {
		i -= len(m.PemCert)
		copy(dAtA[i:], m.PemCert)
		i = encodeVarintCertDistributionIservice(dAtA, i, uint64(len(m.PemCert)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCertDistributionIservice(dAtA []byte, offset int, v uint64) int {
	offset -= sovCertDistributionIservice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FetchCertificateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServiceType != 0 {
		n += 1 + sovCertDistributionIservice(uint64(m.ServiceType))
	}
	l = len(m.ServiceAccountToken)
	if l > 0 {
		n += 1 + l + sovCertDistributionIservice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FetchCertificateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PemCert)
	if l > 0 {
		n += 1 + l + sovCertDistributionIservice(uint64(l))
	}
	l = len(m.PemKey)
	if l > 0 {
		n += 1 + l + sovCertDistributionIservice(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCertDistributionIservice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCertDistributionIservice(x uint64) (n int) {
	return sovCertDistributionIservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FetchCertificateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertDistributionIservice
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
			return fmt.Errorf("proto: FetchCertificateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchCertificateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceType", wireType)
			}
			m.ServiceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertDistributionIservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceType |= storage.ServiceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAccountToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertDistributionIservice
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
				return ErrInvalidLengthCertDistributionIservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertDistributionIservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceAccountToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCertDistributionIservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertDistributionIservice
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
func (m *FetchCertificateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertDistributionIservice
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
			return fmt.Errorf("proto: FetchCertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchCertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PemCert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertDistributionIservice
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
				return ErrInvalidLengthCertDistributionIservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertDistributionIservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PemCert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PemKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertDistributionIservice
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
				return ErrInvalidLengthCertDistributionIservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertDistributionIservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PemKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCertDistributionIservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertDistributionIservice
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
func skipCertDistributionIservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCertDistributionIservice
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
					return 0, ErrIntOverflowCertDistributionIservice
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
					return 0, ErrIntOverflowCertDistributionIservice
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
				return 0, ErrInvalidLengthCertDistributionIservice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCertDistributionIservice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCertDistributionIservice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCertDistributionIservice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCertDistributionIservice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCertDistributionIservice = fmt.Errorf("proto: unexpected end of group")
)
