package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct{}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go fs.read_loop(conn)
	}
}

func (fs *FileServer) read_loop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		// n, err := io.Copy(buf, conn) //会因为没读到EOF而一直等待
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}

		// n, err := conn.Read(buf)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// file := buf[:n]
		// fmt.Println(file)

		fmt.Println(buf.Bytes())
		fmt.Printf("received %d bytes over the network\n", n)
	}
}

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	binary.Write(conn, binary.LittleEndian, int64(size))         //把文件大小写入到流中
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size)) //防止因为内存限制文件流被分割
	if err != nil {
		return err
	}

	// n, err := conn.Write(file)
	// if err != nil {
	// 	return err
	// }

	fmt.Printf("written %d bytes over ther network\n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		sendFile(200000)
	}()

	server := &FileServer{}
	server.start()
}
