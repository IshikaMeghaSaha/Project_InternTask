package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type empData struct {
	Name string
	Age  string
	City string
}

func main() {

	csvFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	s := make([]empData, 0)
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		emp := empData{
			Name: line[0],
			Age:  line[1],
			City: line[2],
		}
		s = append(s, emp)
	}
	fmt.Println("NAME\tAGE\tCITY")
	for _, emp := range s {
		fmt.Println(emp.Name + "\t" + emp.Age + "\t" + emp.City)
	}
}
