package handler

import (
	context "context"

	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MQEndpointsClient is the client API for MQEndpoints service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MQEndpointsClient interface {
	GetMessages(ctx context.Context, in *RequestMessageData, opts ...grpc.CallOption) (*MessageCollection, error)
	SubmitMessage(ctx context.Context, in *MessageInfo, opts ...grpc.CallOption) (*Status, error)
	CreateQueue(ctx context.Context, in *QueueInfo, opts ...grpc.CallOption) (*Status, error)
	Ack(ctx context.Context, in *AckUpdate, opts ...grpc.CallOption) (*Status, error)
	Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Status, error)
	ChangePassword(ctx context.Context, in *ChangeCredentials, opts ...grpc.CallOption) (*Status, error)
	Nack(ctx context.Context, in *AckUpdate, opts ...grpc.CallOption) (*Status, error)
	DeleteQueue(ctx context.Context, in *DeleteQueueInfo, opts ...grpc.CallOption) (*Status, error)
	AddRoutingKey(ctx context.Context, in *AddRoute, opts ...grpc.CallOption) (*Status, error)
	DeleteRoutingKey(ctx context.Context, in *DeleteRoute, opts ...grpc.CallOption) (*Status, error)
	CreateUser(ctx context.Context, in *UserCreds, opts ...grpc.CallOption) (*Status, error)
	SubmitBatchedMessages(ctx context.Context, in *BatchMessages, opts ...grpc.CallOption) (*Status, error)
	IsLeaderNode(ctx context.Context, in *LeaderNodeRequest, opts ...grpc.CallOption) (*Status, error)
	DeleteUser(ctx context.Context, in *UserInformation, opts ...grpc.CallOption) (*Status, error)
	BatchAck(ctx context.Context, in *BatchAckUpdate, opts ...grpc.CallOption) (*Status, error)
	BatchNack(ctx context.Context, in *BatchNackUpdate, opts ...grpc.CallOption) (*Status, error)
}

type mQEndpointsClient struct {
	cc grpc.ClientConnInterface
}

func NewMQEndpointsClient(cc grpc.ClientConnInterface) MQEndpointsClient {
	return &mQEndpointsClient{cc}
}

func (c *mQEndpointsClient) GetMessages(ctx context.Context, in *RequestMessageData, opts ...grpc.CallOption) (*MessageCollection, error) {
	out := new(MessageCollection)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) SubmitMessage(ctx context.Context, in *MessageInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/SubmitMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) CreateQueue(ctx context.Context, in *QueueInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/CreateQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) Ack(ctx context.Context, in *AckUpdate, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) ChangePassword(ctx context.Context, in *ChangeCredentials, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) Nack(ctx context.Context, in *AckUpdate, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/Nack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) DeleteQueue(ctx context.Context, in *DeleteQueueInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/DeleteQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) AddRoutingKey(ctx context.Context, in *AddRoute, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/AddRoutingKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) DeleteRoutingKey(ctx context.Context, in *DeleteRoute, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/DeleteRoutingKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) CreateUser(ctx context.Context, in *UserCreds, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) SubmitBatchedMessages(ctx context.Context, in *BatchMessages, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/SubmitBatchedMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) IsLeaderNode(ctx context.Context, in *LeaderNodeRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/IsLeaderNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) DeleteUser(ctx context.Context, in *UserInformation, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) BatchAck(ctx context.Context, in *BatchAckUpdate, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/BatchAck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mQEndpointsClient) BatchNack(ctx context.Context, in *BatchNackUpdate, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/MQ.MQEndpoints/BatchNack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
