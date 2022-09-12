// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/api/component/camera/v1/camera.proto

package v1

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CameraServiceClient is the client API for CameraService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CameraServiceClient interface {
	// GetImage returns a frame from a camera of the underlying robot. A specific MIME type
	// can be requested but may not necessarily be the same one returned.
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageResponse, error)
	// RenderFrame renders a frame from a camera of the underlying robot to an HTTP response. A specific MIME type
	// can be requested but may not necessarily be the same one returned.
	RenderFrame(ctx context.Context, in *RenderFrameRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// GetPointCloud returns a point cloud from a camera of the underlying robot. A specific MIME type
	// can be requested but may not necessarily be the same one returned.
	GetPointCloud(ctx context.Context, in *GetPointCloudRequest, opts ...grpc.CallOption) (*GetPointCloudResponse, error)
	// GetProperties returns the camera intrinsic parameters and camera distortion parameters from a camera of the underlying robot, if available.
	GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...grpc.CallOption) (*GetPropertiesResponse, error)
}

type cameraServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCameraServiceClient(cc grpc.ClientConnInterface) CameraServiceClient {
	return &cameraServiceClient{cc}
}

func (c *cameraServiceClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageResponse, error) {
	out := new(GetImageResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.camera.v1.CameraService/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameraServiceClient) RenderFrame(ctx context.Context, in *RenderFrameRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/proto.api.component.camera.v1.CameraService/RenderFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameraServiceClient) GetPointCloud(ctx context.Context, in *GetPointCloudRequest, opts ...grpc.CallOption) (*GetPointCloudResponse, error) {
	out := new(GetPointCloudResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.camera.v1.CameraService/GetPointCloud", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cameraServiceClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...grpc.CallOption) (*GetPropertiesResponse, error) {
	out := new(GetPropertiesResponse)
	err := c.cc.Invoke(ctx, "/proto.api.component.camera.v1.CameraService/GetProperties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CameraServiceServer is the server API for CameraService service.
// All implementations must embed UnimplementedCameraServiceServer
// for forward compatibility
type CameraServiceServer interface {
	// GetImage returns a frame from a camera of the underlying robot. A specific MIME type
	// can be requested but may not necessarily be the same one returned.
	GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error)
	// RenderFrame renders a frame from a camera of the underlying robot to an HTTP response. A specific MIME type
	// can be requested but may not necessarily be the same one returned.
	RenderFrame(context.Context, *RenderFrameRequest) (*httpbody.HttpBody, error)
	// GetPointCloud returns a point cloud from a camera of the underlying robot. A specific MIME type
	// can be requested but may not necessarily be the same one returned.
	GetPointCloud(context.Context, *GetPointCloudRequest) (*GetPointCloudResponse, error)
	// GetProperties returns the camera intrinsic parameters and camera distortion parameters from a camera of the underlying robot, if available.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
	mustEmbedUnimplementedCameraServiceServer()
}

// UnimplementedCameraServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCameraServiceServer struct {
}

func (UnimplementedCameraServiceServer) GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedCameraServiceServer) RenderFrame(context.Context, *RenderFrameRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderFrame not implemented")
}
func (UnimplementedCameraServiceServer) GetPointCloud(context.Context, *GetPointCloudRequest) (*GetPointCloudResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPointCloud not implemented")
}
func (UnimplementedCameraServiceServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProperties not implemented")
}
func (UnimplementedCameraServiceServer) mustEmbedUnimplementedCameraServiceServer() {}

// UnsafeCameraServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CameraServiceServer will
// result in compilation errors.
type UnsafeCameraServiceServer interface {
	mustEmbedUnimplementedCameraServiceServer()
}

func RegisterCameraServiceServer(s grpc.ServiceRegistrar, srv CameraServiceServer) {
	s.RegisterService(&CameraService_ServiceDesc, srv)
}

func _CameraService_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameraServiceServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.camera.v1.CameraService/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameraServiceServer).GetImage(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameraService_RenderFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameraServiceServer).RenderFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.camera.v1.CameraService/RenderFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameraServiceServer).RenderFrame(ctx, req.(*RenderFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameraService_GetPointCloud_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPointCloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameraServiceServer).GetPointCloud(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.camera.v1.CameraService/GetPointCloud",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameraServiceServer).GetPointCloud(ctx, req.(*GetPointCloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CameraService_GetProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPropertiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameraServiceServer).GetProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.component.camera.v1.CameraService/GetProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameraServiceServer).GetProperties(ctx, req.(*GetPropertiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CameraService_ServiceDesc is the grpc.ServiceDesc for CameraService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CameraService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api.component.camera.v1.CameraService",
	HandlerType: (*CameraServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImage",
			Handler:    _CameraService_GetImage_Handler,
		},
		{
			MethodName: "RenderFrame",
			Handler:    _CameraService_RenderFrame_Handler,
		},
		{
			MethodName: "GetPointCloud",
			Handler:    _CameraService_GetPointCloud_Handler,
		},
		{
			MethodName: "GetProperties",
			Handler:    _CameraService_GetProperties_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/component/camera/v1/camera.proto",
}
