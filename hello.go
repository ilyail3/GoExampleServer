package main

import "fmt"
import "net/http"
import (
	"html/template"
	"strings"
)

type Page struct{
	Name string
	Count int
}

func HomePageHandler(page *Page) func (http.ResponseWriter, *http.Request){
	return func (rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			newName := strings.Trim(r.FormValue("name")," ")

			if len(newName) > 0 {
				page.Name = newName
			}

			http.Redirect(rw, r, "/", http.StatusMovedPermanently)
		} else {
			rw.Header().Add("Content-Type", "text/html")
			rw.WriteHeader(200)

			page.Count ++

			t, _ := template.ParseFiles("hello.html")

			t.Execute(rw, page)
		}
	}
}
func main() {
	fmt.Printf("Hello, world.\n")

	page := Page{Name:"Initial", Count:0}

	http.HandleFunc("/",HomePageHandler(&page))
	http.ListenAndServe(":8080", nil)
}



