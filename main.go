package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
)

var PORT = 1080

func init() {
	port := os.Getenv("PORT")
	fmt.Printf("Check PORT ENV %s\n", port)
	if port != "" {
		p, err := strconv.Atoi(port)
		if err == nil {
			PORT = p
		} else {
			fmt.Println(err)
		}
	}
	fmt.Printf("PORT is %d\n", PORT)
}
func main() {
	go tcp()
	go udp()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case sig := <-c:
		fmt.Println("Close：", sig)
	}
}
