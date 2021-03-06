// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package remote

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVClient interface {
	// open a cursor on given position of given bucket
	// if streaming requested - streams all data: stops if client's buffer is full, resumes when client read enough from buffer
	// if streaming not requested - streams next data only when clients sends message to bi-directional channel
	// no full consistency guarantee - server implementation can close/open underlying db transaction at any time
	Seek(ctx context.Context, opts ...grpc.CallOption) (KV_SeekClient, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

var kVSeekStreamDesc = &grpc.StreamDesc{
	StreamName:    "Seek",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *kVClient) Seek(ctx context.Context, opts ...grpc.CallOption) (KV_SeekClient, error) {
	stream, err := c.cc.NewStream(ctx, kVSeekStreamDesc, "/remote.KV/Seek", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVSeekClient{stream}
	return x, nil
}

type KV_SeekClient interface {
	Send(*SeekRequest) error
	Recv() (*Pair, error)
	grpc.ClientStream
}

type kVSeekClient struct {
	grpc.ClientStream
}

func (x *kVSeekClient) Send(m *SeekRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kVSeekClient) Recv() (*Pair, error) {
	m := new(Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVService is the service API for KV service.
// Fields should be assigned to their respective handler implementations only before
// RegisterKVService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type KVService struct {
	// open a cursor on given position of given bucket
	// if streaming requested - streams all data: stops if client's buffer is full, resumes when client read enough from buffer
	// if streaming not requested - streams next data only when clients sends message to bi-directional channel
	// no full consistency guarantee - server implementation can close/open underlying db transaction at any time
	Seek func(KV_SeekServer) error
}

func (s *KVService) seek(_ interface{}, stream grpc.ServerStream) error {
	if s.Seek == nil {
		return status.Errorf(codes.Unimplemented, "method Seek not implemented")
	}
	return s.Seek(&kVSeekServer{stream})
}

type KV_SeekServer interface {
	Send(*Pair) error
	Recv() (*SeekRequest, error)
	grpc.ServerStream
}

type kVSeekServer struct {
	grpc.ServerStream
}

func (x *kVSeekServer) Send(m *Pair) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kVSeekServer) Recv() (*SeekRequest, error) {
	m := new(SeekRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterKVService registers a service implementation with a gRPC server.
func RegisterKVService(s grpc.ServiceRegistrar, srv *KVService) {
	sd := grpc.ServiceDesc{
		ServiceName: "remote.KV",
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Seek",
				Handler:       srv.seek,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "remote/kv.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewKVService creates a new KVService containing the
// implemented methods of the KV service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewKVService(s interface{}) *KVService {
	ns := &KVService{}
	if h, ok := s.(interface{ Seek(KV_SeekServer) error }); ok {
		ns.Seek = h.Seek
	}
	return ns
}

// UnstableKVService is the service API for KV service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableKVService interface {
	// open a cursor on given position of given bucket
	// if streaming requested - streams all data: stops if client's buffer is full, resumes when client read enough from buffer
	// if streaming not requested - streams next data only when clients sends message to bi-directional channel
	// no full consistency guarantee - server implementation can close/open underlying db transaction at any time
	Seek(KV_SeekServer) error
}

// KV2Client is the client API for KV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KV2Client interface {
	Tx(ctx context.Context, opts ...grpc.CallOption) (KV2_TxClient, error)
}

type kV2Client struct {
	cc grpc.ClientConnInterface
}

func NewKV2Client(cc grpc.ClientConnInterface) KV2Client {
	return &kV2Client{cc}
}

var kV2TxStreamDesc = &grpc.StreamDesc{
	StreamName:    "Tx",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *kV2Client) Tx(ctx context.Context, opts ...grpc.CallOption) (KV2_TxClient, error) {
	stream, err := c.cc.NewStream(ctx, kV2TxStreamDesc, "/remote.KV2/Tx", opts...)
	if err != nil {
		return nil, err
	}
	x := &kV2TxClient{stream}
	return x, nil
}

type KV2_TxClient interface {
	Send(*Cursor) error
	Recv() (*Pair2, error)
	grpc.ClientStream
}

type kV2TxClient struct {
	grpc.ClientStream
}

func (x *kV2TxClient) Send(m *Cursor) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kV2TxClient) Recv() (*Pair2, error) {
	m := new(Pair2)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KV2Service is the service API for KV2 service.
// Fields should be assigned to their respective handler implementations only before
// RegisterKV2Service is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type KV2Service struct {
	Tx func(KV2_TxServer) error
}

func (s *KV2Service) tx(_ interface{}, stream grpc.ServerStream) error {
	if s.Tx == nil {
		return status.Errorf(codes.Unimplemented, "method Tx not implemented")
	}
	return s.Tx(&kV2TxServer{stream})
}

type KV2_TxServer interface {
	Send(*Pair2) error
	Recv() (*Cursor, error)
	grpc.ServerStream
}

type kV2TxServer struct {
	grpc.ServerStream
}

func (x *kV2TxServer) Send(m *Pair2) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kV2TxServer) Recv() (*Cursor, error) {
	m := new(Cursor)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterKV2Service registers a service implementation with a gRPC server.
func RegisterKV2Service(s grpc.ServiceRegistrar, srv *KV2Service) {
	sd := grpc.ServiceDesc{
		ServiceName: "remote.KV2",
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Tx",
				Handler:       srv.tx,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "remote/kv.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewKV2Service creates a new KV2Service containing the
// implemented methods of the KV2 service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewKV2Service(s interface{}) *KV2Service {
	ns := &KV2Service{}
	if h, ok := s.(interface{ Tx(KV2_TxServer) error }); ok {
		ns.Tx = h.Tx
	}
	return ns
}

// UnstableKV2Service is the service API for KV2 service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableKV2Service interface {
	Tx(KV2_TxServer) error
}
