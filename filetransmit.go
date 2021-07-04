package main

import (
 	"fmt"
	"os"
	"bufio"
	"net"
	"io"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("usage: "+os.Args[0]+" [-s|-r|-ls|-lr] <IP:PORT|PORT> <file>")
		os.Exit(3)
	}

	switch os.Args[1] {
	
	case "-ls":
		fi, _ := os.Open(os.Args[3]) 
		read := bufio.NewReader(fi)
		
		addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:" + os.Args[2])
		listener, _ := net.ListenTCP("tcp", addr)
		conn, _ := listener.AcceptTCP()	
		io.Copy(conn, read)
	
	case "-lr":
		fo, _ := os.Create(os.Args[3])
			
		addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:" + os.Args[2])
		listener, _ := net.ListenTCP("tcp", addr)
		conn, _ := listener.AcceptTCP()	
		io.Copy(fo, conn)
	
	case "-s":
		fi, _ := os.Open(os.Args[3]) 
		read := bufio.NewReader(fi)
	
		conn, err := net.Dial("tcp", os.Args[2])
		if nil != err {
			if nil != conn {
				conn.Close()
			}
			fmt.Println(err)
		}
		io.Copy(conn, read)
		
	case "-r":	
		fo, _ := os.Create(os.Args[3]) 
		
		conn, err := net.Dial("tcp", os.Args[2])
		if nil != err {
			if nil != conn {
				conn.Close()
			}
			fmt.Println(err)
		}
		io.Copy(fo, conn)
		
	}	
	
}

