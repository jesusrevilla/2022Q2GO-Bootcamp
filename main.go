package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type data struct {
	Index int64
	Item  string
}

func readFile(fileName string) []data {
	items := make([]data, 0)

	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println()
	}

	for _, line := range csvLines {
		index, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		item := data{
			Index: index,
			Item:  line[1],
		}
		items = append(items, item)
	}

	return items
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	rows := readFile("csvfile.csv")
	message := "<h1>File content: </h1>"
	message += "<ul>"
	for _, data := range rows {
		message += fmt.Sprint("<li>Index: ", data.Index, " Item: ", data.Item, "</li>")
	}
	message += "</ul>"
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal()
	}
}

func main() {

	http.HandleFunc("/readcsvfile", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}
