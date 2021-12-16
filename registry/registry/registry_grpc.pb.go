// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package registry

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

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	// GetBrokers returns a BrokerResponse with the brokers field populated
	// with full broker metadata. If the input BrokerRequest.id field is
	// non-nil, a single broker is returned matching the ID specified in the
	// Broker object. Otherwise all brokers are returned, optionally filtered
	// by any provided BrokerRequest.tags parameters.
	GetBrokers(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*BrokerResponse, error)
	// ListBrokers returns a BrokerResponse with the ids field populated
	// with broker IDs. If the input BrokerRequest.id field is non-nil,
	// a single broker ID is returned matching the ID specified in the
	// Broker object if the broker exists. Otherwise all brokers are returned,
	// optionally filtered by any provided BrokerRequest.tags parameters.
	ListBrokers(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*BrokerResponse, error)
	// UnmappedBrokers returns a BrokerResponse with the ids field
	// populated with broker IDs that do not hold any assigned partitions.
	// Any topic names specified in the UnmappedBrokersRequest exclude field
	// are ignored. For example, broker 1000 holds no partitions other
	// than one belonging to the 'test0' topic. If UnmappedBrokers is called
	// with 'test0' specified as an exclude name, broker 1000 will be returned
	// in the BrokerResponse as an unmapped broker.
	UnmappedBrokers(ctx context.Context, in *UnmappedBrokersRequest, opts ...grpc.CallOption) (*BrokerResponse, error)
	// GetTopics returns a TopicResponse with the topics field populated
	// with full topic metadata. If the input TopicRequest.name field is
	// non-nil, a single topic is returned matching the name specified in the
	// Topic object. Otherwise all topics are returned, optionally filtered
	// by any provided TopicRequest.tags parameters.
	GetTopics(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	// ListTopics returns a TopicResponse with the names field populated
	// with topic names. If the input TopicRequest.name field is non-nil,
	// a single topic name is returned matching the name specified in the
	// Topic object if the topic exists. Otherwise all topics are returned,
	// optionally filtered by any provided TopicRequest.tags parameters.
	ListTopics(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	//
	//CreateTopic creates a topic.
	//Example:
	//$ curl -XPOST "localhost:8080/v1/topics/create" -d '{
	//"topic": {
	//"name": "mytopic",
	//"partitions": 32,
	//"replication": 2,
	//"tags": {"env":"staging"}
	//},
	//"target_broker_tags": ["pool:tests"]
	//}'
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*Empty, error)
	//
	//DeleteTopic takes a TopicRequest and deletes the topic specified in the
	//TopicRequest.name field.
	//Example:
	//$ curl -XDELETE "localhost:8080/v1/topics/mytopic"
	DeleteTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*Empty, error)
	// ReassigningTopics returns a TopicResponse with the names field populated
	// with topic names of all topics undergoing a reassignment.
	ReassigningTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicResponse, error)
	// UnderReplicatedTopics returns a TopicResponse with the names field populated
	// with topic names of all under replicated topics.
	UnderReplicatedTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicResponse, error)
	// TopicMappings returns a BrokerResponse with the ids field
	// populated with broker IDs that hold at least one partition
	// for the requested topic. Both a single topic name or specified in the
	// TopicRequest.name field.
	TopicMappings(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*BrokerResponse, error)
	// BrokerMappings returns a TopicResponse with the names field
	// populated with topics that the broker holds at least one partition
	// for the requested broker. The broker is specified in the
	// BrokerRequest.id field.
	BrokerMappings(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	// TagTopic takes a TopicRequest and sets any specified
	// tags for the named topic. Any existing tags that are
	// not specified in the request are left unmodified.
	TagTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TagResponse, error)
	// DeleteTopicTags takes a TopicRequest and deletes any
	// specified tags for the named topic. Tags must be provided
	// as key names only; "key:value" will not target the tag "key".
	DeleteTopicTags(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TagResponse, error)
	// TagBroker takes a BrokerRequest and sets any specified
	// tags for the named broker. Any existing tags that are
	// not specified in the request are left unmodified.
	TagBroker(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*TagResponse, error)
	// DeleteBrokerTags takes a BrokerRequest and deletes any
	// specified tags for the named broker. Tags must be provided
	// as key names only; "key:value" will not target the tag "key".
	DeleteBrokerTags(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*TagResponse, error)
	// TranslateOffsets returns a TranslateOffsetResponse with the
	// the upstream/local offsets for the provided consumer group
	// populated per topic/partition.
	// The remote cluster alias and consumer group id are specified
	// in the TranslateOffsetRequest.remote_cluster_alias and
	// TranslateOffsetRequest.group_id respectively.
	TranslateOffsets(ctx context.Context, in *TranslateOffsetRequest, opts ...grpc.CallOption) (*TranslateOffsetResponse, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) GetBrokers(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*BrokerResponse, error) {
	out := new(BrokerResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetBrokers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListBrokers(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*BrokerResponse, error) {
	out := new(BrokerResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListBrokers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) UnmappedBrokers(ctx context.Context, in *UnmappedBrokersRequest, opts ...grpc.CallOption) (*BrokerResponse, error) {
	out := new(BrokerResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/UnmappedBrokers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetTopics(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListTopics(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/registry.Registry/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeleteTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/registry.Registry/DeleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ReassigningTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/ReassigningTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) UnderReplicatedTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/UnderReplicatedTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) TopicMappings(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*BrokerResponse, error) {
	out := new(BrokerResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/TopicMappings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) BrokerMappings(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/BrokerMappings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) TagTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TagResponse, error) {
	out := new(TagResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/TagTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeleteTopicTags(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TagResponse, error) {
	out := new(TagResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/DeleteTopicTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) TagBroker(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*TagResponse, error) {
	out := new(TagResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/TagBroker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeleteBrokerTags(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*TagResponse, error) {
	out := new(TagResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/DeleteBrokerTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) TranslateOffsets(ctx context.Context, in *TranslateOffsetRequest, opts ...grpc.CallOption) (*TranslateOffsetResponse, error) {
	out := new(TranslateOffsetResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/TranslateOffsets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	// GetBrokers returns a BrokerResponse with the brokers field populated
	// with full broker metadata. If the input BrokerRequest.id field is
	// non-nil, a single broker is returned matching the ID specified in the
	// Broker object. Otherwise all brokers are returned, optionally filtered
	// by any provided BrokerRequest.tags parameters.
	GetBrokers(context.Context, *BrokerRequest) (*BrokerResponse, error)
	// ListBrokers returns a BrokerResponse with the ids field populated
	// with broker IDs. If the input BrokerRequest.id field is non-nil,
	// a single broker ID is returned matching the ID specified in the
	// Broker object if the broker exists. Otherwise all brokers are returned,
	// optionally filtered by any provided BrokerRequest.tags parameters.
	ListBrokers(context.Context, *BrokerRequest) (*BrokerResponse, error)
	// UnmappedBrokers returns a BrokerResponse with the ids field
	// populated with broker IDs that do not hold any assigned partitions.
	// Any topic names specified in the UnmappedBrokersRequest exclude field
	// are ignored. For example, broker 1000 holds no partitions other
	// than one belonging to the 'test0' topic. If UnmappedBrokers is called
	// with 'test0' specified as an exclude name, broker 1000 will be returned
	// in the BrokerResponse as an unmapped broker.
	UnmappedBrokers(context.Context, *UnmappedBrokersRequest) (*BrokerResponse, error)
	// GetTopics returns a TopicResponse with the topics field populated
	// with full topic metadata. If the input TopicRequest.name field is
	// non-nil, a single topic is returned matching the name specified in the
	// Topic object. Otherwise all topics are returned, optionally filtered
	// by any provided TopicRequest.tags parameters.
	GetTopics(context.Context, *TopicRequest) (*TopicResponse, error)
	// ListTopics returns a TopicResponse with the names field populated
	// with topic names. If the input TopicRequest.name field is non-nil,
	// a single topic name is returned matching the name specified in the
	// Topic object if the topic exists. Otherwise all topics are returned,
	// optionally filtered by any provided TopicRequest.tags parameters.
	ListTopics(context.Context, *TopicRequest) (*TopicResponse, error)
	//
	//CreateTopic creates a topic.
	//Example:
	//$ curl -XPOST "localhost:8080/v1/topics/create" -d '{
	//"topic": {
	//"name": "mytopic",
	//"partitions": 32,
	//"replication": 2,
	//"tags": {"env":"staging"}
	//},
	//"target_broker_tags": ["pool:tests"]
	//}'
	CreateTopic(context.Context, *CreateTopicRequest) (*Empty, error)
	//
	//DeleteTopic takes a TopicRequest and deletes the topic specified in the
	//TopicRequest.name field.
	//Example:
	//$ curl -XDELETE "localhost:8080/v1/topics/mytopic"
	DeleteTopic(context.Context, *TopicRequest) (*Empty, error)
	// ReassigningTopics returns a TopicResponse with the names field populated
	// with topic names of all topics undergoing a reassignment.
	ReassigningTopics(context.Context, *Empty) (*TopicResponse, error)
	// UnderReplicatedTopics returns a TopicResponse with the names field populated
	// with topic names of all under replicated topics.
	UnderReplicatedTopics(context.Context, *Empty) (*TopicResponse, error)
	// TopicMappings returns a BrokerResponse with the ids field
	// populated with broker IDs that hold at least one partition
	// for the requested topic. Both a single topic name or specified in the
	// TopicRequest.name field.
	TopicMappings(context.Context, *TopicRequest) (*BrokerResponse, error)
	// BrokerMappings returns a TopicResponse with the names field
	// populated with topics that the broker holds at least one partition
	// for the requested broker. The broker is specified in the
	// BrokerRequest.id field.
	BrokerMappings(context.Context, *BrokerRequest) (*TopicResponse, error)
	// TagTopic takes a TopicRequest and sets any specified
	// tags for the named topic. Any existing tags that are
	// not specified in the request are left unmodified.
	TagTopic(context.Context, *TopicRequest) (*TagResponse, error)
	// DeleteTopicTags takes a TopicRequest and deletes any
	// specified tags for the named topic. Tags must be provided
	// as key names only; "key:value" will not target the tag "key".
	DeleteTopicTags(context.Context, *TopicRequest) (*TagResponse, error)
	// TagBroker takes a BrokerRequest and sets any specified
	// tags for the named broker. Any existing tags that are
	// not specified in the request are left unmodified.
	TagBroker(context.Context, *BrokerRequest) (*TagResponse, error)
	// DeleteBrokerTags takes a BrokerRequest and deletes any
	// specified tags for the named broker. Tags must be provided
	// as key names only; "key:value" will not target the tag "key".
	DeleteBrokerTags(context.Context, *BrokerRequest) (*TagResponse, error)
	// TranslateOffsets returns a TranslateOffsetResponse with the
	// the upstream/local offsets for the provided consumer group
	// populated per topic/partition.
	// The remote cluster alias and consumer group id are specified
	// in the TranslateOffsetRequest.remote_cluster_alias and
	// TranslateOffsetRequest.group_id respectively.
	TranslateOffsets(context.Context, *TranslateOffsetRequest) (*TranslateOffsetResponse, error)
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) GetBrokers(context.Context, *BrokerRequest) (*BrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrokers not implemented")
}
func (UnimplementedRegistryServer) ListBrokers(context.Context, *BrokerRequest) (*BrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBrokers not implemented")
}
func (UnimplementedRegistryServer) UnmappedBrokers(context.Context, *UnmappedBrokersRequest) (*BrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnmappedBrokers not implemented")
}
func (UnimplementedRegistryServer) GetTopics(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopics not implemented")
}
func (UnimplementedRegistryServer) ListTopics(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (UnimplementedRegistryServer) CreateTopic(context.Context, *CreateTopicRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedRegistryServer) DeleteTopic(context.Context, *TopicRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedRegistryServer) ReassigningTopics(context.Context, *Empty) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReassigningTopics not implemented")
}
func (UnimplementedRegistryServer) UnderReplicatedTopics(context.Context, *Empty) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnderReplicatedTopics not implemented")
}
func (UnimplementedRegistryServer) TopicMappings(context.Context, *TopicRequest) (*BrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicMappings not implemented")
}
func (UnimplementedRegistryServer) BrokerMappings(context.Context, *BrokerRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BrokerMappings not implemented")
}
func (UnimplementedRegistryServer) TagTopic(context.Context, *TopicRequest) (*TagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagTopic not implemented")
}
func (UnimplementedRegistryServer) DeleteTopicTags(context.Context, *TopicRequest) (*TagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopicTags not implemented")
}
func (UnimplementedRegistryServer) TagBroker(context.Context, *BrokerRequest) (*TagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagBroker not implemented")
}
func (UnimplementedRegistryServer) DeleteBrokerTags(context.Context, *BrokerRequest) (*TagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrokerTags not implemented")
}
func (UnimplementedRegistryServer) TranslateOffsets(context.Context, *TranslateOffsetRequest) (*TranslateOffsetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateOffsets not implemented")
}
func (UnimplementedRegistryServer) mustEmbedUnimplementedRegistryServer() {}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_GetBrokers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetBrokers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetBrokers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetBrokers(ctx, req.(*BrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListBrokers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListBrokers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListBrokers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListBrokers(ctx, req.(*BrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_UnmappedBrokers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnmappedBrokersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).UnmappedBrokers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/UnmappedBrokers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).UnmappedBrokers(ctx, req.(*UnmappedBrokersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetTopics(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListTopics(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeleteTopic(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ReassigningTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ReassigningTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ReassigningTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ReassigningTopics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_UnderReplicatedTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).UnderReplicatedTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/UnderReplicatedTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).UnderReplicatedTopics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_TopicMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).TopicMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/TopicMappings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).TopicMappings(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_BrokerMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).BrokerMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/BrokerMappings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).BrokerMappings(ctx, req.(*BrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_TagTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).TagTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/TagTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).TagTopic(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeleteTopicTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeleteTopicTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/DeleteTopicTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeleteTopicTags(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_TagBroker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).TagBroker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/TagBroker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).TagBroker(ctx, req.(*BrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeleteBrokerTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeleteBrokerTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/DeleteBrokerTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeleteBrokerTags(ctx, req.(*BrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_TranslateOffsets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateOffsetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).TranslateOffsets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/TranslateOffsets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).TranslateOffsets(ctx, req.(*TranslateOffsetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBrokers",
			Handler:    _Registry_GetBrokers_Handler,
		},
		{
			MethodName: "ListBrokers",
			Handler:    _Registry_ListBrokers_Handler,
		},
		{
			MethodName: "UnmappedBrokers",
			Handler:    _Registry_UnmappedBrokers_Handler,
		},
		{
			MethodName: "GetTopics",
			Handler:    _Registry_GetTopics_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _Registry_ListTopics_Handler,
		},
		{
			MethodName: "CreateTopic",
			Handler:    _Registry_CreateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _Registry_DeleteTopic_Handler,
		},
		{
			MethodName: "ReassigningTopics",
			Handler:    _Registry_ReassigningTopics_Handler,
		},
		{
			MethodName: "UnderReplicatedTopics",
			Handler:    _Registry_UnderReplicatedTopics_Handler,
		},
		{
			MethodName: "TopicMappings",
			Handler:    _Registry_TopicMappings_Handler,
		},
		{
			MethodName: "BrokerMappings",
			Handler:    _Registry_BrokerMappings_Handler,
		},
		{
			MethodName: "TagTopic",
			Handler:    _Registry_TagTopic_Handler,
		},
		{
			MethodName: "DeleteTopicTags",
			Handler:    _Registry_DeleteTopicTags_Handler,
		},
		{
			MethodName: "TagBroker",
			Handler:    _Registry_TagBroker_Handler,
		},
		{
			MethodName: "DeleteBrokerTags",
			Handler:    _Registry_DeleteBrokerTags_Handler,
		},
		{
			MethodName: "TranslateOffsets",
			Handler:    _Registry_TranslateOffsets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/registry.proto",
}
