package main

import (
	"fmt"
	"io"
	"net"
)

func tcp() {
	tadd, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%d", PORT))
	if err != nil {
		panic(err)
	}
	list, err := net.ListenTCP("tcp", tadd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listen TCP :%s\n", tadd)
	for {
		t, e := list.AcceptTCP()
		if err != nil {
			fmt.Println(e)
			continue
		}
		go tcpHandler(t)
	}
}
func tcpHandler(conn *net.TCPConn) {
	buf := make([]byte, 4096)
	conn.Write([]byte(conn.RemoteAddr().String()))
	var i = 0
	var e error
	for {
		i, e = conn.Read(buf)
		conn.LocalAddr()
		if e != nil {
			if e != io.EOF {
				fmt.Printf("LocalAddr:%s, RemoteAddr:%s,Read Err:%e\n", conn.LocalAddr(), conn.RemoteAddr(), e)
				e = conn.Close()
				if e != nil {
					fmt.Printf("LocalAddr:%s, RemoteAddr:%s,Close Err:%e\n", conn.LocalAddr(), conn.RemoteAddr(), e)
				}
				return
			}
		}

		fmt.Printf("LocalAddr:%s, RemoteAddr:%s,Data:%#v\n", conn.LocalAddr(), conn.RemoteAddr(), buf[:i])
		_, e = conn.Write(buf[:i])
		if e != nil {
			fmt.Printf("LocalAddr:%s, RemoteAddr:%s,Writer Err:%e\n", conn.LocalAddr(), conn.RemoteAddr(), e)
		}
	}
}
