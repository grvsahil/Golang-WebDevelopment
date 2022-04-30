package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main()  {
	li,err := net.Listen("tcp",":5500")
	if err!=nil{
		log.Fatalln(err)
	}
	defer li.Close()

	for{
		conn,err:=li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn,"Connection established using TCP")
		fmt.Fprintf(conn,"\nGaurav Sahil\n")

		conn.Close()
	}
}