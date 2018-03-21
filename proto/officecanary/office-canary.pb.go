// Code generated by protoc-gen-go. DO NOT EDIT.
// source: office-canary.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	office-canary.proto

It has these top-level messages:
	GetPublicDeviceRequest
	GetPublicDeviceResponse
	GetPublicDatapointsRequest
	GetPublicDatapointsResponse
	StreamPublicDatapointsRequest
	StreamPublicDatapointsResponse
	PublicDevice
	PublicDatapoint
	Datapoint
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type CO2Status int32

const (
	CO2Status_OK    CO2Status = 0
	CO2Status_Busy  CO2Status = 1
	CO2Status_Runin CO2Status = 16
	CO2Status_Error CO2Status = 128
)

var CO2Status_name = map[int32]string{
	0:   "OK",
	1:   "Busy",
	16:  "Runin",
	128: "Error",
}
var CO2Status_value = map[string]int32{
	"OK":    0,
	"Busy":  1,
	"Runin": 16,
	"Error": 128,
}

func (x CO2Status) String() string {
	return proto1.EnumName(CO2Status_name, int32(x))
}
func (CO2Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetPublicDeviceRequest struct {
	Eui string `protobuf:"bytes,1,opt,name=eui" json:"eui,omitempty"`
}

func (m *GetPublicDeviceRequest) Reset()                    { *m = GetPublicDeviceRequest{} }
func (m *GetPublicDeviceRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetPublicDeviceRequest) ProtoMessage()               {}
func (*GetPublicDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPublicDeviceRequest) GetEui() string {
	if m != nil {
		return m.Eui
	}
	return ""
}

type GetPublicDeviceResponse struct {
	Device *PublicDevice `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
}

func (m *GetPublicDeviceResponse) Reset()                    { *m = GetPublicDeviceResponse{} }
func (m *GetPublicDeviceResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetPublicDeviceResponse) ProtoMessage()               {}
func (*GetPublicDeviceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetPublicDeviceResponse) GetDevice() *PublicDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

type GetPublicDatapointsRequest struct {
	Eui string `protobuf:"bytes,1,opt,name=eui" json:"eui,omitempty"`
}

func (m *GetPublicDatapointsRequest) Reset()                    { *m = GetPublicDatapointsRequest{} }
func (m *GetPublicDatapointsRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetPublicDatapointsRequest) ProtoMessage()               {}
func (*GetPublicDatapointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetPublicDatapointsRequest) GetEui() string {
	if m != nil {
		return m.Eui
	}
	return ""
}

type GetPublicDatapointsResponse struct {
	Datapoints []*PublicDatapoint `protobuf:"bytes,1,rep,name=datapoints" json:"datapoints,omitempty"`
}

func (m *GetPublicDatapointsResponse) Reset()                    { *m = GetPublicDatapointsResponse{} }
func (m *GetPublicDatapointsResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetPublicDatapointsResponse) ProtoMessage()               {}
func (*GetPublicDatapointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetPublicDatapointsResponse) GetDatapoints() []*PublicDatapoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

type StreamPublicDatapointsRequest struct {
	DeviceEui string `protobuf:"bytes,1,opt,name=device_eui,json=deviceEui" json:"device_eui,omitempty"`
}

func (m *StreamPublicDatapointsRequest) Reset()                    { *m = StreamPublicDatapointsRequest{} }
func (m *StreamPublicDatapointsRequest) String() string            { return proto1.CompactTextString(m) }
func (*StreamPublicDatapointsRequest) ProtoMessage()               {}
func (*StreamPublicDatapointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StreamPublicDatapointsRequest) GetDeviceEui() string {
	if m != nil {
		return m.DeviceEui
	}
	return ""
}

type StreamPublicDatapointsResponse struct {
	Datapoint *PublicDatapoint `protobuf:"bytes,1,opt,name=datapoint" json:"datapoint,omitempty"`
}

func (m *StreamPublicDatapointsResponse) Reset()                    { *m = StreamPublicDatapointsResponse{} }
func (m *StreamPublicDatapointsResponse) String() string            { return proto1.CompactTextString(m) }
func (*StreamPublicDatapointsResponse) ProtoMessage()               {}
func (*StreamPublicDatapointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StreamPublicDatapointsResponse) GetDatapoint() *PublicDatapoint {
	if m != nil {
		return m.Datapoint
	}
	return nil
}

type PublicDevice struct {
	Eui  string `protobuf:"bytes,1,opt,name=eui" json:"eui,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PublicDevice) Reset()                    { *m = PublicDevice{} }
func (m *PublicDevice) String() string            { return proto1.CompactTextString(m) }
func (*PublicDevice) ProtoMessage()               {}
func (*PublicDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PublicDevice) GetEui() string {
	if m != nil {
		return m.Eui
	}
	return ""
}

func (m *PublicDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PublicDatapoint struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Co2Ppm    uint32                     `protobuf:"varint,2,opt,name=co2_ppm,json=co2Ppm" json:"co2_ppm,omitempty"`
}

func (m *PublicDatapoint) Reset()                    { *m = PublicDatapoint{} }
func (m *PublicDatapoint) String() string            { return proto1.CompactTextString(m) }
func (*PublicDatapoint) ProtoMessage()               {}
func (*PublicDatapoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PublicDatapoint) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PublicDatapoint) GetCo2Ppm() uint32 {
	if m != nil {
		return m.Co2Ppm
	}
	return 0
}

type Datapoint struct {
	AppEui     string    `protobuf:"bytes,4,opt,name=app_eui,json=appEui" json:"app_eui,omitempty"`
	GatewayEui string    `protobuf:"bytes,5,opt,name=gateway_eui,json=gatewayEui" json:"gateway_eui,omitempty"`
	DataRate   string    `protobuf:"bytes,6,opt,name=data_rate,json=dataRate" json:"data_rate,omitempty"`
	DevAddr    string    `protobuf:"bytes,7,opt,name=dev_addr,json=devAddr" json:"dev_addr,omitempty"`
	Frequency  float64   `protobuf:"fixed64,8,opt,name=frequency" json:"frequency,omitempty"`
	Rssi       int32     `protobuf:"varint,9,opt,name=rssi" json:"rssi,omitempty"`
	Snr        float64   `protobuf:"fixed64,10,opt,name=snr" json:"snr,omitempty"`
	Payload    []byte    `protobuf:"bytes,11,opt,name=payload,proto3" json:"payload,omitempty"`
	Co2Ppm     uint32    `protobuf:"varint,12,opt,name=co2_ppm,json=co2Ppm" json:"co2_ppm,omitempty"`
	Co2Status  CO2Status `protobuf:"varint,13,opt,name=co2_status,json=co2Status,enum=office_canary.CO2Status" json:"co2_status,omitempty"`
	Resistance uint32    `protobuf:"varint,14,opt,name=resistance" json:"resistance,omitempty"`
	// Total Volatile Organic Compound equivalent in Parts Per Billion
	TvocPpb uint32 `protobuf:"varint,15,opt,name=tvoc_ppb,json=tvocPpb" json:"tvoc_ppb,omitempty"`
}

func (m *Datapoint) Reset()                    { *m = Datapoint{} }
func (m *Datapoint) String() string            { return proto1.CompactTextString(m) }
func (*Datapoint) ProtoMessage()               {}
func (*Datapoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Datapoint) GetAppEui() string {
	if m != nil {
		return m.AppEui
	}
	return ""
}

func (m *Datapoint) GetGatewayEui() string {
	if m != nil {
		return m.GatewayEui
	}
	return ""
}

func (m *Datapoint) GetDataRate() string {
	if m != nil {
		return m.DataRate
	}
	return ""
}

func (m *Datapoint) GetDevAddr() string {
	if m != nil {
		return m.DevAddr
	}
	return ""
}

func (m *Datapoint) GetFrequency() float64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *Datapoint) GetRssi() int32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *Datapoint) GetSnr() float64 {
	if m != nil {
		return m.Snr
	}
	return 0
}

func (m *Datapoint) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Datapoint) GetCo2Ppm() uint32 {
	if m != nil {
		return m.Co2Ppm
	}
	return 0
}

func (m *Datapoint) GetCo2Status() CO2Status {
	if m != nil {
		return m.Co2Status
	}
	return CO2Status_OK
}

func (m *Datapoint) GetResistance() uint32 {
	if m != nil {
		return m.Resistance
	}
	return 0
}

func (m *Datapoint) GetTvocPpb() uint32 {
	if m != nil {
		return m.TvocPpb
	}
	return 0
}

func init() {
	proto1.RegisterType((*GetPublicDeviceRequest)(nil), "office_canary.GetPublicDeviceRequest")
	proto1.RegisterType((*GetPublicDeviceResponse)(nil), "office_canary.GetPublicDeviceResponse")
	proto1.RegisterType((*GetPublicDatapointsRequest)(nil), "office_canary.GetPublicDatapointsRequest")
	proto1.RegisterType((*GetPublicDatapointsResponse)(nil), "office_canary.GetPublicDatapointsResponse")
	proto1.RegisterType((*StreamPublicDatapointsRequest)(nil), "office_canary.StreamPublicDatapointsRequest")
	proto1.RegisterType((*StreamPublicDatapointsResponse)(nil), "office_canary.StreamPublicDatapointsResponse")
	proto1.RegisterType((*PublicDevice)(nil), "office_canary.PublicDevice")
	proto1.RegisterType((*PublicDatapoint)(nil), "office_canary.PublicDatapoint")
	proto1.RegisterType((*Datapoint)(nil), "office_canary.Datapoint")
	proto1.RegisterEnum("office_canary.CO2Status", CO2Status_name, CO2Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OfficeCanary service

type OfficeCanaryClient interface {
	GetPublicDevice(ctx context.Context, in *GetPublicDeviceRequest, opts ...grpc.CallOption) (*GetPublicDeviceResponse, error)
	GetPublicDatapoints(ctx context.Context, in *GetPublicDatapointsRequest, opts ...grpc.CallOption) (*GetPublicDatapointsResponse, error)
	StreamPublicDatapoints(ctx context.Context, in *StreamPublicDatapointsRequest, opts ...grpc.CallOption) (OfficeCanary_StreamPublicDatapointsClient, error)
}

type officeCanaryClient struct {
	cc *grpc.ClientConn
}

func NewOfficeCanaryClient(cc *grpc.ClientConn) OfficeCanaryClient {
	return &officeCanaryClient{cc}
}

func (c *officeCanaryClient) GetPublicDevice(ctx context.Context, in *GetPublicDeviceRequest, opts ...grpc.CallOption) (*GetPublicDeviceResponse, error) {
	out := new(GetPublicDeviceResponse)
	err := grpc.Invoke(ctx, "/office_canary.OfficeCanary/GetPublicDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *officeCanaryClient) GetPublicDatapoints(ctx context.Context, in *GetPublicDatapointsRequest, opts ...grpc.CallOption) (*GetPublicDatapointsResponse, error) {
	out := new(GetPublicDatapointsResponse)
	err := grpc.Invoke(ctx, "/office_canary.OfficeCanary/GetPublicDatapoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *officeCanaryClient) StreamPublicDatapoints(ctx context.Context, in *StreamPublicDatapointsRequest, opts ...grpc.CallOption) (OfficeCanary_StreamPublicDatapointsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_OfficeCanary_serviceDesc.Streams[0], c.cc, "/office_canary.OfficeCanary/StreamPublicDatapoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &officeCanaryStreamPublicDatapointsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OfficeCanary_StreamPublicDatapointsClient interface {
	Recv() (*StreamPublicDatapointsResponse, error)
	grpc.ClientStream
}

type officeCanaryStreamPublicDatapointsClient struct {
	grpc.ClientStream
}

func (x *officeCanaryStreamPublicDatapointsClient) Recv() (*StreamPublicDatapointsResponse, error) {
	m := new(StreamPublicDatapointsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for OfficeCanary service

type OfficeCanaryServer interface {
	GetPublicDevice(context.Context, *GetPublicDeviceRequest) (*GetPublicDeviceResponse, error)
	GetPublicDatapoints(context.Context, *GetPublicDatapointsRequest) (*GetPublicDatapointsResponse, error)
	StreamPublicDatapoints(*StreamPublicDatapointsRequest, OfficeCanary_StreamPublicDatapointsServer) error
}

func RegisterOfficeCanaryServer(s *grpc.Server, srv OfficeCanaryServer) {
	s.RegisterService(&_OfficeCanary_serviceDesc, srv)
}

func _OfficeCanary_GetPublicDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeCanaryServer).GetPublicDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/office_canary.OfficeCanary/GetPublicDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeCanaryServer).GetPublicDevice(ctx, req.(*GetPublicDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfficeCanary_GetPublicDatapoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicDatapointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficeCanaryServer).GetPublicDatapoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/office_canary.OfficeCanary/GetPublicDatapoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficeCanaryServer).GetPublicDatapoints(ctx, req.(*GetPublicDatapointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfficeCanary_StreamPublicDatapoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamPublicDatapointsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OfficeCanaryServer).StreamPublicDatapoints(m, &officeCanaryStreamPublicDatapointsServer{stream})
}

type OfficeCanary_StreamPublicDatapointsServer interface {
	Send(*StreamPublicDatapointsResponse) error
	grpc.ServerStream
}

type officeCanaryStreamPublicDatapointsServer struct {
	grpc.ServerStream
}

func (x *officeCanaryStreamPublicDatapointsServer) Send(m *StreamPublicDatapointsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _OfficeCanary_serviceDesc = grpc.ServiceDesc{
	ServiceName: "office_canary.OfficeCanary",
	HandlerType: (*OfficeCanaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicDevice",
			Handler:    _OfficeCanary_GetPublicDevice_Handler,
		},
		{
			MethodName: "GetPublicDatapoints",
			Handler:    _OfficeCanary_GetPublicDatapoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPublicDatapoints",
			Handler:       _OfficeCanary_StreamPublicDatapoints_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "office-canary.proto",
}

func init() { proto1.RegisterFile("office-canary.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x4f, 0xdb, 0x30,
	0x14, 0x25, 0x85, 0x7e, 0xe4, 0xb6, 0x40, 0x67, 0x24, 0xf0, 0xca, 0x80, 0x2a, 0xd2, 0xa6, 0x0e,
	0x8d, 0x30, 0x05, 0xa4, 0xed, 0x61, 0x42, 0x1a, 0x0c, 0xed, 0x61, 0xd2, 0x5a, 0x85, 0xbd, 0x6e,
	0x95, 0x9b, 0xb8, 0x95, 0xa5, 0x26, 0xf6, 0x6c, 0xa7, 0x53, 0xdf, 0xf6, 0x47, 0xf7, 0xbe, 0x9f,
	0x31, 0xc5, 0x49, 0x3f, 0x08, 0xed, 0xe0, 0xed, 0xe6, 0xdc, 0xe3, 0x7b, 0x8f, 0x8f, 0xef, 0x0d,
	0xec, 0xf1, 0xe1, 0x90, 0x05, 0xf4, 0x2c, 0x20, 0x31, 0x91, 0x53, 0x57, 0x48, 0xae, 0x39, 0xda,
	0xce, 0xc0, 0x7e, 0x06, 0xb6, 0x4e, 0x46, 0x9c, 0x8f, 0xc6, 0xf4, 0xdc, 0x24, 0x07, 0xc9, 0xf0,
	0x5c, 0xb3, 0x88, 0x2a, 0x4d, 0x22, 0x91, 0xf1, 0x9d, 0x53, 0xd8, 0xff, 0x4c, 0x75, 0x2f, 0x19,
	0x8c, 0x59, 0xf0, 0x89, 0x4e, 0x58, 0x40, 0x7d, 0xfa, 0x33, 0xa1, 0x4a, 0xa3, 0x26, 0x6c, 0xd2,
	0x84, 0x61, 0xab, 0x6d, 0x75, 0x6c, 0x3f, 0x0d, 0x9d, 0xaf, 0x70, 0xf0, 0x80, 0xab, 0x04, 0x8f,
	0x15, 0x45, 0x17, 0x50, 0x09, 0x0d, 0x62, 0xf8, 0x75, 0xef, 0xd0, 0xbd, 0xa7, 0xc3, 0xbd, 0x77,
	0x28, 0xa7, 0x3a, 0x2e, 0xb4, 0x16, 0xf5, 0x88, 0x26, 0x82, 0xb3, 0x58, 0xab, 0xf5, 0xfd, 0xbf,
	0xc3, 0xe1, 0x4a, 0x7e, 0xae, 0xe1, 0x0a, 0x20, 0x9c, 0xa3, 0xd8, 0x6a, 0x6f, 0x76, 0xea, 0xde,
	0xf1, 0x6a, 0x1d, 0x33, 0x9a, 0xbf, 0x74, 0xc2, 0xb9, 0x82, 0xa3, 0x3b, 0x2d, 0x29, 0x89, 0xd6,
	0x29, 0x3a, 0x02, 0xc8, 0x94, 0xf7, 0x17, 0xc2, 0xec, 0x0c, 0xb9, 0x4d, 0x98, 0xf3, 0x03, 0x8e,
	0xd7, 0x9d, 0xcf, 0x15, 0x7e, 0x00, 0x7b, 0xde, 0x2f, 0x37, 0xea, 0x31, 0x81, 0x8b, 0x03, 0xce,
	0x25, 0x34, 0x96, 0x6d, 0x7c, 0x68, 0x10, 0x42, 0xb0, 0x15, 0x93, 0x88, 0xe2, 0x92, 0x81, 0x4c,
	0xec, 0x84, 0xb0, 0x5b, 0xa8, 0x89, 0xde, 0x83, 0x3d, 0x1f, 0x83, 0x5c, 0x46, 0xcb, 0xcd, 0x06,
	0xc5, 0x9d, 0x0d, 0x8a, 0xfb, 0x6d, 0xc6, 0xf0, 0x17, 0x64, 0x74, 0x00, 0xd5, 0x80, 0x7b, 0x7d,
	0x21, 0x22, 0xd3, 0x63, 0xdb, 0xaf, 0x04, 0xdc, 0xeb, 0x89, 0xc8, 0xf9, 0x5b, 0x02, 0x7b, 0xd1,
	0xe0, 0x00, 0xaa, 0x44, 0x08, 0xe3, 0xd2, 0x96, 0x91, 0x52, 0x21, 0x42, 0xdc, 0x26, 0x0c, 0x9d,
	0x40, 0x7d, 0x44, 0x34, 0xfd, 0x45, 0xa6, 0x26, 0x59, 0x36, 0x49, 0xc8, 0xa1, 0x94, 0x70, 0x98,
	0x39, 0xd4, 0x97, 0x44, 0x53, 0x5c, 0x31, 0xe9, 0x5a, 0x0a, 0xf8, 0x44, 0x53, 0xf4, 0x1c, 0x6a,
	0x21, 0x9d, 0xf4, 0x49, 0x18, 0x4a, 0x5c, 0x35, 0xb9, 0x6a, 0x48, 0x27, 0x1f, 0xc3, 0x50, 0xa2,
	0x17, 0x60, 0x0f, 0x65, 0xfa, 0x4c, 0x71, 0x30, 0xc5, 0xb5, 0xb6, 0xd5, 0xb1, 0xfc, 0x05, 0x90,
	0xfa, 0x22, 0x95, 0x62, 0xd8, 0x6e, 0x5b, 0x9d, 0xb2, 0x6f, 0xe2, 0xd4, 0x3d, 0x15, 0x4b, 0x0c,
	0x86, 0x9b, 0x86, 0x08, 0x43, 0x55, 0x90, 0xe9, 0x98, 0x93, 0x10, 0xd7, 0xdb, 0x56, 0xa7, 0xe1,
	0xcf, 0x3e, 0x97, 0xaf, 0xdd, 0x58, 0xbe, 0x36, 0x7a, 0x07, 0x90, 0x26, 0x94, 0x26, 0x3a, 0x51,
	0x78, 0xbb, 0x6d, 0x75, 0x76, 0x3c, 0x5c, 0x78, 0xd1, 0x9b, 0xae, 0x77, 0x67, 0xf2, 0xbe, 0x1d,
	0xf0, 0x3c, 0x44, 0xc7, 0x00, 0x92, 0x2a, 0xa6, 0x34, 0x89, 0x03, 0x8a, 0x77, 0x4c, 0xd1, 0x25,
	0x24, 0xbd, 0xaa, 0x9e, 0xf0, 0xa0, 0x2f, 0xc4, 0x00, 0xef, 0x9a, 0x6c, 0x35, 0xfd, 0xee, 0x89,
	0xc1, 0xe9, 0x25, 0xd8, 0xf3, 0x92, 0xa8, 0x02, 0xa5, 0xee, 0x97, 0xe6, 0x06, 0xaa, 0xc1, 0xd6,
	0x75, 0xa2, 0xa6, 0x4d, 0x0b, 0xd9, 0x50, 0xf6, 0x93, 0x98, 0xc5, 0xcd, 0x26, 0x02, 0x28, 0xdf,
	0x4a, 0xc9, 0x65, 0xf3, 0xb7, 0xe5, 0xfd, 0x29, 0x41, 0xa3, 0x6b, 0x74, 0xdd, 0x18, 0x59, 0x68,
	0x00, 0xbb, 0x85, 0x65, 0x46, 0x2f, 0x0b, 0xca, 0x57, 0xff, 0x18, 0x5a, 0xaf, 0x1e, 0xa3, 0x65,
	0xd3, 0xee, 0x6c, 0xa0, 0x18, 0xf6, 0x56, 0x2c, 0x2c, 0x7a, 0xbd, 0xb6, 0x40, 0x71, 0xe5, 0x5a,
	0xa7, 0x4f, 0xa1, 0xce, 0xfb, 0x4d, 0x61, 0x7f, 0xf5, 0x06, 0xa2, 0x37, 0x85, 0x3a, 0xff, 0x5d,
	0xf4, 0xd6, 0xd9, 0x13, 0xd9, 0xb3, 0xc6, 0x6f, 0xad, 0xeb, 0x4b, 0x38, 0x09, 0x78, 0xe4, 0x6a,
	0x3a, 0xa6, 0x31, 0x97, 0x21, 0x1b, 0x31, 0x4d, 0xc6, 0x79, 0x99, 0xac, 0xca, 0xf5, 0xb3, 0x65,
	0xff, 0x7b, 0x66, 0xcf, 0x2a, 0x66, 0xdd, 0x2e, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x45,
	0xe2, 0x5b, 0xcb, 0x05, 0x00, 0x00,
}
