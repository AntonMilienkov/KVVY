package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type node struct {
	data string

	hashSelfParent  string
	hashOtherParent string
	hashOwn         string

	timestamp int64
}
type genesisNode struct {
	data string

	hashOwn string

	timestamp int64
}

func (gn *genesisNode) genesisGenerate() bool {
	gn.data = "hello world"
	gn.timestamp = time.Now().UnixMicro()

	gn.hashOwn = getSHA256Hash(gn.data)

	return true
}

func main() {
	var gn genesisNode

	fmt.Println(gn.genesisGenerate())
	fmt.Println(gn.hashOwn)
	fmt.Println(gn.timestamp)
}

func getSHA256Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))

	return hex.EncodeToString(hasher.Sum(nil))
}
