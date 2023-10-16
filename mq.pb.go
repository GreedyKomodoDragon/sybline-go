package handler

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutingKey string `protobuf:"bytes,1,opt,name=routingKey,proto3" json:"routingKey,omitempty"`
	Data       []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MessageInfo) Reset() {
	*x = MessageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageInfo) ProtoMessage() {}

func (x *MessageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageInfo.ProtoReflect.Descriptor instead.
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{0}
}

func (x *MessageInfo) GetRoutingKey() string {
	if x != nil {
		return x.RoutingKey
	}
	return ""
}

func (x *MessageInfo) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type MessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MessageData) Reset() {
	*x = MessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageData) ProtoMessage() {}

func (x *MessageData) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageData.ProtoReflect.Descriptor instead.
func (*MessageData) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{1}
}

func (x *MessageData) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *MessageData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type MessageCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*MessageData `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MessageCollection) Reset() {
	*x = MessageCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCollection) ProtoMessage() {}

func (x *MessageCollection) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCollection.ProtoReflect.Descriptor instead.
func (*MessageCollection) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{2}
}

func (x *MessageCollection) GetMessages() []*MessageData {
	if x != nil {
		return x.Messages
	}
	return nil
}

type RequestMessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueName string `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
	Amount    uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *RequestMessageData) Reset() {
	*x = RequestMessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessageData) ProtoMessage() {}

func (x *RequestMessageData) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessageData.ProtoReflect.Descriptor instead.
func (*RequestMessageData) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{3}
}

func (x *RequestMessageData) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *RequestMessageData) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type QueueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutingKey string `protobuf:"bytes,1,opt,name=routingKey,proto3" json:"routingKey,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size       uint32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	RetryLimit uint32 `protobuf:"varint,4,opt,name=retryLimit,proto3" json:"retryLimit,omitempty"`
	HasDLQueue bool   `protobuf:"varint,5,opt,name=hasDLQueue,proto3" json:"hasDLQueue,omitempty"`
}

func (x *QueueInfo) Reset() {
	*x = QueueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueInfo) ProtoMessage() {}

func (x *QueueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueInfo.ProtoReflect.Descriptor instead.
func (*QueueInfo) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{5}
}

func (x *QueueInfo) GetRoutingKey() string {
	if x != nil {
		return x.RoutingKey
	}
	return ""
}

func (x *QueueInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueueInfo) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *QueueInfo) GetRetryLimit() uint32 {
	if x != nil {
		return x.RetryLimit
	}
	return 0
}

func (x *QueueInfo) GetHasDLQueue() bool {
	if x != nil {
		return x.HasDLQueue
	}
	return false
}

type AckUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueName string `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
	Id        []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AckUpdate) Reset() {
	*x = AckUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckUpdate) ProtoMessage() {}

func (x *AckUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckUpdate.ProtoReflect.Descriptor instead.
func (*AckUpdate) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{6}
}

func (x *AckUpdate) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *AckUpdate) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{7}
}

func (x *Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ChangeCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
}

func (x *ChangeCredentials) Reset() {
	*x = ChangeCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeCredentials) ProtoMessage() {}

func (x *ChangeCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeCredentials.ProtoReflect.Descriptor instead.
func (*ChangeCredentials) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{8}
}

func (x *ChangeCredentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChangeCredentials) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangeCredentials) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type DeleteQueueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueName string `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
}

func (x *DeleteQueueInfo) Reset() {
	*x = DeleteQueueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQueueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQueueInfo) ProtoMessage() {}

func (x *DeleteQueueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQueueInfo.ProtoReflect.Descriptor instead.
func (*DeleteQueueInfo) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteQueueInfo) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

type AddRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueName string `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
	RouteName string `protobuf:"bytes,2,opt,name=routeName,proto3" json:"routeName,omitempty"`
}

func (x *AddRoute) Reset() {
	*x = AddRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoute) ProtoMessage() {}

func (x *AddRoute) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoute.ProtoReflect.Descriptor instead.
func (*AddRoute) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{10}
}

func (x *AddRoute) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *AddRoute) GetRouteName() string {
	if x != nil {
		return x.RouteName
	}
	return ""
}

type DeleteRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueName string `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
	RouteName string `protobuf:"bytes,2,opt,name=routeName,proto3" json:"routeName,omitempty"`
}

func (x *DeleteRoute) Reset() {
	*x = DeleteRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoute) ProtoMessage() {}

func (x *DeleteRoute) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoute.ProtoReflect.Descriptor instead.
func (*DeleteRoute) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteRoute) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *DeleteRoute) GetRouteName() string {
	if x != nil {
		return x.RouteName
	}
	return ""
}

type UserCreds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserCreds) Reset() {
	*x = UserCreds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreds) ProtoMessage() {}

func (x *UserCreds) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreds.ProtoReflect.Descriptor instead.
func (*UserCreds) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{12}
}

func (x *UserCreds) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCreds) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type BatchMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*MessageInfo `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *BatchMessages) Reset() {
	*x = BatchMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchMessages) ProtoMessage() {}

func (x *BatchMessages) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchMessages.ProtoReflect.Descriptor instead.
func (*BatchMessages) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{13}
}

func (x *BatchMessages) GetMessages() []*MessageInfo {
	if x != nil {
		return x.Messages
	}
	return nil
}

type LeaderNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaderNodeRequest) Reset() {
	*x = LeaderNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderNodeRequest) ProtoMessage() {}

func (x *LeaderNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderNodeRequest.ProtoReflect.Descriptor instead.
func (*LeaderNodeRequest) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{14}
}

type UserInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserInformation) Reset() {
	*x = UserInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInformation) ProtoMessage() {}

func (x *UserInformation) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInformation.ProtoReflect.Descriptor instead.
func (*UserInformation) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{15}
}

func (x *UserInformation) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type BatchAckUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueName string   `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
	Id        [][]byte `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *BatchAckUpdate) Reset() {
	*x = BatchAckUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchAckUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchAckUpdate) ProtoMessage() {}

func (x *BatchAckUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchAckUpdate.ProtoReflect.Descriptor instead.
func (*BatchAckUpdate) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{16}
}

func (x *BatchAckUpdate) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *BatchAckUpdate) GetId() [][]byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type BatchNackUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueName string   `protobuf:"bytes,1,opt,name=queueName,proto3" json:"queueName,omitempty"`
	Ids       [][]byte `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *BatchNackUpdate) Reset() {
	*x = BatchNackUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchNackUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchNackUpdate) ProtoMessage() {}

func (x *BatchNackUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchNackUpdate.ProtoReflect.Descriptor instead.
func (*BatchNackUpdate) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{17}
}

func (x *BatchNackUpdate) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *BatchNackUpdate) GetIds() [][]byte {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_mq_proto protoreflect.FileDescriptor

var file_mq_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x4d, 0x51, 0x22, 0x41,
	0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x31, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x51,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x61,
	0x73, 0x44, 0x4c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x68, 0x61, 0x73, 0x44, 0x4c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x09, 0x41, 0x63,
	0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x73, 0x0a, 0x11,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x2f, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x46, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x4d, 0x51, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x0e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0f,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x61, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x69, 0x64, 0x73, 0x32,
	0x8c, 0x06, 0x0a, 0x0b, 0x4d, 0x51, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x4d, 0x51, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x4d, 0x51, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12,
	0x2e, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0f, 0x2e, 0x4d, 0x51, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12,
	0x2a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x0d,
	0x2e, 0x4d, 0x51, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0a, 0x2e,
	0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x03, 0x41,
	0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x4d, 0x51, 0x2e, 0x41, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12,
	0x26, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x4d, 0x51, 0x2e, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x4d, 0x51, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x23,
	0x0a, 0x04, 0x4e, 0x61, 0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x4d, 0x51, 0x2e, 0x41, 0x63, 0x6b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x13, 0x2e, 0x4d, 0x51, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x0c, 0x2e, 0x4d, 0x51, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x0f, 0x2e, 0x4d, 0x51, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x4d, 0x51, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x64, 0x73, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x4d, 0x51, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0a, 0x2e, 0x4d,
	0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c, 0x49, 0x73,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x4d, 0x51, 0x2e,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e,
	0x4d, 0x51, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x12, 0x2c, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x4d,
	0x51, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x61, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x4d, 0x51,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x61, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x1a, 0x0a, 0x2e, 0x4d, 0x51, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mq_proto_rawDescOnce sync.Once
	file_mq_proto_rawDescData = file_mq_proto_rawDesc
)

func file_mq_proto_rawDescGZIP() []byte {
	file_mq_proto_rawDescOnce.Do(func() {
		file_mq_proto_rawDescData = protoimpl.X.CompressGZIP(file_mq_proto_rawDescData)
	})
	return file_mq_proto_rawDescData
}

var file_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_mq_proto_goTypes = []interface{}{
	(*MessageInfo)(nil),        // 0: MQ.MessageInfo
	(*MessageData)(nil),        // 1: MQ.MessageData
	(*MessageCollection)(nil),  // 2: MQ.MessageCollection
	(*RequestMessageData)(nil), // 3: MQ.RequestMessageData
	(*Status)(nil),             // 4: MQ.Status
	(*QueueInfo)(nil),          // 5: MQ.QueueInfo
	(*AckUpdate)(nil),          // 6: MQ.AckUpdate
	(*Credentials)(nil),        // 7: MQ.Credentials
	(*ChangeCredentials)(nil),  // 8: MQ.ChangeCredentials
	(*DeleteQueueInfo)(nil),    // 9: MQ.DeleteQueueInfo
	(*AddRoute)(nil),           // 10: MQ.AddRoute
	(*DeleteRoute)(nil),        // 11: MQ.DeleteRoute
	(*UserCreds)(nil),          // 12: MQ.UserCreds
	(*BatchMessages)(nil),      // 13: MQ.BatchMessages
	(*LeaderNodeRequest)(nil),  // 14: MQ.LeaderNodeRequest
	(*UserInformation)(nil),    // 15: MQ.UserInformation
	(*BatchAckUpdate)(nil),     // 16: MQ.BatchAckUpdate
	(*BatchNackUpdate)(nil),    // 17: MQ.BatchNackUpdate
}
var file_mq_proto_depIdxs = []int32{
	1,  // 0: MQ.MessageCollection.messages:type_name -> MQ.MessageData
	0,  // 1: MQ.BatchMessages.messages:type_name -> MQ.MessageInfo
	3,  // 2: MQ.MQEndpoints.GetMessages:input_type -> MQ.RequestMessageData
	0,  // 3: MQ.MQEndpoints.SubmitMessage:input_type -> MQ.MessageInfo
	5,  // 4: MQ.MQEndpoints.CreateQueue:input_type -> MQ.QueueInfo
	6,  // 5: MQ.MQEndpoints.Ack:input_type -> MQ.AckUpdate
	7,  // 6: MQ.MQEndpoints.Login:input_type -> MQ.Credentials
	8,  // 7: MQ.MQEndpoints.ChangePassword:input_type -> MQ.ChangeCredentials
	6,  // 8: MQ.MQEndpoints.Nack:input_type -> MQ.AckUpdate
	9,  // 9: MQ.MQEndpoints.DeleteQueue:input_type -> MQ.DeleteQueueInfo
	10, // 10: MQ.MQEndpoints.AddRoutingKey:input_type -> MQ.AddRoute
	11, // 11: MQ.MQEndpoints.DeleteRoutingKey:input_type -> MQ.DeleteRoute
	12, // 12: MQ.MQEndpoints.CreateUser:input_type -> MQ.UserCreds
	13, // 13: MQ.MQEndpoints.SubmitBatchedMessages:input_type -> MQ.BatchMessages
	14, // 14: MQ.MQEndpoints.IsLeaderNode:input_type -> MQ.LeaderNodeRequest
	15, // 15: MQ.MQEndpoints.DeleteUser:input_type -> MQ.UserInformation
	16, // 16: MQ.MQEndpoints.BatchAck:input_type -> MQ.BatchAckUpdate
	17, // 17: MQ.MQEndpoints.BatchNack:input_type -> MQ.BatchNackUpdate
	2,  // 18: MQ.MQEndpoints.GetMessages:output_type -> MQ.MessageCollection
	4,  // 19: MQ.MQEndpoints.SubmitMessage:output_type -> MQ.Status
	4,  // 20: MQ.MQEndpoints.CreateQueue:output_type -> MQ.Status
	4,  // 21: MQ.MQEndpoints.Ack:output_type -> MQ.Status
	4,  // 22: MQ.MQEndpoints.Login:output_type -> MQ.Status
	4,  // 23: MQ.MQEndpoints.ChangePassword:output_type -> MQ.Status
	4,  // 24: MQ.MQEndpoints.Nack:output_type -> MQ.Status
	4,  // 25: MQ.MQEndpoints.DeleteQueue:output_type -> MQ.Status
	4,  // 26: MQ.MQEndpoints.AddRoutingKey:output_type -> MQ.Status
	4,  // 27: MQ.MQEndpoints.DeleteRoutingKey:output_type -> MQ.Status
	4,  // 28: MQ.MQEndpoints.CreateUser:output_type -> MQ.Status
	4,  // 29: MQ.MQEndpoints.SubmitBatchedMessages:output_type -> MQ.Status
	4,  // 30: MQ.MQEndpoints.IsLeaderNode:output_type -> MQ.Status
	4,  // 31: MQ.MQEndpoints.DeleteUser:output_type -> MQ.Status
	4,  // 32: MQ.MQEndpoints.BatchAck:output_type -> MQ.Status
	4,  // 33: MQ.MQEndpoints.BatchNack:output_type -> MQ.Status
	18, // [18:34] is the sub-list for method output_type
	2,  // [2:18] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_mq_proto_init() }
func file_mq_proto_init() {
	if File_mq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCollection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessageData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeCredentials); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQueueInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchMessages); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInformation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchAckUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mq_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchNackUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mq_proto_goTypes,
		DependencyIndexes: file_mq_proto_depIdxs,
		MessageInfos:      file_mq_proto_msgTypes,
	}.Build()
	File_mq_proto = out.File
	file_mq_proto_rawDesc = nil
	file_mq_proto_goTypes = nil
	file_mq_proto_depIdxs = nil
}
