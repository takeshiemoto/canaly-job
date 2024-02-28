package main

import (
	"encoding/csv"
	"fmt"
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
	record, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading the CSV", err)
		return
	}
	fmt.Println(record)
}
