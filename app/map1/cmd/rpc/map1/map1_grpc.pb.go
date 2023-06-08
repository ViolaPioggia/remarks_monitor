// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: map1.proto

package map1

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

const (
	MapWork_MapWork_FullMethodName = "/MapWork/MapWork"
)

// MapWorkClient is the client API for MapWork service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapWorkClient interface {
	MapWork(ctx context.Context, in *GetMapWorkReq, opts ...grpc.CallOption) (*GetMapWorkResp, error)
}

type mapWorkClient struct {
	cc grpc.ClientConnInterface
}

func NewMapWorkClient(cc grpc.ClientConnInterface) MapWorkClient {
	return &mapWorkClient{cc}
}

func (c *mapWorkClient) MapWork(ctx context.Context, in *GetMapWorkReq, opts ...grpc.CallOption) (*GetMapWorkResp, error) {
	out := new(GetMapWorkResp)
	err := c.cc.Invoke(ctx, MapWork_MapWork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapWorkServer is the server API for MapWork service.
// All implementations must embed UnimplementedMapWorkServer
// for forward compatibility
type MapWorkServer interface {
	MapWork(context.Context, *GetMapWorkReq) (*GetMapWorkResp, error)
	mustEmbedUnimplementedMapWorkServer()
}

// UnimplementedMapWorkServer must be embedded to have forward compatible implementations.
type UnimplementedMapWorkServer struct {
}

func (UnimplementedMapWorkServer) MapWork(context.Context, *GetMapWorkReq) (*GetMapWorkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapWork not implemented")
}
func (UnimplementedMapWorkServer) mustEmbedUnimplementedMapWorkServer() {}

// UnsafeMapWorkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapWorkServer will
// result in compilation errors.
type UnsafeMapWorkServer interface {
	mustEmbedUnimplementedMapWorkServer()
}

func RegisterMapWorkServer(s grpc.ServiceRegistrar, srv MapWorkServer) {
	s.RegisterService(&MapWork_ServiceDesc, srv)
}

func _MapWork_MapWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapWorkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapWorkServer).MapWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MapWork_MapWork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapWorkServer).MapWork(ctx, req.(*GetMapWorkReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MapWork_ServiceDesc is the grpc.ServiceDesc for MapWork service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapWork_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MapWork",
	HandlerType: (*MapWorkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MapWork",
			Handler:    _MapWork_MapWork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "map1.proto",
}
