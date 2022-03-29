package proto

import (
	"encoding/binary"
)


/*
goim 协议结构
总共16字节
4bytes PacketLen 包长度，在数据流传输过程中，先写入整个包的长度，方便整个包的数据读取。
2bytes HeaderLen 头长度，在处理数据时，会先解析头部，可以知道具体业务操作。
2bytes Version 协议版本号，主要用于上行和下行数据包按版本号进行解析。
4bytes Operation 业务操作码，可以按操作码进行分发数据包到具体业务当中。
4bytes Sequence 序列号，数据包的唯一标记，可以做具体业务处理，或者数据包去重。
PacketLen-HeaderLen Body 实际业务数据，在业务层中会进行数据解码和编码。
*/
const (
	// 包字节数
	packetLen 	= 4
	// 头节数长度
	headerLen 	= 2
	// 版本节数长度
	versionLen 	= 2
	// 操作类型节数长度
	operationLen 	= 4
	// 序列号长度
	sequenceLen 	= 4
	// 包总长度
	RawPacketLen = packetLen + headerLen + versionLen + operationLen + sequenceLen

	// 偏移位置
	PacketOffset = 0
	HeaderOffset = PacketOffset + packetLen
	VersionOffset = HeaderOffset + headerLen
	OperationOffset = VersionOffset + versionLen
	SequenceOffset = OperationOffset + operationLen

	MaxBufferSize = int(1 << 8)
)
type Proto struct {
	Version uint16
	Operation uint32
	Sequence uint32
	Body interface{}
}

func New(version int) *Proto{
	return &Proto{
		Version: uint16(version),
	}
}

// Encode 将消息编码
func (p *Proto)Encode(message string) ([]byte, error) {
	// 包长度
	packetLen := len([]byte(message)) + RawPacketLen
	// 创建缓冲区
	wb := make([]byte, packetLen)
	// 写入包头信息
	binary.BigEndian.PutUint32(wb[PacketOffset:HeaderOffset], uint32(packetLen))
	// 写入header
	binary.BigEndian.PutUint16(wb[HeaderOffset:VersionOffset], RawPacketLen)
	// 写入版本
	binary.BigEndian.PutUint16(wb[VersionOffset:OperationOffset], versionLen)
	// 写入操作类型
	binary.BigEndian.PutUint32(wb[OperationOffset:SequenceOffset], operationLen)
	// 写入序列号
	binary.BigEndian.PutUint32(wb[SequenceOffset:RawPacketLen], sequenceLen)
	// 写入body
	byteBody := []byte(message)
	copy(wb[RawPacketLen:], byteBody)
	return wb, nil
}

// Encode 将消息编码
func (p *Proto)Decode(data []byte) (message string, err error) {
	//包头转为string
	//head := string(data[:RawPacketLen])

	//将包体转为string
	return string(data[RawPacketLen:]), nil
}