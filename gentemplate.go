package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	genFromTemplate()
}

func genFromTemplate() {
	funcs := template.FuncMap{
		"AddFloats": func(x, y float64) float64 {
			return x + y
		},
		"AddInts": func(x, y int) int {
			return x + y
		},
	}
	tmpl, err := template.New("template.html").Funcs(funcs).ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}

	type Cart struct {
		Item   string
		Amount float64
	}
	d := map[string]any{
		"Items": []Cart{
			{
				Item:   "Bread",
				Amount: 24,
			},
			{
				Item:   "Rice",
				Amount: 56.7,
			},
			{
				Item:   "Clothes",
				Amount: 150.45,
			},
			{
				Item:   "Water",
				Amount: 100,
			},
			{
				Item:   "Gas",
				Amount: 100.00,
			},
		},
		"Title": "Inventory list",
	}
	f, err := os.Create("output.html")
	if err != nil {
		log.Fatalln("Output file create err:", err)
	}
	err = tmpl.Execute(f, d)
	if err != nil {
		log.Fatalln("template execute failure:", err)
	}
}
