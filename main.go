package main

import (
	"bufio"
	"fmt"
	"nauchka/graph"
	"os"
	"time"
)

func main() {
	go graph.ServerStart()

	time.Sleep(8 * time.Second)

	genesisNode := graph.GetGenesis()

	graph.ArtifNodeGenerate(genesisNode)
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
