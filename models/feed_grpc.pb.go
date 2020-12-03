// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package models

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RssClient is the client API for Rss service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RssClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Start(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	AddRss(ctx context.Context, in *RssLink, opts ...grpc.CallOption) (*empty.Empty, error)
	GetNews(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*News, error)
	ListNews(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*News, error)
}

type rssClient struct {
	cc grpc.ClientConnInterface
}

func NewRssClient(cc grpc.ClientConnInterface) RssClient {
	return &rssClient{cc}
}

func (c *rssClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/models.Rss/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rssClient) Start(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/models.Rss/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rssClient) Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/models.Rss/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rssClient) AddRss(ctx context.Context, in *RssLink, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/models.Rss/AddRss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rssClient) GetNews(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*News, error) {
	out := new(News)
	err := c.cc.Invoke(ctx, "/models.Rss/GetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rssClient) ListNews(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*News, error) {
	out := new(News)
	err := c.cc.Invoke(ctx, "/models.Rss/ListNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RssServer is the server API for Rss service.
// All implementations must embed UnimplementedRssServer
// for forward compatibility
type RssServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	Start(context.Context, *empty.Empty) (*empty.Empty, error)
	Stop(context.Context, *empty.Empty) (*empty.Empty, error)
	AddRss(context.Context, *RssLink) (*empty.Empty, error)
	GetNews(context.Context, *GetRequest) (*News, error)
	ListNews(context.Context, *empty.Empty) (*News, error)
	mustEmbedUnimplementedRssServer()
}

// UnimplementedRssServer must be embedded to have forward compatible implementations.
type UnimplementedRssServer struct {
}

func (UnimplementedRssServer) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedRssServer) Start(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedRssServer) Stop(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedRssServer) AddRss(context.Context, *RssLink) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRss not implemented")
}
func (UnimplementedRssServer) GetNews(context.Context, *GetRequest) (*News, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}
func (UnimplementedRssServer) ListNews(context.Context, *empty.Empty) (*News, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNews not implemented")
}
func (UnimplementedRssServer) mustEmbedUnimplementedRssServer() {}

// UnsafeRssServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RssServer will
// result in compilation errors.
type UnsafeRssServer interface {
	mustEmbedUnimplementedRssServer()
}

func RegisterRssServer(s grpc.ServiceRegistrar, srv RssServer) {
	s.RegisterService(&_Rss_serviceDesc, srv)
}

func _Rss_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RssServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Rss/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RssServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rss_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RssServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Rss/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RssServer).Start(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rss_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RssServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Rss/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RssServer).Stop(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rss_AddRss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RssLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RssServer).AddRss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Rss/AddRss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RssServer).AddRss(ctx, req.(*RssLink))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rss_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RssServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Rss/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RssServer).GetNews(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rss_ListNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RssServer).ListNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Rss/ListNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RssServer).ListNews(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rss_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Rss",
	HandlerType: (*RssServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Rss_Ping_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Rss_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Rss_Stop_Handler,
		},
		{
			MethodName: "AddRss",
			Handler:    _Rss_AddRss_Handler,
		},
		{
			MethodName: "GetNews",
			Handler:    _Rss_GetNews_Handler,
		},
		{
			MethodName: "ListNews",
			Handler:    _Rss_ListNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "models/feed.proto",
}
