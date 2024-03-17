package main

import (
	"encoding/csv"
	"fmt"
	"github.com/takeshiemoto/canaly-job/leran"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data/example.csv")
	if err != nil {
		fmt.Println("Error opened the CSV", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		fmt.Println("Error reading the header row:", err)
	}
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading a row:", err)
			return
		}
		fmt.Println(record[0])
		fmt.Println(record[4])
		fmt.Println(record[5])
	}

	leran.Run()
}
