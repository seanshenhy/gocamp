package reader

import (
	"encoding/binary"
	"fmt"
	"gocamp/ch9/pkg/proto"
	"io"
	"net"
)

type Reader struct {
	Conn net.Conn
	Buff []byte
	Start int    //数据读取开始位置
	End int    //数据读取结束位置
	BuffLen int    //数据接收缓冲区大小
	HeaderLen int    //包头长度
}

func New(conn net.Conn) *Reader{
	return &Reader{
		Conn:  conn,
		Buff:  make([]byte, proto.MaxBufferSize),
		Start: 0,
		End:   0,
		BuffLen: proto.MaxBufferSize,
		HeaderLen: proto.RawPacketLen,
	}
}
func (r *Reader)Read(acceptData chan string) error {
	for {
		r.Move()
		if r.End == r.BuffLen {
			return fmt.Errorf("one message is too large:%v", r)
		}
		len, err := r.Conn.Read(r.Buff[r.End:])
		if err == io.EOF {
			continue
		} else if err != nil {
			return err
		}
		r.End += len
		r.GetData(acceptData)
	}
}

func (r *Reader)Move() {
	if r.Start == 0 {
		return
	}
	copy(r.Buff, r.Buff[r.Start:r.End])
	r.End -= r.Start
	r.Start = 0
}

func (r *Reader)GetData(acceptData chan string) error {
	//包头的长度不够，继续接收
	if r.End - r.Start < r.HeaderLen {
		return nil
	}


	headerData := r.Buff[r.Start:(r.Start + r.HeaderLen)]
	packetLen := binary.BigEndian.Uint32(headerData[proto.PacketOffset:proto.HeaderOffset])
	bodyLen := packetLen - proto.RawPacketLen
	//包体的长度不够，继续接收
	if r.End - r.Start - r.HeaderLen < int(bodyLen) {
		return nil
	}

	//读取包体数据
	sp := r.Start + r.HeaderLen
	ep := sp + int(bodyLen)
	bodyData := r.Buff[sp:ep]

	//把完整的包用通道传递出去
	acceptData <- string(append(headerData, bodyData...))

	//每读完一次数据 start 后移读取位置
	r.Start += r.HeaderLen + int(bodyLen)
	return r.GetData(acceptData)
}
