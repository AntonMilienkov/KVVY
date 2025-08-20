package graph

import (
	"math/rand"
	"nauchka/files"
	"strconv"
	"time"
)

type Node struct {
	Data string

	HashSelfParent  string
	HashOtherParent string
	HashOwn         string

	Timestamp int64
}

/*
type GenesisNode struct {
	Data string

	HashOwn string

	Timestamp int64
}*/

// генерация последующих событий
func SubGenesisNodeGen(genesisNode *Node) *Node {
	var firstNode Node

	firstNode.Timestamp = time.Now().UnixMicro()
	firstNode.Data = time.Now().Format("02.01.2006 15:04:05")
	firstNode.HashSelfParent = genesisNode.HashOwn
	firstNode.HashOtherParent = genesisNode.HashOwn

	firstNode.HashOwn = getSHA256Hash(firstNode.Data + strconv.FormatInt(gn.Timestamp, 10) + firstNode.HashSelfParent + firstNode.HashOtherParent)

	return &firstNode
}

// генерация событий
// add hashOtherParent
func NodeGenerate(prevNode *Node) *Node {
	var newNode Node

	newNode.Timestamp = time.Now().UnixMicro()
	newNode.Data = time.Now().Format("02.01.2006 15:04:05")
	newNode.HashSelfParent = prevNode.HashOwn
	//here
	newNode.HashOtherParent = prevNode.HashOwn

	newNode.HashOwn = getSHA256Hash(newNode.Data + strconv.FormatInt(prevNode.Timestamp, 10) + newNode.HashSelfParent + newNode.HashOtherParent)

	return &newNode
}

// func syncWriteToFiles() error {
// 	return err
// }

// искуственная генерация событий
func ArtifNodeGenerate(genesisNode *GenesisNode) error {
	firstNode := SubGenesisNodeGen(gn)
	files.WriteToFile(&firstNode)

	rand.Seed(time.Now().UnixNano())
	coef := 1
	end := time.Now().Add(time.Duration(coef) * time.Minute)

	nextNode := firstNode
	for time.Now().Before(end) {
		secToSleep := coef*5 + rand.Intn(20)
		time.Sleep(time.Duration(secToSleep) * time.Second)

		nextNode = NodeGenerate(nextNode)

		files.WriteToFile(&nextNode)
	}

	return nil
}

// генерация генезис события
/*
func (n *Node) GenesisGenerate() bool {
	n.Data = "Hello Graph"
	n.Timestamp = time.Now().UnixMicro()

	n.HashSelfParent = getSHA256Hash(n.Data + strconv.FormatInt(n.Timestamp, 10))
	n.HashOtherParent = n.HashSelfParent

	return files.WriteToFile(n)
}
*/
