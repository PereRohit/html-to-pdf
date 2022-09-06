package main

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Usage: go run main.go <URL> <output>.pdf")
		os.Exit(1)
	}
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatalln("New instance gen error", err)
	}

	page := wkhtmltopdf.NewPage(os.Args[1])
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatalln("Creating error", err)
	}

	err = pdfg.WriteFile("./files/" + os.Args[2])
	if err != nil {
		log.Fatalln("Write to file error", err)
	}
	log.Println("Done")
}
