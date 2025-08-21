package files

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO  проверка поступаемых данных
func WriteToFile(gn interface{}) error {
	filename := "graph/hashgraph.json"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("Error while opening file")
		return err
	}
	defer file.Close()

	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}

	size := fileInfo.Size()

	if fileInfo.Size() == 0 {
		file.WriteString("[")
	} else {
		err = os.Truncate(filename, size-1)
		if err != nil {
			fmt.Println("ошибка при усечении файла: %w", err)
		}

		file.WriteString(",")
	}

	jsonData, err := json.Marshal(&gn)
	if err != nil {
		fmt.Println("Ошибка сериализации в JSON:", err)
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return err
	}

	file.WriteString("]")

	return err
}
