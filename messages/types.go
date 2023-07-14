package messages

import (
	"sync"

	"google.golang.org/protobuf/runtime/protoimpl"
)

var (
	file_mq_proto_rawDescOnce sync.Once
	file_mq_proto_rawDescData = file_mq_proto_rawDesc
)

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
	(*BatchNackUpdate)(nil),    // 16: MQ.BatchNackUpdate
}

func file_mq_proto_rawDescGZIP() []byte {
	file_mq_proto_rawDescOnce.Do(func() {
		file_mq_proto_rawDescData = protoimpl.X.CompressGZIP(file_mq_proto_rawDescData)
	})
	return file_mq_proto_rawDescData
}
