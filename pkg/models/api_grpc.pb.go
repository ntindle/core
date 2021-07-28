// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package models

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

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	// Initializes New Node
	Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*NoResponse, error)
	// Starts Method Host and Connects to Network
	Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*NoResponse, error)
	// Action method handles misceallaneous actions for node
	Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	// Signing Method Request for Data
	Sign(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// Verification Method Request for Signed Data
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	// Update Method proximity/direction and Notify Lobby
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*NoResponse, error)
	// Invite Method Processes Data and Sends Invite to Peer
	Invite(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*NoResponse, error)
	// Respond Method to an Invite with Decision
	Respond(ctx context.Context, in *InviteResponse, opts ...grpc.CallOption) (*NoResponse, error)
	// Mail Method handles request for a message in Mailbox
	Mail(ctx context.Context, in *MailboxRequest, opts ...grpc.CallOption) (*MailboxResponse, error)
	//
	OnStatus(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnStatusClient, error)
	OnTopic(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnTopicClient, error)
	OnInvite(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnInviteClient, error)
	OnReply(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnReplyClient, error)
	OnMail(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnMailClient, error)
	OnProgress(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnProgressClient, error)
	OnComplete(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnCompleteClient, error)
	OnError(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnErrorClient, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Initialize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Sign(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Invite(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Invite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Respond(ctx context.Context, in *InviteResponse, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Respond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Mail(ctx context.Context, in *MailboxRequest, opts ...grpc.CallOption) (*MailboxResponse, error) {
	out := new(MailboxResponse)
	err := c.cc.Invoke(ctx, "/models.NodeService/Mail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) OnStatus(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[0], "/models.NodeService/OnStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnStatusClient interface {
	Recv() (*StatusEvent, error)
	grpc.ClientStream
}

type nodeServiceOnStatusClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnStatusClient) Recv() (*StatusEvent, error) {
	m := new(StatusEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnTopic(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnTopicClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[1], "/models.NodeService/OnTopic", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnTopicClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnTopicClient interface {
	Recv() (*TopicEvent, error)
	grpc.ClientStream
}

type nodeServiceOnTopicClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnTopicClient) Recv() (*TopicEvent, error) {
	m := new(TopicEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnInvite(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnInviteClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[2], "/models.NodeService/OnInvite", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnInviteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnInviteClient interface {
	Recv() (*InviteRequest, error)
	grpc.ClientStream
}

type nodeServiceOnInviteClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnInviteClient) Recv() (*InviteRequest, error) {
	m := new(InviteRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnReply(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnReplyClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[3], "/models.NodeService/OnReply", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnReplyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnReplyClient interface {
	Recv() (*InviteResponse, error)
	grpc.ClientStream
}

type nodeServiceOnReplyClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnReplyClient) Recv() (*InviteResponse, error) {
	m := new(InviteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnMail(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnMailClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[4], "/models.NodeService/OnMail", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnMailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnMailClient interface {
	Recv() (*MailEvent, error)
	grpc.ClientStream
}

type nodeServiceOnMailClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnMailClient) Recv() (*MailEvent, error) {
	m := new(MailEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnProgress(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[5], "/models.NodeService/OnProgress", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnProgressClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnProgressClient interface {
	Recv() (*ProgressEvent, error)
	grpc.ClientStream
}

type nodeServiceOnProgressClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnProgressClient) Recv() (*ProgressEvent, error) {
	m := new(ProgressEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnComplete(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnCompleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[6], "/models.NodeService/OnComplete", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnCompleteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnCompleteClient interface {
	Recv() (*CompleteEvent, error)
	grpc.ClientStream
}

type nodeServiceOnCompleteClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnCompleteClient) Recv() (*CompleteEvent, error) {
	m := new(CompleteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnError(ctx context.Context, in *NoRequest, opts ...grpc.CallOption) (NodeService_OnErrorClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[7], "/models.NodeService/OnError", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnErrorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnErrorClient interface {
	Recv() (*ErrorEvent, error)
	grpc.ClientStream
}

type nodeServiceOnErrorClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnErrorClient) Recv() (*ErrorEvent, error) {
	m := new(ErrorEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	// Initializes New Node
	Initialize(context.Context, *InitializeRequest) (*NoResponse, error)
	// Starts Method Host and Connects to Network
	Connect(context.Context, *ConnectionRequest) (*NoResponse, error)
	// Action method handles misceallaneous actions for node
	Action(context.Context, *ActionRequest) (*ActionResponse, error)
	// Signing Method Request for Data
	Sign(context.Context, *AuthRequest) (*AuthResponse, error)
	// Verification Method Request for Signed Data
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	// Update Method proximity/direction and Notify Lobby
	Update(context.Context, *UpdateRequest) (*NoResponse, error)
	// Invite Method Processes Data and Sends Invite to Peer
	Invite(context.Context, *InviteRequest) (*NoResponse, error)
	// Respond Method to an Invite with Decision
	Respond(context.Context, *InviteResponse) (*NoResponse, error)
	// Mail Method handles request for a message in Mailbox
	Mail(context.Context, *MailboxRequest) (*MailboxResponse, error)
	//
	OnStatus(*NoRequest, NodeService_OnStatusServer) error
	OnTopic(*NoRequest, NodeService_OnTopicServer) error
	OnInvite(*NoRequest, NodeService_OnInviteServer) error
	OnReply(*NoRequest, NodeService_OnReplyServer) error
	OnMail(*NoRequest, NodeService_OnMailServer) error
	OnProgress(*NoRequest, NodeService_OnProgressServer) error
	OnComplete(*NoRequest, NodeService_OnCompleteServer) error
	OnError(*NoRequest, NodeService_OnErrorServer) error
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) Initialize(context.Context, *InitializeRequest) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initialize not implemented")
}
func (UnimplementedNodeServiceServer) Connect(context.Context, *ConnectionRequest) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedNodeServiceServer) Action(context.Context, *ActionRequest) (*ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}
func (UnimplementedNodeServiceServer) Sign(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedNodeServiceServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedNodeServiceServer) Update(context.Context, *UpdateRequest) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNodeServiceServer) Invite(context.Context, *InviteRequest) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invite not implemented")
}
func (UnimplementedNodeServiceServer) Respond(context.Context, *InviteResponse) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Respond not implemented")
}
func (UnimplementedNodeServiceServer) Mail(context.Context, *MailboxRequest) (*MailboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mail not implemented")
}
func (UnimplementedNodeServiceServer) OnStatus(*NoRequest, NodeService_OnStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method OnStatus not implemented")
}
func (UnimplementedNodeServiceServer) OnTopic(*NoRequest, NodeService_OnTopicServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTopic not implemented")
}
func (UnimplementedNodeServiceServer) OnInvite(*NoRequest, NodeService_OnInviteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnInvite not implemented")
}
func (UnimplementedNodeServiceServer) OnReply(*NoRequest, NodeService_OnReplyServer) error {
	return status.Errorf(codes.Unimplemented, "method OnReply not implemented")
}
func (UnimplementedNodeServiceServer) OnMail(*NoRequest, NodeService_OnMailServer) error {
	return status.Errorf(codes.Unimplemented, "method OnMail not implemented")
}
func (UnimplementedNodeServiceServer) OnProgress(*NoRequest, NodeService_OnProgressServer) error {
	return status.Errorf(codes.Unimplemented, "method OnProgress not implemented")
}
func (UnimplementedNodeServiceServer) OnComplete(*NoRequest, NodeService_OnCompleteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnComplete not implemented")
}
func (UnimplementedNodeServiceServer) OnError(*NoRequest, NodeService_OnErrorServer) error {
	return status.Errorf(codes.Unimplemented, "method OnError not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_Initialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Initialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Initialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Initialize(ctx, req.(*InitializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Connect(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Action(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Sign(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Invite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Invite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Invite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Invite(ctx, req.(*InviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Respond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Respond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Respond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Respond(ctx, req.(*InviteResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Mail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Mail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.NodeService/Mail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Mail(ctx, req.(*MailboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_OnStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnStatus(m, &nodeServiceOnStatusServer{stream})
}

type NodeService_OnStatusServer interface {
	Send(*StatusEvent) error
	grpc.ServerStream
}

type nodeServiceOnStatusServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnStatusServer) Send(m *StatusEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnTopic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnTopic(m, &nodeServiceOnTopicServer{stream})
}

type NodeService_OnTopicServer interface {
	Send(*TopicEvent) error
	grpc.ServerStream
}

type nodeServiceOnTopicServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnTopicServer) Send(m *TopicEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnInvite_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnInvite(m, &nodeServiceOnInviteServer{stream})
}

type NodeService_OnInviteServer interface {
	Send(*InviteRequest) error
	grpc.ServerStream
}

type nodeServiceOnInviteServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnInviteServer) Send(m *InviteRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnReply_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnReply(m, &nodeServiceOnReplyServer{stream})
}

type NodeService_OnReplyServer interface {
	Send(*InviteResponse) error
	grpc.ServerStream
}

type nodeServiceOnReplyServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnReplyServer) Send(m *InviteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnMail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnMail(m, &nodeServiceOnMailServer{stream})
}

type NodeService_OnMailServer interface {
	Send(*MailEvent) error
	grpc.ServerStream
}

type nodeServiceOnMailServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnMailServer) Send(m *MailEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnProgress(m, &nodeServiceOnProgressServer{stream})
}

type NodeService_OnProgressServer interface {
	Send(*ProgressEvent) error
	grpc.ServerStream
}

type nodeServiceOnProgressServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnProgressServer) Send(m *ProgressEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnComplete_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnComplete(m, &nodeServiceOnCompleteServer{stream})
}

type NodeService_OnCompleteServer interface {
	Send(*CompleteEvent) error
	grpc.ServerStream
}

type nodeServiceOnCompleteServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnCompleteServer) Send(m *CompleteEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnError_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnError(m, &nodeServiceOnErrorServer{stream})
}

type NodeService_OnErrorServer interface {
	Send(*ErrorEvent) error
	grpc.ServerStream
}

type nodeServiceOnErrorServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnErrorServer) Send(m *ErrorEvent) error {
	return x.ServerStream.SendMsg(m)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Initialize",
			Handler:    _NodeService_Initialize_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _NodeService_Connect_Handler,
		},
		{
			MethodName: "Action",
			Handler:    _NodeService_Action_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _NodeService_Sign_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _NodeService_Verify_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NodeService_Update_Handler,
		},
		{
			MethodName: "Invite",
			Handler:    _NodeService_Invite_Handler,
		},
		{
			MethodName: "Respond",
			Handler:    _NodeService_Respond_Handler,
		},
		{
			MethodName: "Mail",
			Handler:    _NodeService_Mail_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnStatus",
			Handler:       _NodeService_OnStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTopic",
			Handler:       _NodeService_OnTopic_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnInvite",
			Handler:       _NodeService_OnInvite_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnReply",
			Handler:       _NodeService_OnReply_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnMail",
			Handler:       _NodeService_OnMail_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnProgress",
			Handler:       _NodeService_OnProgress_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnComplete",
			Handler:       _NodeService_OnComplete_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnError",
			Handler:       _NodeService_OnError_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}