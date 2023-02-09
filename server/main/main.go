package main

import (
	"fmt"
	"gotest/chat/client/model"
	"gotest/chat/common"
	"net"
)

func co(conn net.Conn) {

	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}
	err := processor.run()
	if err != nil {
		fmt.Println("客户端与服务端协程错误err=", err)
	}

}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(common.Client)
}

func main() {
	common.NewAutoInc(1, 1)
	common.InitC()
	initUserDao()

	listen, err := net.Listen("tcp", "127.0.0.1:8899")
	defer listen.Close()
	if err != nil {
		fmt.Println("服务器错误", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务器接收错误", err)
		}

		go co(conn)
	}

	//defer autoInc.Close()
}
