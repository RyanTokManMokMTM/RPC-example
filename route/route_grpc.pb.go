// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package route

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

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {
	//unary => send a request and response a request
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	//streaming
	//server,keep streaming Feature from request points
	GetListOfFeature(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_GetListOfFeatureClient, error)
	//client,keep sending point and return a summary
	GetRecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_GetRecordRouteClient, error)
	//2 side streaming
	Recommend(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecommendClient, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/route.RouteGuide/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) GetListOfFeature(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_GetListOfFeatureClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[0], "/route.RouteGuide/GetListOfFeature", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideGetListOfFeatureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_GetListOfFeatureClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideGetListOfFeatureClient struct {
	grpc.ClientStream
}

func (x *routeGuideGetListOfFeatureClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) GetRecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_GetRecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[1], "/route.RouteGuide/GetRecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideGetRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_GetRecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideGetRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideGetRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideGetRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) Recommend(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecommendClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[2], "/route.RouteGuide/Recommend", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecommendClient{stream}
	return x, nil
}

type RouteGuide_RecommendClient interface {
	Send(*RecommendationRequest) error
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideRecommendClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecommendClient) Send(m *RecommendationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecommendClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
// All implementations must embed UnimplementedRouteGuideServer
// for forward compatibility
type RouteGuideServer interface {
	//unary => send a request and response a request
	GetFeature(context.Context, *Point) (*Feature, error)
	//streaming
	//server,keep streaming Feature from request points
	GetListOfFeature(*Rectangle, RouteGuide_GetListOfFeatureServer) error
	//client,keep sending point and return a summary
	GetRecordRoute(RouteGuide_GetRecordRouteServer) error
	//2 side streaming
	Recommend(RouteGuide_RecommendServer) error
	mustEmbedUnimplementedRouteGuideServer()
}

// UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (UnimplementedRouteGuideServer) GetFeature(context.Context, *Point) (*Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedRouteGuideServer) GetListOfFeature(*Rectangle, RouteGuide_GetListOfFeatureServer) error {
	return status.Errorf(codes.Unimplemented, "method GetListOfFeature not implemented")
}
func (UnimplementedRouteGuideServer) GetRecordRoute(RouteGuide_GetRecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRecordRoute not implemented")
}
func (UnimplementedRouteGuideServer) Recommend(RouteGuide_RecommendServer) error {
	return status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}
func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {}

// UnsafeRouteGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServer will
// result in compilation errors.
type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	s.RegisterService(&RouteGuide_ServiceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_GetListOfFeature_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).GetListOfFeature(m, &routeGuideGetListOfFeatureServer{stream})
}

type RouteGuide_GetListOfFeatureServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideGetListOfFeatureServer struct {
	grpc.ServerStream
}

func (x *routeGuideGetListOfFeatureServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_GetRecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).GetRecordRoute(&routeGuideGetRecordRouteServer{stream})
}

type RouteGuide_GetRecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideGetRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideGetRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideGetRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_Recommend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).Recommend(&routeGuideRecommendServer{stream})
}

type RouteGuide_RecommendServer interface {
	Send(*Feature) error
	Recv() (*RecommendationRequest, error)
	grpc.ServerStream
}

type routeGuideRecommendServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecommendServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecommendServer) Recv() (*RecommendationRequest, error) {
	m := new(RecommendationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuide_ServiceDesc is the grpc.ServiceDesc for RouteGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetListOfFeature",
			Handler:       _RouteGuide_GetListOfFeature_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRecordRoute",
			Handler:       _RouteGuide_GetRecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Recommend",
			Handler:       _RouteGuide_Recommend_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "route.proto",
}