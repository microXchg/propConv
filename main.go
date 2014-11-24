package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
)

type entry struct {
	Name      string
	EMail     string
	Telephone string
	Language  string
	Day       string
	Track     string
	Time      string
	Status    string
	Confirmed string
	Title     string
	Abstract  string
	Bio       string
	Remarks   string
}

func recordToStruct(record []string) entry {

	fullsizeArray := make([]string, 13) //prevent index out of bounds
	copy(fullsizeArray, record)

	e := entry{}

	e.Name = fullsizeArray[0]
	e.EMail = fullsizeArray[1]
	e.Telephone = fullsizeArray[2]
	e.Language = fullsizeArray[3]
	e.Day = fullsizeArray[4]
	e.Track = fullsizeArray[5]
	e.Time = fullsizeArray[6]
	e.Status = fullsizeArray[7]
	e.Confirmed = fullsizeArray[8]
	e.Title = fullsizeArray[9]
	e.Abstract = fullsizeArray[10]
	e.Bio = fullsizeArray[11]
	e.Remarks = fullsizeArray[12]
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

	t, _ := template.ParseFiles("schedule-template.ctmpl")

	fmt.Printf("<!-- START day %s track %s -->\n", os.Args[2], os.Args[3])
	for _, line := range content {
		record := recordToStruct(line)
		if record.Day == os.Args[2] && record.Track == os.Args[3] {
			t.Execute(os.Stdout, record)
		}
	}
	fmt.Printf("<!-- END day %s track %s -->\n", os.Args[2], os.Args[3])

}
