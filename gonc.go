package main

import (
	"bufio"
	"fmt"
	"os"
	"net"
	"io"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage: "+os.Args[0]+"<IP|-l> <PORT>")
		os.Exit(3)
	}

	switch os.Args[1] {
	
	case "-l":
		reader := bufio.NewReader(os.Stdin)
		
		addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:" + os.Args[2])
		listener, _ := net.ListenTCP("tcp", addr)
		conn, _ := listener.AcceptTCP()	
		go io.Copy(conn, reader)
		io.Copy(os.Stdout, conn)
		
		
	default:
		reader := bufio.NewReader(os.Stdin)
	
		conn, err := net.Dial("tcp",os.Args[1] + ":" + os.Args[2])
		if nil != err {
			if nil != conn {
				conn.Close()
			}
			fmt.Println(err)
		}
		go io.Copy(conn, reader)
		io.Copy(os.Stdout, conn)
	}
}
