// Code generated by protoc-gen-go.
// source: proto/add_better.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/add_better.proto

It has these top-level messages:
	BetterNumericRequest
	BetterNumericResponse
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

type TestType int32

const (
	TestType_TYPE_0 TestType = 0
	TestType_TYPE_1 TestType = 1
)

var TestType_name = map[int32]string{
	0: "TYPE_0",
	1: "TYPE_1",
}
var TestType_value = map[string]int32{
	"TYPE_0": 0,
	"TYPE_1": 1,
}

func (x TestType) String() string {
	return proto1.EnumName(TestType_name, int32(x))
}

type BetterNumericRequest struct {
	Type   TestType `protobuf:"varint,1,opt,name=type,enum=proto.TestType" json:"type,omitempty"`
	Values []int32  `protobuf:"varint,2,rep,name=values" json:"values,omitempty"`
}

func (m *BetterNumericRequest) Reset()         { *m = BetterNumericRequest{} }
func (m *BetterNumericRequest) String() string { return proto1.CompactTextString(m) }
func (*BetterNumericRequest) ProtoMessage()    {}

type BetterNumericResponse struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	R      int32  `protobuf:"varint,2,opt,name=r" json:"r,omitempty"`
}

func (m *BetterNumericResponse) Reset()         { *m = BetterNumericResponse{} }
func (m *BetterNumericResponse) String() string { return proto1.CompactTextString(m) }
func (*BetterNumericResponse) ProtoMessage()    {}

func init() {
	proto1.RegisterEnum("proto.TestType", TestType_name, TestType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for BetterTest service

type BetterTestClient interface {
	Add(ctx context.Context, in *BetterNumericRequest, opts ...grpc.CallOption) (*BetterNumericResponse, error)
}

type betterTestClient struct {
	cc *grpc.ClientConn
}

func NewBetterTestClient(cc *grpc.ClientConn) BetterTestClient {
	return &betterTestClient{cc}
}

func (c *betterTestClient) Add(ctx context.Context, in *BetterNumericRequest, opts ...grpc.CallOption) (*BetterNumericResponse, error) {
	out := new(BetterNumericResponse)
	err := grpc.Invoke(ctx, "/proto.BetterTest/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BetterTest service

type BetterTestServer interface {
	Add(context.Context, *BetterNumericRequest) (*BetterNumericResponse, error)
}

func RegisterBetterTestServer(s *grpc.Server, srv BetterTestServer) {
	s.RegisterService(&_BetterTest_serviceDesc, srv)
}

func _BetterTest_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BetterNumericRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BetterTestServer).Add(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BetterTest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BetterTest",
	HandlerType: (*BetterTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _BetterTest_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}