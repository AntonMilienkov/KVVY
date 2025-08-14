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
type GenesisNode struct {
	Data string

	HashOwn string

	Timestamp int64
}

// генерация последующих событий
func SubGenesisNodeGen(gn *GenesisNode) *Node {
	var firstNode Node

	firstNode.Timestamp = time.Now().UnixMicro()
	firstNode.Data = time.Now().Format("02.01.2006 15:04:05")
	firstNode.HashSelfParent = gn.HashOwn
	firstNode.HashOtherParent = gn.HashOwn

	firstNode.HashOwn = getSHA256Hash(firstNode.Data + strconv.FormatInt(gn.Timestamp, 10) + firstNode.HashSelfParent + firstNode.HashOtherParent)

	return &firstNode
}

// генерация событий
func NodeGenerate(prevNode *Node) *Node {
	var newNode Node

	newNode.Timestamp = time.Now().UnixMicro()
	newNode.Data = time.Now().Format("02.01.2006 15:04:05")
	newNode.HashSelfParent = prevNode.HashOwn
	newNode.HashOtherParent = prevNode.HashOwn

	newNode.HashOwn = getSHA256Hash(newNode.Data + strconv.FormatInt(prevNode.Timestamp, 10) + newNode.HashSelfParent + newNode.HashOtherParent)

	return &newNode
}

// искуственная генерация событий
func ArtifNodeGenerate(gn *GenesisNode) error {
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
func (gn *GenesisNode) GenesisGenerate() bool {
	gn.Data = time.Now().Format("02.01.2006 15:04:05")
	gn.Timestamp = time.Now().UnixMicro()

	gn.HashOwn = getSHA256Hash(gn.Data + strconv.FormatInt(gn.Timestamp, 10))

	return files.WriteToFile(gn)
}
