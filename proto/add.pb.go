// Code generated by protoc-gen-go.
// source: proto/add.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/add.proto

It has these top-level messages:
	NumericRequest
	NumericResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NumericRequest struct {
	V1 int32 `protobuf:"varint,1,opt,name=v1" json:"v1,omitempty"`
	V2 int32 `protobuf:"varint,2,opt,name=v2" json:"v2,omitempty"`
}

func (m *NumericRequest) Reset()         { *m = NumericRequest{} }
func (m *NumericRequest) String() string { return proto1.CompactTextString(m) }
func (*NumericRequest) ProtoMessage()    {}

type NumericResponse struct {
	R int32 `protobuf:"varint,1,opt,name=r" json:"r,omitempty"`
}

func (m *NumericResponse) Reset()         { *m = NumericResponse{} }
func (m *NumericResponse) String() string { return proto1.CompactTextString(m) }
func (*NumericResponse) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Test service

type TestClient interface {
	Add(ctx context.Context, in *NumericRequest, opts ...grpc.CallOption) (*NumericResponse, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) Add(ctx context.Context, in *NumericRequest, opts ...grpc.CallOption) (*NumericResponse, error) {
	out := new(NumericResponse)
	err := grpc.Invoke(ctx, "/proto.Test/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Test service

type TestServer interface {
	Add(context.Context, *NumericRequest) (*NumericResponse, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NumericRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TestServer).Add(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Test_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}