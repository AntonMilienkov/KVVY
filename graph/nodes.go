package graph

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"nauchka/files"
	"os"
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

func GetGenesis() *Node {
	filename := "graph/hashgraph.json"
	file, _ := os.ReadFile(filename)

	var data []Node
	json.Unmarshal(file, &data)

	return &data[0]
}

// генерация событий
// add hashOtherParent
func NodeGenerate(prevNode *Node, hashOtherParent string) *Node {
	var newNode Node

	newNode.Timestamp = time.Now().UnixMicro()
	newNode.Data = time.Now().Format("02.01.2006 15:04:05")
	newNode.HashSelfParent = prevNode.HashOwn
	//here
	newNode.HashOtherParent = hashOtherParent

	newNode.HashOwn = getSHA256Hash(newNode.Data + strconv.FormatInt(prevNode.Timestamp, 10) + newNode.HashSelfParent + newNode.HashOtherParent)

	return &newNode
}

type Host struct {
	DnsName string
	Port    int
}

var hosts []Host

var count, selfNumber int

func fillHosts(hosts []Host) {
	for i := range hosts {
		hosts[i] = Host{
			DnsName: "node" + strconv.Itoa(i+1),
			Port:    17152,
		}
	}
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Has(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func getRandomNodeNumber(exceptNodes Set[int]) (randomNodeNumber int) {
	randomNodeNumber = selfNumber

	for exceptNodes.Has(randomNodeNumber) {
		randomNodeNumber = rand.Intn(count)
	}

	exceptNodes.Add(randomNodeNumber)

	return
}

func Init() {
	// количество узлов
	count, _ = strconv.Atoi(os.Getenv("COUNT"))

	hosts = make([]Host, count)

	// заполнить массив
	fillHosts(hosts)

	selfNumber, _ = strconv.Atoi(os.Getenv("SELFNUMBER"))
}

func sendNode(node *Node) error {

	exceptNodes := make(Set[int])

	exceptNodes.Add(selfNumber)

	randomNodeNumber := getRandomNodeNumber(exceptNodes)

	client(hosts[randomNodeNumber], node)

	// TODO: повторить

	return nil
}

func syncWriteToFiles(node *Node) error {
	err := files.WriteToFile(node)

	if err != nil {
		fmt.Println("Ошибка при записи ноды в файл: %w", err)
	}

	err = sendNode(node)

	if err != nil {
		fmt.Println("Ошибка при отправке ноды другим серверам: %w", err)
	}

	return nil
}

// TODO:
func getOtherParentNode() *Node {
	return &Node{Data: "otherParentNode", HashSelfParent: "", HashOtherParent: "", HashOwn: "", Timestamp: 12345678}
}

// искусcтвенная генерация событий
func ArtifNodeGenerate(genesisNode *Node) error {
	firstNode := NodeGenerate(genesisNode, "")

	syncWriteToFiles(firstNode)
	/*
		coef := 1
		end := time.Now().Add(time.Duration(coef) * time.Minute)

		nextNode := firstNode
		for time.Now().Before(end) {
			secToSleep := coef*5 + rand.Intn(20)
			time.Sleep(time.Duration(secToSleep) * time.Second)

			otherParentNode := getOtherParentNode()
			nextNode = NodeGenerate(nextNode, otherParentNode.HashOwn)

			files.WriteToFile(&nextNode)
		}
	*/
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
