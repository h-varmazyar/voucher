// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/proto/src/voucher.proto

package voucherApi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VoucherServiceClient is the client API for VoucherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VoucherServiceClient interface {
	Create(ctx context.Context, in *VoucherCreateReq, opts ...grpc.CallOption) (*Voucher, error)
	Apply(ctx context.Context, in *VoucherApplyReq, opts ...grpc.CallOption) (*Void, error)
	Usage(ctx context.Context, in *VoucherUsageReq, opts ...grpc.CallOption) (*Usages, error)
}

type voucherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoucherServiceClient(cc grpc.ClientConnInterface) VoucherServiceClient {
	return &voucherServiceClient{cc}
}

func (c *voucherServiceClient) Create(ctx context.Context, in *VoucherCreateReq, opts ...grpc.CallOption) (*Voucher, error) {
	out := new(Voucher)
	err := c.cc.Invoke(ctx, "/voucher_api.VoucherService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voucherServiceClient) Apply(ctx context.Context, in *VoucherApplyReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/voucher_api.VoucherService/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voucherServiceClient) Usage(ctx context.Context, in *VoucherUsageReq, opts ...grpc.CallOption) (*Usages, error) {
	out := new(Usages)
	err := c.cc.Invoke(ctx, "/voucher_api.VoucherService/Usage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoucherServiceServer is the server API for VoucherService service.
// All implementations must embed UnimplementedVoucherServiceServer
// for forward compatibility
type VoucherServiceServer interface {
	Create(context.Context, *VoucherCreateReq) (*Voucher, error)
	Apply(context.Context, *VoucherApplyReq) (*Void, error)
	Usage(context.Context, *VoucherUsageReq) (*Usages, error)
	mustEmbedUnimplementedVoucherServiceServer()
}

// UnimplementedVoucherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVoucherServiceServer struct {
}

func (UnimplementedVoucherServiceServer) Create(context.Context, *VoucherCreateReq) (*Voucher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVoucherServiceServer) Apply(context.Context, *VoucherApplyReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedVoucherServiceServer) Usage(context.Context, *VoucherUsageReq) (*Usages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Usage not implemented")
}
func (UnimplementedVoucherServiceServer) mustEmbedUnimplementedVoucherServiceServer() {}

// UnsafeVoucherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VoucherServiceServer will
// result in compilation errors.
type UnsafeVoucherServiceServer interface {
	mustEmbedUnimplementedVoucherServiceServer()
}

func RegisterVoucherServiceServer(s grpc.ServiceRegistrar, srv VoucherServiceServer) {
	s.RegisterService(&VoucherService_ServiceDesc, srv)
}

func _VoucherService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoucherCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voucher_api.VoucherService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServiceServer).Create(ctx, req.(*VoucherCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoucherService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoucherApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voucher_api.VoucherService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServiceServer).Apply(ctx, req.(*VoucherApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoucherService_Usage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoucherUsageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServiceServer).Usage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voucher_api.VoucherService/Usage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServiceServer).Usage(ctx, req.(*VoucherUsageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VoucherService_ServiceDesc is the grpc.ServiceDesc for VoucherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VoucherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voucher_api.VoucherService",
	HandlerType: (*VoucherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VoucherService_Create_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _VoucherService_Apply_Handler,
		},
		{
			MethodName: "Usage",
			Handler:    _VoucherService_Usage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/src/voucher.proto",
}
