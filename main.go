package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type IndexData struct {
	PageTitle  string
	QueryParam string
}

type AddData struct {
	Value1 int
	Value2 int
	Result int
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/", indexHandler)
	log.Println("Starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/add.html")

	if err != nil {
		log.Println(err)
		return
	}

	rawVal1 := r.FormValue("value1")
	rawVal2 := r.FormValue("value2")

	val1, err1 := strconv.Atoi(rawVal1)
	val2, err2 := strconv.Atoi(rawVal2)

	if err1 != nil || err2 != nil {
		val1 = 0
		val2 = 0
	}

	result := val1 + val2

	log.Printf("Adding %d and %d", val1, val2)

	pageData := AddData{
		Value1: val1,
		Value2: val2,
		Result: result,
	}

	tmpl.Execute(w, pageData)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Println(err)
		return
	}

	query := r.URL.Query()
	QueryParam := query.Get("QueryParam")

	pageData := IndexData{
		PageTitle:  "Hello World",
		QueryParam: QueryParam,
	}

	tmpl.Execute(w, pageData)
}
