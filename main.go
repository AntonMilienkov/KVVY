package main

import (
	"bufio"
	"fmt"
	"log"
	"nauchka/graph"
	"net"
	"net/rpc"
	"os"
)

type QuotientServer struct {
	node graph.Node
}

type ServerMethods int

func (rcv *ServerMethods) WriteData(args QuotientServer, result *bool) error {
	return nil
}

func serverRole() {
	rpc.Register(new(ServerMethods))

	l, e := net.Listen("tcp", ":17152")

	if e != nil {
		log.Fatal("listen error:", e)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		fmt.Printf("request from %v\n", c.RemoteAddr())
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func clientRole() {
	// to where?
	client, err := rpc.Dial("tcp", "")

	if err != nil {
		log.Fatal("Не удалось подключиться к серверу:", err)
	}
	defer client.Close()

	var gn graph.GenesisNode

	fmt.Println(gn.GenesisGenerate())
	fmt.Println(gn.HashOwn)
	fmt.Println(gn.Timestamp)

	//var n1 node
	// fmt.Println(graph.ArtifNodeGenerate(&gn))
	//fmt.Println(gn.HashOwn)
	//fmt.Println(gn.Timestamp)
}

func main() {
	//go serverRole()

	//time.Sleep(8 * time.Second)

	//clientRole()

	var gn graph.GenesisNode
	fmt.Println(gn.GenesisGenerate())
	fmt.Println(gn.HashOwn)
	fmt.Println(gn.Timestamp)

}

func generateDataBlocks() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	switch text {
	case "Exit\n":
		os.Exit(0)
	default:
		fmt.Println(text)
		generateDataBlocks()
	}
}
