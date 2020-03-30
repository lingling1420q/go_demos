package client

import (
	"log"
	"net"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func TestEpoll_Client(t *testing.T) {
	//conn, err := net.Dial("tcp", "172.16.4.136:8080")
	conn, err := net.Dial("tcp", "172.16.92.132:8080")
	if err != nil {
		log.Println("Error dialing", err.Error())
		return // 终止程序
	}

	n, err := conn.Write([]byte("hello"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("write:", n)

	go func() {
		for {
			var bytes = make([]byte, 100)
			n, err := conn.Read(bytes)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("read:", string(bytes[0:n]))
		}
	}()

	select {}
}