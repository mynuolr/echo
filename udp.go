package main

import (
	"fmt"
	"io"
	"net"
)

func udp() {
	ladd, err := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:%d", PORT))
	if err != nil {
		panic(err)
	}
	list, err := net.ListenUDP("udp", ladd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listen UDP :%s\n", ladd)
	buf:=make([]byte,4096)
	for {
		t,uadd, e := list.ReadFromUDP(buf)
		if e != nil {
			if e != io.EOF {
				fmt.Printf("RemoteAddr:%s,Read Err:%e\n", uadd.String(), e)
				continue
			}
		}
		fmt.Printf("RemoteAddr:%s,Data:%#v\n", uadd.String(), buf[:t])
		_,e=list.WriteToUDP(append([]byte("sd"),buf[:t]...),uadd)
		if e != nil {
			fmt.Printf("RemoteAddr:%s,Writer Err:%e\n", uadd.String(), e)
			continue
		}
	}
}
