package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type node struct {
	data string

	HashSelfParent  string
	HashOtherParent string
	HashOwn         string

	Timestamp int64
}
type genesisNode struct {
	Data string

	HashOwn string

	Timestamp int64
}

func (gn *genesisNode) writeToFile() bool {
	file, err := os.OpenFile("hashgraph.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("Error while opening file")
		return false
	}
	defer file.Close()

	fileInfo, err := os.Stat("hashgraph.json")
	if err != nil {
		return false
	}

	size := fileInfo.Size()
	fmt.Println(size)

	if fileInfo.Size() == 0 {
		file.WriteString("[")
	} else {
		err = os.Truncate("hashgraph.json", size-1)
		if err != nil {
			fmt.Println("ошибка при усечении файла: %w", err)
		}

		file.WriteString(",")
	}

	jsonData, err := json.Marshal(&gn)
	if err != nil {
		fmt.Println("Ошибка сериализации в JSON:", err)
		return false
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return false
	}

	file.WriteString("]")

	return true
}

func (gn *genesisNode) genesisGenerate() bool {
	gn.Data = "hello world"
	gn.Timestamp = time.Now().UnixMicro()

	gn.HashOwn = getSHA256Hash(gn.Data + strconv.FormatInt(gn.Timestamp, 10))

	return gn.writeToFile()
}

func main() {
	var gn genesisNode

	fmt.Println(gn.genesisGenerate())
	fmt.Println(gn.HashOwn)
	fmt.Println(gn.Timestamp)
}

func getSHA256Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))

	return hex.EncodeToString(hasher.Sum(nil))
}
