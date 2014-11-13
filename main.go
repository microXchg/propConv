package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/template"
)

type entry struct {
	Timestamp string
	Name      string
	EMail     string
	Telephone string
	Language  string
	Abstract  string
	Bio       string
}

func recordToStruct(record []string) entry {

	fullsizeArray := make([]string, 7) //prevent index out of bounds
	copy(fullsizeArray, record)

	e := entry{}

	e.Timestamp = fullsizeArray[0]
	e.Name = fullsizeArray[1]
	e.EMail = fullsizeArray[2]
	e.Telephone = fullsizeArray[3]
	e.Language = fullsizeArray[4]
	e.Abstract = fullsizeArray[5]
	e.Bio = fullsizeArray[6]
	return e
}

func main() {

	input, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error opening inputfile")
	}

	defer input.Close()

	csvReader := csv.NewReader(input)

	//drop first line
	csvReader.Read()
	content, _ := csvReader.ReadAll()

	t, _ := template.ParseFiles("md-template.ctmpl")

	for _, line := range content {
		record := recordToStruct(line)
		t.Execute(os.Stdout, record)
	}

}
