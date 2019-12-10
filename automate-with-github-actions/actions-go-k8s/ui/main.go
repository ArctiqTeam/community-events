package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type microservicedata struct {
	Data string `json:"Data"`
	Logo string `json:"Logo"`
	Svc  string `json:"Svc"`
}

func beerme(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://go-web-beer", nil)
	if err != nil {
		//handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(w, "<p align='center'>Doh The Backend Appears Down!! </p>")
		fmt.Fprintf(w, "<p align='center'></p>")
		w.Header().Set("Content-Type", "image/jpeg")
		fmt.Fprintf(w, "<p align='center'><img src='images/home-nobackend.gif' alt='HomeNoooo' style='width:320px;height:320px;'></p>")
		return
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	beer := microservicedata{}
	jsonErr := json.Unmarshal(body, &beer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer resp.Body.Close()

	t, err := template.ParseFiles("templates/beer.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, beer)
}

func simpsons(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://go-web-simpsons", nil)
	if err != nil {
		//handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(w, "<p align='center'>Doh The Backend Appears Down!! </p>")
		fmt.Fprintf(w, "<p align='center'></p>")
		w.Header().Set("Content-Type", "image/jpeg")
		fmt.Fprintf(w, "<p align='center'><img src='images/home-nobackend.gif' alt='HomeNoooo' style='width:320px;height:320px;'></p>")
		return
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	beer := microservicedata{}
	jsonErr := json.Unmarshal(body, &beer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer resp.Body.Close()

	t, err := template.ParseFiles("templates/beer.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, beer)
}
func dadjoke(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://go-web-dadjoke", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(w, "<html>")
		fmt.Fprintf(w, "<h2 align=center>Doh The Backend Appears Down!! </h2>")
		fmt.Fprintf(w, "<p align='center'></p>")
		w.Header().Set("Content-Type", "image/jpeg")
		fmt.Fprintf(w, "<p align='center'><img src='images/home-nobackend.gif' alt='HomeNoooo' style='width:320px;height:320px;'></p>")
		fmt.Fprintf(w, "</html>")
		return
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	beer := microservicedata{}
	jsonErr := json.Unmarshal(body, &beer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer resp.Body.Close()

	t, err := template.ParseFiles("templates/dadjoke.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, beer)
}

func setupRoutes() {
	http.HandleFunc("/beer", beerme)
	http.HandleFunc("/joke", dadjoke)
	http.HandleFunc("/simpsons", simpsons)
	http.Handle("/", http.FileServer(http.Dir("./html")))
}

func main() {
	fmt.Println("Go Web App Started")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
