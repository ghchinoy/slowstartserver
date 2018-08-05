package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const port = 12001

func main() {

	log.Println("starting server...")
	time.Sleep(time.Second * 10)

	r := mux.NewRouter()
	r.Handle("/css/{style}", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.HandleFunc("/", helloHandler)
	http.Handle("/", r)

	log.Printf("Listening on %v ...", port)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), loggedRouter))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	funcs := template.FuncMap{
		"now":   time.Now,
		"fdate": time.Time.Format,
	}
	tmpl := template.Must(template.New("hello.html").Funcs(funcs).ParseFiles("hello.html"))
	//t, err := template.Must(template.Funcs(funcs).ParseFiles("hello.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Println(err)
		return
	}

}
