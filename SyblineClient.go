package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	ErrMissingToken      = errors.New("no valid auth token provided")
	ErrNoServerAddresses = errors.New("no addresses have been provided")
	ErrNoValidAddresses  = errors.New("client must be provided at least one address")

	transportErrString = "transport: Error while dialing dial tcp"
	serverErr          = "error reading from server: EOF"
	leadersErrString   = "cannot communicate with non-leader nodes"
	AuthErrString      = "no valid auth token provided"

	SYB_TOKEN = "syb-token"
)

type SyblineClient interface {
	SubmitMessage(ctx context.Context, routingKey string, data []byte) error
	SubmitBatchMessage(ctx context.Context, msg []Message) error
	GetMessages(ctx context.Context, queue string, amount uint32) ([]*MessageData, error)
	CreateQueue(ctx context.Context, routing, name string, size, retryLimit uint32, hasDLQ bool) error
	DeleteOueue(ctx context.Context, name string) error
	Ack(ctx context.Context, queue string, id []byte) error
	Login(ctx context.Context, username string) error
	ChangePassword(ctx context.Context, username, oldPassword, newPassword string) error
	Nack(ctx context.Context, queue string, id []byte) error
	AddRoutingKey(ctx context.Context, queue, routingKey string) error
	DeleteRoutingKey(ctx context.Context, queue, routingKey string) error
	CreateUser(ctx context.Context, username, password string) error
	DeleteUser(ctx context.Context, usernam string) error
	BatchAck(ctx context.Context, queue string, ids [][]byte) error
	BatchNack(ctx context.Context, queue string, ids [][]byte) error
	Consumer(capacity int, time time.Duration, queue string) (*syblineConsumer, error)
	Logout(ctx context.Context) error
	Close()
}

type syblineClient struct {
	headers         metadata.MD
	gClient         MQEndpointsClient
	conn            *grpc.ClientConn
	opts            []grpc.DialOption
	caFile          string
	addresses       []string
	passwordManager PasswordManager
	username        string
	reconnectLock   *sync.Mutex
	login           *sync.Mutex
	lastLog         time.Time
	config          Config
}

func NewBasicSyblineClient(serverAddr []string, passwordManager PasswordManager, config Config) (SyblineClient, error) {
	if len(serverAddr) == 0 {
		return nil, ErrNoServerAddresses
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	for _, address := range serverAddr {
		conn, err := grpc.Dial(address, opts...)
		if err != nil {
			continue
		}

		gClient := NewMQEndpointsClient(conn)
		status, err := gClient.IsLeaderNode(context.Background(), &LeaderNodeRequest{})
		if err != nil || !status.Status {
			continue
		}

		return &syblineClient{
			gClient:         NewMQEndpointsClient(conn),
			conn:            conn,
			headers:         nil,
			caFile:          "",
			addresses:       serverAddr,
			opts:            opts,
			passwordManager: passwordManager,
			username:        "",
			reconnectLock:   &sync.Mutex{},
			login:           &sync.Mutex{},
			config:          config,
		}, nil
	}

	return nil, ErrNoValidAddresses
}

func NewTLSSyblineClient(serverAddr []string, caFile, certFile, keyFile string, skipVerification bool, passwordManager PasswordManager, config Config) (SyblineClient, error) {
	if len(serverAddr) == 0 {
		return nil, ErrNoServerAddresses
	}

	tlsConfig, err := createTLSConfig(caFile, certFile, keyFile, skipVerification)
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))}

	for _, address := range serverAddr {
		conn, err := grpc.Dial(address, opts...)
		if err != nil {
			continue
		}

		gClient := NewMQEndpointsClient(conn)
		status, err := gClient.IsLeaderNode(context.Background(), &LeaderNodeRequest{})
		if err != nil || !status.Status {
			continue
		}

		return &syblineClient{
			gClient:         NewMQEndpointsClient(conn),
			conn:            conn,
			headers:         nil,
			addresses:       serverAddr,
			opts:            opts,
			passwordManager: passwordManager,
			username:        "",
			reconnectLock:   &sync.Mutex{},
			login:           &sync.Mutex{},
			lastLog:         time.Unix(0, 0),
			config:          config,
		}, nil
	}

	return nil, ErrNoValidAddresses
}

func (c *syblineClient) SubmitMessage(ctx context.Context, routingKey string, data []byte) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateMessage(data, routingKey); err != nil {
		return err
	}

	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	for {
		message, err := c.gClient.SubmitMessage(ctx, &MessageInfo{
			RoutingKey: routingKey,
			Data:       data,
		})

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if message != nil && message.Status {
			return nil
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("message was unable to be submitted")
		}

		timeout.Increment()
		timeout.Sleep()

		ctx = metadata.NewOutgoingContext(ctx, c.headers)

	}
}

func (c *syblineClient) SubmitBatchMessage(ctx context.Context, msgRaw []Message) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	msgs := []*MessageInfo{}
	for _, data := range msgRaw {
		if err := validateMessage(data.Data, data.Rk); err != nil {
			return err
		}

		msgs = append(msgs, &MessageInfo{
			RoutingKey: data.Rk,
			Data:       data.Data,
		})
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	msgsWrapped := &BatchMessages{
		Messages: msgs,
	}

	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)
	for {
		message, err := c.gClient.SubmitBatchedMessages(ctx, msgsWrapped)

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if message != nil && message.Status {
			return nil
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("message was unable to be submitted")
		}

		timeout.Increment()
		timeout.Sleep()

		ctx = metadata.NewOutgoingContext(ctx, c.headers)
	}
}

func (c *syblineClient) GetMessages(ctx context.Context, queue string, amount uint32) ([]*MessageData, error) {
	if c.headers == nil {
		return nil, ErrMissingToken
	}

	if err := validateRequestMessages(queue, amount); err != nil {
		return nil, err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)

	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)
	payload := &RequestMessageData{
		QueueName: queue,
		Amount:    amount,
	}

	for {
		msgs, err := c.gClient.GetMessages(ctx, payload)

		if err := c.handleErr(ctx, err); err != nil {
			return nil, err
		}

		if msgs != nil {
			return msgs.Messages, nil
		}

		if timeout.HasThreadReached() {
			return nil, fmt.Errorf("unable to get messages")
		}

		timeout.Increment()
		timeout.Sleep()

		ctx = metadata.NewOutgoingContext(ctx, c.headers)
	}
}

func (c *syblineClient) CreateQueue(ctx context.Context, routing, name string, size, retryLimit uint32, hasDLQ bool) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateCreateQueue(routing, name, size); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)

	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)
	payload := &QueueInfo{
		RoutingKey: routing,
		Name:       name,
		Size:       size,
		RetryLimit: retryLimit,
		HasDLQueue: hasDLQ,
	}

	for {
		status, err := c.gClient.CreateQueue(ctx, payload)
		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to create queue for %s with size: %v", name, size)
		}

		timeout.Increment()
		timeout.Sleep()

		ctx = metadata.NewOutgoingContext(ctx, c.headers)
	}

}

func (c *syblineClient) Ack(ctx context.Context, queue string, id []byte) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateAck(queue, id); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	payload := &AckUpdate{
		QueueName: queue,
		Id:        id,
	}

	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.Ack(ctx, payload)
		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to acknowledge message")
		}

		timeout.Increment()
		timeout.Sleep()
		ctx = metadata.NewOutgoingContext(ctx, c.headers)
	}
}

func (c *syblineClient) BatchAck(ctx context.Context, queue string, ids [][]byte) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateBatchAck(queue, ids); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	payload := &BatchAckUpdate{
		QueueName: queue,
		Id:        ids,
	}

	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.BatchAck(ctx, payload)
		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to batch acknowledge messages")
		}

		timeout.Increment()
		timeout.Sleep()
		ctx = metadata.NewOutgoingContext(ctx, c.headers)
	}
}

func (c *syblineClient) Login(ctx context.Context, username string) error {
	c.login.Lock()
	defer c.login.Unlock()

	if len(username) == 0 {
		return ErrInvalidUsername
	}

	if time.Since(c.lastLog) < time.Second*30 {
		return nil
	}

	var headerReturn metadata.MD

	password, err := c.passwordManager.GetPassword(username)
	if err != nil {
		return err
	}

	payload := &Credentials{
		Username: username,
		Password: password,
	}

	head := grpc.Header(&headerReturn)
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for !timeout.HasThreadReached() {
		status, err := c.gClient.Login(ctx, payload, head)
		if err == nil && status.Status {
			c.headers = headerReturn
			c.username = username
			c.lastLog = time.Now()
			return nil
		}

		timeout.Increment()
		timeout.Sleep()
	}

	return fmt.Errorf("unable to login")
}

func (c *syblineClient) ChangePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateChangePassword(username, oldPassword, newPassword); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {

		status, err := c.gClient.ChangePassword(ctx, &ChangeCredentials{
			Username:    username,
			OldPassword: oldPassword,
			NewPassword: newPassword,
		})
		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to update password")

		}

		timeout.Increment()
		timeout.Sleep()
	}
}

func (c *syblineClient) Nack(ctx context.Context, queue string, id []byte) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateAck(queue, id); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.Nack(ctx, &AckUpdate{
			QueueName: queue,
			Id:        id,
		})

		if err == nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to nack message on queue")

		}

		timeout.Increment()
		timeout.Sleep()

	}
}

func (c *syblineClient) BatchNack(ctx context.Context, queue string, ids [][]byte) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateBatchAck(queue, ids); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.BatchNack(ctx, &BatchNackUpdate{
			QueueName: queue,
			Ids:       ids,
		})

		if err == nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to nack message on queue")

		}

		timeout.Increment()
		timeout.Sleep()

	}
}

func (c *syblineClient) Consumer(capacity int, timePeriod time.Duration, queue string) (*syblineConsumer, error) {
	if c.headers == nil {
		return nil, ErrMissingToken
	}

	return newSyblineConsumer(capacity, queue, timePeriod, c), nil
}

func (c *syblineClient) DeleteOueue(ctx context.Context, name string) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if len(name) == 0 {
		return ErrInvalidQueueName
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.DeleteQueue(ctx, &DeleteQueueInfo{
			QueueName: name,
		})

		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to delete queue")

		}

		timeout.Increment()
		timeout.Sleep()
	}
}

func (c *syblineClient) AddRoutingKey(ctx context.Context, queue, routingKey string) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateRoutingKeyChange(queue, routingKey); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.AddRoutingKey(ctx, &AddRoute{
			RouteName: routingKey,
			QueueName: queue,
		})

		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to add routing key")
		}

		timeout.Increment()
		timeout.Sleep()
	}

}

func (c *syblineClient) DeleteRoutingKey(ctx context.Context, queue, routingKey string) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateRoutingKeyChange(queue, routingKey); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.DeleteRoutingKey(ctx, &DeleteRoute{
			RouteName: routingKey,
			QueueName: queue,
		})

		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to add routing key")
		}

		timeout.Increment()
		timeout.Sleep()
	}
}

func (c *syblineClient) CreateUser(ctx context.Context, username, password string) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if err := validateUsernamePassword(username, password); err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	payload := &UserCreds{
		Username: username,
		Password: password,
	}
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.CreateUser(ctx, payload)

		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("unable to create user")
		}

		timeout.Increment()
		timeout.Sleep()

	}
}

func (c *syblineClient) DeleteUser(ctx context.Context, username string) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	if len(username) == 0 {
		return ErrInvalidUsername
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)
	payload := &UserInformation{
		Username: username,
	}
	timeout := NewTimeout(c.config.TimeoutAttempts, c.config.TimeoutSec)

	for {
		status, err := c.gClient.DeleteUser(ctx, payload)

		if err == nil && status != nil && status.Status {
			return nil
		}

		if err := c.handleErr(ctx, err); err != nil {
			return err
		}

		if timeout.HasThreadReached() {
			return fmt.Errorf("message was unable to delete user")
		}

		timeout.Increment()
		timeout.Sleep()
	}
}

func (c *syblineClient) Logout(ctx context.Context) error {
	if c.headers == nil {
		return ErrMissingToken
	}

	ctx = metadata.NewOutgoingContext(ctx, c.headers)

	_, err := c.gClient.LogOut(ctx, &LogOutRequest{})

	return err
}

func (c *syblineClient) Close() {
	c.conn.Close()
}

func (c *syblineClient) reconnect(ctx context.Context) error {
	// time for leadership election to finish
	time.Sleep(time.Second * 10)

	c.reconnectLock.Lock()
	defer c.reconnectLock.Unlock()

	// double check another goroutine has no re-establised connection
	status, err := c.gClient.IsLeaderNode(context.Background(), &LeaderNodeRequest{})
	if err == nil && status != nil && status.Status {
		return nil
	}

	for _, address := range c.addresses {
		conn, err := grpc.Dial(address, c.opts...)
		if err != nil {
			continue
		}

		gClient := NewMQEndpointsClient(conn)
		status, err := gClient.IsLeaderNode(context.Background(), &LeaderNodeRequest{})
		if err != nil || !status.Status {
			continue
		}

		c.conn = conn
		c.gClient = gClient

		if err := c.Login(ctx, c.username); err != nil {
			return err
		}

		return nil
	}

	return ErrNoValidAddresses
}

func (c *syblineClient) handleErr(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	if IsAuthErr(err) {
		return c.Login(ctx, c.username)
	}

	if IsNodeInvalid(err) {
		return c.reconnect(ctx)
	}

	return err
}

func IsNodeInvalid(err error) bool {
	errStr := err.Error()

	return strings.Contains(errStr, transportErrString) ||
		strings.Contains(errStr, leadersErrString) ||
		strings.Contains(errStr, serverErr)
}

func IsAuthErr(err error) bool {
	return strings.Contains(err.Error(), AuthErrString)
}

type Message struct {
	Rk   string
	Data []byte
}
