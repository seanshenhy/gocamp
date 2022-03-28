package main

import "gocamp/ch9/server/server"

func main()  {
	srv := server.New("tcp", ":9000")
	srv.Listen()
	srv.Run()
}

