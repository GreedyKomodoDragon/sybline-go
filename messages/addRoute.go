package messages

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type AddRoute struct {
	State         protoimpl.MessageState
	SizeCache     protoimpl.SizeCache
	UnknownFields protoimpl.UnknownFields

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

func (*AddRoute) ProtoMessage() {
	// ProtoMessage acts as a tag to make sure no one accidentally implements the
	// proto.Message interface.
}

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
