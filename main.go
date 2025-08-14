package main

import (
	"fmt"
	"nauchka/graph"
)

func main() {
	var gn graph.GenesisNode

	fmt.Println(gn.GenesisGenerate())
	fmt.Println(gn.HashOwn)
	fmt.Println(gn.Timestamp)

	//var n1 node
	fmt.Println(graph.ArtifNodeGenerate(&gn))
	//fmt.Println(gn.HashOwn)
	//fmt.Println(gn.Timestamp)
}
