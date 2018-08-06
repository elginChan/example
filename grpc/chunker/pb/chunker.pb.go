// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chunker.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	chunker.proto

It has these top-level messages:
	Empty
	Chunk
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorChunker, []int{0} }

type Chunk struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptorChunker, []int{1} }

func (m *Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*Chunk)(nil), "pb.Chunk")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chunker service

type ChunkerClient interface {
	Chunker(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chunker_ChunkerClient, error)
}

type chunkerClient struct {
	cc *grpc.ClientConn
}

func NewChunkerClient(cc *grpc.ClientConn) ChunkerClient {
	return &chunkerClient{cc}
}

func (c *chunkerClient) Chunker(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chunker_ChunkerClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Chunker_serviceDesc.Streams[0], c.cc, "/pb.Chunker/Chunker", opts...)
	if err != nil {
		return nil, err
	}
	x := &chunkerChunkerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chunker_ChunkerClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type chunkerChunkerClient struct {
	grpc.ClientStream
}

func (x *chunkerChunkerClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Chunker service

type ChunkerServer interface {
	Chunker(*Empty, Chunker_ChunkerServer) error
}

func RegisterChunkerServer(s *grpc.Server, srv ChunkerServer) {
	s.RegisterService(&_Chunker_serviceDesc, srv)
}

func _Chunker_Chunker_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChunkerServer).Chunker(m, &chunkerChunkerServer{stream})
}

type Chunker_ChunkerServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type chunkerChunkerServer struct {
	grpc.ServerStream
}

func (x *chunkerChunkerServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Chunker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Chunker",
	HandlerType: (*ChunkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chunker",
			Handler:       _Chunker_Chunker_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chunker.proto",
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Chunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chunk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Chunk) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintChunker(dAtA, i, uint64(len(m.Chunk)))
		i += copy(dAtA[i:], m.Chunk)
	}
	return i, nil
}

func encodeVarintChunker(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Chunk) Size() (n int) {
	var l int
	_ = l
	l = len(m.Chunk)
	if l > 0 {
		n += 1 + l + sovChunker(uint64(l))
	}
	return n
}

func sovChunker(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozChunker(x uint64) (n int) {
	return sovChunker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChunker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipChunker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChunker
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
func (m *Chunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChunker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Chunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChunker
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunk = append(m.Chunk[:0], dAtA[iNdEx:postIndex]...)
			if m.Chunk == nil {
				m.Chunk = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChunker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChunker
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
func skipChunker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChunker
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
					return 0, ErrIntOverflowChunker
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
					return 0, ErrIntOverflowChunker
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthChunker
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowChunker
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
				next, err := skipChunker(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthChunker = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChunker   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("chunker.proto", fileDescriptorChunker) }

var fileDescriptorChunker = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x28, 0xcd,
	0xcb, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x62, 0xe7,
	0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x54, 0x92, 0xe5, 0x62, 0x75, 0x06, 0xc9, 0x0a, 0x89, 0x70,
	0xb1, 0x82, 0x95, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x38, 0x46, 0x7a, 0x5c, 0xec,
	0xce, 0x10, 0xcd, 0x42, 0xca, 0x08, 0x26, 0xa7, 0x5e, 0x41, 0x92, 0x1e, 0x58, 0xbf, 0x14, 0x98,
	0x09, 0x16, 0x57, 0x62, 0x30, 0x60, 0x74, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x5b, 0x6a, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x97, 0x24, 0xb2, 0xcb, 0x85, 0x00, 0x00, 0x00,
}