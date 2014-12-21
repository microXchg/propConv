package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"flag"
)

type entry struct {
	Name      string
	EMail     string
	Photo	  string
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
	Url   string
}

func recordToStruct(record []string) entry {

	fullsizeArray := make([]string, 14) //prevent index out of bounds
	copy(fullsizeArray, record)

	e := entry{}

	e.Name = fullsizeArray[0]
	e.EMail = fullsizeArray[1]
	e.Photo = fullsizeArray[2]
	e.Telephone = fullsizeArray[3]
	e.Language = fullsizeArray[4]
	e.Day = fullsizeArray[5]
	e.Track = fullsizeArray[6]
	e.Time = fullsizeArray[7]
	e.Status = fullsizeArray[8]
	e.Confirmed = fullsizeArray[9]
	e.Title = fullsizeArray[10]
	e.Abstract = fullsizeArray[11]
	e.Bio = fullsizeArray[12]
	e.Url = fullsizeArray[13]
	return e
}

func main() {

	var infilename  string
	flag.StringVar(&infilename,"i", "", "Csv file containing the input data.")
	var day string
	flag.StringVar(&day, "d", "1", "Day for which schedule should be generated.")
	var track string
	flag.StringVar(&track, "t", "1", "Track for which schedule should be generated.")

	flag.Parse()

	input, err := os.Open(infilename)

	if err != nil {
		fmt.Println("Error opening inputfile")
	}

	defer input.Close()

	csvReader := csv.NewReader(input)

	//drop first line
	csvReader.Read()
	content, _ := csvReader.ReadAll()

	t, _ := template.ParseFiles("schedule-template.ctmpl")

	fmt.Printf("<!-- START day %s track %s -->\n", day, track)
	for _, line := range content {
		record := recordToStruct(line)
		if record.Day == day && record.Track == track {
			t.Execute(os.Stdout, record)
		}
	}
	fmt.Printf("<!-- END day %s track %s -->\n", day, track)

}
