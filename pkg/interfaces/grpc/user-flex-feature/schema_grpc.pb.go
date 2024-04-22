// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/schema.proto

package user_flex_feature

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
	UserFlexFeatureService_RuleUpdate_FullMethodName = "/user_flex_feature.v1.UserFlexFeatureService/RuleUpdate"
)

// UserFlexFeatureServiceClient is the client API for UserFlexFeatureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserFlexFeatureServiceClient interface {
	RuleUpdate(ctx context.Context, in *RuleUpdateRequest, opts ...grpc.CallOption) (*RuleUpdateResponse, error)
}

type userFlexFeatureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserFlexFeatureServiceClient(cc grpc.ClientConnInterface) UserFlexFeatureServiceClient {
	return &userFlexFeatureServiceClient{cc}
}

func (c *userFlexFeatureServiceClient) RuleUpdate(ctx context.Context, in *RuleUpdateRequest, opts ...grpc.CallOption) (*RuleUpdateResponse, error) {
	out := new(RuleUpdateResponse)
	err := c.cc.Invoke(ctx, UserFlexFeatureService_RuleUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserFlexFeatureServiceServer is the server API for UserFlexFeatureService service.
// All implementations should embed UnimplementedUserFlexFeatureServiceServer
// for forward compatibility
type UserFlexFeatureServiceServer interface {
	RuleUpdate(context.Context, *RuleUpdateRequest) (*RuleUpdateResponse, error)
}

// UnimplementedUserFlexFeatureServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserFlexFeatureServiceServer struct {
}

func (UnimplementedUserFlexFeatureServiceServer) RuleUpdate(context.Context, *RuleUpdateRequest) (*RuleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuleUpdate not implemented")
}

// UnsafeUserFlexFeatureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserFlexFeatureServiceServer will
// result in compilation errors.
type UnsafeUserFlexFeatureServiceServer interface {
	mustEmbedUnimplementedUserFlexFeatureServiceServer()
}

func RegisterUserFlexFeatureServiceServer(s grpc.ServiceRegistrar, srv UserFlexFeatureServiceServer) {
	s.RegisterService(&UserFlexFeatureService_ServiceDesc, srv)
}

func _UserFlexFeatureService_RuleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFlexFeatureServiceServer).RuleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserFlexFeatureService_RuleUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFlexFeatureServiceServer).RuleUpdate(ctx, req.(*RuleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserFlexFeatureService_ServiceDesc is the grpc.ServiceDesc for UserFlexFeatureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserFlexFeatureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_flex_feature.v1.UserFlexFeatureService",
	HandlerType: (*UserFlexFeatureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RuleUpdate",
			Handler:    _UserFlexFeatureService_RuleUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/schema.proto",
}
