package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type dadjoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func newjoke(w http.ResponseWriter, r *http.Request) {
	//this function will return a dad joke

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	thisjoke := dadjoke{}
	jsonErr := json.Unmarshal(body, &thisjoke)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer resp.Body.Close()
	joke := thisjoke.Joke

	items := struct {
		Data string
		Logo string
		Svc  string
	}{
		Data: joke,
		Logo: "https://icanhazdadjoke.com/static/smile.svg",
		Svc:  "Dad Joke",
	}

	b, err := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func setupRoutes() {
	http.HandleFunc("/", newjoke)
}

func main() {
	fmt.Println("Go Web App Started")
	setupRoutes()
	http.ListenAndServe(":3000", nil)

}
