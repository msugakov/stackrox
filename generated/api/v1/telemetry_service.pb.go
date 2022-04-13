// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/telemetry_service.proto

package v1

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

type ConfigureTelemetryRequest struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigureTelemetryRequest) Reset()         { *m = ConfigureTelemetryRequest{} }
func (m *ConfigureTelemetryRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigureTelemetryRequest) ProtoMessage()    {}
func (*ConfigureTelemetryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d29ceed52498e29, []int{0}
}
func (m *ConfigureTelemetryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigureTelemetryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigureTelemetryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigureTelemetryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureTelemetryRequest.Merge(m, src)
}
func (m *ConfigureTelemetryRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConfigureTelemetryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureTelemetryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureTelemetryRequest proto.InternalMessageInfo

func (m *ConfigureTelemetryRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ConfigureTelemetryRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ConfigureTelemetryRequest) Clone() *ConfigureTelemetryRequest {
	if m == nil {
		return nil
	}
	cloned := new(ConfigureTelemetryRequest)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*ConfigureTelemetryRequest)(nil), "v1.ConfigureTelemetryRequest")
}

func init() { proto.RegisterFile("api/v1/telemetry_service.proto", fileDescriptor_3d29ceed52498e29) }

var fileDescriptor_3d29ceed52498e29 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0xd4, 0x2f, 0x49, 0xcd, 0x49, 0xcd, 0x4d, 0x2d, 0x29, 0xaa, 0x8c, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x29, 0x4d, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x90, 0x12, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0x45,
	0x18, 0x01, 0x95, 0x10, 0x82, 0x1a, 0x9d, 0x9a, 0x5b, 0x50, 0x02, 0x15, 0x53, 0x32, 0xe5, 0x92,
	0x74, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x2f, 0x2d, 0x4a, 0x0d, 0x81, 0xa9, 0x0f, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0xcd, 0x4b, 0x4c, 0xca, 0x49, 0x4d, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x71, 0x8d, 0xbe, 0x31, 0x72, 0x09, 0xc0, 0x95, 0x07, 0x43,
	0x1c, 0x28, 0x94, 0xc1, 0x25, 0xe9, 0x9e, 0x5a, 0x02, 0x17, 0x86, 0x99, 0x0b, 0x76, 0x9c, 0x10,
	0xa7, 0x5e, 0x99, 0xa1, 0x9e, 0x2b, 0xc8, 0x66, 0x29, 0x79, 0x3d, 0xa8, 0x0b, 0xf5, 0xb0, 0xab,
	0x55, 0x92, 0x6f, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xa4, 0x90, 0x38, 0x4a, 0x40, 0xe8, 0x27, 0xc3,
	0x1c, 0x2a, 0x54, 0xc1, 0x25, 0x84, 0xe9, 0x6a, 0x21, 0x59, 0x90, 0x15, 0x38, 0x7d, 0x43, 0xd8,
	0x5a, 0x25, 0xb0, 0xb5, 0x32, 0x52, 0xb8, 0xac, 0xb5, 0x62, 0xd4, 0x72, 0xd2, 0x3b, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0xe0, 0x92,
	0xc8, 0xcc, 0xd7, 0x2b, 0x2e, 0x49, 0x4c, 0xce, 0x2e, 0xca, 0xaf, 0x80, 0x04, 0xaa, 0x5e, 0x62,
	0x41, 0xa6, 0x5e, 0x99, 0x61, 0x14, 0x53, 0x99, 0x61, 0x04, 0x43, 0x12, 0x1b, 0x58, 0xcc, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x42, 0x4d, 0x40, 0xa5, 0xd9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TelemetryServiceClient is the client API for TelemetryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type TelemetryServiceClient interface {
	GetTelemetryConfiguration(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*storage.TelemetryConfiguration, error)
	ConfigureTelemetry(ctx context.Context, in *ConfigureTelemetryRequest, opts ...grpc.CallOption) (*storage.TelemetryConfiguration, error)
}

type telemetryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTelemetryServiceClient(cc grpc.ClientConnInterface) TelemetryServiceClient {
	return &telemetryServiceClient{cc}
}

func (c *telemetryServiceClient) GetTelemetryConfiguration(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*storage.TelemetryConfiguration, error) {
	out := new(storage.TelemetryConfiguration)
	err := c.cc.Invoke(ctx, "/v1.TelemetryService/GetTelemetryConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServiceClient) ConfigureTelemetry(ctx context.Context, in *ConfigureTelemetryRequest, opts ...grpc.CallOption) (*storage.TelemetryConfiguration, error) {
	out := new(storage.TelemetryConfiguration)
	err := c.cc.Invoke(ctx, "/v1.TelemetryService/ConfigureTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServiceServer is the server API for TelemetryService service.
type TelemetryServiceServer interface {
	GetTelemetryConfiguration(context.Context, *Empty) (*storage.TelemetryConfiguration, error)
	ConfigureTelemetry(context.Context, *ConfigureTelemetryRequest) (*storage.TelemetryConfiguration, error)
}

// UnimplementedTelemetryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTelemetryServiceServer struct {
}

func (*UnimplementedTelemetryServiceServer) GetTelemetryConfiguration(ctx context.Context, req *Empty) (*storage.TelemetryConfiguration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTelemetryConfiguration not implemented")
}
func (*UnimplementedTelemetryServiceServer) ConfigureTelemetry(ctx context.Context, req *ConfigureTelemetryRequest) (*storage.TelemetryConfiguration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureTelemetry not implemented")
}

func RegisterTelemetryServiceServer(s *grpc.Server, srv TelemetryServiceServer) {
	s.RegisterService(&_TelemetryService_serviceDesc, srv)
}

func _TelemetryService_GetTelemetryConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServiceServer).GetTelemetryConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TelemetryService/GetTelemetryConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServiceServer).GetTelemetryConfiguration(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryService_ConfigureTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServiceServer).ConfigureTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TelemetryService/ConfigureTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServiceServer).ConfigureTelemetry(ctx, req.(*ConfigureTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TelemetryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.TelemetryService",
	HandlerType: (*TelemetryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTelemetryConfiguration",
			Handler:    _TelemetryService_GetTelemetryConfiguration_Handler,
		},
		{
			MethodName: "ConfigureTelemetry",
			Handler:    _TelemetryService_ConfigureTelemetry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/telemetry_service.proto",
}

func (m *ConfigureTelemetryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigureTelemetryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigureTelemetryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTelemetryService(dAtA []byte, offset int, v uint64) int {
	offset -= sovTelemetryService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfigureTelemetryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTelemetryService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTelemetryService(x uint64) (n int) {
	return sovTelemetryService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfigureTelemetryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelemetryService
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
			return fmt.Errorf("proto: ConfigureTelemetryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigureTelemetryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetryService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTelemetryService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTelemetryService
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
func skipTelemetryService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTelemetryService
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
					return 0, ErrIntOverflowTelemetryService
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
					return 0, ErrIntOverflowTelemetryService
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
				return 0, ErrInvalidLengthTelemetryService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTelemetryService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTelemetryService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTelemetryService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTelemetryService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTelemetryService = fmt.Errorf("proto: unexpected end of group")
)
