// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: mp3_SDFS.proto

package protobuf

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
	SDFS_GetFile_FullMethodName            = "/cs425_mp3.SDFS/GetFile"
	SDFS_PutFile_FullMethodName            = "/cs425_mp3.SDFS/PutFile"
	SDFS_DeleteFileLeader_FullMethodName   = "/cs425_mp3.SDFS/DeleteFileLeader"
	SDFS_DeleteFileFollower_FullMethodName = "/cs425_mp3.SDFS/DeleteFileFollower"
	SDFS_LsFile_FullMethodName             = "/cs425_mp3.SDFS/LsFile"
	SDFS_StoreFile_FullMethodName          = "/cs425_mp3.SDFS/StoreFile"
	SDFS_ReplicateFile_FullMethodName      = "/cs425_mp3.SDFS/ReplicateFile"
)

// SDFSClient is the client API for SDFS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SDFSClient interface {
	GetFile(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	PutFile(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	DeleteFileLeader(ctx context.Context, in *DeleteRequestLeader, opts ...grpc.CallOption) (*DeleteResponseLeader, error)
	DeleteFileFollower(ctx context.Context, in *DeleteRequestFollower, opts ...grpc.CallOption) (*DeleteResponseFollower, error)
	LsFile(ctx context.Context, in *LsRequest, opts ...grpc.CallOption) (*LsResponse, error)
	StoreFile(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	ReplicateFile(ctx context.Context, in *ReplicationRequest, opts ...grpc.CallOption) (*ReplicationResponse, error)
}

type sDFSClient struct {
	cc grpc.ClientConnInterface
}

func NewSDFSClient(cc grpc.ClientConnInterface) SDFSClient {
	return &sDFSClient{cc}
}

func (c *sDFSClient) GetFile(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, SDFS_GetFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDFSClient) PutFile(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, SDFS_PutFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDFSClient) DeleteFileLeader(ctx context.Context, in *DeleteRequestLeader, opts ...grpc.CallOption) (*DeleteResponseLeader, error) {
	out := new(DeleteResponseLeader)
	err := c.cc.Invoke(ctx, SDFS_DeleteFileLeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDFSClient) DeleteFileFollower(ctx context.Context, in *DeleteRequestFollower, opts ...grpc.CallOption) (*DeleteResponseFollower, error) {
	out := new(DeleteResponseFollower)
	err := c.cc.Invoke(ctx, SDFS_DeleteFileFollower_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDFSClient) LsFile(ctx context.Context, in *LsRequest, opts ...grpc.CallOption) (*LsResponse, error) {
	out := new(LsResponse)
	err := c.cc.Invoke(ctx, SDFS_LsFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDFSClient) StoreFile(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, SDFS_StoreFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDFSClient) ReplicateFile(ctx context.Context, in *ReplicationRequest, opts ...grpc.CallOption) (*ReplicationResponse, error) {
	out := new(ReplicationResponse)
	err := c.cc.Invoke(ctx, SDFS_ReplicateFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SDFSServer is the server API for SDFS service.
// All implementations must embed UnimplementedSDFSServer
// for forward compatibility
type SDFSServer interface {
	GetFile(context.Context, *GetRequest) (*GetResponse, error)
	PutFile(context.Context, *PutRequest) (*PutResponse, error)
	DeleteFileLeader(context.Context, *DeleteRequestLeader) (*DeleteResponseLeader, error)
	DeleteFileFollower(context.Context, *DeleteRequestFollower) (*DeleteResponseFollower, error)
	LsFile(context.Context, *LsRequest) (*LsResponse, error)
	StoreFile(context.Context, *StoreRequest) (*StoreResponse, error)
	ReplicateFile(context.Context, *ReplicationRequest) (*ReplicationResponse, error)
	mustEmbedUnimplementedSDFSServer()
}

// UnimplementedSDFSServer must be embedded to have forward compatible implementations.
type UnimplementedSDFSServer struct {
}

func (UnimplementedSDFSServer) GetFile(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedSDFSServer) PutFile(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutFile not implemented")
}
func (UnimplementedSDFSServer) DeleteFileLeader(context.Context, *DeleteRequestLeader) (*DeleteResponseLeader, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileLeader not implemented")
}
func (UnimplementedSDFSServer) DeleteFileFollower(context.Context, *DeleteRequestFollower) (*DeleteResponseFollower, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileFollower not implemented")
}
func (UnimplementedSDFSServer) LsFile(context.Context, *LsRequest) (*LsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LsFile not implemented")
}
func (UnimplementedSDFSServer) StoreFile(context.Context, *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFile not implemented")
}
func (UnimplementedSDFSServer) ReplicateFile(context.Context, *ReplicationRequest) (*ReplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateFile not implemented")
}
func (UnimplementedSDFSServer) mustEmbedUnimplementedSDFSServer() {}

// UnsafeSDFSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SDFSServer will
// result in compilation errors.
type UnsafeSDFSServer interface {
	mustEmbedUnimplementedSDFSServer()
}

func RegisterSDFSServer(s grpc.ServiceRegistrar, srv SDFSServer) {
	s.RegisterService(&SDFS_ServiceDesc, srv)
}

func _SDFS_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDFSServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SDFS_GetFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDFSServer).GetFile(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDFS_PutFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDFSServer).PutFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SDFS_PutFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDFSServer).PutFile(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDFS_DeleteFileLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequestLeader)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDFSServer).DeleteFileLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SDFS_DeleteFileLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDFSServer).DeleteFileLeader(ctx, req.(*DeleteRequestLeader))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDFS_DeleteFileFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequestFollower)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDFSServer).DeleteFileFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SDFS_DeleteFileFollower_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDFSServer).DeleteFileFollower(ctx, req.(*DeleteRequestFollower))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDFS_LsFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDFSServer).LsFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SDFS_LsFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDFSServer).LsFile(ctx, req.(*LsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDFS_StoreFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDFSServer).StoreFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SDFS_StoreFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDFSServer).StoreFile(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDFS_ReplicateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDFSServer).ReplicateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SDFS_ReplicateFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDFSServer).ReplicateFile(ctx, req.(*ReplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SDFS_ServiceDesc is the grpc.ServiceDesc for SDFS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SDFS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs425_mp3.SDFS",
	HandlerType: (*SDFSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFile",
			Handler:    _SDFS_GetFile_Handler,
		},
		{
			MethodName: "PutFile",
			Handler:    _SDFS_PutFile_Handler,
		},
		{
			MethodName: "DeleteFileLeader",
			Handler:    _SDFS_DeleteFileLeader_Handler,
		},
		{
			MethodName: "DeleteFileFollower",
			Handler:    _SDFS_DeleteFileFollower_Handler,
		},
		{
			MethodName: "LsFile",
			Handler:    _SDFS_LsFile_Handler,
		},
		{
			MethodName: "StoreFile",
			Handler:    _SDFS_StoreFile_Handler,
		},
		{
			MethodName: "ReplicateFile",
			Handler:    _SDFS_ReplicateFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mp3_SDFS.proto",
}
