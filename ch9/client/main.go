package main

import (
	"fmt"
	"gocamp/ch9/pkg/proto"
	"log"
	"net"
)

func main()  {
	conn, err := net.Dial("tcp", ":9000")
	defer conn.Close()
	if err != nil {
		log.Printf("net dial error %v", err)
		return
	}
	proto := proto.New( 1)
	for i :=0; i<1000 ;i++  {
		str := fmt.Sprintf("[这是世界的一部分是不是]：%d", i)
		//str := "ab"
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

