package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"gomail/response"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}
	_, _ = conn.Write([]byte("1262193323@qq.com:kwjklcboqznsbabc"))
	rec := make([]byte, 1000)
	for {
		_, err = conn.Read(rec)
		if err != nil {
			log.Println(err)
			break
		}
		mail := &response.Mail{}
		err = proto.Unmarshal(rec, mail)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf(" get mail: %+v \n", mail)
	}

	_ = conn.Close()
}
