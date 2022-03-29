package client

import (
	"gocamp/ch9/pkg/proto"
	"log"
	"net"
	"strings"
)

func Send()  {
	conn, err := net.Dial("tcp", ":9000")
	defer conn.Close()
	if err != nil {
		log.Printf("net dial error %v", err)
		return
	}
	proto := proto.New( 1)
	for i :=0; i<33;i++  {
		//str := fmt.Sprintf("[这是世界的一部分是不是]：%d", i)
		str := strings.Repeat("a",16)
		data, err := proto.Encode(str)
		if err != nil {
			log.Printf("proto encode error %v", err)
			return
		}
		//fmt.Printf("%d\n",data)
		conn.Write(data)
	}
	log.Println("send success and existed.")
}
