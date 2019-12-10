package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/alyarctiq/actions-go-acme/randombeer"
	randomsimpson "github.com/alyarctiq/actions-go-acme/simpsons"
)

func thisbeer(w http.ResponseWriter, r *http.Request) {
	joke := randombeer.Randombeerrequest{}
	joke.GetData()
	t, err := template.ParseFiles("templates/randombeer.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, joke)
}
func simpson(w http.ResponseWriter, r *http.Request) {
	joke := randomsimpson.Randomsimpsonrequest{}
	joke.GetData()
	t, err := template.ParseFiles("templates/randombeer.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, joke)
}
func setupRoutes() {
	http.HandleFunc("/beer", thisbeer)
	http.HandleFunc("/simpson", simpson)
	http.Handle("/", http.FileServer(http.Dir("./html")))
}

func main() {
	fmt.Println("Go Web App Started")
	setupRoutes()
	http.ListenAndServe(":3000", nil)

}
// Adding some commenting
