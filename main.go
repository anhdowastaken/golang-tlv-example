package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net"

	"github.com/anhdowastaken/golang-tlv-example/tlv"
)

func main() {

	network := "tcp"
	address := "127.0.0.1:8081"
	ln, err := net.Listen(network, address)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Printf("Listening %s://%s...", network, address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("%v", err)
			continue
		}

		go read(&conn)
	}
}

func simpleTest() {
	buf := new(bytes.Buffer)
	codec := &tlv.Codec{TypeBytes: tlv.TwoBytes, LenBytes: tlv.TwoBytes}
	wr := tlv.NewWriter(buf, codec)
	record := &tlv.Record{
		Payload: []byte("hello, go!"),
		Type:    8,
	}
	wr.Write(record)
	log.Println(hex.Dump(buf.Bytes()))

	reader := bytes.NewReader(buf.Bytes())
	tlvReader := tlv.NewReader(reader, codec)
	next, _ := tlvReader.Next()
	log.Println("type: ", next.Type)
	log.Println("payload: ", string(next.Payload))
}

func read(conn *net.Conn) {
	defer (*conn).Close()

	buf, _ := ioutil.ReadAll(bufio.NewReader(*conn))
	codec := &tlv.Codec{TypeBytes: tlv.TwoBytes, LenBytes: tlv.TwoBytes}
	reader := bytes.NewReader(buf)
	tlvReader := tlv.NewReader(reader, codec)
	next, err := tlvReader.Next()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("type: %d", next.Type)
	log.Printf("payload: %s", string(next.Payload))
}
