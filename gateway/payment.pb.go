// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment.proto

/*
Package gateway is a generated protocol buffer package.

It is generated from these files:
	payment.proto

It has these top-level messages:
	PaymentRequest
	PaymentResponse
*/
package gateway

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PaymentRequest struct {
	PaymentId         int64                  `protobuf:"varint,1,opt,name=payment_id,json=paymentId" json:"payment_id,omitempty"`
	Merchant          string                 `protobuf:"bytes,2,opt,name=merchant" json:"merchant,omitempty"`
	MerchantPaymentId string                 `protobuf:"bytes,3,opt,name=merchant_payment_id,json=merchantPaymentId" json:"merchant_payment_id,omitempty"`
	Amount            float64                `protobuf:"fixed64,4,opt,name=amount" json:"amount,omitempty"`
	Fee               float64                `protobuf:"fixed64,5,opt,name=fee" json:"fee,omitempty"`
	User              *PaymentRequest_User   `protobuf:"bytes,6,opt,name=user" json:"user,omitempty"`
	Item              []*PaymentRequest_Item `protobuf:"bytes,7,rep,name=item" json:"item,omitempty"`
}

func (m *PaymentRequest) Reset()                    { *m = PaymentRequest{} }
func (m *PaymentRequest) String() string            { return proto.CompactTextString(m) }
func (*PaymentRequest) ProtoMessage()               {}
func (*PaymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PaymentRequest) GetPaymentId() int64 {
	if m != nil {
		return m.PaymentId
	}
	return 0
}

func (m *PaymentRequest) GetMerchant() string {
	if m != nil {
		return m.Merchant
	}
	return ""
}

func (m *PaymentRequest) GetMerchantPaymentId() string {
	if m != nil {
		return m.MerchantPaymentId
	}
	return ""
}

func (m *PaymentRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PaymentRequest) GetFee() float64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *PaymentRequest) GetUser() *PaymentRequest_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PaymentRequest) GetItem() []*PaymentRequest_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type PaymentRequest_User struct {
	Id    int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *PaymentRequest_User) Reset()                    { *m = PaymentRequest_User{} }
func (m *PaymentRequest_User) String() string            { return proto.CompactTextString(m) }
func (*PaymentRequest_User) ProtoMessage()               {}
func (*PaymentRequest_User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *PaymentRequest_User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PaymentRequest_User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PaymentRequest_User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type PaymentRequest_Item struct {
	Name     string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Price    float64 `protobuf:"fixed64,2,opt,name=price" json:"price,omitempty"`
	Quantity int64   `protobuf:"varint,3,opt,name=quantity" json:"quantity,omitempty"`
}

func (m *PaymentRequest_Item) Reset()                    { *m = PaymentRequest_Item{} }
func (m *PaymentRequest_Item) String() string            { return proto.CompactTextString(m) }
func (*PaymentRequest_Item) ProtoMessage()               {}
func (*PaymentRequest_Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *PaymentRequest_Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PaymentRequest_Item) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *PaymentRequest_Item) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type PaymentResponse struct {
	Status  int32  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PaymentResponse) Reset()                    { *m = PaymentResponse{} }
func (m *PaymentResponse) String() string            { return proto.CompactTextString(m) }
func (*PaymentResponse) ProtoMessage()               {}
func (*PaymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PaymentResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PaymentResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PaymentRequest)(nil), "gateway.PaymentRequest")
	proto.RegisterType((*PaymentRequest_User)(nil), "gateway.PaymentRequest.User")
	proto.RegisterType((*PaymentRequest_Item)(nil), "gateway.PaymentRequest.Item")
	proto.RegisterType((*PaymentResponse)(nil), "gateway.PaymentResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Gateway service

type GatewayClient interface {
	MakePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) MakePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := grpc.Invoke(ctx, "/gateway.Gateway/MakePayment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gateway service

type GatewayServer interface {
	MakePayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_MakePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).MakePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.Gateway/MakePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).MakePayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakePayment",
			Handler:    _Gateway_MakePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}

func init() { proto.RegisterFile("payment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xd1, 0x4a, 0xc3, 0x40,
	0x10, 0xf4, 0x9a, 0xb4, 0xb1, 0x5b, 0xac, 0xba, 0x8a, 0x1e, 0x41, 0x21, 0xf4, 0x29, 0x4f, 0x41,
	0xea, 0x0f, 0x88, 0x3e, 0x48, 0xc1, 0x42, 0x39, 0xf0, 0xb9, 0x9c, 0xed, 0x5a, 0x83, 0x5e, 0x92,
	0xe6, 0x2e, 0x48, 0x3f, 0xc5, 0xbf, 0x95, 0x5c, 0x2e, 0xc1, 0x22, 0xfa, 0xb6, 0x73, 0x3b, 0x33,
	0x99, 0x59, 0x02, 0x47, 0x85, 0xdc, 0x29, 0xca, 0x4c, 0x52, 0x94, 0xb9, 0xc9, 0x31, 0xd8, 0x48,
	0x43, 0x9f, 0x72, 0x37, 0xf9, 0xf2, 0x60, 0xbc, 0x68, 0x56, 0x82, 0xb6, 0x15, 0x69, 0x83, 0xd7,
	0x00, 0x8e, 0xbc, 0x4c, 0xd7, 0x9c, 0x45, 0x2c, 0xf6, 0xc4, 0xd0, 0xbd, 0xcc, 0xd6, 0x18, 0xc2,
	0xa1, 0xa2, 0x72, 0xf5, 0x26, 0x33, 0xc3, 0x7b, 0x11, 0x8b, 0x87, 0xa2, 0xc3, 0x98, 0xc0, 0x59,
	0x3b, 0x2f, 0x7f, 0x78, 0x78, 0x96, 0x76, 0xda, 0xae, 0x16, 0x9d, 0xd7, 0x05, 0x0c, 0xa4, 0xca,
	0xab, 0xcc, 0x70, 0x3f, 0x62, 0x31, 0x13, 0x0e, 0xe1, 0x09, 0x78, 0xaf, 0x44, 0xbc, 0x6f, 0x1f,
	0xeb, 0x11, 0x6f, 0xc0, 0xaf, 0x34, 0x95, 0x7c, 0x10, 0xb1, 0x78, 0x34, 0xbd, 0x4a, 0x5c, 0xfe,
	0x64, 0x3f, 0x7b, 0xf2, 0xac, 0xa9, 0x14, 0x96, 0x59, 0x2b, 0x52, 0x43, 0x8a, 0x07, 0x91, 0xf7,
	0x9f, 0x62, 0x66, 0x48, 0x09, 0xcb, 0x0c, 0xef, 0xc0, 0xaf, 0xf5, 0x38, 0x86, 0x5e, 0x57, 0xbc,
	0x97, 0xae, 0x11, 0xc1, 0xcf, 0xa4, 0x22, 0xd7, 0xd6, 0xce, 0x78, 0x0e, 0x7d, 0x52, 0x32, 0xfd,
	0x70, 0xdd, 0x1a, 0x10, 0x3e, 0x81, 0x5f, 0xfb, 0x75, 0x0a, 0xb6, 0xaf, 0x28, 0xca, 0x74, 0xd5,
	0xd8, 0x30, 0xd1, 0x80, 0xfa, 0x9a, 0xdb, 0x4a, 0x66, 0x26, 0x35, 0x3b, 0x6b, 0xe5, 0x89, 0x0e,
	0x4f, 0x1e, 0xe0, 0xb8, 0x0b, 0xab, 0x8b, 0x3c, 0xd3, 0x54, 0x1f, 0x4c, 0x1b, 0x69, 0x2a, 0x6d,
	0xad, 0xfb, 0xc2, 0x21, 0xe4, 0x10, 0x28, 0xd2, 0x5a, 0x6e, 0xda, 0x94, 0x2d, 0x9c, 0xce, 0x21,
	0x78, 0x6c, 0x9a, 0xe3, 0x3d, 0x8c, 0xe6, 0xf2, 0x9d, 0x9c, 0x27, 0x5e, 0xfe, 0x71, 0x92, 0x90,
	0xff, 0x5e, 0x34, 0x9f, 0x9f, 0x1c, 0xbc, 0x0c, 0xec, 0xff, 0x73, 0xfb, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x68, 0x87, 0xa2, 0xd8, 0x50, 0x02, 0x00, 0x00,
}