package files

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO  проверка поступаемых данных
func WriteToFile(gn interface{}) bool {
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
