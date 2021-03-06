// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loadtest.proto

package genproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StartRequest struct {
	// The GCP project. This must be set even for Kafka, as we use it to export metrics.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The Pub/Sub or Kafka topic name.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// The time at which the load test should start. If this is less than the current time, we start immediately.
	StartTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The duration the load test should run for.
	TestDuration *duration.Duration `protobuf:"bytes,4,opt,name=test_duration,json=testDuration,proto3" json:"test_duration,omitempty"`
	// Whether to include ids in check responses.
	IncludeIds bool `protobuf:"varint,5,opt,name=include_ids,json=includeIds,proto3" json:"include_ids,omitempty"`
	// Types that are valid to be assigned to Options:
	//	*StartRequest_PubsubOptions
	//	*StartRequest_KafkaOptions
	Options isStartRequest_Options `protobuf_oneof:"options"`
	// Types that are valid to be assigned to ClientOptions:
	//	*StartRequest_PublisherOptions
	//	*StartRequest_SubscriberOptions
	ClientOptions isStartRequest_ClientOptions `protobuf_oneof:"client_options"`
	// The cpu scaling of the worker.  A multiple of the number of logical processors
	// on the machine.  The number of threads for the worker is calculated by
	// max((numCpus * cpu_scaling), 1) for languages which use thread parallelism.
	// Languages which use process parallelism ignore this setting.
	CpuScaling           int32    `protobuf:"varint,10,opt,name=cpu_scaling,json=cpuScaling,proto3" json:"cpu_scaling,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{0}
}
func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (dst *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(dst, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *StartRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *StartRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *StartRequest) GetTestDuration() *duration.Duration {
	if m != nil {
		return m.TestDuration
	}
	return nil
}

func (m *StartRequest) GetIncludeIds() bool {
	if m != nil {
		return m.IncludeIds
	}
	return false
}

type isStartRequest_Options interface {
	isStartRequest_Options()
}

type StartRequest_PubsubOptions struct {
	PubsubOptions *PubsubOptions `protobuf:"bytes,6,opt,name=pubsub_options,json=pubsubOptions,proto3,oneof"`
}

type StartRequest_KafkaOptions struct {
	KafkaOptions *KafkaOptions `protobuf:"bytes,7,opt,name=kafka_options,json=kafkaOptions,proto3,oneof"`
}

func (*StartRequest_PubsubOptions) isStartRequest_Options() {}

func (*StartRequest_KafkaOptions) isStartRequest_Options() {}

func (m *StartRequest) GetOptions() isStartRequest_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *StartRequest) GetPubsubOptions() *PubsubOptions {
	if x, ok := m.GetOptions().(*StartRequest_PubsubOptions); ok {
		return x.PubsubOptions
	}
	return nil
}

func (m *StartRequest) GetKafkaOptions() *KafkaOptions {
	if x, ok := m.GetOptions().(*StartRequest_KafkaOptions); ok {
		return x.KafkaOptions
	}
	return nil
}

type isStartRequest_ClientOptions interface {
	isStartRequest_ClientOptions()
}

type StartRequest_PublisherOptions struct {
	PublisherOptions *PublisherOptions `protobuf:"bytes,8,opt,name=publisher_options,json=publisherOptions,proto3,oneof"`
}

type StartRequest_SubscriberOptions struct {
	SubscriberOptions *SubscriberOptions `protobuf:"bytes,9,opt,name=subscriber_options,json=subscriberOptions,proto3,oneof"`
}

func (*StartRequest_PublisherOptions) isStartRequest_ClientOptions() {}

func (*StartRequest_SubscriberOptions) isStartRequest_ClientOptions() {}

func (m *StartRequest) GetClientOptions() isStartRequest_ClientOptions {
	if m != nil {
		return m.ClientOptions
	}
	return nil
}

func (m *StartRequest) GetPublisherOptions() *PublisherOptions {
	if x, ok := m.GetClientOptions().(*StartRequest_PublisherOptions); ok {
		return x.PublisherOptions
	}
	return nil
}

func (m *StartRequest) GetSubscriberOptions() *SubscriberOptions {
	if x, ok := m.GetClientOptions().(*StartRequest_SubscriberOptions); ok {
		return x.SubscriberOptions
	}
	return nil
}

func (m *StartRequest) GetCpuScaling() int32 {
	if m != nil {
		return m.CpuScaling
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StartRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StartRequest_OneofMarshaler, _StartRequest_OneofUnmarshaler, _StartRequest_OneofSizer, []interface{}{
		(*StartRequest_PubsubOptions)(nil),
		(*StartRequest_KafkaOptions)(nil),
		(*StartRequest_PublisherOptions)(nil),
		(*StartRequest_SubscriberOptions)(nil),
	}
}

func _StartRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StartRequest)
	// options
	switch x := m.Options.(type) {
	case *StartRequest_PubsubOptions:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PubsubOptions); err != nil {
			return err
		}
	case *StartRequest_KafkaOptions:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KafkaOptions); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StartRequest.Options has unexpected type %T", x)
	}
	// client_options
	switch x := m.ClientOptions.(type) {
	case *StartRequest_PublisherOptions:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PublisherOptions); err != nil {
			return err
		}
	case *StartRequest_SubscriberOptions:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SubscriberOptions); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StartRequest.ClientOptions has unexpected type %T", x)
	}
	return nil
}

func _StartRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StartRequest)
	switch tag {
	case 6: // options.pubsub_options
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PubsubOptions)
		err := b.DecodeMessage(msg)
		m.Options = &StartRequest_PubsubOptions{msg}
		return true, err
	case 7: // options.kafka_options
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaOptions)
		err := b.DecodeMessage(msg)
		m.Options = &StartRequest_KafkaOptions{msg}
		return true, err
	case 8: // client_options.publisher_options
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PublisherOptions)
		err := b.DecodeMessage(msg)
		m.ClientOptions = &StartRequest_PublisherOptions{msg}
		return true, err
	case 9: // client_options.subscriber_options
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SubscriberOptions)
		err := b.DecodeMessage(msg)
		m.ClientOptions = &StartRequest_SubscriberOptions{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StartRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StartRequest)
	// options
	switch x := m.Options.(type) {
	case *StartRequest_PubsubOptions:
		s := proto.Size(x.PubsubOptions)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StartRequest_KafkaOptions:
		s := proto.Size(x.KafkaOptions)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// client_options
	switch x := m.ClientOptions.(type) {
	case *StartRequest_PublisherOptions:
		s := proto.Size(x.PublisherOptions)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StartRequest_SubscriberOptions:
		s := proto.Size(x.SubscriberOptions)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{1}
}
func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (dst *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(dst, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

type PublisherOptions struct {
	// The max messages-per-second publishing rate.  If unset, no rate limit will
	// be imposed.
	Rate float32 `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	// The max duration for coalescing a batch of published messages.
	BatchDuration *duration.Duration `protobuf:"bytes,2,opt,name=batch_duration,json=batchDuration,proto3" json:"batch_duration,omitempty"`
	// The number of user messages of size message_size to publish together.
	BatchSize int32 `protobuf:"varint,3,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	// The size in bytes of messages to publish
	MessageSize          int32    `protobuf:"varint,4,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublisherOptions) Reset()         { *m = PublisherOptions{} }
func (m *PublisherOptions) String() string { return proto.CompactTextString(m) }
func (*PublisherOptions) ProtoMessage()    {}
func (*PublisherOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{2}
}
func (m *PublisherOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublisherOptions.Unmarshal(m, b)
}
func (m *PublisherOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublisherOptions.Marshal(b, m, deterministic)
}
func (dst *PublisherOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublisherOptions.Merge(dst, src)
}
func (m *PublisherOptions) XXX_Size() int {
	return xxx_messageInfo_PublisherOptions.Size(m)
}
func (m *PublisherOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_PublisherOptions.DiscardUnknown(m)
}

var xxx_messageInfo_PublisherOptions proto.InternalMessageInfo

func (m *PublisherOptions) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *PublisherOptions) GetBatchDuration() *duration.Duration {
	if m != nil {
		return m.BatchDuration
	}
	return nil
}

func (m *PublisherOptions) GetBatchSize() int32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *PublisherOptions) GetMessageSize() int32 {
	if m != nil {
		return m.MessageSize
	}
	return 0
}

type SubscriberOptions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriberOptions) Reset()         { *m = SubscriberOptions{} }
func (m *SubscriberOptions) String() string { return proto.CompactTextString(m) }
func (*SubscriberOptions) ProtoMessage()    {}
func (*SubscriberOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{3}
}
func (m *SubscriberOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberOptions.Unmarshal(m, b)
}
func (m *SubscriberOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberOptions.Marshal(b, m, deterministic)
}
func (dst *SubscriberOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberOptions.Merge(dst, src)
}
func (m *SubscriberOptions) XXX_Size() int {
	return xxx_messageInfo_SubscriberOptions.Size(m)
}
func (m *SubscriberOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberOptions proto.InternalMessageInfo

type PubsubOptions struct {
	// The Cloud Pub/Sub subscription name
	Subscription         string   `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubsubOptions) Reset()         { *m = PubsubOptions{} }
func (m *PubsubOptions) String() string { return proto.CompactTextString(m) }
func (*PubsubOptions) ProtoMessage()    {}
func (*PubsubOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{4}
}
func (m *PubsubOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubsubOptions.Unmarshal(m, b)
}
func (m *PubsubOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubsubOptions.Marshal(b, m, deterministic)
}
func (dst *PubsubOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubsubOptions.Merge(dst, src)
}
func (m *PubsubOptions) XXX_Size() int {
	return xxx_messageInfo_PubsubOptions.Size(m)
}
func (m *PubsubOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_PubsubOptions.DiscardUnknown(m)
}

var xxx_messageInfo_PubsubOptions proto.InternalMessageInfo

func (m *PubsubOptions) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

type KafkaOptions struct {
	// The network address of the Kafka broker.
	Broker string `protobuf:"bytes,1,opt,name=broker,proto3" json:"broker,omitempty"`
	// The length of time to poll for.
	PollDuration *duration.Duration `protobuf:"bytes,2,opt,name=poll_duration,json=pollDuration,proto3" json:"poll_duration,omitempty"`
	// The network address(es) of the Zookeeper host(s).
	ZookeeperIpAddress string `protobuf:"bytes,3,opt,name=zookeeper_ip_address,json=zookeeperIpAddress,proto3" json:"zookeeper_ip_address,omitempty"`
	// The replication factor of the Kafka topic.
	ReplicationFactor int32 `protobuf:"varint,4,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty"`
	// The number of partitions of the Kafka topic.
	Partitions           int32    `protobuf:"varint,5,opt,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaOptions) Reset()         { *m = KafkaOptions{} }
func (m *KafkaOptions) String() string { return proto.CompactTextString(m) }
func (*KafkaOptions) ProtoMessage()    {}
func (*KafkaOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{5}
}
func (m *KafkaOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaOptions.Unmarshal(m, b)
}
func (m *KafkaOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaOptions.Marshal(b, m, deterministic)
}
func (dst *KafkaOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaOptions.Merge(dst, src)
}
func (m *KafkaOptions) XXX_Size() int {
	return xxx_messageInfo_KafkaOptions.Size(m)
}
func (m *KafkaOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaOptions.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaOptions proto.InternalMessageInfo

func (m *KafkaOptions) GetBroker() string {
	if m != nil {
		return m.Broker
	}
	return ""
}

func (m *KafkaOptions) GetPollDuration() *duration.Duration {
	if m != nil {
		return m.PollDuration
	}
	return nil
}

func (m *KafkaOptions) GetZookeeperIpAddress() string {
	if m != nil {
		return m.ZookeeperIpAddress
	}
	return ""
}

func (m *KafkaOptions) GetReplicationFactor() int32 {
	if m != nil {
		return m.ReplicationFactor
	}
	return 0
}

func (m *KafkaOptions) GetPartitions() int32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

type MessageIdentifier struct {
	// The unique id of the client that published the message.
	PublisherClientId int64 `protobuf:"varint,1,opt,name=publisher_client_id,json=publisherClientId,proto3" json:"publisher_client_id,omitempty"`
	// Sequence number of the published message with the given publish_client_id.
	SequenceNumber       int32    `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageIdentifier) Reset()         { *m = MessageIdentifier{} }
func (m *MessageIdentifier) String() string { return proto.CompactTextString(m) }
func (*MessageIdentifier) ProtoMessage()    {}
func (*MessageIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{6}
}
func (m *MessageIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageIdentifier.Unmarshal(m, b)
}
func (m *MessageIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageIdentifier.Marshal(b, m, deterministic)
}
func (dst *MessageIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageIdentifier.Merge(dst, src)
}
func (m *MessageIdentifier) XXX_Size() int {
	return xxx_messageInfo_MessageIdentifier.Size(m)
}
func (m *MessageIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_MessageIdentifier proto.InternalMessageInfo

func (m *MessageIdentifier) GetPublisherClientId() int64 {
	if m != nil {
		return m.PublisherClientId
	}
	return 0
}

func (m *MessageIdentifier) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

// A message for kafka which does not support adding metadata to the message.
type KafkaMessage struct {
	Id                   *MessageIdentifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PublishTime          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	Payload              []byte               `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KafkaMessage) Reset()         { *m = KafkaMessage{} }
func (m *KafkaMessage) String() string { return proto.CompactTextString(m) }
func (*KafkaMessage) ProtoMessage()    {}
func (*KafkaMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{7}
}
func (m *KafkaMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaMessage.Unmarshal(m, b)
}
func (m *KafkaMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaMessage.Marshal(b, m, deterministic)
}
func (dst *KafkaMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaMessage.Merge(dst, src)
}
func (m *KafkaMessage) XXX_Size() int {
	return xxx_messageInfo_KafkaMessage.Size(m)
}
func (m *KafkaMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaMessage.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaMessage proto.InternalMessageInfo

func (m *KafkaMessage) GetId() *MessageIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KafkaMessage) GetPublishTime() *timestamp.Timestamp {
	if m != nil {
		return m.PublishTime
	}
	return nil
}

func (m *KafkaMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Request a statistics update.
type CheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{8}
}
func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (dst *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(dst, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

type CheckResponse struct {
	// Histogram of latencies, each one a delta from the previous CheckResponse sent.
	// The bounds of the nth bucket (starting from the 0th bucket) are
	// [1.5^(n-1), 1.5^n) milliseconds.  The lower bound of the 0th bucket is 0 seconds.
	BucketValues []int64 `protobuf:"varint,1,rep,packed,name=bucket_values,json=bucketValues,proto3" json:"bucket_values,omitempty"`
	// The duration from the start of the loadtest to its completion or now if is_finished is false.
	RunningDuration *duration.Duration `protobuf:"bytes,2,opt,name=running_duration,json=runningDuration,proto3" json:"running_duration,omitempty"`
	// True if the load test has finished running.
	IsFinished bool `protobuf:"varint,3,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	// MessageIdentifiers of all received messages since the last Check.
	ReceivedMessages []*MessageIdentifier `protobuf:"bytes,4,rep,name=received_messages,json=receivedMessages,proto3" json:"received_messages,omitempty"`
	// Number of failed messages since the last check.
	Failed               int64    `protobuf:"varint,5,opt,name=failed,proto3" json:"failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_loadtest_9b64c8b2750912ff, []int{9}
}
func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (dst *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(dst, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetBucketValues() []int64 {
	if m != nil {
		return m.BucketValues
	}
	return nil
}

func (m *CheckResponse) GetRunningDuration() *duration.Duration {
	if m != nil {
		return m.RunningDuration
	}
	return nil
}

func (m *CheckResponse) GetIsFinished() bool {
	if m != nil {
		return m.IsFinished
	}
	return false
}

func (m *CheckResponse) GetReceivedMessages() []*MessageIdentifier {
	if m != nil {
		return m.ReceivedMessages
	}
	return nil
}

func (m *CheckResponse) GetFailed() int64 {
	if m != nil {
		return m.Failed
	}
	return 0
}

func init() {
	proto.RegisterType((*StartRequest)(nil), "google.pubsub.loadtest.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "google.pubsub.loadtest.StartResponse")
	proto.RegisterType((*PublisherOptions)(nil), "google.pubsub.loadtest.PublisherOptions")
	proto.RegisterType((*SubscriberOptions)(nil), "google.pubsub.loadtest.SubscriberOptions")
	proto.RegisterType((*PubsubOptions)(nil), "google.pubsub.loadtest.PubsubOptions")
	proto.RegisterType((*KafkaOptions)(nil), "google.pubsub.loadtest.KafkaOptions")
	proto.RegisterType((*MessageIdentifier)(nil), "google.pubsub.loadtest.MessageIdentifier")
	proto.RegisterType((*KafkaMessage)(nil), "google.pubsub.loadtest.KafkaMessage")
	proto.RegisterType((*CheckRequest)(nil), "google.pubsub.loadtest.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "google.pubsub.loadtest.CheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadtestWorkerClient is the client API for LoadtestWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadtestWorkerClient interface {
	// Starts a worker
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	// Check the status of a load test worker.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type loadtestWorkerClient struct {
	cc *grpc.ClientConn
}

func NewLoadtestWorkerClient(cc *grpc.ClientConn) LoadtestWorkerClient {
	return &loadtestWorkerClient{cc}
}

func (c *loadtestWorkerClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/google.pubsub.loadtest.LoadtestWorker/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadtestWorkerClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/google.pubsub.loadtest.LoadtestWorker/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadtestWorkerServer is the server API for LoadtestWorker service.
type LoadtestWorkerServer interface {
	// Starts a worker
	Start(context.Context, *StartRequest) (*StartResponse, error)
	// Check the status of a load test worker.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

func RegisterLoadtestWorkerServer(s *grpc.Server, srv LoadtestWorkerServer) {
	s.RegisterService(&_LoadtestWorker_serviceDesc, srv)
}

func _LoadtestWorker_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadtestWorkerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.pubsub.loadtest.LoadtestWorker/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadtestWorkerServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadtestWorker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadtestWorkerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.pubsub.loadtest.LoadtestWorker/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadtestWorkerServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadtestWorker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.pubsub.loadtest.LoadtestWorker",
	HandlerType: (*LoadtestWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _LoadtestWorker_Start_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _LoadtestWorker_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loadtest.proto",
}

func init() { proto.RegisterFile("loadtest.proto", fileDescriptor_loadtest_9b64c8b2750912ff) }

var fileDescriptor_loadtest_9b64c8b2750912ff = []byte{
	// 860 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0xfc, 0x93, 0xd4, 0x27, 0x92, 0x63, 0xb3, 0x45, 0xa1, 0x05, 0x68, 0xeb, 0x69, 0x2b,
	0xe6, 0x5d, 0x4c, 0x1d, 0xda, 0xab, 0x5e, 0x6c, 0xd8, 0xd2, 0xa2, 0x58, 0xd0, 0xad, 0x2b, 0x98,
	0xa0, 0x05, 0x76, 0x23, 0x48, 0x14, 0xed, 0x70, 0x96, 0x45, 0x8e, 0xa4, 0x0a, 0x2c, 0xaf, 0xb0,
	0xb7, 0xd8, 0xe5, 0x9e, 0x61, 0xef, 0x33, 0x60, 0x4f, 0x31, 0xf0, 0x47, 0xb6, 0xe3, 0xd4, 0x09,
	0x72, 0xe7, 0xf3, 0x9d, 0xef, 0x7c, 0xe4, 0x39, 0xfa, 0x78, 0x0c, 0xc3, 0x8a, 0xe7, 0xa5, 0xa6,
	0x4a, 0xa7, 0x42, 0x72, 0xcd, 0xd1, 0xfd, 0x39, 0xe7, 0xf3, 0x8a, 0xa6, 0xa2, 0x29, 0x54, 0x53,
	0xa4, 0x6d, 0xf6, 0xe8, 0xa1, 0xc3, 0x9f, 0x58, 0x56, 0xd1, 0xcc, 0x9e, 0x94, 0x8d, 0xcc, 0x35,
	0xe3, 0xb5, 0xab, 0x3b, 0x7a, 0xb4, 0x9d, 0xd7, 0x6c, 0x49, 0x95, 0xce, 0x97, 0xc2, 0x11, 0x92,
	0xff, 0x7a, 0x10, 0x9e, 0xea, 0x5c, 0x6a, 0x4c, 0x7f, 0x6f, 0xa8, 0xd2, 0x28, 0x86, 0x7d, 0x21,
	0xf9, 0x6f, 0x94, 0xe8, 0x38, 0x98, 0x04, 0xd3, 0x01, 0x6e, 0x43, 0x74, 0x0f, 0xfa, 0x9a, 0x0b,
	0x46, 0xe2, 0x8e, 0xc5, 0x5d, 0x80, 0x9e, 0x03, 0x28, 0x53, 0x9f, 0x19, 0xe5, 0xb8, 0x3b, 0x09,
	0xa6, 0x07, 0x4f, 0x8f, 0xd2, 0xf6, 0xba, 0xfe, 0xd8, 0xf4, 0xac, 0x3d, 0x16, 0x0f, 0x2c, 0xdb,
	0xc4, 0xe8, 0x3b, 0x88, 0x4c, 0x13, 0x59, 0x7b, 0xe7, 0xb8, 0x67, 0xab, 0x3f, 0xbd, 0x52, 0xfd,
	0xd2, 0x13, 0x70, 0x68, 0xf8, 0x6d, 0x84, 0x1e, 0xc1, 0x01, 0xab, 0x49, 0xd5, 0x94, 0x34, 0x63,
	0xa5, 0x8a, 0xfb, 0x93, 0x60, 0x7a, 0x07, 0x83, 0x87, 0x4e, 0x4a, 0x85, 0xde, 0xc0, 0xd0, 0x0d,
	0x2c, 0xe3, 0xc2, 0x54, 0xa8, 0x78, 0xcf, 0x9e, 0xf0, 0x38, 0xfd, 0xf8, 0x38, 0xd3, 0xb7, 0x36,
	0xfe, 0xc5, 0x91, 0x7f, 0xfc, 0x04, 0x47, 0x62, 0x13, 0x40, 0xaf, 0x21, 0x5a, 0xe4, 0xb3, 0x45,
	0xbe, 0x92, 0xdb, 0xb7, 0x72, 0x5f, 0xec, 0x92, 0x7b, 0x6d, 0xc8, 0x6b, 0xb5, 0x70, 0xb1, 0x11,
	0xa3, 0xf7, 0x30, 0x16, 0x4d, 0x51, 0x31, 0x75, 0x4e, 0xe5, 0x4a, 0xf0, 0x8e, 0x15, 0x9c, 0x5e,
	0x73, 0x3f, 0x57, 0xd0, 0x8a, 0x06, 0x78, 0x24, 0xb6, 0x30, 0xf4, 0x2b, 0x20, 0xd5, 0x14, 0x8a,
	0x48, 0x56, 0x6c, 0x28, 0x0f, 0xac, 0xf2, 0x57, 0xbb, 0x94, 0x4f, 0x57, 0x15, 0x6b, 0xe9, 0xb1,
	0xda, 0x06, 0xcd, 0xc8, 0x89, 0x68, 0x32, 0x45, 0xf2, 0x8a, 0xd5, 0xf3, 0x18, 0x26, 0xc1, 0xb4,
	0x8f, 0x81, 0x88, 0xe6, 0xd4, 0x21, 0xc7, 0x03, 0xd8, 0xf7, 0x27, 0x1e, 0x8f, 0x60, 0x48, 0x2a,
	0x46, 0x6b, 0xdd, 0xde, 0x21, 0x39, 0x84, 0xc8, 0x7b, 0x4d, 0x09, 0x5e, 0x2b, 0x9a, 0xfc, 0x1d,
	0xc0, 0x68, 0xbb, 0x27, 0x84, 0xa0, 0x27, 0x73, 0x4d, 0xad, 0xfd, 0x3a, 0xd8, 0xfe, 0x46, 0xdf,
	0xc3, 0xb0, 0xc8, 0x35, 0x39, 0x5f, 0x7b, 0xa5, 0x73, 0x93, 0x57, 0x22, 0x5b, 0xb0, 0x32, 0xcb,
	0x03, 0x00, 0xa7, 0xa0, 0xd8, 0x85, 0xf3, 0x69, 0x1f, 0x0f, 0x2c, 0x72, 0xca, 0x2e, 0x28, 0xfa,
	0x0c, 0xc2, 0x25, 0x55, 0x2a, 0x9f, 0x53, 0x47, 0xe8, 0x59, 0xc2, 0x81, 0xc7, 0x0c, 0x25, 0xb9,
	0x0b, 0xe3, 0x2b, 0x53, 0x4a, 0x9e, 0x41, 0x74, 0xc9, 0x34, 0x28, 0x81, 0xd0, 0x8f, 0xcd, 0x02,
	0xfe, 0x11, 0x5d, 0xc2, 0x92, 0x7f, 0x03, 0x08, 0x37, 0xbd, 0x81, 0xee, 0xc3, 0x5e, 0x21, 0xf9,
	0x82, 0x4a, 0x4f, 0xf7, 0x91, 0x79, 0x21, 0x82, 0x57, 0xd5, 0x2d, 0xba, 0x0e, 0x0d, 0x7f, 0xd5,
	0xf4, 0x37, 0x70, 0xef, 0x82, 0xf3, 0x05, 0xa5, 0x82, 0xca, 0x8c, 0x89, 0x2c, 0x2f, 0x4b, 0x49,
	0x95, 0xb2, 0xed, 0x0f, 0x30, 0x5a, 0xe5, 0x4e, 0xc4, 0x0f, 0x2e, 0x83, 0xbe, 0x06, 0x24, 0xa9,
	0xa8, 0x18, 0xb1, 0x02, 0xd9, 0x2c, 0x27, 0x9a, 0x4b, 0x3f, 0x8d, 0xf1, 0x46, 0xe6, 0x95, 0x4d,
	0xa0, 0x87, 0x00, 0x22, 0x97, 0x9a, 0x39, 0x8f, 0xf5, 0x9d, 0x1d, 0xd6, 0x48, 0x52, 0xc1, 0xf8,
	0x67, 0x37, 0xc2, 0x93, 0x92, 0xd6, 0x9a, 0xcd, 0x18, 0x95, 0x28, 0x85, 0xbb, 0x6b, 0xe7, 0x7b,
	0x8b, 0xb0, 0xd2, 0xb6, 0xde, 0xc5, 0xeb, 0x47, 0xf1, 0xc2, 0x66, 0x4e, 0x4a, 0xf4, 0x25, 0x1c,
	0x2a, 0xb3, 0x9d, 0x6a, 0x42, 0xb3, 0xba, 0x59, 0x16, 0x54, 0xda, 0x39, 0xf4, 0xf1, 0xb0, 0x85,
	0xdf, 0x58, 0x34, 0xf9, 0xab, 0x9d, 0xab, 0x3f, 0x13, 0x3d, 0x87, 0x8e, 0x17, 0xbe, 0xc6, 0xfa,
	0x57, 0x2e, 0x88, 0x3b, 0xac, 0x44, 0xdf, 0x42, 0xe8, 0x6f, 0xe2, 0x36, 0x5b, 0xe7, 0xc6, 0xcd,
	0x76, 0xe0, 0xf9, 0x76, 0xb7, 0x99, 0x35, 0x9a, 0xff, 0x61, 0xce, 0xb0, 0xc3, 0x0e, 0x71, 0x1b,
	0x26, 0x43, 0x08, 0x5f, 0x9c, 0x53, 0xb2, 0xf0, 0x0b, 0x37, 0xf9, 0xb3, 0x03, 0x91, 0x07, 0xdc,
	0xab, 0x40, 0x9f, 0x43, 0x54, 0x34, 0x64, 0x41, 0x75, 0xf6, 0x21, 0xaf, 0x1a, 0xaa, 0xe2, 0x60,
	0xd2, 0x9d, 0x76, 0x71, 0xe8, 0xc0, 0x77, 0x16, 0x43, 0x2f, 0x61, 0x24, 0x9b, 0xba, 0x66, 0xf5,
	0xfc, 0x16, 0xee, 0x38, 0xf4, 0x25, 0x97, 0x56, 0xa8, 0xca, 0x66, 0xac, 0x36, 0x13, 0x77, 0x57,
	0x35, 0x2b, 0x54, 0xbd, 0xf2, 0x08, 0x7a, 0x07, 0x63, 0x49, 0x09, 0x65, 0x1f, 0x68, 0x99, 0xf9,
	0xc7, 0xa0, 0xe2, 0xde, 0xa4, 0x7b, 0xbb, 0x81, 0x8e, 0x5a, 0x0d, 0x9f, 0xb2, 0x8e, 0x9f, 0xe5,
	0xac, 0xa2, 0xa5, 0x35, 0x4d, 0x17, 0xfb, 0xe8, 0xe9, 0x3f, 0x01, 0x0c, 0x7f, 0xf2, 0x42, 0xef,
	0xb9, 0x34, 0x8f, 0xe0, 0x0c, 0xfa, 0x76, 0x6b, 0xa0, 0x9d, 0x7b, 0x76, 0xf3, 0x0f, 0xec, 0xe8,
	0xf1, 0x0d, 0x2c, 0x3f, 0xe4, 0x33, 0xe8, 0xdb, 0xa9, 0xef, 0x56, 0xdd, 0xfc, 0x4a, 0xbb, 0x55,
	0x2f, 0x7d, 0xba, 0xe3, 0x14, 0x1e, 0x10, 0xbe, 0xdc, 0xe2, 0xce, 0x2a, 0x46, 0x52, 0xc2, 0x97,
	0x4b, 0x5e, 0x1f, 0x47, 0x6d, 0x73, 0x6f, 0xed, 0xc7, 0xd9, 0xb3, 0xdf, 0xe8, 0xd9, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x07, 0x11, 0x92, 0x7a, 0xf0, 0x07, 0x00, 0x00,
}
