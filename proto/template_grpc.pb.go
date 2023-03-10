// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateClient interface {
	// one message is sent and one is recieved
	Put(ctx context.Context, in *KeyValPair, opts ...grpc.CallOption) (*PutAck, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*GetAck, error)
}

type templateClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateClient(cc grpc.ClientConnInterface) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) Put(ctx context.Context, in *KeyValPair, opts ...grpc.CallOption) (*PutAck, error) {
	out := new(PutAck)
	err := c.cc.Invoke(ctx, "/proto.Template/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*GetAck, error) {
	out := new(GetAck)
	err := c.cc.Invoke(ctx, "/proto.Template/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServer is the server API for Template service.
// All implementations must embed UnimplementedTemplateServer
// for forward compatibility
type TemplateServer interface {
	// one message is sent and one is recieved
	Put(context.Context, *KeyValPair) (*PutAck, error)
	Get(context.Context, *Key) (*GetAck, error)
	mustEmbedUnimplementedTemplateServer()
}

// UnimplementedTemplateServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateServer struct {
}

func (UnimplementedTemplateServer) Put(context.Context, *KeyValPair) (*PutAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedTemplateServer) Get(context.Context, *Key) (*GetAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTemplateServer) mustEmbedUnimplementedTemplateServer() {}

// UnsafeTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServer will
// result in compilation errors.
type UnsafeTemplateServer interface {
	mustEmbedUnimplementedTemplateServer()
}

func RegisterTemplateServer(s grpc.ServiceRegistrar, srv TemplateServer) {
	s.RegisterService(&Template_ServiceDesc, srv)
}

func _Template_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Put(ctx, req.(*KeyValPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// Template_ServiceDesc is the grpc.ServiceDesc for Template service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Template_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Template_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Template_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/template.proto",
}
