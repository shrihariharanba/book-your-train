// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: train_ticketing_system.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TrainTicketingService_GetTrainDetails_FullMethodName             = "/TrainTicketingService/GetTrainDetails"
	TrainTicketingService_BookTrainTicket_FullMethodName             = "/TrainTicketingService/BookTrainTicket"
	TrainTicketingService_ModifyTrainSeat_FullMethodName             = "/TrainTicketingService/ModifyTrainSeat"
	TrainTicketingService_GetTicketPurchaseDetails_FullMethodName    = "/TrainTicketingService/GetTicketPurchaseDetails"
	TrainTicketingService_DeleteTicketPurchaseDetails_FullMethodName = "/TrainTicketingService/DeleteTicketPurchaseDetails"
)

// TrainTicketingServiceClient is the client API for TrainTicketingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainTicketingServiceClient interface {
	GetTrainDetails(ctx context.Context, in *TrainDetails, opts ...grpc.CallOption) (*TrainDetails, error)
	BookTrainTicket(ctx context.Context, in *TicketPurchaseDetails, opts ...grpc.CallOption) (*TicketPurchaseDetails, error)
	ModifyTrainSeat(ctx context.Context, in *TicketPurchaseDetails, opts ...grpc.CallOption) (*TicketPurchaseDetails, error)
	GetTicketPurchaseDetails(ctx context.Context, in *TicketPurchaseDetails, opts ...grpc.CallOption) (*TicketPurchaseDetails, error)
	DeleteTicketPurchaseDetails(ctx context.Context, in *TicketPurchaseDetails, opts ...grpc.CallOption) (*TicketPurchaseDetails, error)
}

type trainTicketingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainTicketingServiceClient(cc grpc.ClientConnInterface) TrainTicketingServiceClient {
	return &trainTicketingServiceClient{cc}
}

func (c *trainTicketingServiceClient) GetTrainDetails(ctx context.Context, in *TrainDetails, opts ...grpc.CallOption) (*TrainDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrainDetails)
	err := c.cc.Invoke(ctx, TrainTicketingService_GetTrainDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketingServiceClient) BookTrainTicket(ctx context.Context, in *TicketPurchaseDetails, opts ...grpc.CallOption) (*TicketPurchaseDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketPurchaseDetails)
	err := c.cc.Invoke(ctx, TrainTicketingService_BookTrainTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketingServiceClient) ModifyTrainSeat(ctx context.Context, in *TicketPurchaseDetails, opts ...grpc.CallOption) (*TicketPurchaseDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketPurchaseDetails)
	err := c.cc.Invoke(ctx, TrainTicketingService_ModifyTrainSeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketingServiceClient) GetTicketPurchaseDetails(ctx context.Context, in *TicketPurchaseDetails, opts ...grpc.CallOption) (*TicketPurchaseDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketPurchaseDetails)
	err := c.cc.Invoke(ctx, TrainTicketingService_GetTicketPurchaseDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketingServiceClient) DeleteTicketPurchaseDetails(ctx context.Context, in *TicketPurchaseDetails, opts ...grpc.CallOption) (*TicketPurchaseDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketPurchaseDetails)
	err := c.cc.Invoke(ctx, TrainTicketingService_DeleteTicketPurchaseDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainTicketingServiceServer is the server API for TrainTicketingService service.
// All implementations must embed UnimplementedTrainTicketingServiceServer
// for forward compatibility.
type TrainTicketingServiceServer interface {
	GetTrainDetails(context.Context, *TrainDetails) (*TrainDetails, error)
	BookTrainTicket(context.Context, *TicketPurchaseDetails) (*TicketPurchaseDetails, error)
	ModifyTrainSeat(context.Context, *TicketPurchaseDetails) (*TicketPurchaseDetails, error)
	GetTicketPurchaseDetails(context.Context, *TicketPurchaseDetails) (*TicketPurchaseDetails, error)
	DeleteTicketPurchaseDetails(context.Context, *TicketPurchaseDetails) (*TicketPurchaseDetails, error)
	mustEmbedUnimplementedTrainTicketingServiceServer()
}

// UnimplementedTrainTicketingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrainTicketingServiceServer struct{}

func (UnimplementedTrainTicketingServiceServer) GetTrainDetails(context.Context, *TrainDetails) (*TrainDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrainDetails not implemented")
}
func (UnimplementedTrainTicketingServiceServer) BookTrainTicket(context.Context, *TicketPurchaseDetails) (*TicketPurchaseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTrainTicket not implemented")
}
func (UnimplementedTrainTicketingServiceServer) ModifyTrainSeat(context.Context, *TicketPurchaseDetails) (*TicketPurchaseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyTrainSeat not implemented")
}
func (UnimplementedTrainTicketingServiceServer) GetTicketPurchaseDetails(context.Context, *TicketPurchaseDetails) (*TicketPurchaseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicketPurchaseDetails not implemented")
}
func (UnimplementedTrainTicketingServiceServer) DeleteTicketPurchaseDetails(context.Context, *TicketPurchaseDetails) (*TicketPurchaseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicketPurchaseDetails not implemented")
}
func (UnimplementedTrainTicketingServiceServer) mustEmbedUnimplementedTrainTicketingServiceServer() {}
func (UnimplementedTrainTicketingServiceServer) testEmbeddedByValue()                               {}

// UnsafeTrainTicketingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainTicketingServiceServer will
// result in compilation errors.
type UnsafeTrainTicketingServiceServer interface {
	mustEmbedUnimplementedTrainTicketingServiceServer()
}

func RegisterTrainTicketingServiceServer(s grpc.ServiceRegistrar, srv TrainTicketingServiceServer) {
	// If the following call pancis, it indicates UnimplementedTrainTicketingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrainTicketingService_ServiceDesc, srv)
}

func _TrainTicketingService_GetTrainDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServiceServer).GetTrainDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketingService_GetTrainDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServiceServer).GetTrainDetails(ctx, req.(*TrainDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketingService_BookTrainTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketPurchaseDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServiceServer).BookTrainTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketingService_BookTrainTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServiceServer).BookTrainTicket(ctx, req.(*TicketPurchaseDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketingService_ModifyTrainSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketPurchaseDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServiceServer).ModifyTrainSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketingService_ModifyTrainSeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServiceServer).ModifyTrainSeat(ctx, req.(*TicketPurchaseDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketingService_GetTicketPurchaseDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketPurchaseDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServiceServer).GetTicketPurchaseDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketingService_GetTicketPurchaseDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServiceServer).GetTicketPurchaseDetails(ctx, req.(*TicketPurchaseDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketingService_DeleteTicketPurchaseDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketPurchaseDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServiceServer).DeleteTicketPurchaseDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketingService_DeleteTicketPurchaseDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServiceServer).DeleteTicketPurchaseDetails(ctx, req.(*TicketPurchaseDetails))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainTicketingService_ServiceDesc is the grpc.ServiceDesc for TrainTicketingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainTicketingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TrainTicketingService",
	HandlerType: (*TrainTicketingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTrainDetails",
			Handler:    _TrainTicketingService_GetTrainDetails_Handler,
		},
		{
			MethodName: "BookTrainTicket",
			Handler:    _TrainTicketingService_BookTrainTicket_Handler,
		},
		{
			MethodName: "ModifyTrainSeat",
			Handler:    _TrainTicketingService_ModifyTrainSeat_Handler,
		},
		{
			MethodName: "GetTicketPurchaseDetails",
			Handler:    _TrainTicketingService_GetTicketPurchaseDetails_Handler,
		},
		{
			MethodName: "DeleteTicketPurchaseDetails",
			Handler:    _TrainTicketingService_DeleteTicketPurchaseDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "train_ticketing_system.proto",
}
