package server

import (
	"bufio"
	"fmt"
	"gocamp/ch9/pkg/proto"
	"io"
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
	reader := bufio.NewReader(conn)
	for  {
		msg, err := s.proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		if msg == "" {
			fmt.Printf("recive clientn error：%v\n", err)
			return
		}
		fmt.Printf("收到client发来的数：%v\n", msg)
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

