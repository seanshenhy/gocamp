package server

import (
	"fmt"
	"gocamp/ch9/pkg/proto"
	"gocamp/ch9/pkg/reader"
	"log"
	"net"
)

type Server struct {
	network string
	addr string
	proto *proto.Proto
	listener net.Listener
}

func New(network, addr string) *Server {
	return &Server{
		network: network,
		addr:    addr,
		proto:   proto.New(1),
	}
}

func (s *Server)handler(conn net.Conn)  {
	defer conn.Close()
	//创建接收数据的通道
	acceptData := make(chan string, 10)
	defer close(acceptData)
	go s.accept(acceptData)

	reader := reader.New(conn)
	err := reader.Read(acceptData)
	if err != nil {
		fmt.Println("read data error:", err)
	}
	//buffer := make([]byte, 1024)
	//for  {
	//	n, err := conn.Read(buffer)
	//	fmt.Println(n,err)
	//	if err == io.EOF {
	//		break
	//	}
	//	time.Sleep(time.Second)
	//}
}
func (s *Server)accept(acceptData chan string)  {
	for {
		value, isOk := <- acceptData
		if !isOk {
			break
		}
		v,_ := s.proto.Decode([]byte(value))
		fmt.Println("接收数据：",v)
	}
}
func (s *Server)Listen()  {
	listener, err := net.Listen(s.network, s.addr)
	if err != nil {
		log.Printf("listen error %+v", err)
	}
	s.listener = listener
}

func (s *Server)Run()  {
	//go client.Send()
	log.Printf("%v server start %s", s.network, s.addr)
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("connection accept error %+v", err)
			continue
		}
		go s.handler(conn)
	}
}

