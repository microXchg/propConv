package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"html/template"
	"log"
)

type entry struct {
	Name      string
	EMail     string
	Photo	  string
	Twitter   string
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

	fullsizeArray := make([]string, 15) //prevent index out of bounds
	copy(fullsizeArray, record)

	e := entry{}

	e.Name = fullsizeArray[0]
	e.EMail = fullsizeArray[1]
	e.Photo = fullsizeArray[2]
	e.Twitter = fullsizeArray[3]
	e.Telephone = fullsizeArray[4]
	e.Language = fullsizeArray[5]
	e.Day = fullsizeArray[6]
	e.Track = fullsizeArray[7]
	e.Time = fullsizeArray[8]
	e.Status = fullsizeArray[9]
	e.Confirmed = fullsizeArray[10]
	e.Title = fullsizeArray[11]
	e.Abstract = fullsizeArray[12]
	e.Bio = fullsizeArray[13]
	e.Url = fullsizeArray[14]
	return e
}

func main() {

	var infilename  string
	flag.StringVar(&infilename,"i", "", "Csv file containing the input data.")
	var day string
	flag.StringVar(&day, "d", "1", "Day for which schedule should be generated.")
	var track string
	flag.StringVar(&track, "t", "1", "Track for which schedule should be generated.")
	var detail bool
	flag.BoolVar(&detail, "detail", false, "Set to generate the single pages instead of schedule-html. Day and Track will be ignored.")
	var outdir string
	flag.StringVar(&outdir, "outdir", "talk", "Locaction where the generated single pages should be saved.")


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

	if detail {
		writeDetailPages(outdir, content)
	} else {
		writeScheduleHtml(day, track, content)
	}
}

func writeScheduleHtml(day string, track string, content [][]string) {
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

func writeDetailPages(outdir string, content [][]string) {
	t, _ := template.ParseFiles("singlepage-template.ctmpl")

	for _, line := range content {
		record := recordToStruct(line)
		filename := outdir + "/" + record.Url + ".html"
		outfile, err := os.Create(filename)

		if err!=nil {
			log.Fatalf("Could not create Outputfile with name: %s", filename)
		}

		defer outfile.Close()

		t.Execute(outfile, record)
	}
}
