package main

import (
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	c, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		log.Println(err)
		return
	}

	if _, err := c.Write([]byte("GET / HTTP/1.1\r\nHost: localhost\r\n\r\n")); err != nil {
		log.Println(err)
		return
	}

	b := make([]byte, 40)

	if _, err := c.Read(b); err != nil {
		log.Println(err)
		return
	}

	log.Println(string(b))

	if _, err := c.Write([]byte("my name is gorilla\r\n")); err != nil {
		log.Println(err)
		return
	}

	b = make([]byte, 1024)

	if _, err := c.Read(b); err != nil {
		log.Println(err)
		return
	}

	log.Println(string(b))
}
